package service

import (
	"catchpoint-provider/internal/models"
	"catchpoint-provider/internal/testutil"
	"testing"
)

func newAlertRuleConfig() models.AlertRuleConfig {
	return models.AlertRuleConfig{
		Trigger: models.TriggerConfig{
			CriticalReminderFrequency: models.IDNameConfig{ID: 1, Name: "CriticalReminder"},
			CriticalTrigger:           2.0,
			WarningReminderFrequency:  models.IDNameConfig{ID: 1, Name: "WarningReminder"},
			WarningTrigger:            1.0,
			TriggerType:               models.IDNameConfig{ID: 1, Name: "specific value"},
			OperationType:             models.IDNameConfig{ID: 1, Name: "OperationType"},
			StatisticalType:           models.IDNameConfig{ID: 1, Name: "StatisticalType"},
			HistoricalInterval:        models.IDNameConfig{ID: 1, Name: "TrailingHistoricalInterval"},
			ThresholdInterval:         models.IDNameConfig{ID: 1, Name: "AlertThresholdInterval"},
			Expression:                "expr",
			DNSTTL:                    500,
			DNSResolvedName:           "example.com",
			DNSRecordType:             models.IDNameConfig{ID: 1, Name: "A"},
			FilterType:                models.IDNameConfig{ID: 1, Name: "FilterType"},
			FilterValue:               "filterValue",
			UseIntervalRollingWindow:  true,
			Monitor:                   "Synthetic",
		},
		NodeThresholdConfig: models.NodeThresholdConfig{
			NumberOfUnits:          5,
			NodeThresholdType:      models.IDNameConfig{ID: 0, Name: "NodeType"},
			PercentageOfUnits:      50.0,
			NumberOfFailingUnits:   2,
			ConsecutiveRunsEnabled: true,
			ConsecutiveRuns:        3,
		},
		AlertType:    models.IDNameConfig{ID: 1, Name: "AlertType"},
		AlertSubType: models.IDNameConfig{ID: 1, Name: "SubType"},
		NotificationGroups: []models.NotificationGroupConfig{
			{
				Subject:         "GroupSubject",
				Emails:          []string{},
				ContactGroupIDs: []int{},
			},
		},
	}
}

func checkAlertGroupFields(t *testing.T, testObj *models.AlertGroupJSON, config *models.CommonConfig) {
	testutil.AssertEqual(t, "AlertGroupType.ID", testObj.AlertSettingType.ID, config.AlertSettingsConfig.AlertSettingType.ID)
	testutil.AssertEqual(t, "AlertGroupType.Name", testObj.AlertSettingType.Name, config.AlertSettingsConfig.AlertSettingType.Name)
	testutil.AssertEqual(t, "NotificationGroup.Subject", testObj.NotificationGroup.Subject, alertSubject)

	// At the top-level, an AlertGroupStruct contains a singular NotificationGroup, which in turn contains Recipients and Webhooks.
	compareRecipients(t, testObj.NotificationGroup.Recipients, config.AlertSettingsConfig.NotificationGroup.Emails, config.AlertSettingsConfig.NotificationGroup.ContactGroupIDs)
	// However, an AlertGroupStruct also contains AlertGroupItems and each AlertGroupItem contains its own NotificationGroup*s*.
	checkAlertGroupItems(t, testObj.AlertGroupItems, config.AlertSettingsConfig.AlertRules)
	compareWebhooks(t, testObj.NotificationGroup.AlertWebhooks, config.AlertSettingsConfig.NotificationGroup.WebhookIDs)
}

func checkAlertGroupItems(t *testing.T, alertGroupItems []models.AlertGroupItemJSON, configAlertRuleConfigs []models.AlertRuleConfig) {
	if len(alertGroupItems) != len(configAlertRuleConfigs) {
		t.Errorf("AlertGroupItems count mismatch: got %d, want %d", len(alertGroupItems), len(configAlertRuleConfigs))
	}

	for i, item := range alertGroupItems {
		configItem := configAlertRuleConfigs[i]

		testutil.AssertEqual(t, "EnforceTestFailure", item.EnforceTestFailure, configItem.EnforceTestFailure)
		testutil.AssertEqual(t, "MatchAllRecords", item.MatchAllRecords, configItem.MatchAllRecords)
		testutil.AssertEqual(t, "OmitScatterplot", item.OmitScatterplot, configItem.OmitScatterplot)
		testutil.AssertEqual(t, "AlertType.ID", item.AlertType.ID, configItem.AlertType.ID)
		testutil.AssertEqual(t, "AlertType.Name", item.AlertType.Name, configItem.AlertType.Name)

		testutil.AssertEqual(t, "NotificationType.ID", item.NotificationType.ID, configItem.NotificationType.ID)
		testutil.AssertEqual(t, "NotificationType.Name", item.NotificationType.Name, configItem.NotificationType.Name)

		if item.AlertSubType != nil {
			testutil.AssertEqual(t, "AlertSubType.ID", *item.AlertSubType.ID, configItem.AlertSubType.ID)
			testutil.AssertEqual(t, "AlertSubType.Name", *item.AlertSubType.Name, configItem.AlertSubType.Name)
		}

		checkAlertGroupItemFields(t, &item, configItem)
		checkNotificationGroups(t, &item, configItem.NotificationGroups[0])
	}
}

func checkAlertGroupItemFields(t *testing.T, item *models.AlertGroupItemJSON, config models.AlertRuleConfig) {
	checkNodeThreshold(t, item, config.NodeThresholdConfig)
	checkTriggerFields(t, item, config.Trigger)
	checkNotificationGroups(t, item, config.NotificationGroups[0])
}

func checkNodeThreshold(t *testing.T, item *models.AlertGroupItemJSON, config models.NodeThresholdConfig) {
	testutil.AssertEqual(t, "NodeThreshold.Name", item.NodeThreshold.Name, config.Name)
	testutil.AssertEqual(t, "NodeThreshold.ConsecutiveRunsEnabled", item.NodeThreshold.ConsecutiveRunsEnabled, config.ConsecutiveRunsEnabled)

	if item.NodeThreshold.NumberOfConsecutiveRuns != nil {
		testutil.AssertEqual(t, "NodeThreshold.ConsecutiveRuns", *item.NodeThreshold.NumberOfConsecutiveRuns, config.ConsecutiveRuns)
	} else {
		testutil.AssertNil(t, "NodeThreshold.NumberOfConsecutiveRuns", item.NodeThreshold.NumberOfConsecutiveRuns)
	}

	if item.NodeThreshold.NumberOfUnits != nil {
		testutil.AssertEqual(t, "NodeThreshold.NumberOfUnits", *item.NodeThreshold.NumberOfUnits, config.NumberOfUnits)
	} else {
		testutil.AssertNil(t, "NodeThreshold.NumberOfUnits", item.NodeThreshold.NumberOfUnits)
	}

	if item.NodeThreshold.NumberOfFailingUnits != nil {
		testutil.AssertEqual(t, "NodeThreshold.NumberOfFailingUnits", *item.NodeThreshold.NumberOfFailingUnits, config.NumberOfFailingUnits)
	} else {
		testutil.AssertNil(t, "NodeThreshold.NumberOfFailingUnits", item.NodeThreshold.NumberOfFailingUnits)
	}

	if item.NodeThreshold.PercentageOfUnits != nil {
		testutil.AssertEqual(t, "NodeThreshold.PercentageOfUnits", *item.NodeThreshold.PercentageOfUnits, config.PercentageOfUnits)
	} else {
		testutil.AssertNil(t, "NodeThreshold.PercentageOfUnits", item.NodeThreshold.PercentageOfUnits)
	}

	testutil.AssertEqual(t, "NodeThreshold.UtilizePerNodeHistoricalAverage", item.NodeThreshold.UtilizePerNodeHistoricalAverage, config.UtilizePerNodeHistoricalAverage)
}

func checkTriggerFields(t *testing.T, item *models.AlertGroupItemJSON, config models.TriggerConfig) {
	testutil.AssertEqual(t, "Trigger.CriticalReminderFrequency.ID", item.Trigger.CriticalReminderFrequency.ID, config.CriticalReminderFrequency.ID)
	testutil.AssertEqual(t, "Trigger.CriticalReminderFrequency.Name", item.Trigger.CriticalReminderFrequency.Name, config.CriticalReminderFrequency.Name)
	testutil.AssertEqual(t, "Trigger.CriticalTrigger", *item.Trigger.CriticalTrigger, config.CriticalTrigger)
	testutil.AssertEqual(t, "Trigger.WarningReminderFrequency.ID", item.Trigger.WarningReminderFrequency.ID, config.WarningReminderFrequency.ID)
	testutil.AssertEqual(t, "Trigger.WarningReminderFrequency.Name", item.Trigger.WarningReminderFrequency.Name, config.WarningReminderFrequency.Name)
	testutil.AssertEqual(t, "Trigger.WarningTrigger", *item.Trigger.WarningTrigger, config.WarningTrigger)
	testutil.AssertEqual(t, "Trigger.DNSRecordType.ID", *item.Trigger.DNSRecordType.ID, config.DNSRecordType.ID)
	testutil.AssertEqual(t, "Trigger.DNSRecordType.Name", *item.Trigger.DNSRecordType.Name, config.DNSRecordType.Name)
	testutil.AssertEqual(t, "Trigger.DNSResolvedName", *item.Trigger.DNSResolvedName, config.DNSResolvedName)
	testutil.AssertEqual(t, "Trigger.DNSttl", *item.Trigger.DNSTTL, config.DNSTTL)
	testutil.AssertEqual(t, "Trigger.Expression", *item.Trigger.Expression, config.Expression)
	testutil.AssertEqual(t, "Trigger.FilterType.ID", *item.Trigger.FilterType.ID, config.FilterType.ID)
	testutil.AssertEqual(t, "Trigger.FilterType.Name", *item.Trigger.FilterType.Name, config.FilterType.Name)
	testutil.AssertEqual(t, "Trigger.FilterValue", *item.Trigger.FilterValue, config.FilterValue)
	testutil.AssertEqual(t, "Trigger.HistoricalInterval.ID", *item.Trigger.HistoricalInterval.ID, config.HistoricalInterval.ID)
	testutil.AssertEqual(t, "Trigger.HistoricalInterval.Name", *item.Trigger.HistoricalInterval.Name, config.HistoricalInterval.Name)
	testutil.AssertEqual(t, "Trigger.Monitor", *item.Trigger.Monitor, config.Monitor)
	testutil.AssertEqual(t, "Trigger.OperationType.ID", item.Trigger.OperationType.ID, config.OperationType.ID)
	testutil.AssertEqual(t, "Trigger.OperationType.Name", item.Trigger.OperationType.Name, config.OperationType.Name)
	testutil.AssertEqual(t, "Trigger.StatisticalType.ID", *item.Trigger.StatisticalType.ID, config.StatisticalType.ID)
	testutil.AssertEqual(t, "Trigger.StatisticalType.Name", *item.Trigger.StatisticalType.Name, config.StatisticalType.Name)
	testutil.AssertEqual(t, "Trigger.ThresholdInterval.ID", item.Trigger.ThresholdInterval.ID, config.ThresholdInterval.ID)
	testutil.AssertEqual(t, "Trigger.ThresholdInterval.Name", item.Trigger.ThresholdInterval.Name, config.ThresholdInterval.Name)
	testutil.AssertEqual(t, "Trigger.TriggerType.ID", item.Trigger.TriggerType.ID, config.TriggerType.ID)
	testutil.AssertEqual(t, "Trigger.TriggerType.Name", item.Trigger.TriggerType.Name, config.TriggerType.Name)
	testutil.AssertEqual(t, "Trigger.UseIntervalRollingWindow", item.Trigger.UseIntervalRollingWindow, config.UseIntervalRollingWindow)
}

func checkNotificationGroups(t *testing.T, item *models.AlertGroupItemJSON, config models.NotificationGroupConfig) {
	testutil.AssertEqual(t, "NotificationGroups[0].Subject", item.NotificationGroups[0].Subject, config.Subject)
	testutil.AssertEqual(t, "NotificationGroups[0].NotifyOnCritical", item.NotificationGroups[0].NotifyOnCritical, config.NotifyOnCritical)
	testutil.AssertEqual(t, "NotificationGroups[0].NotifyOnImproved", item.NotificationGroups[0].NotifyOnImproved, config.NotifyOnImproved)
	testutil.AssertEqual(t, "NotificationGroups[0].NotifyOnWarning", item.NotificationGroups[0].NotifyOnWarning, config.NotifyOnWarning)
	compareWebhooks(t, item.NotificationGroups[0].AlertWebhooks, config.WebhookIDs)
	compareRecipients(t, item.NotificationGroups[0].Recipients, config.Emails, config.ContactGroupIDs)
}
