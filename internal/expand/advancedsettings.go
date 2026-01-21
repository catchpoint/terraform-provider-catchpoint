package expand

import (
	"catchpoint-provider/internal/fields"
	"catchpoint-provider/internal/models"
	cptypes "catchpoint-provider/internal/types"
	"catchpoint-provider/internal/validation"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// ExpandAdvancedSettingsConfig expands the plan Object for advancedSettings into the Configuration object.
func ExpandAdvancedSettingsConfig(advancedSetting types.Object, config *models.AdvancedSettingsConfig) (diags diag.Diagnostics) {
	attrs := advancedSetting.Attributes()

	if isBlockEmpty(advancedSetting) ||
		(attrExistsAndNotNull(attrs, fields.AdvancedSettingType) && attrs[fields.AdvancedSettingType].Equal(types.StringValue(cptypes.Inherit))) {
		// If the advanced settings are empty, or the block is set to inherit, then set defaults and return no diagnostics.
		config.AppliedTestFlags = []int{}
		config.AdvancedSettingType = validation.GetGenericSettingTypeOrDefault(cptypes.Inherit)
		return
	}

	// If anything has been set, we assume the user wants to override the defaults.
	config.AdvancedSettingType = validation.GetGenericSettingTypeOrDefault(cptypes.Override)

	expandStringSettingFromAttrs(attrs, fields.AdditionalMonitor, func(val string) {
		config.AdditionalMonitorType = validation.GetAdditionalMonitorTypeOrDefault(val)
	})

	expandStringSettingFromAttrs(attrs, fields.BandwidthThrottling, func(val string) {
		config.BandwidthThrottling = validation.GetBandwidthThrottlingTypeOrDefault(val)
	})

	expandStringSettingFromAttrs(attrs, fields.EDNSSubnet, func(val string) {
		config.EDNSSubnet = val
	})

	expandIntSettingFromAttrs(attrs, fields.EnforceTestFailureIfRunsLongerThan, func(val int) {
		config.MaxStepRuntimeSecOverride = val
	})

	expandIntSettingFromAttrs(attrs, fields.PingCount, func(val int) {
		config.TraceroutePingCount = val
	})

	expandIntSettingFromAttrs(attrs, fields.FailureHopCount, func(val int) {
		config.TracerouteFailureHopCount = val
	})

	expandIntSettingFromAttrs(attrs, fields.ViewportHeight, func(val int) {
		config.ViewportHeight = val
	})

	expandIntSettingFromAttrs(attrs, fields.ViewportWidth, func(val int) {
		config.ViewportWidth = val
	})

	if attrExistsAndNotNull(attrs, fields.StopTestOnDocumentComplete) &&
		attrExistsAndNotNull(attrs, fields.WaitForNoActivity) {

		waitForNoActivity := attrs[fields.WaitForNoActivity].(types.Int64)
		val := int(waitForNoActivity.ValueInt64())
		config.WaitForNoActivityOnDocComplete = &val
	}

	expandAppliedTestFlags(attrs, config)

	return diags
}

// Expands the advanced settings into TestFlags for a CommonConfig.
//
// Note: this is specifically for the TestFlags, not any other Advanced Setting.
func expandAppliedTestFlags(attrs map[string]attr.Value, config *models.AdvancedSettingsConfig) {
	testFlags := cptypes.ValidTestFlagNames
	appliedTestFlagIDs := make([]int, 0, len(testFlags))

	for _, testFlag := range testFlags {
		if attrVal, exists := attrs[testFlag]; exists && !attrVal.IsNull() && !attrVal.IsUnknown() {
			// For boolean test flags, only add if explicitly set to true
			if boolVal, ok := attrVal.(types.Bool); ok && boolVal.ValueBool() {
				if id, ok := cptypes.GetTestFlagID(testFlag); ok {
					appliedTestFlagIDs = append(appliedTestFlagIDs, id)
				}
			}
		}
	}

	config.AppliedTestFlags = appliedTestFlagIDs
}
