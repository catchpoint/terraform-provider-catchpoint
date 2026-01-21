package service

import (
	"testing"

	"catchpoint-provider/internal/helpers"
	"catchpoint-provider/internal/labels"
	"catchpoint-provider/internal/models"
	"catchpoint-provider/internal/testutil"
	"catchpoint-provider/internal/types"
)

const (
	errorUnmarshal       = "Failed to unmarshal JSON: %v"
	errorEmptyJSONString = "Expected non-empty JSON string"
)

const (
	alertSubject = "Test Subject"
)

// #region Comparison Helpers

func compareWebhooks(t *testing.T, webhooks []models.AlertWebhookJSON, configWebhookIds []int) {
	if len(webhooks) != len(configWebhookIds) {
		t.Errorf("Webhook count mismatch: got %d, want %d", len(webhooks), len(configWebhookIds))
	}

	webhookSet := make(map[int]bool)
	for _, id := range configWebhookIds {
		webhookSet[id] = false
	}

	for _, webhook := range webhooks {
		if _, ok := webhookSet[*webhook.ID]; ok {
			webhookSet[*webhook.ID] = true
		} else {
			t.Errorf("Unexpected webhook ID: %d", *webhook.ID)
		}
	}

	for id, found := range webhookSet {
		if !found {
			t.Errorf("Expected webhook ID not found: %d", id)
		}
	}
}

func compareRecipients(t *testing.T, got []models.RecipientJSON, emails []string, groups []int) {
	expectedCount := len(emails) + len(groups)
	if len(got) != expectedCount {
		t.Errorf("Recipients count mismatch: got %d, want %d", len(got), expectedCount)
		return
	}

	// Extract emails and group IDs from recipients
	gotEmails := []string{}
	gotGroups := []int{}

	for _, r := range got {
		if r.Email != types.EmptyString {
			gotEmails = append(gotEmails, r.Email)
		}
		if r.ID != nil && *r.ID != 0 {
			gotGroups = append(gotGroups, *r.ID)
		}

	}
	// Compare using existing testutil function
	testutil.AssertDeepEqual(t, "Recipient emails", gotEmails, emails)
	testutil.AssertDeepEqual(t, "Recipient groups", gotGroups, groups)
}

func compareRequestSettingFields(t *testing.T, testObj *models.RequestSettingsJSON, config *models.CommonConfig) {
	if testObj.LibraryCertificateIDs != nil {
		testutil.AssertDeepEqual(t, "LibraryCertificateIDs", *testObj.LibraryCertificateIDs, config.RequestSettingsConfig.CertificateIDs)
	}
	if testObj.Authentication.AuthenticationMethodType != nil && config.RequestSettingsConfig.AuthenticationType != (models.IDNameConfig{}) {
		testutil.AssertEqual(t, "AuthenticationMethodType.ID", *testObj.Authentication.AuthenticationMethodType.ID, config.RequestSettingsConfig.AuthenticationType.ID)
		testutil.AssertEqual(t, "AuthenticationMethodType.Name", *testObj.Authentication.AuthenticationMethodType.Name, config.RequestSettingsConfig.AuthenticationType.Name)
	}

	if config.RequestSettingsConfig.PasswordIDs != nil {
		testutil.AssertDeepEqual(t, "PasswordIDs", *testObj.Authentication.PasswordIDs, config.RequestSettingsConfig.PasswordIDs)
	}

	if testObj.TokenIDs != nil {
		testutil.AssertDeepEqual(t, "TokenIDs", *testObj.TokenIDs, config.RequestSettingsConfig.TokenIDs)
	}
	testutil.AssertEqual(t, "RequestSettingType.ID", testObj.RequestSettingType.ID, config.RequestSettingsConfig.RequestSettingType.ID)
	testutil.AssertEqual(t, "RequestSettingType.Name", testObj.RequestSettingType.Name, config.RequestSettingsConfig.RequestSettingType.Name)
}

func compareInsightDataStructFields(t *testing.T, testObj *models.InsightDataJSON, config *models.CommonConfig) {
	testutil.AssertEqual(t, "InsightSettingType.ID", testObj.InsightSettingType.ID, config.InsightSettingsConfig.InsightSettingType.ID)
	testutil.AssertEqual(t, "InsightSettingType.Name", testObj.InsightSettingType.Name, config.InsightSettingsConfig.InsightSettingType.Name)
	testutil.AssertDeepEqual(t, "Indicators", helpers.FlattenToIDs(testObj.Indicators, func(x models.GenericIDNameJSON) int { return x.ID }), config.InsightSettingsConfig.IndicatorIDs)
	testutil.AssertDeepEqual(t, "Tracepoints", helpers.FlattenToIDs(testObj.Tracepoints, func(x models.GenericIDNameJSON) int { return x.ID }), config.InsightSettingsConfig.TracepointIDs)
}

func compareScheduleSettingSectionFields(t *testing.T, testObj *models.ScheduleSettingsJSON, config *models.CommonConfig) {
	testutil.AssertEqual(t, "ScheduleSettingType.ID", testObj.ScheduleSettingType.ID, config.ScheduleSettingsConfig.ScheduleSettingType.ID)
	testutil.AssertEqual(t, "ScheduleSettingType.Name", testObj.ScheduleSettingType.Name, config.ScheduleSettingsConfig.ScheduleSettingType.Name)
	testutil.AssertDeepEqual(t, "Nodes", helpers.FlattenToIDs(testObj.Nodes, func(x models.NodeJSON) int { return *x.ID }), config.ScheduleSettingsConfig.NodeIDs)
	testutil.AssertEqual(t, "Frequency.ID", testObj.Frequency.ID, config.ScheduleSettingsConfig.Frequency.ID)
	testutil.AssertEqual(t, "Frequency.Name", testObj.Frequency.Name, config.ScheduleSettingsConfig.Frequency.Name)
	testutil.AssertEqual(t, "TestNodeDistribution.ID", testObj.TestNodeDistribution.ID, config.ScheduleSettingsConfig.NodeDistribution.ID)
	testutil.AssertEqual(t, "TestNodeDistribution.Name", testObj.TestNodeDistribution.Name, config.ScheduleSettingsConfig.NodeDistribution.Name)
}

func compareAdvancedSettingsFields(t *testing.T, testObj *models.AdvancedSettingsJSON, config *models.CommonConfig) {
	testutil.AssertEqual(t, "AdvancedSettingType.ID", testObj.AdvancedSettingType.ID, config.AdvancedSettingsConfig.AdvancedSettingType.ID)
	testutil.AssertEqual(t, "AdvancedSettingType.Name", testObj.AdvancedSettingType.Name, config.AdvancedSettingsConfig.AdvancedSettingType.Name)
	testutil.AssertDeepEqual(t, "AppliedTestFlags", helpers.FlattenToIDs(testObj.AppliedTestFlags, func(x models.GenericIDNameOmitEmptyJSON) int { return *x.ID }), config.AdvancedSettingsConfig.AppliedTestFlags)
	testutil.AssertEqual(t, "EDNSSubnet", *testObj.EDNSSubnet, config.AdvancedSettingsConfig.EDNSSubnet)
	testutil.AssertEqual(t, "MaxStepRuntimeSecOverride", *testObj.MaxStepRuntimeSecOverride, config.AdvancedSettingsConfig.MaxStepRuntimeSecOverride)
	testutil.AssertEqual(t, "FailureHopCount", *testObj.FailureHopCount, config.AdvancedSettingsConfig.TracerouteFailureHopCount)
	testutil.AssertEqual(t, "PingCount", *testObj.PingCount, config.AdvancedSettingsConfig.TraceroutePingCount)
	testutil.AssertEqual(t, "ViewportHeight", *testObj.ViewportHeight, config.AdvancedSettingsConfig.ViewportHeight)
	testutil.AssertEqual(t, "ViewportWidth", *testObj.ViewportWidth, config.AdvancedSettingsConfig.ViewportWidth)

	if testObj.WaitForNoActivity != nil && config.AdvancedSettingsConfig.WaitForNoActivityOnDocComplete != nil {
		testutil.AssertEqual(t, "WaitForNoActivity", *testObj.WaitForNoActivity, *config.AdvancedSettingsConfig.WaitForNoActivityOnDocComplete)
	}

	if testObj.AdditionalMonitor != nil && config.AdvancedSettingsConfig.AdditionalMonitorType != (models.IDNameConfig{}) {
		testutil.AssertEqual(t, "AdditionalMonitor.ID", *testObj.AdditionalMonitor.ID, config.AdvancedSettingsConfig.AdditionalMonitorType.ID)
		testutil.AssertEqual(t, "AdditionalMonitor.Name", *testObj.AdditionalMonitor.Name, config.AdvancedSettingsConfig.AdditionalMonitorType.Name)
	}

	if testObj.TestBandwidthThrottling != nil && config.AdvancedSettingsConfig.BandwidthThrottling != (models.IDNameConfig{}) {
		testutil.AssertEqual(t, "TestBandwidthThrottling.ID", *testObj.TestBandwidthThrottling.ID, config.AdvancedSettingsConfig.BandwidthThrottling.ID)
		testutil.AssertEqual(t, "TestBandwidthThrottling.Name", *testObj.TestBandwidthThrottling.Name, config.AdvancedSettingsConfig.BandwidthThrottling.Name)
	}
}

// #endregion

// #region New Objects

// Create a test HTTP header request for testing purposes.
func newTestHTTPHeaders() []models.TestHTTPHeaderRequestConfig {
	return []models.TestHTTPHeaderRequestConfig{
		{
			RequestHeaderType: models.IDNameConfig{
				ID:   1,
				Name: "Content-Type",
			},
			RequestValue:     "application/json",
			ChildHostPattern: "*.example.com",
			HeaderName:       "X-Example-Header",
		},
	}
}

func newCommonConfig() models.CommonConfig {
	return models.CommonConfig{
		DivisionID: 1,
		AlertSettingsConfig: models.AlertSettingsConfig{
			AlertSettingType: models.IDNameConfig{ID: 2, Name: "CustomAlertSetting"},
			NotificationGroup: models.NotificationGroupConfig{
				Subject:         alertSubject,
				Emails:          []string{"user1@example.com", "user2@example.com"},
				ContactGroupIDs: []int{123, 456},
				WebhookIDs:      []int{101, 202},
			},
			AlertRules: []models.AlertRuleConfig{
				newAlertRuleConfig(),
			},
		},
		RequestSettingsConfig: models.RequestSettingsConfig{
			RequestSettingType:     models.IDNameConfig{ID: 1, Name: "Override"},
			TestHTTPHeaderRequests: newTestHTTPHeaders(),
			AuthenticationType: models.IDNameConfig{
				ID:   3,
				Name: "Basic",
			},
			PasswordIDs:    []int{1, 2, 3},
			TokenIDs:       []int{4, 5, 6},
			CertificateIDs: []int{7, 8, 9},
		},
		InsightSettingsConfig: models.InsightSettingsConfig{
			InsightSettingType: models.IDNameConfig{ID: 1, Name: "Override"},
		},
		ScheduleSettingsConfig: models.ScheduleSettingsConfig{
			ScheduleSettingType: models.IDNameConfig{ID: 1, Name: "Override"},
			Frequency:           models.IDNameConfig{ID: 0, Name: labels.None},
			NodeDistribution:    models.IDNameConfig{ID: 0, Name: labels.Random},
		},
		AdvancedSettingsConfig: models.AdvancedSettingsConfig{
			AdvancedSettingType: models.IDNameConfig{ID: 1, Name: "Override"},
			VerifyTestOnFailure: true,
		},
	}
}

func newCommonConfigWithAdvancedSettings() models.CommonConfig {
	return models.CommonConfig{
		AdvancedSettingsConfig: models.AdvancedSettingsConfig{
			AdditionalMonitorType: models.IDNameConfig{
				ID:   1,
				Name: "Additional Monitor",
			},
			BandwidthThrottling: models.IDNameConfig{
				ID:   2,
				Name: "Bandwidth Throttling",
			},
			AdvancedSettingType:            models.IDNameConfig{ID: 1, Name: "Override"},
			AppliedTestFlags:               []int{8},
			MaxStepRuntimeSecOverride:      60,
			WaitForNoActivityOnDocComplete: testutil.ToIntPtr(1),
			ViewportHeight:                 800,
			ViewportWidth:                  600,
			TracerouteFailureHopCount:      5,
			TraceroutePingCount:            4,
			EDNSSubnet:                     "0.0.0.0/0",
		},
	}
}

// #endregion
