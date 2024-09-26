package catchpoint

import (
	"strings"
)

func flattenLabels(labels []Label) []interface{} {
	labelMaps := make([]interface{}, len(labels))
	for i, label := range labels {
		labelMaps[i] = map[string]interface{}{
			//"color":  label.Color,
			"key":    label.Name,
			"values": label.Values,
		}
	}
	return labelMaps
}

func flattenThresholds(thresholds Thresholds) []interface{} {

	if thresholds == (Thresholds{}) {
		return nil
	}

	thresholdsMap := map[string]interface{}{
		"test_time_warning":     thresholds.TestTimeApdexThresholdWarning,
		"test_time_critical":    thresholds.TestTimeApdexThresholdCritical,
		"availability_warning":  thresholds.AvailabilityApdexThresholdWarning,
		"availability_critical": thresholds.AvailabilityApdexThresholdCritical,
	}
	return []interface{}{thresholdsMap}
}
func flattenHttpHeaderRequests(requestSetting RequestSetting) []interface{} {
	httpHeaderRequests := make([]interface{}, 0, len(requestSetting.HttpHeaderRequests))
	for _, header := range requestSetting.HttpHeaderRequests {
		// get the header type name
		key := getReqHeaderTypeName(header.RequestHeaderType.Id)
		userAgentHeader := map[string]interface{}{
			"value":              header.RequestValue,
			"child_host_pattern": header.ChildHostPattern,
		}
		httpHeaderRequests = append(httpHeaderRequests, map[string]interface{}{
			key: []interface{}{userAgentHeader},
		})
	}
	return httpHeaderRequests
}
func flattenRequestSetting(requestSetting RequestSetting) []interface{} {
	requestSettingMap := map[string]interface{}{
		"authentication":       flattenAuthenticationStruct(requestSetting.Authentication),
		"http_request_headers": flattenHttpHeaderRequests(requestSetting),
	}
	if len(requestSetting.LibraryCertificateIds) > 0 {
		requestSettingMap["library_certificate_ids"] = requestSetting.LibraryCertificateIds
	}
	if len(requestSetting.TokenIds) > 0 {
		requestSettingMap["token_ids"] = requestSetting.TokenIds
	}
	return []interface{}{requestSettingMap}
}

func flattenAuthenticationStruct(authentication *AuthenticationStruct) []interface{} {
	if authentication != nil {
		authMap := map[string]interface{}{
			"authentication_type": strings.ToLower(authentication.AuthenticationMethodType.Name),
			"password_ids":        authentication.PasswordIds,
		}
		return []interface{}{authMap}
	}
	return nil
}

func flattenInsightDataStruct(insightData InsightDataStruct) []interface{} {

	if len(insightData.Indicators) == 0 || len(insightData.Tracepoints) == 0 {
		return nil
	}

	insightMap := map[string]interface{}{
		"indicator_ids":  flattenGenericIdNamesToIds(insightData.Indicators),
		"tracepoint_ids": flattenGenericIdNamesToIds(insightData.Tracepoints),
	}
	return []interface{}{insightMap}
}

func flattenScheduleSetting(scheduleSetting ScheduleSetting) []interface{} {
	scheduleMap := map[string]interface{}{
		"run_schedule_id":         scheduleSetting.RunScheduleId,
		"maintenance_schedule_id": scheduleSetting.MaintenanceScheduleId,
		"frequency":               getFrequencyName(scheduleSetting.Frequency.Id),
		"node_distribution":       getNodeDistributionName(scheduleSetting.TestNodeDistribution.Id),
		"no_of_subset_nodes":      scheduleSetting.NoOfSubsetNodes,
		//Backbone network type is currently supported. Remove comment if more types are added
		//"network_type":            flattenGenericIdName(scheduleSetting.NetworkType),
	}

	nodes := make([]int, len(scheduleSetting.Nodes))
	for i, node := range scheduleSetting.Nodes {
		nodes[i] = node.Id
	}

	if len(nodes) > 0 {
		scheduleMap["node_ids"] = nodes
	}

	nodeGroups := make([]int, len(scheduleSetting.NodeGroups))
	for i, group := range scheduleSetting.NodeGroups {
		nodeGroups[i] = group.Id
	}

	if len(nodeGroups) > 0 {
		scheduleMap["node_group_ids"] = nodeGroups
	}

	return []interface{}{scheduleMap}
}

func flattenAdvancedSetting(advancedSetting AdvancedSetting) []interface{} {

	additionalMonitor := ""
	if advancedSetting.AdditionalMonitor != nil {
		additionalMonitor = getAdditionalMonitorTypeName(advancedSetting.AdditionalMonitor.Id)
	}

	testBandwidthThrottling := ""
	if advancedSetting.TestBandwidthThrottling != nil {
		testBandwidthThrottling = getBandwidthThrottlingTypeName(advancedSetting.TestBandwidthThrottling.Id)
	}

	advSettingMap := map[string]interface{}{
		"additional_monitor":   additionalMonitor,
		"bandwidth_throttling": testBandwidthThrottling,
	}

	if advancedSetting.MaxStepRuntimeSecOverride != 0 {
		advSettingMap["enforce_test_failure_if_runs_longer_than"] = advancedSetting.MaxStepRuntimeSecOverride
	}

	if advancedSetting.WaitForNoActivity != nil {
		advSettingMap["wait_for_no_activity"] = advancedSetting.WaitForNoActivity
	}

	if advancedSetting.ViewportHeight != 0 {
		advSettingMap["viewport_height"] = advancedSetting.ViewportHeight
	}

	if advancedSetting.ViewportWidth != 0 {
		advSettingMap["viewport_width"] = advancedSetting.ViewportWidth
	}

	if advancedSetting.FailureHopCount != 0 {
		advSettingMap["failure_hop_count"] = advancedSetting.FailureHopCount
	}

	if advancedSetting.PingCount != 0 {
		advSettingMap["ping_count"] = advancedSetting.PingCount
	}

	if advancedSetting.EdnsSubnet != "" {
		advSettingMap["edns_subnet"] = advancedSetting.EdnsSubnet
	}

	for _, flag := range advancedSetting.AppliedTestFlags {
		lowerCasedSpaceReplacedFlagName := getTestFlagName(flag.Id)
		if lowerCasedSpaceReplacedFlagName != "" {
			advSettingMap[lowerCasedSpaceReplacedFlagName] = true
		}
	}

	return []interface{}{advSettingMap}
}

func flattenGenericIdNamesToIds(genericIdNames []GenericIdName) []int {
	var ids []int

	for _, idName := range genericIdNames {
		ids = append(ids, flattenGenericIdName(idName)["id"].(int))
	}
	return ids
}

func flattenGenericIdName(idName GenericIdName) map[string]interface{} {
	return map[string]interface{}{
		"id":   idName.Id,
		"name": idName.Name,
	}
}

func flattenRecipient(recipient Recipient) map[string]interface{} {
	return map[string]interface{}{
		"id":            recipient.Id,
		"email":         recipient.Email,
		"recipientType": recipient.RecipientType.Name,
		"name":          recipient.Name,
	}
}

func flattenNotificationGroup(notificationGroup NotificationGroupStruct, includeNotify bool) []interface{} {
	alertWebhooks := make([]int, len(notificationGroup.AlertWebhooks))

	for i, webhook := range notificationGroup.AlertWebhooks {
		alertWebhooks[i] = webhook.Id
	}

	var recipients []string
	var contactGroups []string
	for _, recipient := range notificationGroup.Recipients {
		recipientFlattened := flattenRecipient(recipient)
		var value = recipientFlattened["email"].(string)
		if isValidEmail(value) {
			recipients = append(recipients, value)
		} else {
			contactGroups = append(contactGroups, value)
		}
	}

	notifGroupMap := map[string]interface{}{
		"recipient_email_ids": recipients,
		"subject":             notificationGroup.Subject,
		"contact_groups":      contactGroups,
	}

	if includeNotify {
		notifGroupMap["notify_on_warning"] = notificationGroup.NotifyOnWarning
		notifGroupMap["notify_on_critical"] = notificationGroup.NotifyOnCritical
		notifGroupMap["notify_on_improved"] = notificationGroup.NotifyOnImproved
	}

	if len(alertWebhooks) > 0 {
		notifGroupMap["alert_webhook_ids"] = alertWebhooks
	}
	return []interface{}{notifGroupMap}
}

func flattenAlertRuleNotificationGroup(notificationGroups []NotificationGroupStruct, includeNotify bool) []interface{} {
	var flattenedGroups []interface{}

	for _, notificationGroup := range notificationGroups {
		alertWebhooks := make([]int, len(notificationGroup.AlertWebhooks))

		for i, webhook := range notificationGroup.AlertWebhooks {
			alertWebhooks[i] = webhook.Id
		}

		var recipients []string
		var contactGroups []string
		for _, recipient := range notificationGroup.Recipients {
			recipientFlattened := flattenRecipient(recipient)
			var value = recipientFlattened["email"].(string)
			if isValidEmail(value) {
				recipients = append(recipients, value)
			} else {
				contactGroups = append(contactGroups, value)
			}
		}

		notifGroupMap := map[string]interface{}{
			"recipient_email_ids": recipients,
			"subject":             notificationGroup.Subject,
			"contact_groups":      contactGroups,
		}

		if includeNotify {
			notifGroupMap["notify_on_warning"] = notificationGroup.NotifyOnWarning
			notifGroupMap["notify_on_critical"] = notificationGroup.NotifyOnCritical
			notifGroupMap["notify_on_improved"] = notificationGroup.NotifyOnImproved
		}

		if len(alertWebhooks) > 0 {
			notifGroupMap["alert_webhook_ids"] = alertWebhooks
		}

		flattenedGroups = append(flattenedGroups, notifGroupMap)
	}

	return flattenedGroups
}

func flattenAlertGroupItem(alertGroupItem AlertGroupItem) map[string]interface{} {
	nodeThreshold := alertGroupItem.NodeThreshold
	trigger := alertGroupItem.Trigger
	alertGroupItemMap := map[string]interface{}{
		//Only default contacts notification type is supported as of now
		"notification_type":            "default contacts",
		"alert_type":                   getAlertTypeName(alertGroupItem.AlertType.Id),
		"enforce_test_failure":         alertGroupItem.EnforceTestFailure,
		"omit_scatterplot":             alertGroupItem.OmitScatterplot,
		"node_threshold_type":          getNodeThresholdTypeName(nodeThreshold.NodeThresholdType.Id),
		"threshold_number_of_runs":     nodeThreshold.NumberOfUnits,
		"consecutive_number_of_runs":   nodeThreshold.NumberOfConsecutiveRuns,
		"threshold_percentage_of_runs": nodeThreshold.PercentageOfUnits,
		"number_of_failing_nodes":      nodeThreshold.NumberOfFailingUnits,
		"enable_consecutive":           nodeThreshold.ConsecutiveRunsEnabled,
		"warning_reminder":             getReminderName(trigger.WarningReminderFrequency.Id),
		"critical_reminder":            getReminderName(trigger.CriticalReminderFrequency.Id),
		"trigger_type":                 getTriggerTypeName(trigger.TriggerType.Id),
		"operation_type":               getOperationTypeName(trigger.OperationType.Id),
		"threshold_interval":           getThresholdIntervalName(trigger.ThresholdInterval.Id),
		"warning_trigger":              trigger.WarningTrigger,
		"critical_trigger":             trigger.CriticalTrigger,
		"use_rolling_window":           trigger.UseIntervalRollingWindow,
		"expression":                   trigger.Expression,
	}

	if len(alertGroupItem.NotificationGroups) > 0 {
		alertGroupItemMap["notification_group"] = flattenAlertRuleNotificationGroup(alertGroupItem.NotificationGroups, true)
	}

	if alertGroupItem.AlertSubType != nil {
		alertGroupItemMap["alert_sub_type"] = getAlertSubTypeName(alertGroupItem.AlertSubType.Id)
	}
	if trigger.StatisticalType != nil {
		alertGroupItemMap["statistical_type"] = strings.ToLower(trigger.StatisticalType.Name)
	}
	if trigger.HistoricalInterval != nil {
		alertGroupItemMap["historical_interval"] = getHistoricalIntervalName(trigger.HistoricalInterval.Id)
	}

	return alertGroupItemMap
}

func flattenAlertGroupStruct(alertGroup AlertGroupStruct) []interface{} {
	alertGroupItems := make([]interface{}, len(alertGroup.AlertGroupItems))
	for i, item := range alertGroup.AlertGroupItems {
		alertGroupItems[i] = flattenAlertGroupItem(item)
	}

	alertGroupMap := map[string]interface{}{
		"notification_group": flattenNotificationGroup(alertGroup.NotificationGroup, false),
		"alert_rule":         alertGroupItems,
	}
	return []interface{}{alertGroupMap}
}

func flattenTest(test *Test) map[string]interface{} {
	dnsQueryType := ""
	userAgentType := ""
	chromeVersion := ""
	if test.DnsQueryType != nil {
		dnsQueryType = getDnsQueryName(test.DnsQueryType.Id)
	}
	if test.UserAgentType != nil {
		userAgentType = getUserAgentTypeName(test.UserAgentType.Id)
	}
	if test.ChromeMonitorVersion != nil {
		chromeVersion = strings.ToLower(test.ChromeMonitorVersion.ApplicationVersionType.Name)
	}
	testMap := map[string]interface{}{
		"id":                              test.Id,
		"division_id":                     test.DivisionId,
		"product_id":                      test.ProductId,
		"folder_id":                       test.FolderId,
		"test_name":                       test.Name,
		"test_description":                test.Description,
		"test_url":                        test.Url,
		"gateway_address_or_host":         test.GatewayAddressOrHost,
		"label":                           flattenLabels(test.Labels),
		"thresholds":                      flattenThresholds(test.TestThresholds),
		"enforce_certificate_pinning":     test.EnforceCertificatePinning,
		"enforce_certificate_key_pinning": test.EnforceCertificateKeyPinning,
		"enable_test_data_webhook":        test.EnableTestDataWebhook,
		"alerts_paused":                   test.AlertsPaused,
		"change_date":                     test.ChangeDate,
		"start_time":                      test.StartTime,
		"end_time":                        test.EndTime,
		"status":                          strings.ToLower(test.Status.Name),
		"monitor":                         getMonitorName(test.Monitor.Id),
		"dns_server":                      test.DnsServer,
		"query_type":                      dnsQueryType,
		"user_agent_type":                 userAgentType,
		"chrome_version":                  chromeVersion,
		"request_settings":                flattenRequestSetting(test.RequestSettings),
		"alert_settings":                  flattenAlertGroupStruct(test.AlertGroup),
		"insights":                        flattenInsightDataStruct(test.InsightData),
		"schedule_settings":               flattenScheduleSetting(test.ScheduleSettings),
		"advanced_settings":               flattenAdvancedSetting(test.AdvancedSettings),
	}

	if test.TestRequestData != nil {
		if test.TestRequestData.RequestData != "" {
			testMap["test_script"] = test.TestRequestData.RequestData
		}
		if test.TestRequestData.TransactionScriptType.Name != "" {
			testMap["test_script_type"] = strings.ToLower(test.TestRequestData.TransactionScriptType.Name)
		}
	}
	return testMap
}
