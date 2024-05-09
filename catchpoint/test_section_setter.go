package catchpoint

import (
	"errors"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func setRequestData(testTypeId int, testScript string, monitorId int, testScriptTypeId int, testConfig *TestConfig) {
	testConfig.Script.RequestData = testScript
	testConfig.Script.Monitor = monitorId
	testConfig.Script.TransactionScriptType = testScriptTypeId
	testConfig.Script.TestType = testTypeId
	//This is hardcoded because TestId property in TestRequestData object is required to have a value > 0
	testConfig.Script.TestId = 1

}

func setLabels(testTypeId int, labels []interface{}, testConfig *TestConfig) {

	for i := range labels {
		label_map := labels[i].(map[string]interface{})
		label_name := label_map["key"].(string)
		label_values := label_map["values"].([]interface{})
		label_values_list := make([]string, len(label_values))
		for i, label_value := range label_values {
			label_values_list[i] = label_value.(string)
		}
		testConfig.Labels = append(testConfig.Labels, TestLabel{Name: label_name, Values: label_values_list})
	}
}

func setThresholds(testTypeId int, threshold map[string]interface{}, testConfig *TestConfig) {

	test_time_warning := threshold["test_time_warning"].(float64)
	test_time_critical := threshold["test_time_critical"].(float64)
	availability_warning := threshold["availability_warning"].(float64)
	availability_critical := threshold["availability_critical"].(float64)
	testConfig.TestTimeThresholdWarning = test_time_warning
	testConfig.TestTimeThresholdCritical = test_time_critical
	testConfig.AvailabilityThresholdWarning = availability_warning
	testConfig.AvailabilityThresholdCritical = availability_critical
}

func setAdvancedSettings(testTypeId int, advanced_setting map[string]interface{}, testConfig *TestConfig) {
	var applied_test_flag_ids []int

	test_flags := [28]string{"verify_test_on_failure", "debug_primary_host_on_failure", "enable_http2", "debug_referenced_hosts_on_failure", "capture_http_headers", "capture_response_content", "ignore_ssl_failures", "host_data_collection_enabled", "zone_data_collection_enabled", "f40x_or_50x_http_mark_successful", "t30x_redirects_do_not_follow", "enable_self_versus_third_party_zones", "allow_test_download_limit_override", "capture_filmstrip", "capture_screenshot", "stop_test_on_document_complete", "disable_cross_origin_iframe_access", "stop_test_on_dom_content_load", "enable_path_mtu_discovery", "protocol", "disable_recursive_resolution", "enable_bind_hostname", "enable_tcp_protocol", "enable_nsid", "enable_dnssec", "favor_fastest_round_trip_nameserver", "try_next_nameserver_on_failure", "certificate_revocation_disabled"}
	applied_test_flag_ids = make([]int, len(test_flags))
	for i, test_flag := range test_flags {
		if advanced_setting[test_flag] != nil && advanced_setting[test_flag].(bool) {
			applied_test_flag_ids[i] = getTestFlagId(test_flag)
		}
	}

	if testTypeId == int(TestType(Web)) ||
		testTypeId == int(TestType(Api)) ||
		testTypeId == int(TestType(Transaction)) ||
		testTypeId == int(TestType(Playwright)) ||
		testTypeId == int(TestType(Puppeteer)) {
		enforce_test_failure_if_runs_longer_than := advanced_setting["enforce_test_failure_if_runs_longer_than"].(int)
		viewport_height := advanced_setting["viewport_height"].(int)
		viewport_width := advanced_setting["viewport_width"].(int)
		testConfig.MaxStepRuntimeSecOverride = enforce_test_failure_if_runs_longer_than
		testConfig.ViewportHeight = viewport_height
		testConfig.ViewportWidth = viewport_width
		// if stop_test_on_document_complete is enabled then wait for no activity time is required. Defaults to 0 ms value.
		// Pointer for WaitForNoActivityOnDocComplete will be nil if stop_test_on_document_complete is not enabled by user and waitForNoActivity will be omitted in JSON.
		wait_for_no_activity := advanced_setting["wait_for_no_activity"].(int)
		if advanced_setting["stop_test_on_document_complete"].(bool) {
			testConfig.WaitForNoActivityOnDocComplete = &wait_for_no_activity
		}
		bandwidth_throttling := advanced_setting["bandwidth_throttling"].(string)
		bandwidth_throttling_id, bandwidth_throttling_name := getBandwidthThrottlingTypeId(bandwidth_throttling)
		testConfig.BandwidthThrottling.Id = bandwidth_throttling_id
		testConfig.BandwidthThrottling.Name = bandwidth_throttling_name
	}

	if testTypeId == int(TestType(Dns)) {
		edns_subnet := advanced_setting["edns_subnet"].(string)
		testConfig.EdnsSubnet = edns_subnet
	}

	if testTypeId == int(TestType(Ssl)) ||
		testTypeId == int(TestType(Dns)) ||
		testTypeId == int(TestType(Web)) ||
		testTypeId == int(TestType(Api)) ||
		testTypeId == int(TestType(Transaction)) ||
		testTypeId == int(TestType(Playwright)) ||
		testTypeId == int(TestType(Puppeteer)) {
		additional_monitor := advanced_setting["additional_monitor"].(string)
		if additional_monitor != "" {
			additional_monitor_id, additional_monitor_name := getAdditionalMonitorTypeId(additional_monitor)
			testConfig.AdditionalMonitorType.Id = additional_monitor_id
			testConfig.AdditionalMonitorType.Name = additional_monitor_name
		}
	}

	if testTypeId == int(TestType(Traceroute)) {
		ping_count := advanced_setting["ping_count"].(int)
		failure_hop_count := advanced_setting["failure_hop_count"].(int)
		testConfig.TracerouteFailureHopCount = failure_hop_count
		testConfig.TraceroutePingCount = ping_count
	}

	testConfig.AdvancedSettingType = 1
	testConfig.AppliedTestFlags = applied_test_flag_ids

}

func setRequestSettings(testTypeId int, request_setting map[string]interface{}, testConfig *TestConfig) error {
	authentication_list := request_setting["authentication"].(*schema.Set).List()
	for i := range authentication_list {
		authentication := authentication_list[i].(map[string]interface{})
		authentication_type := authentication["authentication_type"].(string)
		authentication_type_id, authentication_type_name := getAuthenticationTypeId(authentication_type)
		if authentication_type_id == -1 {
			return errors.New("invalid authentication_type provided. valid types are 'basic','ntlm','digest','login'")
		}
		testConfig.AuthenticationType.Id = authentication_type_id
		testConfig.AuthenticationType.Name = authentication_type_name

		tfpassword_ids := authentication["password_ids"].([]interface{})
		password_ids := make([]int, len(tfpassword_ids))
		for i, password_id := range tfpassword_ids {
			password_ids[i] = password_id.(int)
		}
		testConfig.AuthenticationPasswordIds = password_ids
	}

	tftoken_ids := request_setting["token_ids"].([]interface{})
	token_ids := make([]int, len(tftoken_ids))
	for i, token_id := range tftoken_ids {
		token_ids[i] = token_id.(int)
	}
	testConfig.AuthenticationTokenIds = token_ids

	if request_setting["library_certificate_ids"] != nil {
		tfcertificate_ids := request_setting["library_certificate_ids"].([]interface{})
		certificate_ids := make([]int, len(tfcertificate_ids))
		for i, certificate_id := range tfcertificate_ids {
			certificate_ids[i] = certificate_id.(int)
		}
		testConfig.AuthenticationCertificateIds = certificate_ids
	}

	http_request_headers_list := request_setting["http_request_headers"].(*schema.Set).List()
	for i := range http_request_headers_list {
		http_request_headers := http_request_headers_list[i].(map[string]interface{})
		request_headers := [14]string{"user_agent", "accept", "accept_encoding", "accept_language", "accept_charset", "cookie", "cache_control", "pragma", "referer", "host", "request_override", "dns_override", "request_block", "request_delay"}

		for _, request_header := range request_headers {
			request_header_list := http_request_headers[request_header].(*schema.Set).List()
			for i := range request_header_list {
				request_header_map := request_header_list[i].(map[string]interface{})
				request_value := request_header_map["value"].(string)
				child_host_pattern := request_header_map["child_host_pattern"].(string)
				if request_value != "" {
					request_header_id, request_header_name := getReqHeaderTypeId(request_header)
					httpHeaderRequest := TestHttpHeaderRequest{RequestHeaderType: IdName{Id: request_header_id, Name: request_header_name}, RequestValue: request_value, ChildHostPattern: child_host_pattern}
					testConfig.TestHttpHeaderRequests = append(testConfig.TestHttpHeaderRequests, httpHeaderRequest)
				}
			}
		}
	}
	testConfig.RequestSettingType = 1

	return nil
}

func setInsightSettings(testTypeId int, insight_setting map[string]interface{}, testConfig *TestConfig) {

	tftracepoint_ids := insight_setting["tracepoint_ids"].([]interface{})
	tracepoint_ids := make([]int, len(tftracepoint_ids))
	for i, tracepoint_id := range tftracepoint_ids {
		tracepoint_ids[i] = tracepoint_id.(int)
	}
	tfindicator_ids := insight_setting["indicator_ids"].([]interface{})
	indicator_ids := make([]int, len(tfindicator_ids))
	for i, indicator_id := range tfindicator_ids {
		indicator_ids[i] = indicator_id.(int)
	}
	testConfig.InsightSettingType = 1
	testConfig.TracepointIds = tracepoint_ids
	testConfig.IndicatorIds = indicator_ids
}

func setScheduleSettings(testTypeId int, schedule_setting map[string]interface{}, testConfig *TestConfig) error {
	frequency := schedule_setting["frequency"].(string)
	frequency_id, frequency_name := getFrequencyId(frequency)
	if frequency_id == -1 {
		return errors.New("invalid test scheduling frequency string provided. acceptable values are 1 minute, 2 minutes, 4 minutes, 5 minutes, 10 minutes, 15 minutes, 20 minutes, 30 minutes, 60 minutes, 2 hours, 3 hours, 4 hours, 6 hours, 8 hours, 12 hours, 24 hours")
	}
	node_distribution := schedule_setting["node_distribution"].(string)
	node_distribution_id, node_distribution_name := getNodeDistributionId(node_distribution)
	if node_distribution_id == -1 {
		return errors.New("invalid node distribution string provided. acceptable values are random and concurrent")
	}
	no_of_subset_nodes := schedule_setting["no_of_subset_nodes"].(int)
	tfnode_ids := schedule_setting["node_ids"].([]interface{})
	node_ids := make([]int, len(tfnode_ids))
	for i, tfnode := range tfnode_ids {
		node_ids[i] = tfnode.(int)
	}
	tfnode_group_ids := schedule_setting["node_group_ids"].([]interface{})

	var nodeGroupIds []NodeGroup
	networkType := GenericIdName{Id: 0, Name: "Backbone"}
	for _, id := range tfnode_group_ids {
		nodeGroup := NodeGroup{
			Id:                   id.(int),
			Name:                 "DefaultNodeGroupName",
			Description:          "",
			SyntheticNetworkType: networkType,
			Nodes:                []Node{{Id: 123, Name: "DefaultNodeName", NetworkType: networkType}},
		}
		nodeGroupIds = append(nodeGroupIds, nodeGroup)
	}

	if len(tfnode_ids) == 0 && len(tfnode_group_ids) == 0 {
		return errors.New("must specify at least 1 node_ids or node_group_ids in a list. to inherit, remove the schedule_settings attribute")
	}
	run_schedule_id := schedule_setting["run_schedule_id"].(int)
	maintenance_schedule_id := schedule_setting["maintenance_schedule_id"].(int)

	testConfig.ScheduleSettingType = 1
	testConfig.ScheduleRunScheduleId = run_schedule_id
	testConfig.ScheduleMaintenanceScheduleId = maintenance_schedule_id
	testConfig.TestFrequency.Id = frequency_id
	testConfig.TestFrequency.Name = frequency_name
	testConfig.NodeDistribution.Id = node_distribution_id
	testConfig.NodeDistribution.Name = node_distribution_name
	testConfig.NodeIds = node_ids
	testConfig.NodeGroupIds = nodeGroupIds

	if no_of_subset_nodes > 0 {
		testConfig.NoOfSubsetNodes = no_of_subset_nodes
	}

	return nil
}

func setAlertSettings(testTypeId int, alert_setting map[string]interface{}, testConfig *TestConfig) error {

	alert_rule_list := alert_setting["alert_rule"].(*schema.Set).List()
	for i := range alert_rule_list {
		alert_rule := alert_rule_list[i].(map[string]interface{})
		node_threshold_type := alert_rule["node_threshold_type"].(string)
		node_threshold_type_id, node_threshold_type_name := getNodeThresholdTypeId(node_threshold_type)
		if node_threshold_type_id == -1 {
			return errors.New("invalid node threshold type string provided. acceptable values are runs and node")
		}
		threshold_number_of_runs := alert_rule["threshold_number_of_runs"].(int)
		consecutive_number_of_runs := alert_rule["consecutive_number_of_runs"].(int)
		threshold_percentage_of_runs := alert_rule["threshold_percentage_of_runs"].(float64)
		enable_consecutive := alert_rule["enable_consecutive"].(bool)
		warning_reminder := alert_rule["warning_reminder"].(string)
		warning_reminder_id, warning_reminder_name := getReminderId(warning_reminder)
		if warning_reminder_id == 0 {
			log.Printf("[INFO] warning_reminder was not set or an invalid interval string was provided and defaulted to none")
		}
		critical_reminder := alert_rule["critical_reminder"].(string)
		critical_reminder_id, critical_reminder_name := getReminderId(critical_reminder)
		if critical_reminder_id == 0 {
			log.Printf("[INFO] critical_reminder was not set or an invalid interval string was provided and defaulted to none")
		}
		threshold_interval := alert_rule["threshold_interval"].(string)
		threshold_interval_id, threshold_interval_name := getThresholdIntervalId(threshold_interval)
		if threshold_interval_id == 0 {
			log.Printf("[INFO] threshold_interval was not set or an invalid interval string was provided and defaulted to default interval")
		}
		use_rolling_window := alert_rule["use_rolling_window"].(bool)
		notification_type := alert_rule["notification_type"].(string)
		notification_type_id := getNotificationTypeId(notification_type)
		if notification_type_id == 0 {
			log.Printf("[INFO] notification_type was not set or an invalid interval string was provided and defaulted to default contacts")
		}
		alert_type := alert_rule["alert_type"].(string)
		alert_type_id, alert_type_name := getAlertTypeId(alert_type)
		if alert_type_id == -1 {
			return errors.New("invalid alert type string provided")
		}
		enforce_test_failure := alert_rule["enforce_test_failure"].(bool)
		omit_scatterplot := alert_rule["omit_scatterplot"].(bool)

		number_of_failing_nodes := alert_rule["number_of_failing_nodes"].(int)
		trigger_type := alert_rule["trigger_type"].(string)
		trigger_type_id, trigger_type_name := getTriggerTypeId(trigger_type)

		operation_type := alert_rule["operation_type"].(string)
		operation_type_id, operation_type_name := getOperationTypeId(operation_type)
		if operation_type_id == -1 && trigger_type_id != 3 {
			if alert_type_id == 15 || alert_type_id == 7 || alert_type_id == 12 || alert_type_id == 20 {
				return errors.New("operation_type is required. acceptable values are 'greater than', 'greater than or equals', 'less than', 'less than or equals'")
			}
		}

		if alert_type_id == 9 || trigger_type_id == 3 || alert_type_id == 4 {
			operation_type_id = 0
		}

		var historical_interval_id int
		var statistical_type_id int
		var historical_interval_name string
		var statistical_type_name string
		historical_interval := alert_rule["historical_interval"].(string)
		statistical_type := alert_rule["statistical_type"].(string)
		if trigger_type_id == 2 {
			historical_interval_id, historical_interval_name = getHistoricalIntervalId(historical_interval)
			statistical_type_id, statistical_type_name = getStatisticalTypeId(statistical_type)
		}

		if node_threshold_type_id != 1 && threshold_number_of_runs == 0 && threshold_percentage_of_runs == 0 {
			return errors.New("must specify at least 1 node threshold type: threshold_number_of_runs or threshold_percentage_of_runs")
		}

		var expression string
		if alert_rule["expression"] != nil {
			expression = alert_rule["expression"].(string)
		}
		warning_trigger := alert_rule["warning_trigger"].(float64)
		critical_trigger := alert_rule["critical_trigger"].(float64)
		alert_sub_type := alert_rule["alert_sub_type"].(string)
		var alert_sub_type_id int
		var alert_sub_type_name string
		if alert_type_id != 9 && alert_type_id != 4 {
			alert_sub_type_id, alert_sub_type_name = getAlertSubTypeId(alert_sub_type)
			if alert_sub_type_id == -1 {
				return errors.New("must specify the alert sub type. for example 'test' for alert_type 'availability'")
			}
		}
		alert_notif_group_list := alert_rule["notification_group"].(*schema.Set).List()

		var notificationGroups []NotificationGroupStruct
		recipientType := GenericIdName{Id: 2, Name: "Email"}
		contactGroupType := GenericIdName{Id: 1, Name: "ContactGroup"}

		for _, notif_group_item := range alert_notif_group_list {
			var subject string
			var notifyOnCritical bool
			var notifyOnImproved bool
			var notifyOnWarning bool
			var all_email_ids []Recipient
			notification_group := notif_group_item.(map[string]interface{})
			subject = notification_group["subject"].(string)
			notifyOnCritical = notification_group["notify_on_critical"].(bool)
			notifyOnWarning = notification_group["notify_on_warning"].(bool)
			notifyOnImproved = notification_group["notify_on_improved"].(bool)
			var emailIds []interface{}

			if notificationGroup, ok := notification_group["recipient_email_ids"]; ok {
				if email_ids, ok := notificationGroup.([]interface{}); ok {
					emailIds = email_ids
				}
			}
			var contactGroups []interface{}

			if notificationGroup, ok := notification_group["contact_groups"]; ok {
				if groups, ok := notificationGroup.([]interface{}); ok {
					contactGroups = groups
				}
			}

			for _, emailID := range emailIds {
				email, ok := emailID.(string)
				if !ok {
					continue
				}
				all_email_ids = append(all_email_ids, Recipient{Email: email, RecipientType: recipientType})
			}

			for i, contactGroup := range contactGroups {
				contact, ok := contactGroup.(string)
				if !ok {
					continue
				}
				all_email_ids = append(all_email_ids, Recipient{Id: i + 1, RecipientType: contactGroupType, Name: contact})
			}

			notificationGroups = append(notificationGroups, NotificationGroupStruct{Subject: subject,
				NotifyOnWarning:  notifyOnWarning,
				NotifyOnCritical: notifyOnCritical,
				NotifyOnImproved: notifyOnImproved,
				AlertWebhooks:    []AlertWebhook{},
				Recipients:       all_email_ids})
		}

		testConfig.AlertRuleConfigs = append(testConfig.AlertRuleConfigs, AlertRuleConfig{AlertNodeThresholdType: IdName{Id: node_threshold_type_id, Name: node_threshold_type_name}, AlertThresholdNumOfRuns: threshold_number_of_runs, AlertConsecutiveNumOfRuns: consecutive_number_of_runs, AlertThresholdPercentOfRuns: threshold_percentage_of_runs,
			AlertThresholdNumOfFailingNodes: number_of_failing_nodes, TriggerType: IdName{Id: trigger_type_id, Name: trigger_type_name},
			OperationType:              IdName{Id: operation_type_id, Name: operation_type_name},
			StatisticalType:            IdName{Id: statistical_type_id, Name: statistical_type_name},
			TrailingHistoricalInterval: IdName{Id: historical_interval_id, Name: historical_interval_name},
			AlertWarningTrigger:        warning_trigger, AlertCriticalTrigger: critical_trigger, Expression: expression,
			AlertEnableConsecutive: enable_consecutive, AlertWarningReminder: IdName{Id: warning_reminder_id, Name: warning_reminder_name},
			AlertCriticalReminder:  IdName{Id: critical_reminder_id, Name: critical_reminder_name},
			AlertThresholdInterval: IdName{Id: threshold_interval_id, Name: threshold_interval_name},
			AlertUseRollingWindow:  use_rolling_window, AlertNotificationType: notification_type_id,
			AlertType:               IdName{Id: alert_type_id, Name: alert_type_name},
			AlertSubType:            IdName{Id: alert_sub_type_id, Name: alert_sub_type_name},
			AlertEnforceTestFailure: enforce_test_failure,
			AlertOmitScatterplot:    omit_scatterplot,
			NotificationGroups:      notificationGroups,
		})
	}

	notif_group_list := alert_setting["notification_group"].(*schema.Set).List()

	var all_alert_webhook_ids []int
	var all_email_ids []string
	var all_contact_groups []string
	var subject string

	for _, notif_group_item := range notif_group_list {
		notification_group := notif_group_item.(map[string]interface{})

		tfalert_webhooks := notification_group["alert_webhook_ids"].([]interface{})
		for _, tfalert_webhook := range tfalert_webhooks {
			all_alert_webhook_ids = append(all_alert_webhook_ids, tfalert_webhook.(int))
		}

		tfemail_ids := notification_group["recipient_email_ids"].([]interface{})
		for _, email_id := range tfemail_ids {
			all_email_ids = append(all_email_ids, email_id.(string))
		}

		ContactGroups := notification_group["contact_groups"].([]interface{})
		for _, contactGroup := range ContactGroups {
			all_contact_groups = append(all_contact_groups, contactGroup.(string))
		}

		subject = subject + notification_group["subject"].(string)
	}

	testConfig.AlertSettingType = 1
	testConfig.AlertWebhookIds = all_alert_webhook_ids
	testConfig.AlertRecipientEmails = all_email_ids
	testConfig.AlertContactGroups = all_contact_groups
	testConfig.AlertSubject = subject

	return nil
}
