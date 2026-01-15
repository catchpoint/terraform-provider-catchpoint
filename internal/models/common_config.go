package models

// IDNameConfig represents a generic object with an integer ID and a name.
// Used throughout the codebase for backend objects that have both an ID and a Name field.
type IDNameConfig struct {
	ID   int
	Name string
}

// TestHTTPHeaderRequestConfig defines the configuration for a test HTTP header request,
// including the host pattern, header name, header type, and request value.
type TestHTTPHeaderRequestConfig struct {
	ChildHostPattern  string
	HeaderName        string
	RequestHeaderType IDNameConfig
	RequestValue      string
}

// NotificationGroupConfig represents a group of notification recipients and settings
// for alerting, including email addresses, contact groups, webhooks, and notification triggers.
type NotificationGroupConfig struct {
	Subject          string
	Emails           []string
	ContactGroupIDs  []int
	WebhookIDs       []int
	NotifyOnCritical bool
	NotifyOnImproved bool
	NotifyOnWarning  bool
}

// AlertSettingsConfig contains alerting configuration, including the alert setting type,
// associated alert rules, and the notification group to use for alerts.
type AlertSettingsConfig struct {
	AlertSettingType  IDNameConfig
	AlertRules        []AlertRuleConfig
	NotificationGroup NotificationGroupConfig
}

// NodeThresholdConfig defines threshold settings for nodes, including run counts,
// threshold types, and criteria for triggering alerts based on node performance.
type NodeThresholdConfig struct {
	ConsecutiveRuns                 int
	ConsecutiveRunsEnabled          bool
	Name                            string
	NodeThresholdType               IDNameConfig
	NumberOfUnits                   int // Number of Runs AND Number of Nodes
	NumberOfFailingUnits            int
	PercentageOfUnits               float64 // Percentage of Runs AND Percentage of Nodes
	UtilizePerNodeHistoricalAverage bool
}

// TriggerConfig specifies the configuration for alert triggers, including thresholds,
// DNS settings, filter types, operation types, and reminder frequencies.
type TriggerConfig struct {
	CriticalReminderFrequency IDNameConfig
	CriticalTrigger           float64
	DNSRecordType             IDNameConfig
	DNSResolvedName           string
	DNSTTL                    int
	Expression                string
	FilterType                IDNameConfig
	FilterValue               string
	HistoricalInterval        IDNameConfig
	Monitor                   string // Not to be confused with the test "Monitor" field. This is typically "Synthetic".
	OperationType             IDNameConfig
	StatisticalType           IDNameConfig
	ThresholdInterval         IDNameConfig
	TriggerType               IDNameConfig
	UseIntervalRollingWindow  bool
	WarningReminderFrequency  IDNameConfig
	WarningTrigger            float64
}

// AlertRuleConfig defines a single alert rule, including its type, thresholds,
// notification groups, and trigger configuration.
type AlertRuleConfig struct {
	AlertSubType        IDNameConfig
	AlertType           IDNameConfig
	EnforceTestFailure  bool
	MatchAllRecords     bool
	NodeThresholdConfig NodeThresholdConfig
	NotificationGroups  []NotificationGroupConfig
	NotificationType    IDNameConfig
	OmitScatterplot     bool
	Trigger             TriggerConfig
}

// RequestSettingsConfig contains configuration for request settings, such as
// authentication type, certificate IDs, and HTTP header requests.
type RequestSettingsConfig struct {
	RequestSettingType     IDNameConfig
	CertificateIDs         []int
	PasswordIDs            []int
	TokenIDs               []int
	AuthenticationType     IDNameConfig
	TestHTTPHeaderRequests []TestHTTPHeaderRequestConfig
}

// ScheduleSettingsConfig defines scheduling settings for tests or jobs,
// including schedule type, frequency, node distribution, and node groups.
type ScheduleSettingsConfig struct {
	ScheduleSettingType   IDNameConfig
	RunScheduleID         int
	MaintenanceScheduleID int
	Frequency             IDNameConfig
	NodeDistribution      IDNameConfig
	NodeIDs               []int
	NodeGroupIDs          []IDNameConfig
	NoOfSubsetNodes       int
}

// InsightSettingsConfig contains configuration for insights, such as
// the type of insight and associated tracepoint and indicator IDs.
type InsightSettingsConfig struct {
	InsightSettingType IDNameConfig
	TracepointIDs      []int
	IndicatorIDs       []int
}

// AdvancedSettingsConfig holds advanced configuration options, including
// monitor types, bandwidth throttling, viewport settings, and other advanced flags.
type AdvancedSettingsConfig struct {
	AdvancedSettingType            IDNameConfig
	AdditionalMonitorType          IDNameConfig
	AppliedTestFlags               []int
	BandwidthThrottling            IDNameConfig
	EDNSSubnet                     string
	MaxStepRuntimeSecOverride      int
	TracerouteFailureHopCount      int
	TraceroutePingCount            int
	VerifyTestOnFailure            bool
	ViewportHeight                 int
	ViewportWidth                  int
	WaitForNoActivityOnDocComplete *int // Pointer to allow nil values, indicating no wait time set, since 0 is a valid value.
}

// CommonConfig contains fields that are common across different configurations like Test, Product, and Folder.
type CommonConfig struct {
	DivisionID             int
	AlertSettingsConfig    AlertSettingsConfig
	RequestSettingsConfig  RequestSettingsConfig
	AdvancedSettingsConfig AdvancedSettingsConfig
	ScheduleSettingsConfig ScheduleSettingsConfig
	InsightSettingsConfig  InsightSettingsConfig
}

// ConfigUpdate defines an interface for updating configuration sections and retrieving updated values.
type ConfigUpdate interface {
	SetInsightSection([]map[string]int)
	SetScheduleSection(any)
	SetSectionToUpdate(string)
	GetSectionToUpdate() string
	GetUpdatedFieldValue() string
	GetUpdatedAdvancedSettingsSection() AdvancedSettingsJSON
	GetUpdatedRequestSettingsSection() RequestSettingsJSON
	GetUpdatedScheduleSettingsSection() any
	GetUpdatedInsightSettingsSection() []map[string]int
	GetUpdatedAlertSettingsSection() AlertGroupJSON
	GetUpdatedLabels() []LabelsJSON
	GetUpdatedTestThresholds() TestThresholdsJSON
	GetUpdatedTestRequestData() TestRequestDataJSON
	GetUpdatedChromeVersionSection() ChromeMonitorVersionStructJSON
}
