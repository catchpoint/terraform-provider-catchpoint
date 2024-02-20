package catchpoint

import (
	"math/rand"
	"time"
)

type Config struct {
	ApiToken    string
	LogJson     bool
	Environment string
}

func newConfig(apiToken string, logJson bool, cpEnvironment string) *Config {
	return &Config{
		ApiToken:    apiToken,
		LogJson:     logJson,
		Environment: cpEnvironment,
	}
}

func getTime() string {
	t := time.Now()
	timeCurrent := t.Format(time.RFC3339)
	return string(timeCurrent)
}

func randHexString() string {
	const hexLetters = "abcdef0123456789"
	const numOfLetters = 6
	b := make([]byte, numOfLetters)
	for i := range b {
		b[i] = hexLetters[rand.Intn(len(hexLetters))]
	}
	return "#" + string(b)
}

func getTestStatusTypeId(testStatus string) int {
	testStatusTypes := map[int]string{
		0: "active",
		1: "inactive",
	}
	for id, testStatusType := range testStatusTypes {
		if testStatusType == testStatus {
			return id
		}
	}
	return 0
}

func getMonitorId(monitor string) int {
	monitorTypes := map[int]string{
		2:  "object",
		3:  "emulated",
		18: "chrome",
		19: "playback",
		20: "mobile playback",
		26: "mobile",
		25: "api",
		8:  "ping icmp",
		11: "ping tcp",
		12: "dns experience",
		13: "dns direct",
		23: "ping udp",
		9:  "traceroute icmp",
		14: "traceroute udp",
		29: "traceroute tcp",
		31: "ssl",
		34: "bgp",
		41: "bgp basic",
	}
	for id, monitorType := range monitorTypes {
		if monitorType == monitor {
			return id
		}
	}
	return -1
}

func getMonitorName(monitor int) string {
	monitorTypes := map[int]string{
		2:  "object",
		3:  "emulated",
		18: "chrome",
		19: "playback",
		20: "mobile playback",
		26: "mobile",
		25: "api",
		8:  "ping icmp",
		11: "ping tcp",
		12: "dns experience",
		13: "dns direct",
		23: "ping udp",
		9:  "traceroute icmp",
		14: "traceroute udp",
		29: "traceroute tcp",
		31: "ssl",
		34: "bgp",
		41: "bgp basic",
	}
	for id, monitorType := range monitorTypes {
		if id == monitor {
			return monitorType
		}
	}
	return ""
}

func getApiScriptTypeId(scriptType string) int {
	apiScriptTypes := map[int]string{
		1: "selenium",
		2: "javascript",
	}
	for id, apiScriptType := range apiScriptTypes {
		if apiScriptType == scriptType {
			return id
		}
	}
	return -1
}

func getUserAgentTypeId(userAgentType string) int {
	userAgentTypes := map[int]string{
		3:  "android",
		4:  "iphone",
		5:  "ipad 2",
		6:  "kindle fire",
		7:  "galaxy tab",
		8:  "iphone 5",
		9:  "ipad mini",
		10: "galaxy note",
		11: "nexus 7",
		12: "nexus 4",
		13: "nokia lumia920",
		14: "iphone 6",
		15: "blackberry z30",
		16: "galaxy s4",
		17: "htc onex",
		18: "lg optimusg",
		19: "droid razr hd",
		20: "nexus 6",
		21: "iphone 6s",
		22: "galaxy s6",
		23: "iphone 7",
		24: "google pixel",
		25: "galaxy s8",
	}
	for id, userAgentTypeString := range userAgentTypes {
		if userAgentTypeString == userAgentType {
			return id
		}
	}
	return 0
}

func getUserAgentTypeName(userAgentType int) string {
	userAgentTypes := map[int]string{
		3:  "android",
		4:  "iphone",
		5:  "ipad 2",
		6:  "kindle fire",
		7:  "galaxy tab",
		8:  "iphone 5",
		9:  "ipad mini",
		10: "galaxy note",
		11: "nexus 7",
		12: "nexus 4",
		13: "nokia lumia920",
		14: "iphone 6",
		15: "blackberry z30",
		16: "galaxy s4",
		17: "htc onex",
		18: "lg optimusg",
		19: "droid razr hd",
		20: "nexus 6",
		21: "iphone 6s",
		22: "galaxy s6",
		23: "iphone 7",
		24: "google pixel",
		25: "galaxy s8",
	}
	for id, userAgentTypeString := range userAgentTypes {
		if id == userAgentType {
			return userAgentTypeString
		}
	}
	return ""
}

func getChromeVersionId(chromeVersion string) (int, string) {
	chromeVersions := map[int][]string{
		1: {"stable"},
		2: {"preview"},
		3: {"108", "89", "87", "85", "75", "71", "66", "63", "59", "53"},
	}
	for id, chromeVer := range chromeVersions {
		for _, specifcVersion := range chromeVer {
			if specifcVersion == chromeVersion {
				return id, specifcVersion
			}
		}
	}
	return 0, ""
}

func getChromeApplicationVersionId(chromeApplicationVersion string) (int, string) {
	chromeApplicationVersions := map[int]string{
		1:     "53",
		3:     "59",
		4:     "63",
		5:     "66",
		7:     "75",
		8:     "71",
		12:    "85",
		13:    "87",
		14:    "89",
		28558: "108",
	}
	for id, chromeApplicationVer := range chromeApplicationVersions {

		if chromeApplicationVer == chromeApplicationVersion {
			return id, chromeApplicationVer
		}
	}
	return 0, ""
}

func getDnsQueryTypeId(queryType string) (int, string) {
	queryTypes := map[int]string{
		0:     "none",
		1:     "a",
		2:     "ns",
		5:     "cname",
		6:     "soa",
		7:     "mb",
		8:     "mg",
		9:     "mr",
		10:    "null",
		11:    "wks",
		12:    "ptr",
		13:    "hinfo",
		14:    "minfo",
		15:    "mx",
		16:    "txt",
		17:    "rp",
		18:    "afsdb",
		19:    "x25",
		20:    "isdn",
		21:    "rt",
		22:    "nsap",
		24:    "sig",
		25:    "key",
		26:    "px",
		28:    "aaaa",
		29:    "loc",
		31:    "eid",
		32:    "nimloc",
		33:    "srv",
		34:    "atma",
		35:    "naptr",
		36:    "kx",
		37:    "cert",
		38:    "a6",
		39:    "dname",
		40:    "sink",
		41:    "opt",
		42:    "apl",
		43:    "ds",
		44:    "sshfp",
		45:    "ipseckey",
		46:    "rrsig",
		47:    "nsec",
		48:    "dnskey",
		49:    "dhcid",
		50:    "nsec3",
		51:    "nsec3param",
		55:    "hip",
		99:    "spf",
		100:   "uinfo",
		101:   "uid",
		102:   "gid",
		103:   "unspec",
		249:   "tkey",
		250:   "tsig",
		251:   "ixfr",
		252:   "axfr",
		253:   "mailb",
		255:   "any",
		32768: "ta",
		32769: "dlv",
	}
	for id, queryTypeString := range queryTypes {
		if queryTypeString == queryType {
			return id, queryTypeString
		}
	}
	return 0, ""
}

func getDnsQueryName(queryType int) string {
	queryTypes := map[int]string{
		0:     "none",
		1:     "a",
		2:     "ns",
		5:     "cname",
		6:     "soa",
		7:     "mb",
		8:     "mg",
		9:     "mr",
		10:    "null",
		11:    "wks",
		12:    "ptr",
		13:    "hinfo",
		14:    "minfo",
		15:    "mx",
		16:    "txt",
		17:    "rp",
		18:    "afsdb",
		19:    "x25",
		20:    "isdn",
		21:    "rt",
		22:    "nsap",
		24:    "sig",
		25:    "key",
		26:    "px",
		28:    "aaaa",
		29:    "loc",
		31:    "eid",
		32:    "nimloc",
		33:    "srv",
		34:    "atma",
		35:    "naptr",
		36:    "kx",
		37:    "cert",
		38:    "a6",
		39:    "dname",
		40:    "sink",
		41:    "opt",
		42:    "apl",
		43:    "ds",
		44:    "sshfp",
		45:    "ipseckey",
		46:    "rrsig",
		47:    "nsec",
		48:    "dnskey",
		49:    "dhcid",
		50:    "nsec3",
		51:    "nsec3param",
		55:    "hip",
		99:    "spf",
		100:   "uinfo",
		101:   "uid",
		102:   "gid",
		103:   "unspec",
		249:   "tkey",
		250:   "tsig",
		251:   "ixfr",
		252:   "axfr",
		253:   "mailb",
		255:   "any",
		32768: "ta",
		32769: "dlv",
	}
	for id, queryTypeString := range queryTypes {
		if id == queryType {
			return queryTypeString
		}
	}
	return ""
}

func getFrequencyId(frequency string) (int, string) {
	frequencies := map[int]string{
		0:  "none",
		1:  "1 minute",
		2:  "5 minutes",
		3:  "10 minutes",
		4:  "15 minutes",
		5:  "20 minutes",
		6:  "30 minutes",
		7:  "60 minutes",
		8:  "2 hours",
		9:  "3 hours",
		10: "4 hours",
		11: "6 hours",
		12: "8 hours",
		13: "12 hours",
		14: "24 hours",
		15: "4 minutes",
		16: "2 minutes",
	}
	for id, freq := range frequencies {
		if freq == frequency {
			return id, freq
		}
	}
	return -1, ""
}

func getFrequencyName(frequency int) string {
	frequencies := map[int]string{
		0:  "none",
		1:  "1 minute",
		2:  "5 minutes",
		3:  "10 minutes",
		4:  "15 minutes",
		5:  "20 minutes",
		6:  "30 minutes",
		7:  "60 minutes",
		8:  "2 hours",
		9:  "3 hours",
		10: "4 hours",
		11: "6 hours",
		12: "8 hours",
		13: "12 hours",
		14: "24 hours",
		15: "4 minutes",
		16: "2 minutes",
	}
	for id, freq := range frequencies {
		if id == frequency {
			return freq
		}
	}
	return ""
}

func getNodeDistributionId(nodeDistribution string) (int, string) {
	nodeDistributions := map[int]string{
		0: "random",
		1: "concurrent",
	}
	for id, nodeDist := range nodeDistributions {
		if nodeDist == nodeDistribution {
			return id, nodeDist
		}
	}
	return -1, ""
}

func getNodeDistributionName(nodeDistribution int) string {
	nodeDistributions := map[int]string{
		0: "random",
		1: "concurrent",
	}
	for id, nodeDist := range nodeDistributions {
		if id == nodeDistribution {
			return nodeDist
		}
	}
	return ""
}

func getNodeThresholdTypeId(nodeThresholdType string) (int, string) {
	nodeThresholdTypes := map[int]string{
		0: "runs",
		1: "average across nodes",
		2: "node",
	}
	for id, nodeThreshold := range nodeThresholdTypes {
		if nodeThreshold == nodeThresholdType {
			return id, nodeThreshold
		}
	}
	return -1, ""
}

func getNodeThresholdTypeName(nodeThresholdType int) string {
	nodeThresholdTypes := map[int]string{
		0: "runs",
		1: "average across nodes",
		2: "node",
	}
	for id, nodeThreshold := range nodeThresholdTypes {
		if id == nodeThresholdType {
			return nodeThreshold
		}
	}
	return ""
}

func getOperationTypeId(operationType string) (int, string) {
	operationTypes := map[int]string{
		0: "not equals",
		1: "equals",
		2: "greater than",
		3: "greater than or equals",
		4: "less than",
		5: "less than or equals",
	}
	for id, opType := range operationTypes {
		if opType == operationType {
			return id, opType
		}
	}
	return -1, ""
}

func getOperationTypeName(operationType int) string {
	operationTypes := map[int]string{
		0: "not equals",
		1: "equals",
		2: "greater than",
		3: "greater than or equals",
		4: "less than",
		5: "less than or equals",
	}
	for id, opType := range operationTypes {
		if id == operationType {
			return opType
		}
	}
	return ""
}

func getTriggerTypeId(triggerType string) (int, string) {
	TriggerTypes := map[int]string{
		1: "specific value",
		2: "trailing value",
		3: "trendshift",
	}
	for id, triggType := range TriggerTypes {
		if triggType == triggerType {
			return id, triggType
		}
	}
	return 1, "specific value"
}

func getTriggerTypeName(triggerType int) string {
	TriggerTypes := map[int]string{
		1: "specific value",
		2: "trailing value",
		3: "trendshift",
	}
	for id, triggType := range TriggerTypes {
		if id == triggerType {
			return triggType
		}
	}
	//Default to specific value alert trigger type
	return "specific value"
}

func getReminderId(reminder string) (int, string) {
	reminders := map[int]string{
		0:    "none",
		1:    "1 minute",
		5:    "5 minutes",
		10:   "10 minutes",
		15:   "15 minutes",
		30:   "30 minutes",
		60:   "1 hour",
		1440: "daily",
	}
	for id, reminderInterval := range reminders {
		if reminderInterval == reminder {
			return id, reminderInterval
		}
	}
	return 0, "none"
}

func getReminderName(reminder int) string {
	reminders := map[int]string{
		0:    "none",
		1:    "1 minute",
		5:    "5 minutes",
		10:   "10 minutes",
		15:   "15 minutes",
		30:   "30 minutes",
		60:   "1 hour",
		1440: "daily",
	}
	for id, reminderInterval := range reminders {
		if id == reminder {
			return reminderInterval
		}
	}
	return "none"
}

func getThresholdIntervalId(thresholdInterval string) (int, string) {
	thresholdIntervals := map[int]string{
		0:   "default",
		5:   "5 minutes",
		10:  "10 minutes",
		15:  "15 minutes",
		30:  "30 minutes",
		60:  "1 hour",
		120: "2 hours",
		360: "6 hours",
		720: "12 hours",
	}
	for id, alertThresholdInterval := range thresholdIntervals {
		if alertThresholdInterval == thresholdInterval {
			return id, alertThresholdInterval
		}
	}
	return 0, "default"
}

func getThresholdIntervalName(thresholdInterval int) string {
	thresholdIntervals := map[int]string{
		0:   "default",
		5:   "5 minutes",
		10:  "10 minutes",
		15:  "15 minutes",
		30:  "30 minutes",
		60:  "1 hour",
		120: "2 hours",
		360: "6 hours",
		720: "12 hours",
	}
	for id, alertThresholdInterval := range thresholdIntervals {
		if id == thresholdInterval {
			return alertThresholdInterval
		}
	}
	return "default"
}

func getHistoricalIntervalId(historicalInterval string) (int, string) {
	historicalIntervals := map[int]string{
		5:     "5 minutes",
		10:    "10 minutes",
		15:    "15 minutes",
		30:    "30 minutes",
		60:    "1 hour",
		120:   "2 hours",
		360:   "6 hours",
		720:   "12 hours",
		1440:  "1 day",
		10080: "1 week",
	}
	for id, alertHistoricalInterval := range historicalIntervals {
		if alertHistoricalInterval == historicalInterval {
			return id, alertHistoricalInterval
		}
	}
	return 5, "5 minutes"
}

func getHistoricalIntervalName(historicalInterval int) string {
	historicalIntervals := map[int]string{
		5:     "5 minutes",
		10:    "10 minutes",
		15:    "15 minutes",
		30:    "30 minutes",
		60:    "1 hour",
		120:   "2 hours",
		360:   "6 hours",
		720:   "12 hours",
		1440:  "1 day",
		10080: "1 week",
	}
	for id, alertHistoricalInterval := range historicalIntervals {
		if id == historicalInterval {
			return alertHistoricalInterval
		}
	}
	return "5 minutes"
}

func getNotificationTypeId(notificationType string) int {
	notificationTypes := map[int]string{
		0: "default contacts",
	}
	for id, alertNotificationType := range notificationTypes {
		if alertNotificationType == notificationType {
			return id
		}
	}
	return 0
}

func getAlertTypeId(alertType string) (int, string) {
	alertTypes := map[int]string{
		2:  "byte length",
		3:  "content match",
		4:  "host failure",
		9:  "test failure",
		7:  "timing",
		12: "ping",
		13: "requests",
		15: "availability",
		20: "path",
		23: "asn",
	}
	for id, alertTypeString := range alertTypes {
		if alertTypeString == alertType {
			return id, alertTypeString
		}
	}
	return -1, ""
}

func getAlertTypeName(alertType int) string {
	alertTypes := map[int]string{
		2:  "byte length",
		3:  "content match",
		4:  "host failure",
		9:  "test failure",
		7:  "timing",
		12: "ping",
		13: "requests",
		15: "availability",
		20: "path",
		23: "asn",
	}
	for id, alertTypeString := range alertTypes {
		if id == alertType {
			return alertTypeString
		}
	}
	return ""
}

func getAlertSubTypeId(alertSubType string) (int, string) {
	alertSubTypes := map[int]string{
		1:   "byte length",
		2:   "page",
		3:   "file size",
		10:  "regular expression",
		14:  "response code",
		15:  "response headers",
		50:  "dns",
		51:  "connect",
		52:  "send",
		53:  "wait",
		54:  "load",
		55:  "ttfb",
		57:  "content load",
		58:  "response",
		59:  "test time",
		61:  "dom load",
		63:  "test time with suspect",
		64:  "server response",
		66:  "document complete",
		67:  "redirect",
		100: "ping rtt",
		101: "ping packet loss",
		110: "# requests",
		111: "# hosts",
		112: "# connections",
		113: "# redirects",
		114: "# other",
		115: "# images",
		116: "# scripts",
		117: "# html",
		118: "# css",
		119: "# xml",
		120: "# flash",
		121: "# media",
		140: "test",
		141: "content",
		142: "% downtime",
		190: "# cities",
		191: "# asns",
		193: "# countries",
		194: "# hops",
		195: "handshake_time",
		196: "days_to_expiration",
		210: "origin as",
		211: "path as",
		212: "origin neighbor",
		213: "prefix mismatch",
	}
	for id, alertSubTypeString := range alertSubTypes {
		if alertSubTypeString == alertSubType {
			return id, alertSubTypeString
		}
	}
	return -1, ""
}

func getAlertSubTypeName(alertSubType int) string {
	alertSubTypes := map[int]string{
		1:   "byte length",
		2:   "page",
		3:   "file size",
		10:  "regular expression",
		14:  "response code",
		15:  "response headers",
		50:  "dns",
		51:  "connect",
		52:  "send",
		53:  "wait",
		54:  "load",
		55:  "ttfb",
		57:  "content load",
		58:  "response",
		59:  "test time",
		61:  "dom load",
		63:  "test time with suspect",
		64:  "server response",
		66:  "document complete",
		67:  "redirect",
		100: "ping rtt",
		101: "ping packet loss",
		110: "# requests",
		111: "# hosts",
		112: "# connections",
		113: "# redirects",
		114: "# other",
		115: "# images",
		116: "# scripts",
		117: "# html",
		118: "# css",
		119: "# xml",
		120: "# flash",
		121: "# media",
		140: "test",
		141: "content",
		142: "% downtime",
		190: "# cities",
		191: "# asns",
		193: "# countries",
		194: "# hops",
		195: "handshake_time",
		196: "days_to_expiration",
		210: "origin as",
		211: "path as",
		212: "origin neighbor",
		213: "prefix mismatch",
	}
	for id, alertSubTypeString := range alertSubTypes {
		if id == alertSubType {
			return alertSubTypeString
		}
	}
	return ""
}

func getStatisticalTypeId(statisticalType string) (int, string) {
	statisticalTypes := map[int]string{
		1: "average",
	}
	for id, statisticalTypeString := range statisticalTypes {
		if statisticalTypeString == statisticalType {
			return id, statisticalTypeString
		}
	}
	return 1, "average"
}

func getTestFlagId(testFlag string) int {
	testFlagTypes := map[int]string{
		2:  "verify_test_on_failure",
		3:  "debug_primary_host_on_failure",
		4:  "enable_http2",
		8:  "debug_referenced_hosts_on_failure",
		9:  "capture_http_headers",
		11: "capture_response_content",
		13: "capture_filmstrip",
		14: "capture_screenshot",
		17: "ignore_ssl_failures",
		19: "enable_bind_hostname",
		20: "enable_tcp_protocol",
		21: "enable_nsid",
		22: "disable_recursive_resolution",
		23: "host_data_collection_enabled",
		24: "zone_data_collection_enabled",
		25: "stop_test_on_document_complete",
		26: "try_next_nameserver_on_failure",
		27: "f40x_or_50x_http_mark_successful",
		31: "favor_fastest_round_trip_nameserver",
		33: "t30x_redirects_do_not_follow",
		36: "enable_self_versus_third_party_zones",
		37: "allow_test_download_limit_override",
		38: "disable_cross_origin_iframe_access",
		39: "stop_test_on_dom_content_load",
		40: "initiated_from_api",
		41: "instant_test_charged",
		42: "certificate_revocation_disabled",
		43: "enforce_certificate_pinning",
		44: "enforce_public_key_pinning",
		48: "enable_dnssec",
		50: "enable_path_mtu_discovery",
		51: "enable_tracing",
		52: "is_continuous",
		53: "enable_ecn",
		54: "enable_dscp",
		55: "is_accurate_ecn",
		56: "enable_dns_query_limits",
	}
	for id, testFlagString := range testFlagTypes {
		if testFlagString == testFlag {
			return id
		}
	}
	return 0

}

func getTestFlagName(testFlag int) string {
	testFlagTypes := map[int]string{
		2:  "verify_test_on_failure",
		3:  "debug_primary_host_on_failure",
		4:  "enable_http2",
		8:  "debug_referenced_hosts_on_failure",
		9:  "capture_http_headers",
		11: "capture_response_content",
		13: "capture_filmstrip",
		14: "capture_screenshot",
		17: "ignore_ssl_failures",
		19: "enable_bind_hostname",
		20: "enable_tcp_protocol",
		21: "enable_nsid",
		22: "disable_recursive_resolution",
		23: "host_data_collection_enabled",
		24: "zone_data_collection_enabled",
		25: "stop_test_on_document_complete",
		26: "try_next_nameserver_on_failure",
		27: "f40x_or_50x_http_mark_successful",
		31: "favor_fastest_round_trip_nameserver",
		33: "t30x_redirects_do_not_follow",
		36: "enable_self_versus_third_party_zones",
		37: "allow_test_download_limit_override",
		38: "disable_cross_origin_iframe_access",
		39: "stop_test_on_dom_content_load",
		40: "initiated_from_api",
		41: "instant_test_charged",
		42: "certificate_revocation_disabled",
		43: "enforce_certificate_pinning",
		44: "enforce_public_key_pinning",
		48: "enable_dnssec",
		50: "enable_path_mtu_discovery",
		51: "enable_tracing",
		52: "is_continuous",
		53: "enable_ecn",
		54: "enable_dscp",
		55: "is_accurate_ecn",
		56: "enable_dns_query_limits",
	}
	for id, testFlagString := range testFlagTypes {
		if id == testFlag {
			return testFlagString
		}
	}
	return ""
}

func getAdditionalMonitorTypeId(additionalMonitorType string) (int, string) {
	additionalMonitorTypes := map[int]string{
		8:  "ping icmp",
		11: "ping tcp",
		23: "ping udp",
		9:  "traceroute icmp",
		14: "traceroute udp",
		29: "traceroute tcp",
	}
	for id, additionalMonitorTypeString := range additionalMonitorTypes {
		if additionalMonitorTypeString == additionalMonitorType {
			return id, additionalMonitorTypeString
		}
	}
	return -1, ""
}

func getAdditionalMonitorTypeName(additionalMonitorType int) string {
	additionalMonitorTypes := map[int]string{
		8:  "ping icmp",
		11: "ping tcp",
		23: "ping udp",
		9:  "traceroute icmp",
		14: "traceroute udp",
		29: "traceroute tcp",
	}
	for id, additionalMonitorTypeString := range additionalMonitorTypes {
		if id == additionalMonitorType {
			return additionalMonitorTypeString
		}
	}
	return ""
}

func getBandwidthThrottlingTypeId(bandwidthThrottlingType string) (int, string) {
	bandwidthThrottlingTypes := map[int]string{
		1: "gprs",
		2: "regular 2g",
		3: "good 2g",
		4: "regular 3g",
		5: "good 3g",
		6: "regular 4g",
		7: "dsl",
		8: "wifi",
	}
	for id, bandwidthThrottlingTypeString := range bandwidthThrottlingTypes {
		if bandwidthThrottlingTypeString == bandwidthThrottlingType {
			return id, bandwidthThrottlingTypeString
		}
	}
	return 0, ""
}

func getBandwidthThrottlingTypeName(bandwidthThrottlingType int) string {
	bandwidthThrottlingTypes := map[int]string{
		1: "gprs",
		2: "regular 2g",
		3: "good 2g",
		4: "regular 3g",
		5: "good 3g",
		6: "regular 4g",
		7: "dsl",
		8: "wifi",
	}
	for id, bandwidthThrottlingTypeString := range bandwidthThrottlingTypes {
		if id == bandwidthThrottlingType {
			return bandwidthThrottlingTypeString
		}
	}
	return ""
}

func getReqHeaderTypeId(requestHeader string) (int, string) {
	requestHeaderTypes := map[int]string{
		1:  "user_agent",
		2:  "accept",
		3:  "accept_encoding",
		4:  "accept_language",
		5:  "accept_charset",
		6:  "cookie",
		7:  "cache_control",
		9:  "pragma",
		10: "referer",
		12: "host",
		13: "request_override",
		14: "dns_override",
		15: "request_block",
		16: "request_delay",
	}
	for id, requestHeaderType := range requestHeaderTypes {
		if requestHeaderType == requestHeader {
			return id, requestHeaderType
		}
	}
	return 0, ""
}

func getReqHeaderTypeName(requestHeader int) string {
	requestHeaderTypes := map[int]string{
		1:  "user_agent",
		2:  "accept",
		3:  "accept_encoding",
		4:  "accept_language",
		5:  "accept_charset",
		6:  "cookie",
		7:  "cache_control",
		9:  "pragma",
		10: "referer",
		12: "host",
		13: "request_override",
		14: "dns_override",
		15: "request_block",
		16: "request_delay",
	}
	for id, requestHeaderType := range requestHeaderTypes {
		if id == requestHeader {
			return requestHeaderType
		}
	}
	return ""
}

func getAuthenticationTypeId(authenticationType string) (int, string) {
	authenticationTypes := map[int]string{
		1: "basic",
		2: "digest",
		3: "ntlm",
		5: "login",
	}
	for id, authenticationTypeString := range authenticationTypes {
		if authenticationTypeString == authenticationType {
			return id, authenticationTypeString
		}
	}
	return -1, ""
}
