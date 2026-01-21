package transform

import (
	"context"
	"testing"

	"catchpoint-provider/internal/fields"
	"catchpoint-provider/internal/models"
	"catchpoint-provider/internal/testutil"

	"github.com/hashicorp/terraform-plugin-framework/types"
)

func TestJSONToTerraformInsights(t *testing.T) {
	ctx := context.Background()

	// Create test data with all fields populated
	insightDataJSON := &models.InsightDataJSON{
		InsightSettingType: models.GenericIDNameJSON{
			ID:   1,
			Name: "override",
		},
		Indicators: []models.GenericIDNameJSON{
			{ID: 101, Name: "Response Time"},
			{ID: 102, Name: "Availability"},
			{ID: 103, Name: "Throughput"},
		},
		Tracepoints: []models.GenericIDNameJSON{
			{ID: 201, Name: "DNS Lookup"},
			{ID: 202, Name: "TCP Connect"},
			{ID: 203, Name: "SSL Handshake"},
		},
	}

	// Execute the function
	result, diags := JSONToTerraformInsights(ctx, insightDataJSON)

	// Verify no errors occurred
	testutil.AssertDiagsHasNoErrors(t, diags)

	// Verify the result is not null
	testutil.AssertEqual(t, "result should not be null", result.Insights.IsNull(), false)

	// Get the outer attributes
	outerAttrs := result.Insights.Attributes()

	// Verify insight setting type
	if insightSettingType, exists := outerAttrs[fields.InsightSettingType]; exists {
		insightSettingTypeStr := insightSettingType.(types.String)
		testutil.AssertEqual(t, "insight setting type", insightSettingTypeStr.ValueString(), "override")
	}

	// Verify indicator IDs list
	if indicatorIDs, exists := outerAttrs[fields.IndicatorIDs]; exists {
		indicatorIDsList := indicatorIDs.(types.List)
		testutil.AssertEqual(t, "indicator ids should not be null", indicatorIDsList.IsNull(), false)

		indicatorIDsElements := indicatorIDsList.Elements()
		testutil.AssertEqual(t, "should have 3 indicator ids", len(indicatorIDsElements), 3)
		testutil.AssertEqual(t, "first indicator id", indicatorIDsElements[0].(types.Int64).ValueInt64(), int64(101))
		testutil.AssertEqual(t, "second indicator id", indicatorIDsElements[1].(types.Int64).ValueInt64(), int64(102))
		testutil.AssertEqual(t, "third indicator id", indicatorIDsElements[2].(types.Int64).ValueInt64(), int64(103))
	}

	// Verify tracepoint IDs list
	if tracepointIDs, exists := outerAttrs[fields.TracepointIDs]; exists {
		tracepointIDsList := tracepointIDs.(types.List)
		testutil.AssertEqual(t, "tracepoint ids should not be null", tracepointIDsList.IsNull(), false)

		tracepointIDsElements := tracepointIDsList.Elements()
		testutil.AssertEqual(t, "should have 3 tracepoint ids", len(tracepointIDsElements), 3)
		testutil.AssertEqual(t, "first tracepoint id", tracepointIDsElements[0].(types.Int64).ValueInt64(), int64(201))
		testutil.AssertEqual(t, "second tracepoint id", tracepointIDsElements[1].(types.Int64).ValueInt64(), int64(202))
		testutil.AssertEqual(t, "third tracepoint id", tracepointIDsElements[2].(types.Int64).ValueInt64(), int64(203))
	}

	// Verify that we have the expected number of attributes
	expectedFieldCount := 3 // InsightSettingType, IndicatorIDs, TracepointIDs
	testutil.AssertEqual(t, "should have expected number of fields", len(outerAttrs), expectedFieldCount)
}
