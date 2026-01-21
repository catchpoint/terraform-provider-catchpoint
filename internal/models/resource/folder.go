package resource

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type FolderResourceModel struct {
	ID            types.Int64         `tfsdk:"id"`
	DivisionID    types.Int64         `tfsdk:"division_id"`
	ProductID     types.Int64         `tfsdk:"product_id"`
	ParentID      types.Int64         `tfsdk:"parent_id"`
	FolderName    types.String        `tfsdk:"folder_name"`
	AlertSettings *AlertSettingsModel `tfsdk:"alert_settings"`
	// Note: as dynamic types.Objects, these cannot be pointers or hold complex objects right now.
	AdvancedSettingsModel
	RequestSettingsModel
	ScheduleSettingsModel
	InsightSettingsModel
}

func (m *FolderResourceModel) IsNull() bool {
	return m == nil
}

// GetAdvancedSettingsModel returns a pointer to the embedded AdvancedSettingsModel,
// which contains advanced configuration options for the Web test.
func (m *FolderResourceModel) GetAdvancedSettingsModel() *AdvancedSettingsModel {
	return &m.AdvancedSettingsModel
}

// SetAdvancedSettingsModel sets the embedded AdvancedSettingsModel with the provided model.
func (m *FolderResourceModel) SetAdvancedSettingsModel(model *AdvancedSettingsModel) {
	m.AdvancedSettingsModel = *model
}

// GetAlertSettingsModel returns a pointer to the embedded AlertSettingsModel,
// which contains alert configuration for the Web test.
func (m *FolderResourceModel) GetAlertSettingsModel() *AlertSettingsModel {
	return m.AlertSettings
}

// SetAlertSettingsModel sets the embedded AlertSettingsModel with the provided model.
func (m *FolderResourceModel) SetAlertSettingsModel(model *AlertSettingsModel) {
	m.AlertSettings = model
}

// GetRequestSettingsModel returns a pointer to the embedded RequestSettingsModel,
// which contains request configuration for the Web test.
func (m *FolderResourceModel) GetRequestSettingsModel() *RequestSettingsModel {
	return &m.RequestSettingsModel
}

// SetRequestSettingsModel sets the embedded RequestSettingsModel with the provided model.
func (m *FolderResourceModel) SetRequestSettingsModel(model *RequestSettingsModel) {
	m.RequestSettingsModel = *model
}

// GetScheduleSettingsModel returns a pointer to the embedded ScheduleSettingsModel,
// which contains scheduling configuration for the Web test.
func (m *FolderResourceModel) GetScheduleSettingsModel() *ScheduleSettingsModel {
	return &m.ScheduleSettingsModel
}

// SetScheduleSettingsModel sets the embedded ScheduleSettingsModel with the provided model.
func (m *FolderResourceModel) SetScheduleSettingsModel(model *ScheduleSettingsModel) {
	m.ScheduleSettingsModel = *model
}

// GetInsightSettingsModel returns a pointer to the embedded InsightSettingsModel,
// which contains insight configuration for the Web test.
func (m *FolderResourceModel) GetInsightSettingsModel() *InsightSettingsModel {
	return &m.InsightSettingsModel
}

// SetInsightSettingsModel sets the embedded InsightSettingsModel with the provided model.
func (m *FolderResourceModel) SetInsightSettingsModel(model *InsightSettingsModel) {
	m.InsightSettingsModel = *model
}
