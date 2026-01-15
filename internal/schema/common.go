package schema

import (
	maps0 "maps"
	"slices"

	cptypes "catchpoint-provider/internal/types"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var (
	StatusValidator             = oneOfStringValidator(cptypes.ValidStatusTypes)
	genericSettingTypeValidator = oneOfStringValidator(cptypes.ValidGenericSettingTypeNames)
	frequencyValidator          = oneOfStringValidator(cptypes.ValidFrequencyNames)
	nodeDistributionValidator   = oneOfStringValidator(cptypes.ValidNodeDistributions)
)

func MergeAttributes(maps ...map[string]schema.Attribute) map[string]schema.Attribute {
	result := map[string]schema.Attribute{}
	for _, m := range maps {
		maps0.Copy(result, m)
	}
	return result
}

func MergeBlocks(maps ...map[string]schema.Block) map[string]schema.Block {
	result := map[string]schema.Block{}
	for _, m := range maps {
		maps0.Copy(result, m)
	}
	return result
}

func GetAttributeKeys(attrs map[string]schema.Attribute) []string {
	keys := make([]string, 0, len(attrs))
	for k := range attrs {
		keys = append(keys, k)
	}
	return keys
}

func GetBlockKeys(blocks map[string]schema.Block) []string {
	keys := make([]string, 0, len(blocks))
	for k := range blocks {
		keys = append(keys, k)
	}
	return keys
}

func GetSingleNestedAttributes(attrs map[string]schema.Attribute, key string) (map[string]schema.Attribute, bool) {
	attr, ok := attrs[key]
	if !ok {
		return nil, false
	}
	single, ok := attr.(schema.SingleNestedAttribute)
	if !ok {
		return nil, false
	}
	return single.Attributes, true
}

func GetSingleNestedBlockAttributes(blocks map[string]schema.Block, blockName string) (map[string]schema.Attribute, bool) {
	if block, exists := blocks[blockName]; exists {
		if singleNestedBlock, ok := block.(schema.SingleNestedBlock); ok {
			return singleNestedBlock.Attributes, true
		}
	}
	return nil, false
}

func GetSingleNestedBlocks(blocks map[string]schema.Block, blockName string) (map[string]schema.Block, bool) {
	if block, exists := blocks[blockName]; exists {
		if singleNestedBlock, ok := block.(schema.SingleNestedBlock); ok {
			return singleNestedBlock.Blocks, true
		}
	}
	return nil, false
}

func ExtractAllAttributeTypes(attrs map[string]schema.Attribute, blocks map[string]schema.Block) map[string]attr.Type {
	attrTypes := make(map[string]attr.Type)

	// Handle attributes
	for name, attr := range attrs {
		attrTypes[name] = getAttributeType(attr)
	}

	// Handle blocks
	for name, block := range blocks {
		attrTypes[name] = getBlockType(block)
	}

	return attrTypes
}

// OneOfStringValidator returns a string validator built from the provided values.
// It copies and sorts the slice so the produced description is deterministic.
func oneOfStringValidator(values []string) validator.String {
	copied := slices.Clone(values)
	slices.Sort(copied)
	return stringvalidator.OneOf(copied...)
}

// Helper function to extract attr.Type from schema.Attribute map
func extractAttributeTypesFromSchema(attrs map[string]schema.Attribute) map[string]attr.Type {
	attrTypes := make(map[string]attr.Type)
	for name, attr := range attrs {
		attrTypes[name] = getAttributeType(attr)
	}
	return attrTypes
}

// Helper function to convert schema.Attribute to attr.Type
func getAttributeType(attr schema.Attribute) attr.Type {
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
		return types.ObjectType{AttrTypes: extractAttributeTypesFromSchema(a.Attributes)}
	case schema.ListNestedAttribute:
		return types.ListType{ElemType: types.ObjectType{AttrTypes: extractAttributeTypesFromSchema(a.NestedObject.Attributes)}}
	case schema.SetNestedAttribute:
		return types.SetType{ElemType: types.ObjectType{AttrTypes: extractAttributeTypesFromSchema(a.NestedObject.Attributes)}}
	default:
		// Fallback for unknown types
		return types.StringType
	}
}

func getBlockType(block schema.Block) attr.Type {
	switch b := block.(type) {
	case schema.SingleNestedBlock:
		innerAttrTypes := make(map[string]attr.Type)
		for innerName, innerAttr := range b.Attributes {
			innerAttrTypes[innerName] = getAttributeType(innerAttr)
		}
		// Also handle nested blocks within this block
		for innerName, innerBlock := range b.Blocks {
			innerAttrTypes[innerName] = getBlockType(innerBlock)
		}
		return types.ObjectType{AttrTypes: innerAttrTypes}
	case schema.SetNestedBlock:
		innerAttrTypes := make(map[string]attr.Type)
		for innerName, innerAttr := range b.NestedObject.Attributes {
			innerAttrTypes[innerName] = getAttributeType(innerAttr)
		}
		// Also handle nested blocks within this block
		for innerName, innerBlock := range b.NestedObject.Blocks {
			innerAttrTypes[innerName] = getBlockType(innerBlock)
		}
		return types.SetType{ElemType: types.ObjectType{AttrTypes: innerAttrTypes}}
	case schema.ListNestedBlock:
		innerAttrTypes := make(map[string]attr.Type)
		for innerName, innerAttr := range b.NestedObject.Attributes {
			innerAttrTypes[innerName] = getAttributeType(innerAttr)
		}
		// Also handle nested blocks within this block
		for innerName, innerBlock := range b.NestedObject.Blocks {
			innerAttrTypes[innerName] = getBlockType(innerBlock)
		}
		return types.ListType{ElemType: types.ObjectType{AttrTypes: innerAttrTypes}}
	default:
		return types.StringType
	}
}
