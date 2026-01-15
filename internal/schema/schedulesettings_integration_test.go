package schema

import (
	"context"
	"reflect"
	"sync"
	"testing"

	"catchpoint-provider/internal/fields"
	"catchpoint-provider/internal/testutil"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func TestBuildScheduleAttributes(t *testing.T) {
	ctx := context.Background()

	attrs := BuildScheduleSettingsBlock(ctx)

	// Should have exactly one attribute
	testutil.AssertEqual(t, "should have one attribute", len(attrs), 1)

	// Should contain ScheduleSettings
	scheduleSettings, exists := attrs[fields.ScheduleSettings]
	testutil.AssertEqual(t, "should contain ScheduleSettings", exists, true)

	// Should be SingleNestedBlock
	singleNested, ok := scheduleSettings.(schema.SingleNestedBlock)
	testutil.AssertEqual(t, "should be SingleNestedBlock", ok, true)

	// Verify nested attributes
	nestedAttrs := singleNested.Attributes
	expectedFields := []string{
		fields.ScheduleSettingType,
		fields.RunScheduleID,
		fields.MaintenanceScheduleID,
		fields.Frequency,
		fields.NodeDistribution,
		fields.NodeIDs,
		fields.NodeGroupIDs,
		fields.NoOfSubsetNodes,
	}

	testutil.AssertEqual(t, "should have correct number of nested attributes", len(nestedAttrs), len(expectedFields))

	for _, field := range expectedFields {
		_, exists := nestedAttrs[field]
		testutil.AssertEqual(t, "should contain field "+field, exists, true)
	}
}

func TestBuildScheduleAttributesScheduleSettingType(t *testing.T) {
	ctx := context.Background()

	attrs := BuildScheduleSettingsBlock(ctx)
	scheduleSettings := attrs[fields.ScheduleSettings].(schema.SingleNestedBlock)

	scheduleSettingType := scheduleSettings.Attributes[fields.ScheduleSettingType].(schema.StringAttribute)

	// Should not be optional or required (computed only)
	testutil.AssertEqual(t, "should be optional", scheduleSettingType.Optional, true)
	testutil.AssertEqual(t, "should not be required", scheduleSettingType.Required, false)
	testutil.AssertEqual(t, "should be computed", scheduleSettingType.Computed, true)

	// Should have validators
	testutil.AssertEqual(t, "should have validators", len(scheduleSettingType.Validators), 1)

	// Should have description
	testutil.AssertStringContains(t, "should contain setting type in description", scheduleSettingType.Description, "schedule setting type")
}

func TestBuildScheduleAttributesRequiredFields(t *testing.T) {
	ctx := context.Background()

	attrs := BuildScheduleSettingsBlock(ctx)
	scheduleSettings := attrs[fields.ScheduleSettings].(schema.SingleNestedBlock)

	// Frequency should be required
	frequency := scheduleSettings.Attributes[fields.Frequency].(schema.StringAttribute)
	testutil.AssertEqual(t, "frequency should not be required", frequency.Required, false)
	testutil.AssertEqual(t, "frequency should be optional", frequency.Optional, true)
	testutil.AssertEqual(t, "frequency should have validators", len(frequency.Validators), 1)
	testutil.AssertStringContains(t, "frequency description should mention seconds", frequency.Description, "seconds")

	// NodeDistribution should be required
	nodeDistribution := scheduleSettings.Attributes[fields.NodeDistribution].(schema.StringAttribute)
	testutil.AssertEqual(t, "node distribution should not be required", nodeDistribution.Required, false)
	testutil.AssertEqual(t, "node distribution should be optional", nodeDistribution.Optional, true)
	testutil.AssertEqual(t, "node distribution should have validators", len(nodeDistribution.Validators), 1)
	testutil.AssertStringContains(t, "node distribution description should mention distribution", nodeDistribution.Description, "distribution")
}

func TestBuildScheduleAttributesOptionalFields(t *testing.T) {
	ctx := context.Background()

	attrs := BuildScheduleSettingsBlock(ctx)
	scheduleSettings := attrs[fields.ScheduleSettings].(schema.SingleNestedBlock)

	// RunScheduleID should be optional
	runScheduleID := scheduleSettings.Attributes[fields.RunScheduleID].(schema.Int64Attribute)
	testutil.AssertEqual(t, "run schedule id should be optional", runScheduleID.Optional, true)
	testutil.AssertEqual(t, "run schedule id should not be required", runScheduleID.Required, false)
	testutil.AssertStringContains(t, "run schedule id description should mention Run Schedule", runScheduleID.Description, "Run Schedule")

	// MaintenanceScheduleID should be optional
	maintenanceScheduleID := scheduleSettings.Attributes[fields.MaintenanceScheduleID].(schema.Int64Attribute)
	testutil.AssertEqual(t, "maintenance schedule id should be optional", maintenanceScheduleID.Optional, true)
	testutil.AssertEqual(t, "maintenance schedule id should not be required", maintenanceScheduleID.Required, false)
	testutil.AssertStringContains(t, "maintenance schedule id description should mention Maintenance Schedule", maintenanceScheduleID.Description, "Maintenance Schedule")

	// NoOfSubsetNodes should be optional
	noOfSubsetNodes := scheduleSettings.Attributes[fields.NoOfSubsetNodes].(schema.Int64Attribute)
	testutil.AssertEqual(t, "number of subset nodes should be optional", noOfSubsetNodes.Optional, true)
	testutil.AssertEqual(t, "number of subset nodes should not be required", noOfSubsetNodes.Required, false)
	testutil.AssertStringContains(t, "number of subset nodes description should mention subset", noOfSubsetNodes.Description, "subset")
}

func TestBuildScheduleAttributesListFields(t *testing.T) {
	ctx := context.Background()

	attrs := BuildScheduleSettingsBlock(ctx)
	scheduleSettings := attrs[fields.ScheduleSettings].(schema.SingleNestedBlock)

	// NodeIDs should be list of Int64
	nodeIDs := scheduleSettings.Attributes[fields.NodeIDs].(schema.ListAttribute)
	testutil.AssertEqual(t, "node ids should be optional", nodeIDs.Optional, true)
	testutil.AssertEqual(t, "node ids should not be required", nodeIDs.Required, false)
	// Compare element type using reflect.DeepEqual to avoid type mismatch error
	testutil.AssertEqual(t, "node ids element type should be Int64", reflect.DeepEqual(nodeIDs.ElementType, types.Int64Type), true)
	testutil.AssertStringContains(t, "node ids description should mention Node IDs", nodeIDs.Description, "Node IDs")

	// NodeGroupIDs should be list of Int64
	nodeGroupIDs := scheduleSettings.Attributes[fields.NodeGroupIDs].(schema.ListAttribute)
	testutil.AssertEqual(t, "node group ids should be optional", nodeGroupIDs.Optional, true)
	testutil.AssertEqual(t, "node group ids should not be required", nodeGroupIDs.Required, false)
	testutil.AssertEqual(t, "node group ids element type should be Int64", reflect.DeepEqual(nodeGroupIDs.ElementType, types.Int64Type), true)
	testutil.AssertStringContains(t, "node group ids description should mention Node Group IDs", nodeGroupIDs.Description, "Node Group IDs")
}

func TestGetScheduleSettingsAttributeTypes(t *testing.T) {
	attrTypes := GetScheduleSettingsAttributeTypes()

	// Should return map of attribute types
	testutil.AssertNotNil(t, "attribute types should not be nil", attrTypes)
	testutil.AssertEqual(t, "should have correct number of attributes", len(attrTypes), 8)

	// Verify all expected fields exist
	expectedFields := []string{
		fields.ScheduleSettingType,
		fields.RunScheduleID,
		fields.MaintenanceScheduleID,
		fields.Frequency,
		fields.NodeDistribution,
		fields.NodeIDs,
		fields.NodeGroupIDs,
		fields.NoOfSubsetNodes,
	}

	for _, field := range expectedFields {
		attrType, exists := attrTypes[field]
		testutil.AssertEqual(t, "should contain field "+field, exists, true)
		testutil.AssertNotNil(t, "attribute type should not be nil for "+field, attrType)
	}

	// Verify specific attribute types
	testutil.AssertEqual(t, "schedule setting type should be StringType", reflect.DeepEqual(attrTypes[fields.ScheduleSettingType], types.StringType), true)
	testutil.AssertEqual(t, "run schedule id should be Int64Type", reflect.DeepEqual(attrTypes[fields.RunScheduleID], types.Int64Type), true)
	testutil.AssertEqual(t, "maintenance schedule id should be Int64Type", reflect.DeepEqual(attrTypes[fields.MaintenanceScheduleID], types.Int64Type), true)
	testutil.AssertEqual(t, "frequency should be StringType", reflect.DeepEqual(attrTypes[fields.Frequency], types.StringType), true)
	testutil.AssertEqual(t, "node distribution should be StringType", reflect.DeepEqual(attrTypes[fields.NodeDistribution], types.StringType), true)
	testutil.AssertEqual(t, "number of subset nodes should be Int64Type", reflect.DeepEqual(attrTypes[fields.NoOfSubsetNodes], types.Int64Type), true)

	// Verify list types
	expectedNodeIDsType := types.ListType{ElemType: types.Int64Type}
	testutil.AssertEqual(t, "node ids should be ListType with Int64 elements", reflect.DeepEqual(attrTypes[fields.NodeIDs], expectedNodeIDsType), true)

	expectedNodeGroupIDsType := types.ListType{ElemType: types.Int64Type}
	testutil.AssertEqual(t, "node group ids should be ListType with Int64 elements", reflect.DeepEqual(attrTypes[fields.NodeGroupIDs], expectedNodeGroupIDsType), true)
}

func TestGetScheduleSettingsAttributeTypesConsistency(t *testing.T) {
	// Get attribute types multiple times to ensure consistency
	attrTypes1 := GetScheduleSettingsAttributeTypes()
	attrTypes2 := GetScheduleSettingsAttributeTypes()

	testutil.AssertEqual(t, "should return same number of attributes", len(attrTypes1), len(attrTypes2))

	// Verify all fields are the same
	for field, attrType1 := range attrTypes1 {
		attrType2, exists := attrTypes2[field]
		testutil.AssertEqual(t, "field should exist in both calls for "+field, exists, true)
		testutil.AssertEqual(t, "attribute types should be equal for "+field, attrType1, attrType2)
	}
}

func TestGetScheduleSettingsSchemaOnceInitialization(t *testing.T) {
	// Reset the once to test initialization
	scheduleSettingsSchemaOnce = sync.Once{}

	// First call should initialize
	schema1 := getScheduleSettingsSchema()
	testutil.AssertNotNil(t, "schema should not be nil", schema1)

	// Second call should return same instance
	schema2 := getScheduleSettingsSchema()
	testutil.AssertEqual(t, "should return same schema instance", schema1.Description, schema2.Description)
	testutil.AssertEqual(t, "should have same number of attributes", len(schema1.Attributes), len(schema2.Attributes))
}

func TestGetScheduleSettingsSchemaStructure(t *testing.T) {
	schema := getScheduleSettingsSchema()

	// Should have correct number of attributes
	testutil.AssertEqual(t, "should have 8 attributes", len(schema.Attributes), 8)
}

func TestBuildScheduleAttributesFieldValidation(t *testing.T) {
	ctx := context.Background()

	attrs := BuildScheduleSettingsBlock(ctx)
	scheduleSettings := attrs[fields.ScheduleSettings].(schema.SingleNestedBlock)

	// Test that all required fields have proper validation
	frequency := scheduleSettings.Attributes[fields.Frequency].(schema.StringAttribute)
	testutil.AssertEqual(t, "frequency should have exactly one validator", len(frequency.Validators), 1)

	nodeDistribution := scheduleSettings.Attributes[fields.NodeDistribution].(schema.StringAttribute)
	testutil.AssertEqual(t, "node distribution should have exactly one validator", len(nodeDistribution.Validators), 1)

	scheduleSettingType := scheduleSettings.Attributes[fields.ScheduleSettingType].(schema.StringAttribute)
	testutil.AssertEqual(t, "schedule setting type should have exactly one validator", len(scheduleSettingType.Validators), 1)

	// Optional fields should not have validators (except for basic type validation)
	runScheduleID := scheduleSettings.Attributes[fields.RunScheduleID].(schema.Int64Attribute)
	testutil.AssertEqual(t, "run schedule id should have no validators", len(runScheduleID.Validators), 0)

	maintenanceScheduleID := scheduleSettings.Attributes[fields.MaintenanceScheduleID].(schema.Int64Attribute)
	testutil.AssertEqual(t, "maintenance schedule id should have no validators", len(maintenanceScheduleID.Validators), 0)
}

func TestBuildScheduleAttributesDescriptions(t *testing.T) {
	ctx := context.Background()

	attrs := BuildScheduleSettingsBlock(ctx)
	scheduleSettings := attrs[fields.ScheduleSettings].(schema.SingleNestedBlock)

	// Verify all attributes have meaningful descriptions
	for fieldName, attr := range scheduleSettings.Attributes {
		var description string

		switch a := attr.(type) {
		case schema.StringAttribute:
			description = a.Description
		case schema.Int64Attribute:
			description = a.Description
		case schema.ListAttribute:
			description = a.Description
		}

		testutil.AssertNotEqual(t, "description should not be empty for "+fieldName, description, "")
	}
}
