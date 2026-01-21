package transform

import (
	"catchpoint-provider/internal/fields"
	"catchpoint-provider/internal/helpers"
	"catchpoint-provider/internal/logger"
	"catchpoint-provider/internal/models"
	"catchpoint-provider/internal/models/resource"
	cpschema "catchpoint-provider/internal/schema"
	cptypes "catchpoint-provider/internal/types"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func JSONToTerraformScheduleSettings(requestSetting *models.ScheduleSettingsJSON) (resource resource.ScheduleSettingsModel, diags diag.Diagnostics) {
	var attrsType = cpschema.GetScheduleSettingsAttributeTypes()

	obj := types.ObjectNull(attrsType)
	resource.ScheduleSettings = obj

	if requestSetting == nil {
		return
	}

	scheduleSettingTypeName, ok := cptypes.GetGenericSettingTypeName(requestSetting.ScheduleSettingType.ID)
	if !ok {
		logger.WarnBG("Invalid ScheduleSettingType ID: %d", requestSetting.ScheduleSettingType.ID)
		diags.Append(diag.NewWarningDiagnostic("Invalid ScheduleSettingType ID", "The provided ScheduleSettingType ID is not recognized."))
	}

	frequencyName, ok := cptypes.GetFrequencyName(requestSetting.Frequency.ID)
	if !ok {
		logger.WarnBG("Invalid Frequency ID: %d", requestSetting.Frequency.ID)
		diags.Append(diag.NewWarningDiagnostic("Invalid Frequency ID", "The provided Frequency ID is not recognized."))
	}

	nodeDistributionName, ok := cptypes.GetNodeDistributionName(requestSetting.TestNodeDistribution.ID)
	if !ok {
		logger.WarnBG("Invalid NodeDistribution ID: %d", requestSetting.TestNodeDistribution.ID)
		diags.Append(diag.NewWarningDiagnostic("Invalid NodeDistribution ID", "The provided NodeDistribution ID is not recognized."))
	}

	nodeIDList, nodeDiags := IntSliceToTerraformList(
		helpers.FlattenToIDs(requestSetting.Nodes, func(x models.NodeJSON) int { return *x.ID }))
	diags.Append(nodeDiags...)
	if diags.HasError() {
		return
	}

	nodeGroupIDList, nodeGroupDiags := IntSliceToTerraformList(
		helpers.FlattenToIDs(requestSetting.NodeGroups, func(x models.NodeGroupJSON) int {
			// For whatever reason, the backend sometimes returns NodeGroupID and sometimes ID
			// when dealing with NodeGroups. So we check both here. They should never both be nil though.
			if x.NodeGroupID != nil {
				return *x.NodeGroupID
			}
			if x.ID != nil {
				return *x.ID
			}
			return 0
		}))
	diags.Append(nodeGroupDiags...)
	if diags.HasError() {
		return
	}

	attrs := map[string]attr.Value{
		fields.ScheduleSettingType: types.StringValue(scheduleSettingTypeName),
		fields.Frequency:           types.StringValue(frequencyName),
		fields.NodeDistribution:    types.StringValue(nodeDistributionName),
		fields.NodeIDs:             nodeIDList,
		fields.NodeGroupIDs:        nodeGroupIDList,
	}

	SetIntOrNull(attrs, attrsType, fields.RunScheduleID, requestSetting.RunScheduleID)
	SetIntOrNull(attrs, attrsType, fields.MaintenanceScheduleID, requestSetting.MaintenanceScheduleID)
	SetIntOrNull(attrs, attrsType, fields.NoOfSubsetNodes, requestSetting.NoOfSubsetNodes)

	obj, objDiags := types.ObjectValue(attrsType, attrs)
	diags.Append(objDiags...)
	if diags.HasError() {
		return
	}
	resource.ScheduleSettings = obj
	return
}
