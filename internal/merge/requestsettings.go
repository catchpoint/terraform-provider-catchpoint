package merge

import (
	"catchpoint-provider/internal/fields"
	cpresource "catchpoint-provider/internal/models/resource"
	cpschema "catchpoint-provider/internal/schema"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func RequestSettings(plan, apiState *cpresource.RequestSettingsModel, objSchema schema.Schema) (result *cpresource.RequestSettingsModel) {
	if plan.RequestSettings.IsNull() {
		result = &cpresource.RequestSettingsModel{RequestSettings: types.ObjectNull(cpschema.GetRequestSettingsAttributeTypes())}
		return
	}

	requestBlock := objSchema.Blocks[fields.RequestSettings].(schema.SingleNestedBlock)
	mergedObj := MergeBlockPreservingComputedFields(
		plan.RequestSettings,
		apiState.RequestSettings,
		requestBlock.Attributes,
		cpschema.GetRequestSettingsAttributeTypes(),
	)

	mergedObj = mergeRequestSettingsHeaders(plan, apiState, mergedObj)

	result = &cpresource.RequestSettingsModel{RequestSettings: mergedObj}
	return
}

// mergeRequestSettingsHeaders handles the merging of HTTPRequestHeaders for RequestSettings.
func mergeRequestSettingsHeaders(plan, apiState *cpresource.RequestSettingsModel, mergedObj types.Object) types.Object {
	planHeaders, planHas := plan.RequestSettings.Attributes()[fields.HTTPRequestHeaders]
	apiHeaders, apiHas := apiState.RequestSettings.Attributes()[fields.HTTPRequestHeaders]

	attrs := mergedObj.Attributes()

	if planHas && apiHas {
		planHeadersObj, planObjOk := planHeaders.(types.Object)
		apiHeadersObj, apiObjOk := apiHeaders.(types.Object)

		if planObjOk && apiObjOk {
			if planHeadersObj.IsNull() || planHeadersObj.IsUnknown() {
				mergedHeaders := types.ObjectNull(cpschema.GetHTTPRequestHeadersAttributeTypes())
				attrs[fields.HTTPRequestHeaders] = mergedHeaders
			} else {
				mergedHeaders := mergeHTTPRequestHeaders(planHeadersObj, apiHeadersObj)
				attrs[fields.HTTPRequestHeaders] = mergedHeaders
			}
			mergedObj, _ = types.ObjectValue(cpschema.GetRequestSettingsAttributeTypes(), attrs)
			return mergedObj
		}
	}

	attrs[fields.HTTPRequestHeaders] = types.ObjectNull(cpschema.GetHTTPRequestHeadersAttributeTypes())
	mergedObj, _ = types.ObjectValue(cpschema.GetRequestSettingsAttributeTypes(), attrs)
	return mergedObj
}

func mergeHTTPRequestHeaders(planHeadersObj, apiHeadersObj types.Object) types.Object {
	mergedAttrs := make(map[string]attr.Value)

	for headerName, planValRaw := range planHeadersObj.Attributes() {
		mergedAttrs[headerName] = mergeSingleHeader(headerName, planValRaw, apiHeadersObj)
	}

	mergedObj, _ := types.ObjectValue(cpschema.GetHTTPRequestHeadersAttributeTypes(), mergedAttrs)
	return mergedObj
}

// mergeSingleHeader merges a single header's attributes from plan and api objects.
func mergeSingleHeader(headerName string, planValRaw attr.Value, apiHeadersObj types.Object) attr.Value {
	planVal, planObjOk := planValRaw.(types.Object)
	if !planObjOk {
		// Defensive: return as-is if not an object
		return planValRaw
	}
	if planVal.IsNull() || planVal.IsUnknown() {
		// Return null/unknown as-is, do not merge
		return planVal
	}

	// Try to get the corresponding API value
	apiValRaw := apiHeadersObj.Attributes()[headerName]
	apiVal, apiObjOk := apiValRaw.(types.Object)

	attrTypes := cpschema.GetHTTPHeaderAttributeTypes(headerName)
	mergedHeaderAttrs := make(map[string]attr.Value)
	for field := range attrTypes {
		mergedHeaderAttrs[field] = mergeHeaderField(field, attrTypes[field], planVal, apiVal, apiObjOk)
	}

	mergedHeaderObj, _ := types.ObjectValue(attrTypes, mergedHeaderAttrs)
	return mergedHeaderObj
}

// mergeHeaderField merges a single field of a header.
func mergeHeaderField(field string, fieldType attr.Type, planVal types.Object, apiVal types.Object, apiObjOk bool) attr.Value {
	planField := planVal.Attributes()[field]
	var apiField attr.Value
	if apiObjOk {
		apiField = apiVal.Attributes()[field]
	}

	if planField != nil && !planField.IsUnknown() && !planField.IsNull() {
		return planField
	}
	if apiObjOk && apiField != nil && !apiField.IsUnknown() && !apiField.IsNull() {
		return apiField
	}

	switch fieldType {
	case types.StringType:
		return types.StringValue("")
	case types.Int64Type:
		return types.Int64Null()
	case types.BoolType:
		return types.BoolNull()
	default:
		return nil
	}
}
