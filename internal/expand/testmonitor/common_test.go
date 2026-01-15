package testmonitor

import (
	"testing"

	"catchpoint-provider/internal/models"
	"catchpoint-provider/internal/models/resource"
	"catchpoint-provider/internal/testutil"
	cptypes "catchpoint-provider/internal/types"

	"github.com/hashicorp/terraform-plugin-framework/types"
)

func TestExpandTestConfigFromPlan(t *testing.T) {
	// Create a fully populated plan
	plan := resource.BaseTestResourceModel{
		// Required fields
		DivisionID: types.Int64Value(123),
		ProductID:  types.Int64Value(456),
		Name:       types.StringValue("Test Monitor"),

		// Optional fields with values
		Monitor:               types.StringValue("api"),
		FolderID:              types.Int64Value(789),
		Description:           types.StringValue("Test description"),
		Status:                types.StringValue("active"),
		AlertsPaused:          types.BoolValue(true),
		StartTime:             types.StringValue("2025-01-01T00:00:00Z"),
		EndTime:               types.StringValue("2025-12-31T23:59:59Z"),
		EnableTestDataWebhook: types.BoolValue(false),
	}

	// Create empty config to populate
	config := &models.TestConfig{}

	// Execute the function
	diags := expandTestConfigFromTestResourceModel(plan, config, cptypes.APIType)

	// Verify no errors occurred
	testutil.AssertDiagsHasNoErrors(t, diags)

	// Verify required fields were set correctly
	testutil.AssertEqual(t, "division ID", config.CommonConfig.DivisionID, 123)
	testutil.AssertEqual(t, "product ID", config.ProductID, 456)
	testutil.AssertEqual(t, "test name", config.TestName, "Test Monitor")

	// Verify optional fields were set correctly
	testutil.AssertEqual(t, "folder ID", config.FolderID, 789)
	testutil.AssertEqual(t, "test description", config.TestDescription, "Test description")
	testutil.AssertEqual(t, "alerts paused", config.AlertsPaused, true)
	testutil.AssertEqual(t, "start time", config.StartTime, "2025-01-01T00:00:00Z")
	testutil.AssertEqual(t, "end time", config.EndTime, "2025-12-31T23:59:59Z")
	testutil.AssertEqual(t, "enable test data webhook", config.EnableTestDataWebhook, false)

	// Verify status was processed correctly
	testutil.AssertEqual(t, "status ID", config.Status.ID, 0)
	testutil.AssertEqual(t, "status name", config.Status.Name, "active")
}

func TestExpandTestConfigFromMinimalPlan(t *testing.T) {
	// Create a plan with only required fields
	plan := resource.BaseTestResourceModel{
		// Required fields only
		DivisionID: types.Int64Value(123),
		ProductID:  types.Int64Value(456),
		Name:       types.StringValue("Minimal Test Monitor"),
		Monitor:    types.StringValue("api"),

		// Optional fields left null/empty
		FolderID:              types.Int64Null(),
		Description:           types.StringNull(),
		Status:                types.StringNull(),
		AlertsPaused:          types.BoolNull(),
		StartTime:             types.StringNull(),
		EndTime:               types.StringNull(),
		EnableTestDataWebhook: types.BoolNull(),
	}

	// Create empty config to populate
	config := &models.TestConfig{}

	// Execute the function
	diags := expandTestConfigFromTestResourceModel(plan, config, cptypes.APIType)

	// Verify no errors occurred
	testutil.AssertDiagsHasNoErrors(t, diags)

	// Verify required fields were set correctly
	testutil.AssertEqual(t, "division ID", config.CommonConfig.DivisionID, 123)
	testutil.AssertEqual(t, "product ID", config.ProductID, 456)
	testutil.AssertEqual(t, "test name", config.TestName, "Minimal Test Monitor")

	// Verify optional fields have reasonable defaults or remain unset
	testutil.AssertEqual(t, "folder ID should be zero", config.FolderID, 0)
	testutil.AssertEqual(t, "test description should be empty", config.TestDescription, "")

	// Verify status defaults to active
	testutil.AssertEqual(t, "status should default to active", config.Status.Name, "active")
}
