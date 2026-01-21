package expand

import (
	"catchpoint-provider/internal/fields"
	"catchpoint-provider/internal/models"
	cptypes "catchpoint-provider/internal/types"
	"catchpoint-provider/internal/validation"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// ExpandInsightSettingsConfig expands the plan Object for insightSettings into the Configuration object.
func ExpandInsightSettingsConfig(insightSetting types.Object, config *models.InsightSettingsConfig) (diags diag.Diagnostics) {
	attrs := insightSetting.Attributes()

	if isBlockEmpty(insightSetting) ||
		(attrExistsAndNotNull(attrs, fields.InsightSettingType) && attrs[fields.InsightSettingType].Equal(types.StringValue(cptypes.Inherit))) {
		// If the insight settings are empty, set defaults and return no diagnostics.
		config.InsightSettingType = validation.GetInsightSettingTypeOrDefault(cptypes.Inherit)
		config.IndicatorIDs = []int{}
		config.TracepointIDs = []int{}
		return
	}

	// If the user has specified these settings, we can assume that they're trying to override the defaults.
	config.InsightSettingType = validation.GetInsightSettingTypeOrDefault(cptypes.Override)

	tracepointIDs, getDiags := getIntListFromAttr(fields.TracepointIDs, attrs)
	diags.Append(getDiags...)
	if diags.HasError() {
		return
	}
	config.TracepointIDs = tracepointIDs

	indicatorIDs, getDiags := getIntListFromAttr(fields.IndicatorIDs, attrs)
	diags.Append(getDiags...)
	if diags.HasError() {
		return
	}
	config.IndicatorIDs = indicatorIDs

	return
}
