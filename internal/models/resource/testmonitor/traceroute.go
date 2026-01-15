package testmonitor

import (
	"catchpoint-provider/internal/models/resource"
	cptypes "catchpoint-provider/internal/types"

	"github.com/hashicorp/terraform-plugin-framework/types"
)

type TracerouteTestResourceModel struct {
	resource.BaseTestResourceModel
	resource.AdvancedSettingsModel
	resource.ScheduleSettingsModel
	AlertSettings *resource.AlertSettingsModel `tfsdk:"alert_settings"`
	Thresholds    *resource.ThresholdModel     `tfsdk:"thresholds"`
	TestLocation  types.String                 `tfsdk:"test_location"`
}

func (m *TracerouteTestResourceModel) New() resource.TestResourceModelProvider {
	return &TracerouteTestResourceModel{}
}

func (m *TracerouteTestResourceModel) IsNull() bool {
	return m == nil
}

func (m *TracerouteTestResourceModel) GetTestType() cptypes.TestType {
	return cptypes.TracerouteType
}

// GetID returns the ID for the TracerouteTestResourceModel.
func (m *TracerouteTestResourceModel) GetID() types.Int64 {
	return m.ID
}

// SetID sets the ID for the TracerouteTestResourceModel.
func (m *TracerouteTestResourceModel) SetID(id types.Int64) {
	m.ID = id
}

// GetTestResourceModel returns the TestResourceModel for Traceroute tests.
func (m *TracerouteTestResourceModel) GetTestResourceModel() *resource.BaseTestResourceModel {
	if m == nil {
		return nil
	}
	return &m.BaseTestResourceModel
}

// GetURLField for TracerouteTestResourceModel returns the Location as the URL field.
func (m *TracerouteTestResourceModel) GetURLField() (types.String, bool) {
	if m.TestLocation.IsNull() || m.TestLocation.IsUnknown() {
		return m.TestLocation, false
	}
	return m.TestLocation, true
}

// SetURLField for TracerouteTestResourceModel sets the Location as the URL field.
func (m *TracerouteTestResourceModel) SetURLField(url types.String) {
	m.TestLocation = url
}

// GetTestThresholdModel returns a pointer to the ThresholdModel for the Traceroute test,
// which contains threshold configuration for test metrics.
func (m *TracerouteTestResourceModel) GetTestThresholdModel() *resource.ThresholdModel {
	if m == nil {
		return nil
	}
	return m.Thresholds
}

// SetTestThresholdModel sets the ThresholdModel for the Traceroute test with the provided threshold.
func (m *TracerouteTestResourceModel) SetTestThresholdModel(threshold *resource.ThresholdModel) {
	m.Thresholds = threshold
}

// GetAdvancedSettingsModel returns a pointer to the embedded AdvancedSettingsModel,
// which contains advanced configuration options for the Traceroute test.
func (m *TracerouteTestResourceModel) GetAdvancedSettingsModel() *resource.AdvancedSettingsModel {
	if m == nil {
		return nil
	}
	return &m.AdvancedSettingsModel
}

// SetAdvancedSettingsModel sets the embedded AdvancedSettingsModel with the provided model.
func (m *TracerouteTestResourceModel) SetAdvancedSettingsModel(model *resource.AdvancedSettingsModel) {
	m.AdvancedSettingsModel = *model
}

// GetAlertSettingsModel returns a pointer to the embedded AlertSettingsModel,
// which contains alert configuration for the Traceroute test.
func (m *TracerouteTestResourceModel) GetAlertSettingsModel() *resource.AlertSettingsModel {
	if m == nil {
		return nil
	}
	return m.AlertSettings
}

// SetAlertSettingsModel sets the embedded AlertSettingsModel with the provided model.
func (m *TracerouteTestResourceModel) SetAlertSettingsModel(model *resource.AlertSettingsModel) {
	m.AlertSettings = model
}

// GetScheduleSettingsModel returns a pointer to the embedded ScheduleSettingsModel,
// which contains scheduling configuration for the Traceroute test.
func (m *TracerouteTestResourceModel) GetScheduleSettingsModel() *resource.ScheduleSettingsModel {
	if m == nil {
		return nil
	}
	return &m.ScheduleSettingsModel
}

// SetScheduleSettingsModel sets the embedded ScheduleSettingsModel with the provided model.
func (m *TracerouteTestResourceModel) SetScheduleSettingsModel(model *resource.ScheduleSettingsModel) {
	m.ScheduleSettingsModel = *model
}
