package schema

import (
	"context"
	"sync"

	"catchpoint-provider/internal/fields"
	"catchpoint-provider/internal/modifier"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var (
	insightSettingsSchemaOnce sync.Once
	insightSettingsBlocks     map[string]schema.Block
)

func BuildInsightsBlock(ctx context.Context) map[string]schema.Block {
	insightSettingsSchemaOnce.Do(func() {
		insightSettingsBlocks = map[string]schema.Block{
			fields.Insights: schema.SingleNestedBlock{
				Description: `Used for overriding the insight settings section. 
			Note: omitting this field entirely means that the insight settings for this object will not be tracked.
			If you specifically want the object to inherit the parent insight settings then supply an object with only the insight_setting_type
			field, e.g. '{ insight_setting_type = "Inherit" }'.`,
				Attributes: map[string]schema.Attribute{
					fields.InsightSettingType: schema.StringAttribute{
						Optional:    true,
						Computed:    true,
						Description: "Set the insight setting type, " + genericSettingTypeValidator.Description(ctx),
						Validators:  []validator.String{genericSettingTypeValidator},
					},
					fields.TracepointIDs: schema.ListAttribute{
						ElementType: types.Int64Type,
						Optional:    true,
						Computed:    true,
						Description: "The list of Tracepoint IDs to use for the Test",
					},
					fields.IndicatorIDs: schema.ListAttribute{
						ElementType: types.Int64Type,
						Optional:    true,
						Computed:    true,
						Description: "The list of Indicator IDs to use for the Test",
					},
				},
				PlanModifiers: []planmodifier.Object{
					modifier.InheritBlockPlanModifier(fields.InsightSettingType),
				},
			},
		}
	})
	return insightSettingsBlocks
}

func GetInsightsAttributeTypes() map[string]attr.Type {
	ctx := context.Background()
	schemaAttrs := BuildInsightsBlock(ctx)

	// Extract the inner attribute types from the nested schema
	if insightsAttr, exists := schemaAttrs[fields.Insights]; exists {
		if singleNested, ok := insightsAttr.(schema.SingleNestedBlock); ok {
			return extractAttributeTypesFromSchema(singleNested.Attributes)
		}
	}
	return make(map[string]attr.Type)
}
