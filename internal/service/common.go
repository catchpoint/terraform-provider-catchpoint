package service

import (
	"encoding/json"
	"strings"

	"catchpoint-provider/internal/fields"
	"catchpoint-provider/internal/labels"
	"catchpoint-provider/internal/models"
	"catchpoint-provider/internal/types"
	"catchpoint-provider/internal/validation"
)

const (
	notificationSubject = "${NotificationLevel}:  test=#${TestId} - ${TestName}, alert=${AlertType}"
)

func createJSONPatchDocument(config models.ConfigUpdate, path string, isMetaData bool) string {
	var jsonPatchDoc []byte
	operation := "replace"

	if config.GetSectionToUpdate() == fields.LabelsSection {
		jsonPatchObject := models.JSONPatchLabel{
			LabelValue: config.GetUpdatedLabels(),
			Path:       path,
			Op:         operation,
		}
		jsonPatchDoc, _ = json.Marshal(jsonPatchObject)
		return string(jsonPatchDoc)
	}

	if config.GetSectionToUpdate() == fields.ThresholdRestModelSection {
		jsonPatchObject := models.JSONPatchThreshold{
			ThresholdValue: config.GetUpdatedTestThresholds(),
			Path:           path,
			Op:             operation,
		}
		jsonPatchDoc, _ = json.Marshal(jsonPatchObject)
		return string(jsonPatchDoc)
	}

	if config.GetSectionToUpdate() == fields.TestRequestDataSection {
		jsonPatchObject := models.JSONPatchRequestData{
			TestRequestDataValue: config.GetUpdatedTestRequestData(),
			Path:                 path,
			Op:                   operation,
		}
		jsonPatchDoc, _ = json.Marshal(jsonPatchObject)
		return string(jsonPatchDoc)
	}

	if config.GetSectionToUpdate() == fields.ChromeMonitorVersionSection {
		jsonPatchObject := models.JSONPatchChromeVersion{
			ChromeVersionValue: config.GetUpdatedChromeVersionSection(),
			Path:               path,
			Op:                 operation,
		}
		jsonPatchDoc, _ = json.Marshal(jsonPatchObject)
		return string(jsonPatchDoc)
	}

	if isMetaData {
		jsonPatchObject := models.JSONPatch{
			Value: config.GetUpdatedFieldValue(),
			Path:  path,
			Op:    operation,
		}
		jsonPatchDoc, _ = json.Marshal(jsonPatchObject)
		return string(jsonPatchDoc)
	}

	// advancedSettings is for Folder and Test, advancedSettingsModel is for Product
	if strings.Contains(config.GetSectionToUpdate(), fields.AdvancedSettingsSection) || strings.Contains(config.GetSectionToUpdate(), fields.AdvancedSettingsModelSection) {
		jsonPatchObject := models.JSONPatchAdvanced{
			AdvancedSettingValue: config.GetUpdatedAdvancedSettingsSection(),
			Path:                 path,
			Op:                   operation,
		}
		jsonPatchDoc, _ = json.Marshal(jsonPatchObject)
		return string(jsonPatchDoc)
	}
	if strings.Contains(config.GetSectionToUpdate(), fields.RequestSettingSection) || strings.Contains(config.GetSectionToUpdate(), fields.RequestSettingsSection) {
		jsonPatchObject := models.JSONPatchRequest{
			RequestSettingValue: config.GetUpdatedRequestSettingsSection(),
			Path:                path,
			Op:                  operation,
		}
		jsonPatchDoc, _ = json.Marshal(jsonPatchObject)
		return string(jsonPatchDoc)
	}
	if strings.Contains(config.GetSectionToUpdate(), fields.InsightsSection) || strings.Contains(config.GetSectionToUpdate(), fields.InsightDataSection) {
		jsonPatchObject := models.JSONPatchInsight{
			InsightDataValue: config.GetUpdatedInsightSettingsSection(),
			Path:             path,
			Op:               operation,
		}
		jsonPatchDoc, _ = json.Marshal(jsonPatchObject)
		return string(jsonPatchDoc)
	}
	// Check both /scheduleSetting (Folder) and /scheduleSetting*s* (Product, Test)
	if strings.Contains(config.GetSectionToUpdate(), fields.ScheduleSettingSection) || strings.Contains(config.GetSectionToUpdate(), fields.ScheduleSettingsSection) {
		jsonPatchObject := models.JSONPatchSchedule{
			ScheduleSettingValue: config.GetUpdatedScheduleSettingsSection(),
			Path:                 path,
			Op:                   operation,
		}
		jsonPatchDoc, _ = json.Marshal(jsonPatchObject)
		return string(jsonPatchDoc)
	}
	if strings.Contains(config.GetSectionToUpdate(), fields.AlertGroupSection) {
		jsonPatchObject := models.JSONPatchAlert{
			AlertSettingValue: config.GetUpdatedAlertSettingsSection(),
			Path:              path,
			Op:                operation,
		}
		jsonPatchDoc, _ = json.Marshal(jsonPatchObject)
		return string(jsonPatchDoc)
	}
	return types.EmptyString
}

func buildAdvancedSettings(config *models.AdvancedSettingsConfig) models.AdvancedSettingsJSON {
	advancedSettingId := 0
	advancedSettingType := models.GenericIDNameJSON{ID: config.AdvancedSettingType.ID, Name: config.AdvancedSettingType.Name}

	appliedTestFlags := buildAppliedTestFlags(config.AppliedTestFlags)

	advancedSettings := models.AdvancedSettingsJSON{
		AdvancedSettingType: advancedSettingType,
		AppliedTestFlags:    appliedTestFlags,
		ID:                  advancedSettingId,
	}

	setOptionalAdvancedSettings(config, &advancedSettings)

	return advancedSettings
}

func buildAppliedTestFlags(appliedTestFlags []int) []models.GenericIDNameOmitEmptyJSON {
	result := []models.GenericIDNameOmitEmptyJSON{}
	if len(appliedTestFlags) > 0 {
		for i := range appliedTestFlags {
			if appliedTestFlags[i] != 0 {
				if name, ok := types.GetTestFlagName(appliedTestFlags[i]); ok {
					result = append(result, models.GenericIDNameOmitEmptyJSON{ID: &appliedTestFlags[i], Name: &name})
				}
			}
		}
	}
	return result
}

func setOptionalAdvancedSettings(config *models.AdvancedSettingsConfig, advancedSettings *models.AdvancedSettingsJSON) {
	if config.AdditionalMonitorType != (models.IDNameConfig{}) {
		additionalMonitor := models.GenericIDNameOmitEmptyJSON{ID: &config.AdditionalMonitorType.ID, Name: &config.AdditionalMonitorType.Name}
		advancedSettings.AdditionalMonitor = &additionalMonitor
	}

	if config.BandwidthThrottling != (models.IDNameConfig{}) {
		bandwidthThrottling := models.GenericIDNameOmitEmptyJSON{ID: &config.BandwidthThrottling.ID, Name: &config.BandwidthThrottling.Name}
		advancedSettings.TestBandwidthThrottling = &bandwidthThrottling
	}

	if config.ViewportHeight != 0 {
		advancedSettings.ViewportHeight = &config.ViewportHeight
	}

	if config.ViewportWidth != 0 {
		advancedSettings.ViewportWidth = &config.ViewportWidth
	}

	if config.MaxStepRuntimeSecOverride != 0 {
		advancedSettings.MaxStepRuntimeSecOverride = &config.MaxStepRuntimeSecOverride
	}

	if config.TracerouteFailureHopCount != 0 {
		advancedSettings.FailureHopCount = &config.TracerouteFailureHopCount
	}

	if config.TraceroutePingCount != 0 {
		advancedSettings.PingCount = &config.TraceroutePingCount
	}

	if config.WaitForNoActivityOnDocComplete != nil {
		advancedSettings.WaitForNoActivity = config.WaitForNoActivityOnDocComplete
	}

	if config.EDNSSubnet != types.EmptyString {
		advancedSettings.EDNSSubnet = &config.EDNSSubnet
	}
}

func buildAlertSettings(config *models.AlertSettingsConfig) models.AlertGroupJSON {
	alertSettingType := models.GenericIDNameJSON{ID: config.AlertSettingType.ID, Name: config.AlertSettingType.Name}
	alertGroup := models.AlertGroupJSON{
		AlertSettingType:  alertSettingType,
		AlertGroupItems:   nil,
		NotificationGroup: nil,
	}

	if config.AlertSettingType.Name == types.Inherit {
		return alertGroup
	}

	recipients := buildRecipients(&config.NotificationGroup)
	alertWebhooks := buildAlertWebhooks(&config.NotificationGroup)
	alertGroupItems := buildAlertGroupItems(config.AlertRules)

	notifSubject := notificationSubject
	if config.NotificationGroup.Subject != types.EmptyString {
		notifSubject = config.NotificationGroup.Subject
	}

	// For a top-level notification group, the "NotifyXxx" fields should always be true.
	// The struct is the same as for nested groups, but the API expects these to be true.
	// Nested groups may customize the settings.
	notificationGroup := models.NotificationGroupJSON{
		Subject:          notifSubject,
		NotifyOnWarning:  true,
		NotifyOnCritical: true,
		NotifyOnImproved: true,
		AlertWebhooks:    alertWebhooks,
		Recipients:       recipients,
	}

	alertGroup.AlertGroupItems = alertGroupItems
	alertGroup.NotificationGroup = &notificationGroup

	return alertGroup
}

func buildAlertGroupItems(config []models.AlertRuleConfig) []models.AlertGroupItemJSON {
	alertGroupItems := []models.AlertGroupItemJSON{}
	for i := range config {
		// Build the NodeThresholdJSON from the NodeThresholdConfig
		nodeThreshold := buildNodeThresholdConfig(config[i].NodeThresholdConfig)
		trigger := buildTriggerConfig(config[i].Trigger)

		alertType := models.GenericIDNameJSON{ID: config[i].AlertType.ID, Name: config[i].AlertType.Name}
		notificationType := models.GenericIDNameJSON{ID: config[i].NotificationType.ID, Name: config[i].NotificationType.Name}

		notificationGroups := config[i].NotificationGroups
		notificationGroupStructs := []models.NotificationGroupJSON{}
		// Ensure that each notification group has initialized recipients and alert webhooks or get API errors due to null field.
		for j := range notificationGroups {
			notifGroup := models.NotificationGroupJSON{
				Subject:          notificationGroups[j].Subject,
				Recipients:       buildRecipients(&notificationGroups[j]),
				AlertWebhooks:    buildAlertWebhooks(&notificationGroups[j]),
				NotifyOnWarning:  notificationGroups[j].NotifyOnWarning,
				NotifyOnCritical: notificationGroups[j].NotifyOnCritical,
				NotifyOnImproved: notificationGroups[j].NotifyOnImproved,
			}
			notificationGroupStructs = append(notificationGroupStructs, notifGroup)
		}

		alertGroupItem := models.AlertGroupItemJSON{
			NodeThreshold:      nodeThreshold,
			Trigger:            trigger,
			NotificationType:   notificationType,
			AlertType:          alertType,
			EnforceTestFailure: config[i].EnforceTestFailure,
			OmitScatterplot:    config[i].OmitScatterplot,
			MatchAllRecords:    config[i].MatchAllRecords,
			NotificationGroups: notificationGroupStructs,
		}

		if config[i].AlertSubType != (models.IDNameConfig{}) {
			alertGroupItem.AlertSubType = &models.GenericIDNameOmitEmptyJSON{ID: &config[i].AlertSubType.ID, Name: &config[i].AlertSubType.Name}
		}

		alertGroupItems = append(alertGroupItems, alertGroupItem)
	}
	return alertGroupItems
}

func buildNodeThresholdConfig(config models.NodeThresholdConfig) models.NodeThresholdJSON {
	nodeThresholdType := models.GenericIDNameJSON{ID: config.NodeThresholdType.ID, Name: config.NodeThresholdType.Name}
	nodeThreshold := models.NodeThresholdJSON{
		NodeThresholdType:               nodeThresholdType,
		UtilizePerNodeHistoricalAverage: config.UtilizePerNodeHistoricalAverage,
		ConsecutiveRunsEnabled:          config.ConsecutiveRunsEnabled,
	}

	if config.Name != types.EmptyString {
		nodeThreshold.Name = config.Name
	}

	if config.ConsecutiveRuns != 0 {
		nodeThreshold.NumberOfConsecutiveRuns = &config.ConsecutiveRuns
	}

	if config.NumberOfFailingUnits != 0 {
		nodeThreshold.NumberOfFailingUnits = &config.NumberOfFailingUnits
	}

	if config.NumberOfUnits != 0 {
		nodeThreshold.NumberOfUnits = &config.NumberOfUnits
	}

	if config.PercentageOfUnits != 0 {
		nodeThreshold.PercentageOfUnits = &config.PercentageOfUnits
	}

	return nodeThreshold
}

func buildTriggerConfig(config models.TriggerConfig) models.TriggerJSON {
	trigger := models.TriggerJSON{
		ID:                       0,
		UseIntervalRollingWindow: config.UseIntervalRollingWindow,
	}

	assignIDNameJSON(&trigger.CriticalReminderFrequency, config.CriticalReminderFrequency)

	if config.CriticalTrigger != 0 {
		trigger.CriticalTrigger = &config.CriticalTrigger
	}

	assignIDNameJSON(&trigger.WarningReminderFrequency, config.WarningReminderFrequency)

	if config.WarningTrigger != 0 {
		trigger.WarningTrigger = &config.WarningTrigger
	}

	assignIDNameOmitEmptyJSONPtr(&trigger.DNSRecordType, config.DNSRecordType)

	assignStringJSON(&trigger.DNSResolvedName, config.DNSResolvedName)

	if config.DNSTTL != 0 {
		trigger.DNSTTL = &config.DNSTTL
	}

	assignStringJSON(&trigger.Expression, config.Expression)
	assignIDNameOmitEmptyJSONPtr(&trigger.FilterType, config.FilterType)

	assignStringJSON(&trigger.FilterValue, config.FilterValue)

	assignIDNameOmitEmptyJSONPtr(&trigger.StatisticalType, config.StatisticalType)
	assignIDNameOmitEmptyJSONPtr(&trigger.HistoricalInterval, config.HistoricalInterval)

	assignStringJSON(&trigger.Monitor, config.Monitor)

	assignIDNameJSON(&trigger.OperationType, config.OperationType)
	assignIDNameJSON(&trigger.ThresholdInterval, config.ThresholdInterval)
	assignIDNameJSON(&trigger.TriggerType, config.TriggerType)

	return trigger
}

func buildRecipients(config *models.NotificationGroupConfig) []models.RecipientJSON {
	// Always return a non-nil slice
	recipients := make([]models.RecipientJSON, 0)

	if len(config.Emails) > 0 {
		recipientType := models.GenericIDNameJSON{ID: 2, Name: labels.Email}
		for i := range config.Emails {
			recipients = append(recipients, models.RecipientJSON{
				Email:         config.Emails[i],
				RecipientType: recipientType,
			})
		}
	}

	if len(config.ContactGroupIDs) > 0 {
		recipientType := models.GenericIDNameJSON{ID: 1, Name: labels.ContactGroup}
		for i := range config.ContactGroupIDs {
			recipients = append(recipients, models.RecipientJSON{
				ID:            &config.ContactGroupIDs[i],
				RecipientType: recipientType,
			})
		}
	}

	return recipients
}

func buildAlertWebhooks(config *models.NotificationGroupConfig) []models.AlertWebhookJSON {
	alertWebhooks := make([]models.AlertWebhookJSON, 0) // Non-nil empty slice

	if len(config.WebhookIDs) > 0 {
		for i := range config.WebhookIDs {
			alertWebhooks = append(alertWebhooks, models.AlertWebhookJSON{ID: &config.WebhookIDs[i]})
		}
	}

	return alertWebhooks
}

func buildScheduleSettings(config *models.ScheduleSettingsConfig) models.ScheduleSettingsJSON {
	scheduleSettingType := models.GenericIDNameJSON{ID: config.ScheduleSettingType.ID, Name: config.ScheduleSettingType.Name}
	frequency := models.GenericIDNameJSON{ID: config.Frequency.ID, Name: config.Frequency.Name}
	testNodeDistribution := models.GenericIDNameJSON{ID: config.NodeDistribution.ID, Name: config.NodeDistribution.Name}
	networkType := models.GenericIDNameJSON{ID: 0, Name: labels.Backbone}

	nodes, nodeGroups := buildNodesAndNodeGroups(config)

	scheduleSettings := models.ScheduleSettingsJSON{
		ScheduleSettingType:  scheduleSettingType,
		Frequency:            frequency,
		TestNodeDistribution: testNodeDistribution,
		NetworkType:          networkType,
		Nodes:                nodes,
		NodeGroups:           nodeGroups,
	}

	if config.NoOfSubsetNodes > 0 {
		scheduleSettings.NoOfSubsetNodes = &config.NoOfSubsetNodes
	}

	if config.RunScheduleID != 0 {
		scheduleSettings.RunScheduleID = &config.RunScheduleID
	}

	if config.MaintenanceScheduleID != 0 {
		scheduleSettings.MaintenanceScheduleID = &config.MaintenanceScheduleID
	}

	return scheduleSettings
}

func buildNodesAndNodeGroups(config *models.ScheduleSettingsConfig) ([]models.NodeJSON, []models.NodeGroupJSON) {
	networkType := models.GenericIDNameJSON{ID: 0, Name: labels.Backbone}
	nodes := []models.NodeJSON{}
	if len(config.NodeIDs) > 0 {
		for i := range config.NodeIDs {
			nodes = append(nodes, models.NodeJSON{ID: &config.NodeIDs[i], Name: labels.Node, NetworkType: networkType})
		}
	}

	// Initialize an empty slice of NodeGroup
	nodeGroups := []models.NodeGroupJSON{}
	if len(config.NodeGroupIDs) > 0 {
		for i := range config.NodeGroupIDs {
			nodeGroup := models.NodeGroupJSON{
				ID:          &config.NodeGroupIDs[i].ID,
				NodeGroupID: &config.NodeGroupIDs[i].ID,
				Name:        labels.DefaultNodeGroupName,
			}
			nodeGroups = append(nodeGroups, nodeGroup)
		}
	}
	return nodes, nodeGroups
}

func buildInsightSettings(config *models.InsightSettingsConfig) models.InsightDataJSON {
	tracepoints := []models.GenericIDNameJSON{}
	indicators := []models.GenericIDNameJSON{}

	insightSettingType := models.GenericIDNameJSON{ID: config.InsightSettingType.ID, Name: config.InsightSettingType.Name}

	if len(config.TracepointIDs) > 0 {
		for i := range config.TracepointIDs {
			tracepoints = append(tracepoints, models.GenericIDNameJSON{ID: config.TracepointIDs[i], Name: labels.Tracepoint})
		}
	}
	if len(config.IndicatorIDs) > 0 {
		for i := range config.IndicatorIDs {
			indicators = append(indicators, models.GenericIDNameJSON{ID: config.IndicatorIDs[i], Name: labels.Indicator})
		}
	}

	insightData := models.InsightDataJSON{InsightSettingType: insightSettingType, Indicators: indicators, Tracepoints: tracepoints}

	return insightData
}

func buildRequestSettings(config *models.RequestSettingsConfig) models.RequestSettingsJSON {
	httpHeaderRequests := []models.HTTPHeaderRequestJSON{}
	requestSettingType := models.GenericIDNameJSON{ID: config.RequestSettingType.ID, Name: config.RequestSettingType.Name}

	if len(config.TestHTTPHeaderRequests) > 0 {
		for i := range config.TestHTTPHeaderRequests {
			requestHeaderType := models.GenericIDNameJSON{ID: config.TestHTTPHeaderRequests[i].RequestHeaderType.ID, Name: config.TestHTTPHeaderRequests[i].RequestHeaderType.Name}
			httpHeaderRequests = append(httpHeaderRequests, models.HTTPHeaderRequestJSON{
				RequestValue:      config.TestHTTPHeaderRequests[i].RequestValue,
				RequestHeaderType: requestHeaderType,
				ChildHostPattern:  &config.TestHTTPHeaderRequests[i].ChildHostPattern,
				HeaderName:        &config.TestHTTPHeaderRequests[i].HeaderName,
			})
		}
	}

	authentication := buildAuthentication(config)

	requestSetting := models.RequestSettingsJSON{
		RequestSettingType:    requestSettingType,
		HTTPHeaderRequests:    &httpHeaderRequests,
		TokenIDs:              &config.TokenIDs,
		LibraryCertificateIDs: &config.CertificateIDs,
		Authentication:        &authentication,
	}

	return requestSetting
}

func buildAuthentication(config *models.RequestSettingsConfig) models.AuthenticationJSON {
	authenticationID := 0
	authentication := models.AuthenticationJSON{
		ID: &authenticationID,
	}

	if config.PasswordIDs == nil {
		authentication.PasswordIDs = &[]int{}
	} else {
		authentication.PasswordIDs = &config.PasswordIDs
	}

	if config.AuthenticationType != (models.IDNameConfig{}) {
		authentication.AuthenticationMethodType = &models.GenericIDNameOmitEmptyJSON{ID: &config.AuthenticationType.ID, Name: &config.AuthenticationType.Name}
	} else {
		authType := validation.GetAuthenticationTypeOrDefault(types.None)
		authentication.AuthenticationMethodType = &models.GenericIDNameOmitEmptyJSON{ID: &authType.ID, Name: &authType.Name}
	}

	return authentication
}

func assignStringJSON(target **string, source string) {
	if source != types.EmptyString {
		*target = &source
	}
}

func assignIDNameJSON(target *models.GenericIDNameJSON, source models.IDNameConfig) {
	if source != (models.IDNameConfig{}) {
		*target = models.GenericIDNameJSON(source)
	}
}

func assignIDNameOmitEmptyJSONPtr(target **models.GenericIDNameOmitEmptyJSON, source models.IDNameConfig) {
	if source != (models.IDNameConfig{}) {
		*target = &models.GenericIDNameOmitEmptyJSON{
			ID:   &source.ID,
			Name: &source.Name,
		}
	}
}
