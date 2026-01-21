package types

type Interval string

// Various intervals used throughout.
const (
	// FrequencyID 1
	// ReminderID 1
	OneMinute Interval = "1 minute"
	// FrequencyID 16
	TwoMinutes Interval = "2 minutes"
	// FrequencyID 15
	FourMinutes Interval = "4 minutes"
	// HistoricalTypeID 5
	// FrequencyID 2
	// ReminderID 5
	// ThresholdIntervalID 5
	FiveMinutes Interval = "5 minutes"
	// HistoricalTypeID 10
	// FrequencyID 3
	// ReminderID 10
	// ThresholdIntervalID 10
	TenMinutes Interval = "10 minutes"
	// HistoricalTypeID 15
	// FrequencyID 4
	// ReminderID 15
	// ThresholdIntervalID 15
	FifteenMinutes Interval = "15 minutes"
	// FrequencyID 5
	TwentyMinutes Interval = "20 minutes"
	// HistoricalTypeID 30
	// FrequencyID 6
	// ReminderID 30
	// ThresholdIntervalID 30
	ThirtyMinutes Interval = "30 minutes"
	// FrequencyID 7
	SixtyMinutes Interval = "60 minutes"
	// HistoricalTypeID 60
	// ReminderID 60
	// ThresholdIntervalID 60
	OneHour Interval = "1 hour"
	// HistoricalTypeID 120
	// FrequencyID 8
	// ThresholdIntervalID 120
	TwoHours Interval = "2 hours"
	// FrequencyID 9
	ThreeHours Interval = "3 hours"
	// FrequencyID 10
	FourHours Interval = "4 hours"
	// HistoricalTypeID 360
	// FrequencyID 11
	// ThresholdIntervalID 360
	SixHours Interval = "6 hours"
	// FrequencyID 12
	EightHours Interval = "8 hours"
	// HistoricalTypeID 720
	// FrequencyID 13
	// ThresholdIntervalID 720
	TwelveHours Interval = "12 hours"
	// FrequencyID 14
	TwentyFourHours Interval = "24 hours"
	// ReminderID 1440
	Daily Interval = "daily"
	// HistoricalTypeID 1440
	OneDay Interval = "1 day"
	// HistoricalTypeID 10080
	OneWeek Interval = "1 week"
)

// Can be used by various types for inheriting/overriding settings.
const (
	// AlertSettingTypeID 1, GenericSettingTypeID 1
	Override = "override"
	// AlertSettingTypeID 0, GenericSettingTypeID 0
	Inherit = "inherit"
	// AlertSettingTypeID 2
	InheritAndAdd = "inherit & add"
	// InsightSettingTypeID 3
	NoSettings = "no settings"
)

// Node Distributions (Schedule Setting).
const (
	// NodeDistributionID 0
	Random = "random"
	// NodeDistributionID 1
	Concurrent = "concurrent"
)

// Status Types.
const (
	Active   = "active"
	Inactive = "inactive"
)

// Various values.
const (
	EmptyString = ""
	// StatisticalTypeID 1
	Average = "average"
	// NotificationTypeID 0
	DefaultContacts = "default contacts"
	// ReminderID 0
	// FrequencyID 0
	None = "none"
	// ThresholdIntervalID 0
	Default = "default"
)

// Auth types (Request Setting).
const (
	BasicAuth  = "basic"
	DigestAuth = "digest"
	NTLMAuth   = "ntlm"
	LoginAuth  = "login"
)

// Request Header types (Request Setting).
const (
	AcceptCharsetHeader       = "accept_charset"
	AcceptEncodingHeader      = "accept_encoding"
	AcceptHeader              = "accept"
	AcceptLanguageHeader      = "accept_language"
	CacheControlHeader        = "cache_control"
	ConnectionHeader          = "connection"
	CookieHeader              = "cookie"
	CustomHeader              = "custom"
	DNSOverrideHeader         = "dns_override"
	DNSResolverOverrideHeader = "dns_resolver_override"
	HostHeader                = "host"
	PragmaHeader              = "pragma"
	RefererHeader             = "referer"
	RequestBlockHeader        = "request_block"
	RequestDelayHeader        = "request_delay"
	RequestOverrideHeader     = "request_override"
	SNIOverrideHeader         = "sni_override"
	UserAgentHeader           = "user_agent"
)

// DNS Query Types, used for both Test settings and specific alerting.
const (
	A          = "a"
	A6         = "a6"
	AAAA       = "aaaa"
	AFSDB      = "afsdb"
	ANY        = "any"
	AorAAAA    = "AorAAAA"
	APL        = "apl"
	ATMA       = "atma"
	AXFR       = "axfr"
	CERT       = "cert"
	CName      = "cname"
	DHCID      = "dhcid"
	DLV        = "dlv"
	DName      = "dname"
	DNSKEY     = "dnskey"
	DS         = "ds"
	EID        = "eid"
	GID        = "gid"
	HInfo      = "hinfo"
	HIP        = "hip"
	IPSECKEY   = "ipseckey"
	ISDN       = "isdn"
	IXFR       = "ixfr"
	Key        = "key"
	KX         = "kx"
	Loc        = "loc"
	MAILB      = "mailb"
	MB         = "mb"
	MG         = "mg"
	MInfo      = "minfo"
	MR         = "mr"
	MX         = "mx"
	NAPTR      = "naptr"
	NIMLOC     = "nimloc"
	NS         = "ns"
	NSAP       = "nsap"
	NSEC       = "nsec"
	NSEC3      = "nsec3"
	NSEC3PARAM = "nsec3param"
	Null       = "null"
	OPT        = "opt"
	PTR        = "ptr"
	PX         = "px"
	RP         = "rp"
	RRSIG      = "rrsig"
	RT         = "rt"
	Sig        = "sig"
	SINK       = "sink"
	SOA        = "soa"
	SPF        = "spf"
	SRV        = "srv"
	SSHFP      = "sshfp"
	TA         = "ta"
	TKEY       = "tkey"
	TSIG       = "tsig"
	TXT        = "txt"
	UID        = "uid"
	UINFO      = "uinfo"
	UNSPEC     = "unspec"
	WKS        = "wks"
	X25        = "x25"
)

// TestFlags (Advanced Settings) applied to multiple resource types (Product, Folder, Tests).
const (
	// TestFlagID 37
	AllowTestDownloadLimitOverride = "allow_test_download_limit_override"
	// TestFlagID 13
	CaptureFilmstrip = "capture_filmstrip"
	// TestFlagID 9
	CaptureHTTPHeaders = "capture_http_headers"
	// TestFlagID 11
	CaptureResponseContent = "capture_response_content"
	// TestFlagID 14
	CaptureScreenshot = "capture_screenshot"
	// TestFlagID 42
	CertificateRevocationDisabled = "certificate_revocation_disabled"
	// TestFlagID 3
	DebugPrimaryHostOnFailure = "debug_primary_host_on_failure"
	// TestFlagID 8
	DebugReferencedHostsOnFailure = "debug_referenced_hosts_on_failure"
	// TestFlagID 38
	DisableCrossOriginIframeAccess = "disable_cross_origin_iframe_access"
	// TestFlagID 22
	DisableRecursiveResolution = "disable_recursive_resolution"
	// TestFlagID 19
	EnableBindHostname = "enable_bind_hostname"
	// TestFlagID 48
	EnableDNSSEC = "enable_dnssec"
	// TestFlagID 21
	EnableNSID = "enable_nsid"
	// TestFlagID 50
	EnablePathMTUDiscovery = "enable_path_mtu_discovery"
	// TestFlagID 36
	EnableSelfVersusThirdPartyZones = "enable_self_versus_third_party_zones"
	// TestFlagID 20
	EnableTCPProtocol = "enable_tcp_protocol"
	// TestFlagID 27
	F40xOr50xHTTPMarkSuccessful = "f40x_or_50x_http_mark_successful"
	// TestFlagID 31
	FavorFastestRoundTripNameserver = "favor_fastest_round_trip_nameserver"
	// TestFlagID 23
	HostDataCollectionEnabled = "host_data_collection_enabled"
	// TestFlagID 17
	IgnoreSSLFailures = "ignore_ssl_failures"
	// TestFlagID 25
	StopTestOnDocumentComplete = "stop_test_on_document_complete"
	// TestFlagID 39
	StopTestOnDOMContentLoad = "stop_test_on_dom_content_load"
	// TestFlagID 33
	T30xRedirectsDoNotFollow = "t30x_redirects_do_not_follow"
	// TestFlagID 26
	TryNextNameserverOnFailure = "try_next_nameserver_on_failure"
	// TestFlagID 2
	VerifyTestOnFailure = "verify_test_on_failure"
	// TestFlagID 24
	ZoneDataCollectionEnabled = "zone_data_collection_enabled"
)

// BandwidthThrottling Types (Advanced setting).
const (
	GPRS      = "gprs"
	Regular2G = "regular 2g"
	Good2G    = "good 2g"
	Regular3G = "regular 3g"
	Good3G    = "good 3g"
	Regular4G = "regular 4g"
	DSL       = "dsl"
	WIFI      = "wifi"
)
