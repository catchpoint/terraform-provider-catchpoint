package expand

import (
	"catchpoint-provider/internal/fields"
	"catchpoint-provider/internal/labels"
	"catchpoint-provider/internal/models"
	cptypes "catchpoint-provider/internal/types"
	"catchpoint-provider/internal/validation"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// ExpandScheduleSettingsConfig expands the plan Object for scheduleSettings into the Configuration object.
func ExpandScheduleSettingsConfig(ScheduleSettingSection types.Object, config *models.ScheduleSettingsConfig) (diags diag.Diagnostics) {
	attrs := ScheduleSettingSection.Attributes()

	if isBlockEmpty(ScheduleSettingSection) ||
		(attrExistsAndNotNull(attrs, fields.ScheduleSettingType) && attrs[fields.ScheduleSettingType].Equal(types.StringValue(cptypes.Inherit))) {
		// If the schedule settings are empty, set defaults and return no diagnostics.
		config.ScheduleSettingType = validation.GetGenericSettingTypeOrDefault(cptypes.Inherit)
		config.NodeIDs = []int{}
		config.NodeGroupIDs = []models.IDNameConfig{}
		return
	}

	// This function is only called if the ScheduleSetting section is not null or unknown.
	// Thus, if the user has specified these settings, we can assume that they're trying to override the defaults.
	scheduleSetting := validation.GetGenericSettingTypeOrDefault(cptypes.Override)
	config.ScheduleSettingType = scheduleSetting

	expandIntSettingFromAttrs(attrs, fields.RunScheduleID, func(val int) {
		config.RunScheduleID = val
	})

	expandIntSettingFromAttrs(attrs, fields.MaintenanceScheduleID, func(val int) {
		config.MaintenanceScheduleID = val
	})

	if attrExistsAndNotNull(attrs, fields.Frequency) {
		frequencyStr := attrs[fields.Frequency].(types.String)
		frequency := validation.GetFrequencyOrDefault(frequencyStr.ValueString())
		config.Frequency = frequency
	} else {
		config.Frequency = validation.GetFrequencyOrDefault(string(cptypes.FiveMinutes)) // Default
	}

	if attrExistsAndNotNull(attrs, fields.NodeDistribution) {
		nodeDistributionStr := attrs[fields.NodeDistribution].(types.String)
		nodeDistribution := validation.GetNodeDistributionOrDefault(nodeDistributionStr.ValueString())
		config.NodeDistribution = nodeDistribution
	}

	// Note: advanced validation is done in the schema configuration to make sure one or more of these
	// (node_ids, node_group_ids) is set.
	nodeIDs, getDiags := getIntListFromAttr(fields.NodeIDs, attrs)
	diags.Append(getDiags...)
	if diags.HasError() {
		return
	}
	config.NodeIDs = nodeIDs

	nodeGroupIDs, getDiags := getIntListFromAttr(fields.NodeGroupIDs, attrs)
	diags.Append(getDiags...)
	if diags.HasError() {
		return
	}

	nodeGroups := []models.IDNameConfig{}
	for _, id := range nodeGroupIDs {
		nodeGroup := models.IDNameConfig{
			ID:   id,
			Name: labels.DefaultNodeGroupName,
		}
		nodeGroups = append(nodeGroups, nodeGroup)
	}
	config.NodeGroupIDs = nodeGroups

	if attrExistsAndNotNull(attrs, fields.NoOfSubsetNodes) {
		noOfSubsetNodes := attrs[fields.NoOfSubsetNodes].(types.Int64)
		config.NoOfSubsetNodes = int(noOfSubsetNodes.ValueInt64())
	}

	return
}
