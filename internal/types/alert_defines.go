package types

// Trigger Types.
const (
	// TriggerTypeID 1
	SpecificValue = "specific value"
	// TriggerTypeID 2
	TrailingValue = "trailing value"
	// TriggerTypeID 3
	TrendShift = "trendshift"
)

// Operation Types.
const (
	// OperationTypeID 0
	NotEquals = "not equals"
	// OperationTypeID 1
	Equals = "equals"
	// OperationTypeID 2
	GreaterThan = "greater than"
	// OperationTypeID 3
	GreaterThanOrEquals = "greater than or equals"
	// OperationTypeID 4
	LessThan = "less than"
	// OperationTypeID 5
	LessThanOrEquals = "less than or equals"
)

// AlertTypes.
const (
	// AlertTypeID 2, AlertSubTypeID 1
	ByteLength = "byte length"
	// AlertTypeID 3
	ContentMatch = "content match"
	// AlertTypeID 4
	HostFailure = "host failure"
	// AlertTypeID 9
	TestFailure = "test failure"
	// AlertTypeID 7, AKA 'Response Time'
	Timing = "timing"
	// AlertTypeID 10, AKA 'Custom-Test-Data'
	Insight = "insight"
	// AlertTypeID 12
	Ping = "ping"
	// AlertTypeID 13
	Requests = "requests"
	// AlertTypeID 15
	Availability = "availability"
	// AlertTypeID 17
	DNS = "dns"
	// AlertTypeID 20
	Path = "path"
	// AlertTypeID 23
	ASN = "asn"
	// AlertTypeID 26
	ExperienceScore = "experience score"
)

// AlertSubTypes.
const (
	Page                = "page"
	FileSize            = "file size"
	RegularExpression   = "regular expression"
	CompareToPrevious   = "compare to previous"
	CompareToFile       = "compare to file"
	ResponseCode        = "response code"
	ResponseHeaders     = "response headers"
	Connect             = "connect"
	Send                = "send"
	Wait                = "wait"
	Load                = "load"
	TTFB                = "ttfb"
	ContentLoad         = "content load"
	Response            = "response"
	TestTime            = "test time"
	DOMLoad             = "dom load"
	TestTimeWithSuspect = "test time with suspect"
	ServerResponse      = "server response"
	DocumentComplete    = "document complete"
	Redirect            = "redirect"
	PingRTT             = "ping rtt"
	PingPacketLoss      = "ping packet loss"
	RequestsNum         = "# requests"
	HostsNum            = "# hosts"
	ConnectionsNum      = "# connections"
	RedirectsNum        = "# redirects"
	OtherNum            = "# other"
	ImagesNum           = "# images"
	ScriptsNum          = "# scripts"
	HTMLNum             = "# html"
	CSSNum              = "# css"
	XMLNum              = "# xml"
	FlashNum            = "# flash"
	MediaNum            = "# media"
	Test                = "test"
	Content             = "content"
	DowntimePercent     = "% downtime"
	DNSGeneral          = "dns general"
	DNSAnswer           = "dns answer"
	DNSAuthority        = "dns authority"
	DNSAdditional       = "dns additional"
	CitiesNum           = "# cities"
	ASNsNum             = "# asns"
	CountriesNum        = "# countries"
	HopsNum             = "# hops"
	HandshakeTime       = "handshake time"
	DaysToExpiration    = "days to expiration"
	OriginAS            = "origin as"
	PathAS              = "path as"
	OriginNeighbor      = "origin neighbor"
	PrefixMismatch      = "prefix mismatch"
	Reachability        = "reachability"
	TestURLs            = "test urls"
	AnyURLs             = "any urls"
	ChildURLs           = "child urls"
	Tracepoints         = "tracepoints"
	Indicators          = "indicators"
)

// FilterTypes.
const (
	Index   = "index"
	Last    = "last"
	Name    = "name"
	Address = "address"
)

// Node Thresholds.
const (
	// NodeThresholdTypeID 0
	Runs = "runs"
	// NodeThresholdTypeID 1
	AverageAcrossNodes = "average across nodes"
	// NodeThresholdTypeID 2
	Node = "node"
)
