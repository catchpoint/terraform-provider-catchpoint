package schema

import (
	"context"
	"testing"

	"catchpoint-provider/internal/fields"
	"catchpoint-provider/internal/testutil"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

func TestBuildInsightsAttributes(t *testing.T) {
	ctx := context.Background()

	attrs := BuildInsightsBlock(ctx)

	// Should have exactly one attribute
	testutil.AssertEqual(t, "should have one attribute", len(attrs), 1)

	// Should contain Insights field
	insights, exists := attrs[fields.Insights]
	testutil.AssertEqual(t, "should contain Insights field", exists, true)

	// Should be SingleNestedBlock
	singleNested, ok := insights.(schema.SingleNestedBlock)
	testutil.AssertEqual(t, "should be SingleNestedBlock", ok, true)

	// Verify nested attributes
	nestedAttrs := singleNested.Attributes
	testutil.AssertEqual(t, "should have 3 nested attributes", len(nestedAttrs), 3)

	// Verify InsightSettingType field
	insightSettingType, exists := nestedAttrs[fields.InsightSettingType]
	testutil.AssertEqual(t, "should contain InsightSettingType", exists, true)

	settingTypeAttr := insightSettingType.(schema.StringAttribute)
	testutil.AssertEqual(t, "insight setting type should be optional", settingTypeAttr.Optional, true)
	testutil.AssertEqual(t, "insight setting type should not be required", settingTypeAttr.Required, false)
	testutil.AssertEqual(t, "insight setting type should be computed", settingTypeAttr.Computed, true)
	testutil.AssertEqual(t, "insight setting type should have one validator", len(settingTypeAttr.Validators), 1)
	testutil.AssertStringContains(t, "insight setting type description should mention setting type", settingTypeAttr.Description, "insight setting type")

	// Verify TracepointIDs field
	tracepointIDs, exists := nestedAttrs[fields.TracepointIDs]
	testutil.AssertEqual(t, "should contain TracepointIDs", exists, true)

	tracepointIDsAttr := tracepointIDs.(schema.ListAttribute)
	testutil.AssertEqual(t, "tracepoint ids should be optional", tracepointIDsAttr.Optional, true)
	testutil.AssertEqual(t, "tracepoint ids description should mention Tracepoint IDs", tracepointIDsAttr.Description, "The list of Tracepoint IDs to use for the Test")

	// Verify IndicatorIDs field
	indicatorIDs, exists := nestedAttrs[fields.IndicatorIDs]
	testutil.AssertEqual(t, "should contain IndicatorIDs", exists, true)

	indicatorIDsAttr := indicatorIDs.(schema.ListAttribute)
	testutil.AssertEqual(t, "indicator ids should be optional", indicatorIDsAttr.Optional, true)
	testutil.AssertEqual(t, "indicator ids description should mention Indicator IDs", indicatorIDsAttr.Description, "The list of Indicator IDs to use for the Test")
}
