package models

type ProductConfig struct {
	CommonConfig
	AlertGroupID      int
	ProductName       string
	Status            IDNameConfig
	TestDataWebhookID int
}

type ProductConfigUpdate struct {
	UpdatedFieldValue              string
	UpdatedAdvancedSettingsSection AdvancedSettingsJSON
	UpdatedRequestSettingsSection  RequestSettingsJSON
	UpdatedScheduleSettingsSection any
	UpdatedInsightSettingsSection  []map[string]int
	UpdatedAlertSettingsSection    AlertGroupJSON
	SectionToUpdate                string
}

func (c *ProductConfigUpdate) SetInsightSection(val []map[string]int) {
	c.UpdatedInsightSettingsSection = val
}
func (c *ProductConfigUpdate) SetScheduleSection(val any) {
	c.UpdatedScheduleSettingsSection = val
}
func (c *ProductConfigUpdate) SetSectionToUpdate(val string) {
	c.SectionToUpdate = val
}
func (c *ProductConfigUpdate) GetSectionToUpdate() string {
	return c.SectionToUpdate
}
func (c *ProductConfigUpdate) GetUpdatedFieldValue() string {
	return c.UpdatedFieldValue
}
func (c *ProductConfigUpdate) GetUpdatedAdvancedSettingsSection() AdvancedSettingsJSON {
	return c.UpdatedAdvancedSettingsSection
}
func (c *ProductConfigUpdate) GetUpdatedRequestSettingsSection() RequestSettingsJSON {
	return c.UpdatedRequestSettingsSection
}
func (c *ProductConfigUpdate) GetUpdatedScheduleSettingsSection() any {
	return c.UpdatedScheduleSettingsSection
}
func (c *ProductConfigUpdate) GetUpdatedInsightSettingsSection() []map[string]int {
	return c.UpdatedInsightSettingsSection
}
func (c *ProductConfigUpdate) GetUpdatedAlertSettingsSection() AlertGroupJSON {
	return c.UpdatedAlertSettingsSection
}
func (c *ProductConfigUpdate) GetUpdatedLabels() []LabelsJSON {
	return nil // Not a ProductConfigUpdate field
}
func (c *ProductConfigUpdate) GetUpdatedTestThresholds() TestThresholdsJSON {
	return TestThresholdsJSON{} // Not a ProductConfigUpdate field
}
func (c *ProductConfigUpdate) GetUpdatedTestRequestData() TestRequestDataJSON {
	return TestRequestDataJSON{} // Not a ProductConfigUpdate field
}
func (c *ProductConfigUpdate) GetUpdatedChromeVersionSection() ChromeMonitorVersionStructJSON {
	return ChromeMonitorVersionStructJSON{} // Not a ProductConfigUpdate field
}
