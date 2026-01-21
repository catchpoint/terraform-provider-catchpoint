package types

var (
	ValidAuthenticationTypeNames      []string
	ValidAdditionalMonitorTypeNames   []string
	ValidBandwidthThrottlingTypeNames []string
	ValidTestFlagNames                []string
	ValidHTTPRequestHeaders           []string
	testFlagNames                     map[int]string
	additionalMonitorTypeNames        map[int]string
	bandwidthThrottlingTypeNames      map[int]string
	reqHeaderTypeNames                map[int]string
	authenticationTypeNames           map[int]string
)

// #region TestFlags

var TestFlagIDs = map[string]int{
	//CacheTLDQueryLevels: 1 TODO: enable this
	VerifyTestOnFailure:             2,
	DebugPrimaryHostOnFailure:       3,
	DebugReferencedHostsOnFailure:   8,
	CaptureHTTPHeaders:              9,
	CaptureResponseContent:          11,
	CaptureFilmstrip:                13,
	CaptureScreenshot:               14,
	IgnoreSSLFailures:               17,
	EnableBindHostname:              19,
	EnableTCPProtocol:               20,
	EnableNSID:                      21, // This doesn't seem to exist in appliedTestFlags for Folders or Products
	DisableRecursiveResolution:      22,
	HostDataCollectionEnabled:       23,
	ZoneDataCollectionEnabled:       24,
	StopTestOnDocumentComplete:      25,
	TryNextNameserverOnFailure:      26,
	F40xOr50xHTTPMarkSuccessful:     27,
	FavorFastestRoundTripNameserver: 31,
	T30xRedirectsDoNotFollow:        33,
	EnableSelfVersusThirdPartyZones: 36,
	AllowTestDownloadLimitOverride:  37,
	DisableCrossOriginIframeAccess:  38,
	StopTestOnDOMContentLoad:        39, // This doesn't seem to exist in appliedTestFlags for Folders or Products
	CertificateRevocationDisabled:   42,
	EnableDNSSEC:                    48,
	EnablePathMTUDiscovery:          50,
	// Tracing: 51 TODO: enable this
	// EnableECN: 53 TODO: enable this
	// EnableDSCP: 54 TODO: enable this
	// EnableQueryLimits: 56 TODO: enable this
}

func GetTestFlagID(name string) (int, bool) {
	id, ok := TestFlagIDs[name]
	return id, ok
}

func GetTestFlagName(id int) (string, bool) {
	name, ok := testFlagNames[id]
	return name, ok
}

// #endregion

// #region AdditionalMonitorTypes

var AdditionalMonitorTypeIDs = map[string]int{
	PingICMPString:       8,
	PingTCPString:        11,
	PingUDPString:        23,
	TracerouteICMPString: 9,
	TracerouteUDPString:  14,
	TracerouteTCPString:  29,
	//TODO: add Traceroute Insession, Traceroute QUIC.
}

func GetAdditionalMonitorTypeID(name string) (int, bool) {
	id, ok := AdditionalMonitorTypeIDs[name]
	return id, ok
}

func GetAdditionalMonitorTypeName(id int) (string, bool) {
	name, ok := additionalMonitorTypeNames[id]
	return name, ok
}

// #endregion

// #region BandwidthThrottlingTypes

var BandwidthThrottlingTypeIDs = map[string]int{
	GPRS:      1,
	Regular2G: 2,
	Good2G:    3,
	Regular3G: 4,
	Good3G:    5,
	Regular4G: 6,
	DSL:       7,
	WIFI:      8,
}

func GetBandwidthThrottlingTypeID(name string) (int, bool) {
	id, ok := BandwidthThrottlingTypeIDs[name]
	return id, ok
}

func GetBandwidthThrottlingTypeName(id int) (string, bool) {
	name, ok := bandwidthThrottlingTypeNames[id]
	return name, ok
}

// #endregion

// #region RequestHeaderTypes

var ReqHeaderTypeIDs = map[string]int{
	UserAgentHeader:           1,
	AcceptHeader:              2,
	AcceptEncodingHeader:      3,
	AcceptLanguageHeader:      4,
	AcceptCharsetHeader:       5,
	CookieHeader:              6,
	CacheControlHeader:        7,
	ConnectionHeader:          8,
	PragmaHeader:              9,
	RefererHeader:             10,
	CustomHeader:              11,
	HostHeader:                12,
	RequestOverrideHeader:     13,
	DNSOverrideHeader:         14,
	RequestBlockHeader:        15,
	RequestDelayHeader:        16,
	DNSResolverOverrideHeader: 17,
}

func GetReqHeaderTypeID(name string) (int, bool) {
	id, ok := ReqHeaderTypeIDs[name]

	// Special case: SNI Override is not in the ReqHeaderTypeIDs map because it's a Custom header.
	if name == SNIOverrideHeader {
		id = 11
		ok = true
	}

	return id, ok
}

func GetReqHeaderTypeName(id int) (string, bool) {
	name, ok := reqHeaderTypeNames[id]
	return name, ok
}

// #endregion

// #region AuthenticationTypes

var AuthenticationTypeIDs = map[string]int{
	None:       0,
	BasicAuth:  1,
	DigestAuth: 2,
	NTLMAuth:   3,
	LoginAuth:  5,
}

func GetAuthenticationTypeID(name string) (int, bool) {
	id, ok := AuthenticationTypeIDs[name]
	return id, ok
}

func GetAuthenticationTypeName(id int) (string, bool) {
	name, ok := authenticationTypeNames[id]
	return name, ok
}

// #endregion

// On init, generate reverse mappings for later O(1) lookups.
func init() {
	testFlagNames = populateReverseMap(TestFlagIDs)
	ValidTestFlagNames = populateValidNames(TestFlagIDs)

	additionalMonitorTypeNames = populateReverseMap(AdditionalMonitorTypeIDs)
	ValidAdditionalMonitorTypeNames = populateValidNames(AdditionalMonitorTypeIDs)

	bandwidthThrottlingTypeNames = populateReverseMap(BandwidthThrottlingTypeIDs)
	ValidBandwidthThrottlingTypeNames = populateValidNames(BandwidthThrottlingTypeIDs)

	reqHeaderTypeNames = populateReverseMap(ReqHeaderTypeIDs)
	ValidHTTPRequestHeaders = populateValidNames(ReqHeaderTypeIDs)

	authenticationTypeNames = populateReverseMap(AuthenticationTypeIDs)
	ValidAuthenticationTypeNames = populateValidNames(AuthenticationTypeIDs)
}
