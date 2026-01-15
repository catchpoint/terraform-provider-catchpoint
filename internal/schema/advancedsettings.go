package schema

import (
	"context"

	"catchpoint-provider/internal/fields"
	"catchpoint-provider/internal/logger"
	"catchpoint-provider/internal/modifier"
	"catchpoint-provider/internal/types"
	"catchpoint-provider/internal/validation/advancedsettings"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

func GetAdvancedSettingsAttributeTypesForTest(testType types.TestType) map[string]attr.Type {
	// Get the dynamic schema attributes for the specific test type
	ctx := context.Background()
	schemaAttrs := BuildAdvancedSettingsBlockForTest(ctx, testType)

	// Extract the inner attribute types from the nested schema
	if advancedSettingsAttr, exists := schemaAttrs[fields.AdvancedSettings]; exists {
		if singleNested, ok := advancedSettingsAttr.(schema.SingleNestedBlock); ok {
			return extractAttributeTypesFromSchema(singleNested.Attributes)
		}
	}
	return make(map[string]attr.Type)
}

func GetAdvancedSettingsAttributeTypesForProductAndFolder() map[string]attr.Type {
	// Get the dynamic schema attributes
	ctx := context.Background()
	schemaAttrs := BuildAdvancedSettingsAttributesForProductAndFolder(ctx)

	// Extract the inner attribute types from the nested schema
	if advancedSettingsAttr, exists := schemaAttrs[fields.AdvancedSettings]; exists {
		if singleNested, ok := advancedSettingsAttr.(schema.SingleNestedBlock); ok {
			return extractAttributeTypesFromSchema(singleNested.Attributes)
		} else {
			logger.WarnBG("AdvancedSettings is not a SingleNestedBlock")
		}
	} else {
		logger.WarnBG("AdvancedSettings attribute not found in schema")
		logger.WarnBG("Schema attributes: %+v", schemaAttrs)
	}

	return make(map[string]attr.Type)
}

func BuildAdvancedSettingsBlockForTest(ctx context.Context, testType types.TestType) map[string]schema.Block {
	return buildAdvancedSettingsBlock(ctx, advancedsettings.GetAdvancedSettingsForTestType(testType))
}

func BuildAdvancedSettingsAttributesForProductAndFolder(ctx context.Context) map[string]schema.Block {
	return buildAdvancedSettingsBlock(ctx, advancedsettings.GetAllAdvancedSettings())
}

func buildAdvancedSettingsBlock(ctx context.Context, keys []string) map[string]schema.Block {
	attrs := make(map[string]schema.Attribute)
	for _, key := range keys {
		if builder, ok := advancedSettingAttributes[key]; ok {
			attrs[key] = builder(ctx)
		}
	}

	var typeAttribute map[string]schema.Attribute

	// Only add the alert setting type if there are any attributes to merge.
	// (Since BGP has no advanced settings, it also doesn't need an advancedSettingsType.)
	if len(attrs) > 0 {
		typeAttribute = map[string]schema.Attribute{
			fields.AdvancedSettingType: schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Set the advanced settings type, " + genericSettingTypeValidator.Description(ctx),
				Validators:  []validator.String{genericSettingTypeValidator},
			},
		}
	}

	return map[string]schema.Block{
		fields.AdvancedSettings: schema.SingleNestedBlock{
			Description: `Used for overriding the advanced settings section. 
			Note: omitting this field entirely means that the advanced settings for this object will not be tracked.
			If you specifically want the object to inherit the parent advanced settings then supply an object with only the advanced_setting_type
			field, e.g. '{ advanced_setting_type = "Inherit" }'.`,
			Attributes: MergeAttributes(
				typeAttribute,
				attrs,
			),
			PlanModifiers: []planmodifier.Object{
				modifier.InheritBlockPlanModifier(fields.AdvancedSettingType),
			},
		},
	}
}
