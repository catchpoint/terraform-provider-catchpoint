package expand

import (
	"testing"

	"catchpoint-provider/internal/fields"
	"catchpoint-provider/internal/models"
	"catchpoint-provider/internal/testutil"
	cptypes "catchpoint-provider/internal/types"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func TestExpandInsightSettingsConfigNullObject(t *testing.T) {
	obj := types.ObjectNull(map[string]attr.Type{})
	config := &models.InsightSettingsConfig{}

	diags := ExpandInsightSettingsConfig(obj, config)

	testutil.AssertDiagsHasNoErrors(t, diags)
	// Config should remain unchanged when object is null
}

func TestExpandInsightSettingsConfigEmptyObject(t *testing.T) {
	attrs := map[string]attr.Value{}

	obj, _ := types.ObjectValue(map[string]attr.Type{}, attrs)
	config := &models.InsightSettingsConfig{}

	diags := ExpandInsightSettingsConfig(obj, config)

	testutil.AssertDiagsHasNoErrors(t, diags)
	// Should set insight setting type to Inherit by default
	testutil.AssertEqual(t, "insight setting type name", config.InsightSettingType.Name, cptypes.Inherit)
	testutil.AssertEqual(t, "insight setting type id", config.InsightSettingType.ID, 0)
	testutil.AssertEqual(t, "tracepoint ids length", len(config.TracepointIDs), 0)
	testutil.AssertEqual(t, "indicator ids length", len(config.IndicatorIDs), 0)
}

func TestExpandInsightSettingsConfigTracepointIDsOnly(t *testing.T) {
	tracepointIDList, _ := types.ListValue(types.Int64Type, []attr.Value{
		types.Int64Value(1),
		types.Int64Value(2),
		types.Int64Value(3),
	})

	attrs := map[string]attr.Value{
		fields.TracepointIDs: tracepointIDList,
	}

	obj, _ := types.ObjectValue(map[string]attr.Type{
		fields.TracepointIDs: types.ListType{ElemType: types.Int64Type},
	}, attrs)

	config := &models.InsightSettingsConfig{}

	diags := ExpandInsightSettingsConfig(obj, config)

	testutil.AssertDiagsHasNoErrors(t, diags)
	testutil.AssertEqual(t, "insight setting type name", config.InsightSettingType.Name, cptypes.Override)
	testutil.AssertEqual(t, "tracepoint ids length", len(config.TracepointIDs), 3)
	testutil.AssertEqual(t, "first tracepoint id", config.TracepointIDs[0], 1)
	testutil.AssertEqual(t, "second tracepoint id", config.TracepointIDs[1], 2)
	testutil.AssertEqual(t, "third tracepoint id", config.TracepointIDs[2], 3)
	testutil.AssertEqual(t, "indicator ids length", len(config.IndicatorIDs), 0)
}

func TestExpandInsightSettingsConfigIndicatorIDsOnly(t *testing.T) {
	indicatorIDList, _ := types.ListValue(types.Int64Type, []attr.Value{
		types.Int64Value(10),
		types.Int64Value(20),
		types.Int64Value(30),
	})

	attrs := map[string]attr.Value{
		fields.IndicatorIDs: indicatorIDList,
	}

	obj, _ := types.ObjectValue(map[string]attr.Type{
		fields.IndicatorIDs: types.ListType{ElemType: types.Int64Type},
	}, attrs)

	config := &models.InsightSettingsConfig{}

	diags := ExpandInsightSettingsConfig(obj, config)

	testutil.AssertDiagsHasNoErrors(t, diags)
	testutil.AssertEqual(t, "insight setting type name", config.InsightSettingType.Name, cptypes.Override)
	testutil.AssertEqual(t, "indicator ids length", len(config.IndicatorIDs), 3)
	testutil.AssertEqual(t, "first indicator id", config.IndicatorIDs[0], 10)
	testutil.AssertEqual(t, "second indicator id", config.IndicatorIDs[1], 20)
	testutil.AssertEqual(t, "third indicator id", config.IndicatorIDs[2], 30)
	testutil.AssertEqual(t, "tracepoint ids length", len(config.TracepointIDs), 0)
}

func TestExpandInsightSettingsConfigBothIDsPresent(t *testing.T) {
	tracepointIDList, _ := types.ListValue(types.Int64Type, []attr.Value{
		types.Int64Value(1),
		types.Int64Value(2),
	})

	indicatorIDList, _ := types.ListValue(types.Int64Type, []attr.Value{
		types.Int64Value(10),
		types.Int64Value(20),
	})

	attrs := map[string]attr.Value{
		fields.TracepointIDs: tracepointIDList,
		fields.IndicatorIDs:  indicatorIDList,
	}

	obj, _ := types.ObjectValue(map[string]attr.Type{
		fields.TracepointIDs: types.ListType{ElemType: types.Int64Type},
		fields.IndicatorIDs:  types.ListType{ElemType: types.Int64Type},
	}, attrs)

	config := &models.InsightSettingsConfig{}

	diags := ExpandInsightSettingsConfig(obj, config)

	testutil.AssertDiagsHasNoErrors(t, diags)
	testutil.AssertEqual(t, "insight setting type name", config.InsightSettingType.Name, cptypes.Override)
	testutil.AssertEqual(t, "tracepoint ids length", len(config.TracepointIDs), 2)
	testutil.AssertEqual(t, "indicator ids length", len(config.IndicatorIDs), 2)
	testutil.AssertEqual(t, "first tracepoint id", config.TracepointIDs[0], 1)
	testutil.AssertEqual(t, "first indicator id", config.IndicatorIDs[0], 10)
}

func TestExpandInsightSettingsConfigNullLists(t *testing.T) {
	attrs := map[string]attr.Value{
		fields.TracepointIDs: types.ListNull(types.Int64Type),
		fields.IndicatorIDs:  types.ListNull(types.Int64Type),
	}

	obj, _ := types.ObjectValue(map[string]attr.Type{
		fields.TracepointIDs: types.ListType{ElemType: types.Int64Type},
		fields.IndicatorIDs:  types.ListType{ElemType: types.Int64Type},
	}, attrs)

	config := &models.InsightSettingsConfig{}

	diags := ExpandInsightSettingsConfig(obj, config)

	testutil.AssertDiagsHasNoErrors(t, diags)
	testutil.AssertEqual(t, "insight setting type name", config.InsightSettingType.Name, cptypes.Inherit)
	testutil.AssertEqual(t, "tracepoint ids length", len(config.TracepointIDs), 0)
	testutil.AssertEqual(t, "indicator ids length", len(config.IndicatorIDs), 0)
}

func TestExpandInsightSettingsConfigEmptyLists(t *testing.T) {
	emptyTracepointList, _ := types.ListValue(types.Int64Type, []attr.Value{})
	emptyIndicatorList, _ := types.ListValue(types.Int64Type, []attr.Value{})

	attrs := map[string]attr.Value{
		fields.TracepointIDs: emptyTracepointList,
		fields.IndicatorIDs:  emptyIndicatorList,
	}

	obj, _ := types.ObjectValue(map[string]attr.Type{
		fields.TracepointIDs: types.ListType{ElemType: types.Int64Type},
		fields.IndicatorIDs:  types.ListType{ElemType: types.Int64Type},
	}, attrs)

	config := &models.InsightSettingsConfig{}

	diags := ExpandInsightSettingsConfig(obj, config)

	testutil.AssertDiagsHasNoErrors(t, diags)
	testutil.AssertEqual(t, "insight setting type name", config.InsightSettingType.Name, cptypes.Inherit)
	testutil.AssertEqual(t, "tracepoint ids length", len(config.TracepointIDs), 0)
	testutil.AssertEqual(t, "indicator ids length", len(config.IndicatorIDs), 0)
}

func TestExpandInsightSettingsConfigLargeValues(t *testing.T) {
	tracepointIDList, _ := types.ListValue(types.Int64Type, []attr.Value{
		types.Int64Value(999999),
		types.Int64Value(888888),
	})

	indicatorIDList, _ := types.ListValue(types.Int64Type, []attr.Value{
		types.Int64Value(777777),
		types.Int64Value(666666),
	})

	attrs := map[string]attr.Value{
		fields.TracepointIDs: tracepointIDList,
		fields.IndicatorIDs:  indicatorIDList,
	}

	obj, _ := types.ObjectValue(map[string]attr.Type{
		fields.TracepointIDs: types.ListType{ElemType: types.Int64Type},
		fields.IndicatorIDs:  types.ListType{ElemType: types.Int64Type},
	}, attrs)

	config := &models.InsightSettingsConfig{}

	diags := ExpandInsightSettingsConfig(obj, config)

	testutil.AssertDiagsHasNoErrors(t, diags)
	testutil.AssertEqual(t, "first large tracepoint id", config.TracepointIDs[0], 999999)
	testutil.AssertEqual(t, "first large indicator id", config.IndicatorIDs[0], 777777)
}

func TestExpandInsightSettingsConfigZeroValues(t *testing.T) {
	tracepointIDList, _ := types.ListValue(types.Int64Type, []attr.Value{
		types.Int64Value(0),
	})

	indicatorIDList, _ := types.ListValue(types.Int64Type, []attr.Value{
		types.Int64Value(0),
	})

	attrs := map[string]attr.Value{
		fields.TracepointIDs: tracepointIDList,
		fields.IndicatorIDs:  indicatorIDList,
	}

	obj, _ := types.ObjectValue(map[string]attr.Type{
		fields.TracepointIDs: types.ListType{ElemType: types.Int64Type},
		fields.IndicatorIDs:  types.ListType{ElemType: types.Int64Type},
	}, attrs)

	config := &models.InsightSettingsConfig{}

	diags := ExpandInsightSettingsConfig(obj, config)

	testutil.AssertDiagsHasNoErrors(t, diags)
	testutil.AssertEqual(t, "zero tracepoint id", config.TracepointIDs[0], 0)
	testutil.AssertEqual(t, "zero indicator id", config.IndicatorIDs[0], 0)
}

func TestExpandInsightSettingsConfigSingleValues(t *testing.T) {
	tracepointIDList, _ := types.ListValue(types.Int64Type, []attr.Value{
		types.Int64Value(42),
	})

	indicatorIDList, _ := types.ListValue(types.Int64Type, []attr.Value{
		types.Int64Value(84),
	})

	attrs := map[string]attr.Value{
		fields.TracepointIDs: tracepointIDList,
		fields.IndicatorIDs:  indicatorIDList,
	}

	obj, _ := types.ObjectValue(map[string]attr.Type{
		fields.TracepointIDs: types.ListType{ElemType: types.Int64Type},
		fields.IndicatorIDs:  types.ListType{ElemType: types.Int64Type},
	}, attrs)

	config := &models.InsightSettingsConfig{}

	diags := ExpandInsightSettingsConfig(obj, config)

	testutil.AssertDiagsHasNoErrors(t, diags)
	testutil.AssertEqual(t, "tracepoint ids length", len(config.TracepointIDs), 1)
	testutil.AssertEqual(t, "indicator ids length", len(config.IndicatorIDs), 1)
	testutil.AssertEqual(t, "single tracepoint id", config.TracepointIDs[0], 42)
	testutil.AssertEqual(t, "single indicator id", config.IndicatorIDs[0], 84)
}

func TestExpandInsightSettingsConfigManyValues(t *testing.T) {
	// Test with a larger number of IDs
	tracepointValues := []attr.Value{}
	indicatorValues := []attr.Value{}

	for i := 1; i <= 10; i++ {
		tracepointValues = append(tracepointValues, types.Int64Value(int64(i)))
		indicatorValues = append(indicatorValues, types.Int64Value(int64(i*10)))
	}

	tracepointIDList, _ := types.ListValue(types.Int64Type, tracepointValues)
	indicatorIDList, _ := types.ListValue(types.Int64Type, indicatorValues)

	attrs := map[string]attr.Value{
		fields.TracepointIDs: tracepointIDList,
		fields.IndicatorIDs:  indicatorIDList,
	}

	obj, _ := types.ObjectValue(map[string]attr.Type{
		fields.TracepointIDs: types.ListType{ElemType: types.Int64Type},
		fields.IndicatorIDs:  types.ListType{ElemType: types.Int64Type},
	}, attrs)

	config := &models.InsightSettingsConfig{}

	diags := ExpandInsightSettingsConfig(obj, config)

	testutil.AssertDiagsHasNoErrors(t, diags)
	testutil.AssertEqual(t, "tracepoint ids length", len(config.TracepointIDs), 10)
	testutil.AssertEqual(t, "indicator ids length", len(config.IndicatorIDs), 10)

	// Verify first and last values
	testutil.AssertEqual(t, "first tracepoint id", config.TracepointIDs[0], 1)
	testutil.AssertEqual(t, "last tracepoint id", config.TracepointIDs[9], 10)
	testutil.AssertEqual(t, "first indicator id", config.IndicatorIDs[0], 10)
	testutil.AssertEqual(t, "last indicator id", config.IndicatorIDs[9], 100)
}

// Integration test with complete configuration
func TestExpandInsightSettingsConfigCompleteConfiguration(t *testing.T) {
	tracepointIDList, _ := types.ListValue(types.Int64Type, []attr.Value{
		types.Int64Value(100),
		types.Int64Value(200),
		types.Int64Value(300),
	})

	indicatorIDList, _ := types.ListValue(types.Int64Type, []attr.Value{
		types.Int64Value(1000),
		types.Int64Value(2000),
	})

	attrs := map[string]attr.Value{
		fields.TracepointIDs: tracepointIDList,
		fields.IndicatorIDs:  indicatorIDList,
	}

	obj, _ := types.ObjectValue(map[string]attr.Type{
		fields.TracepointIDs: types.ListType{ElemType: types.Int64Type},
		fields.IndicatorIDs:  types.ListType{ElemType: types.Int64Type},
	}, attrs)

	config := &models.InsightSettingsConfig{}

	diags := ExpandInsightSettingsConfig(obj, config)

	testutil.AssertDiagsHasNoErrors(t, diags)

	// Verify all fields
	testutil.AssertEqual(t, "insight setting type name", config.InsightSettingType.Name, cptypes.Override)
	testutil.AssertEqual(t, "tracepoint ids length", len(config.TracepointIDs), 3)
	testutil.AssertEqual(t, "indicator ids length", len(config.IndicatorIDs), 2)

	// Verify all tracepoint IDs
	expectedTracepointIDs := []int{100, 200, 300}
	for i, expectedID := range expectedTracepointIDs {
		testutil.AssertEqual(t, "tracepoint id "+string(rune(i)), config.TracepointIDs[i], expectedID)
	}

	// Verify all indicator IDs
	expectedIndicatorIDs := []int{1000, 2000}
	for i, expectedID := range expectedIndicatorIDs {
		testutil.AssertEqual(t, "indicator id "+string(rune(i)), config.IndicatorIDs[i], expectedID)
	}
}

// Test error handling for invalid list data
func TestExpandInsightSettingsConfigTracepointIDsError(t *testing.T) {
	// Create invalid list with actual string data that can't be converted to int
	invalidList, _ := types.ListValue(types.StringType, []attr.Value{
		types.StringValue("not-a-number"),
		types.StringValue("also-invalid"),
	})

	validIndicatorList, _ := types.ListValue(types.Int64Type, []attr.Value{types.Int64Value(100)})

	attrs := map[string]attr.Value{
		fields.TracepointIDs: invalidList,
		fields.IndicatorIDs:  validIndicatorList, // Valid list
	}

	obj, _ := types.ObjectValue(map[string]attr.Type{
		fields.TracepointIDs: types.ListType{ElemType: types.StringType}, // Wrong type
		fields.IndicatorIDs:  types.ListType{ElemType: types.Int64Type},
	}, attrs)

	config := &models.InsightSettingsConfig{}

	diags := ExpandInsightSettingsConfig(obj, config)

	// Should have errors for invalid tracepoint IDs
	testutil.AssertEqual(t, "should have errors for invalid tracepoint data", diags.HasError(), true)

	// Insight setting type should still be set since that happens first
	testutil.AssertEqual(t, "insight setting type should be set", config.InsightSettingType.Name, cptypes.Override)
}

func TestExpandInsightSettingsConfigIndicatorIDsError(t *testing.T) {
	// Create valid tracepoint list and invalid indicator list
	validList, _ := types.ListValue(types.Int64Type, []attr.Value{types.Int64Value(100)})

	// Create invalid list with actual string data that can't be converted to int
	invalidList, _ := types.ListValue(types.StringType, []attr.Value{
		types.StringValue("invalid-indicator"),
		types.StringValue("another-bad-value"),
	})

	attrs := map[string]attr.Value{
		fields.TracepointIDs: validList,
		fields.IndicatorIDs:  invalidList,
	}

	obj, _ := types.ObjectValue(map[string]attr.Type{
		fields.TracepointIDs: types.ListType{ElemType: types.Int64Type},
		fields.IndicatorIDs:  types.ListType{ElemType: types.StringType}, // Wrong type
	}, attrs)

	config := &models.InsightSettingsConfig{}

	diags := ExpandInsightSettingsConfig(obj, config)

	// Should have errors for invalid indicator IDs
	testutil.AssertEqual(t, "should have errors for invalid indicator data", diags.HasError(), true)

	// Insight setting type should still be set since that happens first
	testutil.AssertEqual(t, "insight setting type should be set", config.InsightSettingType.Name, cptypes.Override)

	// Tracepoint IDs should have been processed successfully before the error
	testutil.AssertEqual(t, "tracepoint ids should be set", len(config.TracepointIDs), 1)
	testutil.AssertEqual(t, "tracepoint id value", config.TracepointIDs[0], 100)
}

// Test that config clears existing values when fields are not provided
func TestExpandInsightSettingsConfigClearsFieldsWhenNotProvided(t *testing.T) {
	// Start with a config that has existing values
	config := &models.InsightSettingsConfig{
		TracepointIDs: []int{999, 888},
		IndicatorIDs:  []int{777, 666},
	}

	// Create an object without the ID fields
	attrs := map[string]attr.Value{
		// No TracepointIDs or IndicatorIDs provided
	}

	obj, _ := types.ObjectValue(map[string]attr.Type{}, attrs)

	originalTracepointCount := len(config.TracepointIDs)
	originalIndicatorCount := len(config.IndicatorIDs)

	diags := ExpandInsightSettingsConfig(obj, config)

	testutil.AssertDiagsHasNoErrors(t, diags)

	// Values should be CLEARED, not preserved (following Terraform's declarative model)
	testutil.AssertEqual(t, "tracepoint ids should be cleared", len(config.TracepointIDs), 0)
	testutil.AssertEqual(t, "indicator ids should be cleared", len(config.IndicatorIDs), 0)
	testutil.AssertEqual(t, "insight setting type should be set", config.InsightSettingType.Name, cptypes.Inherit)

	// Verify original values were actually different (test validity)
	testutil.AssertNotEqual(t, "original tracepoint count was not zero", originalTracepointCount, 0)
	testutil.AssertNotEqual(t, "original indicator count was not zero", originalIndicatorCount, 0)
}

// Test early return on first error
func TestExpandInsightSettingsConfigEarlyReturnOnTracepointError(t *testing.T) {
	// Create invalid tracepoint list with actual string data - this should cause early return
	invalidTracepointList, _ := types.ListValue(types.StringType, []attr.Value{
		types.StringValue("this-will-fail"),
	})

	// This valid indicator list should not be processed due to early return
	validIndicatorList, _ := types.ListValue(types.Int64Type, []attr.Value{types.Int64Value(100)})

	attrs := map[string]attr.Value{
		fields.TracepointIDs: invalidTracepointList,
		fields.IndicatorIDs:  validIndicatorList,
	}

	obj, _ := types.ObjectValue(map[string]attr.Type{
		fields.TracepointIDs: types.ListType{ElemType: types.StringType}, // Wrong type
		fields.IndicatorIDs:  types.ListType{ElemType: types.Int64Type},
	}, attrs)

	config := &models.InsightSettingsConfig{}

	diags := ExpandInsightSettingsConfig(obj, config)

	// Should have errors and return early
	testutil.AssertEqual(t, "should have errors", diags.HasError(), true)

	// Due to early return, indicator IDs should not have been processed
	testutil.AssertEqual(t, "indicator ids should not be processed", len(config.IndicatorIDs), 0)

	// Insight setting type should still be set since that happens first
	testutil.AssertEqual(t, "insight setting type should be set", config.InsightSettingType.Name, cptypes.Override)
}
