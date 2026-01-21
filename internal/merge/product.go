package merge

import (
	"context"

	cpresource "catchpoint-provider/internal/models/resource"
	productSchema "catchpoint-provider/internal/schema/product"
	cptypes "catchpoint-provider/internal/types"
)

func ProductPlanWithState(plan, apiState cpresource.ProductResourceModel) (result cpresource.ProductResourceModel) {
	ctx := context.Background()
	productSchema := productSchema.BuildProductSchema(ctx)

	// Merge each nested block dynamically
	result.ScheduleSettings = ScheduleSettings(&plan.ScheduleSettingsModel, &apiState.ScheduleSettingsModel, productSchema).ScheduleSettings
	result.AlertSettings = AlertSettings(plan.AlertSettings, apiState.AlertSettings)
	result.RequestSettings = RequestSettings(&plan.RequestSettingsModel, &apiState.RequestSettingsModel, productSchema).RequestSettings
	result.AdvancedSettings = AdvancedSettings(&plan.AdvancedSettingsModel, &apiState.AdvancedSettingsModel, productSchema).AdvancedSettings
	result.Insights = InsightSettings(&plan.InsightSettingsModel, &apiState.InsightSettingsModel, productSchema).Insights

	result.DivisionID = plan.DivisionID
	result.ProductName = plan.ProductName

	// Computed, read-only setting - always set from API state
	result.ID = apiState.ID

	// Computed settings, but the user may or may not omit them. If they don't, set from the backend.
	// If the status is not set in the plan, take it from the API state
	if plan.Status.IsNull() || plan.Status.ValueString() == cptypes.EmptyString {
		result.Status = apiState.Status
	} else {
		result.Status = plan.Status
	}

	if plan.TestDataWebhookID.IsNull() || plan.TestDataWebhookID.ValueInt64() == 0 {
		result.TestDataWebhookID = apiState.TestDataWebhookID
	} else {
		result.TestDataWebhookID = plan.TestDataWebhookID
	}

	if plan.AlertGroupID.IsNull() || plan.AlertGroupID.ValueInt64() == 0 {
		result.AlertGroupID = apiState.AlertGroupID
	} else {
		result.AlertGroupID = plan.AlertGroupID
	}
	return
}
