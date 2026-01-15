package catchpoint

import (
	"context"
	"fmt"
	"slices"
	"strconv"
	"strings"

	"catchpoint-provider/internal"
	"catchpoint-provider/internal/client"
	"catchpoint-provider/internal/fields"
	"catchpoint-provider/internal/logger"
	"catchpoint-provider/internal/models"
	cpresource "catchpoint-provider/internal/models/resource"
	"catchpoint-provider/internal/service"
	cptypes "catchpoint-provider/internal/types"
	"catchpoint-provider/internal/validation/advancedsettings"
	"catchpoint-provider/internal/validation/alerttypes"
	testmonitor "catchpoint-provider/internal/validation/testmonitor"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

type TestResource[T cpresource.TestResourceModelProvider] struct {
	providerConfig *internal.Config
	testType       cptypes.TestType
	schemaBuilder  func(context.Context) schema.Schema
	expandFunc     func(T, *models.TestConfig) diag.Diagnostics
	transformFunc  func(context.Context, T, *models.TestJSON, T) diag.Diagnostics
}

func (r *TestResource[T]) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = cptypes.GetTestTypeName(r.testType) + "_test"
}

func (r *TestResource[T]) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = r.schemaBuilder(ctx)
}

func (r *TestResource[T]) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	r.providerConfig = req.ProviderData.(*internal.Config)
}

func (r *TestResource[T]) ConfigValidators(ctx context.Context) []resource.ConfigValidator {
	return []resource.ConfigValidator{
		testmonitor.NewTestConfigValidator(r.testType),
	}
}

func (r *TestResource[T]) ValidateConfig(ctx context.Context, req resource.ValidateConfigRequest, resp *resource.ValidateConfigResponse) {
	var config T
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

}

// ModifyPlan is called after defaults are applied but before the plan is finalized.
// We validate alert rules here because the "monitor" may be a default that isn't set during config validation.
func (r *TestResource[T]) ModifyPlan(ctx context.Context, req resource.ModifyPlanRequest, resp *resource.ModifyPlanResponse) {
	// Get the plan with all defaults applied
	var plan T
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)

	// ModifyPlan also runs during a destroy op but may be at least partially null.
	// If there are any errors or the plan is null, skip validation.
	if resp.Diagnostics.HasError() || plan.GetTestResourceModel().IsNull() {
		return
	}

	monitor := plan.GetTestResourceModel().Monitor.ValueString()
	if monitor == cptypes.EmptyString {
		return
	}

	alertProvider, ok := any(plan).(cpresource.AlertSettingsModelProvider)
	if !ok {
		// This resource type doesn't support alert settings
		return
	}

	// Get the alert settings
	alertSettings := alertProvider.GetAlertSettingsModel()
	if alertSettings == nil || len(alertSettings.AlertRule) == 0 {
		// No alert settings to validate
		return
	}

	// Validate alert rules
	validAlertTypes := alerttypes.GetMonitorAlertTypes(r.testType, monitor)
	if validAlertTypes == nil {
		resp.Diagnostics.AddAttributeError(
			path.Root("monitor"),
			"Invalid Monitor Type",
			fmt.Sprintf("Monitor type %q is not valid for test type %q.",
				monitor,
				cptypes.GetTestTypeName(r.testType)),
		)
		return
	}

	r.validateAlertRules(resp, alertSettings, validAlertTypes, monitor)
	r.validateRequestSettings(plan, monitor, resp)
	r.validateAdvancedSettings(plan, monitor, resp)
}

func (r *TestResource[T]) validateAlertRules(
	resp *resource.ModifyPlanResponse,
	alertSettings *cpresource.AlertSettingsModel,
	validAlertTypes *alerttypes.MonitorAlertTypes,
	monitor string,
) {
	for index, alertRule := range alertSettings.AlertRule {
		if !validAlertTypes.IsAlertTypeValid(alertRule.AlertType.ValueString()) {
			resp.Diagnostics.AddAttributeError(
				path.Root("alert_settings").AtName("alert_rule").AtListIndex(index).AtName("alert_type"),
				"Invalid Alert Type",
				fmt.Sprintf("Alert type %q is not valid for test type %q with monitor type %q. Valid alert types are: %v",
					alertRule.AlertType.ValueString(),
					cptypes.GetTestTypeName(r.testType),
					monitor,
					strings.Join(validAlertTypes.ListAllValidAlertTypes(), ", ")),
			)
			continue
		}

		if !alertRule.AlertSubType.IsNull() && !alertRule.AlertSubType.IsUnknown() {
			alertType := validAlertTypes.AlertTypes[alertRule.AlertType.ValueString()]
			if alertType == nil || !alertType.IsAlertSubTypeValid(alertRule.AlertSubType.ValueString()) {
				resp.Diagnostics.AddAttributeError(
					path.Root("alert_settings").AtName("alert_rule").AtListIndex(index).AtName("alert_sub_type"),
					"Invalid Alert SubType",
					fmt.Sprintf("Alert sub type %q is not valid for alert type %q for test type %q with monitor type %q. Valid alert sub types are: %v",
						alertRule.AlertSubType.ValueString(),
						alertRule.AlertType.ValueString(),
						cptypes.GetTestTypeName(r.testType),
						monitor,
						strings.Join(alertType.SubTypes, ", ")),
				)
			}
		}
	}
}

func (r *TestResource[T]) validateAdvancedSettings(config T, monitor string, resp *resource.ModifyPlanResponse) {
	advancedSettingsProvider, planOK := any(config).(cpresource.AdvancedSettingsModelProvider)
	if !planOK {
		return
	}
	advancedSettingsModel := advancedSettingsProvider.GetAdvancedSettingsModel()
	if advancedSettingsModel == nil || advancedSettingsModel.AdvancedSettings.IsNull() || advancedSettingsModel.AdvancedSettings.IsUnknown() {
		return
	}

	advancedSettingsObj, objOK := any(advancedSettingsModel.AdvancedSettings).(types.Object)
	if !objOK {
		resp.Diagnostics.AddAttributeError(
			path.Root("advanced_settings"),
			"Invalid Advanced Settings",
			"Advanced settings is not an object.",
		)
		return
	}

	if advancedSettingsObj.IsUnknown() || advancedSettingsObj.IsNull() {
		return
	}

	attributes := advancedSettingsObj.Attributes()
	validAdvancedSettings := advancedsettings.GetAdvancedSettingsForTestAndMonitorCombination(r.testType, monitor)
	if validAdvancedSettings == nil {
		resp.Diagnostics.AddAttributeError(
			path.Root("monitor"),
			"Invalid Monitor Type",
			fmt.Sprintf("Monitor type %q is not valid for test type %q.",
				monitor,
				cptypes.GetTestTypeName(r.testType)),
		)
		return
	}

	for attrName, attrValue := range attributes {
		// Only validate if the user actually set the value
		if attrValue.IsNull() || attrValue.IsUnknown() || attrName == "advanced_setting_type" {
			continue
		}
		if !slices.Contains(validAdvancedSettings, attrName) {
			resp.Diagnostics.AddAttributeError(
				path.Root("advanced_settings").AtName(attrName),
				"Invalid Advanced Setting",
				fmt.Sprintf("Advanced setting %q is not valid for test type %q with monitor type %q. Valid advanced settings are: %v",
					attrName,
					cptypes.GetTestTypeName(r.testType),
					monitor,
					strings.Join(validAdvancedSettings, ", ")),
			)
		}
	}
}

func (r *TestResource[T]) validateRequestSettings(config T, monitor string, resp *resource.ModifyPlanResponse) {
	planRequestSettingsProvider, planOK := any(config).(cpresource.RequestSettingsModelProvider)
	if !planOK {
		return
	}
	requestSettingsModel := planRequestSettingsProvider.GetRequestSettingsModel()
	if requestSettingsModel == nil || requestSettingsModel.RequestSettings.IsNull() || requestSettingsModel.RequestSettings.IsUnknown() {
		return
	}

	// API, Playwright, and Puppeteer may include certificates regardless of monitor type.
	includeCertificates := r.testType == cptypes.APIType || r.testType == cptypes.PlaywrightType || r.testType == cptypes.PuppeteerType

	// Web and Transaction tests may include certificates if the monitor type is Emulated, HTTP, or Edge.
	if (r.testType == cptypes.WebType || r.testType == cptypes.TransactionType) &&
		(monitor == cptypes.EmulatedString || monitor == cptypes.HTTPString || monitor == cptypes.EdgeString) {
		includeCertificates = true
	}

	certs := requestSettingsModel.RequestSettings.Attributes()[fields.LibraryCertificateIDs].(types.List)
	// If certs are not allowed but the attribute is set, raise an error.

	logger.DebugBG("IncludeCerts: %v, Certs isNull: %v, IsUnknown: %v", includeCertificates, certs.IsNull(), certs.IsUnknown())

	if !includeCertificates {
		if len(certs.Elements()) > 0 {
			resp.Diagnostics.AddAttributeError(
				path.Root("request_settings").AtName("library_certificate_ids"),
				"Invalid Attribute for Test Type and Monitor Type",
				fmt.Sprintf("The attribute library_certificate_ids is not valid for test type %q with monitor type %q.",
					cptypes.GetTestTypeName(r.testType),
					monitor),
			)
		}
	}
}

func (r *TestResource[T]) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan T
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Initialize empty config
	testConfig := models.TestConfig{}

	// Use the expand function to transform the plan into test config
	expandDiags := r.expandFunc(plan, &testConfig)
	resp.Diagnostics.Append(expandDiags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// The service package creates the JSON payload for the API request.
	clientPayload := service.CreateTestJSON(testConfig)

	if r.providerConfig.LogJSON {
		logger.Debug(ctx, "Creating test with JSON: %v", clientPayload)
	}

	// The client package sends the request to the Catchpoint API.
	respBody, respStatus, testID, err := client.CreateTest(r.providerConfig.APIToken, clientPayload)
	if err != nil {
		logger.Debug(ctx, "Error creating Test: %v. Body: %s, Status: %s", err, respBody, respStatus)
		resp.Diagnostics.AddError("Error creating Test", err.Error())
		return
	}

	// Set the test ID in the context for logging.
	ctx = tflog.SetField(ctx, "test_id", fmt.Sprintf("%d", testID))

	plan.SetID(types.Int64Value(testID))

	// Capture the original plan before computed values, we use this during transformation so we can
	// omit any complex objects that the user did not specify.
	var originalPlan T
	req.Config.Get(ctx, &originalPlan)

	// After creation, read the resource back from the API to get all computed values.
	apiState, stateDiags := r.getAPITestAsState(ctx, testID, r.providerConfig.APIToken, originalPlan)
	resp.Diagnostics.Append(stateDiags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// The `plan` object already has the blocks added by the plan modifier.
	// We now merge the computed values from the `apiState` into the `plan`.
	finalState := r.mergeTestPlanWithState(ctx, plan, apiState)

	// Set the final merged state. This now contains the blocks from the plan modifier
	// and the computed values from the API.
	resp.Diagnostics.Append(resp.State.Set(ctx, &finalState)...)
	if !resp.Diagnostics.HasError() {
		logger.Debug(ctx, "CREATE - State set successfully")
	} else {
		logger.Debug(ctx, "CREATE - State set failed: %v", resp.Diagnostics.Errors())
	}
}

func (r *TestResource[T]) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var plan T
	diags := req.State.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Get the test ID from state.
	testID := plan.GetID().ValueInt64()

	// Set the test ID in the context for logging.
	ctx = tflog.SetField(ctx, "test_id", fmt.Sprintf("%d", testID))

	// Transform the API response to Terraform state
	apiState, stateDiags := r.getAPITestAsState(ctx, testID, r.providerConfig.APIToken, plan)
	resp.Diagnostics.Append(stateDiags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// If we couldn't find the test, remove it from state.
	if apiState.IsNull() {
		resp.State.RemoveResource(ctx)
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &apiState)...)
	if !resp.Diagnostics.HasError() {
		logger.Debug(ctx, "READ - State set successfully")
	} else {
		logger.Debug(ctx, "READ - State set failed: %v", resp.Diagnostics.Errors())
	}
}

func (r *TestResource[T]) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan T
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var state T
	diags = req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	testID := state.GetID().ValueInt64()
	if testID == 0 {
		resp.Diagnostics.AddError("Invalid Test ID", "Test ID is required for update operation")
		return
	}

	// Set the test ID in the context for logging.
	ctx = tflog.SetField(ctx, "test_id", fmt.Sprintf("%d", testID))

	jsonPatchDocs := r.generatePatchDocuments(plan, state, resp)
	if resp.Diagnostics.HasError() {
		return
	}

	if len(jsonPatchDocs) == 0 {
		// Even if no patch docs are generated, we should still refresh state
		// to handle potential drift or out-of-band changes.
		logger.Debug(ctx, "No attribute changes detected, refreshing state from API.")
		r.refreshTestState(ctx, testID, plan, resp, req)
		return
	}

	r.executeUpdate(ctx, testID, jsonPatchDocs, plan, resp, req)
}

func (r *TestResource[T]) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state T
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	testID := state.GetID().ValueInt64()
	if testID == 0 {
		resp.Diagnostics.AddError("Invalid Test ID", "Test ID is required for update operation")
		return
	}

	client.DeleteTest(r.providerConfig.APIToken, testID)
}

func (r *TestResource[T]) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	// Expect the import ID to be the numeric test ID.
	if req.ID == cptypes.EmptyString {
		resp.Diagnostics.AddError("Invalid import ID", "empty import id")
		return
	}

	id, err := strconv.ParseInt(req.ID, 10, 64)
	if err != nil {
		resp.Diagnostics.AddError("Invalid import ID", fmt.Sprintf("expected numeric test id, got: %q", req.ID))
		return
	}

	if r.providerConfig == nil {
		resp.Diagnostics.AddError("Provider not configured", "the provider must be configured prior to import")
		return
	}

	testData, status, getErr := client.GetTest(r.providerConfig.APIToken, id)
	if _, ok := getErr.(*models.ObjectNotFoundError); ok {
		logger.Warn(ctx, "Import: test with ID %d not found", id)
		resp.Diagnostics.AddError("Not Found", fmt.Sprintf("test %d not found (status: %s)", id, status))
		return
	}
	if getErr != nil {
		resp.Diagnostics.AddError("Error reading Test", fmt.Sprintf("error reading test %d: %v (status: %s)", id, getErr, status))
		return
	}
	if testData == nil {
		resp.Diagnostics.AddError("Test Not Found", fmt.Sprintf("test %d not found after API call", id))
		return
	}

	// Create a new model instance and populate it from the API response.
	var model T
	state := model.New().(T)

	transformDiags := r.transformFunc(ctx, state, testData, model)
	resp.Diagnostics.Append(transformDiags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Ensure ID is set on the state
	state.SetID(types.Int64Value(id))

	// Persist the imported state
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		logger.Debug(ctx, "IMPORT - State set failed: %v", resp.Diagnostics.Errors())
		return
	}

	logger.Debug(ctx, "IMPORT - State set successfully for test %d", id)
}

func (r *TestResource[T]) getAPITestAsState(ctx context.Context, testID int64, apiToken string, base T) (state T, diags diag.Diagnostics) {
	testData, readStatus, readErr := client.GetTest(apiToken, testID)
	if _, ok := readErr.(*models.ObjectNotFoundError); ok {
		logger.Warn(ctx, "Test with ID %d not found, removing from state", testID)

		return
	}
	if readErr != nil {
		diags.AddError("Error reading Test", fmt.Sprintf("error reading Test %d (status: %s): %v", testID, readStatus, readErr))
		return
	}

	// Always start with a fresh model for state in Read.
	state = base.New().(T)
	state.SetID(types.Int64Value(testID))

	// Overlay API fields onto existing structure.
	transformDiags := r.transformFunc(ctx, state, testData, base)
	diags.Append(transformDiags...)
	return
}
