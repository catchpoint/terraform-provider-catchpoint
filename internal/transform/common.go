package transform

import (
	"catchpoint-provider/internal/logger"
	cptypes "catchpoint-provider/internal/types"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func IntSliceToTerraformList(intSlice []int) (types.List, diag.Diagnostics) {
	var diags diag.Diagnostics

	// Return null list if slice is empty
	if len(intSlice) == 0 {
		return types.ListNull(types.Int64Type), diags
	}

	// Convert []int to []attr.Value
	attrValues := make([]attr.Value, len(intSlice))
	for i, v := range intSlice {
		attrValues[i] = types.Int64Value(int64(v))
	}

	// Create the types.List
	list, conversionDiags := types.ListValue(types.Int64Type, attrValues)
	diags.Append(conversionDiags...)

	return list, diags
}

func SetNullValuesForMissingFields(attrs map[string]attr.Value, attrTypes map[string]attr.Type) {
	for fieldName, attrType := range attrTypes {
		if _, exists := attrs[fieldName]; !exists {
			attrs[fieldName] = GetNullValueForType(attrType)
		}
	}
}

func GetNullValueForType(attrType attr.Type) attr.Value {
	switch attrType {
	case types.StringType:
		return types.StringNull()
	case types.Int64Type:
		return types.Int64Null()
	case types.Float64Type:
		return types.Float64Null()
	case types.BoolType:
		return types.BoolNull()
	default:
		if listType, ok := attrType.(types.ListType); ok {
			return types.ListNull(listType.ElemType)
		}
		if setType, ok := attrType.(types.SetType); ok {
			return types.SetNull(setType.ElemType)
		}
		if objType, ok := attrType.(types.ObjectType); ok {
			return types.ObjectNull(objType.AttrTypes)
		}
		return types.StringNull() // fallback
	}
}

func SetStringOrNull(attrs map[string]attr.Value, attrTypes map[string]attr.Type, key string, value *string) {
	if _, exists := attrTypes[key]; exists {
		if value != nil {
			attrs[key] = types.StringValue(*value)
		} else {
			attrs[key] = types.StringValue(cptypes.EmptyString)
		}
	}
}

func SetFloatOrNull(attrs map[string]attr.Value, attrTypes map[string]attr.Type, key string, value *float64) {
	if _, exists := attrTypes[key]; exists {
		if value != nil {
			attrs[key] = types.Float64Value(*value)
		} else {
			attrs[key] = types.Float64Null()
		}
	}
}

func SetIntOrNull(attrs map[string]attr.Value, attrTypes map[string]attr.Type, key string, value *int) {
	if _, exists := attrTypes[key]; exists {
		if value != nil {
			attrs[key] = types.Int64Value(int64(*value))
		} else {
			attrs[key] = types.Int64Null()
		}
	}
}

// Helper function to extract attribute types from resource schema
func ExtractAttributeTypesFromSchema(schemaAttrs map[string]schema.Attribute) map[string]attr.Type {
	attrTypes := make(map[string]attr.Type)
	for name, attr := range schemaAttrs {
		attrTypes[name] = GetAttributeType(attr)
	}
	return attrTypes
}

// Helper function to get the attr.Type from a resource schema.Attribute
func GetAttributeType(attr schema.Attribute) attr.Type {
	switch a := attr.(type) {
	case schema.StringAttribute:
		return types.StringType
	case schema.Int64Attribute:
		return types.Int64Type
	case schema.Float64Attribute:
		return types.Float64Type
	case schema.BoolAttribute:
		return types.BoolType
	case schema.ListAttribute:
		return types.ListType{ElemType: a.ElementType}
	case schema.SetAttribute:
		return types.SetType{ElemType: a.ElementType}
	case schema.SingleNestedAttribute:
		return types.ObjectType{AttrTypes: ExtractAttributeTypesFromSchema(a.Attributes)}
	case schema.ListNestedAttribute:
		return types.ListType{ElemType: types.ObjectType{AttrTypes: ExtractAttributeTypesFromSchema(a.NestedObject.Attributes)}}
	case schema.SetNestedAttribute:
		return types.SetType{ElemType: types.ObjectType{AttrTypes: ExtractAttributeTypesFromSchema(a.NestedObject.Attributes)}}
	default:
		// Handle other attribute types as needed
		return types.StringType // fallback
	}
}

func SetFieldFromLookup(attrs map[string]attr.Value, fieldName string, id int, lookupFunc func(int) (string, bool), diags *diag.Diagnostics) {
	name, ok := lookupFunc(id)
	if ok {
		attrs[fieldName] = types.StringValue(name)
	} else {
		logger.WarnBG("Unknown field from lookup %s %d", fieldName, id)
		diags.Append(diag.NewWarningDiagnostic(fmt.Sprintf("Invalid %s ID", fieldName), fmt.Sprintf("The provided %s ID %d is not recognized.", fieldName, id)))
		attrs[fieldName] = types.StringNull()
	}
}
