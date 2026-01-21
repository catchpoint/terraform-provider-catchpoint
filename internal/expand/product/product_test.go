package product

import (
	"testing"

	"catchpoint-provider/internal/fields"
	"catchpoint-provider/internal/models"
	"catchpoint-provider/internal/models/resource"
	"catchpoint-provider/internal/testutil"

	"github.com/hashicorp/terraform-plugin-framework/types"
)

func TestExpandProductConfigFromPlanEmptyPlan(t *testing.T) {
	plan := resource.ProductResourceModel{}
	config := models.ProductConfig{}
	diags := ExpandProductConfigFromPlan(plan, &config)

	// Empty plan should not produce any errors but create a default config.
	testutil.AssertDiagsHasNoErrors(t, diags)

	// Note: unless changed, the schema should never allow these to be empty anyway.
	testutil.AssertEqual(t, fields.DivisionID, config.CommonConfig.DivisionID, 0)
	testutil.AssertEqual(t, fields.ProductName, config.ProductName, "")
	testutil.AssertEqual(t, fields.Status, config.Status.ID, 0)          // Default to "active"
	testutil.AssertEqual(t, fields.Status, config.Status.Name, "active") // Default to "active"

	// Schedule settings is required for Product in the schema, but we do not require it here.
	// So anything not set should default.
	scheduleConfig := config.CommonConfig.ScheduleSettingsConfig
	testutil.AssertEqual(t, fields.ScheduleSettingType, scheduleConfig.ScheduleSettingType.ID, 1) // Default to "Override"
	// Slices should not be nil, but empty.
	testutil.AssertDeepEqual(t, fields.NodeIDs, scheduleConfig.NodeIDs, []int{})
	testutil.AssertNotNil(t, fields.Nodes, scheduleConfig.NodeIDs)
	testutil.AssertDeepEqual(t, fields.NodeGroupIDs, scheduleConfig.NodeGroupIDs, []models.IDNameConfig{})
	testutil.AssertNotNil(t, fields.NodeGroups, scheduleConfig.NodeGroupIDs)

	advancedConfig := config.CommonConfig.AdvancedSettingsConfig
	testutil.AssertEqual(t, fields.AdvancedSettingType, advancedConfig.AdvancedSettingType.ID, 1) // Default to "Override"
	testutil.AssertDeepEqual(t, fields.AdvancedSettings, advancedConfig.AppliedTestFlags, []int{})
	testutil.AssertNotNil(t, fields.AdvancedSettings, advancedConfig.AppliedTestFlags)

	requestConfig := config.CommonConfig.RequestSettingsConfig
	testutil.AssertEqual(t, fields.RequestSettingType, requestConfig.RequestSettingType.ID, 1) // Default to "Override"
	testutil.AssertNil(t, fields.HTTPRequestHeaders, requestConfig.TestHTTPHeaderRequests)

	// Insights should default to "No Settings"
	testutil.AssertEqual(t, fields.InsightSettingType, config.CommonConfig.InsightSettingsConfig.InsightSettingType.ID, 3)
}

func TestExpandProductConfigFromPlanAllMainFields(t *testing.T) {
	plan := resource.ProductResourceModel{
		DivisionID:            types.Int64Value(100),
		ProductName:           types.StringValue("Test Product1"),
		ScheduleSettingsModel: resource.ScheduleSettingsModel{},
		AdvancedSettingsModel: resource.AdvancedSettingsModel{},
		RequestSettingsModel:  resource.RequestSettingsModel{},
		InsightSettingsModel:  resource.InsightSettingsModel{},
	}

	config := models.ProductConfig{}
	diags := ExpandProductConfigFromPlan(plan, &config)

	testutil.AssertDiagsHasNoErrors(t, diags)
	testutil.AssertEqual(t, fields.DivisionID, config.CommonConfig.DivisionID, 100)
	testutil.AssertEqual(t, fields.ProductName, config.ProductName, "Test Product1")

	testutil.AssertEqual(t, fields.ScheduleSettingType, config.CommonConfig.ScheduleSettingsConfig.ScheduleSettingType.ID, 1) // Default to "Override"
	testutil.AssertEqual(t, fields.ScheduleSettingType, config.CommonConfig.ScheduleSettingsConfig.ScheduleSettingType.Name, "override")

	testutil.AssertEqual(t, fields.AdvancedSettingType, config.CommonConfig.AdvancedSettingsConfig.AdvancedSettingType.ID, 1) // Default to "Override"
	testutil.AssertEqual(t, fields.AdvancedSettingType, config.CommonConfig.AdvancedSettingsConfig.AdvancedSettingType.Name, "override")

	testutil.AssertEqual(t, fields.RequestSettingType, config.CommonConfig.RequestSettingsConfig.RequestSettingType.ID, 1) // Default to "Override"
	testutil.AssertEqual(t, fields.RequestSettingType, config.CommonConfig.RequestSettingsConfig.RequestSettingType.Name, "override")

	testutil.AssertEqual(t, fields.InsightSettingType, config.CommonConfig.InsightSettingsConfig.InsightSettingType.ID, 3) // Default to "No Settings"
	testutil.AssertEqual(t, fields.InsightSettingType, config.CommonConfig.InsightSettingsConfig.InsightSettingType.Name, "no settings")

}
