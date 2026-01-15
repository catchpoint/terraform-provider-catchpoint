package merge

import (
	"maps"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func MergeBlockPreservingComputedFields(planBlock, apiBlock types.Object, schemaAttrs map[string]schema.Attribute, expectedAttrTypes map[string]attr.Type) types.Object {
	if planBlock.IsNull() {
		return apiBlock
	}
	if apiBlock.IsNull() {
		return planBlock
	}

	planAttrs := planBlock.Attributes()
	apiAttrs := apiBlock.Attributes()

	mergedAttrs := mergeAPIAttrs(apiAttrs)
	mergePlanAttrs(planAttrs, schemaAttrs, mergedAttrs)
	handleComputedUnknowns(planAttrs, apiAttrs, expectedAttrTypes, mergedAttrs)
	ensureAllAttrsExist(expectedAttrTypes, mergedAttrs)

	mergedObj, diags := types.ObjectValue(expectedAttrTypes, mergedAttrs)
	if diags.HasError() {
		return apiBlock
	}

	return mergedObj
}

func mergeAPIAttrs(apiAttrs map[string]attr.Value) map[string]attr.Value {
	mergedAttrs := make(map[string]attr.Value)
	maps.Copy(mergedAttrs, apiAttrs)
	return mergedAttrs
}

func mergePlanAttrs(planAttrs map[string]attr.Value, schemaAttrs map[string]schema.Attribute, mergedAttrs map[string]attr.Value) {
	for name, value := range planAttrs {
		if value.IsUnknown() {
			continue
		}
		if !value.IsNull() {
			mergedAttrs[name] = value
		}
	}
}

func handleComputedUnknowns(planAttrs, apiAttrs map[string]attr.Value, expectedAttrTypes map[string]attr.Type, mergedAttrs map[string]attr.Value) {
	for attrName, attrType := range expectedAttrTypes {
		if planValue, existsInPlan := planAttrs[attrName]; existsInPlan && planValue.IsUnknown() {
			if apiValue, existsInAPI := apiAttrs[attrName]; existsInAPI && !apiValue.IsUnknown() {
				mergedAttrs[attrName] = apiValue
			} else {
				mergedAttrs[attrName] = createNullValueForType(attrType)
			}
		}
	}
}

func ensureAllAttrsExist(expectedAttrTypes map[string]attr.Type, mergedAttrs map[string]attr.Value) {
	for attrName, attrType := range expectedAttrTypes {
		if _, exists := mergedAttrs[attrName]; !exists {
			mergedAttrs[attrName] = createNullValueForType(attrType)
		}
	}
}

func createNullValueForType(attrType attr.Type) attr.Value {
	switch attrType {
	case types.StringType:
		return types.StringNull()
	case types.BoolType:
		return types.BoolNull()
	case types.Int64Type:
		return types.Int64Null()
	case types.Float64Type:
		return types.Float64Null()
	default:
		// For complex types, need to handle differently
		switch t := attrType.(type) {
		case types.ListType:
			return types.ListNull(t.ElemType)
		case types.SetType:
			return types.SetNull(t.ElemType)
		case types.MapType:
			return types.MapNull(t.ElemType)
		case types.ObjectType:
			return types.ObjectNull(t.AttrTypes)
		default:
			return types.StringNull() // Fallback
		}
	}
}

func isFieldComputed(attr schema.Attribute) bool {
	if attr == nil {
		return false
	}

	// Use reflection or type assertion to check if field is computed-only
	switch a := attr.(type) {
	case schema.StringAttribute:
		return a.Computed
	case schema.BoolAttribute:
		return a.Computed
	case schema.Int64Attribute:
		return a.Computed
	case schema.Float64Attribute:
		return a.Computed
	case schema.ListAttribute:
		return a.Computed
	case schema.SetAttribute:
		return a.Computed
	}
	return false
}
