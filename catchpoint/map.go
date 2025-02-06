package catchpoint

func getTestStatusTypeIds() map[string]int {
	testStatusTypes := map[string]int{
		"active":   0,
		"inactive": 1,
	}
	return testStatusTypes
}

func getMonitorTypeIds() map[string]int {
	monitorTypes := map[string]int{
		"object":          2,
		"emulated":        3,
		"chrome":          18,
		"playback":        19,
		"mobile playback": 20,
		"mobile":          26,
		"api":             25,
		"ping icmp":       8,
		"ping tcp":        11,
		"dns experience":  12,
		"dns direct":      13,
		"ping udp":        23,
		"traceroute icmp": 9,
		"traceroute udp":  14,
		"traceroute tcp":  29,
		"ssl":             31,
		"bgp":             34,
		"playwright":      39,
		"bgp basic":       41,
	}
	return monitorTypes
}

func getMonitorTypeNames() map[int]string {
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
		39: "playwright",
		41: "bgp basic",
	}
	return monitorTypes
}

func getApiScriptTypeIds() map[string]int {
	apiScriptTypes := map[string]int{
		"selenium":   1,
		"javascript": 2,
		"playwright": 3,
		"puppeteer":  4,
	}
	return apiScriptTypes
}

func getUserAgentTypeIds() map[string]int {
	userAgentTypes := map[string]int{
		"ie":             1,
		"chrome":         2,
		"android":        3,
		"iphone":         4,
		"ipad 2":         5,
		"kindle fire":    6,
		"galaxy tab":     7,
		"iphone 5":       8,
		"ipad mini":      9,
		"galaxy note":    10,
		"nexus 7":        11,
		"nexus 4":        12,
		"nokia lumia920": 13,
		"iphone 6":       14,
		"blackberry z30": 15,
		"galaxy s4":      16,
		"htc onex":       17,
		"lg optimusg":    18,
		"droid razr hd":  19,
		"nexus 6":        20,
		"iphone 6s":      21,
		"galaxy s6":      22,
		"iphone 7":       23,
		"google pixel":   24,
		"galaxy s8":      25,
	}
	return userAgentTypes

}

func getUserAgentTypeNames() map[int]string {
	userAgentTypes := map[int]string{
		1:  "ie",
		2:  "chrome",
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
	return userAgentTypes
}
func getChromeVersionIds() map[int][]string {
	chromeVersions := map[int][]string{
		1: {"stable"},
		2: {"preview"},
		3: {"120", "108", "89", "87", "85", "79", "75", "71", "66", "63", "59", "53"},
	}
	return chromeVersions
}

func getChromeApplicationVersionIds() map[string]int {
	if catchpoint_environment == "prod" || catchpoint_environment == "" {
		return map[string]int{
			"53":  1,
			"59":  3,
			"63":  4,
			"66":  5,
			"75":  7,
			"71":  8,
			"85":  12,
			"87":  13,
			"89":  14,
			"108": 28558,
			"120": 31965,
		}
	}
	return map[string]int{
		"53":  1,
		"59":  3,
		"63":  4,
		"66":  5,
		"71":  6,
		"75":  7,
		"79":  9,
		"85":  10,
		"87":  13,
		"89":  12,
		"108": 15,
		"120": 44,
	}

}

func getDnsQueryTypeIds() map[string]int {
	queryTypes := map[string]int{
		"none":       0,
		"a":          1,
		"ns":         2,
		"cname":      5,
		"soa":        6,
		"mb":         7,
		"mg":         8,
		"mr":         9,
		"null":       10,
		"wks":        11,
		"ptr":        12,
		"hinfo":      13,
		"minfo":      14,
		"mx":         15,
		"txt":        16,
		"rp":         17,
		"afsdb":      18,
		"x25":        19,
		"isdn":       20,
		"rt":         21,
		"nsap":       22,
		"sig":        24,
		"key":        25,
		"px":         26,
		"aaaa":       28,
		"loc":        29,
		"eid":        31,
		"nimloc":     32,
		"srv":        33,
		"atma":       34,
		"naptr":      35,
		"kx":         36,
		"cert":       37,
		"a6":         38,
		"dname":      39,
		"sink":       40,
		"opt":        41,
		"apl":        42,
		"ds":         43,
		"sshfp":      44,
		"ipseckey":   45,
		"rrsig":      46,
		"nsec":       47,
		"dnskey":     48,
		"dhcid":      49,
		"nsec3":      50,
		"nsec3param": 51,
		"hip":        55,
		"spf":        99,
		"uinfo":      100,
		"uid":        101,
		"gid":        102,
		"unspec":     103,
		"tkey":       249,
		"tsig":       250,
		"ixfr":       251,
		"axfr":       252,
		"mailb":      253,
		"any":        255,
		"ta":         32768,
		"dlv":        32769,
		"AorAAAA":    32770,
	}
	return queryTypes
}

func getDnsQueryTypeNames() map[int]string {
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
		32770: "AorAAAA",
	}
	return queryTypes
}

func getFrequencyIds() map[string]int {
	frequencies := map[string]int{
		"none":       0,
		"1 minute":   1,
		"5 minutes":  2,
		"10 minutes": 3,
		"15 minutes": 4,
		"20 minutes": 5,
		"30 minutes": 6,
		"60 minutes": 7,
		"2 hours":    8,
		"3 hours":    9,
		"4 hours":    10,
		"6 hours":    11,
		"8 hours":    12,
		"12 hours":   13,
		"24 hours":   14,
		"4 minutes":  15,
		"2 minutes":  16,
	}
	return frequencies
}

func getFrequencyNames() map[int]string {
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
	return frequencies
}

func getNodeDistributionIds() map[string]int {
	nodeDistributions := map[string]int{
		"random":     0,
		"concurrent": 1,
	}
	return nodeDistributions
}

func getNodeDistributionNames() map[int]string {
	nodeDistributions := map[int]string{
		0: "random",
		1: "concurrent",
	}
	return nodeDistributions
}

func getNodeThresholdTypeIds() map[string]int {
	nodeThresholdTypes := map[string]int{
		"runs":                 0,
		"average across nodes": 1,
		"node":                 2,
	}
	return nodeThresholdTypes
}

func getNodeThresholdTypeNames() map[int]string {
	nodeThresholdTypes := map[int]string{
		0: "runs",
		1: "average across nodes",
		2: "node",
	}
	return nodeThresholdTypes
}

func getOperationTypeIds() map[string]int {

	operationTypes := map[string]int{
		"not equals":             0,
		"equals":                 1,
		"greater than":           2,
		"greater than or equals": 3,
		"less than":              4,
		"less than or equals":    5,
	}
	return operationTypes
}

func getOperationTypeNames() map[int]string {
	operationTypes := map[int]string{
		0: "not equals",
		1: "equals",
		2: "greater than",
		3: "greater than or equals",
		4: "less than",
		5: "less than or equals",
	}
	return operationTypes
}

func getTriggerTypeIds() map[string]int {
	triggerTypes := map[string]int{
		"specific value": 1,
		"trailing value": 2,
		"trendshift":     3,
	}
	return triggerTypes
}

func getTriggerTypeNames() map[int]string {
	triggerTypes := map[int]string{
		1: "specific value",
		2: "trailing value",
		3: "trendshift",
	}
	return triggerTypes
}

func getReminderIds() map[string]int {
	reminders := map[string]int{
		"none":       0,
		"1 minute":   1,
		"5 minutes":  5,
		"10 minutes": 10,
		"15 minutes": 15,
		"30 minutes": 30,
		"1 hour":     60,
		"daily":      1440,
	}
	return reminders
}

func getReminderNames() map[int]string {
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
	return reminders
}

func getThresholdIntervalIds() map[string]int {
	thresholdIntervals := map[string]int{
		"default":    0,
		"5 minutes":  5,
		"10 minutes": 10,
		"15 minutes": 15,
		"30 minutes": 30,
		"1 hour":     60,
		"2 hours":    120,
		"6 hours":    360,
		"12 hours":   720,
	}
	return thresholdIntervals
}

func getThresholdIntervalNames() map[int]string {
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
	return thresholdIntervals
}

func getHistoricalIntervalIds() map[string]int {
	historicalIntervals := map[string]int{
		"5 minutes":  5,
		"10 minutes": 10,
		"15 minutes": 15,
		"30 minutes": 30,
		"1 hour":     60,
		"2 hours":    120,
		"6 hours":    360,
		"12 hours":   720,
		"1 day":      1440,
		"1 week":     10080,
	}
	return historicalIntervals
}

func getHistoricalIntervalNames() map[int]string {
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
	return historicalIntervals
}

func getNotificationTypeIds() map[string]int {
	notificationTypes := map[string]int{
		"default contacts": 0,
	}
	return notificationTypes
}

func getAlertTypeIds() map[string]int {
	alertTypes := map[string]int{
		"byte length":   2,
		"content match": 3,
		"host failure":  4,
		"test failure":  9,
		"timing":        7,
		"ping":          12,
		"requests":      13,
		"availability":  15,
		"dns":           17,
		"path":          20,
		"asn":           23,
	}
	return alertTypes
}

func getAlertTypeNames() map[int]string {
	alertTypes := map[int]string{
		2:  "byte length",
		3:  "content match",
		4:  "host failure",
		9:  "test failure",
		7:  "timing",
		12: "ping",
		13: "requests",
		15: "availability",
		17: "dns",
		20: "path",
		23: "asn",
	}
	return alertTypes
}

func getAlertSubTypeIds() map[string]int {
	alertSubTypes := map[string]int{
		"byte length":            1,
		"page":                   2,
		"file size":              3,
		"regular expression":     10,
		"response code":          14,
		"response headers":       15,
		"dns":                    50,
		"connect":                51,
		"send":                   52,
		"wait":                   53,
		"load":                   54,
		"ttfb":                   55,
		"content load":           57,
		"response":               58,
		"test time":              59,
		"dom load":               61,
		"test time with suspect": 63,
		"server response":        64,
		"document complete":      66,
		"redirect":               67,
		"ping rtt":               100,
		"ping packet loss":       101,
		"# requests":             110,
		"# hosts":                111,
		"# connections":          112,
		"# redirects":            113,
		"# other":                114,
		"# images":               115,
		"# scripts":              116,
		"# html":                 117,
		"# css":                  118,
		"# xml":                  119,
		"# flash":                120,
		"# media":                121,
		"test":                   140,
		"content":                141,
		"% downtime":             142,
		"dns answer":             161,
		"# cities":               190,
		"# asns":                 191,
		"# countries":            193,
		"# hops":                 194,
		"handshake_time":         195,
		"days_to_expiration":     196,
		"origin as":              210,
		"path as":                211,
		"origin neighbor":        212,
		"prefix mismatch":        213,
	}
	return alertSubTypes
}

func getAlertSubTypeNames() map[int]string {
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
		161: "dns answer",
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
	return alertSubTypes
}

func getStatisticalTypeIds() map[string]int {
	statisticalTypes := map[string]int{
		"average": 1,
	}
	return statisticalTypes
}

func getTestFlagIds() map[string]int {
	testFlagTypes := map[string]int{
		"verify_test_on_failure":               2,
		"debug_primary_host_on_failure":        3,
		"enable_http2":                         4,
		"debug_referenced_hosts_on_failure":    8,
		"capture_http_headers":                 9,
		"capture_response_content":             11,
		"capture_filmstrip":                    13,
		"capture_screenshot":                   14,
		"ignore_ssl_failures":                  17,
		"enable_bind_hostname":                 19,
		"enable_tcp_protocol":                  20,
		"enable_nsid":                          21,
		"disable_recursive_resolution":         22,
		"host_data_collection_enabled":         23,
		"zone_data_collection_enabled":         24,
		"stop_test_on_document_complete":       25,
		"try_next_nameserver_on_failure":       26,
		"f40x_or_50x_http_mark_successful":     27,
		"favor_fastest_round_trip_nameserver":  31,
		"t30x_redirects_do_not_follow":         33,
		"enable_self_versus_third_party_zones": 36,
		"allow_test_download_limit_override":   37,
		"disable_cross_origin_iframe_access":   38,
		"stop_test_on_dom_content_load":        39,
		"certificate_revocation_disabled":      42,
		"enable_dnssec":                        48,
		"enable_path_mtu_discovery":            50,
	}
	return testFlagTypes
}

func getTestFlagNames() map[int]string {
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
		42: "certificate_revocation_disabled",
		48: "enable_dnssec",
		50: "enable_path_mtu_discovery",
	}
	return testFlagTypes
}

func getAdditionalMonitorTypeIds() map[string]int {
	additionalMonitorTypes := map[string]int{
		"ping icmp":       8,
		"ping tcp":        11,
		"ping udp":        23,
		"traceroute icmp": 9,
		"traceroute udp":  14,
		"traceroute tcp":  29,
	}
	return additionalMonitorTypes

}

func getAdditionalMonitorTypeNames() map[int]string {
	additionalMonitorTypes := map[int]string{
		8:  "ping icmp",
		11: "ping tcp",
		23: "ping udp",
		9:  "traceroute icmp",
		14: "traceroute udp",
		29: "traceroute tcp",
	}
	return additionalMonitorTypes
}

func getBandwidthThrottlingTypeIds() map[string]int {
	bandwidthThrottlingTypes := map[string]int{
		"gprs":       1,
		"regular 2g": 2,
		"good 2g":    3,
		"regular 3g": 4,
		"good 3g":    5,
		"regular 4g": 6,
		"dsl":        7,
		"wifi":       8,
	}
	return bandwidthThrottlingTypes
}

func getBandwidthThrottlingTypeNames() map[int]string {
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
	return bandwidthThrottlingTypes
}

func getReqHeaderTypeIds() map[string]int {
	requestHeaderTypes := map[string]int{
		"user_agent":            1,
		"accept":                2,
		"accept_encoding":       3,
		"accept_language":       4,
		"accept_charset":        5,
		"cookie":                6,
		"cache_control":         7,
		"connection":            8,
		"pragma":                9,
		"referer":               10,
		"custom":                11,
		"sni_override":          11,
		"host":                  12,
		"request_override":      13,
		"dns_override":          14,
		"request_block":         15,
		"request_delay":         16,
		"dns_resolver_override": 17,
	}
	return requestHeaderTypes

}

func getReqHeaderTypeNames() map[int]string {
	requestHeaderTypes := map[int]string{
		1:  "user_agent",
		2:  "accept",
		3:  "accept_encoding",
		4:  "accept_language",
		5:  "accept_charset",
		6:  "cookie",
		7:  "cache_control",
		8:  "connection",
		9:  "pragma",
		10: "referer",
		12: "host",
		13: "request_override",
		14: "dns_override",
		15: "request_block",
		16: "request_delay",
		17: "dns_resolver_override",
	}
	return requestHeaderTypes

}

func getAuthenticationTypeIds() map[string]int {
	authenticationTypes := map[string]int{
		"basic":  1,
		"digest": 2,
		"ntlm":   3,
		"login":  5,
	}
	return authenticationTypes
}

func getDnsTypeIds() map[string]int {
	dnsRecordTypes := map[string]int{
		"a":       1,
		"ns":      2,
		"cname":   5,
		"aaaa":    28,
		"aoraaaa": 32770,
	}
	return dnsRecordTypes
}

func getFilterTypeIds() map[string]int {
	filterTypes := map[string]int{
		"index":   1,
		"last":    1,
		"address": 3,
	}
	return filterTypes

}

func getAlertSettingTypeIds() map[string]int {
	alertSettingTypes := map[string]int{
		"override":      1,
		"inherit & add": 2,
	}
	return alertSettingTypes
}

func getAlertSettingTypeNames() map[int]string {
	alertSettingTypes := map[int]string{
		0: "inherit",
		1: "override",
		2: "inherit & add",
	}
	return alertSettingTypes
}
