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
	scheduleSettingsSchemaOnce sync.Once
	scheduleSettingsSchema     schema.SingleNestedBlock
	scheduleSettingsBlockOnce  sync.Once
	scheduleSettingsBlock      map[string]schema.Block
)

// BuildScheduleSettingsBlock builds the block for schedule settings.
func BuildScheduleSettingsBlock(ctx context.Context) map[string]schema.Block {
	scheduleSettingsBlockOnce.Do(func() {
		scheduleSettingsBlock = buildScheduleSettingsBlock(ctx)
	})
	return scheduleSettingsBlock
}

func buildScheduleSettingsBlock(ctx context.Context) map[string]schema.Block {
	return map[string]schema.Block{
		fields.ScheduleSettings: schema.SingleNestedBlock{
			Description: `Used for overriding the schedule settings section. 
			Note: omitting this field entirely means that the schedule settings for this object will not be tracked.
			If you specifically want the object to inherit the parent schedule settings then supply an object with only the schedule_setting_type
			field, e.g. '{ schedule_setting_type = "Inherit" }'.`,
			PlanModifiers: []planmodifier.Object{
				modifier.ScheduleConfigurationModifier(),
				modifier.InheritBlockPlanModifier(fields.ScheduleSettingType),
			},
			Attributes: map[string]schema.Attribute{
				fields.ScheduleSettingType: schema.StringAttribute{
					Optional:    true,
					Computed:    true,
					Description: "Set the schedule setting type, " + genericSettingTypeValidator.Description(ctx),
					Validators:  []validator.String{genericSettingTypeValidator},
				},
				fields.RunScheduleID: schema.Int64Attribute{
					Optional:    true,
					Description: "The Run Schedule ID to use for the Test",
				},
				fields.MaintenanceScheduleID: schema.Int64Attribute{
					Optional:    true,
					Description: "The Maintenance Schedule ID to use for the Test",
				},
				fields.Frequency: schema.StringAttribute{
					Optional:    true,
					Computed:    true,
					Description: "The frequency in seconds for the Test to run, " + frequencyValidator.Description(ctx),
					Validators:  []validator.String{frequencyValidator},
				},
				fields.NodeDistribution: schema.StringAttribute{
					Optional:    true,
					Computed:    true,
					Description: "The node distribution for the Test, " + nodeDistributionValidator.Description(ctx),
					Validators:  []validator.String{nodeDistributionValidator},
				},
				fields.NodeIDs: schema.ListAttribute{
					Optional:    true,
					ElementType: types.Int64Type,
					Description: "The list of Node IDs to use for the Test",
				},
				fields.NodeGroupIDs: schema.ListAttribute{
					Optional:    true,
					ElementType: types.Int64Type,
					Description: "The list of Node Group IDs to use for the Test",
				},
				fields.NoOfSubsetNodes: schema.Int64Attribute{
					Optional:    true,
					Description: "The number of subset nodes to use for the Test",
				},
			},
		},
	}
}

func GetScheduleSettingsAttributeTypes() map[string]attr.Type {
	return extractAttributeTypesFromSchema(getScheduleSettingsSchema().Attributes)
}

func getScheduleSettingsSchema() schema.SingleNestedBlock {
	scheduleSettingsSchemaOnce.Do(func() {
		attrs := BuildScheduleSettingsBlock(context.Background())
		scheduleSettingsSchema = attrs[fields.ScheduleSettings].(schema.SingleNestedBlock)
	})
	return scheduleSettingsSchema
}
