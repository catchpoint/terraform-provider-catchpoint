package models

type GetTestDataJSON struct {
	Tests []TestJSON `json:"tests"`
}

type GetTestResponse struct {
	ResponseData GetTestDataJSON `json:"data"`
	CommonResponse
}

type DeleteTestResponse struct {
	ResponseData DeleteDataJSON `json:"data"`
	CommonResponse
}

type LabelsJSON struct {
	Color  string   `json:"color"`
	Name   string   `json:"name"`
	Values []string `json:"values"`
}

type TestThresholdsJSON struct {
	TestTimeApdexThresholdWarning      *float64 `json:"testTimeApdexThresholdWarning,omitempty"`
	TestTimeApdexThresholdCritical     *float64 `json:"testTimeApdexThresholdCritical,omitempty"`
	AvailabilityApdexThresholdWarning  *float64 `json:"availabilityApdexThresholdWarning,omitempty"`
	AvailabilityApdexThresholdCritical *float64 `json:"availabilityApdexThresholdCritical,omitempty"`
}

type ChromeMonitorVersionStructJSON struct {
	ApplicationVersionType *GenericIDNameOmitEmptyJSON `json:"applicationVersionType,omitempty"`
	ApplicationVersionID   *int                        `json:"applicationVersionId,omitempty"`
}

type TestRequestDataJSON struct {
	TestID                *int                        `json:"testID,omitempty"`
	RequestData           *string                     `json:"requestData,omitempty"`
	TransactionScriptType *GenericIDNameOmitEmptyJSON `json:"transactionScriptType,omitempty"`
	TestType              *GenericIDNameOmitEmptyJSON `json:"testType,omitempty"`
	Monitor               *GenericIDNameOmitEmptyJSON `json:"monitor,omitempty"`
}

type TestJSON struct {
	// These fields are always returned by the API for a Test.
	AlertsPaused          bool              `json:"alertsPaused"`
	ChangeDate            string            `json:"changeDate"`
	DivisionID            int               `json:"divisionId"`
	EnableTestDataWebhook bool              `json:"enableTestDataWebhook"`
	ID                    int               `json:"id"`
	Monitor               GenericIDNameJSON `json:"monitor"`
	Name                  string            `json:"name"`
	ProductID             int               `json:"productId"`
	StartTime             string            `json:"startTime"`
	Status                GenericIDNameJSON `json:"status"`
	TestType              GenericIDNameJSON `json:"testType"`

	// Depending on the test, these fields may or may not be returned by the API.
	// These are over-rides that should be set to inherit if nil.
	// They may also exist and already be set to inherit.
	AlertGroup       *AlertGroupJSON       `json:"alertGroup"`
	AdvancedSettings *AdvancedSettingsJSON `json:"advancedSettings,omitempty"`
	InsightData      *InsightDataJSON      `json:"insightData,omitempty"`
	RequestSettings  *RequestSettingsJSON  `json:"requestSettings,omitempty"`
	ScheduleSettings *ScheduleSettingsJSON `json:"scheduleSettings,omitempty"`

	ApplicationVersion         *string                         `json:"applicationVersion,omitempty"`
	CertificateName            *string                         `json:"certificateName,omitempty"`
	CertificateThumbprintValue *string                         `json:"certificateThumbprintValue,omitempty"`
	ChromeMonitorVersion       *ChromeMonitorVersionStructJSON `json:"chromeMonitorVersion,omitempty"`
	Description                *string                         `json:"description,omitempty"`
	DNSQueryType               *GenericIDNameOmitEmptyJSON     `json:"dnsQueryType,omitempty"`
	DNSServer                  *string                         `json:"dnsServer,omitempty"`
	// NOTE: EndTime being required or not is a client-level setting. Thus, we need to assume that it may not be set.
	EndTime                      *string                     `json:"endTime,omitempty"`
	EnforceCertificateKeyPinning *bool                       `json:"enforceCertificateKeyPinning,omitempty"`
	EnforceCertificatePinning    *bool                       `json:"enforceCertificatePinning,omitempty"`
	FileData                     *string                     `json:"fileData,omitempty"`
	FolderID                     *int                        `json:"folderId,omitempty"`
	GatewayAddressOrHost         *string                     `json:"gatewayAddressOrHost,omitempty"`
	Labels                       *[]LabelsJSON               `json:"labels,omitempty"`
	PassPhrase                   *string                     `json:"passPhrase,omitempty"`
	PublicKeyThumbprintValue     *string                     `json:"publicKeyThumbprintValue,omitempty"`
	RequestHTTPMethod            *GenericIDNameJSON          `json:"requestHTTPMethod,omitempty"`
	TestRequestData              *TestRequestDataJSON        `json:"testRequestData,omitempty"`
	TestThresholds               *TestThresholdsJSON         `json:"thresholdRestModel,omitempty"`
	UserAgentType                *GenericIDNameOmitEmptyJSON `json:"userAgentTypeId,omitempty"`
	URL                          *string                     `json:"url,omitempty"`
}
