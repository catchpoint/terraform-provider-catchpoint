package catchpoint

import (
	"context"
	"strconv"
	"strings"

	"catchpoint-provider/internal"
	"catchpoint-provider/internal/client"
	productExpand "catchpoint-provider/internal/expand/product"
	"catchpoint-provider/internal/fields"
	"catchpoint-provider/internal/logger"
	"catchpoint-provider/internal/merge"
	"catchpoint-provider/internal/models"
	cpresource "catchpoint-provider/internal/models/resource"
	productSchema "catchpoint-provider/internal/schema/product"
	"catchpoint-provider/internal/service"
	productTransform "catchpoint-provider/internal/transform/product"
	cptypes "catchpoint-provider/internal/types"
	cpvalidation "catchpoint-provider/internal/validation"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ProductResource struct {
	providerConfig *internal.Config
}

func NewProductResource() resource.Resource {
	return &ProductResource{}
}

func (r *ProductResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = "manage_product"
}

func (r *ProductResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = productSchema.BuildProductSchema(ctx)
}

func (r *ProductResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	r.providerConfig = req.ProviderData.(*internal.Config)
}

func (r *ProductResource) ValidateConfig(ctx context.Context, req resource.ValidateConfigRequest, resp *resource.ValidateConfigResponse) {
	var config cpresource.ProductResourceModel
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	if config.ScheduleSettings.IsNull() || config.ScheduleSettings.IsUnknown() {
		resp.Diagnostics.AddError(
			"Missing Schedule Settings",
			"The product resource minimally requires schedule_settings to be configured identifying target nodes.",
		)
	}
}

func (r *ProductResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan cpresource.ProductResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Initialize empty config
	productConfig := models.ProductConfig{}

	// Expand the plan into the config
	expandDiags := productExpand.ExpandProductConfigFromPlan(plan, &productConfig)
	resp.Diagnostics.Append(expandDiags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// The service package creates the JSON payload for the API request.
	clientPayload := service.CreateProductJSON(productConfig)

	if r.providerConfig.LogJSON {
		logger.Debug(ctx, "Creating product with JSON: %v", clientPayload)
	}

	// The client package sends the request to the Catchpoint API.
	respBody, respStatus, productID, err := client.CreateProduct(r.providerConfig.APIToken, clientPayload)
	if err != nil {
		logger.Debug(ctx, "Error creating Product: %v. Body: %s, Status: %s", err, respBody, respStatus)
		resp.Diagnostics.AddError("Error creating Product", err.Error())
		return
	}

	// Set the ID in the plan
	plan.ID = types.Int64Value(productID)

	updatedState, getDiags := getAPIProductAsState(ctx, productID, r.providerConfig.APIToken, &plan)
	resp.Diagnostics.Append(getDiags...)
	if resp.Diagnostics.HasError() {
		return
	}

	finalState := merge.ProductPlanWithState(plan, *updatedState)

	resp.Diagnostics.Append(resp.State.Set(ctx, &finalState)...)
}

func (r *ProductResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var currentState cpresource.ProductResourceModel
	diags := req.State.Get(ctx, &currentState)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Get product from API
	updatedState, getDiags := getAPIProductAsState(ctx, currentState.ID.ValueInt64(), r.providerConfig.APIToken, &currentState)
	resp.Diagnostics.Append(getDiags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// If we couldn't find the resource, remove it from state.
	if updatedState.IsNull() {
		logger.Warn(ctx, "Product with ID %d not found, removing from state", currentState.ID.ValueInt64())
		resp.State.RemoveResource(ctx)
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &updatedState)...)
}

func (r *ProductResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	plan, state, productID := r.extractUpdateData(ctx, req, resp)
	if resp.Diagnostics.HasError() {
		return
	}

	jsonPatchDocs := r.generatePatchDocuments(plan, state, resp)
	if resp.Diagnostics.HasError() {
		return
	}

	if len(jsonPatchDocs) == 0 {
		resp.Diagnostics.AddWarning("No Changes", "No changes detected. Your infrastructure matches the configuration.")
		return
	}

	r.executeUpdate(ctx, productID, jsonPatchDocs, plan, resp)
}

func (r *ProductResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	resp.Diagnostics.AddError(
		"Product Deletion Not Supported",
		"Products cannot be deleted via the API. To remove this resource from Terraform state without deleting the actual product, use 'terraform state rm manage_product.terraform_product'.",
	)
}

func (r *ProductResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	if req.ID == cptypes.EmptyString {
		resp.Diagnostics.AddError("Invalid import ID", "empty import id")
		return
	}

	id, err := strconv.ParseInt(req.ID, 10, 64)
	if err != nil {
		resp.Diagnostics.AddError("Invalid import ID", "expected numeric product id")
		return
	}

	if r.providerConfig == nil {
		resp.Diagnostics.AddError("Provider not configured", "the provider must be configured prior to import")
		return
	}

	productData, _, getErr := client.GetProduct(r.providerConfig.APIToken, id)
	if _, ok := getErr.(*models.ObjectNotFoundError); ok {
		logger.Warn(ctx, "Import: product with ID %d not found", id)
		resp.Diagnostics.AddError("Not Found", "product not found")
		return
	}
	if getErr != nil {
		resp.Diagnostics.AddError("Error reading product", getErr.Error())
		return
	}
	if productData == nil {
		resp.Diagnostics.AddError("Product Not Found", "the product was not found after API call")
		return
	}

	apiState, transformDiags := productTransform.JSONToTerraformProduct(ctx, productData, nil)
	resp.Diagnostics.Append(transformDiags...)
	if resp.Diagnostics.HasError() {
		return
	}

	apiState.ID = types.Int64Value(id)

	resp.Diagnostics.Append(resp.State.Set(ctx, &apiState)...)
	if resp.Diagnostics.HasError() {
		logger.Debug(ctx, "IMPORT - State set failed: %v", resp.Diagnostics.Errors())
		return
	}

	logger.Debug(ctx, "IMPORT - State set successfully for product %d", id)
}

func getAPIProductAsState(ctx context.Context, productID int64, apiToken string, currentState *cpresource.ProductResourceModel) (state *cpresource.ProductResourceModel, diags diag.Diagnostics) {
	productData, readStatus, readErr := client.GetProduct(apiToken, productID)
	if _, ok := readErr.(*models.ObjectNotFoundError); ok {
		return
	}

	if readErr != nil {
		logger.Debug(ctx, "Error reading created Product: %v. Status: %s", readErr, readStatus)
		diags.AddError("Error reading created Product", readErr.Error())
		return
	}

	if productData == nil {
		diags.AddError("Product Not Found", "The product was not found after creation.")
		return
	}

	state, transformDiags := productTransform.JSONToTerraformProduct(ctx, productData, currentState)
	diags.Append(transformDiags...)
	if diags.HasError() {
		return
	}

	return
}

func (r *ProductResource) extractUpdateData(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) (cpresource.ProductResourceModel, cpresource.ProductResourceModel, int64) {
	var plan, state cpresource.ProductResourceModel

	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return plan, state, 0
	}

	diags = req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return plan, state, 0
	}

	productID := state.ID.ValueInt64()
	if productID == 0 {
		resp.Diagnostics.AddError("Invalid Product ID", "Product ID is required for update operation")
		return plan, state, 0
	}

	return plan, state, productID
}

func (r *ProductResource) generatePatchDocuments(plan, state cpresource.ProductResourceModel, resp *resource.UpdateResponse) []string {
	var jsonPatchDocs []string

	jsonPatchDocs = append(jsonPatchDocs, r.handleBasicFieldUpdates(plan, state)...)
	jsonPatchDocs = append(jsonPatchDocs, r.handleComplexSectionUpdates(plan, state, resp)...)

	return jsonPatchDocs
}

func (r *ProductResource) handleBasicFieldUpdates(plan, state cpresource.ProductResourceModel) []string {
	var jsonPatchDocs []string

	if !plan.ProductName.Equal(state.ProductName) {
		productConfigUpdate := models.ProductConfigUpdate{
			UpdatedFieldValue: plan.ProductName.ValueString(),
		}
		jsonPatchDocs = append(jsonPatchDocs, service.CreateJSONProductPatchDocument(&productConfigUpdate, fields.NameSection, true))
	}

	if !plan.Status.IsNull() && plan.Status.ValueString() != cptypes.EmptyString && !plan.Status.Equal(state.Status) {
		updatedStatus := cpvalidation.GetStatusTypeOrDefault(plan.Status.ValueString())
		productConfigUpdate := models.ProductConfigUpdate{
			UpdatedFieldValue: strconv.Itoa(updatedStatus.ID),
		}
		jsonPatchDocs = append(jsonPatchDocs, service.CreateJSONProductPatchDocument(&productConfigUpdate, fields.StatusSection, true))
	}

	if !plan.TestDataWebhookID.IsNull() && plan.TestDataWebhookID.ValueInt64() != 0 && !plan.TestDataWebhookID.Equal(state.TestDataWebhookID) {
		productConfigUpdate := models.ProductConfigUpdate{
			UpdatedFieldValue: strconv.Itoa(int(plan.TestDataWebhookID.ValueInt64())),
		}
		jsonPatchDocs = append(jsonPatchDocs, service.CreateJSONProductPatchDocument(&productConfigUpdate, fields.TestDataWebhookIDSection, true))
	}

	return jsonPatchDocs
}

func (r *ProductResource) handleComplexSectionUpdates(plan, state cpresource.ProductResourceModel, resp *resource.UpdateResponse) []string {
	var jsonPatchDocs []string

	if patch := handleAdvancedSettingsUpdate(plan, state, resp); patch != cptypes.EmptyString {
		jsonPatchDocs = append(jsonPatchDocs, patch)
	}
	if patch := handleRequestSettingsUpdate(plan, state, resp); patch != cptypes.EmptyString {
		jsonPatchDocs = append(jsonPatchDocs, patch)
	}
	if patch := handleAlertSettingsUpdate(plan, state, resp); patch != cptypes.EmptyString {
		jsonPatchDocs = append(jsonPatchDocs, patch)
	}
	if patch := handleInsightsUpdate(plan, state); patch != nil {
		jsonPatchDocs = append(jsonPatchDocs, patch...)
	}
	if patch := handleScheduleSettingsUpdate(plan, state, resp); patch != cptypes.EmptyString {
		jsonPatchDocs = append(jsonPatchDocs, patch)
	}

	return jsonPatchDocs
}

func (r *ProductResource) executeUpdate(ctx context.Context, productID int64, jsonPatchDocs []string, plan cpresource.ProductResourceModel, resp *resource.UpdateResponse) {
	jsonPatchDoc := "[" + strings.Join(jsonPatchDocs, ",") + "]"

	if r.providerConfig.LogJSON {
		logger.Debug(ctx, "Updating product %d with JSON PATCH: %v", productID, jsonPatchDoc)
	}

	respBody, respStatus, completed, err := client.UpdateProduct(r.providerConfig.APIToken, productID, jsonPatchDoc)
	if err != nil {
		logger.Debug(ctx, "Error updating Product: %v. Body: %s, Status: %s", err, respBody, respStatus)
		resp.Diagnostics.AddError("Error updating Product", err.Error())
		return
	}

	if !completed {
		resp.Diagnostics.AddWarning("Update may not be complete", "The API indicated the update may not have completed successfully")
	}

	// Fetch updated product data
	updatedState, getDiags := getAPIProductAsState(ctx, productID, r.providerConfig.APIToken, &plan)
	resp.Diagnostics.Append(getDiags...)
	if resp.Diagnostics.HasError() {
		return
	}

	finalState := merge.ProductPlanWithState(plan, *updatedState)
	finalState.ID = types.Int64Value(productID)

	// Set the state
	resp.Diagnostics.Append(resp.State.Set(ctx, &finalState)...)
}
