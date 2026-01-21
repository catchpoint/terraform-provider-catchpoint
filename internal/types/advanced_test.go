package types

import (
	"catchpoint-provider/internal/testutil"
	"testing"
)

func TestGetTestFlagIDAndName(t *testing.T) {
	test := []struct {
		testFlag   string
		testFlagID int
	}{
		{VerifyTestOnFailure, 2},
		{DebugPrimaryHostOnFailure, 3},
		{DebugReferencedHostsOnFailure, 8},
		{CaptureHTTPHeaders, 9},
		{CaptureResponseContent, 11},
		{CaptureFilmstrip, 13},
		{CaptureScreenshot, 14},
		{IgnoreSSLFailures, 17},
		{EnableBindHostname, 19},
		{EnableTCPProtocol, 20},
		{EnableNSID, 21},
		{DisableRecursiveResolution, 22},
		{HostDataCollectionEnabled, 23},
		{ZoneDataCollectionEnabled, 24},
		{StopTestOnDocumentComplete, 25},
		{TryNextNameserverOnFailure, 26},
		{F40xOr50xHTTPMarkSuccessful, 27},
		{FavorFastestRoundTripNameserver, 31},
		{T30xRedirectsDoNotFollow, 33},
		{EnableSelfVersusThirdPartyZones, 36},
		{AllowTestDownloadLimitOverride, 37},
		{DisableCrossOriginIframeAccess, 38},
		{StopTestOnDOMContentLoad, 39},
		{CertificateRevocationDisabled, 42},
		{EnableDNSSEC, 48},
		{EnablePathMTUDiscovery, 50},
	}
	for _, tt := range test {
		id := testutil.GetOrFail(t, func() (int, bool) { return GetTestFlagID(tt.testFlag) }, "GetTestFlagID", tt.testFlag)
		name := testutil.GetOrFail(t, func() (string, bool) { return GetTestFlagName(id) }, "GetTestFlagName", id)

		testutil.AssertEqual(t, "testFlagID", id, tt.testFlagID)
		testutil.AssertEqual(t, "testFlagName", name, tt.testFlag)
	}
}

func TestGetAdditionalMonitorTypeIDAndName(t *testing.T) {
	tests := []struct {
		name string
		ID   int
	}{
		{PingICMPString, 8},
		{PingTCPString, 11},
		{PingUDPString, 23},
		{TracerouteICMPString, 9},
		{TracerouteUDPString, 14},
		{TracerouteTCPString, 29},
	}
	for _, tt := range tests {
		id := testutil.GetOrFail(t, func() (int, bool) { return GetAdditionalMonitorTypeID(tt.name) }, "GetAdditionalMonitorTypeID", tt.name)
		name := testutil.GetOrFail(t, func() (string, bool) { return GetAdditionalMonitorTypeName(id) }, "GetAdditionalMonitorTypeName", id)

		testutil.AssertEqual(t, "additionalMonitorTypeID", id, tt.ID)
		testutil.AssertEqual(t, "additionalMonitorTypeName", name, tt.name)
	}
}

func TestGetBandwidthThrottlingTypeIDAndName(t *testing.T) {
	tests := []struct {
		name string
		ID   int
	}{
		{GPRS, 1},
		{Regular2G, 2},
		{Good2G, 3},
		{Regular3G, 4},
		{Good3G, 5},
		{Regular4G, 6},
		{DSL, 7},
		{WIFI, 8},
	}
	for _, tt := range tests {
		id := testutil.GetOrFail(t, func() (int, bool) { return GetBandwidthThrottlingTypeID(tt.name) }, "GetBandwidthThrottlingTypeID", tt.name)
		name := testutil.GetOrFail(t, func() (string, bool) { return GetBandwidthThrottlingTypeName(id) }, "GetBandwidthThrottlingTypeName", id)

		testutil.AssertEqual(t, "bandwidthThrottlingTypeID", id, tt.ID)
		testutil.AssertEqual(t, "bandwidthThrottlingTypeName", name, tt.name)
	}
}

func TestGetReqHeaderTypeIDAndName(t *testing.T) {
	test := []struct {
		name string
		id   int
	}{
		{UserAgentHeader, 1},
		{AcceptHeader, 2},
		{AcceptEncodingHeader, 3},
		{AcceptLanguageHeader, 4},
		{AcceptCharsetHeader, 5},
		{CookieHeader, 6},
		{CacheControlHeader, 7},
		{ConnectionHeader, 8},
		{PragmaHeader, 9},
		{RefererHeader, 10},
		{HostHeader, 12},
		{RequestOverrideHeader, 13},
		{DNSOverrideHeader, 14},
		{RequestBlockHeader, 15},
		{RequestDelayHeader, 16},
		{DNSResolverOverrideHeader, 17},
	}
	for _, tt := range test {
		id := testutil.GetOrFail(t, func() (int, bool) { return GetReqHeaderTypeID(tt.name) }, "GetReqHeaderTypeID", tt.name)
		name := testutil.GetOrFail(t, func() (string, bool) { return GetReqHeaderTypeName(id) }, "GetReqHeaderTypeName", id)

		testutil.AssertEqual(t, "reqHeaderTypeID", id, tt.id)
		testutil.AssertEqual(t, "reqHeaderTypeName", name, tt.name)
	}
}

func TestGetAuthenticationTypeIDAndName(t *testing.T) {
	test := []struct {
		name string
		id   int
	}{
		{BasicAuth, 1},
		{DigestAuth, 2},
		{NTLMAuth, 3},
		{LoginAuth, 5},
	}
	for _, tt := range test {
		id := testutil.GetOrFail(t, func() (int, bool) { return GetAuthenticationTypeID(tt.name) }, "GetAuthenticationTypeID", tt.name)
		name := testutil.GetOrFail(t, func() (string, bool) { return GetAuthenticationTypeName(id) }, "GetAuthenticationTypeName", id)

		testutil.AssertEqual(t, "authenticationTypeID", id, tt.id)
		testutil.AssertEqual(t, "authenticationTypeName", name, tt.name)
	}
}
