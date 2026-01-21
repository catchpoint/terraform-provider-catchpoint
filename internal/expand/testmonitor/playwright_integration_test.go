package testmonitor

import (
	"testing"

	"catchpoint-provider/internal/fields"
	"catchpoint-provider/internal/helpers"
	"catchpoint-provider/internal/models"
	"catchpoint-provider/internal/models/resource"
	testmonitor "catchpoint-provider/internal/models/resource/testmonitor"
	"catchpoint-provider/internal/testutil"
	cptypes "catchpoint-provider/internal/types"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func TestExpandPlaywrightTestConfigFromPlan(t *testing.T) {
	script := `// Step - 1 \nname("step")\nname("step 1")\nopen("https://www.amazon.com")`

	// Create a fully populated plan
	plan := testmonitor.PlaywrightTestResourceModel{
		BaseTestResourceModel: resource.BaseTestResourceModel{
			// Required fields
			DivisionID: types.Int64Value(123),
			ProductID:  types.Int64Value(456),
			Name:       types.StringValue("Test Monitor"),

			// Optional fields with values
			Monitor:               types.StringValue("edge"),
			FolderID:              types.Int64Value(789),
			Description:           types.StringValue("Test description"),
			Status:                types.StringValue("active"),
			AlertsPaused:          types.BoolValue(true),
			StartTime:             types.StringValue("2025-01-01T00:00:00Z"),
			EndTime:               types.StringValue("2025-12-31T23:59:59Z"),
			EnableTestDataWebhook: types.BoolValue(false),
			Label: []resource.LabelModel{
				{
					Key: types.StringValue("env"),
					Values: types.ListValueMust(types.StringType, []attr.Value{
						types.StringValue("prod"),
						types.StringValue("staging")}),
				},
			},
		},
		TestScriptResourceModel: resource.TestScriptResourceModel{
			Script:     types.StringValue(script),
			ScriptType: types.StringValue("playwright"),
		},

		AdvancedSettingsModel: resource.AdvancedSettingsModel{
			AdvancedSettings: types.Object{},
		},

		AlertSettings: nil,

		RequestSettingsModel: resource.RequestSettingsModel{
			RequestSettings: types.Object{},
		},

		ScheduleSettingsModel: resource.ScheduleSettingsModel{
			ScheduleSettings: types.Object{},
		},

		InsightSettingsModel: resource.InsightSettingsModel{
			Insights: types.Object{},
		},

		GatewayAddressOrHostModel: resource.GatewayAddressOrHostModel{
			GatewayAddressOrHost: types.StringValue("gateway.example.com"),
		},
	}

	// Create empty config to populate
	config := &models.TestConfig{}

	// Execute the function
	diags := ExpandTestConfig(&plan, config)

	// Verify no errors occurred
	testutil.AssertDiagsHasNoErrors(t, diags)

	// Verify required fields were set correctly
	testutil.AssertEqual(t, fields.DivisionID, config.CommonConfig.DivisionID, 123)
	testutil.AssertEqual(t, fields.ProductID, config.ProductID, 456)
	testutil.AssertEqual(t, fields.TestName, config.TestName, "Test Monitor")

	// Verify optional fields were set correctly
	testutil.AssertEqual(t, fields.FolderID, config.FolderID, 789)
	testutil.AssertEqual(t, fields.TestDescription, config.TestDescription, "Test description")
	testutil.AssertEqual(t, fields.AlertsPaused, config.AlertsPaused, true)
	testutil.AssertEqual(t, fields.StartTime, config.StartTime, "2025-01-01T00:00:00Z")
	testutil.AssertEqual(t, fields.EndTime, config.EndTime, "2025-12-31T23:59:59Z")
	testutil.AssertEqual(t, fields.EnableTestDataWebhook, config.EnableTestDataWebhook, false)

	// Verify status was processed correctly
	testutil.AssertEqual(t, fields.Status+"ID", config.Status.ID, 0)
	testutil.AssertEqual(t, fields.Status+"Name", config.Status.Name, "active")

	// Verify that the Playwright monitor fields were set correctly
	testutil.AssertEqual(t, fields.Monitor+"ID", config.Monitor.ID, 39)
	testutil.AssertEqual(t, fields.Monitor+"Name", config.Monitor.Name, "edge")
	testutil.AssertEqual(t, "test type ID", config.TestType.ID, 25)
	testutil.AssertEqual(t, "test type name", config.TestType.Name, "playwright")

	// For an Playwright test, the URL should not be set.
	testutil.AssertEqual(t, fields.TestURL, config.TestURL, cptypes.EmptyString)
	testutil.AssertNotNil(t, "RequestData", config.Script.RequestData)
	testutil.AssertEqual(t, fields.TestScript, config.Script.RequestData, helpers.NormalizeScript(script))
	testutil.AssertEqual(t, fields.TestScriptType, config.Script.TransactionScriptType.Name, "playwright")
	testutil.AssertEqual(t, fields.TestScriptType+"ID", config.Script.TransactionScriptType.ID, 3)
	testutil.AssertEqual(t, fields.TestScript+"Type", config.Script.TransactionScriptType.Name, "playwright")

	// Monitor and TestType should be replicated into the Script block.
	testutil.AssertEqual(t, fields.Monitor+"ID_ScriptBlock", config.Script.Monitor.ID, 39)
	testutil.AssertEqual(t, fields.Monitor+"Name_ScriptBlock", config.Script.Monitor.Name, "edge")
	testutil.AssertEqual(t, "ScriptBlock testtype.ID", config.Script.TestType.ID, 25)
	testutil.AssertEqual(t, "ScriptBlock testtype.name", config.Script.TestType.Name, "playwright")

	// Verify that nested config objects were set to inherit from parent since they're all empty.
	testutil.AssertEqual(t, fields.AlertSettings, config.CommonConfig.AlertSettingsConfig.AlertSettingType.Name, cptypes.Inherit)
	testutil.AssertEqual(t, fields.AdvancedSettings, config.CommonConfig.AdvancedSettingsConfig.AdvancedSettingType.Name, cptypes.Inherit)
	testutil.AssertEqual(t, fields.RequestSettings, config.CommonConfig.RequestSettingsConfig.RequestSettingType.Name, cptypes.Inherit)
	testutil.AssertEqual(t, fields.Insights, config.CommonConfig.InsightSettingsConfig.InsightSettingType.Name, cptypes.Inherit)
	testutil.AssertEqual(t, fields.ScheduleSettings, config.CommonConfig.ScheduleSettingsConfig.ScheduleSettingType.Name, cptypes.Inherit)

	// Verify GatewayAddressOrHost was set correctly.
	testutil.AssertEqual(t, fields.GatewayAddressOrHost, config.GatewayAddressOrHost, "gateway.example.com")

	// Verify labels were expanded correctly
	if len(config.Labels) != 1 {
		t.Fatalf("Expected 1 label, got %d", len(config.Labels))
	}
	label := config.Labels[0]
	testutil.AssertEqual(t, "label key", label.Name, "env")
	if len(label.Values) != 2 {
		t.Fatalf("Expected 2 label values, got %d", len(label.Values))
	}
	testutil.AssertEqual(t, "first label value", label.Values[0], "prod")
	testutil.AssertEqual(t, "second label value", label.Values[1], "staging")
}
