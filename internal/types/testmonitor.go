package types

import (
	"catchpoint-provider/internal"
	"slices"
)

// #region TestTypes

var (
	monitorTypeNames    map[int]string
	apiScriptTypeNames  map[int]string
	userAgentTypeNames  map[int]string
	dnsQueryTypeNames   map[int]string
	ValidDNSQueryTypes  []string
	dnsRecordTypeNames  map[int]string
	ValidDNSRecordTypes []string
	ValidChromeVersions []string
	ValidUserAgentTypes []string
)

type TestType int

const (
	WebType         TestType = 0
	TransactionType TestType = 1
	DNSType         TestType = 5
	PingType        TestType = 6
	APIType         TestType = 9
	TracerouteType  TestType = 12
	SSLType         TestType = 18
	BGPType         TestType = 20
	PlaywrightType  TestType = 25
	PuppeteerType   TestType = 26
)

func GetTestTypeName(testType TestType) string {
	switch testType {
	case APIType:
		return "api"
	case BGPType:
		return "bgp"
	case DNSType:
		return "dns"
	case SSLType:
		return "ssl"
	case TracerouteType:
		return "traceroute"
	case TransactionType:
		return "transaction"
	case PingType:
		return "ping"
	case PlaywrightType:
		return "playwright"
	case PuppeteerType:
		return "puppeteer"
	case WebType:
		return "web"
	default:
		return ""
	}
}

// #endregion

// #region MonitorTypes

var MonitorTypeIDs = map[string]int{
	HTTPString:           2,
	EmulatedString:       3,
	ChromeString:         18,
	PlaybackString:       19,
	MobilePlaybackString: 20,
	MobileString:         26,
	APIString:            25,
	PingICMPString:       8,
	PingTCPString:        11,
	DNSExperienceString:  12,
	DNSDirectString:      13,
	PingUDPString:        23,
	TracerouteICMPString: 9,
	TracerouteUDPString:  14,
	TracerouteTCPString:  29,
	SSLString:            31,
	BGPString:            34,
	EdgeString:           39,
	BGPBasicString:       41,
}

func GetMonitorTypeID(name string) (int, bool) {
	id, ok := MonitorTypeIDs[name]
	return id, ok
}

func GetMonitorTypeName(id int) (string, bool) {
	name, ok := monitorTypeNames[id]
	return name, ok
}

// #endregion

// #region APIScriptTypes

var APIScriptTypeIDs = map[string]int{
	SeleniumString:   1,
	JavaScriptString: 2,
	PlaywrightString: 3,
	PuppeteerString:  4,
}

func GetAPIScriptTypeID(name string) (int, bool) {
	id, ok := APIScriptTypeIDs[name]
	return id, ok
}

func GetAPIScriptTypeName(id int) (string, bool) {
	name, ok := apiScriptTypeNames[id]
	return name, ok
}

// #endregion

// #region UserAgentTypes

var UserAgentTypeIDs = map[string]int{
	UserAgentIE:            1,
	UserAgentChrome:        2,
	UserAgentAndroid:       3,
	UserAgentiPhone:        4,
	UserAgentiPad2:         5,
	UserAgentKindleFire:    6,
	UserAgentGalaxyTab:     7,
	UserAgentiPhone5:       8,
	UserAgentiPadMini:      9,
	UserAgentGalaxyNote:    10,
	UserAgentNexus7:        11,
	UserAgentNexus4:        12,
	UserAgentNokiaLumia920: 13,
	UserAgentIPhone6:       14,
	UserAgentBlackberryZ30: 15,
	UserAgentGalaxyS4:      16,
	UserAgentHTCOneX:       17,
	UserAgentLGOptimusG:    18,
	UserAgentDroidRazrHD:   19,
	UserAgentNexus6:        20,
	UserAgentIPhone6S:      21,
	UserAgentGalaxyS6:      22,
	UserAgentIPhone7:       23,
	UserAgentGooglePixel:   24,
	UserAgentGalaxyS8:      25,
}

func GetUserAgentTypeID(name string) (int, bool) {
	id, ok := UserAgentTypeIDs[name]
	return id, ok
}

func GetUserAgentTypeName(id int) (string, bool) {
	name, ok := userAgentTypeNames[id]
	return name, ok
}

// #endregion

// #region ChromeVersionIDs

var ChromeVersionIDs = map[int][]string{
	1: {Stable},
	2: {Preview},
	3: {Chrome120, Chrome108, Chrome97, Chrome89, Chrome87, Chrome85, Chrome79, Chrome75, Chrome71, Chrome66, Chrome63, Chrome59, Chrome53}, // "Specific"
}

var ChromeApplicationVersionIDsProd = map[string]int{
	Chrome53: 1,
	Chrome59: 3,
	Chrome63: 4,
	Chrome66: 5,
	Chrome75: 7,
	Chrome71: 8,
	// Note: 79 never added to prod.
	Chrome85:  12,
	Chrome87:  13,
	Chrome89:  14,
	Chrome97:  28357,
	Chrome108: 28558,
	Chrome120: 31965,
}

var ChromeApplicationVersionIDsQA = map[string]int{
	Chrome53:  1,
	Chrome59:  3,
	Chrome63:  4,
	Chrome66:  5,
	Chrome71:  6,
	Chrome75:  7,
	Chrome79:  9,
	Chrome85:  10,
	Chrome87:  13,
	Chrome89:  12,
	Chrome108: 15,
	Chrome120: 44,
}

var ChromeApplicationVersionIDsStage = map[string]int{
	Chrome53:  1,
	Chrome59:  3,
	Chrome63:  4,
	Chrome66:  5,
	Chrome71:  6,
	Chrome75:  7,
	Chrome79:  9,
	Chrome85:  10,
	Chrome87:  11,
	Chrome89:  12,
	Chrome97:  3744,
	Chrome108: 3753,
	Chrome120: 4181,
}

func GetChromeVersionIDs() map[int][]string {
	return ChromeVersionIDs
}

func GetChromeVersionID(name string) (int, bool) {
	for id, names := range ChromeVersionIDs {
		if slices.Contains(names, name) {
			return id, true
		}
	}
	return 0, false
}

func GetChromeApplicationVersionIDs() map[string]int {
	switch internal.CatchpointEnvironment {
	case "stage":
		return ChromeApplicationVersionIDsStage
	case "qa":
		return ChromeApplicationVersionIDsQA
	default:
		return ChromeApplicationVersionIDsProd
	}
}

func GetChromeApplicationVersionID(name string) (int, bool) {
	vid, ok := GetChromeApplicationVersionIDs()[name]
	return vid, ok
}

func GetChromeApplicationVersionString(id int) (string, bool) {
	for name, vid := range GetChromeApplicationVersionIDs() {
		if vid == id {
			return name, true
		}
	}
	return EmptyString, false
}

// #endregion

// #region DNS

var DNSQueryTypeIDs = map[string]int{
	None:       0,
	A:          1,
	NS:         2,
	CName:      5,
	SOA:        6,
	MB:         7,
	MG:         8,
	MR:         9,
	Null:       10,
	WKS:        11,
	PTR:        12,
	HInfo:      13,
	MInfo:      14,
	MX:         15,
	TXT:        16,
	RP:         17,
	AFSDB:      18,
	X25:        19,
	ISDN:       20,
	RT:         21,
	NSAP:       22,
	Sig:        24,
	Key:        25,
	PX:         26,
	AAAA:       28,
	Loc:        29,
	EID:        31,
	NIMLOC:     32,
	SRV:        33,
	ATMA:       34,
	NAPTR:      35,
	KX:         36,
	CERT:       37,
	A6:         38,
	DName:      39,
	SINK:       40,
	OPT:        41,
	APL:        42,
	DS:         43,
	SSHFP:      44,
	IPSECKEY:   45,
	RRSIG:      46,
	NSEC:       47,
	DNSKEY:     48,
	DHCID:      49,
	NSEC3:      50,
	NSEC3PARAM: 51,
	HIP:        55,
	SPF:        99,
	UINFO:      100,
	UID:        101,
	GID:        102,
	UNSPEC:     103,
	TKEY:       249,
	TSIG:       250,
	IXFR:       251,
	AXFR:       252,
	MAILB:      253,
	ANY:        255,
	TA:         32768,
	DLV:        32769,
	AorAAAA:    32770,
}

func GetDNSQueryTypeID(name string) (int, bool) {
	id, ok := DNSQueryTypeIDs[name]
	return id, ok
}

func GetDNSQueryTypeName(id int) (string, bool) {
	name, ok := dnsQueryTypeNames[id]
	return name, ok
}

// This is for Alerting, as opposed to the longer DNSQueryTypeIDs which is used for DNS Test creation.
var DNSRecordTypeIDs = map[string]int{
	A:       1,
	NS:      2,
	CName:   5,
	AAAA:    28,
	AorAAAA: 32770,
}

func GetDNSRecordTypeID(name string) (int, bool) {
	id, ok := DNSRecordTypeIDs[name]
	return id, ok
}

func GetDNSRecordTypeName(id int) (string, bool) {
	name, ok := dnsRecordTypeNames[id]
	return name, ok
}

// #endregion

// On init, generate reverse mappings for later O(1) lookups.
func init() {
	monitorTypeNames = populateReverseMap(MonitorTypeIDs)
	apiScriptTypeNames = populateReverseMap(APIScriptTypeIDs)
	userAgentTypeNames = populateReverseMap(UserAgentTypeIDs)

	dnsQueryTypeNames = populateReverseMap(DNSQueryTypeIDs)
	ValidDNSQueryTypes = populateValidNames(DNSQueryTypeIDs)

	dnsRecordTypeNames = populateReverseMap(DNSRecordTypeIDs)
	ValidDNSRecordTypes = populateValidNames(DNSRecordTypeIDs)

	// ValidChromeVersions can be either the specific versions or "stable" or "preview".
	ValidChromeVersions = populateValidNames(ChromeApplicationVersionIDsProd)
	ValidChromeVersions = append(ValidChromeVersions, Stable, Preview)

	ValidUserAgentTypes = populateValidNames(UserAgentTypeIDs)
}
