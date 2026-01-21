package transform

import (
	"testing"

	"catchpoint-provider/internal/fields"
	"catchpoint-provider/internal/models"
	"catchpoint-provider/internal/testutil"

	"github.com/hashicorp/terraform-plugin-framework/types"
)

func TestJSONToTerraformScheduleSettings(t *testing.T) {
	// Create test data with all fields populated
	scheduleSettingsJSON := &models.ScheduleSettingsJSON{
		ScheduleSettingType: models.GenericIDNameJSON{
			ID:   1,
			Name: "override",
		},
		RunScheduleID:         testutil.ToIntPtr(123),
		MaintenanceScheduleID: testutil.ToIntPtr(456),
		Frequency: models.GenericIDNameJSON{
			ID:   1,
			Name: "1 minute",
		},
		TestNodeDistribution: models.GenericIDNameJSON{
			ID:   1,
			Name: "concurrent",
		},
		Nodes: []models.NodeJSON{
			{ID: testutil.ToIntPtr(1001), Name: "Node 1"},
			{ID: testutil.ToIntPtr(1002), Name: "Node 2"},
			{ID: testutil.ToIntPtr(1003), Name: "Node 3"},
		},
		NodeGroups: []models.NodeGroupJSON{
			{NodeGroupID: testutil.ToIntPtr(2001), Name: "Group 1"},
			{NodeGroupID: testutil.ToIntPtr(2002), Name: "Group 2"},
		},
		NoOfSubsetNodes: testutil.ToIntPtr(5),
	}

	// Execute the function
	result, diags := JSONToTerraformScheduleSettings(scheduleSettingsJSON)

	// Verify no errors occurred
	testutil.AssertDiagsHasNoErrors(t, diags)

	// Verify the result is not null
	testutil.AssertEqual(t, "result should not be null", result.ScheduleSettings.IsNull(), false)

	// Get the attributes from the result
	attrs := result.ScheduleSettings.Attributes()

	// Verify schedule setting type
	scheduleSettingType := attrs[fields.ScheduleSettingType].(types.String)
	testutil.AssertEqual(t, "schedule setting type", scheduleSettingType.ValueString(), "override")

	// Verify run schedule ID
	runScheduleID := attrs[fields.RunScheduleID].(types.Int64)
	testutil.AssertEqual(t, "run schedule id", runScheduleID.ValueInt64(), int64(123))

	// Verify maintenance schedule ID
	maintenanceScheduleID := attrs[fields.MaintenanceScheduleID].(types.Int64)
	testutil.AssertEqual(t, "maintenance schedule id", maintenanceScheduleID.ValueInt64(), int64(456))

	// Verify frequency
	frequency := attrs[fields.Frequency].(types.String)
	testutil.AssertEqual(t, "frequency", frequency.ValueString(), "1 minute")

	// Verify node distribution
	nodeDistribution := attrs[fields.NodeDistribution].(types.String)
	testutil.AssertEqual(t, "node distribution", nodeDistribution.ValueString(), "concurrent")

	// Verify node IDs list
	nodeIDsList := attrs[fields.NodeIDs].(types.List)
	testutil.AssertEqual(t, "node ids should not be null", nodeIDsList.IsNull(), false)

	nodeIDsElements := nodeIDsList.Elements()
	testutil.AssertEqual(t, "should have 3 node ids", len(nodeIDsElements), 3)
	testutil.AssertEqual(t, "first node id", nodeIDsElements[0].(types.Int64).ValueInt64(), int64(1001))
	testutil.AssertEqual(t, "second node id", nodeIDsElements[1].(types.Int64).ValueInt64(), int64(1002))
	testutil.AssertEqual(t, "third node id", nodeIDsElements[2].(types.Int64).ValueInt64(), int64(1003))

	// Verify node group IDs list
	nodeGroupIDsList := attrs[fields.NodeGroupIDs].(types.List)
	testutil.AssertEqual(t, "node group ids should not be null", nodeGroupIDsList.IsNull(), false)

	nodeGroupIDsElements := nodeGroupIDsList.Elements()
	testutil.AssertEqual(t, "should have 2 node group ids", len(nodeGroupIDsElements), 2)
	testutil.AssertEqual(t, "first node group id", nodeGroupIDsElements[0].(types.Int64).ValueInt64(), int64(2001))
	testutil.AssertEqual(t, "second node group id", nodeGroupIDsElements[1].(types.Int64).ValueInt64(), int64(2002))

	// Verify number of subset nodes
	noOfSubsetNodes := attrs[fields.NoOfSubsetNodes].(types.Int64)
	testutil.AssertEqual(t, "number of subset nodes", noOfSubsetNodes.ValueInt64(), int64(5))

	// Verify all expected fields are present
	expectedFieldCount := 8
	testutil.AssertEqual(t, "should have all expected fields", len(attrs), expectedFieldCount)
}
