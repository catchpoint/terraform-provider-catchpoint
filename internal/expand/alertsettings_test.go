package expand

import (
	"testing"

	"catchpoint-provider/internal/fields"
	"catchpoint-provider/internal/models"
	"catchpoint-provider/internal/models/resource"
	"catchpoint-provider/internal/testutil"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func TestExpandConfigAlertSettingsNil(t *testing.T) {
	obj := resource.AlertSettingsModel{}
	config := &models.AlertSettingsConfig{}

	diags := ExpandAlertSettingsConfig(&obj, config)

	testutil.AssertDiagsHasNoErrors(t, diags)
	// Config should remain unchanged when object is null
}

func TestExpandConfigAlertSettingsValidAlertSettingType(t *testing.T) {
	obj := resource.AlertSettingsModel{
		AlertSettingType:  types.StringValue("inherit"),
		AlertRule:         make([]resource.AlertRuleModel, 0),
		NotificationGroup: nil,
	}

	config := &models.AlertSettingsConfig{}

	diags := ExpandAlertSettingsConfig(&obj, config)

	testutil.AssertDiagsHasNoErrors(t, diags)
	testutil.AssertNotNil(t, "alert setting type", config.AlertSettingType)
}

func TestExpandConfigAlertSettingsWithAlertRules(t *testing.T) {
	// Create a valid alert rule object
	alertRule := resource.AlertRuleModel{
		EnforceTestFailure: types.BoolValue(true),
		OmitScatterplot:    types.BoolValue(false),
		TriggerModel: resource.TriggerModel{
			UseRollingWindow:  types.BoolValue(false),
			WarningReminder:   types.StringValue("10 minutes"),
			CriticalReminder:  types.StringValue("1 hour"),
			ThresholdInterval: types.StringValue("5 minutes"),
			TriggerType:       types.StringValue("specific value"),
			OperationType:     types.StringValue("greater than"),
			Expression:        types.StringValue("avg"),
			WarningTrigger:    types.Float64Value(100.0),
			CriticalTrigger:   types.Float64Value(200.0),
			Level: &resource.LevelModel{
				FilterType:  types.StringValue("name"),
				FilterValue: types.StringValue("test"),
			},
			DNSResolvedName: types.StringValue("example.com"),
			DNSRecordType:   types.StringValue("a"),
			DNSTTL:          types.Int64Value(300),
		},
		NodeThresholdModel: resource.NodeThresholdModel{
			EnableConsecutive:         types.BoolValue(true),
			NumberOfFailingNodes:      types.Int64Value(5),
			NodeThresholdType:         types.StringValue("runs"),
			ThresholdNumberOfRuns:     types.Int64Value(3),
			ConsecutiveNumberOfRuns:   types.Int64Value(2),
			ThresholdPercentageOfRuns: types.Float64Value(75.5),
		},
		NotificationType:  types.StringValue("default contacts"),
		AlertType:         types.StringValue("byte length"),
		AlertSubType:      types.StringValue("response"),
		NotificationGroup: nil,
	}

	alertRules := []resource.AlertRuleModel{alertRule}

	obj := resource.AlertSettingsModel{
		AlertSettingType:  types.StringValue("override"),
		AlertRule:         alertRules,
		NotificationGroup: nil,
	}

	config := &models.AlertSettingsConfig{}

	diags := ExpandAlertSettingsConfig(&obj, config)

	testutil.AssertDiagsHasNoErrors(t, diags)
	testutil.AssertEqual(t, "alert rules count", len(config.AlertRules), 1)

	testutil.AssertEqual(t, fields.EnableConsecutive, config.AlertRules[0].NodeThresholdConfig.ConsecutiveRunsEnabled, true)
	testutil.AssertEqual(t, fields.UseRollingWindow, config.AlertRules[0].Trigger.UseIntervalRollingWindow, false)
	testutil.AssertEqual(t, fields.EnforceTestFailure, config.AlertRules[0].EnforceTestFailure, true)
	testutil.AssertEqual(t, fields.OmitScatterplot, config.AlertRules[0].OmitScatterplot, false)
	testutil.AssertEqual(t, fields.NumberOfFailingNodes, config.AlertRules[0].NodeThresholdConfig.NumberOfFailingUnits, 5)
	testutil.AssertEqual(t, fields.NodeThresholdType, config.AlertRules[0].NodeThresholdConfig.NodeThresholdType.Name, "runs")
	testutil.AssertEqual(t, fields.ThresholdNumberOfRuns, config.AlertRules[0].NodeThresholdConfig.NumberOfUnits, 3)
	testutil.AssertEqual(t, fields.ConsecutiveNumberOfRuns, config.AlertRules[0].NodeThresholdConfig.ConsecutiveRuns, 2)
	testutil.AssertEqual(t, fields.ThresholdPercentageOfRuns, config.AlertRules[0].NodeThresholdConfig.PercentageOfUnits, 75.5)
	testutil.AssertEqual(t, fields.WarningReminder, config.AlertRules[0].Trigger.WarningReminderFrequency.Name, "10 minutes")
	testutil.AssertEqual(t, fields.CriticalReminder, config.AlertRules[0].Trigger.CriticalReminderFrequency.Name, "1 hour")
	testutil.AssertEqual(t, fields.ThresholdInterval, config.AlertRules[0].Trigger.ThresholdInterval.Name, "5 minutes")
	testutil.AssertEqual(t, fields.NotificationType, config.AlertRules[0].NotificationType.Name, "default contacts")
	testutil.AssertEqual(t, fields.AlertType, config.AlertRules[0].AlertType.Name, "byte length")
	testutil.AssertEqual(t, fields.TriggerType, config.AlertRules[0].Trigger.TriggerType.Name, "specific value")
	testutil.AssertEqual(t, fields.OperationType, config.AlertRules[0].Trigger.OperationType.Name, "greater than")
	testutil.AssertEqual(t, fields.AlertSubType, config.AlertRules[0].AlertSubType.Name, "response")
	testutil.AssertEqual(t, fields.Expression, config.AlertRules[0].Trigger.Expression, "avg")
	testutil.AssertEqual(t, fields.WarningTrigger, config.AlertRules[0].Trigger.WarningTrigger, 100.0)
	testutil.AssertEqual(t, fields.CriticalTrigger, config.AlertRules[0].Trigger.CriticalTrigger, 200.0)
	testutil.AssertEqual(t, fields.DNSResolvedName, config.AlertRules[0].Trigger.DNSResolvedName, "example.com")
	testutil.AssertEqual(t, fields.DNSRecordType, config.AlertRules[0].Trigger.DNSRecordType.Name, "a")
	testutil.AssertEqual(t, fields.DNSttl, config.AlertRules[0].Trigger.DNSTTL, 300)
	testutil.AssertEqual(t, fields.FilterType, config.AlertRules[0].Trigger.FilterType.Name, "name")
	testutil.AssertEqual(t, fields.FilterValue, config.AlertRules[0].Trigger.FilterValue, "test")
}

func TestExpandConfigAlertSettingsWithNotificationGroup(t *testing.T) {
	notifGroup := resource.NotificationGroupModel{
		NotifyOnWarning:  types.BoolValue(true),
		NotifyOnCritical: types.BoolValue(true),
		NotifyOnImproved: types.BoolValue(false),
		Subject:          types.StringValue("Test Alert"),
		AlertWebhookIDs:  types.ListNull(types.Int64Type),
		Emails:           types.ListNull(types.StringType),
		ContactGroupIDs:  types.ListNull(types.Int64Type),
	}

	obj := resource.AlertSettingsModel{
		NotificationGroup: &notifGroup,
	}

	config := &models.AlertSettingsConfig{}

	diags := ExpandAlertSettingsConfig(&obj, config)

	testutil.AssertDiagsHasNoErrors(t, diags)
	testutil.AssertEqual(t, "notification group subject", config.NotificationGroup.Subject, "Test Alert")
	testutil.AssertEqual(t, "notify on warning", config.NotificationGroup.NotifyOnWarning, true)
	testutil.AssertEqual(t, "notify on critical", config.NotificationGroup.NotifyOnCritical, true)
	testutil.AssertEqual(t, "notify on improved", config.NotificationGroup.NotifyOnImproved, false)
	testutil.AssertDeepEqual(t, "alert webhook IDs", config.NotificationGroup.WebhookIDs, []int{})
	testutil.AssertDeepEqual(t, "recipient emails", config.NotificationGroup.Emails, []string{})
	testutil.AssertDeepEqual(t, "contact group IDs", config.NotificationGroup.ContactGroupIDs, []int{})
}

func TestExpandAlertRuleConfigAllBooleanFields(t *testing.T) {
	alertRule := resource.AlertRuleModel{
		NodeThresholdModel: resource.NodeThresholdModel{
			EnableConsecutive: types.BoolValue(true),
		},
		TriggerModel: resource.TriggerModel{
			UseRollingWindow: types.BoolValue(true),
		},
		EnforceTestFailure: types.BoolValue(true),
		OmitScatterplot:    types.BoolValue(true),
		NotificationGroup:  nil,
	}

	config := &models.AlertRuleConfig{}

	diags := expandAlertRuleConfig(alertRule, config)

	testutil.AssertDiagsHasNoErrors(t, diags)
	testutil.AssertEqual(t, "enable consecutive", config.NodeThresholdConfig.ConsecutiveRunsEnabled, true)
	testutil.AssertEqual(t, "use rolling window", config.Trigger.UseIntervalRollingWindow, true)
	testutil.AssertEqual(t, "enforce test failure", config.EnforceTestFailure, true)
	testutil.AssertEqual(t, "omit scatterplot", config.OmitScatterplot, true)
}

func TestExpandAlertRuleConfigIntegerFields(t *testing.T) {
	alertRule := resource.AlertRuleModel{
		NodeThresholdModel: resource.NodeThresholdModel{
			NumberOfFailingNodes:    types.Int64Value(10),
			ThresholdNumberOfRuns:   types.Int64Value(5),
			ConsecutiveNumberOfRuns: types.Int64Value(3),
		},
	}

	config := &models.AlertRuleConfig{}

	diags := expandAlertRuleConfig(alertRule, config)

	testutil.AssertDiagsHasNoErrors(t, diags)
	testutil.AssertEqual(t, "number of failing nodes", config.NodeThresholdConfig.NumberOfFailingUnits, 10)
	testutil.AssertEqual(t, "threshold number of runs", config.NodeThresholdConfig.NumberOfUnits, 5)
	testutil.AssertEqual(t, "consecutive number of runs", config.NodeThresholdConfig.ConsecutiveRuns, 3)
}

func TestExpandAlertRuleConfigFloatFields(t *testing.T) {
	alertRule := resource.AlertRuleModel{
		NodeThresholdModel: resource.NodeThresholdModel{
			ThresholdPercentageOfRuns: types.Float64Value(85.5),
		},
		TriggerModel: resource.TriggerModel{
			WarningTrigger:  types.Float64Value(100.0),
			CriticalTrigger: types.Float64Value(200.0),
		},
		NotificationGroup: nil,
	}

	config := &models.AlertRuleConfig{}

	diags := expandAlertRuleConfig(alertRule, config)

	testutil.AssertDiagsHasNoErrors(t, diags)
	testutil.AssertEqual(t, "threshold percentage of runs", config.NodeThresholdConfig.PercentageOfUnits, 85.5)
	testutil.AssertEqual(t, "warning trigger", config.Trigger.WarningTrigger, 100.0)
	testutil.AssertEqual(t, "critical trigger", config.Trigger.CriticalTrigger, 200.0)
}

func TestExpandAlertRuleConfigTrailingValueTriggerType(t *testing.T) {
	alertRule := resource.AlertRuleModel{
		TriggerModel: resource.TriggerModel{
			TriggerType:        types.StringValue("trailing value"),
			HistoricalInterval: types.StringValue("1 hour"),
			StatisticalType:    types.StringValue("average"),
		},
	}

	config := &models.AlertRuleConfig{}

	diags := expandAlertRuleConfig(alertRule, config)

	testutil.AssertDiagsHasNoErrors(t, diags)
	testutil.AssertNotNil(t, "trailing historical interval", config.Trigger.HistoricalInterval)
	testutil.AssertNotNil(t, "statistical type", config.Trigger.StatisticalType)
}

func TestExpandNotificationGroupAllFields(t *testing.T) {
	webhookList, _ := types.ListValue(types.Int64Type, []attr.Value{
		types.Int64Value(1), types.Int64Value(2), types.Int64Value(3),
	})

	emailList, _ := types.ListValue(types.StringType, []attr.Value{
		types.StringValue("test1@example.com"),
		types.StringValue("test2@example.com"),
	})

	contactGroupList, _ := types.ListValue(types.Int64Type, []attr.Value{
		types.Int64Value(10), types.Int64Value(20),
	})

	notifGroup := resource.NotificationGroupModel{
		NotifyOnWarning:  types.BoolValue(true),
		NotifyOnCritical: types.BoolValue(true),
		NotifyOnImproved: types.BoolValue(false),
		Subject:          types.StringValue("Test Alert Subject"),
		AlertWebhookIDs:  webhookList,
		Emails:           emailList,
		ContactGroupIDs:  contactGroupList,
	}

	group := &models.NotificationGroupConfig{}

	diags := expandNotificationGroup(&notifGroup, group)

	testutil.AssertDiagsHasNoErrors(t, diags)
	testutil.AssertEqual(t, "notify on warning", group.NotifyOnWarning, true)
	testutil.AssertEqual(t, "notify on critical", group.NotifyOnCritical, true)
	testutil.AssertEqual(t, "notify on improved", group.NotifyOnImproved, false)
	testutil.AssertEqual(t, "subject", group.Subject, "Test Alert Subject")
	testutil.AssertEqual(t, "webhook ids count", len(group.WebhookIDs), 3)
	testutil.AssertEqual(t, "emails count", len(group.Emails), 2)
	testutil.AssertEqual(t, "contact group ids count", len(group.ContactGroupIDs), 2)
}

func TestExpandNotificationGroupEmptyLists(t *testing.T) {
	notifGroup := resource.NotificationGroupModel{
		NotifyOnWarning:  types.BoolValue(false),
		NotifyOnCritical: types.BoolValue(false),
		NotifyOnImproved: types.BoolValue(false),
		Subject:          types.StringValue(""),
		AlertWebhookIDs:  types.ListNull(types.Int64Type),
		Emails:           types.ListNull(types.StringType),
		ContactGroupIDs:  types.ListNull(types.Int64Type),
	}

	group := &models.NotificationGroupConfig{}

	diags := expandNotificationGroup(&notifGroup, group)

	testutil.AssertDiagsHasNoErrors(t, diags)
	testutil.AssertEqual(t, "webhook ids", len(group.WebhookIDs), 0)
	testutil.AssertEqual(t, "emails", len(group.Emails), 0)
	testutil.AssertEqual(t, "contact group ids", len(group.ContactGroupIDs), 0)
}

func TestExpandAlertRuleConfigWithNotificationGroups(t *testing.T) {
	// Create first notification group
	webhookList1, _ := types.ListValue(types.Int64Type, []attr.Value{types.Int64Value(100), types.Int64Value(200)})
	emailList1, _ := types.ListValue(types.StringType, []attr.Value{types.StringValue("user1@example.com"), types.StringValue("user2@example.com")})
	contactList1, _ := types.ListValue(types.Int64Type, []attr.Value{types.Int64Value(10)})

	notifGroup := resource.NotificationGroupModel{
		NotifyOnWarning:  types.BoolValue(true),
		NotifyOnCritical: types.BoolValue(true),
		NotifyOnImproved: types.BoolValue(false),
		Subject:          types.StringValue("Alert Group 1"),
		AlertWebhookIDs:  webhookList1,
		Emails:           emailList1,
		ContactGroupIDs:  contactList1,
	}

	// Create second notification group
	webhookList2, _ := types.ListValue(types.Int64Type, []attr.Value{types.Int64Value(300)})
	emailList2, _ := types.ListValue(types.StringType, []attr.Value{types.StringValue("admin@example.com")})
	contactList2, _ := types.ListValue(types.Int64Type, []attr.Value{types.Int64Value(20), types.Int64Value(30)})

	notifGroup2 := resource.NotificationGroupModel{
		NotifyOnWarning:  types.BoolValue(false),
		NotifyOnCritical: types.BoolValue(true),
		NotifyOnImproved: types.BoolValue(true),
		Subject:          types.StringValue("Alert Group 2"),
		AlertWebhookIDs:  webhookList2,
		Emails:           emailList2,
		ContactGroupIDs:  contactList2,
	}

	notifGroups := []resource.NotificationGroupModel{notifGroup, notifGroup2}

	// Create alert rule with notification groups
	alertRule := resource.AlertRuleModel{
		NotificationType:  types.StringValue("default contacts"),
		AlertType:         types.StringValue("timing"),
		AlertSubType:      types.StringValue("response"),
		NotificationGroup: notifGroups,
	}

	config := &models.AlertRuleConfig{}

	diags := expandAlertRuleConfig(alertRule, config)

	testutil.AssertDiagsHasNoErrors(t, diags)

	// Verify basic alert rule fields
	testutil.AssertEqual(t, "notification type name", config.NotificationType.Name, "default contacts")
	testutil.AssertEqual(t, "alert type name", config.AlertType.Name, "timing")
	testutil.AssertEqual(t, "alert sub type name", config.AlertSubType.Name, "response")

	// Verify notification groups count
	testutil.AssertEqual(t, "notification groups length", len(config.NotificationGroups), 2)

	// Find and verify notification groups by subject (since order isn't guaranteed)
	var group1, group2 *models.NotificationGroupConfig
	for i := range config.NotificationGroups {
		group := &config.NotificationGroups[i]
		switch group.Subject {
		case "Alert Group 1":
			group1 = group
		case "Alert Group 2":
			group2 = group
		}
	}

	// Verify first notification group
	testutil.AssertNotNil(t, "notification group 1 should exist", group1)
	if group1 != nil {
		testutil.AssertEqual(t, "group1 notify on warning", group1.NotifyOnWarning, true)
		testutil.AssertEqual(t, "group1 notify on critical", group1.NotifyOnCritical, true)
		testutil.AssertEqual(t, "group1 notify on improved", group1.NotifyOnImproved, false)
		testutil.AssertEqual(t, "group1 subject", group1.Subject, "Alert Group 1")
		testutil.AssertEqual(t, "group1 webhook ids length", len(group1.WebhookIDs), 2)
		testutil.AssertEqual(t, "group1 emails length", len(group1.Emails), 2)
		testutil.AssertEqual(t, "group1 contact group ids length", len(group1.ContactGroupIDs), 1)

		// Verify specific values
		testutil.AssertContains(t, "group1 webhook ids", group1.WebhookIDs, 100)
		testutil.AssertContains(t, "group1 webhook ids", group1.WebhookIDs, 200)
		testutil.AssertContains(t, "group1 emails", group1.Emails, "user1@example.com")
		testutil.AssertContains(t, "group1 emails", group1.Emails, "user2@example.com")
		testutil.AssertContains(t, "group1 contact group ids", group1.ContactGroupIDs, 10)
	}

	// Verify second notification group
	testutil.AssertNotNil(t, "notification group 2 should exist", group2)
	if group2 != nil {
		testutil.AssertEqual(t, "group2 notify on warning", group2.NotifyOnWarning, false)
		testutil.AssertEqual(t, "group2 notify on critical", group2.NotifyOnCritical, true)
		testutil.AssertEqual(t, "group2 notify on improved", group2.NotifyOnImproved, true)
		testutil.AssertEqual(t, "group2 subject", group2.Subject, "Alert Group 2")
		testutil.AssertEqual(t, "group2 webhook ids length", len(group2.WebhookIDs), 1)
		testutil.AssertEqual(t, "group2 emails length", len(group2.Emails), 1)
		testutil.AssertEqual(t, "group2 contact group ids length", len(group2.ContactGroupIDs), 2)

		// Verify specific values
		testutil.AssertContains(t, "group2 webhook ids", group2.WebhookIDs, 300)
		testutil.AssertContains(t, "group2 emails", group2.Emails, "admin@example.com")
		testutil.AssertContains(t, "group2 contact group ids", group2.ContactGroupIDs, 20)
		testutil.AssertContains(t, "group2 contact group ids", group2.ContactGroupIDs, 30)
	}
}

func TestExpandAlertRuleConfigEmptyNotificationGroups(t *testing.T) {
	// Create alert rule with empty notification groups
	alertRule := resource.AlertRuleModel{
		NotificationType:  types.StringValue("email"),
		AlertType:         types.StringValue("timing"),
		NotificationGroup: []resource.NotificationGroupModel{},
	}

	config := &models.AlertRuleConfig{}

	diags := expandAlertRuleConfig(alertRule, config)

	testutil.AssertDiagsHasNoErrors(t, diags)
	testutil.AssertEqual(t, "notification groups should be empty", len(config.NotificationGroups), 0)
}
