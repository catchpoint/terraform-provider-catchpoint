package testmonitor

import (
	"catchpoint-provider/internal/models/resource"
	cptypes "catchpoint-provider/internal/types"

	"github.com/hashicorp/terraform-plugin-framework/types"
)

type PingTestResourceModel struct {
	resource.BaseTestResourceModel
	resource.AdvancedSettingsModel
	resource.ScheduleSettingsModel
	AlertSettings *resource.AlertSettingsModel `tfsdk:"alert_settings"`
	Thresholds    *resource.ThresholdModel     `tfsdk:"thresholds"`
	TestLocation  types.String                 `tfsdk:"test_location"`
}

func (m *PingTestResourceModel) New() resource.TestResourceModelProvider {
	return &PingTestResourceModel{}
}

func (m *PingTestResourceModel) IsNull() bool {
	return m == nil
}

func (m *PingTestResourceModel) GetTestType() cptypes.TestType {
	return cptypes.PingType
}

// GetID returns the ID for the PingTestResourceModel.
func (m *PingTestResourceModel) GetID() types.Int64 {
	return m.ID
}

// SetID sets the ID for the PingTestResourceModel.
func (m *PingTestResourceModel) SetID(id types.Int64) {
	m.ID = id
}

// GetTestResourceModel returns the TestResourceModel for Ping tests.
func (m *PingTestResourceModel) GetTestResourceModel() *resource.BaseTestResourceModel {
	if m == nil {
		return nil
	}
	return &m.BaseTestResourceModel
}

// GetURLField for PingTestResourceModel returns the Location as the URL field.
func (m *PingTestResourceModel) GetURLField() (types.String, bool) {
	if m.TestLocation.IsNull() || m.TestLocation.IsUnknown() {
		return m.TestLocation, false
	}
	return m.TestLocation, true
}

// SetURLField for PingTestResourceModel sets the Location as the URL field.
func (m *PingTestResourceModel) SetURLField(url types.String) {
	m.TestLocation = url
}

// GetTestThresholdModel returns a pointer to the ThresholdModel for the Ping test,
// which contains threshold configuration for test metrics.
func (m *PingTestResourceModel) GetTestThresholdModel() *resource.ThresholdModel {
	if m == nil {
		return nil
	}
	return m.Thresholds
}

// SetTestThresholdModel sets the ThresholdModel for the Ping test with the provided threshold.
func (m *PingTestResourceModel) SetTestThresholdModel(threshold *resource.ThresholdModel) {
	m.Thresholds = threshold
}

// GetAdvancedSettingsModel returns a pointer to the embedded AdvancedSettingsModel,
// which contains advanced configuration options for the Ping test.
func (m *PingTestResourceModel) GetAdvancedSettingsModel() *resource.AdvancedSettingsModel {
	if m == nil {
		return nil
	}
	return &m.AdvancedSettingsModel
}

// SetAdvancedSettingsModel sets the embedded AdvancedSettingsModel with the provided model.
func (m *PingTestResourceModel) SetAdvancedSettingsModel(model *resource.AdvancedSettingsModel) {
	m.AdvancedSettingsModel = *model
}

// GetAlertSettingsModel returns a pointer to the embedded AlertSettingsModel,
// which contains alert configuration for the Ping test.
func (m *PingTestResourceModel) GetAlertSettingsModel() *resource.AlertSettingsModel {
	if m == nil {
		return nil
	}
	return m.AlertSettings
}

// SetAlertSettingsModel sets the embedded AlertSettingsModel with the provided model.
func (m *PingTestResourceModel) SetAlertSettingsModel(model *resource.AlertSettingsModel) {
	m.AlertSettings = model
}

// GetScheduleSettingsModel returns a pointer to the embedded ScheduleSettingsModel,
// which contains scheduling configuration for the Ping test.
func (m *PingTestResourceModel) GetScheduleSettingsModel() *resource.ScheduleSettingsModel {
	if m == nil {
		return nil
	}
	return &m.ScheduleSettingsModel
}

// SetScheduleSettingsModel sets the embedded ScheduleSettingsModel with the provided model.
func (m *PingTestResourceModel) SetScheduleSettingsModel(model *resource.ScheduleSettingsModel) {
	m.ScheduleSettingsModel = *model
}
