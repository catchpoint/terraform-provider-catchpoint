package models

type TestLabel struct {
	Name   string
	Values []string
}

type TestRequestData struct {
	TestID                int
	RequestData           string
	TransactionScriptType IDNameConfig
	TestType              IDNameConfig
	Monitor               IDNameConfig
}

type ChromeMonitorVersion struct {
	ApplicationVersionType IDNameConfig
	ApplicationVersionID   int
}

type TestThresholds struct {
	TestTimeThresholdWarning      float64
	TestTimeThresholdCritical     float64
	AvailabilityThresholdWarning  float64
	AvailabilityThresholdCritical float64
}

type TestConfig struct {
	CommonConfig
	AlertsPaused                 bool
	TestThresholds               TestThresholds
	CertificateName              string
	ChromeApplicationVersion     ChromeMonitorVersion
	DNSQueryType                 IDNameConfig
	DNSServer                    string
	EnableTestDataWebhook        bool
	EndTime                      string
	EnforceCertificateKeyPinning bool
	EnforceCertificatePinning    bool
	FileData                     string
	FolderID                     int
	GatewayAddressOrHost         string
	Labels                       []TestLabel
	Monitor                      IDNameConfig
	Passphrase                   string
	ProductID                    int
	Script                       TestRequestData
	SimulateDevice               IDNameConfig
	StartTime                    string
	TestDescription              string
	TestName                     string
	Status                       IDNameConfig
	TestType                     IDNameConfig
	TestURL                      string
}

type TestConfigUpdate struct {
	UpdatedFieldValue              string
	UpdatedAdvancedSettingsSection AdvancedSettingsJSON
	UpdatedRequestSettingsSection  RequestSettingsJSON
	UpdatedScheduleSettingsSection any
	UpdatedInsightSettingsSection  []map[string]int
	UpdatedAlertSettingsSection    AlertGroupJSON
	UpdatedLabels                  []LabelsJSON
	UpdatedTestThresholds          TestThresholdsJSON
	UpdatedTestRequestData         TestRequestDataJSON
	UpdatedChromeVersionSection    ChromeMonitorVersionStructJSON
	SectionToUpdate                string
}

func (c *TestConfigUpdate) SetInsightSection(val []map[string]int) {
	c.UpdatedInsightSettingsSection = val
}
func (c *TestConfigUpdate) SetScheduleSection(val any) {
	c.UpdatedScheduleSettingsSection = val
}
func (c *TestConfigUpdate) SetSectionToUpdate(val string) {
	c.SectionToUpdate = val
}
func (c *TestConfigUpdate) GetSectionToUpdate() string {
	return c.SectionToUpdate
}
func (c *TestConfigUpdate) GetUpdatedFieldValue() string {
	return c.UpdatedFieldValue
}
func (c *TestConfigUpdate) GetUpdatedAdvancedSettingsSection() AdvancedSettingsJSON {
	return c.UpdatedAdvancedSettingsSection
}
func (c *TestConfigUpdate) GetUpdatedRequestSettingsSection() RequestSettingsJSON {
	return c.UpdatedRequestSettingsSection
}
func (c *TestConfigUpdate) GetUpdatedScheduleSettingsSection() any {
	return c.UpdatedScheduleSettingsSection
}
func (c *TestConfigUpdate) GetUpdatedInsightSettingsSection() []map[string]int {
	return c.UpdatedInsightSettingsSection
}
func (c *TestConfigUpdate) GetUpdatedAlertSettingsSection() AlertGroupJSON {
	return c.UpdatedAlertSettingsSection
}
func (c *TestConfigUpdate) GetUpdatedLabels() []LabelsJSON {
	return c.UpdatedLabels
}
func (c *TestConfigUpdate) GetUpdatedTestThresholds() TestThresholdsJSON {
	return c.UpdatedTestThresholds
}
func (c *TestConfigUpdate) GetUpdatedTestRequestData() TestRequestDataJSON {
	return c.UpdatedTestRequestData
}
func (c *TestConfigUpdate) GetUpdatedChromeVersionSection() ChromeMonitorVersionStructJSON {
	return c.UpdatedChromeVersionSection
}
