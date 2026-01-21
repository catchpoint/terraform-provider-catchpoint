package folder

import (
	"context"

	"catchpoint-provider/internal/models"
	cpresource "catchpoint-provider/internal/models/resource"
	cpschema "catchpoint-provider/internal/schema"
	"catchpoint-provider/internal/transform"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func JSONToTerraformFolder(ctx context.Context, folder *models.FolderJSON, currentState *cpresource.FolderResourceModel) (resource *cpresource.FolderResourceModel, diags diag.Diagnostics) {
	if folder == nil {
		return
	}

	resource = &cpresource.FolderResourceModel{}

	// Build all the folder attributes using the full schema attribute types
	buildDiags := transformFolderFields(resource, folder)
	if buildDiags.HasError() {
		diags.Append(buildDiags...)
		return
	}

	// Create the object with the correct attribute types
	objDiags := transformNestedFolderFields(ctx, resource, folder, currentState)
	diags.Append(objDiags...)
	return
}

func transformFolderFields(resource *cpresource.FolderResourceModel, folder *models.FolderJSON) (diags diag.Diagnostics) {
	resource.ID = types.Int64Value(int64(folder.ID))
	resource.DivisionID = types.Int64Value(int64(folder.DivisionID))
	resource.ProductID = types.Int64Value(int64(folder.ProductID))
	resource.FolderName = types.StringValue(folder.Name)

	if folder.ParentID != nil {
		resource.ParentID = types.Int64Value(int64(*folder.ParentID))
	} else {
		resource.ParentID = types.Int64Null()
	}

	return
}

func transformNestedFolderFields(ctx context.Context, folderResource *cpresource.FolderResourceModel, folder *models.FolderJSON, currentState *cpresource.FolderResourceModel) (diags diag.Diagnostics) {
	// If there is no current state, we cannot transform any nested objects as we have no idea what the user originally provided.
	// However, Terraform still demands *something* be provided for these attributes as they are defined in the schema as required.
	// This happens during Import.
	if currentState == nil {
		folderResource.AdvancedSettings = types.ObjectNull(cpschema.GetAdvancedSettingsAttributeTypesForProductAndFolder())
		folderResource.RequestSettings = types.ObjectNull(cpschema.GetRequestSettingsAttributeTypes())
		folderResource.ScheduleSettings = types.ObjectNull(cpschema.GetScheduleSettingsAttributeTypes())
		folderResource.Insights = types.ObjectNull(cpschema.GetInsightsAttributeTypes())
		return
	}

	// AlertSettings is different from the others as it uses a fully mapped model which is assigned as a pointer.
	if currentState.AlertSettings != nil {
		alertSettings, d := transform.JSONToTerraformAlertGroup(&folder.AlertGroup, currentState.AlertSettings)
		diags.Append(d...)
		folderResource.AlertSettings = alertSettings
	}

	// Only transform nested settings if the user originally provided them in the .tf file. Otherwise, they *must* be null objects. We can
	// neither add the object when it wasn't provided, nor can we skip it entirely if it wasn't provided.
	// By providing the resource when the user did not specify it, we trigger an error in Terraform that the block was absent and is now not.
	// By not returning a properly-nulled object, Terraform will error that various value conversions failed.
	if !currentState.AdvancedSettings.IsNull() && !currentState.AdvancedSettings.IsUnknown() {
		advancedSettings, d := transform.JSONToTerraformAdvancedSettingsForProductAndFolder(&folder.AdvancedSettings)
		diags.Append(d...)
		folderResource.AdvancedSettings = advancedSettings.AdvancedSettings
	} else {
		folderResource.AdvancedSettings = types.ObjectNull(cpschema.GetAdvancedSettingsAttributeTypesForProductAndFolder())
	}

	if !currentState.RequestSettings.IsNull() && !currentState.RequestSettings.IsUnknown() {
		requestSettings, diag := transform.JSONToTerraformRequestSettings(&folder.RequestSettings)
		diags.Append(diag...)
		folderResource.RequestSettings = requestSettings.RequestSettings
	} else {
		folderResource.RequestSettings = types.ObjectNull(cpschema.GetRequestSettingsAttributeTypes())
	}

	if !currentState.ScheduleSettings.IsNull() && !currentState.ScheduleSettings.IsUnknown() {
		scheduleSettings, d := transform.JSONToTerraformScheduleSettings(&folder.ScheduleSettings)
		diags.Append(d...)
		folderResource.ScheduleSettings = scheduleSettings.ScheduleSettings
	} else {
		folderResource.ScheduleSettings = types.ObjectNull(cpschema.GetScheduleSettingsAttributeTypes())
	}

	if !currentState.Insights.IsNull() && !currentState.Insights.IsUnknown() {
		insightSettings, d := transform.JSONToTerraformInsights(ctx, &folder.InsightData)
		diags.Append(d...)
		folderResource.Insights = insightSettings.Insights
	} else {
		folderResource.Insights = types.ObjectNull(cpschema.GetInsightsAttributeTypes())
	}

	return
}
