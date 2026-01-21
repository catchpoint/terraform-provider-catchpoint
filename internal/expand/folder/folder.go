package folder

import (
	"catchpoint-provider/internal/expand"
	"catchpoint-provider/internal/models"
	"catchpoint-provider/internal/models/resource"

	"github.com/hashicorp/terraform-plugin-framework/diag"
)

// ExpandFolderConfigFromPlan expands the Terraform plan (FolderResourceModel) into a FolderConfig model.
func ExpandFolderConfigFromPlan(plan resource.FolderResourceModel, config *models.FolderConfig) (diags diag.Diagnostics) {
	// Expand main fields directly from plan
	if !plan.DivisionID.IsNull() {
		config.CommonConfig.DivisionID = int(plan.DivisionID.ValueInt64())
	}
	if !plan.FolderName.IsNull() {
		config.FolderName = plan.FolderName.ValueString()
	}
	if !plan.ProductID.IsNull() {
		config.ProductID = int(plan.ProductID.ValueInt64())
	}
	if !plan.ParentID.IsNull() {
		config.ParentID = int(plan.ParentID.ValueInt64())
	}

	// Expand the nested objects.
	// For other resource types, these will be set to inherit from parents if the values are not specified.
	// But for FolderConfig, these must be specified explicitly.
	advancedDiags := expand.ExpandAdvancedSettingsConfig(plan.AdvancedSettings, &config.CommonConfig.AdvancedSettingsConfig)
	diags.Append(advancedDiags...)
	if diags.HasError() {
		return
	}

	requestDiags := expand.ExpandRequestSettingsConfig(plan.RequestSettings, &config.CommonConfig.RequestSettingsConfig)
	diags.Append(requestDiags...)
	if diags.HasError() {
		return
	}

	insightDiags := expand.ExpandInsightSettingsConfig(plan.Insights, &config.CommonConfig.InsightSettingsConfig)
	diags.Append(insightDiags...)
	if diags.HasError() {
		return
	}

	scheduleDiags := expand.ExpandScheduleSettingsConfig(plan.ScheduleSettings, &config.CommonConfig.ScheduleSettingsConfig)
	diags.Append(scheduleDiags...)
	if diags.HasError() {
		return
	}

	if plan.AlertSettings != nil {
		alertDiags := expand.ExpandAlertSettingsConfig(plan.AlertSettings, &config.CommonConfig.AlertSettingsConfig)
		diags.Append(alertDiags...)
		if diags.HasError() {
			return
		}
	}

	return
}
