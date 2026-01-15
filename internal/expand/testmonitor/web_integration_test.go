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

func TestExpandWebTestConfigFromPlan(t *testing.T) {
	test := []struct {
		monitor   string
		monitorID int
	}{
		{"http", 2},
		{"emulated", 3},
		{"chrome", 18},
		{"mobile", 26},
		{"playback", 19},
		{"mobile playback", 20},
	}
	for _, tt := range test {
		// Create a fully populated plan
		plan := testmonitor.WebTestResourceModel{
			BaseTestResourceModel: resource.BaseTestResourceModel{
				// Required fields
				DivisionID: types.Int64Value(123),
				ProductID:  types.Int64Value(456),
				Name:       types.StringValue("Test Monitor"),

				// Optional fields with values
				Monitor:               types.StringValue(tt.monitor),
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
			TestUrl: types.StringValue("http://www.example.com"),
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

		// Verify that the Web monitor fields were set correctly
		testutil.AssertEqual(t, fields.Monitor+"ID", config.Monitor.ID, tt.monitorID)
		testutil.AssertEqual(t, fields.Monitor+"Name", config.Monitor.Name, tt.monitor)
		testutil.AssertEqual(t, "test type ID", config.TestType.ID, 0)
		testutil.AssertEqual(t, "test type name", config.TestType.Name, "web")

		// For an Web test, the URL should be set
		testutil.AssertEqual(t, fields.URL, config.TestURL, "http://www.example.com")

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
}

// Note: Prod values being used here. Some IDs are different in QA and Stage.
func TestExpandWebTestConfigFromPlanChromeVersion(t *testing.T) {
	test := []struct {
		chromeVersion              string
		chromeVersionID            int
		chromeApplicationVersionID *int
	}{
		{"stable", 1, nil},
		{"preview", 2, nil},
		{"53", 3, testutil.ToIntPtr(1)},
		{"59", 3, testutil.ToIntPtr(3)},
		{"63", 3, testutil.ToIntPtr(4)},
		{"66", 3, testutil.ToIntPtr(5)},
		{"71", 3, testutil.ToIntPtr(8)},
		{"75", 3, testutil.ToIntPtr(7)},
		{"85", 3, testutil.ToIntPtr(12)},
		{"87", 3, testutil.ToIntPtr(13)},
		{"89", 3, testutil.ToIntPtr(14)},
		{"97", 3, testutil.ToIntPtr(28357)},
		{"108", 3, testutil.ToIntPtr(28558)},
		{"120", 3, testutil.ToIntPtr(31965)},
	}
	for _, tt := range test {
		// Create a fully populated plan
		plan := testmonitor.WebTestResourceModel{
			BaseTestResourceModel: resource.BaseTestResourceModel{
				// Required fields
				DivisionID: types.Int64Value(123),
				ProductID:  types.Int64Value(456),
				Name:       types.StringValue("Test Monitor"),

				// Optional fields with values
				Monitor:               types.StringValue("chrome"),
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
			TestUrl:       types.StringValue("http://www.example.com"),
			ChromeVersion: types.StringValue(tt.chromeVersion),
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

		testutil.AssertEqual(t, fields.ChromeVersion+"ApplicationVersionType.ID", config.ChromeApplicationVersion.ApplicationVersionType.ID, tt.chromeVersionID)
		testutil.AssertEqual(t, fields.ChromeVersion+"ApplicationVersionType.Name", config.ChromeApplicationVersion.ApplicationVersionType.Name, tt.chromeVersion)

		if tt.chromeVersionID == 3 {
			testutil.AssertEqual(t, fields.ChromeVersion+"ApplicationVersionID", config.ChromeApplicationVersion.ApplicationVersionID, *tt.chromeApplicationVersionID)
		}

	}
}

func TestExpandWebTestConfigFromPlanSimulateField(t *testing.T) {
	test := []struct {
		userAgent   string
		userAgentID int
	}{
		{"ie", 1},
		{"chrome", 2},
		{"android", 3},
		{"iphone", 4},
		{"ipad 2", 5},
		{"kindle fire", 6},
		{"galaxy tab", 7},
		{"iphone 5", 8},
		{"ipad mini", 9},
		{"galaxy note", 10},
		{"nexus 7", 11},
		{"nexus 4", 12},
		{"nokia lumia920", 13},
		{"iphone 6", 14},
		{"blackberry z30", 15},
		{"galaxy s4", 16},
		{"htc onex", 17},
		{"lg optimusg", 18},
		{"droid razr hd", 19},
		{"nexus 6", 20},
		{"iphone 6s", 21},
		{"galaxy s6", 22},
		{"iphone 7", 23},
		{"google pixel", 24},
		{"galaxy s8", 25},
	}
	for _, tt := range test {
		// Create a fully populated plan
		plan := testmonitor.WebTestResourceModel{
			BaseTestResourceModel: resource.BaseTestResourceModel{
				// Required fields
				DivisionID: types.Int64Value(123),
				ProductID:  types.Int64Value(456),
				Name:       types.StringValue("Test Monitor"),

				// Optional fields with values
				Monitor:               types.StringValue("chrome"),
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
			TestUrl:  types.StringValue("http://www.example.com"),
			Simulate: types.StringValue(tt.userAgent),
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

		testutil.AssertEqual(t, fields.Simulate+"ID", config.SimulateDevice.ID, tt.userAgentID)
		testutil.AssertEqual(t, fields.Simulate+"Name", config.SimulateDevice.Name, tt.userAgent)
	}
}
