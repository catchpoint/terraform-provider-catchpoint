package catchpoint

type IdName struct {
	Id   int
	Name string
}

type TestConfig struct {
	TestType                       int
	Monitor                        int
	SimulateDevice                 int
	ChromeVersion                  IdName
	ChromeApplicationVersion       IdName
	Script                         TestRequestData
	DivisionId                     int
	ProductId                      int
	FolderId                       int
	TestName                       string
	DnsQueryType                   IdName
	DnsServer                      string
	EdnsSubnet                     string
	TestUrl                        string
	TestDescription                string
	GatewayAddressOrHost           string
	Labels                         []TestLabel
	TestTimeThresholdWarning       float64
	TestTimeThresholdCritical      float64
	AvailabilityThresholdWarning   float64
	AvailabilityThresholdCritical  float64
	EnforceCertificateKeyPinning   bool
	EnforceCertificatePinning      bool
	EnableTestDataWebhook          bool
	AlertsPaused                   bool
	StartTime                      string
	EndTime                        string
	TestStatus                     int
	RequestSettingType             int
	AuthenticationType             IdName
	Username                       string
	Password                       string
	AuthenticationPasswordIds      []int
	AuthenticationTokenIds         []int
	AuthenticationCertificateIds   []int
	TestHttpHeaderRequests         []TestHttpHeaderRequest
	InsightSettingType             int
	TracepointIds                  []int
	IndicatorIds                   []int
	ScheduleSettingType            int
	ScheduleRunScheduleId          int
	ScheduleMaintenanceScheduleId  int
	TestFrequency                  IdName
	NodeDistribution               IdName
	NodeIds                        []int
	NodeGroupIds                   []int
	AlertSettingType               int
	AlertRuleConfigs               []AlertRuleConfig
	AlertWebhookIds                []int
	AlertRecipientEmails           []string
	AdvancedSettingType            int
	AppliedTestFlags               []int
	MaxStepRuntimeSecOverride      int
	AdditionalMonitorType          IdName
	BandwidthThrottling            IdName
	WaitForNoActivityOnDocComplete *int
	ViewportHeight                 int
	ViewportWidth                  int
	TracerouteFailureHopCount      int
	TraceroutePingCount            int
	AlertSubject                   string
}

type TestHttpHeaderRequest struct {
	RequestHeaderType IdName
	RequestValue      string
	ChildHostPattern  string
}

type TestRequestData struct {
	TestId                int
	RequestData           string
	TransactionScriptType int
	TestType              int
	Monitor               int
}

type AlertRuleConfig struct {
	AlertNodeThresholdType          IdName
	AlertThresholdNumOfRuns         int
	AlertConsecutiveNumOfRuns       int
	AlertThresholdPercentOfRuns     float64
	AlertThresholdNumOfFailingNodes int
	TriggerType                     IdName
	OperationType                   IdName
	StatisticalType                 IdName
	TrailingHistoricalInterval      IdName
	Expression                      string
	AlertWarningTrigger             float64
	AlertCriticalTrigger            float64
	AlertEnableConsecutive          bool
	AlertWarningReminder            IdName
	AlertCriticalReminder           IdName
	AlertThresholdInterval          IdName
	AlertUseRollingWindow           bool
	AlertNotificationType           int
	AlertType                       IdName
	AlertSubType                    IdName
	AlertEnforceTestFailure         bool
	AlertOmitScatterplot            bool
	Subject                         string
	NotifyOnWarning                 bool
	NotifyOnCritical                bool
	NotifyOnImproved                bool
	AlertWebhookIds                 []int
	AlertRecipientEmails            []string
}

type TestLabel struct {
	Name   string
	Values []string
}

type TestConfigUpdate struct {
	UpdatedFieldValue              string
	UpdatedAdvancedSettingsSection AdvancedSetting
	UpdatedRequestSettingsSection  RequestSetting
	UpdatedScheduleSettingsSection ScheduleSetting
	UpdatedInsightSettingsSection  InsightDataStruct
	UpdatedAlertSettingsSection    AlertGroupStruct
	UpdatedLabels                  []Label
	UpdatedTestThresholds          Thresholds
	UpdatedTestRequestData         TestRequestDataStruct
	SectionToUpdate                string
}
