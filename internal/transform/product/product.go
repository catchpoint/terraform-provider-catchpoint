package product

import (
	"context"

	"catchpoint-provider/internal/models"
	"catchpoint-provider/internal/models/resource"
	cpresource "catchpoint-provider/internal/models/resource"
	cpschema "catchpoint-provider/internal/schema"
	"catchpoint-provider/internal/transform"
	cptypes "catchpoint-provider/internal/types"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func JSONToTerraformProduct(ctx context.Context, product *models.ProductJSON, currentState *cpresource.ProductResourceModel) (resource *cpresource.ProductResourceModel, diags diag.Diagnostics) {
	if product == nil {
		return
	}
	resource = &cpresource.ProductResourceModel{}

	// Build all the product attributes using the full schema attribute types
	buildDiags := transformProductFields(resource, product)
	if buildDiags.HasError() {
		diags.Append(buildDiags...)
		return
	}

	// Create the object with the correct attribute types
	objDiags := transformNestedProductFields(ctx, resource, product, currentState)
	diags.Append(objDiags...)
	return
}

func transformProductFields(resource *resource.ProductResourceModel, product *models.ProductJSON) (diags diag.Diagnostics) {
	if resource.IsNull() {
		diags.Append(diag.NewErrorDiagnostic("Null ProductResourceModel", "The provided ProductResourceModel is null."))
		return
	}

	resource.ID = types.Int64Value(int64(product.ID))
	resource.DivisionID = types.Int64Value(int64(product.DivisionID))
	resource.ProductName = types.StringValue(product.Name)

	// Handle status with lookup
	statusName, ok := cptypes.GetStatusTypeName(product.Status.ID)
	if ok {
		resource.Status = types.StringValue(statusName)
	} else {
		resource.Status = types.StringValue(cptypes.Active)
	}

	// Set optional fields
	if product.AlertGroupID != 0 {
		resource.AlertGroupID = types.Int64Value(int64(product.AlertGroupID))
	} else {
		resource.AlertGroupID = types.Int64Null()
	}

	if product.TestDataWebhookID != 0 {
		resource.TestDataWebhookID = types.Int64Value(int64(product.TestDataWebhookID))
	} else {
		resource.TestDataWebhookID = types.Int64Null()
	}

	return
}

func transformNestedProductFields(ctx context.Context, resource *resource.ProductResourceModel, product *models.ProductJSON, currentState *resource.ProductResourceModel) (diags diag.Diagnostics) {
	// If there is no current state, we cannot transform any nested objects as we have no idea what the user originally provided.
	// However, Terraform still demands *something* be provided for these attributes as they are defined in the schema as required.
	// This happens during Import.
	if currentState == nil {
		resource.AdvancedSettings = types.ObjectNull(cpschema.GetAdvancedSettingsAttributeTypesForProductAndFolder())
		resource.RequestSettings = types.ObjectNull(cpschema.GetRequestSettingsAttributeTypes())
		resource.ScheduleSettings = types.ObjectNull(cpschema.GetScheduleSettingsAttributeTypes())
		resource.Insights = types.ObjectNull(cpschema.GetInsightsAttributeTypes())
		return
	}

	// AlertSettings is different from the others as it uses a fully mapped model which is assigned as a pointer.
	alertSettings, d := transform.JSONToTerraformAlertGroup(&product.AlertGroup, currentState.AlertSettings)
	diags.Append(d...)
	resource.AlertSettings = alertSettings

	// The other nested objects are anonymous type.Objects that are mapped much differently.
	advancedSettings, d := transform.JSONToTerraformAdvancedSettingsForProductAndFolder(&product.AdvancedSettings)
	diags.Append(d...)
	resource.AdvancedSettings = advancedSettings.AdvancedSettings

	requestSettings, diag := transform.JSONToTerraformRequestSettings(&product.RequestSettings)
	diags.Append(diag...)
	resource.RequestSettings = requestSettings.RequestSettings

	scheduleSettings, d := transform.JSONToTerraformScheduleSettings(&product.ScheduleSettings)
	diags.Append(d...)
	resource.ScheduleSettings = scheduleSettings.ScheduleSettings

	insightSettings, d := transform.JSONToTerraformInsights(ctx, &product.InsightData)
	diags.Append(d...)
	resource.Insights = insightSettings.Insights

	return
}
