package catchpoint

import (
	"catchpoint-provider/internal/fields"

	"github.com/hashicorp/terraform-plugin-framework/types"
)

func listToIntSlice(list types.List) []int64 {
	if list.IsNull() || list.IsUnknown() {
		return []int64{}
	}

	elements := list.Elements()
	result := make([]int64, len(elements))

	for i, elem := range elements {
		if int64Val, ok := elem.(types.Int64); ok && !int64Val.IsNull() && !int64Val.IsUnknown() {
			result[i] = int64Val.ValueInt64()
		}
	}

	return result
}

func extractInsightsAttributes(insights types.Object) (types.List, types.List) {
	// Default to null lists
	nullList := types.ListNull(types.Int64Type)

	if insights.IsNull() || insights.IsUnknown() {
		return nullList, nullList
	}

	attrs := insights.Attributes()

	// Extract indicator_ids
	indicators := nullList
	if indicatorAttr, exists := attrs[fields.IndicatorIDs]; exists && !indicatorAttr.IsNull() {
		if indicatorList, ok := indicatorAttr.(types.List); ok {
			indicators = indicatorList
		}
	}

	// Extract tracepoint_ids
	tracepoints := nullList
	if tracepointAttr, exists := attrs[fields.TracepointIDs]; exists && !tracepointAttr.IsNull() {
		if tracepointList, ok := tracepointAttr.(types.List); ok {
			tracepoints = tracepointList
		}
	}

	return indicators, tracepoints
}
