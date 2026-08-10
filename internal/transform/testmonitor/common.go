package testmonitor

import (
	"context"
	"fmt"

	"catchpoint-provider/internal/helpers"
	"catchpoint-provider/internal/models"
	"catchpoint-provider/internal/models/resource"
	cpschema "catchpoint-provider/internal/schema"
	"catchpoint-provider/internal/transform"
	cptypes "catchpoint-provider/internal/types"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func JSONToTerraformTest[T resource.TestResourceModelProvider](ctx context.Context, model T, test *models.TestJSON, config T) (diags diag.Diagnostics) {
	diags.Append(transformCommonFields(ctx, model, test)...)
	if diags.HasError() {
		return
	}

	diags.Append(transformComplexFields(model, ctx, test, config)...)
	if diags.HasError() {
		return
	}

	diags.Append(handleScriptData(model, test, config)...)
	if diags.HasError() {
		return
	}

	diags.Append(handleGatewayAddressOrHost(model, test)...)
	if diags.HasError() {
		return
	}

	diags.Append(handleChromeVersion(model, test)...)
	if diags.HasError() {
		return
	}

	diags.Append(handleSimulateField(model, test)...)
	if diags.HasError() {
		return
	}

	diags.Append(setDNSDataForTest(model, test)...)
	if diags.HasError() {
		return
	}

	diags.Append(setSSLDataForTest(model, test)...)
	if diags.HasError() {
		return
	}

	diags.Append(setThresholds(model, test, config)...)

	return
}

// handleScriptData populates script data if the model supports it.

func handleScriptData[T any](model any, test *models.TestJSON, config T) (diags diag.Diagnostics) {
	provider, ok := model.(resource.TestScriptResourceProvider)
	if !ok {
		return
	}

	if test.TestRequestData != nil {
		diags.Append(setTestRequestDataForTest(provider.GetTestScriptResourceModel(), test)...)
	}

	configProvider, ok := any(config).(resource.TestScriptResourceProvider)
	if !ok || configProvider.GetTestScriptResourceModel() == nil {
		return
	}

	configuredScript := configProvider.GetTestScriptResourceModel().Script
	if !configuredScript.IsNull() && !configuredScript.IsUnknown() {
		provider.GetTestScriptResourceModel().Script = configuredScript
	}
	return
}

// handleGatewayAddressOrHost populates gateway address or host if the model supports it.
func handleGatewayAddressOrHost(model any, test *models.TestJSON) (diags diag.Diagnostics) {
	if provider, ok := model.(resource.GatewayAddressOrHostModelProvider); ok {
		diags.Append(setGatewayAddressOrHostModel(provider, test)...)
	}
	return
}

func handleSimulateField(model any, test *models.TestJSON) (diags diag.Diagnostics) {
	if provider, ok := model.(resource.SimulateProvider); ok {
		if test.UserAgentType == nil || test.UserAgentType.ID == nil {
			provider.SetSimulateField(types.StringValue(cptypes.EmptyString))
		} else if simulate, ok := cptypes.GetUserAgentTypeName(*test.UserAgentType.ID); ok {
			provider.SetSimulateField(types.StringValue(simulate))
		} else {
			diags.AddError("Invalid Simulate Value", fmt.Sprintf("Unknown user agent type ID: %d", *test.UserAgentType.ID))
		}
	}
	return
}

// handleChromeVersion populates Chrome version if the model supports it.
func handleChromeVersion(model any, test *models.TestJSON) (diags diag.Diagnostics) {
	if provider, ok := model.(resource.ChromeVersionProvider); ok {
		if test.ChromeMonitorVersion == nil || test.ChromeMonitorVersion.ApplicationVersionType == nil {
			provider.SetChromeVersionField(types.StringValue(cptypes.EmptyString))
		} else {
			version, err := getChromeVersion(test)
			if err != nil {
				diags.AddError("Invalid Chrome Version", err.Error())
			}
			provider.SetChromeVersionField(types.StringValue(version))
		}
	}
	return
}

func getChromeVersion(test *models.TestJSON) (version string, err error) {
	switch *test.ChromeMonitorVersion.ApplicationVersionType.ID {
	case 1:
		version = cptypes.Stable
	case 2:
		version = cptypes.Preview
	case 3:
		if test.ChromeMonitorVersion.ApplicationVersionID == nil {
			err = fmt.Errorf("chrome version ID is nil for ApplicationVersionType 3")
			return
		}
		versionID := *test.ChromeMonitorVersion.ApplicationVersionID
		if versionName, found := cptypes.GetChromeApplicationVersionString(versionID); found {
			version = versionName
		} else {
			err = fmt.Errorf("unknown Chrome version ID: %d", versionID)
		}
	default:
		if test.ChromeMonitorVersion.ApplicationVersionID != nil {
			err = fmt.Errorf("unknown Chrome version ID: %d", *test.ChromeMonitorVersion.ApplicationVersionID)
		} else {
			err = fmt.Errorf("chrome version ID is nil for unknown ApplicationVersionType")
		}
	}
	return
}

func setGatewayAddressOrHostModel(model resource.GatewayAddressOrHostModelProvider, test *models.TestJSON) (diags diag.Diagnostics) {
	var gateway resource.GatewayAddressOrHostModel

	if test.GatewayAddressOrHost == nil || *test.GatewayAddressOrHost == cptypes.EmptyString {
		gateway.GatewayAddressOrHost = types.StringValue(cptypes.EmptyString)
	} else {
		gateway.GatewayAddressOrHost = types.StringValue(*test.GatewayAddressOrHost)
	}

	model.SetGatewayAddressOrHostModel(&gateway)

	return
}

func setDNSDataForTest(model resource.TestResourceModelProvider, test *models.TestJSON) (diags diag.Diagnostics) {
	if provider, ok := any(model).(resource.DNSTestModelProvider); ok {
		if test.DNSQueryType != nil {
			if queryTypeName, found := cptypes.GetDNSQueryTypeName(*test.DNSQueryType.ID); found {
				provider.SetQueryTypeField(types.StringValue(queryTypeName))
			} else {
				diags.AddError("Invalid DNS Query Type", fmt.Sprintf("Unknown DNS query type ID: %d", test.DNSQueryType.ID))
				return diags
			}
		} else {
			diags.AddError("Missing DNS Query Type", "The DNS query type is missing from the API response.")
			return diags
		}
		if test.DNSServer != nil {
			provider.SetDNSServerField(types.StringValue(*test.DNSServer))
		}
	}
	return
}

func setSSLDataForTest(model resource.TestResourceModelProvider, test *models.TestJSON) (diags diag.Diagnostics) {
	if provider, ok := any(model).(resource.SSLSettingsModelProvider); ok {
		var ssl resource.SSLSettingsModel
		if test.EnforceCertificateKeyPinning != nil {
			ssl.EnforceCertificateKeyPinning = types.BoolValue(*test.EnforceCertificateKeyPinning)
		} else {
			ssl.EnforceCertificateKeyPinning = types.BoolValue(false)
		}
		if test.EnforceCertificatePinning != nil {
			ssl.EnforceCertificatePinning = types.BoolValue(*test.EnforceCertificatePinning)
		} else {
			ssl.EnforceCertificatePinning = types.BoolValue(false)
		}
		if test.FileData != nil {
			ssl.FileData = types.StringValue(*test.FileData)
		}
		if test.PassPhrase != nil {
			ssl.PassPhrase = types.StringValue(*test.PassPhrase)
		}
		provider.SetSSLSettingsModel(&ssl)
	}
	return
}

func setTestRequestDataForTest(resource *resource.TestScriptResourceModel, test *models.TestJSON) (diags diag.Diagnostics) {
	// If the TestRequestData is not nil, this shouldn't be nil either.
	script := types.StringValue(helpers.NormalizeScript(*test.TestRequestData.RequestData))
	resource.Script = script

	if test.TestRequestData.TransactionScriptType != nil {
		scriptType, ok := cptypes.GetAPIScriptTypeName(*test.TestRequestData.TransactionScriptType.ID)
		if !ok {
			diags.AddError("Invalid Script Type", fmt.Sprintf("Unknown script type ID: %d", *test.TestRequestData.TransactionScriptType.ID))
			return diags
		}
		scriptTypeString := types.StringValue(scriptType)
		resource.ScriptType = scriptTypeString
	}

	return
}

func transformComplexFields[T resource.TestResourceModelProvider](model T, ctx context.Context, test *models.TestJSON, config T) (diags diag.Diagnostics) {
	diags.Append(handleAlertSettings(model, test, config)...)
	diags.Append(handleAdvancedSettings(model, test, config)...)
	diags.Append(handleRequestSettings(model, test, config)...)
	diags.Append(handleScheduleSettings(model, test, config)...)
	diags.Append(handleInsightSettings(model, ctx, test, config)...)
	return
}

func handleAlertSettings[T resource.TestResourceModelProvider](model T, test *models.TestJSON, config T) (diags diag.Diagnostics) {
	if provider, ok := any(model).(resource.AlertSettingsModelProvider); ok {
		configProvider, _ := any(config).(resource.AlertSettingsModelProvider)
		if configProvider != nil && !configProvider.GetAlertSettingsModel().IsNull() && test.AlertGroup != nil {
			settings, d := transform.JSONToTerraformAlertGroup(test.AlertGroup, configProvider.GetAlertSettingsModel())
			diags.Append(d...)
			provider.SetAlertSettingsModel(settings)
		} else {
			alertSettingsModel := resource.AlertSettingsModel{
				AlertRule:        []resource.AlertRuleModel{},
				AlertSettingType: types.StringValue(cptypes.Inherit),
			}
			provider.SetAlertSettingsModel(&alertSettingsModel)
		}
	}
	return
}

func handleAdvancedSettings[T resource.TestResourceModelProvider](model T, test *models.TestJSON, config T) (diags diag.Diagnostics) {
	if provider, ok := any(model).(resource.AdvancedSettingsModelProvider); ok {
		configProvider, _ := any(config).(resource.AdvancedSettingsModelProvider)
		if configProvider != nil &&
			!configProvider.GetAdvancedSettingsModel().IsNull() &&
			test.AdvancedSettings != nil {
			settings, d := transform.JSONToTerraformAdvancedSettingsForTest(test.AdvancedSettings, model.GetTestType())
			diags.Append(d...)
			provider.SetAdvancedSettingsModel(&settings)
		} else {
			obj := types.ObjectNull(cpschema.GetAdvancedSettingsAttributeTypesForTest(model.GetTestType()))
			advancedSettingsModel := resource.AdvancedSettingsModel{
				AdvancedSettings: obj,
			}
			provider.SetAdvancedSettingsModel(&advancedSettingsModel)
		}
	}
	return
}

func handleRequestSettings[T resource.TestResourceModelProvider](model T, test *models.TestJSON, config T) (diags diag.Diagnostics) {
	if provider, ok := any(model).(resource.RequestSettingsModelProvider); ok {
		configProvider, _ := any(config).(resource.RequestSettingsModelProvider)
		if configProvider != nil &&
			!configProvider.GetRequestSettingsModel().IsNull() &&
			test.RequestSettings != nil {
			settings, d := transform.JSONToTerraformRequestSettings(test.RequestSettings)
			diags.Append(d...)
			provider.SetRequestSettingsModel(&settings)
		} else {
			attrTypes := cpschema.GetRequestSettingsAttributeTypes()
			obj := types.ObjectNull(attrTypes)
			requestSettingsModel := resource.RequestSettingsModel{
				RequestSettings: obj,
			}
			provider.GetRequestSettingsModel().RequestSettings = requestSettingsModel.RequestSettings
		}
	}
	return
}

func handleScheduleSettings[T resource.TestResourceModelProvider](model T, test *models.TestJSON, config T) (diags diag.Diagnostics) {
	if provider, ok := any(model).(resource.ScheduleSettingsModelProvider); ok {
		configProvider, _ := any(config).(resource.ScheduleSettingsModelProvider)
		if configProvider != nil &&
			!configProvider.GetScheduleSettingsModel().IsNull() &&
			test.ScheduleSettings != nil {
			settings, d := transform.JSONToTerraformScheduleSettings(test.ScheduleSettings)
			diags.Append(d...)
			provider.SetScheduleSettingsModel(&settings)
		} else {
			obj := types.ObjectNull(cpschema.GetScheduleSettingsAttributeTypes())
			scheduleSettingsModel := resource.ScheduleSettingsModel{
				ScheduleSettings: obj,
			}
			provider.SetScheduleSettingsModel(&scheduleSettingsModel)
		}
	}
	return
}

func handleInsightSettings[T resource.TestResourceModelProvider](model T, ctx context.Context, test *models.TestJSON, config T) (diags diag.Diagnostics) {
	if provider, ok := any(model).(resource.InsightSettingsModelProvider); ok {
		configProvider, _ := any(config).(resource.InsightSettingsModelProvider)
		if configProvider != nil &&
			!configProvider.GetInsightSettingsModel().IsNull() &&
			test.InsightData != nil {
			settings, d := transform.JSONToTerraformInsights(ctx, test.InsightData)
			diags.Append(d...)
			provider.SetInsightSettingsModel(&settings)
		} else {
			obj := types.ObjectNull(cpschema.GetInsightsAttributeTypes())
			insightSettingsModel := resource.InsightSettingsModel{
				Insights: obj,
			}
			provider.SetInsightSettingsModel(&insightSettingsModel)
		}
	}
	return
}

// transformCommonFields populates the common fields of any test resource model from the API response.
// It uses the TestResourceModelInterface to access the base model and the URLProvider interface
// to generically handle URL-like fields (e.g., test_url, prefix, etc.).
func transformCommonFields(ctx context.Context, model resource.TestResourceModelProvider, test *models.TestJSON) (diags diag.Diagnostics) {
	if test == nil {
		panic("transformCommonFields called with nil TestJSON - this indicates a bug in the calling code")
	}

	testResourceModel := model.GetTestResourceModel()

	// Set basic fields
	setBasicFields(testResourceModel, test)

	// Set status and monitor with validation
	diags.Append(setStatusAndMonitor(testResourceModel, test)...)
	if diags.HasError() {
		return
	}

	// Set optional fields
	setOptionalFields(testResourceModel, test)

	// Set labels if present
	setLabels(testResourceModel, test)

	// Set URL if provider supports it
	setURLField(model, test)

	return
}

// setBasicFields sets the basic required fields
func setBasicFields(testResourceModel *resource.BaseTestResourceModel, test *models.TestJSON) {
	testResourceModel.ID = types.Int64Value(int64(test.ID))
	testResourceModel.AlertsPaused = types.BoolValue(test.AlertsPaused)
	testResourceModel.DivisionID = types.Int64Value(int64(test.DivisionID))
	testResourceModel.EnableTestDataWebhook = types.BoolValue(test.EnableTestDataWebhook)
	testResourceModel.ProductID = types.Int64Value(int64(test.ProductID))
	testResourceModel.StartTime = types.StringValue(test.StartTime)
	testResourceModel.Name = types.StringValue(test.Name)
}

// setStatusAndMonitor sets status and monitor fields with validation
func setStatusAndMonitor(testResourceModel *resource.BaseTestResourceModel, test *models.TestJSON) (diags diag.Diagnostics) {
	if statusName, ok := cptypes.GetStatusTypeName(test.Status.ID); ok {
		testResourceModel.Status = types.StringValue(statusName)
	} else {
		diags.AddError("Invalid Status Type", fmt.Sprintf("Unknown status type ID: %d", test.Status.ID))
	}

	if monitorName, ok := cptypes.GetMonitorTypeName(test.Monitor.ID); ok {
		testResourceModel.Monitor = types.StringValue(monitorName)
	} else {
		diags.AddError("Invalid Monitor Type", fmt.Sprintf("Unknown monitor type ID: %d", test.Monitor.ID))
	}

	return
}

// setOptionalFields sets optional fields if they exist
func setOptionalFields(testResourceModel *resource.BaseTestResourceModel, test *models.TestJSON) {
	if test.FolderID != nil {
		testResourceModel.FolderID = types.Int64Value(int64(*test.FolderID))
	}

	if test.Description != nil {
		testResourceModel.Description = types.StringValue(*test.Description)
	}

	if test.EndTime != nil {
		testResourceModel.EndTime = types.StringValue(*test.EndTime)
	}
}

// setLabels sets the labels field if present
func setLabels(testResourceModel *resource.BaseTestResourceModel, test *models.TestJSON) {
	if test.Labels != nil {
		labels := make([]resource.LabelModel, len(*test.Labels))
		for i, label := range *test.Labels {
			// Convert []string to []attr.Value
			values := make([]attr.Value, len(label.Values))
			for j, v := range label.Values {
				values[j] = types.StringValue(v)
			}
			labels[i] = resource.LabelModel{
				Key:    types.StringValue(label.Name),
				Values: types.ListValueMust(types.StringType, values),
			}
		}
		testResourceModel.Label = labels
	}
}

// setURLField sets URL field if the model supports it
func setURLField(model resource.TestResourceModelProvider, test *models.TestJSON) {
	if urlProvider, ok := model.(resource.URLProvider); ok {
		if test.URL != nil {
			urlProvider.SetURLField(types.StringValue(*test.URL))
		}
	}
}

// setThresholds sets threshold fields if the model supports it
func setThresholds[T resource.TestResourceModelProvider](model resource.TestResourceModelProvider, test *models.TestJSON, config T) (diags diag.Diagnostics) {
	if provider, ok := any(model).(resource.ThresholdModelProvider); ok {
		configProvider, _ := any(config).(resource.ThresholdModelProvider)
		if !configProvider.GetTestThresholdModel().IsNull() && test.TestThresholds != nil {
			thresholdModel := resource.ThresholdModel{}
			setThresholdField(&thresholdModel.AvailabilityCritical, test.TestThresholds.AvailabilityApdexThresholdCritical)
			setThresholdField(&thresholdModel.AvailabilityWarning, test.TestThresholds.AvailabilityApdexThresholdWarning)
			setThresholdField(&thresholdModel.TestTimeCritical, test.TestThresholds.TestTimeApdexThresholdCritical)
			setThresholdField(&thresholdModel.TestTimeWarning, test.TestThresholds.TestTimeApdexThresholdWarning)
			provider.SetTestThresholdModel(&thresholdModel)
		}
	}
	return
}

// setThresholdField sets a float64 threshold field if the value is not nil
func setThresholdField(field *types.Float64, value *float64) {
	if value != nil {
		*field = types.Float64Value(*value)
	} else {
		*field = types.Float64Null()
	}
}
