package product

import (
	"catchpoint-provider/internal/expand"
	"catchpoint-provider/internal/models"
	"catchpoint-provider/internal/models/resource"
	"catchpoint-provider/internal/types"
	"catchpoint-provider/internal/validation"

	"github.com/hashicorp/terraform-plugin-framework/diag"
)

// ExpandProductConfigFromPlan expands the Terraform plan (ProductResourceModel) into a ProductConfig model.
func ExpandProductConfigFromPlan(plan resource.ProductResourceModel, config *models.ProductConfig) (diags diag.Diagnostics) {
	// Expand main fields directly from plan
	if !plan.DivisionID.IsNull() {
		config.CommonConfig.DivisionID = int(plan.DivisionID.ValueInt64())
	}
	if !plan.ProductName.IsNull() {
		config.ProductName = plan.ProductName.ValueString()
	}
	if !plan.AlertGroupID.IsNull() {
		config.AlertGroupID = int(plan.AlertGroupID.ValueInt64())
	}

	config.Status = expand.GetStatusFromStringOrActive(plan.Status)

	if !plan.TestDataWebhookID.IsNull() {
		config.TestDataWebhookID = int(plan.TestDataWebhookID.ValueInt64())
	}

	// Expand the nested objects.
	// For other resource types, these will be set to inherit from parents if the values are not specified.
	// But for ProductConfig, these must be specified explicitly.
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

	if !plan.Insights.IsUnknown() && !plan.Insights.IsNull() {
		insightDiags := expand.ExpandInsightSettingsConfig(plan.Insights, &config.CommonConfig.InsightSettingsConfig)
		diags.Append(insightDiags...)
		if diags.HasError() {
			return
		}
	} else {
		// If Insights is not specified, set it to the default "no settings" type.
		// Note: This is different from other resources where we override or inherit from parent. Only Product can do this.
		config.CommonConfig.InsightSettingsConfig.InsightSettingType = validation.GetInsightSettingTypeOrDefault(types.NoSettings)
	}

	// Schedule settings are always required for Product.
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

	// For Product, these all get set to "Override" whether they have settings or not.
	override := validation.GetGenericSettingTypeOrDefault(types.Override)
	config.CommonConfig.AdvancedSettingsConfig.AdvancedSettingType = override
	config.CommonConfig.RequestSettingsConfig.RequestSettingType = override
	config.CommonConfig.ScheduleSettingsConfig.ScheduleSettingType = override
	config.CommonConfig.AlertSettingsConfig.AlertSettingType = override

	return
}
