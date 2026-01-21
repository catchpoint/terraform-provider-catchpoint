package resource

// This file is for both the Plan and State models. This needs to be expanded to use strict-types instead
// of generic types, but for now it will work as-is. This represents both the initial user input as well
// as the final state of the resource after creation or update. The fields contained within represent anything
// that can be set in the Terraform configuration for a given resource as well as anything that the backend
// may set and the customer would want to know about (e.g. IDs can't be set, but the customer still needs to see them).
//
// TODO: Refactor to strict types for better validation and type safety.
//
// IMPORTANT: Everything specified in the schema must be here and vice-versa or Terraform will error.
// This means that if we create specific XxxResourceModel objects to replace types.Object usage of the alert,
// schedule, request, and insights settings, then we cannot dynamically create those schemas or the model
// will not always match.
// We can take this route, but it will require a lot of refactoring and validation will have to move to the
// resource level.

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// The definition for a Product Resource to align with the schema.
type ProductResourceModel struct {
	ID                types.Int64         `tfsdk:"id"`
	DivisionID        types.Int64         `tfsdk:"division_id"`
	ProductName       types.String        `tfsdk:"product_name"`
	AlertGroupID      types.Int64         `tfsdk:"alert_group_id"`
	Status            types.String        `tfsdk:"status"`
	TestDataWebhookID types.Int64         `tfsdk:"test_data_webhook_id"`
	AlertSettings     *AlertSettingsModel `tfsdk:"alert_settings"`
	// Note: as dynamic types.Objects, these cannot be pointers or hold complex objects right now.
	AdvancedSettingsModel
	RequestSettingsModel
	ScheduleSettingsModel
	InsightSettingsModel
}

// The TestScriptResourceModel defines the fields for test types that support testRequestData objects. i.e. any test
// type that takes a script (requestData) definition instead of a URL.
type TestScriptResourceModel struct {
	// API, Playwright, Puppeteer, and Transaction tests don't use a URL-mapped field, they use these settings instead.
	// These get mapped to a testRequestData object with the script as 'requestData'. See TestRequestDataJSON object.
	Script     types.String `tfsdk:"test_script"`
	ScriptType types.String `tfsdk:"test_script_type"` // Type of script (e.g., "selenium", "javascript", etc.)
}

// This struct outlines the fields that *all* Tests share.
type BaseTestResourceModel struct {
	// Common fields
	AlertsPaused          types.Bool   `tfsdk:"alerts_paused"`
	DivisionID            types.Int64  `tfsdk:"division_id"`
	EnableTestDataWebhook types.Bool   `tfsdk:"enable_test_data_webhook"`
	ID                    types.Int64  `tfsdk:"id"`
	Label                 []LabelModel `tfsdk:"label"`
	Monitor               types.String `tfsdk:"monitor"`
	ProductID             types.Int64  `tfsdk:"product_id"`
	StartTime             types.String `tfsdk:"start_time"`
	Status                types.String `tfsdk:"status"`
	Name                  types.String `tfsdk:"test_name"`

	// Optional fields that all tests share but may not be set.
	EndTime     types.String `tfsdk:"end_time"`
	FolderID    types.Int64  `tfsdk:"folder_id"`
	Description types.String `tfsdk:"test_description"`
}

// An optional block that some test types have which sets warning/critical thresholds.
type ThresholdModel struct {
	TestTimeWarning      types.Float64 `tfsdk:"test_time_warning"`
	TestTimeCritical     types.Float64 `tfsdk:"test_time_critical"`
	AvailabilityWarning  types.Float64 `tfsdk:"availability_warning"`
	AvailabilityCritical types.Float64 `tfsdk:"availability_critical"`
}

// A label that can be attached to a test. Tests may apply more than one label.
// Colorization is random and handled during expansion.
type LabelModel struct {
	Key    types.String `tfsdk:"key"`
	Values types.List   `tfsdk:"values"`
}

// The AdvancedSettingsModel defines the object for AdvancedSettings for all resources that support
// advancedSettings and advancedSettingsModel. Not all tests have this block.
type AdvancedSettingsModel struct {
	AdvancedSettings types.Object `tfsdk:"advanced_settings"`
}

// The RequestSettingsModel defines the object for RequestSettings for all resources that support
// requestSettings. Not all tests have this block.
type RequestSettingsModel struct {
	RequestSettings types.Object `tfsdk:"request_settings"`
}

// The ScheduleSettingsModel defines the object for ScheduleSettings for all resources that support
// scheduleSetting and scheduleSettings. Not all tests have this block.
type ScheduleSettingsModel struct {
	ScheduleSettings types.Object `tfsdk:"schedule_settings"`
}

// The InsightSettingsModel defines the object for InsightSettings for all resources that support
// insights and insightsData. Not all tests have this block.
type InsightSettingsModel struct {
	Insights types.Object `tfsdk:"insights"`
}

// The GatewayAddressOrHostModel sets the optional gateway_address_or_host field that is only available
// for certain test types.
type GatewayAddressOrHostModel struct {
	GatewayAddressOrHost types.String `tfsdk:"gateway_address_or_host"`
}

type SSLSettingsModel struct {
	EnforceCertificatePinning    types.Bool   `tfsdk:"enforce_certificate_pinning"`
	EnforceCertificateKeyPinning types.Bool   `tfsdk:"enforce_certificate_key_pinning"`
	FileData                     types.String `tfsdk:"file_data"`
	PassPhrase                   types.String `tfsdk:"passphrase"`
}

func (m *ProductResourceModel) IsNull() bool {
	return m == nil
}
