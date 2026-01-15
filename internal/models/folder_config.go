package models

type FolderConfig struct {
	CommonConfig
	FolderName string
	ParentID   int
	ProductID  int
}

type FolderConfigUpdate struct {
	UpdatedFieldValue              string
	UpdatedAdvancedSettingsSection AdvancedSettingsJSON
	UpdatedRequestSettingsSection  RequestSettingsJSON
	UpdatedScheduleSettingsSection any
	UpdatedInsightSettingsSection  []map[string]int
	UpdatedAlertSettingsSection    AlertGroupJSON
	SectionToUpdate                string
}

func (c *FolderConfigUpdate) SetInsightSection(val []map[string]int) {
	c.UpdatedInsightSettingsSection = val
}
func (c *FolderConfigUpdate) SetScheduleSection(val any) {
	c.UpdatedScheduleSettingsSection = val
}
func (c *FolderConfigUpdate) SetSectionToUpdate(val string) {
	c.SectionToUpdate = val
}
func (c *FolderConfigUpdate) GetSectionToUpdate() string {
	return c.SectionToUpdate
}
func (c *FolderConfigUpdate) GetUpdatedFieldValue() string {
	return c.UpdatedFieldValue
}
func (c *FolderConfigUpdate) GetUpdatedAdvancedSettingsSection() AdvancedSettingsJSON {
	return c.UpdatedAdvancedSettingsSection
}
func (c *FolderConfigUpdate) GetUpdatedRequestSettingsSection() RequestSettingsJSON {
	return c.UpdatedRequestSettingsSection
}
func (c *FolderConfigUpdate) GetUpdatedScheduleSettingsSection() any {
	return c.UpdatedScheduleSettingsSection
}
func (c *FolderConfigUpdate) GetUpdatedInsightSettingsSection() []map[string]int {
	return c.UpdatedInsightSettingsSection
}
func (c *FolderConfigUpdate) GetUpdatedAlertSettingsSection() AlertGroupJSON {
	return c.UpdatedAlertSettingsSection
}
func (c *FolderConfigUpdate) GetUpdatedLabels() []LabelsJSON {
	return nil // Not a FolderConfigUpdate field
}
func (c *FolderConfigUpdate) GetUpdatedTestThresholds() TestThresholdsJSON {
	return TestThresholdsJSON{} // Not a FolderConfigUpdate field
}
func (c *FolderConfigUpdate) GetUpdatedTestRequestData() TestRequestDataJSON {
	return TestRequestDataJSON{} // Not a FolderConfigUpdate field
}
func (c *FolderConfigUpdate) GetUpdatedChromeVersionSection() ChromeMonitorVersionStructJSON {
	return ChromeMonitorVersionStructJSON{} // Not a FolderConfigUpdate field
}
