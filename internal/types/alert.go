package types

var (
	ValidAlertSettingTypeNames   []string
	ValidAlertSubTypeNames       []string
	ValidAlertTypeNames          []string
	ValidHistoricalIntervalNames []string
	ValidNodeThresholdTypes      []string
	ValidNotificationTypeNames   []string
	ValidOperationTypeNames      []string
	ValidReminderNames           []string
	ValidStatisticalTypeNames    []string
	ValidThresholdIntervalNames  []string
	ValidTriggerTypeNames        []string
	ValidFilterTypes             []string
	alertSettingTypeNames        map[int]string
	alertSubTypeNames            map[int]string
	alertTypeNames               map[int]string
	historicalIntervalNames      map[int]string
	notificationTypeNames        map[int]string
	operationTypeNames           map[int]string
	reminderNames                map[int]string
	statisticalTypeNames         map[int]string
	thresholdIntervalNames       map[int]string
	triggerTypeNames             map[int]string
	filterTypeNames              map[int]string
)

// #region NodeThresholdTypes

var NodeThresholdTypeIDs = map[string]int{
	Runs:               0,
	AverageAcrossNodes: 1,
	Node:               2,
}

func GetNodeThresholdTypeID(name string) (int, bool) {
	id, ok := NodeThresholdTypeIDs[name]
	return id, ok
}

func GetNodeThresholdTypeName(id int) (string, bool) {
	name, ok := nodeThresholdTypeNames[id]
	return name, ok
}

// #endregion

// #region NotificationTypes

var NotificationTypeIDs = map[string]int{
	DefaultContacts: 0,
	// TODO: add "AdditionalContacts" id 3 when supported.
}

func GetNotificationTypeID(name string) (int, bool) {
	id, ok := NotificationTypeIDs[name]
	return id, ok
}

func GetNotificationTypeName(id int) (string, bool) {
	name, ok := notificationTypeNames[id]
	return name, ok
}

// #endregion

// #region AlertTypes

var AlertTypeIDs = map[string]int{
	ByteLength:      2,
	ContentMatch:    3,
	HostFailure:     4,
	TestFailure:     9,
	Timing:          7,
	Insight:         10, // Custom-Test-Data
	Ping:            12,
	Requests:        13,
	Availability:    15,
	DNS:             17,
	Path:            20,
	ASN:             23,
	ExperienceScore: 26, // Synthetic-Experience-Score
}

func GetAlertTypeID(alertType string) (int, bool) {
	id, ok := AlertTypeIDs[alertType]
	return id, ok
}

func GetAlertTypeName(id int) (string, bool) {
	name, ok := alertTypeNames[id]
	return name, ok
}

// #endregion

// #region AlertSubTypes

var AlertSubTypeIDs = map[string]int{
	ByteLength:          1,
	Page:                2,
	FileSize:            3,
	RegularExpression:   10,
	CompareToPrevious:   11,
	CompareToFile:       12,
	ResponseCode:        14,
	ResponseHeaders:     15,
	DNS:                 50,
	Connect:             51,
	Send:                52,
	Wait:                53,
	Load:                54,
	TTFB:                55,
	ContentLoad:         57,
	Response:            58,
	TestTime:            59,
	DOMLoad:             61,
	TestTimeWithSuspect: 63,
	ServerResponse:      64,
	DocumentComplete:    66,
	Redirect:            67,
	Tracepoints:         90, // Custom-Test-Data-Dimension
	Indicators:          91, // Custom-Test-Data-Metric
	PingRTT:             100,
	PingPacketLoss:      101,
	RequestsNum:         110,
	HostsNum:            111,
	ConnectionsNum:      112,
	RedirectsNum:        113,
	OtherNum:            114,
	ImagesNum:           115,
	ScriptsNum:          116,
	HTMLNum:             117,
	CSSNum:              118,
	XMLNum:              119,
	FlashNum:            120,
	MediaNum:            121,
	Test:                140,
	Content:             141,
	DowntimePercent:     142,
	Reachability:        143, // Reachability-Bgp
	TestURLs:            150, // Address-Test-Url
	ChildURLs:           151, // Address-Child
	AnyURLs:             152, // Address-Page
	DNSGeneral:          160, // DNS-General
	DNSAnswer:           161, // DNS-Answer
	DNSAuthority:        162, // DNS-Authority
	DNSAdditional:       163, // DNS-Additional
	CitiesNum:           190,
	ASNsNum:             191,
	CountriesNum:        193,
	HopsNum:             194,
	HandshakeTime:       195,
	DaysToExpiration:    196,
	OriginAS:            210,
	PathAS:              211,
	OriginNeighbor:      212,
	PrefixMismatch:      213,
}

func GetAlertSubTypeID(alertType string) (int, bool) {
	id, ok := AlertSubTypeIDs[alertType]
	return id, ok
}

func GetAlertSubTypeName(id int) (string, bool) {
	name, ok := alertSubTypeNames[id]
	return name, ok
}

// #endregion

// #region AlertSettingTypes

var AlertSettingTypeIDs = map[string]int{
	Inherit:       0,
	Override:      1,
	InheritAndAdd: 2,
}

func GetAlertSettingTypeID(alertType string) (int, bool) {
	id, ok := AlertSettingTypeIDs[alertType]
	return id, ok
}

func GetAlertSettingTypeName(id int) (string, bool) {
	name, ok := alertSettingTypeNames[id]
	return name, ok
}

// #endregion

// #region OperationTypes

var OperationTypeIDs = map[string]int{
	Equals:              1,
	NotEquals:           0,
	GreaterThan:         2,
	GreaterThanOrEquals: 3,
	LessThan:            4,
	LessThanOrEquals:    5,
}

func GetOperationTypeID(name string) (int, bool) {
	id, ok := OperationTypeIDs[name]
	return id, ok
}

func GetOperationTypeName(id int) (string, bool) {
	name, ok := operationTypeNames[id]
	return name, ok
}

// #endregion

// #region TriggerTypes

var TriggerTypeIDs = map[string]int{
	SpecificValue: 1,
	TrailingValue: 2,
	TrendShift:    3,
}

func GetTriggerTypeID(triggerType string) (int, bool) {
	id, ok := TriggerTypeIDs[triggerType]
	return id, ok
}

func GetTriggerTypeName(triggerType int) (string, bool) {
	name, ok := triggerTypeNames[triggerType]
	return name, ok
}

// #endregion

// #region ReminderTypes

var ReminderIDs = map[string]int{
	None:                   0,
	string(OneMinute):      1,
	string(FiveMinutes):    5,
	string(TenMinutes):     10,
	string(FifteenMinutes): 15,
	string(ThirtyMinutes):  30,
	string(OneHour):        60,
	string(Daily):          1440,
}

func GetReminderID(reminder string) (int, bool) {
	id, ok := ReminderIDs[reminder]
	return id, ok
}

func GetReminderName(reminder int) (string, bool) {
	name, ok := reminderNames[reminder]
	return name, ok
}

// #endregion

// #region ThresholdIntervalTypes

var ThresholdIntervalIDs = map[string]int{
	Default:                0,
	string(FiveMinutes):    5,
	string(TenMinutes):     10,
	string(FifteenMinutes): 15,
	string(ThirtyMinutes):  30,
	string(OneHour):        60,
	string(TwoHours):       120,
	string(SixHours):       360,
	string(TwelveHours):    720,
}

func GetThresholdIntervalID(thresholdInterval string) (int, bool) {
	id, ok := ThresholdIntervalIDs[thresholdInterval]
	return id, ok
}

func GetThresholdIntervalName(thresholdInterval int) (string, bool) {
	name, ok := thresholdIntervalNames[thresholdInterval]
	return name, ok
}

// #endregion

// #region HistoricalIntervalTypes

var HistoricalIntervalIDs = map[string]int{
	string(FiveMinutes):    5,
	string(TenMinutes):     10,
	string(FifteenMinutes): 15,
	string(ThirtyMinutes):  30,
	string(OneHour):        60,
	string(TwoHours):       120,
	string(SixHours):       360,
	string(TwelveHours):    720,
	string(OneDay):         1440,
	string(OneWeek):        10080,
}

func GetHistoricalIntervalID(historicalInterval string) (int, bool) {
	id, ok := HistoricalIntervalIDs[historicalInterval]
	return id, ok
}

func GetHistoricalIntervalName(historicalInterval int) (string, bool) {
	name, ok := historicalIntervalNames[historicalInterval]
	return name, ok
}

// #endregion

// #region StatisticalTypes

var StatisticalTypeIDs = map[string]int{
	Average: 1,
}

func GetStatisticalTypeID(statisticalType string) (int, bool) {
	id, ok := StatisticalTypeIDs[statisticalType]
	return id, ok
}

func GetStatisticalTypeName(statisticalType int) (string, bool) {
	name, ok := statisticalTypeNames[statisticalType]
	return name, ok
}

// #endregion

// #region FilterTypes

// Note: There's an additional filterType in the FE called "Step: All" that has no filterType or filterValue.
//
//	"fileName": null,
//	"filterItems": [],
//	"filterType": null,
//	"filterValue": null,
//	"historicalInterval": null,
//	"operationType": 2,
//	"statisticalType": null,
//	"thresholdInterval": null,
//	"triggerType": 1,
var FilterTypeIDs = map[string]int{
	Index: 1, // Index is index id 1 with a filterValue giving the specific index.
	//Last:    1, // Last is index id 1 with filterValue "".
	Name:    2, // Step Name or Regex.
	Address: 3, // Address is RootIP with filterValue being an address.
}

func GetFilterTypeID(name string) (int, bool) {
	id, ok := FilterTypeIDs[name]
	return id, ok
}

func GetFilterTypeName(id int) (string, bool) {
	name, ok := filterTypeNames[id]
	return name, ok
}

// #endregion

// On init, generate reverse mappings for later O(1) lookups.
func init() {
	notificationTypeNames = populateReverseMap(NotificationTypeIDs)
	ValidNotificationTypeNames = populateValidNames(NotificationTypeIDs)

	alertTypeNames = populateReverseMap(AlertTypeIDs)
	ValidAlertTypeNames = populateValidNames(AlertTypeIDs)

	alertSubTypeNames = populateReverseMap(AlertSubTypeIDs)
	ValidAlertSubTypeNames = populateValidNames(AlertSubTypeIDs)

	alertSettingTypeNames = populateReverseMap(AlertSettingTypeIDs)
	ValidAlertSettingTypeNames = populateValidNames(AlertSettingTypeIDs)

	operationTypeNames = populateReverseMap(OperationTypeIDs)
	ValidOperationTypeNames = populateValidNames(OperationTypeIDs)

	triggerTypeNames = populateReverseMap(TriggerTypeIDs)
	ValidTriggerTypeNames = populateValidNames(TriggerTypeIDs)

	reminderNames = populateReverseMap(ReminderIDs)
	ValidReminderNames = populateValidNames(ReminderIDs)

	thresholdIntervalNames = populateReverseMap(ThresholdIntervalIDs)
	ValidThresholdIntervalNames = populateValidNames(ThresholdIntervalIDs)

	historicalIntervalNames = populateReverseMap(HistoricalIntervalIDs)
	ValidHistoricalIntervalNames = populateValidNames(HistoricalIntervalIDs)

	statisticalTypeNames = populateReverseMap(StatisticalTypeIDs)
	ValidStatisticalTypeNames = populateValidNames(StatisticalTypeIDs)

	nodeThresholdTypeNames = populateReverseMap(NodeThresholdTypeIDs)
	ValidNodeThresholdTypes = populateValidNames(NodeThresholdTypeIDs)

	filterTypeNames = populateReverseMap(FilterTypeIDs)
	ValidFilterTypes = populateValidNames(FilterTypeIDs)
}
