package merge

import (
	"catchpoint-provider/internal/fields"
	cpresource "catchpoint-provider/internal/models/resource"
	cpschema "catchpoint-provider/internal/schema"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func ScheduleSettings(plan, apiState *cpresource.ScheduleSettingsModel, objSchema schema.Schema) (result *cpresource.ScheduleSettingsModel) {
	if !plan.ScheduleSettings.IsNull() {
		scheduleBlock := objSchema.Blocks[fields.ScheduleSettings].(schema.SingleNestedBlock)
		mergedObj := MergeBlockPreservingComputedFields(
			plan.ScheduleSettings,
			apiState.ScheduleSettings,
			scheduleBlock.Attributes,
			cpschema.GetScheduleSettingsAttributeTypes(),
		)

		result = &cpresource.ScheduleSettingsModel{ScheduleSettings: mergedObj}
	} else {
		result = &cpresource.ScheduleSettingsModel{ScheduleSettings: types.ObjectNull(cpschema.GetScheduleSettingsAttributeTypes())}
	}
	return
}
