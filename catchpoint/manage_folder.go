package catchpoint

import (
	"context"
	"reflect"
	"strconv"
	"strings"

	"catchpoint-provider/internal"
	"catchpoint-provider/internal/client"
	folderExpand "catchpoint-provider/internal/expand/folder"
	"catchpoint-provider/internal/fields"
	"catchpoint-provider/internal/logger"
	"catchpoint-provider/internal/merge"
	"catchpoint-provider/internal/models"
	cpresource "catchpoint-provider/internal/models/resource"
	folderSchema "catchpoint-provider/internal/schema/folder"
	"catchpoint-provider/internal/service"
	folderTransform "catchpoint-provider/internal/transform/folder"
	cptypes "catchpoint-provider/internal/types"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type FolderResource struct {
	providerConfig *internal.Config
}

func NewFolderResource() resource.Resource {
	return &FolderResource{}
}

func (r *FolderResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = "manage_folder"
}

func (r *FolderResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = folderSchema.BuildFolderSchema(ctx)
}

func (r *FolderResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	r.providerConfig = req.ProviderData.(*internal.Config)
}

func (r *FolderResource) ValidateConfig(ctx context.Context, req resource.ValidateConfigRequest, resp *resource.ValidateConfigResponse) {
	var config cpresource.FolderResourceModel
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *FolderResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan cpresource.FolderResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Initialize empty config
	folderConfig := models.FolderConfig{}

	// Expand the plan into the config
	expandDiags := folderExpand.ExpandFolderConfigFromPlan(plan, &folderConfig)
	resp.Diagnostics.Append(expandDiags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// The service package creates the JSON payload for the API request.
	clientPayload := service.CreateFolderJSON(folderConfig)

	if r.providerConfig.LogJSON {
		logger.Debug(ctx, "Creating folder with JSON: %v", clientPayload)
	}

	// The client package sends the request to the Catchpoint API.
	respBody, respStatus, folderID, err := client.CreateFolder(r.providerConfig.APIToken, clientPayload)
	if err != nil {
		logger.Debug(ctx, "Error creating Folder: %v. Body: %s, Status: %s", err, respBody, respStatus)
		resp.Diagnostics.AddError("Error creating Folder", err.Error())
		return
	}

	// Set the ID in the plan
	plan.ID = types.Int64Value(folderID)

	updatedState, getDiags := getAPIFolderAsState(ctx, folderID, r.providerConfig.APIToken, &plan)
	resp.Diagnostics.Append(getDiags...)
	if resp.Diagnostics.HasError() {
		return
	}

	finalState := merge.FolderPlanWithState(plan, *updatedState)
	resp.Diagnostics.Append(resp.State.Set(ctx, &finalState)...)
}

func (r *FolderResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var currentState cpresource.FolderResourceModel
	diags := req.State.Get(ctx, &currentState)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Get folder from API
	updatedState, getDiags := getAPIFolderAsState(ctx, currentState.ID.ValueInt64(), r.providerConfig.APIToken, &currentState)
	resp.Diagnostics.Append(getDiags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// If we couldn't find the resource, remove it from state.
	if updatedState.IsNull() {
		resp.State.RemoveResource(ctx)
		return
	}

	logger.Debug(ctx, "READ - Merging state: %v with updatedState: %v", currentState, updatedState)

	resp.Diagnostics.Append(resp.State.Set(ctx, &updatedState)...)
}

func (r *FolderResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	plan, state, folderID := r.extractUpdateData(ctx, req, resp)
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

	r.executeUpdate(ctx, folderID, jsonPatchDocs, plan, resp)
}

func (r *FolderResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	resp.Diagnostics.AddError(
		"Folder Deletion Not Supported",
		"Folders cannot be deleted via the API. To remove this resource from Terraform state without deleting the actual folder, use 'terraform state rm manage_folder.terraform_folder'.",
	)
}

func (r *FolderResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	if req.ID == cptypes.EmptyString {
		resp.Diagnostics.AddError("Invalid import ID", "empty import id")
		return
	}

	id, err := strconv.ParseInt(req.ID, 10, 64)
	if err != nil {
		resp.Diagnostics.AddError("Invalid import ID", "expected numeric folder id")
		return
	}

	if r.providerConfig == nil {
		resp.Diagnostics.AddError("Provider not configured", "the provider must be configured prior to import")
		return
	}

	folderData, _, getErr := client.GetFolder(r.providerConfig.APIToken, id)
	if _, ok := getErr.(*models.ObjectNotFoundError); ok {
		logger.Warn(ctx, "Import: folder with ID %d not found", id)
		resp.Diagnostics.AddError("Not Found", "folder not found")
		return
	}
	if getErr != nil {
		resp.Diagnostics.AddError("Error reading folder", getErr.Error())
		return
	}
	if folderData == nil {
		resp.Diagnostics.AddError("Folder Not Found", "the folder was not found after API call")
		return
	}

	apiState, transformDiags := folderTransform.JSONToTerraformFolder(ctx, folderData, nil)
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

	logger.Debug(ctx, "IMPORT - State set successfully for folder %d", id)
}

func getAPIFolderAsState(ctx context.Context, folderID int64, apiToken string, currentState *cpresource.FolderResourceModel) (state *cpresource.FolderResourceModel, diags diag.Diagnostics) {
	folderData, readStatus, readErr := client.GetFolder(apiToken, folderID)
	if _, ok := readErr.(*models.ObjectNotFoundError); ok {
		logger.Warn(ctx, "Folder with ID %d not found, removing from state", folderID)

		return
	}

	if readErr != nil {
		logger.Debug(ctx, "Error reading created Folder: %v. Status: %s", readErr, readStatus)
		diags.AddError("Error reading created Folder", readErr.Error())
		return
	}

	if folderData == nil {
		diags.AddError("Folder Not Found", "The folder was not found after creation.")
		return
	}

	state, transformDiags := folderTransform.JSONToTerraformFolder(ctx, folderData, currentState)
	diags.Append(transformDiags...)
	if diags.HasError() {
		return
	}

	return
}

func (r *FolderResource) extractUpdateData(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) (cpresource.FolderResourceModel, cpresource.FolderResourceModel, int64) {
	var plan, state cpresource.FolderResourceModel

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

	folderID := state.ID.ValueInt64()
	if folderID == 0 {
		resp.Diagnostics.AddError("Invalid Folder ID", "Folder ID is required for update operation")
		return plan, state, 0
	}

	return plan, state, folderID
}

func (r *FolderResource) generatePatchDocuments(plan, state cpresource.FolderResourceModel, resp *resource.UpdateResponse) []string {
	var jsonPatchDocs []string

	jsonPatchDocs = append(jsonPatchDocs, r.handleBasicFieldUpdates(plan, state)...)
	jsonPatchDocs = append(jsonPatchDocs, r.handleComplexSectionUpdates(plan, state, resp)...)

	return jsonPatchDocs
}

func (r *FolderResource) handleBasicFieldUpdates(plan, state cpresource.FolderResourceModel) (jsonPatchDocs []string) {
	addPatch := func(value string, section string) {
		testConfigUpdate := models.TestConfigUpdate{
			UpdatedFieldValue: value,
		}
		jsonPatchDocs = append(jsonPatchDocs, service.CreateJSONTestPatchDocument(&testConfigUpdate, section, true))
	}

	// Simple fields
	basicFields := []struct {
		planValue, stateValue any
		getValue              func() string
		section               string
	}{
		{plan.ProductID, state.ProductID, plan.ProductID.String, fields.ProductIDSection},
		{plan.FolderName, state.FolderName, plan.FolderName.ValueString, fields.NameSection},
	}

	for _, f := range basicFields {
		if !reflect.ValueOf(f.planValue).MethodByName("Equal").Call([]reflect.Value{reflect.ValueOf(f.stateValue)})[0].Bool() {
			addPatch(f.getValue(), f.section)
		}
	}

	// Special handling for ParentID since it is nullable
	if !plan.ParentID.Equal(state.ParentID) {
		parentID := int64(0)
		if !plan.ParentID.IsNull() {
			parentID = plan.ParentID.ValueInt64()
		}
		folderConfigUpdate := models.FolderConfigUpdate{
			UpdatedFieldValue: strconv.FormatInt(parentID, 10),
		}
		jsonPatchDocs = append(jsonPatchDocs, service.CreateJSONFolderPatchDocument(&folderConfigUpdate, fields.ParentIDSection, true))
	}

	return jsonPatchDocs
}

func (r *FolderResource) handleComplexSectionUpdates(plan, state cpresource.FolderResourceModel, resp *resource.UpdateResponse) []string {
	var jsonPatchDocs []string

	if patch := handleFolderAdvancedSettingsUpdate(plan, state, resp); patch != cptypes.EmptyString {
		jsonPatchDocs = append(jsonPatchDocs, patch)
	}
	if patch := handleFolderRequestSettingsUpdate(plan, state, resp); patch != cptypes.EmptyString {
		jsonPatchDocs = append(jsonPatchDocs, patch)
	}
	if patch := handleFolderAlertSettingsUpdate(plan, state, resp); patch != cptypes.EmptyString {
		jsonPatchDocs = append(jsonPatchDocs, patch)
	}
	if patch := handleFolderInsightsUpdate(plan, state); patch != nil {
		jsonPatchDocs = append(jsonPatchDocs, patch...)
	}
	if patch := handleFolderScheduleSettingsUpdate(plan, state, resp); patch != cptypes.EmptyString {
		jsonPatchDocs = append(jsonPatchDocs, patch)
	}

	return jsonPatchDocs
}

func (r *FolderResource) executeUpdate(ctx context.Context, folderID int64, jsonPatchDocs []string, plan cpresource.FolderResourceModel, resp *resource.UpdateResponse) {
	jsonPatchDoc := "[" + strings.Join(jsonPatchDocs, ",") + "]"

	if r.providerConfig.LogJSON {
		logger.Debug(ctx, "Updating folder %d with JSON PATCH: %v", folderID, jsonPatchDoc)
	}

	respBody, respStatus, completed, err := client.UpdateFolder(r.providerConfig.APIToken, folderID, jsonPatchDoc)
	if err != nil {
		logger.Debug(ctx, "Error updating Folder: %v. Body: %s, Status: %s", err, respBody, respStatus)
		resp.Diagnostics.AddError("Error updating Folder", err.Error())
		return
	}

	if !completed {
		resp.Diagnostics.AddWarning("Update may not be complete", "The API indicated the update may not have completed successfully")
	}

	// Fetch updated folder data
	updatedState, getDiags := getAPIFolderAsState(ctx, folderID, r.providerConfig.APIToken, &plan)
	resp.Diagnostics.Append(getDiags...)
	if resp.Diagnostics.HasError() {
		return
	}

	finalState := merge.FolderPlanWithState(plan, *updatedState)
	finalState.ID = types.Int64Value(folderID)

	// Set the state
	resp.Diagnostics.Append(resp.State.Set(ctx, &finalState)...)
}
