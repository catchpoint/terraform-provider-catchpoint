package merge

import (
	"context"

	cpresource "catchpoint-provider/internal/models/resource"
	folderSchema "catchpoint-provider/internal/schema/folder"
)

func FolderPlanWithState(plan, apiState cpresource.FolderResourceModel) (result cpresource.FolderResourceModel) {
	ctx := context.Background()
	folderSchema := folderSchema.BuildFolderSchema(ctx)

	// Merge each nested block dynamically
	result.ScheduleSettings = ScheduleSettings(&plan.ScheduleSettingsModel, &apiState.ScheduleSettingsModel, folderSchema).ScheduleSettings
	result.AlertSettings = AlertSettings(plan.AlertSettings, apiState.AlertSettings)
	result.RequestSettings = RequestSettings(&plan.RequestSettingsModel, &apiState.RequestSettingsModel, folderSchema).RequestSettings
	result.AdvancedSettings = AdvancedSettings(&plan.AdvancedSettingsModel, &apiState.AdvancedSettingsModel, folderSchema).AdvancedSettings
	result.Insights = InsightSettings(&plan.InsightSettingsModel, &apiState.InsightSettingsModel, folderSchema).Insights

	result.DivisionID = plan.DivisionID
	result.FolderName = plan.FolderName
	result.ProductID = plan.ProductID
	result.ParentID = plan.ParentID

	// Computed, read-only setting - always set from API state
	result.ID = apiState.ID

	return
}
