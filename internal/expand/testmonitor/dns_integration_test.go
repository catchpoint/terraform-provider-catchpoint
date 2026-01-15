package testmonitor

import (
	"testing"

	"catchpoint-provider/internal/fields"
	"catchpoint-provider/internal/models"
	"catchpoint-provider/internal/models/resource"
	testmonitor "catchpoint-provider/internal/models/resource/testmonitor"
	"catchpoint-provider/internal/testutil"
	cptypes "catchpoint-provider/internal/types"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func TestExpandDNSDirectTestConfigFromPlan(t *testing.T) {

	// Create a fully populated plan
	plan := testmonitor.DNSTestResourceModel{
		BaseTestResourceModel: resource.BaseTestResourceModel{
			// Required fields
			DivisionID: types.Int64Value(123),
			ProductID:  types.Int64Value(456),
			Name:       types.StringValue("DNS Direct Monitor"),

			// Optional fields with values
			Monitor:               types.StringValue("dns direct"),
			FolderID:              types.Int64Value(789),
			Description:           types.StringValue("DNS Direct description"),
			Status:                types.StringValue("active"),
			AlertsPaused:          types.BoolValue(true),
			StartTime:             types.StringValue("2025-01-02T00:00:00Z"),
			EndTime:               types.StringValue("2025-12-30T23:59:59Z"),
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
		TestDomain: types.StringValue("example.com"),

		AdvancedSettingsModel: resource.AdvancedSettingsModel{
			AdvancedSettings: types.Object{},
		},

		AlertSettings: nil,

		ScheduleSettingsModel: resource.ScheduleSettingsModel{
			ScheduleSettings: types.Object{},
		},

		DNSServer: types.StringValue("8.8.8.8"),
		QueryType: types.StringValue("a"),
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
	testutil.AssertEqual(t, fields.TestName, config.TestName, "DNS Direct Monitor")

	// Verify optional fields were set correctly
	testutil.AssertEqual(t, fields.FolderID, config.FolderID, 789)
	testutil.AssertEqual(t, fields.TestDescription, config.TestDescription, "DNS Direct description")
	testutil.AssertEqual(t, fields.AlertsPaused, config.AlertsPaused, true)
	testutil.AssertEqual(t, fields.StartTime, config.StartTime, "2025-01-02T00:00:00Z")
	testutil.AssertEqual(t, fields.EndTime, config.EndTime, "2025-12-30T23:59:59Z")
	testutil.AssertEqual(t, fields.EnableTestDataWebhook, config.EnableTestDataWebhook, false)

	// Verify status was processed correctly
	testutil.AssertEqual(t, fields.Status+"ID", config.Status.ID, 0)
	testutil.AssertEqual(t, fields.Status+"Name", config.Status.Name, "active")

	// Verify that the DNS monitor fields were set correctly
	testutil.AssertEqual(t, fields.Monitor+"ID", config.Monitor.ID, 13)
	testutil.AssertEqual(t, fields.Monitor+"Name", config.Monitor.Name, "dns direct")
	testutil.AssertEqual(t, "test type ID", config.TestType.ID, 5)
	testutil.AssertEqual(t, "test type name", config.TestType.Name, "dns")

	// For an DNS test, the URL should not be set.
	testutil.AssertEqual(t, fields.TestURL, config.TestURL, "example.com")
	testutil.AssertEqual(t, "RequestData", config.Script.RequestData, cptypes.EmptyString)

	// Verify that nested config objects were set to inherit from parent since they're all empty.
	testutil.AssertEqual(t, fields.AlertSettings, config.CommonConfig.AlertSettingsConfig.AlertSettingType.Name, cptypes.Inherit)
	testutil.AssertEqual(t, fields.AdvancedSettings, config.CommonConfig.AdvancedSettingsConfig.AdvancedSettingType.Name, cptypes.Inherit)
	testutil.AssertEqual(t, fields.ScheduleSettings, config.CommonConfig.ScheduleSettingsConfig.ScheduleSettingType.Name, cptypes.Inherit)

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

	testutil.AssertEqual(t, "DNS Server", config.DNSServer, "8.8.8.8")
	testutil.AssertEqual(t, "QueryTypeName", config.DNSQueryType.Name, "a")
	testutil.AssertEqual(t, "QueryTypeID", config.DNSQueryType.ID, 1)
}

func TestExpandDNSExperienceTestConfigFromPlan(t *testing.T) {

	// Create a fully populated plan
	plan := testmonitor.DNSTestResourceModel{
		BaseTestResourceModel: resource.BaseTestResourceModel{
			// Required fields
			DivisionID: types.Int64Value(123),
			ProductID:  types.Int64Value(456),
			Name:       types.StringValue("Test Monitor"),

			// Optional fields with values
			Monitor:               types.StringValue("dns experience"),
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
		TestDomain: types.StringValue("example.com"),

		AdvancedSettingsModel: resource.AdvancedSettingsModel{
			AdvancedSettings: types.Object{},
		},

		AlertSettings: nil,

		ScheduleSettingsModel: resource.ScheduleSettingsModel{
			ScheduleSettings: types.Object{},
		},

		DNSServer: types.StringValue("1.1.1.1"),
		QueryType: types.StringValue("cname"),
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

	// Verify that the DNS monitor fields were set correctly
	testutil.AssertEqual(t, fields.Monitor+"ID", config.Monitor.ID, 12)
	testutil.AssertEqual(t, fields.Monitor+"Name", config.Monitor.Name, "dns experience")
	testutil.AssertEqual(t, "test type ID", config.TestType.ID, 5)
	testutil.AssertEqual(t, "test type name", config.TestType.Name, "dns")

	// For an DNS test, the URL should not be set.
	testutil.AssertEqual(t, fields.TestURL, config.TestURL, "example.com")
	testutil.AssertEqual(t, "RequestData", config.Script.RequestData, cptypes.EmptyString)

	// Verify that nested config objects were set to inherit from parent since they're all empty.
	testutil.AssertEqual(t, fields.AlertSettings, config.CommonConfig.AlertSettingsConfig.AlertSettingType.Name, cptypes.Inherit)
	testutil.AssertEqual(t, fields.AdvancedSettings, config.CommonConfig.AdvancedSettingsConfig.AdvancedSettingType.Name, cptypes.Inherit)
	testutil.AssertEqual(t, fields.ScheduleSettings, config.CommonConfig.ScheduleSettingsConfig.ScheduleSettingType.Name, cptypes.Inherit)

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

	testutil.AssertEqual(t, "DNS Server", config.DNSServer, "1.1.1.1")
	testutil.AssertEqual(t, "QueryTypeName", config.DNSQueryType.Name, "cname")
	testutil.AssertEqual(t, "QueryTypeID", config.DNSQueryType.ID, 5)
}
