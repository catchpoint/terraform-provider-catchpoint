package catchpoint

import (
	"context"

	"catchpoint-provider/internal/merge"
	cpresource "catchpoint-provider/internal/models/resource"
	cptypes "catchpoint-provider/internal/types"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func (r *TestResource[T]) mergeTestPlanWithState(ctx context.Context, plan, apiState T) T {
	// Copy the plan to start with, then merge in any computed fields from the API state.
	// Critical: do not merge anything that isn't computed. Blocks cannot be computed.
	result := plan
	testSchema := r.schemaBuilder(ctx)

	result.SetID(apiState.GetID()) // Always set from backend, not plan

	r.mergeComplexBlocks(result, plan, apiState, testSchema)

	r.mergeOptionalFields(result, plan, apiState)

	r.mergeDNSFields(result, apiState)

	r.mergeSimulateFields(result, apiState)

	r.mergeChromeVersionFields(result, apiState)

	r.mergeGatewayAddressOrHostUpdate(result, apiState)

	r.mergeThresholds(result, apiState)

	// The ID is always set from the API state, this is not a user-settable field.
	result.GetTestResourceModel().ID = apiState.GetTestResourceModel().ID

	return result
}

func (r *TestResource[T]) mergeComplexBlocks(result, plan, apiState T, testSchema schema.Schema) {

	if _, ok := any(result).(cpresource.AdvancedSettingsModelProvider); ok {
		r.mergeAdvancedSettings(result, plan, apiState, testSchema)
	}

	if _, ok := any(result).(cpresource.AlertSettingsModelProvider); ok {
		r.mergeAlertSettings(result, plan, apiState)
	}

	if _, ok := any(result).(cpresource.RequestSettingsModelProvider); ok {
		r.mergeRequestSettings(result, plan, apiState, testSchema)
	}

	if _, ok := any(result).(cpresource.ScheduleSettingsModelProvider); ok {
		r.mergeScheduleSettings(result, plan, apiState, testSchema)
	}

	if _, ok := any(result).(cpresource.InsightSettingsModelProvider); ok {
		r.mergeInsightSettings(result, plan, apiState, testSchema)
	}
}

func (r *TestResource[T]) mergeGatewayAddressOrHostUpdate(result, apiState T) {
	if planModel, ok := any(result).(cpresource.GatewayAddressOrHostModelProvider); ok {
		apiModel := any(apiState).(cpresource.GatewayAddressOrHostModelProvider).GetGatewayAddressOrHostModel()
		resultModel := any(result).(cpresource.GatewayAddressOrHostModelProvider).GetGatewayAddressOrHostModel()

		if planModel.GetGatewayAddressOrHostModel().GatewayAddressOrHost.IsNull() || planModel.GetGatewayAddressOrHostModel().GatewayAddressOrHost.IsUnknown() {
			resultModel.GatewayAddressOrHost = apiModel.GatewayAddressOrHost
		}
	}
}

func (r *TestResource[T]) mergeAdvancedSettings(result, plan, apiState T, testSchema schema.Schema) {
	planModel := any(plan).(cpresource.AdvancedSettingsModelProvider).GetAdvancedSettingsModel()
	apiModel := any(apiState).(cpresource.AdvancedSettingsModelProvider).GetAdvancedSettingsModel()
	resultModel := any(result).(cpresource.AdvancedSettingsModelProvider)

	resultModel.SetAdvancedSettingsModel(merge.AdvancedSettingsForTest(planModel, apiModel, testSchema, r.testType))
}

func (r *TestResource[T]) mergeAlertSettings(result, plan, apiState T) {
	planModel := any(plan).(cpresource.AlertSettingsModelProvider).GetAlertSettingsModel()
	apiModel := any(apiState).(cpresource.AlertSettingsModelProvider).GetAlertSettingsModel()
	resultModel := any(result).(cpresource.AlertSettingsModelProvider)

	resultModel.SetAlertSettingsModel(merge.AlertSettings(planModel, apiModel))
}

func (r *TestResource[T]) mergeRequestSettings(result, plan, apiState T, testSchema schema.Schema) {
	planModel := any(plan).(cpresource.RequestSettingsModelProvider).GetRequestSettingsModel()
	apiModel := any(apiState).(cpresource.RequestSettingsModelProvider).GetRequestSettingsModel()
	resultModel := any(result).(cpresource.RequestSettingsModelProvider)

	resultModel.SetRequestSettingsModel(merge.RequestSettings(planModel, apiModel, testSchema))
}

func (r *TestResource[T]) mergeScheduleSettings(result, plan, apiState T, testSchema schema.Schema) {
	planModel := any(plan).(cpresource.ScheduleSettingsModelProvider).GetScheduleSettingsModel()
	apiModel := any(apiState).(cpresource.ScheduleSettingsModelProvider).GetScheduleSettingsModel()
	resultModel := any(result).(cpresource.ScheduleSettingsModelProvider)

	resultModel.SetScheduleSettingsModel(merge.ScheduleSettings(planModel, apiModel, testSchema))
}

func (r *TestResource[T]) mergeInsightSettings(result, plan, apiState T, testSchema schema.Schema) {
	planModel := any(plan).(cpresource.InsightSettingsModelProvider).GetInsightSettingsModel()
	apiModel := any(apiState).(cpresource.InsightSettingsModelProvider).GetInsightSettingsModel()
	resultModel := any(result).(cpresource.InsightSettingsModelProvider)

	resultModel.SetInsightSettingsModel(merge.InsightSettings(planModel, apiModel, testSchema))
}

func (r *TestResource[T]) mergeOptionalFields(result, plan, apiState T) {
	planModel := plan.GetTestResourceModel()
	apiModel := apiState.GetTestResourceModel()
	resultModel := result.GetTestResourceModel()

	mergeField(&resultModel.EndTime, planModel.EndTime, apiModel.EndTime)
	mergeField(&resultModel.FolderID, planModel.FolderID, apiModel.FolderID)

	mergeLabels(&resultModel.Label, planModel.Label, apiModel.Label)

	mergeField(&resultModel.EnableTestDataWebhook, planModel.EnableTestDataWebhook, apiModel.EnableTestDataWebhook)
	mergeField(&resultModel.AlertsPaused, planModel.AlertsPaused, apiModel.AlertsPaused)
	mergeField(&resultModel.Monitor, planModel.Monitor, apiModel.Monitor)
	mergeField(&resultModel.Description, planModel.Description, apiModel.Description)

	if planModel.StartTime.IsNull() || planModel.StartTime.ValueString() == cptypes.EmptyString {
		resultModel.StartTime = apiModel.StartTime
	} else {
		resultModel.StartTime = planModel.StartTime
	}
}

func (r *TestResource[T]) mergeThresholds(result, apiState T) {
	if _, ok := any(result).(cpresource.ThresholdModelProvider); !ok {
		return // T does not support Thresholds, nothing to do
	}

	resultModel := any(result).(cpresource.ThresholdModelProvider).GetTestThresholdModel()
	apiModel := any(apiState).(cpresource.ThresholdModelProvider).GetTestThresholdModel()

	// If the result is currently null, it's because the plan was null. We cannot present a block
	// when it was absent in the plan.
	// If both are null, there's also nothing to merge.
	if (resultModel.IsNull() && apiModel.IsNull()) || resultModel.IsNull() {
		return
	}

	mergeField(&resultModel.TestTimeWarning, resultModel.TestTimeWarning, apiModel.TestTimeWarning)
	mergeField(&resultModel.TestTimeCritical, resultModel.TestTimeCritical, apiModel.TestTimeCritical)
	mergeField(&resultModel.AvailabilityWarning, resultModel.AvailabilityWarning, apiModel.AvailabilityWarning)
	mergeField(&resultModel.AvailabilityCritical, resultModel.AvailabilityCritical, apiModel.AvailabilityCritical)

	any(result).(cpresource.ThresholdModelProvider).SetTestThresholdModel(resultModel)

}

func (r *TestResource[T]) mergeDNSFields(result, apiState T) {
	if dnsProvider, ok := any(result).(cpresource.DNSTestModelProvider); ok {
		if apiDNSProvider, ok := any(apiState).(cpresource.DNSTestModelProvider); ok {
			if dnsServer, ok := apiDNSProvider.GetDNSServerField(); ok {
				dnsProvider.SetDNSServerField(dnsServer)
			} else {
				dnsProvider.SetDNSServerField(types.StringNull())
			}

			if dnsQueryType, ok := apiDNSProvider.GetQueryTypeField(); ok {
				dnsProvider.SetQueryTypeField(dnsQueryType)
			} else {
				dnsProvider.SetQueryTypeField(types.StringNull())
			}
		}
	}
}

func (r *TestResource[T]) mergeSimulateFields(result, apiState T) {
	provider, ok := any(result).(cpresource.SimulateProvider)
	if !ok {
		return // T does not support Simulate, nothing to do
	}

	planSimulate, planOk := provider.GetSimulateField()
	if planOk && !planSimulate.IsNull() && !planSimulate.IsUnknown() {
		provider.SetSimulateField(planSimulate)
		return
	}

	apiProvider, apiOk := any(apiState).(cpresource.SimulateProvider)
	if apiOk {
		apiSimulate, apiSimOk := apiProvider.GetSimulateField()
		if apiSimOk && !apiSimulate.IsNull() && !apiSimulate.IsUnknown() {
			provider.SetSimulateField(apiSimulate)
			return
		}
	}

	// Defensive: always set to empty string if both plan and API are null/unknown
	provider.SetSimulateField(types.StringValue(""))
}

func (r *TestResource[T]) mergeChromeVersionFields(result, apiState T) {
	provider, ok := any(result).(cpresource.ChromeVersionProvider)
	if !ok {
		return // T does not support ChromeVersion, nothing to do
	}

	planChromeVersion, planOk := provider.GetChromeVersionField()
	if planOk && !planChromeVersion.IsNull() && !planChromeVersion.IsUnknown() {
		provider.SetChromeVersionField(planChromeVersion)
		return
	}

	apiProvider, apiOk := any(apiState).(cpresource.ChromeVersionProvider)
	if apiOk {
		apiChromeVersion, apiChromeOk := apiProvider.GetChromeVersionField()
		if apiChromeOk && !apiChromeVersion.IsNull() && !apiChromeVersion.IsUnknown() {
			provider.SetChromeVersionField(apiChromeVersion)
			return
		}
	}

	// Defensive: always set to empty string if both plan and API are null/unknown
	provider.SetChromeVersionField(types.StringValue(""))
}

// mergeField is a helper method to merge optional fields with standard null/unknown checks
func mergeField[V any](result *V, planValue, apiValue V) {
	type nullableField interface {
		IsUnknown() bool
		IsNull() bool
	}

	if field, ok := any(planValue).(nullableField); ok && (field.IsUnknown() || field.IsNull()) {
		*result = apiValue
	} else {
		*result = planValue
	}
}

// mergeLabels unions plan + api label slices. If plan == nil -> use api (preserve drift).
// Otherwise keep plan entries and append any api entries with names not present in plan.
func mergeLabels(result *[]cpresource.LabelModel, plan, api []cpresource.LabelModel) {
	// If plan is nil, treat as "not provided" and preserve the API state entirely.
	if plan == nil {
		*result = api
		return
	}

	// Build a set of label names from the plan and copy plan entries first.
	nameSet := make(map[string]struct{}, len(plan))
	merged := make([]cpresource.LabelModel, 0, len(plan)+len(api))

	for _, l := range plan {
		merged = append(merged, l)
		nameSet[l.Key.String()] = struct{}{}
	}

	// Append API labels that are not present in the plan (by name).
	for _, l := range api {
		if _, found := nameSet[l.Key.String()]; !found {
			merged = append(merged, l)
			nameSet[l.Key.String()] = struct{}{}
		}
	}

	*result = merged
}
