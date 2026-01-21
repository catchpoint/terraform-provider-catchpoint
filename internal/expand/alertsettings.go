package expand

import (
	"catchpoint-provider/internal/models"
	"catchpoint-provider/internal/models/resource"
	"catchpoint-provider/internal/types"
	cptypes "catchpoint-provider/internal/types"
	"catchpoint-provider/internal/validation"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/diag"
)

// ExpandAlertSettingsConfig expands the plan Object for alertSettings into the Configuration object.
func ExpandAlertSettingsConfig(alertSettings *resource.AlertSettingsModel, config *models.AlertSettingsConfig) (diags diag.Diagnostics) {
	if alertSettings.IsNull() || alertSettings.IsUnknown() {
		// If alertSettings is nil, set to default "no settings" type.
		config.AlertSettingType = validation.GetAlertSettingTypeOrDefault(types.Inherit)
		config.AlertRules = []models.AlertRuleConfig{}
		config.NotificationGroup = models.NotificationGroupConfig{}
		return
	}

	config.AlertSettingType = validation.GetAlertSettingTypeOrDefault(alertSettings.AlertSettingType.ValueString())

	// If the plan calls for any AlertRule objects, expand them into the config.
	if len(alertSettings.AlertRule) == 0 {
		config.AlertRules = []models.AlertRuleConfig{}
	} else {
		for _, alertRuleObj := range alertSettings.AlertRule {
			alertRuleConfig := models.AlertRuleConfig{}
			expandAlertRuleConfig(alertRuleObj, &alertRuleConfig)
			config.AlertRules = append(config.AlertRules, alertRuleConfig)
		}
	}

	// The top-level NotificationGroup is a SingledNestedAttribute so there is only one. But only set it if it exists.
	if alertSettings.NotificationGroup == nil {
		config.NotificationGroup = models.NotificationGroupConfig{}
	} else {
		expandNotificationGroup(alertSettings.NotificationGroup, &config.NotificationGroup)
	}

	return
}

func expandAlertRuleConfig(alertSettings resource.AlertRuleModel, config *models.AlertRuleConfig) (diags diag.Diagnostics) {
	diags = append(diags, expandStringSettingWithLookup(alertSettings.AlertType, cptypes.GetAlertTypeID, &config.AlertType)...)
	diags = append(diags, expandStringSettingWithLookup(alertSettings.AlertSubType, cptypes.GetAlertSubTypeID, &config.AlertSubType)...)
	diags = append(diags, expandStringSettingWithLookup(alertSettings.NotificationType, cptypes.GetNotificationTypeID, &config.NotificationType)...)
	if diags.HasError() {
		return
	}

	for _, notifGroup := range alertSettings.NotificationGroup {
		notifGroupConfig := models.NotificationGroupConfig{}
		diags = append(diags, expandNotificationGroup(&notifGroup, &notifGroupConfig)...)
		config.NotificationGroups = append(config.NotificationGroups, notifGroupConfig)
	}

	expandBoolSetting(alertSettings.EnforceTestFailure, &config.EnforceTestFailure)
	expandBoolSetting(alertSettings.OmitScatterplot, &config.OmitScatterplot)
	expandBoolSetting(alertSettings.AllMatchRecords, &config.MatchAllRecords)

	thresholdDiag := expandNodeThresholdConfig(alertSettings.NodeThresholdModel, &config.NodeThresholdConfig)
	diags.Append(thresholdDiag...)
	if diags.HasError() {
		return
	}

	triggerDiag := expandTriggerConfig(alertSettings.TriggerModel, &config.Trigger)
	diags.Append(triggerDiag...)
	if diags.HasError() {
		return
	}

	// If the AlertType is HostFailure or TestFailure, or the TriggerType is TrendShift, the operation type defaults to NotEquals
	// because these types do not support an operation.
	if config.AlertType.ID == 9 || config.Trigger.TriggerType.ID == 3 || config.AlertType.ID == 4 {
		config.Trigger.OperationType = models.IDNameConfig{ID: 0, Name: cptypes.NotEquals}
	}

	return
}

func expandTriggerConfig(triggerModel resource.TriggerModel, config *models.TriggerConfig) (diags diag.Diagnostics) {
	expandStringSetting(triggerModel.Expression, &config.Expression)

	diags = append(diags, expandStringSettingWithLookup(triggerModel.TriggerType, cptypes.GetTriggerTypeID, &config.TriggerType)...)
	diags = append(diags, expandStringSettingWithLookup(triggerModel.OperationType, cptypes.GetOperationTypeID, &config.OperationType)...)
	diags = append(diags, expandStringSettingWithLookup(triggerModel.StatisticalType, cptypes.GetStatisticalTypeID, &config.StatisticalType)...)
	diags = append(diags, expandStringSettingWithLookup(triggerModel.HistoricalInterval, cptypes.GetHistoricalIntervalID, &config.HistoricalInterval)...)
	if diags.HasError() {
		return
	}

	expandFloat64Setting(triggerModel.WarningTrigger, &config.WarningTrigger)
	expandFloat64Setting(triggerModel.CriticalTrigger, &config.CriticalTrigger)

	msg := fmt.Sprintf("Expanding alert rule UseRollingWindow: %v (IsNull: %v, IsUnknown: %v)",
		triggerModel.UseRollingWindow.ValueBool(),
		triggerModel.UseRollingWindow.IsNull(),
		triggerModel.UseRollingWindow.IsUnknown())

	diags.AddWarning("Debug Info", msg)

	config.UseIntervalRollingWindow = triggerModel.UseRollingWindow.ValueBool()

	diags = append(diags, expandStringSettingWithLookup(triggerModel.WarningReminder, cptypes.GetReminderID, &config.WarningReminderFrequency)...)
	diags = append(diags, expandStringSettingWithLookup(triggerModel.CriticalReminder, cptypes.GetReminderID, &config.CriticalReminderFrequency)...)
	diags = append(diags, expandStringSettingWithLookup(triggerModel.ThresholdInterval, cptypes.GetThresholdIntervalID, &config.ThresholdInterval)...)
	if diags.HasError() {
		return
	}

	// DNS settings.
	expandStringSetting(triggerModel.DNSResolvedName, &config.DNSResolvedName)
	expandIntSetting(triggerModel.DNSTTL, &config.DNSTTL)
	diags = append(diags, expandStringSettingWithLookup(triggerModel.DNSRecordType, cptypes.GetDNSRecordTypeID, &config.DNSRecordType)...)

	// "Level" settings.
	if triggerModel.Level != nil {
		expandStringSetting(triggerModel.Level.FilterValue, &config.FilterValue)
		diags = append(diags, expandStringSettingWithLookup(triggerModel.Level.FilterType, cptypes.GetFilterTypeID, &config.FilterType)...)
	}

	return
}

func expandNodeThresholdConfig(alertSettings resource.NodeThresholdModel, config *models.NodeThresholdConfig) (diags diag.Diagnostics) {
	diags = append(diags, expandStringSettingWithLookup(alertSettings.NodeThresholdType, cptypes.GetNodeThresholdTypeID, &config.NodeThresholdType)...)
	if diags.HasError() {
		return
	}

	expandIntSetting(alertSettings.ConsecutiveNumberOfRuns, &config.ConsecutiveRuns)
	expandIntSetting(alertSettings.ThresholdNumberOfRuns, &config.NumberOfUnits)
	expandFloat64Setting(alertSettings.ThresholdPercentageOfRuns, &config.PercentageOfUnits)
	expandIntSetting(alertSettings.NumberOfFailingNodes, &config.NumberOfFailingUnits)
	expandBoolSetting(alertSettings.EnableConsecutive, &config.ConsecutiveRunsEnabled)

	return
}

func expandNotificationGroup(notifGroup *resource.NotificationGroupModel, group *models.NotificationGroupConfig) (diags diag.Diagnostics) {
	if notifGroup == nil {
		return
	}

	expandBoolSetting(notifGroup.NotifyOnWarning, &group.NotifyOnWarning)
	expandBoolSetting(notifGroup.NotifyOnCritical, &group.NotifyOnCritical)
	expandBoolSetting(notifGroup.NotifyOnImproved, &group.NotifyOnImproved)

	diags = append(diags, expandIntListSetting(notifGroup.AlertWebhookIDs, &group.WebhookIDs)...)
	if diags.HasError() {
		return
	}

	diags = append(diags, expandStringListSetting(notifGroup.Emails, &group.Emails)...)
	if diags.HasError() {
		return
	}

	diags = append(diags, expandIntListSetting(notifGroup.ContactGroupIDs, &group.ContactGroupIDs)...)
	if diags.HasError() {
		return
	}

	expandStringSetting(notifGroup.Subject, &group.Subject)

	return
}
