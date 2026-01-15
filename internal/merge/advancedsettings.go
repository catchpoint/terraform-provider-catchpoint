package merge

import (
	"catchpoint-provider/internal/fields"
	cpresource "catchpoint-provider/internal/models/resource"
	cpschema "catchpoint-provider/internal/schema"
	cptypes "catchpoint-provider/internal/types"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

func AdvancedSettings(plan, apiState *cpresource.AdvancedSettingsModel, objSchema schema.Schema) (result *cpresource.AdvancedSettingsModel) {
	if !plan.AdvancedSettings.IsNull() {
		advancedBlock := objSchema.Blocks[fields.AdvancedSettings].(schema.SingleNestedBlock)
		mergedObj := MergeBlockPreservingComputedFields(
			plan.AdvancedSettings,
			apiState.AdvancedSettings,
			advancedBlock.Attributes,
			cpschema.GetAdvancedSettingsAttributeTypesForProductAndFolder(),
		)
		result = &cpresource.AdvancedSettingsModel{AdvancedSettings: mergedObj}
	} else {
		result = &cpresource.AdvancedSettingsModel{AdvancedSettings: plan.AdvancedSettings}
	}
	return
}

func AdvancedSettingsForTest(plan, apiState *cpresource.AdvancedSettingsModel, objSchema schema.Schema, testType cptypes.TestType) (result *cpresource.AdvancedSettingsModel) {
	if !plan.AdvancedSettings.IsNull() {
		advancedBlock := objSchema.Blocks[fields.AdvancedSettings].(schema.SingleNestedBlock)
		mergedObj := MergeBlockPreservingComputedFields(
			plan.AdvancedSettings,
			apiState.AdvancedSettings,
			advancedBlock.Attributes,
			cpschema.GetAdvancedSettingsAttributeTypesForTest(testType),
		)
		result = &cpresource.AdvancedSettingsModel{AdvancedSettings: mergedObj}
	} else {
		result = &cpresource.AdvancedSettingsModel{AdvancedSettings: plan.AdvancedSettings}
	}
	return
}
