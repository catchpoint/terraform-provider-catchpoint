package transform

import (
	"testing"

	"catchpoint-provider/internal/models"
	"catchpoint-provider/internal/models/resource"
	"catchpoint-provider/internal/testutil"

	"github.com/hashicorp/terraform-plugin-framework/types"
)

func TestJSONToTerraformAlertGroup(t *testing.T) {
	// Create test data with all fields populated
	alertGroupJSON := &models.AlertGroupJSON{
		AlertSettingType: models.GenericIDNameJSON{
			ID:   1,
			Name: "override",
		},
		AlertGroupItems: []models.AlertGroupItemJSON{
			{
				NodeThreshold: models.NodeThresholdJSON{
					NodeThresholdType: models.GenericIDNameJSON{
						ID:   1,
						Name: "average across nodes",
					},
					NumberOfUnits:           testutil.ToIntPtr(5),
					PercentageOfUnits:       testutil.ToFloat64Ptr(75.5),
					NumberOfFailingUnits:    testutil.ToIntPtr(2),
					ConsecutiveRunsEnabled:  true,
					NumberOfConsecutiveRuns: testutil.ToIntPtr(3),
				},
				Trigger: models.TriggerJSON{
					TriggerType: models.GenericIDNameJSON{
						ID:   1,
						Name: "specific value",
					},
					StatisticalType: &models.GenericIDNameOmitEmptyJSON{
						ID:   testutil.ToIntPtr(1),
						Name: testutil.ToStringPtr("average"),
					},
					HistoricalInterval: &models.GenericIDNameOmitEmptyJSON{
						ID:   testutil.ToIntPtr(60),
						Name: testutil.ToStringPtr("1 hour"),
					},
					WarningTrigger:  testutil.ToFloat64Ptr(500.0),
					CriticalTrigger: testutil.ToFloat64Ptr(1000.0),
					Expression:      testutil.ToStringPtr("response time > 500"),
					WarningReminderFrequency: models.GenericIDNameJSON{
						ID:   15,
						Name: "15 minutes",
					},
					CriticalReminderFrequency: models.GenericIDNameJSON{
						ID:   30,
						Name: "30 minutes",
					},
					ThresholdInterval: models.GenericIDNameJSON{
						ID:   5,
						Name: "5 minutes",
					},
					UseIntervalRollingWindow: true,
					FilterType: &models.GenericIDNameOmitEmptyJSON{
						ID: testutil.ToIntPtr(1),
					},
					FilterValue: testutil.ToStringPtr("filter"),
					DNSRecordType: &models.GenericIDNameOmitEmptyJSON{
						ID: testutil.ToIntPtr(1),
					},
					DNSResolvedName: testutil.ToStringPtr("example.com"),
					DNSTTL:          testutil.ToIntPtr(300),
				},
				NotificationType: models.GenericIDNameJSON{
					ID:   0,
					Name: "default contacts",
				},
				AlertType: models.GenericIDNameJSON{
					ID:   7,
					Name: "timing",
				},
				AlertSubType: &models.GenericIDNameOmitEmptyJSON{
					ID:   testutil.ToIntPtr(1),
					Name: testutil.ToStringPtr("byte length"),
				},
				EnforceTestFailure: true,
				OmitScatterplot:    false,
				NotificationGroups: []models.NotificationGroupJSON{
					{
						NotifyOnWarning:  true,
						NotifyOnCritical: true,
						NotifyOnImproved: false,
						Subject:          "Alert: Test Performance Issue",
						AlertWebhooks: []models.AlertWebhookJSON{
							{ID: testutil.ToIntPtr(100)},
							{ID: testutil.ToIntPtr(200)},
						},
						Recipients: []models.RecipientJSON{
							{
								RecipientType: models.GenericIDNameJSON{ID: 2, Name: "email"},
								Email:         "admin@example.com",
								Name:          "Admin User",
							},
							{
								RecipientType: models.GenericIDNameJSON{ID: 1, Name: "ContactGroup"},
								Email:         "",
								Name:          "Operations Team",
								ID:            testutil.ToIntPtr(500),
							},
						},
					},
					{
						NotifyOnWarning:  false,
						NotifyOnCritical: true,
						NotifyOnImproved: true,
						Subject:          "Critical Alert: System Down",
						AlertWebhooks: []models.AlertWebhookJSON{
							{ID: testutil.ToIntPtr(300)},
						},
						Recipients: []models.RecipientJSON{
							{
								RecipientType: models.GenericIDNameJSON{ID: 2, Name: "email"},
								Email:         "oncall@example.com",
								Name:          "On-Call Engineer",
							},
						},
					},
				},
			},
		},
		NotificationGroup: &models.NotificationGroupJSON{
			NotifyOnWarning:  true,
			NotifyOnCritical: true,
			NotifyOnImproved: false,
			Subject:          "Global Alert Notification",
			AlertWebhooks: []models.AlertWebhookJSON{
				{ID: testutil.ToIntPtr(400)},
			},
			Recipients: []models.RecipientJSON{
				{
					RecipientType: models.GenericIDNameJSON{ID: 2, Name: "email"},
					Email:         "global@example.com",
					Name:          "Global Admin",
				},
			},
		},
	}

	// This represents the plan. It must have some sort of value for the transform to work.
	config := &resource.AlertSettingsModel{
		AlertSettingType: types.StringValue("override"),
	}

	// Execute the function
	result, diags := JSONToTerraformAlertGroup(alertGroupJSON, config)

	// Verify no errors occurred
	testutil.AssertDiagsHasNoErrors(t, diags)

	// Verify the main fields
	testutil.AssertEqual(t, "AlertSettingType", result.AlertSettingType, types.StringValue("override"))

	// Verify AlertGroupItems
	if len(result.AlertRule) != 1 {
		t.Fatalf("Expected 1 AlertRule, got %d", len(result.AlertRule))
	}
	alertRule := result.AlertRule[0]

	// Verify NodeThreshold fields
	testutil.AssertEqual(t, "NodeThreshold.NodeThresholdType", alertRule.NodeThresholdType, types.StringValue("average across nodes"))
	testutil.AssertEqual(t, "NodeThreshold.ThresholdNumberOfRuns", alertRule.ThresholdNumberOfRuns, types.Int64Value(5))
	testutil.AssertFloat64Equal(t, "NodeThreshold.ThresholdPercentageOfRuns", alertRule.ThresholdPercentageOfRuns, types.Float64Value(75.5))
	testutil.AssertEqual(t, "NodeThreshold.NumberOfFailingNodes", alertRule.NumberOfFailingNodes, types.Int64Value(2))
	testutil.AssertEqual(t, "NodeThreshold.EnableConsecutive", alertRule.EnableConsecutive, types.BoolValue(true))
	testutil.AssertEqual(t, "NodeThreshold.ConsecutiveNumberOfRuns", alertRule.ConsecutiveNumberOfRuns, types.Int64Value(3))

	// Verify Trigger fields
	testutil.AssertEqual(t, "Trigger.TriggerType", alertRule.TriggerType, types.StringValue("specific value"))
	testutil.AssertEqual(t, "Trigger.StatisticalType", alertRule.StatisticalType, types.StringValue("average"))
	testutil.AssertEqual(t, "Trigger.HistoricalInterval", alertRule.HistoricalInterval, types.StringValue("1 hour"))
	testutil.AssertFloat64Equal(t, "Trigger.WarningTrigger", alertRule.WarningTrigger, types.Float64Value(500.0))
	testutil.AssertFloat64Equal(t, "Trigger.CriticalTrigger", alertRule.CriticalTrigger, types.Float64Value(1000.0))
	testutil.AssertEqual(t, "Trigger.Expression", alertRule.Expression, types.StringValue("response time > 500"))
	testutil.AssertEqual(t, "Trigger.WarningReminder", alertRule.WarningReminder, types.StringValue("15 minutes"))
	testutil.AssertEqual(t, "Trigger.CriticalReminder", alertRule.CriticalReminder, types.StringValue("30 minutes"))
	testutil.AssertEqual(t, "Trigger.ThresholdInterval", alertRule.ThresholdInterval, types.StringValue("5 minutes"))
	testutil.AssertEqual(t, "Trigger.UseRollingWindow", alertRule.UseRollingWindow, types.BoolValue(true))
	testutil.AssertEqual(t, "Trigger.FilterType", alertRule.Level.FilterType, types.StringValue("index"))
	testutil.AssertEqual(t, "Trigger.FilterValue", alertRule.Level.FilterValue, types.StringValue("filter"))

	// Verify AlertType and AlertSubType
	testutil.AssertEqual(t, "AlertType", alertRule.AlertType, types.StringValue("timing"))
	testutil.AssertEqual(t, "AlertSubType", alertRule.AlertSubType, types.StringValue("byte length"))

	// Verify NotificationType
	testutil.AssertEqual(t, "NotificationType", alertRule.NotificationType, types.StringValue("default contacts"))

	// Verify other AlertRule fields
	testutil.AssertEqual(t, "EnforceTestFailure", alertRule.EnforceTestFailure, types.BoolValue(true))
	testutil.AssertEqual(t, "OmitScatterplot", alertRule.OmitScatterplot, types.BoolValue(false))

	// Verify NotificationGroups
	if len(alertRule.NotificationGroup) != 2 {
		t.Fatalf("Expected 2 NotificationGroups, got %d", len(alertRule.NotificationGroup))
	}
	ng1 := alertRule.NotificationGroup[0]
	testutil.AssertEqual(t, "NotificationGroup[0].NotifyOnWarning", ng1.NotifyOnWarning, types.BoolValue(true))
	testutil.AssertEqual(t, "NotificationGroup[0].NotifyOnCritical", ng1.NotifyOnCritical, types.BoolValue(true))
	testutil.AssertEqual(t, "NotificationGroup[0].NotifyOnImproved", ng1.NotifyOnImproved, types.BoolValue(false))
	testutil.AssertEqual(t, "NotificationGroup[0].Subject", ng1.Subject, types.StringValue("Alert: Test Performance Issue"))
	if len(ng1.AlertWebhookIDs.Elements()) != 2 {
		t.Fatalf("Expected 2 AlertWebhookIDs in NotificationGroup[0], got %d", len(ng1.AlertWebhookIDs.Elements()))
	}
	testutil.AssertEqual(t, "NotificationGroup[0].AlertWebhookIDs[0]", ng1.AlertWebhookIDs.Elements()[0].(types.Int64), types.Int64Value(100))
	testutil.AssertEqual(t, "NotificationGroup[0].AlertWebhookIDs[1]", ng1.AlertWebhookIDs.Elements()[1].(types.Int64), types.Int64Value(200))
	if len(ng1.Emails.Elements()) != 1 {
		t.Fatalf("Expected 1 RecipientEmail in NotificationGroup[0], got %d", len(ng1.Emails.Elements()))
	}
	testutil.AssertEqual(t, "NotificationGroup[0].RecipientEmails[0]", ng1.Emails.Elements()[0].(types.String), types.StringValue("admin@example.com"))

	if len(ng1.ContactGroupIDs.Elements()) != 1 {
		t.Fatalf("Expected 1 ContactGroupID in NotificationGroup[0], got %d", len(ng1.ContactGroupIDs.Elements()))
	}
	testutil.AssertEqual(t, "NotificationGroup[0].ContactGroupIDs[0]", ng1.ContactGroupIDs.Elements()[0].(types.Int64), types.Int64Value(500))

	testutil.AssertEqual(t, "DNSRecordType", alertRule.DNSRecordType, types.StringValue("a"))
	testutil.AssertEqual(t, "DNSResolvedName", alertRule.DNSResolvedName, types.StringValue("example.com"))
	testutil.AssertEqual(t, "DNSTTL", alertRule.DNSTTL, types.Int64Value(300))

}
