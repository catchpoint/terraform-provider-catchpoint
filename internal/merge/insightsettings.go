package merge

import (
	"catchpoint-provider/internal/fields"
	cpresource "catchpoint-provider/internal/models/resource"
	cpschema "catchpoint-provider/internal/schema"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func InsightSettings(plan, apiState *cpresource.InsightSettingsModel, objSchema schema.Schema) (result *cpresource.InsightSettingsModel) {
	if !plan.Insights.IsNull() {
		insightsBlock := objSchema.Blocks[fields.Insights].(schema.SingleNestedBlock)
		mergedObj := MergeBlockPreservingComputedFields(
			plan.Insights,
			apiState.Insights,
			insightsBlock.Attributes,
			cpschema.GetInsightsAttributeTypes(),
		)
		result = &cpresource.InsightSettingsModel{Insights: mergedObj}
	} else {
		result = &cpresource.InsightSettingsModel{Insights: types.ObjectNull(cpschema.GetInsightsAttributeTypes())}
	}
	return
}
