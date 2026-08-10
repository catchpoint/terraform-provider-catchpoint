package catchpoint

import (
	"context"
	"reflect"
	"strconv"
	"strings"

	"catchpoint-provider/internal/client"
	"catchpoint-provider/internal/expand"
	expandtestmonitor "catchpoint-provider/internal/expand/testmonitor"
	"catchpoint-provider/internal/fields"
	"catchpoint-provider/internal/helpers"
	"catchpoint-provider/internal/logger"
	"catchpoint-provider/internal/models"
	cpresource "catchpoint-provider/internal/models/resource"
	"catchpoint-provider/internal/service"
	cptypes "catchpoint-provider/internal/types"
	cpvalidation "catchpoint-provider/internal/validation"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func (r *TestResource[T]) generatePatchDocuments(plan, state T, resp *resource.UpdateResponse) (jsonPatchDocs []string) {
	logger.DebugBG("Generating patch documents for test update")

	jsonPatchDocs = append(jsonPatchDocs, r.handleBasicFieldUpdates(*plan.GetTestResourceModel(), *state.GetTestResourceModel())...)
	jsonPatchDocs = append(jsonPatchDocs, r.handleComplexSectionUpdates(plan, state, resp)...)

	if patch, ok := r.handleGatewayAddressOrHostUpdate(plan, state); ok {
		jsonPatchDocs = append(jsonPatchDocs, patch)
	}

	if patch, ok := r.handleURLFieldUpdate(plan, state); ok {
		jsonPatchDocs = append(jsonPatchDocs, patch)
	}

	if patch, ok := r.handleDNSFieldUpdate(plan, state); ok {
		jsonPatchDocs = append(jsonPatchDocs, patch...)
	}

	if patch, ok := r.handleSSLFieldUpdate(plan, state); ok {
		jsonPatchDocs = append(jsonPatchDocs, patch)
	}

	if patch, ok := r.handleChromeVersionUpdate(plan, state); ok {
		jsonPatchDocs = append(jsonPatchDocs, patch)
	}

	if patch, ok := r.handleSimulateFieldUpdate(plan, state); ok {
		jsonPatchDocs = append(jsonPatchDocs, patch)
	}

	if patch, ok := r.handleLabelsUpdate(plan, state); ok {
		jsonPatchDocs = append(jsonPatchDocs, patch)
	}

	if patch, ok := r.handleThresholdUpdate(plan, state); ok {
		jsonPatchDocs = append(jsonPatchDocs, patch)
	}

	if patch, ok := r.handleTestScriptUpdate(plan, state, resp); ok {
		jsonPatchDocs = append(jsonPatchDocs, patch)
	}

	return jsonPatchDocs
}

func (r *TestResource[T]) handleBasicFieldUpdates(plan, state cpresource.BaseTestResourceModel) (jsonPatchDocs []string) {
	addPatch := func(value string, section string) {
		testConfigUpdate := models.TestConfigUpdate{
			UpdatedFieldValue: value,
		}
		jsonPatchDocs = append(jsonPatchDocs, service.CreateJSONTestPatchDocument(&testConfigUpdate, section, true))
	}

	// Simple fields
	basicFields := []struct {
		planValue, stateValue any
		getValue              func() string
		section               string
	}{
		{plan.ProductID, state.ProductID, plan.ProductID.String, fields.ProductIDSection},
		{plan.AlertsPaused, state.AlertsPaused, plan.AlertsPaused.String, fields.AlertsPausedSection},
		{plan.Description, state.Description, plan.Description.ValueString, fields.DescriptionSection},
		{plan.EndTime, state.EndTime, plan.EndTime.ValueString, fields.EndTimeSection},
		{plan.EnableTestDataWebhook, state.EnableTestDataWebhook, plan.EnableTestDataWebhook.String, fields.EnableTestDataWebhookSection},
		{plan.Name, state.Name, plan.Name.ValueString, fields.NameSection},
	}

	for _, f := range basicFields {
		if !reflect.ValueOf(f.planValue).MethodByName("Equal").Call([]reflect.Value{reflect.ValueOf(f.stateValue)})[0].Bool() {
			addPatch(f.getValue(), f.section)
		}
	}

	// FolderID special handling
	if !plan.FolderID.Equal(state.FolderID) {
		folderID := plan.FolderID.String()
		if plan.FolderID.IsNull() {
			folderID = cptypes.EmptyString
		}
		addPatch(folderID, fields.FolderIDSection)
	}

	// Monitor special handling
	if !plan.Monitor.Equal(state.Monitor) {
		if monitorID, ok := cptypes.GetMonitorTypeID(plan.Monitor.ValueString()); ok {
			addPatch(strconv.Itoa(monitorID), fields.MonitorSection)
		} else {
			logger.WarnBG("Unable to determine Monitor ID for monitor type '%s', skipping update", plan.Monitor.String())
		}
	}

	// Status special handling
	if !plan.Status.IsNull() && plan.Status.ValueString() != cptypes.EmptyString && !plan.Status.Equal(state.Status) {
		updatedStatus := cpvalidation.GetStatusTypeOrDefault(plan.Status.ValueString())
		addPatch(strconv.Itoa(updatedStatus.ID), fields.StatusSection)
	}

	return jsonPatchDocs
}

func (r *TestResource[T]) handleLabelsUpdate(plan, state T) (out string, hasChanges bool) {
	planLabels := plan.GetTestResourceModel().Label
	stateLabels := state.GetTestResourceModel().Label

	// If labels haven't changed, do nothing
	if reflect.DeepEqual(planLabels, stateLabels) {
		return
	}

	// Labels have changed - convert plan labels to API format and create patch
	if len(planLabels) > 0 {
		var apiLabels []models.LabelsJSON
		for _, labelModel := range planLabels {
			// Convert the values list to a string slice
			labelValues := make([]string, 0, len(labelModel.Values.Elements()))
			for _, value := range labelModel.Values.Elements() {
				if stringValue, ok := value.(types.String); ok {
					labelValues = append(labelValues, stringValue.ValueString())
				}
			}

			// Add the label to the API format using LabelsJSON struct
			apiLabels = append(apiLabels, models.LabelsJSON{
				Name:   labelModel.Key.ValueString(),
				Values: labelValues,
				Color:  helpers.RandomHexString(),
			})
		}

		testConfigUpdate := models.TestConfigUpdate{
			UpdatedLabels:   apiLabels,
			SectionToUpdate: fields.LabelsSection,
		}
		out = service.CreateJSONTestPatchDocument(&testConfigUpdate, testConfigUpdate.SectionToUpdate, false)
		hasChanges = true
	} else {
		// Plan has no labels - remove all labels by setting empty array
		testConfigUpdate := models.TestConfigUpdate{
			UpdatedLabels:   []models.LabelsJSON{},
			SectionToUpdate: fields.LabelsSection,
		}
		out = service.CreateJSONTestPatchDocument(&testConfigUpdate, testConfigUpdate.SectionToUpdate, false)
		hasChanges = true
	}

	return
}

func (r *TestResource[T]) handleTestScriptUpdate(plan, state T, resp *resource.UpdateResponse) (out string, hasChanges bool) {
	planScriptProvider, planOK := any(plan).(cpresource.TestScriptResourceProvider)
	stateScriptProvider, stateOK := any(state).(cpresource.TestScriptResourceProvider)

	if !planOK || !stateOK {
		return
	}

	planScript := planScriptProvider.GetTestScriptResourceModel()
	stateScript := stateScriptProvider.GetTestScriptResourceModel()
	if planScript == nil || stateScript == nil {
		return
	}

	if planScript.Script.Equal(stateScript.Script) && planScript.ScriptType.Equal(stateScript.ScriptType) {
		return
	}

	testConfig := models.TestConfig{}
	expandDiags := expandtestmonitor.ExpandTestConfig(plan, &testConfig)
	resp.Diagnostics.Append(expandDiags...)
	if resp.Diagnostics.HasError() {
		return
	}

	testConfigUpdate := models.TestConfigUpdate{
		UpdatedTestRequestData: service.TestRequestDataFromTestConfig(&testConfig),
		SectionToUpdate:        fields.TestRequestDataSection,
	}
	out = service.CreateJSONTestPatchDocument(&testConfigUpdate, testConfigUpdate.SectionToUpdate, false)
	hasChanges = true
	return
}

func (r *TestResource[T]) handleSSLFieldUpdate(plan, state T) (out string, hasChanges bool) {
	planSSLField, planOK := any(plan).(cpresource.SSLSettingsModelProvider)
	stateSSLField, stateOK := any(state).(cpresource.SSLSettingsModelProvider)

	if !planOK || !stateOK {
		return
	}

	planSSLSettings := planSSLField.GetSSLSettingsModel()
	stateSSLSettings := stateSSLField.GetSSLSettingsModel()

	if planSSLSettings == stateSSLSettings {
		return
	}

	if !planSSLSettings.EnforceCertificateKeyPinning.Equal(stateSSLSettings.EnforceCertificateKeyPinning) && !planSSLSettings.EnforceCertificateKeyPinning.IsNull() {
		testConfigUpdate := models.TestConfigUpdate{
			UpdatedFieldValue: planSSLSettings.EnforceCertificateKeyPinning.String(),
		}
		out = service.CreateJSONTestPatchDocument(&testConfigUpdate, fields.EnforceCertificateKeyPinningSection, true)
		hasChanges = true
	}

	if !planSSLSettings.EnforceCertificatePinning.Equal(stateSSLSettings.EnforceCertificatePinning) && !planSSLSettings.EnforceCertificatePinning.IsNull() {
		testConfigUpdate := models.TestConfigUpdate{
			UpdatedFieldValue: planSSLSettings.EnforceCertificatePinning.String(),
		}
		out = service.CreateJSONTestPatchDocument(&testConfigUpdate, fields.EnforceCertificatePinningSection, true)
		hasChanges = true
	}

	if !planSSLSettings.FileData.Equal(stateSSLSettings.FileData) && !planSSLSettings.FileData.IsNull() {
		testConfigUpdate := models.TestConfigUpdate{
			UpdatedFieldValue: planSSLSettings.FileData.String(),
		}
		out = service.CreateJSONTestPatchDocument(&testConfigUpdate, fields.FileData, true)
		hasChanges = true
	}

	if !planSSLSettings.PassPhrase.Equal(stateSSLSettings.PassPhrase) && !planSSLSettings.PassPhrase.IsNull() {
		testConfigUpdate := models.TestConfigUpdate{
			UpdatedFieldValue: planSSLSettings.PassPhrase.String(),
		}
		out = service.CreateJSONTestPatchDocument(&testConfigUpdate, fields.PassPhrase, true)
		hasChanges = true
	}

	return
}

func (r *TestResource[T]) handleDNSFieldUpdate(plan, state T) (jsonPatchDocs []string, hasChanges bool) {
	planDNSField, planOK := any(plan).(cpresource.DNSTestModelProvider)
	stateDNSField, stateOK := any(state).(cpresource.DNSTestModelProvider)

	if !planOK || !stateOK {
		return
	}

	planDNSServer, planOK := planDNSField.GetDNSServerField()
	stateDNSServer, stateOK := stateDNSField.GetDNSServerField()

	if planOK && stateOK && !planDNSServer.Equal(stateDNSServer) {
		testConfigUpdate := models.TestConfigUpdate{
			UpdatedFieldValue: planDNSServer.ValueString(),
		}
		jsonPatchDocs = append(jsonPatchDocs, service.CreateJSONTestPatchDocument(&testConfigUpdate, fields.DNSServerSection, true))
		hasChanges = true
	}

	planQueryType, planOK := planDNSField.GetQueryTypeField()
	stateQueryType, stateOK := stateDNSField.GetQueryTypeField()

	if planOK && stateOK && !planQueryType.Equal(stateQueryType) {
		testConfigUpdate := models.TestConfigUpdate{
			UpdatedFieldValue: planQueryType.ValueString(),
		}
		jsonPatchDocs = append(jsonPatchDocs, service.CreateJSONTestPatchDocument(&testConfigUpdate, fields.DNSQueryTypeSection, true))
		hasChanges = true
	}

	return
}

func (r *TestResource[T]) handleGatewayAddressOrHostUpdate(plan, state T) (out string, hasChanges bool) {
	planGatewayAddressOrHost, planOK := any(plan).(cpresource.GatewayAddressOrHostModelProvider)
	stateGatewayAddressOrHost, stateOK := any(state).(cpresource.GatewayAddressOrHostModelProvider)

	if !planOK || !stateOK {
		return
	}

	planGatewayAddressOrHostModel := planGatewayAddressOrHost.GetGatewayAddressOrHostModel()
	stateGatewayAddressOrHostModel := stateGatewayAddressOrHost.GetGatewayAddressOrHostModel()

	if planGatewayAddressOrHostModel.GatewayAddressOrHost.Equal(stateGatewayAddressOrHostModel.GatewayAddressOrHost) {
		return
	}

	testConfigUpdate := models.TestConfigUpdate{
		UpdatedFieldValue: planGatewayAddressOrHostModel.GatewayAddressOrHost.ValueString(),
	}
	out = service.CreateJSONTestPatchDocument(&testConfigUpdate, fields.GatewayAddressOrHostSection, true)
	hasChanges = true

	return
}

func (r *TestResource[T]) handleChromeVersionUpdate(plan, state T) (jsonPatchDoc string, hasChanges bool) {
	planApplicationVersionProvider, planOK := any(plan).(cpresource.ChromeVersionProvider)
	stateApplicationVersionProvider, stateOK := any(state).(cpresource.ChromeVersionProvider)

	if !planOK || !stateOK {
		return
	}

	applicationVersion, ok := planApplicationVersionProvider.GetChromeVersionField()
	if !ok || applicationVersion.IsNull() || applicationVersion.ValueString() == cptypes.EmptyString {
		return
	}

	stateApplicationVersion, ok := stateApplicationVersionProvider.GetChromeVersionField()
	if !ok || stateApplicationVersion.IsNull() || stateApplicationVersion.ValueString() == cptypes.EmptyString {
		return
	}

	if applicationVersion.Equal(stateApplicationVersion) {
		return
	}

	if chromeVersionID, ok := cptypes.GetChromeVersionID(applicationVersion.ValueString()); ok {
		update := models.TestConfigUpdate{}
		model := models.ChromeMonitorVersionStructJSON{
			ApplicationVersionType: &models.GenericIDNameOmitEmptyJSON{
				ID:   &chromeVersionID,
				Name: applicationVersion.ValueStringPointer(),
			},
		}

		if chromeVersionID == 3 {
			if chromeApplicationVersionID, ok := cptypes.GetChromeApplicationVersionID(applicationVersion.ValueString()); ok {
				model.ApplicationVersionID = &chromeApplicationVersionID
				model.ApplicationVersionType.Name = types.StringValue("Specific").ValueStringPointer()
			} else {
				// Can't be found in the specific version mapping (e.g. "53", "59", "63", etc)
				logger.WarnBG("Unable to determine Chrome application version ID for version '%s', skipping update", applicationVersion.ValueString())
				return
			}
		}

		update.SectionToUpdate = fields.ChromeMonitorVersionSection
		update.UpdatedChromeVersionSection = model
		jsonPatchDoc = service.CreateJSONTestPatchDocument(&update, fields.ChromeMonitorVersionSection, false)
		hasChanges = true
	} else {
		// Can't be found in the top-level map (e.g. "stable", "preview", "53", "59", etc)
		logger.WarnBG("Unable to determine Chrome version ID for version '%s', skipping update", applicationVersion.ValueString())
	}

	return
}

func (r *TestResource[T]) handleSimulateFieldUpdate(plan, state T) (out string, hasChanges bool) {
	planSimulateField, planOK := any(plan).(cpresource.SimulateProvider)
	stateSimulateField, stateOK := any(state).(cpresource.SimulateProvider)

	if !planOK || !stateOK {
		return
	}

	planSimulate, ok := planSimulateField.GetSimulateField()
	if !ok || planSimulate.IsNull() || planSimulate.ValueString() == cptypes.EmptyString {
		return
	}

	stateSimulate, ok := stateSimulateField.GetSimulateField()
	if !ok || stateSimulate.IsNull() || stateSimulate.ValueString() == cptypes.EmptyString {
		return
	}

	if userAgentID, ok := cptypes.GetUserAgentTypeID(planSimulate.ValueString()); ok {
		testConfigUpdate := models.TestConfigUpdate{
			UpdatedFieldValue: strconv.Itoa(userAgentID),
		}
		out = service.CreateJSONTestPatchDocument(&testConfigUpdate, fields.SimulateSection, true)
		hasChanges = true
	} else {
		// Note: only warning because this should be caught in schema validation - i.e. it should never get here.
		logger.WarnBG("Unable to determine User Agent ID for simulate type '%s', skipping update", planSimulate.String())
	}

	return
}

func (r *TestResource[T]) handleComplexSectionUpdates(plan, state T, resp *resource.UpdateResponse) (jsonPatchDocs []string) {
	if patch, ok := r.handleAdvancedSettingsUpdate(plan, state, resp); ok {
		jsonPatchDocs = append(jsonPatchDocs, patch)
	}
	if patch, ok := r.handleRequestSettingsUpdate(plan, state, resp); ok {
		jsonPatchDocs = append(jsonPatchDocs, patch)
	}
	if patch, ok := r.handleAlertSettingsUpdate(plan, state, resp); ok {
		jsonPatchDocs = append(jsonPatchDocs, patch)
	}
	if patch, ok := r.handleScheduleSettingsUpdate(plan, state, resp); ok {
		jsonPatchDocs = append(jsonPatchDocs, patch)
	}
	if patch, ok := r.handleInsightsUpdate(plan, state); ok {
		jsonPatchDocs = append(jsonPatchDocs, patch...)
	}

	return jsonPatchDocs
}

func (r *TestResource[T]) handleAdvancedSettingsUpdate(plan, state T, resp *resource.UpdateResponse) (out string, hasChanges bool) {
	planAdvancedSettingsProvider, planOK := any(plan).(cpresource.AdvancedSettingsModelProvider)
	stateAdvancedSettingsProvider, stateOK := any(state).(cpresource.AdvancedSettingsModelProvider)

	if !planOK || !stateOK {
		return
	}

	planAdvancedSettings := planAdvancedSettingsProvider.GetAdvancedSettingsModel()
	stateAdvancedSettings := stateAdvancedSettingsProvider.GetAdvancedSettingsModel()
	if planAdvancedSettings.AdvancedSettings.IsNull() && stateAdvancedSettings.AdvancedSettings.IsNull() {
		return
	}

	if !planAdvancedSettings.AdvancedSettings.Equal(stateAdvancedSettings.AdvancedSettings) {
		testConfig := models.TestConfig{}
		expandDiags := expand.ExpandAdvancedSettingsConfig(planAdvancedSettings.AdvancedSettings, &testConfig.CommonConfig.AdvancedSettingsConfig)
		resp.Diagnostics.Append(expandDiags...)
		if resp.Diagnostics.HasError() {
			logger.ErrorBG("Advanced settings update failed: %v", resp.Diagnostics)
			return
		}

		testConfigUpdate := models.TestConfigUpdate{
			UpdatedAdvancedSettingsSection: service.AdvancedSettingFromTestConfig(&testConfig),
			SectionToUpdate:                fields.AdvancedSettingsSection,
		}
		out = service.CreateJSONTestPatchDocument(&testConfigUpdate, testConfigUpdate.SectionToUpdate, false)
		hasChanges = true
	}
	return
}

func (r *TestResource[T]) handleRequestSettingsUpdate(plan, state T, resp *resource.UpdateResponse) (out string, hasChanges bool) {
	planRequestSettingsProvider, planOK := any(plan).(cpresource.RequestSettingsModelProvider)
	stateRequestSettingsProvider, stateOK := any(state).(cpresource.RequestSettingsModelProvider)

	if !planOK || !stateOK {
		return
	}

	planRequestSettings := planRequestSettingsProvider.GetRequestSettingsModel()
	stateRequestSettings := stateRequestSettingsProvider.GetRequestSettingsModel()
	if planRequestSettings.RequestSettings.IsNull() && stateRequestSettings.RequestSettings.IsNull() {
		return
	}

	if !planRequestSettings.RequestSettings.Equal(stateRequestSettings.RequestSettings) {
		testConfig := models.TestConfig{}
		expandDiags := expand.ExpandRequestSettingsConfig(
			planRequestSettings.RequestSettings,
			&testConfig.CommonConfig.RequestSettingsConfig,
		)
		resp.Diagnostics.Append(expandDiags...)
		if resp.Diagnostics.HasError() {
			return
		}

		testConfigUpdate := models.TestConfigUpdate{
			UpdatedRequestSettingsSection: service.RequestSettingFromTestConfig(&testConfig),
			SectionToUpdate:               fields.RequestSettingsSection,
		}
		out = service.CreateJSONTestPatchDocument(&testConfigUpdate, testConfigUpdate.SectionToUpdate, false)
		hasChanges = true
	}
	return
}

func (r *TestResource[T]) handleAlertSettingsUpdate(plan, state T, resp *resource.UpdateResponse) (out string, hasChanges bool) {
	planAlertSettingsProvider, planOK := any(plan).(cpresource.AlertSettingsModelProvider)
	stateAlertSettingsProvider, stateOK := any(state).(cpresource.AlertSettingsModelProvider)

	if !planOK || !stateOK {
		return
	}

	planAlertSettings := planAlertSettingsProvider.GetAlertSettingsModel()
	stateAlertSettings := stateAlertSettingsProvider.GetAlertSettingsModel()
	if planAlertSettings.IsNull() && stateAlertSettings.IsNull() {
		return
	}

	if planAlertSettings != stateAlertSettings {
		testConfig := models.TestConfig{}
		expandDiags := expand.ExpandAlertSettingsConfig(planAlertSettings, &testConfig.CommonConfig.AlertSettingsConfig)
		resp.Diagnostics.Append(expandDiags...)
		if resp.Diagnostics.HasError() {
			return
		}

		testConfigUpdate := models.TestConfigUpdate{
			UpdatedAlertSettingsSection: service.AlertSettingsFromTestConfig(&testConfig),
			SectionToUpdate:             fields.AlertGroupSection,
		}
		out = service.CreateJSONTestPatchDocument(&testConfigUpdate, testConfigUpdate.SectionToUpdate, false)
		hasChanges = true
	}
	return
}

func (r *TestResource[T]) handleScheduleSettingsUpdate(plan, state T, resp *resource.UpdateResponse) (out string, hasChanges bool) {
	planScheduleSettingsProvider, planOK := any(plan).(cpresource.ScheduleSettingsModelProvider)
	stateScheduleSettingsProvider, stateOK := any(state).(cpresource.ScheduleSettingsModelProvider)

	if !planOK || !stateOK {
		return
	}

	planScheduleSettings := planScheduleSettingsProvider.GetScheduleSettingsModel()
	stateScheduleSettings := stateScheduleSettingsProvider.GetScheduleSettingsModel()
	if planScheduleSettings.ScheduleSettings.IsNull() && stateScheduleSettings.ScheduleSettings.IsNull() {
		return
	}

	if !planScheduleSettings.ScheduleSettings.Equal(stateScheduleSettings.ScheduleSettings) {
		testConfig := models.TestConfig{}
		expandDiags := expand.ExpandScheduleSettingsConfig(planScheduleSettings.ScheduleSettings, &testConfig.CommonConfig.ScheduleSettingsConfig)
		resp.Diagnostics.Append(expandDiags...)
		if resp.Diagnostics.HasError() {
			return
		}

		testConfigUpdate := models.TestConfigUpdate{
			UpdatedScheduleSettingsSection: service.ScheduleSettingsFromTestConfig(&testConfig),
			SectionToUpdate:                fields.ScheduleSettingsSection,
		}
		out = service.CreateJSONTestPatchDocument(&testConfigUpdate, testConfigUpdate.SectionToUpdate, false)
		hasChanges = true
	}
	return
}

// This function is a bit of an oddity because the backend wants a more advanced schema than we currently use
// to set the values to begin with. Thus, instead of patching the entire insight section, we individually
// patch the indicators and tracepoints separately.
func (r *TestResource[T]) handleInsightsUpdate(plan, state T) (response []string, hasChanges bool) {
	planInsightSettingsProvider, planOK := any(plan).(cpresource.InsightSettingsModelProvider)
	stateInsightSettingsProvider, stateOK := any(state).(cpresource.InsightSettingsModelProvider)

	if !planOK || !stateOK {
		return
	}

	planInsightSettings := planInsightSettingsProvider.GetInsightSettingsModel()
	stateInsightSettings := stateInsightSettingsProvider.GetInsightSettingsModel()
	if planInsightSettings.Insights.IsNull() && stateInsightSettings.Insights.IsNull() {
		return
	}

	// Extract plan attributes
	planIndicators, planTracepoints := extractInsightsAttributes(planInsightSettings.Insights)

	// Extract state attributes
	stateIndicators, stateTracepoints := extractInsightsAttributes(stateInsightSettings.Insights)

	// Compare individual components
	indicatorsChanged := !planIndicators.Equal(stateIndicators)
	tracepointsChanged := !planTracepoints.Equal(stateTracepoints)

	// Only update if there are actual changes
	if !indicatorsChanged && !tracepointsChanged {
		return
	}

	if indicatorsChanged {
		indicatorValues := listToIntSlice(planIndicators)
		indicatorsConfigUpdate := models.TestConfigUpdate{
			SectionToUpdate: fields.InsightDataSection + fields.IndicatorsSection,
		}

		for _, val := range indicatorValues {
			indicatorsConfigUpdate.UpdatedInsightSettingsSection = append(indicatorsConfigUpdate.UpdatedInsightSettingsSection, map[string]int{"id": int(val)})
		}

		response = append(response, service.CreateJSONTestPatchDocument(&indicatorsConfigUpdate, indicatorsConfigUpdate.SectionToUpdate, false))
	}
	if tracepointsChanged {
		tracepointValues := listToIntSlice(planTracepoints)
		tracepointsConfigUpdate := models.TestConfigUpdate{
			SectionToUpdate: fields.InsightDataSection + fields.TracepointsSection,
		}
		for _, val := range tracepointValues {
			tracepointsConfigUpdate.UpdatedInsightSettingsSection = append(tracepointsConfigUpdate.UpdatedInsightSettingsSection, map[string]int{"id": int(val)})
		}
		response = append(response, service.CreateJSONTestPatchDocument(&tracepointsConfigUpdate, tracepointsConfigUpdate.SectionToUpdate, false))
	}

	settingType := 0
	if !planInsightSettings.IsNull() {
		attrs := planInsightSettings.Insights.Attributes()
		settingType = cpvalidation.GetGenericSettingTypeOrDefault(attrs[fields.InsightSettingType].(types.String).ValueString()).ID
	}

	insightsConfigUpdate := models.TestConfigUpdate{
		SectionToUpdate:   fields.InsightDataSection + fields.InsightSettingTypeSection,
		UpdatedFieldValue: strconv.Itoa(settingType),
	}

	response = append(response, service.CreateJSONTestPatchDocument(&insightsConfigUpdate, insightsConfigUpdate.SectionToUpdate, true))
	hasChanges = true

	return
}

func (r *TestResource[T]) handleThresholdUpdate(plan, state T) (out string, hasChanges bool) {
	planThresholdProvider, planOK := any(plan).(cpresource.ThresholdModelProvider)
	stateThresholdProvider, stateOK := any(state).(cpresource.ThresholdModelProvider)

	if !planOK || !stateOK {
		return
	}

	planThresholds := planThresholdProvider.GetTestThresholdModel()
	stateThresholds := stateThresholdProvider.GetTestThresholdModel()

	if planThresholds.IsNull() && stateThresholds.IsNull() {
		return
	}

	if planThresholds != stateThresholds {
		testThresholds := models.TestThresholdsJSON{}

		if !planThresholds.IsNull() {
			if !planThresholds.AvailabilityCritical.IsNull() {
				testThresholds.AvailabilityApdexThresholdCritical = planThresholds.AvailabilityCritical.ValueFloat64Pointer()
			}

			if !planThresholds.AvailabilityWarning.IsNull() {
				testThresholds.AvailabilityApdexThresholdWarning = planThresholds.AvailabilityWarning.ValueFloat64Pointer()
			}

			if !planThresholds.TestTimeCritical.IsNull() {
				testThresholds.TestTimeApdexThresholdCritical = planThresholds.TestTimeCritical.ValueFloat64Pointer()
			}

			if !planThresholds.TestTimeWarning.IsNull() {
				testThresholds.TestTimeApdexThresholdWarning = planThresholds.TestTimeWarning.ValueFloat64Pointer()
			}
		}

		testConfigUpdate := models.TestConfigUpdate{
			UpdatedTestThresholds: testThresholds,
			SectionToUpdate:       fields.ThresholdRestModelSection,
		}
		out = service.CreateJSONTestPatchDocument(&testConfigUpdate, testConfigUpdate.SectionToUpdate, false)
		hasChanges = true
	}

	return
}

func (r *TestResource[T]) handleURLFieldUpdate(plan, state T) (out string, hasChanges bool) {
	// Use a type assertion to check if the model can provide a URL.
	planURLProvider, planOK := any(plan).(cpresource.URLProvider)
	stateURLProvider, stateOK := any(state).(cpresource.URLProvider)

	// If the model type doesn't support a URL field, do nothing.
	if !planOK || !stateOK {
		return
	}

	planURL, planHasURL := planURLProvider.GetURLField()
	stateURL, stateHasURL := stateURLProvider.GetURLField()

	// If the model has a URL field and it has changed, create a patch.
	if planHasURL && stateHasURL && !planURL.Equal(stateURL) {
		testConfigUpdate := models.TestConfigUpdate{
			UpdatedFieldValue: planURL.ValueString(),
		}
		// The backend expects all URL-like fields to be in the URL section.
		out = service.CreateJSONTestPatchDocument(&testConfigUpdate, fields.URLSection, true)
		hasChanges = true
	}

	return
}

func (r *TestResource[T]) executeUpdate(ctx context.Context, testID int64, jsonPatchDocs []string, plan T, resp *resource.UpdateResponse, req resource.UpdateRequest) {
	jsonPatchDoc := "[" + strings.Join(jsonPatchDocs, ",") + "]"

	if r.providerConfig.LogJSON {
		logger.Debug(ctx, "Updating test %d with JSON PATCH: %v", testID, jsonPatchDoc)
	}

	respBody, respStatus, completed, err := client.UpdateTest(r.providerConfig.APIToken, testID, jsonPatchDoc)
	if err != nil {
		logger.Debug(ctx, "Error updating Test: %v. Body: %s, Status: %s", err, respBody, respStatus)
		resp.Diagnostics.AddError("Error updating Test", err.Error())
		return
	}

	if !completed {
		resp.Diagnostics.AddWarning("Update may not be complete", "The API indicated the update may not have completed successfully")
	}

	r.refreshTestState(ctx, testID, plan, resp, req)
}

func (r *TestResource[T]) refreshTestState(ctx context.Context, testID int64, plan T, resp *resource.UpdateResponse, req resource.UpdateRequest) {
	var originalPlan T
	req.Config.Get(ctx, &originalPlan)

	// Transform the API response to Terraform state
	state, stateDiags := r.getAPITestAsState(ctx, testID, r.providerConfig.APIToken, originalPlan)
	resp.Diagnostics.Append(stateDiags...)
	if resp.Diagnostics.HasError() {
		return
	}

	finalState := r.mergeTestPlanWithState(ctx, plan, state)
	finalState.SetID(types.Int64Value(testID))

	// Set the state
	resp.Diagnostics.Append(resp.State.Set(ctx, &finalState)...)
	if !resp.Diagnostics.HasError() {
		logger.Debug(ctx, "State set successfully")
	} else {
		logger.Debug(ctx, "State set failed: %v", resp.Diagnostics.Errors())
	}
}
