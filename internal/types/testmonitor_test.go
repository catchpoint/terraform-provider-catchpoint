package types

import (
	"testing"

	"catchpoint-provider/internal"
	"catchpoint-provider/internal/testutil"
)

func TestGetMonitorTypeIDAndName(t *testing.T) {
	tests := []struct {
		monitor   string
		monitorID int
	}{
		{HTTPString, 2},
		{EmulatedString, 3},
		{ChromeString, 18},
		{PlaybackString, 19},
		{MobilePlaybackString, 20},
		{MobileString, 26},
		{APIString, 25},
		{PingICMPString, 8},
		{PingTCPString, 11},
		{DNSExperienceString, 12},
		{DNSDirectString, 13},
		{PingUDPString, 23},
		{TracerouteICMPString, 9},
		{TracerouteUDPString, 14},
		{TracerouteTCPString, 29},
		{SSLString, 31},
		{BGPString, 34},
		{EdgeString, 39},
		{BGPBasicString, 41},
	}
	for _, tt := range tests {
		id := testutil.GetOrFail(t, func() (int, bool) { return GetMonitorTypeID(tt.monitor) }, "GetMonitorTypeID", tt.monitor)
		name := testutil.GetOrFail(t, func() (string, bool) { return GetMonitorTypeName(tt.monitorID) }, "GetMonitorTypeName", tt.monitorID)

		testutil.AssertEqual(t, "monitorID", id, tt.monitorID)
		testutil.AssertEqual(t, "monitorName", name, tt.monitor)
	}
}

func TestGetAPIScriptTypeIDAndName(t *testing.T) {
	tests := []struct {
		scriptType   string
		scriptTypeID int
	}{
		{SeleniumString, 1},
		{JavaScriptString, 2},
		{PlaywrightString, 3},
		{PuppeteerString, 4},
	}
	for _, tt := range tests {
		id := testutil.GetOrFail(t, func() (int, bool) { return GetAPIScriptTypeID(tt.scriptType) }, "GetAPIScriptTypeID", tt.scriptType)
		name := testutil.GetOrFail(t, func() (string, bool) { return GetAPIScriptTypeName(tt.scriptTypeID) }, "GetAPIScriptTypeName", tt.scriptTypeID)

		testutil.AssertEqual(t, "scriptTypeID", id, tt.scriptTypeID)
		testutil.AssertEqual(t, "scriptTypeName", name, tt.scriptType)
	}
}

func TestGetUserAgentTypeIDAndName(t *testing.T) {
	tests := []struct {
		userAgentType   string
		userAgentTypeID int
	}{
		{UserAgentIE, 1},
		{UserAgentChrome, 2},
		{UserAgentAndroid, 3},
		{UserAgentiPhone, 4},
		{UserAgentiPad2, 5},
		{UserAgentKindleFire, 6},
		{UserAgentGalaxyTab, 7},
		{UserAgentiPhone5, 8},
		{UserAgentiPadMini, 9},
		{UserAgentGalaxyNote, 10},
		{UserAgentNexus7, 11},
		{UserAgentNexus4, 12},
		{UserAgentNokiaLumia920, 13},
		{UserAgentIPhone6, 14},
		{UserAgentBlackberryZ30, 15},
		{UserAgentGalaxyS4, 16},
		{UserAgentHTCOneX, 17},
		{UserAgentLGOptimusG, 18},
		{UserAgentDroidRazrHD, 19},
		{UserAgentNexus6, 20},
		{UserAgentIPhone6S, 21},
		{UserAgentGalaxyS6, 22},
		{UserAgentIPhone7, 23},
		{UserAgentGooglePixel, 24},
		{UserAgentGalaxyS8, 25},
	}
	for _, tt := range tests {
		id := testutil.GetOrFail(t, func() (int, bool) { return GetUserAgentTypeID(tt.userAgentType) }, "GetUserAgentTypeID", tt.userAgentType)
		name := testutil.GetOrFail(t, func() (string, bool) { return GetUserAgentTypeName(tt.userAgentTypeID) }, "GetUserAgentTypeName", tt.userAgentTypeID)

		testutil.AssertEqual(t, "userAgentTypeID", id, tt.userAgentTypeID)
		testutil.AssertEqual(t, "userAgentTypeName", name, tt.userAgentType)
	}
}

func TestGetChromeVersionIDs(t *testing.T) {
	versions := GetChromeVersionIDs()
	if len(versions) == 0 {
		t.Error("Expected non-empty ChromeVersionIDs")
	}
}

func TestGetChromeVersionID(t *testing.T) {
	tests := []struct {
		chrome     string
		expectedID int
	}{
		{Stable, 1},
		{Preview, 2},
		{Chrome53, 3},
		{Chrome59, 3},
		{Chrome63, 3},
		{Chrome66, 3},
		{Chrome75, 3},
		{Chrome71, 3},
		{Chrome85, 3},
		{Chrome87, 3},
		{Chrome89, 3},
		{Chrome108, 3},
		{Chrome120, 3},
	}
	for _, tt := range tests {
		id := testutil.GetOrFail(t, func() (int, bool) { return GetChromeVersionID(tt.chrome) }, "GetChromeVersionID", tt.chrome)
		testutil.AssertEqual(t, "chromeVersionID", id, tt.expectedID)
	}
}

func TestGetChromeApplicationVersionIDs(t *testing.T) {
	tests := []struct {
		chrome        string
		chromeIDProd  int
		chromeIDStage int
	}{
		{Chrome53, 1, 1},
		{Chrome59, 3, 3},
		{Chrome63, 4, 4},
		{Chrome66, 5, 5},
		{Chrome75, 7, 7},
		{Chrome71, 8, 6},
		{Chrome85, 12, 10},
		{Chrome87, 13, 11},
		{Chrome89, 14, 12},
		{Chrome108, 28558, 3753},
		{Chrome120, 31965, 4181},
	}
	for _, tt := range tests {
		// Save and restore environment
		origEnv := internal.CatchpointEnvironment
		defer func() { internal.CatchpointEnvironment = origEnv }()

		internal.CatchpointEnvironment = "prod"
		prodIDs := GetChromeApplicationVersionIDs()
		testutil.AssertEqual(t, "prodChromeID", prodIDs[tt.chrome], tt.chromeIDProd)

		internal.CatchpointEnvironment = "stage"
		stageIDs := GetChromeApplicationVersionIDs()
		testutil.AssertEqual(t, "stageChromeID", stageIDs[tt.chrome], tt.chromeIDStage)
	}
}

func TestGetDNSQueryTypeIDAndName(t *testing.T) {
	tests := []struct {
		queryType   string
		queryTypeID int
	}{
		{None, 0},
		{A, 1},
		{NS, 2},
		{CName, 5},
		{SOA, 6},
		{MB, 7},
		{MG, 8},
		{MR, 9},
		{Null, 10},
		{WKS, 11},
		{PTR, 12},
		{HInfo, 13},
		{MInfo, 14},
		{MX, 15},
		{TXT, 16},
		{RP, 17},
		{AFSDB, 18},
		{X25, 19},
		{ISDN, 20},
		{RT, 21},
		{NSAP, 22},
		{Sig, 24},
		{Key, 25},
		{PX, 26},
		{AAAA, 28},
		{Loc, 29},
		{EID, 31},
		{NIMLOC, 32},
		{SRV, 33},
		{ATMA, 34},
		{NAPTR, 35},
		{KX, 36},
		{CERT, 37},
		{A6, 38},
		{DName, 39},
		{SINK, 40},
		{OPT, 41},
		{APL, 42},
		{DS, 43},
		{SSHFP, 44},
		{IPSECKEY, 45},
		{RRSIG, 46},
		{NSEC, 47},
		{DNSKEY, 48},
		{DHCID, 49},
		{NSEC3, 50},
		{NSEC3PARAM, 51},
		{HIP, 55},
		{SPF, 99},
		{UINFO, 100},
		{UID, 101},
		{GID, 102},
		{UNSPEC, 103},
		{TKEY, 249},
		{TSIG, 250},
		{IXFR, 251},
		{AXFR, 252},
		{MAILB, 253},
		{ANY, 255},
		{TA, 32768},
		{DLV, 32769},
		{AorAAAA, 32770},
	}
	for _, tt := range tests {
		id := testutil.GetOrFail(t, func() (int, bool) { return GetDNSQueryTypeID(tt.queryType) }, "GetDNSQueryTypeID", tt.queryType)
		name := testutil.GetOrFail(t, func() (string, bool) { return GetDNSQueryTypeName(tt.queryTypeID) }, "GetDNSQueryTypeName", tt.queryTypeID)

		testutil.AssertEqual(t, "queryTypeID", id, tt.queryTypeID)
		testutil.AssertEqual(t, "queryTypeName", name, tt.queryType)
	}
}

func TestGetDNSRecordTypeIDAndName(t *testing.T) {
	test := []struct {
		recordType   string
		recordTypeID int
	}{
		{A, 1},
		{NS, 2},
		{CName, 5},
		{AAAA, 28},
		{AorAAAA, 32770},
	}
	for _, tt := range test {
		id := testutil.GetOrFail(t, func() (int, bool) { return GetDNSRecordTypeID(tt.recordType) }, "GetDNSRecordTypeID", tt.recordType)
		name := testutil.GetOrFail(t, func() (string, bool) { return GetDNSRecordTypeName(tt.recordTypeID) }, "GetDNSRecordTypeName", tt.recordTypeID)

		testutil.AssertEqual(t, "recordTypeID", id, tt.recordTypeID)
		testutil.AssertEqual(t, "recordTypeName", name, tt.recordType)
	}
}
