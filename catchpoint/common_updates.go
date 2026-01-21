package catchpoint

import (
	"catchpoint-provider/internal/expand"
	"catchpoint-provider/internal/fields"
	"catchpoint-provider/internal/models"
	cpresource "catchpoint-provider/internal/models/resource"
	"catchpoint-provider/internal/service"
	cptypes "catchpoint-provider/internal/types"

	"github.com/hashicorp/terraform-plugin-framework/resource"
)

func handleAdvancedSettingsUpdate(plan, state cpresource.ProductResourceModel, resp *resource.UpdateResponse) string {
	if !plan.AdvancedSettings.Equal(state.AdvancedSettings) {
		productConfig := models.ProductConfig{}
		expandDiags := expand.ExpandAdvancedSettingsConfig(plan.AdvancedSettings, &productConfig.CommonConfig.AdvancedSettingsConfig)
		resp.Diagnostics.Append(expandDiags...)
		if resp.Diagnostics.HasError() {
			return cptypes.EmptyString
		}

		productConfigUpdate := models.ProductConfigUpdate{
			UpdatedAdvancedSettingsSection: service.AdvancedSettingFromProductConfig(&productConfig),
			SectionToUpdate:                fields.AdvancedSettingsModelSection,
		}
		return service.CreateJSONProductPatchDocument(&productConfigUpdate, productConfigUpdate.SectionToUpdate, false)
	}
	return cptypes.EmptyString
}

func handleRequestSettingsUpdate(plan, state cpresource.ProductResourceModel, resp *resource.UpdateResponse) string {
	if !plan.RequestSettings.Equal(state.RequestSettings) {
		productConfig := models.ProductConfig{}
		expandDiags := expand.ExpandRequestSettingsConfig(plan.RequestSettings, &productConfig.CommonConfig.RequestSettingsConfig)
		resp.Diagnostics.Append(expandDiags...)
		if resp.Diagnostics.HasError() {
			return cptypes.EmptyString
		}

		productConfigUpdate := models.ProductConfigUpdate{
			UpdatedRequestSettingsSection: service.RequestSettingFromProductConfig(&productConfig),
			SectionToUpdate:               fields.RequestSettingsSection,
		}
		return service.CreateJSONProductPatchDocument(&productConfigUpdate, productConfigUpdate.SectionToUpdate, false)
	}
	return cptypes.EmptyString
}

func handleAlertSettingsUpdate(plan, state cpresource.ProductResourceModel, resp *resource.UpdateResponse) (jsonPatch string) {
	if plan.AlertSettings != state.AlertSettings {
		productConfig := models.ProductConfig{}
		expandDiags := expand.ExpandAlertSettingsConfig(plan.AlertSettings, &productConfig.CommonConfig.AlertSettingsConfig)
		resp.Diagnostics.Append(expandDiags...)
		if resp.Diagnostics.HasError() {
			return
		}

		productConfigUpdate := models.ProductConfigUpdate{
			UpdatedAlertSettingsSection: service.AlertSettingsFromProductConfig(&productConfig),
			SectionToUpdate:             fields.AlertGroupSection,
		}
		jsonPatch = service.CreateJSONProductPatchDocument(&productConfigUpdate, productConfigUpdate.SectionToUpdate, false)
	}
	return
}

// This function is a bit of an oddity because the backend wants a more advanced schema than we currently use
// to set the values to begin with. Thus, instead of patching the entire insight section, we individually
// patch the indicators and tracepoints separately.
func handleInsightsUpdate(plan, state cpresource.ProductResourceModel) (response []string) {
	// If both are null, no change needed
	if plan.Insights.IsNull() && state.Insights.IsNull() {
		return []string{}
	}

	// Extract plan attributes
	planIndicators, planTracepoints := extractInsightsAttributes(plan.Insights)

	// Extract state attributes
	stateIndicators, stateTracepoints := extractInsightsAttributes(state.Insights)

	// Compare individual components
	indicatorsChanged := !planIndicators.Equal(stateIndicators)
	tracepointsChanged := !planTracepoints.Equal(stateTracepoints)

	// Only update if there are actual changes
	if !indicatorsChanged && !tracepointsChanged {
		return []string{}
	}

	if indicatorsChanged {
		indicatorValues := listToIntSlice(planIndicators)
		indicatorsConfigUpdate := models.ProductConfigUpdate{
			SectionToUpdate: fields.InsightsDataSection + fields.IndicatorsSection,
		}

		for _, val := range indicatorValues {
			indicatorsConfigUpdate.UpdatedInsightSettingsSection = append(indicatorsConfigUpdate.UpdatedInsightSettingsSection, map[string]int{"id": int(val)})
		}

		response = append(response, service.CreateJSONProductPatchDocument(&indicatorsConfigUpdate, indicatorsConfigUpdate.SectionToUpdate, false))
	}
	if tracepointsChanged {
		tracepointValues := listToIntSlice(planTracepoints)
		tracepointsConfigUpdate := models.ProductConfigUpdate{
			SectionToUpdate: fields.InsightsDataSection + fields.TracepointsSection,
		}
		for _, val := range tracepointValues {
			tracepointsConfigUpdate.UpdatedInsightSettingsSection = append(tracepointsConfigUpdate.UpdatedInsightSettingsSection, map[string]int{"id": int(val)})
		}
		response = append(response, service.CreateJSONProductPatchDocument(&tracepointsConfigUpdate, tracepointsConfigUpdate.SectionToUpdate, false))
	}

	return
}

func handleScheduleSettingsUpdate(plan, state cpresource.ProductResourceModel, resp *resource.UpdateResponse) string {
	if !plan.ScheduleSettings.Equal(state.ScheduleSettings) {
		productConfig := models.ProductConfig{}
		expandDiags := expand.ExpandScheduleSettingsConfig(plan.ScheduleSettings, &productConfig.CommonConfig.ScheduleSettingsConfig)
		resp.Diagnostics.Append(expandDiags...)
		if resp.Diagnostics.HasError() {
			return cptypes.EmptyString
		}

		productConfigUpdate := models.ProductConfigUpdate{
			UpdatedScheduleSettingsSection: service.ScheduleSettingsFromProductConfig(&productConfig),
			SectionToUpdate:                fields.ScheduleSettingsSection,
		}
		return service.CreateJSONProductPatchDocument(&productConfigUpdate, productConfigUpdate.SectionToUpdate, false)
	}
	return cptypes.EmptyString
}

func handleFolderAdvancedSettingsUpdate(plan, state cpresource.FolderResourceModel, resp *resource.UpdateResponse) string {
	if !plan.AdvancedSettings.Equal(state.AdvancedSettings) {
		productConfig := models.ProductConfig{}
		expandDiags := expand.ExpandAdvancedSettingsConfig(plan.AdvancedSettings, &productConfig.CommonConfig.AdvancedSettingsConfig)
		resp.Diagnostics.Append(expandDiags...)
		if resp.Diagnostics.HasError() {
			return cptypes.EmptyString
		}

		productConfigUpdate := models.ProductConfigUpdate{
			UpdatedAdvancedSettingsSection: service.AdvancedSettingFromProductConfig(&productConfig),
			SectionToUpdate:                fields.AdvancedSettingsSection,
		}
		return service.CreateJSONProductPatchDocument(&productConfigUpdate, productConfigUpdate.SectionToUpdate, false)
	}
	return cptypes.EmptyString
}

func handleFolderRequestSettingsUpdate(plan, state cpresource.FolderResourceModel, resp *resource.UpdateResponse) string {
	if !plan.RequestSettings.Equal(state.RequestSettings) {
		productConfig := models.ProductConfig{}
		expandDiags := expand.ExpandRequestSettingsConfig(plan.RequestSettings, &productConfig.CommonConfig.RequestSettingsConfig)
		resp.Diagnostics.Append(expandDiags...)
		if resp.Diagnostics.HasError() {
			return cptypes.EmptyString
		}

		productConfigUpdate := models.ProductConfigUpdate{
			UpdatedRequestSettingsSection: service.RequestSettingFromProductConfig(&productConfig),
			SectionToUpdate:               fields.RequestSettingSection,
		}
		return service.CreateJSONProductPatchDocument(&productConfigUpdate, productConfigUpdate.SectionToUpdate, false)
	}
	return cptypes.EmptyString
}

func handleFolderAlertSettingsUpdate(plan, state cpresource.FolderResourceModel, resp *resource.UpdateResponse) (jsonPatch string) {
	if plan.AlertSettings != state.AlertSettings {
		productConfig := models.ProductConfig{}
		expandDiags := expand.ExpandAlertSettingsConfig(plan.AlertSettings, &productConfig.CommonConfig.AlertSettingsConfig)
		resp.Diagnostics.Append(expandDiags...)
		if resp.Diagnostics.HasError() {
			return
		}

		productConfigUpdate := models.ProductConfigUpdate{
			UpdatedAlertSettingsSection: service.AlertSettingsFromProductConfig(&productConfig),
			SectionToUpdate:             fields.AlertGroupSection,
		}
		jsonPatch = service.CreateJSONProductPatchDocument(&productConfigUpdate, productConfigUpdate.SectionToUpdate, false)
	}
	return
}

// This function is a bit of an oddity because the backend wants a more advanced schema than we currently use
// to set the values to begin with. Thus, instead of patching the entire insight section, we individually
// patch the indicators and tracepoints separately.
func handleFolderInsightsUpdate(plan, state cpresource.FolderResourceModel) (response []string) {
	// If both are null, no change needed
	if plan.Insights.IsNull() && state.Insights.IsNull() {
		return []string{}
	}

	// Extract plan attributes
	planIndicators, planTracepoints := extractInsightsAttributes(plan.Insights)

	// Extract state attributes
	stateIndicators, stateTracepoints := extractInsightsAttributes(state.Insights)

	// Compare individual components
	indicatorsChanged := !planIndicators.Equal(stateIndicators)
	tracepointsChanged := !planTracepoints.Equal(stateTracepoints)

	// Only update if there are actual changes
	if !indicatorsChanged && !tracepointsChanged {
		return []string{}
	}

	if indicatorsChanged {
		indicatorValues := listToIntSlice(planIndicators)
		indicatorsConfigUpdate := models.ProductConfigUpdate{
			SectionToUpdate: fields.InsightsSection + fields.IndicatorsSection,
		}

		for _, val := range indicatorValues {
			indicatorsConfigUpdate.UpdatedInsightSettingsSection = append(indicatorsConfigUpdate.UpdatedInsightSettingsSection, map[string]int{"id": int(val)})
		}

		response = append(response, service.CreateJSONProductPatchDocument(&indicatorsConfigUpdate, indicatorsConfigUpdate.SectionToUpdate, false))
	}
	if tracepointsChanged {
		tracepointValues := listToIntSlice(planTracepoints)
		tracepointsConfigUpdate := models.ProductConfigUpdate{
			SectionToUpdate: fields.InsightsSection + fields.TracepointsSection,
		}
		for _, val := range tracepointValues {
			tracepointsConfigUpdate.UpdatedInsightSettingsSection = append(tracepointsConfigUpdate.UpdatedInsightSettingsSection, map[string]int{"id": int(val)})
		}
		response = append(response, service.CreateJSONProductPatchDocument(&tracepointsConfigUpdate, tracepointsConfigUpdate.SectionToUpdate, false))
	}

	return
}

func handleFolderScheduleSettingsUpdate(plan, state cpresource.FolderResourceModel, resp *resource.UpdateResponse) string {
	if !plan.ScheduleSettings.Equal(state.ScheduleSettings) {
		productConfig := models.ProductConfig{}
		expandDiags := expand.ExpandScheduleSettingsConfig(plan.ScheduleSettings, &productConfig.CommonConfig.ScheduleSettingsConfig)
		resp.Diagnostics.Append(expandDiags...)
		if resp.Diagnostics.HasError() {
			return cptypes.EmptyString
		}

		productConfigUpdate := models.ProductConfigUpdate{
			UpdatedScheduleSettingsSection: service.ScheduleSettingsFromProductConfig(&productConfig),
			SectionToUpdate:                fields.ScheduleSettingSection,
		}
		return service.CreateJSONProductPatchDocument(&productConfigUpdate, productConfigUpdate.SectionToUpdate, false)
	}
	return cptypes.EmptyString
}
