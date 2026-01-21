package folder

import (
	"testing"

	"catchpoint-provider/internal/fields"
	"catchpoint-provider/internal/models"
	"catchpoint-provider/internal/models/resource"
	"catchpoint-provider/internal/testutil"

	"github.com/hashicorp/terraform-plugin-framework/types"
)

func TestExpandFolderConfigFromPlanEmptyPlan(t *testing.T) {
	plan := resource.FolderResourceModel{}
	config := models.FolderConfig{}
	diags := ExpandFolderConfigFromPlan(plan, &config)

	// Empty plan should not produce any errors but create a default config.
	testutil.AssertDiagsHasNoErrors(t, diags)

	// Note: unless changed, the schema should never allow these to be empty anyway.
	testutil.AssertEqual(t, fields.DivisionID, config.CommonConfig.DivisionID, 0)
	testutil.AssertEqual(t, fields.FolderName, config.FolderName, "")
	testutil.AssertEqual(t, fields.ProductID, config.ProductID, 0)
	testutil.AssertEqual(t, fields.ParentID, config.ParentID, 0)

	// Schedule settings is required for Folder in the schema, but we do not require it here.
	// So anything not set should default.
	scheduleConfig := config.CommonConfig.ScheduleSettingsConfig
	testutil.AssertEqual(t, fields.ScheduleSettingType, scheduleConfig.ScheduleSettingType.ID, 0) // Default to "Inherit"
	// Slices should not be nil, but empty.
	testutil.AssertDeepEqual(t, fields.NodeIDs, scheduleConfig.NodeIDs, []int{})
	testutil.AssertNotNil(t, fields.Nodes, scheduleConfig.NodeIDs)
	testutil.AssertDeepEqual(t, fields.NodeGroupIDs, scheduleConfig.NodeGroupIDs, []models.IDNameConfig{})
	testutil.AssertNotNil(t, fields.NodeGroups, scheduleConfig.NodeGroupIDs)

	advancedConfig := config.CommonConfig.AdvancedSettingsConfig
	testutil.AssertEqual(t, fields.AdvancedSettingType, advancedConfig.AdvancedSettingType.ID, 0) // Default to "Inherit"
	testutil.AssertDeepEqual(t, fields.AdvancedSettings, advancedConfig.AppliedTestFlags, []int{})
	testutil.AssertNotNil(t, fields.AdvancedSettings, advancedConfig.AppliedTestFlags)

	requestConfig := config.CommonConfig.RequestSettingsConfig
	testutil.AssertEqual(t, fields.RequestSettingType, requestConfig.RequestSettingType.ID, 0) // Default to "Inherit"
	testutil.AssertNil(t, fields.HTTPRequestHeaders, requestConfig.TestHTTPHeaderRequests)

	testutil.AssertEqual(t, fields.InsightSettingType, config.CommonConfig.InsightSettingsConfig.InsightSettingType.ID, 0) // Default to "Inherit"
}

func TestExpandFolderConfigFromPlanAllMainFields(t *testing.T) {
	plan := resource.FolderResourceModel{
		DivisionID:            types.Int64Value(100),
		ProductID:             types.Int64Value(200),
		ParentID:              types.Int64Value(300),
		FolderName:            types.StringValue("Test Folder1"),
		ScheduleSettingsModel: resource.ScheduleSettingsModel{},
		AdvancedSettingsModel: resource.AdvancedSettingsModel{},
		RequestSettingsModel:  resource.RequestSettingsModel{},
		InsightSettingsModel:  resource.InsightSettingsModel{},
	}

	config := models.FolderConfig{}
	diags := ExpandFolderConfigFromPlan(plan, &config)

	testutil.AssertDiagsHasNoErrors(t, diags)
	testutil.AssertEqual(t, fields.DivisionID, config.CommonConfig.DivisionID, 100)
	testutil.AssertEqual(t, fields.FolderName, config.FolderName, "Test Folder1")
	testutil.AssertEqual(t, fields.ProductID, config.ProductID, 200)
	testutil.AssertEqual(t, fields.ParentID, config.ParentID, 300)
}
