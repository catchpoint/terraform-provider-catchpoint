package testmonitor

import (
	"catchpoint-provider/internal/models/resource"
	cptypes "catchpoint-provider/internal/types"

	"github.com/hashicorp/terraform-plugin-framework/types"
)

type SSLTestResourceModel struct {
	resource.SSLSettingsModel
	resource.BaseTestResourceModel
	resource.AdvancedSettingsModel
	resource.ScheduleSettingsModel
	resource.GatewayAddressOrHostModel
	AlertSettings *resource.AlertSettingsModel `tfsdk:"alert_settings"`
	Thresholds    *resource.ThresholdModel     `tfsdk:"thresholds"`
	TestLocation  types.String                 `tfsdk:"test_location"`
}

func (m *SSLTestResourceModel) New() resource.TestResourceModelProvider {
	return &SSLTestResourceModel{}
}

func (m *SSLTestResourceModel) IsNull() bool {
	return m == nil
}

func (m *SSLTestResourceModel) GetTestType() cptypes.TestType {
	return cptypes.SSLType
}

// GetID returns the ID for the SSLTestResourceModel.
func (m *SSLTestResourceModel) GetID() types.Int64 {
	return m.ID
}

// SetID sets the ID for the SSLTestResourceModel.
func (m *SSLTestResourceModel) SetID(id types.Int64) {
	m.ID = id
}

// GetTestResourceModel returns the TestResourceModel for SSL tests.
func (m *SSLTestResourceModel) GetTestResourceModel() *resource.BaseTestResourceModel {
	if m == nil {
		return nil
	}
	return &m.BaseTestResourceModel
}

// GetURLField for SSLTestResourceModel returns the Location as the URL field.
func (m *SSLTestResourceModel) GetURLField() (types.String, bool) {
	if m.TestLocation.IsNull() || m.TestLocation.IsUnknown() {
		return m.TestLocation, false
	}
	return m.TestLocation, true
}

// SetURLField for SSLTestResourceModel sets the Location as the URL field.
func (m *SSLTestResourceModel) SetURLField(url types.String) {
	m.TestLocation = url
}

// GetTestThresholdModel returns a pointer to the ThresholdModel for the SSL test,
// which contains threshold configuration for test metrics.
func (m *SSLTestResourceModel) GetTestThresholdModel() *resource.ThresholdModel {
	if m == nil {
		return nil
	}
	return m.Thresholds
}

// SetTestThresholdModel sets the ThresholdModel for the SSL test with the provided threshold.
func (m *SSLTestResourceModel) SetTestThresholdModel(threshold *resource.ThresholdModel) {
	m.Thresholds = threshold
}

// GetAdvancedSettingsModel returns a pointer to the embedded AdvancedSettingsModel,
// which contains advanced configuration options for the SSL test.
func (m *SSLTestResourceModel) GetAdvancedSettingsModel() *resource.AdvancedSettingsModel {
	if m == nil {
		return nil
	}
	return &m.AdvancedSettingsModel
}

// SetAdvancedSettingsModel sets the embedded AdvancedSettingsModel with the provided model.
func (m *SSLTestResourceModel) SetAdvancedSettingsModel(model *resource.AdvancedSettingsModel) {
	m.AdvancedSettingsModel = *model
}

// GetAlertSettingsModel returns a pointer to the embedded AlertSettingsModel,
// which contains alert configuration for the SSL test.
func (m *SSLTestResourceModel) GetAlertSettingsModel() *resource.AlertSettingsModel {
	if m == nil {
		return nil
	}
	return m.AlertSettings
}

// SetAlertSettingsModel sets the embedded AlertSettingsModel with the provided model.
func (m *SSLTestResourceModel) SetAlertSettingsModel(model *resource.AlertSettingsModel) {
	m.AlertSettings = model
}

// GetScheduleSettingsModel returns a pointer to the embedded ScheduleSettingsModel,
// which contains scheduling configuration for the SSL test.
func (m *SSLTestResourceModel) GetScheduleSettingsModel() *resource.ScheduleSettingsModel {
	if m == nil {
		return nil
	}
	return &m.ScheduleSettingsModel
}

// SetScheduleSettingsModel sets the embedded ScheduleSettingsModel with the provided model.
func (m *SSLTestResourceModel) SetScheduleSettingsModel(model *resource.ScheduleSettingsModel) {
	m.ScheduleSettingsModel = *model
}

// GetGatewayAddressOrHostModel returns a pointer to the embedded GatewayAddressOrHostModel,
// which contains gateway address or host configuration for the SSL test.
func (m *SSLTestResourceModel) GetGatewayAddressOrHostModel() *resource.GatewayAddressOrHostModel {
	if m == nil {
		return nil
	}
	return &m.GatewayAddressOrHostModel
}

// SetGatewayAddressOrHostModel sets the embedded GatewayAddressOrHostModel with the provided model.
func (m *SSLTestResourceModel) SetGatewayAddressOrHostModel(model *resource.GatewayAddressOrHostModel) {
	m.GatewayAddressOrHostModel = *model
}

// GetSSLSettingsModel returns a pointer to the embedded SSLSettingsModel,
// which contains SSL-specific configuration for the SSL test.
func (m *SSLTestResourceModel) GetSSLSettingsModel() *resource.SSLSettingsModel {
	if m == nil {
		return nil
	}
	return &m.SSLSettingsModel
}

// SetSSLSettingsModel sets the embedded SSLSettingsModel with the provided model.
func (m *SSLTestResourceModel) SetSSLSettingsModel(model *resource.SSLSettingsModel) {
	m.SSLSettingsModel = *model
}
