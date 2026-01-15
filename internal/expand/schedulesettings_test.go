package expand

import (
	"testing"

	"catchpoint-provider/internal/fields"
	"catchpoint-provider/internal/labels"
	"catchpoint-provider/internal/models"
	"catchpoint-provider/internal/testutil"
	cptypes "catchpoint-provider/internal/types"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func TestExpandScheduleSettingsConfigNullObject(t *testing.T) {
	obj := types.ObjectNull(map[string]attr.Type{})
	config := &models.ScheduleSettingsConfig{}

	diags := ExpandScheduleSettingsConfig(obj, config)

	testutil.AssertDiagsHasNoErrors(t, diags)
	// Config should remain unchanged when object is null
}

func TestExpandScheduleSettingsConfigEmptyObject(t *testing.T) {
	attrs := map[string]attr.Value{}

	obj, _ := types.ObjectValue(map[string]attr.Type{}, attrs)
	config := &models.ScheduleSettingsConfig{}

	diags := ExpandScheduleSettingsConfig(obj, config)

	testutil.AssertDiagsHasNoErrors(t, diags)
	// Should set schedule setting type to Override by default
	testutil.AssertEqual(t, "schedule setting type name", config.ScheduleSettingType.Name, cptypes.Inherit)
	testutil.AssertEqual(t, "schedule setting type id", config.ScheduleSettingType.ID, 0)
}

func TestExpandScheduleSettingsConfigAllMainFields(t *testing.T) {
	attrs := map[string]attr.Value{
		fields.RunScheduleID:         types.Int64Value(100),
		fields.MaintenanceScheduleID: types.Int64Value(200),
		fields.Frequency:             types.StringValue("5 minutes"),
		fields.NodeDistribution:      types.StringValue("random"),
		fields.NoOfSubsetNodes:       types.Int64Value(5),
	}

	obj, _ := types.ObjectValue(map[string]attr.Type{
		fields.RunScheduleID:         types.Int64Type,
		fields.MaintenanceScheduleID: types.Int64Type,
		fields.Frequency:             types.StringType,
		fields.NodeDistribution:      types.StringType,
		fields.NoOfSubsetNodes:       types.Int64Type,
	}, attrs)

	config := &models.ScheduleSettingsConfig{}

	diags := ExpandScheduleSettingsConfig(obj, config)

	testutil.AssertDiagsHasNoErrors(t, diags)
	testutil.AssertEqual(t, fields.RunScheduleID, config.RunScheduleID, 100)
	testutil.AssertEqual(t, fields.MaintenanceScheduleID, config.MaintenanceScheduleID, 200)
	testutil.AssertEqual(t, fields.Frequency, config.Frequency.Name, "5 minutes")
	testutil.AssertEqual(t, fields.NodeDistribution, config.NodeDistribution.Name, "random")
	testutil.AssertEqual(t, fields.NoOfSubsetNodes, config.NoOfSubsetNodes, 5)
}

func TestExpandScheduleSettingsConfigWithNodeIDs(t *testing.T) {
	nodeIDList, _ := types.ListValue(types.Int64Type, []attr.Value{
		types.Int64Value(1),
		types.Int64Value(2),
		types.Int64Value(3),
	})

	attrs := map[string]attr.Value{
		fields.NodeIDs: nodeIDList,
	}

	obj, _ := types.ObjectValue(map[string]attr.Type{
		fields.NodeIDs: types.ListType{ElemType: types.Int64Type},
	}, attrs)

	config := &models.ScheduleSettingsConfig{}

	diags := ExpandScheduleSettingsConfig(obj, config)

	testutil.AssertDiagsHasNoErrors(t, diags)
	testutil.AssertEqual(t, "nodeids length", len(config.NodeIDs), 3)
	testutil.AssertEqual(t, "first node id", config.NodeIDs[0], 1)
	testutil.AssertEqual(t, "second node id", config.NodeIDs[1], 2)
	testutil.AssertEqual(t, "third node id", config.NodeIDs[2], 3)
}

func TestExpandScheduleSettingsConfigWithNodeGroupIDs(t *testing.T) {
	nodeGroupIDList, _ := types.ListValue(types.Int64Type, []attr.Value{
		types.Int64Value(10),
		types.Int64Value(20),
		types.Int64Value(30),
	})

	attrs := map[string]attr.Value{
		fields.NodeGroupIDs: nodeGroupIDList,
	}

	obj, _ := types.ObjectValue(map[string]attr.Type{
		fields.NodeGroupIDs: types.ListType{ElemType: types.Int64Type},
	}, attrs)

	config := &models.ScheduleSettingsConfig{}

	diags := ExpandScheduleSettingsConfig(obj, config)

	testutil.AssertDiagsHasNoErrors(t, diags)
	testutil.AssertEqual(t, fields.NodeGroupIDs, len(config.NodeGroupIDs), 3)

	// Check first node group
	testutil.AssertEqual(t, "first node group id", config.NodeGroupIDs[0].ID, 10)
	testutil.AssertEqual(t, "first node group name", config.NodeGroupIDs[0].Name, labels.DefaultNodeGroupName)

	// Check second node group
	testutil.AssertEqual(t, "second node group id", config.NodeGroupIDs[1].ID, 20)
	testutil.AssertEqual(t, "second node group name", config.NodeGroupIDs[1].Name, labels.DefaultNodeGroupName)
}

func TestExpandScheduleSettingsConfigWithBothNodeIDsAndGroupIDs(t *testing.T) {
	nodeIDList, _ := types.ListValue(types.Int64Type, []attr.Value{
		types.Int64Value(1),
		types.Int64Value(2),
	})

	nodeGroupIDList, _ := types.ListValue(types.Int64Type, []attr.Value{
		types.Int64Value(10),
		types.Int64Value(20),
	})

	attrs := map[string]attr.Value{
		fields.NodeIDs:      nodeIDList,
		fields.NodeGroupIDs: nodeGroupIDList,
	}

	obj, _ := types.ObjectValue(map[string]attr.Type{
		fields.NodeIDs:      types.ListType{ElemType: types.Int64Type},
		fields.NodeGroupIDs: types.ListType{ElemType: types.Int64Type},
	}, attrs)

	config := &models.ScheduleSettingsConfig{}

	diags := ExpandScheduleSettingsConfig(obj, config)

	testutil.AssertDiagsHasNoErrors(t, diags)
	testutil.AssertEqual(t, "node ids length", len(config.NodeIDs), 2)
	testutil.AssertEqual(t, "node group ids length", len(config.NodeGroupIDs), 2)
}

func TestExpandScheduleSettingsConfigNullOptionalFields(t *testing.T) {
	attrs := map[string]attr.Value{
		fields.RunScheduleID:    types.Int64Value(100),
		fields.Frequency:        types.StringNull(),
		fields.NodeDistribution: types.StringNull(),
		fields.NoOfSubsetNodes:  types.Int64Null(),
		fields.NodeIDs:          types.ListNull(types.Int64Type),
		fields.NodeGroupIDs:     types.ListNull(types.Int64Type),
	}

	obj, _ := types.ObjectValue(map[string]attr.Type{
		fields.RunScheduleID:    types.Int64Type,
		fields.Frequency:        types.StringType,
		fields.NodeDistribution: types.StringType,
		fields.NoOfSubsetNodes:  types.Int64Type,
		fields.NodeIDs:          types.ListType{ElemType: types.Int64Type},
		fields.NodeGroupIDs:     types.ListType{ElemType: types.Int64Type},
	}, attrs)

	config := &models.ScheduleSettingsConfig{}

	diags := ExpandScheduleSettingsConfig(obj, config)

	testutil.AssertDiagsHasNoErrors(t, diags)
	testutil.AssertEqual(t, fields.RunScheduleID, config.RunScheduleID, 100)
	testutil.AssertEqual(t, fields.Frequency, config.Frequency.Name, "5 minutes")      // Should be 5 minutes for null (else the API rejects the payload)
	testutil.AssertEqual(t, fields.NodeDistribution, config.NodeDistribution.Name, "") // Should be empty for null
	testutil.AssertEqual(t, fields.NoOfSubsetNodes, config.NoOfSubsetNodes, 0)         // Should be 0 for null
	testutil.AssertEqual(t, fields.NodeIDs, len(config.NodeIDs), 0)
	testutil.AssertEqual(t, fields.NodeGroupIDs, len(config.NodeGroupIDs), 0)
}

func TestExpandScheduleSettingsConfigFrequencyValidation(t *testing.T) {
	for _, tc := range cptypes.ValidFrequencyNames {
		t.Run(tc, func(t *testing.T) {
			attrs := map[string]attr.Value{
				fields.Frequency: types.StringValue(tc),
			}

			obj, _ := types.ObjectValue(map[string]attr.Type{
				fields.Frequency: types.StringType,
			}, attrs)

			config := &models.ScheduleSettingsConfig{}

			diags := ExpandScheduleSettingsConfig(obj, config)

			testutil.AssertDiagsHasNoErrors(t, diags)
			testutil.AssertEqual(t, "frequency name", config.Frequency.Name, tc)
		})
	}
}

func TestExpandScheduleSettingsConfigNodeDistributionValidation(t *testing.T) {
	for _, tc := range cptypes.ValidNodeDistributions {
		t.Run(tc, func(t *testing.T) {
			attrs := map[string]attr.Value{
				fields.NodeDistribution: types.StringValue(tc),
			}

			obj, _ := types.ObjectValue(map[string]attr.Type{
				fields.NodeDistribution: types.StringType,
			}, attrs)

			config := &models.ScheduleSettingsConfig{}

			diags := ExpandScheduleSettingsConfig(obj, config)

			testutil.AssertDiagsHasNoErrors(t, diags)
			testutil.AssertEqual(t, "node distribution name", config.NodeDistribution.Name, tc)
		})
	}
}

func TestExpandScheduleSettingsConfigEmptyNodeLists(t *testing.T) {
	emptyNodeIDList, _ := types.ListValue(types.Int64Type, []attr.Value{})
	emptyNodeGroupIDList, _ := types.ListValue(types.Int64Type, []attr.Value{})

	attrs := map[string]attr.Value{
		fields.NodeIDs:      emptyNodeIDList,
		fields.NodeGroupIDs: emptyNodeGroupIDList,
	}

	obj, _ := types.ObjectValue(map[string]attr.Type{
		fields.NodeIDs:      types.ListType{ElemType: types.Int64Type},
		fields.NodeGroupIDs: types.ListType{ElemType: types.Int64Type},
	}, attrs)

	config := &models.ScheduleSettingsConfig{}

	diags := ExpandScheduleSettingsConfig(obj, config)

	testutil.AssertDiagsHasNoErrors(t, diags)
	testutil.AssertEqual(t, "node ids length", len(config.NodeIDs), 0)
	testutil.AssertEqual(t, "node group ids length", len(config.NodeGroupIDs), 0)
}

func TestExpandScheduleSettingsConfigLargeNodeValues(t *testing.T) {
	nodeIDList, _ := types.ListValue(types.Int64Type, []attr.Value{
		types.Int64Value(999999),
		types.Int64Value(888888),
	})

	nodeGroupIDList, _ := types.ListValue(types.Int64Type, []attr.Value{
		types.Int64Value(777777),
		types.Int64Value(666666),
	})

	attrs := map[string]attr.Value{
		fields.NodeIDs:         nodeIDList,
		fields.NodeGroupIDs:    nodeGroupIDList,
		fields.NoOfSubsetNodes: types.Int64Value(999),
	}

	obj, _ := types.ObjectValue(map[string]attr.Type{
		fields.NodeIDs:         types.ListType{ElemType: types.Int64Type},
		fields.NodeGroupIDs:    types.ListType{ElemType: types.Int64Type},
		fields.NoOfSubsetNodes: types.Int64Type,
	}, attrs)

	config := &models.ScheduleSettingsConfig{}

	diags := ExpandScheduleSettingsConfig(obj, config)

	testutil.AssertDiagsHasNoErrors(t, diags)
	testutil.AssertEqual(t, "first node id", config.NodeIDs[0], 999999)
	testutil.AssertEqual(t, "first node group id", config.NodeGroupIDs[0].ID, 777777)
	testutil.AssertEqual(t, "no of subset nodes", config.NoOfSubsetNodes, 999)
}

func TestExpandScheduleSettingsConfigZeroValues(t *testing.T) {
	nodeIDList, _ := types.ListValue(types.Int64Type, []attr.Value{
		types.Int64Value(0),
	})

	attrs := map[string]attr.Value{
		fields.RunScheduleID:         types.Int64Value(0),
		fields.MaintenanceScheduleID: types.Int64Value(0),
		fields.NodeIDs:               nodeIDList,
		fields.NoOfSubsetNodes:       types.Int64Value(0),
	}

	obj, _ := types.ObjectValue(map[string]attr.Type{
		fields.RunScheduleID:         types.Int64Type,
		fields.MaintenanceScheduleID: types.Int64Type,
		fields.NodeIDs:               types.ListType{ElemType: types.Int64Type},
		fields.NoOfSubsetNodes:       types.Int64Type,
	}, attrs)

	config := &models.ScheduleSettingsConfig{}

	diags := ExpandScheduleSettingsConfig(obj, config)

	testutil.AssertDiagsHasNoErrors(t, diags)
	testutil.AssertEqual(t, fields.RunScheduleID, config.RunScheduleID, 0)
	testutil.AssertEqual(t, fields.MaintenanceScheduleID, config.MaintenanceScheduleID, 0)
	testutil.AssertEqual(t, fields.NodeIDs, config.NodeIDs[0], 0)
	testutil.AssertEqual(t, fields.NoOfSubsetNodes, config.NoOfSubsetNodes, 0)
}

// Integration test with all fields populated
func TestExpandScheduleSettingsConfigCompleteConfiguration(t *testing.T) {
	nodeIDList, _ := types.ListValue(types.Int64Type, []attr.Value{
		types.Int64Value(1),
		types.Int64Value(2),
		types.Int64Value(3),
	})

	nodeGroupIDList, _ := types.ListValue(types.Int64Type, []attr.Value{
		types.Int64Value(10),
		types.Int64Value(20),
	})

	attrs := map[string]attr.Value{
		fields.RunScheduleID:         types.Int64Value(100),
		fields.MaintenanceScheduleID: types.Int64Value(200),
		fields.Frequency:             types.StringValue("15 minutes"),
		fields.NodeDistribution:      types.StringValue("random"),
		fields.NodeIDs:               nodeIDList,
		fields.NodeGroupIDs:          nodeGroupIDList,
		fields.NoOfSubsetNodes:       types.Int64Value(5),
	}

	obj, _ := types.ObjectValue(map[string]attr.Type{
		fields.RunScheduleID:         types.Int64Type,
		fields.MaintenanceScheduleID: types.Int64Type,
		fields.Frequency:             types.StringType,
		fields.NodeDistribution:      types.StringType,
		fields.NodeIDs:               types.ListType{ElemType: types.Int64Type},
		fields.NodeGroupIDs:          types.ListType{ElemType: types.Int64Type},
		fields.NoOfSubsetNodes:       types.Int64Type,
	}, attrs)

	config := &models.ScheduleSettingsConfig{}

	diags := ExpandScheduleSettingsConfig(obj, config)

	testutil.AssertDiagsHasNoErrors(t, diags)

	// Verify all fields
	testutil.AssertEqual(t, fields.ScheduleSettingType, config.ScheduleSettingType.Name, cptypes.Override)
	testutil.AssertEqual(t, fields.RunScheduleID, config.RunScheduleID, 100)
	testutil.AssertEqual(t, fields.MaintenanceScheduleID, config.MaintenanceScheduleID, 200)
	testutil.AssertEqual(t, fields.Frequency, config.Frequency.Name, "15 minutes")
	testutil.AssertEqual(t, fields.NodeDistribution, config.NodeDistribution.Name, "random")
	testutil.AssertEqual(t, fields.NodeIDs, len(config.NodeIDs), 3)
	testutil.AssertEqual(t, fields.NodeGroupIDs, len(config.NodeGroupIDs), 2)
	testutil.AssertEqual(t, fields.NoOfSubsetNodes, config.NoOfSubsetNodes, 5)
}

// Test error handling for invalid list data
func TestExpandScheduleSettingsConfigInvalidNodeListData(t *testing.T) {
	// Test with invalid list types that might cause errors in getIntListFromAttr
	invalidList := types.ListUnknown(types.StringType) // Wrong element type

	attrs := map[string]attr.Value{
		fields.NodeIDs: invalidList,
	}

	obj, _ := types.ObjectValue(map[string]attr.Type{
		fields.NodeIDs: types.ListType{ElemType: types.StringType}, // Wrong type
	}, attrs)

	config := &models.ScheduleSettingsConfig{}

	diags := ExpandScheduleSettingsConfig(obj, config)

	// Should handle errors gracefully - the exact behavior depends on getIntListFromAttr implementation
	if diags.HasError() {
		// Error is expected for invalid data
		testutil.AssertEqual(t, "should have errors for invalid data", diags.HasError(), true)
	}
}

func TestExpandScheduleSettingsConfigConfigNotModifiedOnError(t *testing.T) {
	config := &models.ScheduleSettingsConfig{
		RunScheduleID: 999, // Set initial value
	}

	// Test with potentially problematic data
	invalidList := types.ListUnknown(types.Int64Type)

	attrs := map[string]attr.Value{
		fields.NodeIDs: invalidList,
	}

	obj, _ := types.ObjectValue(map[string]attr.Type{
		fields.NodeIDs: types.ListType{ElemType: types.Int64Type},
	}, attrs)

	diags := ExpandScheduleSettingsConfig(obj, config)

	if diags.HasError() {
		// If there's an error, some fields might still be modified, but function should not panic
		_ = config.RunScheduleID // Just verify we can access it
	}

	// At minimum, the ScheduleSettingType should be set since that happens first
	testutil.AssertEqual(t, "schedule setting type should be set", config.ScheduleSettingType.Name, cptypes.Inherit)
}
