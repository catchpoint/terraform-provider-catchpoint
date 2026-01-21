package transform

import (
	"context"

	"catchpoint-provider/internal/fields"
	"catchpoint-provider/internal/helpers"
	"catchpoint-provider/internal/models"
	"catchpoint-provider/internal/models/resource"
	cpschema "catchpoint-provider/internal/schema"
	cptypes "catchpoint-provider/internal/types"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func JSONToTerraformInsights(ctx context.Context, insightData *models.InsightDataJSON) (resource resource.InsightSettingsModel, diags diag.Diagnostics) {
	// Get the dynamic attribute types
	attrTypes := cpschema.GetInsightsAttributeTypes()
	obj := types.ObjectNull(attrTypes)
	resource.Insights = obj

	if insightData == nil {
		return
	}

	_, ok := cptypes.GetInsightSettingTypeName(insightData.InsightSettingType.ID)
	if !ok {
		diags.Append(diag.NewWarningDiagnostic("Invalid InsightSettingType", "The provided InsightSettingType is not recognized."))
	}

	// Build the nested attributes map
	attrs, buildDiags := buildInsightDataAttrs(insightData, attrTypes)
	diags.Append(buildDiags...)
	if diags.HasError() {
		return
	}

	// Create the object directly
	obj, objDiags := types.ObjectValue(attrTypes, attrs)
	diags.Append(objDiags...)
	if diags.HasError() {
		return
	}

	resource.Insights = obj
	return
}

func buildInsightDataAttrs(insightData *models.InsightDataJSON, attrTypes map[string]attr.Type) (attrs map[string]attr.Value, diags diag.Diagnostics) {
	attrs = make(map[string]attr.Value)

	if insightData.Indicators != nil || insightData.Tracepoints != nil {
		// If either indicators or tracepoints are set, we can assume the user is overriding the defaults.
		attrs[fields.InsightSettingType] = types.StringValue(cptypes.Override)
	} else {
		attrs[fields.InsightSettingType] = types.StringValue(cptypes.Inherit)
	}

	if insightData.Indicators != nil {
		indicatorIDs, indicatorsDiags := IntSliceToTerraformList(helpers.FlattenToIDs(insightData.Indicators, func(x models.GenericIDNameJSON) int { return x.ID }))
		diags.Append(indicatorsDiags...)
		if diags.HasError() {
			return nil, diags
		}
		attrs[fields.IndicatorIDs] = indicatorIDs
	}

	if insightData.Tracepoints != nil {
		tracepointIDs, tracepointDiags := IntSliceToTerraformList(helpers.FlattenToIDs(insightData.Tracepoints, func(x models.GenericIDNameJSON) int { return x.ID }))
		diags.Append(tracepointDiags...)
		if diags.HasError() {
			return nil, diags
		}
		attrs[fields.TracepointIDs] = tracepointIDs
	}

	SetNullValuesForMissingFields(attrs, attrTypes)

	return
}
