package testmonitor

import (
	"catchpoint-provider/internal/models/resource"
	cptypes "catchpoint-provider/internal/types"

	"github.com/hashicorp/terraform-plugin-framework/types"
)

type WebTestResourceModel struct {
	resource.BaseTestResourceModel
	resource.AdvancedSettingsModel
	resource.RequestSettingsModel
	resource.ScheduleSettingsModel
	resource.InsightSettingsModel
	resource.GatewayAddressOrHostModel
	AlertSettings *resource.AlertSettingsModel `tfsdk:"alert_settings"`
	Thresholds    *resource.ThresholdModel     `tfsdk:"thresholds"`
	TestUrl       types.String                 `tfsdk:"test_url"`
	Simulate      types.String                 `tfsdk:"simulate"`
	ChromeVersion types.String                 `tfsdk:"chrome_version"`
}

// New creates a new instance of WebTestResourceModel.
// Note: generally a New() method is not added to the interface but is a "package-level factory" function.
// However, in this case, we need it to create a new instance of the specific type T for our generic programming
// pattern in the TestResource struct.
// This allows us to create a new instance of the same type as the plan, which is necessary for merging state.
func (m *WebTestResourceModel) New() resource.TestResourceModelProvider {
	return &WebTestResourceModel{}
}

func (m *WebTestResourceModel) IsNull() bool {
	return m == nil
}

// GetTestType returns the test type for the Web test resource, which is cptypes.WebType.
func (m *WebTestResourceModel) GetTestType() cptypes.TestType {
	return cptypes.WebType
}

// GetID returns the unique identifier (ID) for the WebTestResourceModel.
func (m *WebTestResourceModel) GetID() types.Int64 {
	return m.ID
}

// SetID sets the unique identifier (ID) for the WebTestResourceModel.
func (m *WebTestResourceModel) SetID(id types.Int64) {
	m.ID = id
}

// GetTestResourceModel returns a pointer to the embedded BaseTestResourceModel,
// which contains the core test resource fields.
func (m *WebTestResourceModel) GetTestResourceModel() *resource.BaseTestResourceModel {
	if m == nil {
		return nil
	}
	return &m.BaseTestResourceModel
}

// GetTestThresholdModel returns a pointer to the ThresholdModel for the Web test,
// which contains threshold configuration for test metrics.
func (m *WebTestResourceModel) GetTestThresholdModel() *resource.ThresholdModel {
	if m == nil {
		return nil
	}
	return m.Thresholds
}

// SetTestThresholdModel sets the ThresholdModel for the Web test with the provided threshold.
func (m *WebTestResourceModel) SetTestThresholdModel(threshold *resource.ThresholdModel) {
	m.Thresholds = threshold
}

// GetAdvancedSettingsModel returns a pointer to the embedded AdvancedSettingsModel,
// which contains advanced configuration options for the Web test.
func (m *WebTestResourceModel) GetAdvancedSettingsModel() *resource.AdvancedSettingsModel {
	if m == nil {
		return nil
	}
	return &m.AdvancedSettingsModel
}

// SetAdvancedSettingsModel sets the embedded AdvancedSettingsModel with the provided model.
func (m *WebTestResourceModel) SetAdvancedSettingsModel(model *resource.AdvancedSettingsModel) {
	m.AdvancedSettingsModel = *model
}

// GetAlertSettingsModel returns a pointer to the embedded AlertSettingsModel,
// which contains alert configuration for the Web test.
func (m *WebTestResourceModel) GetAlertSettingsModel() *resource.AlertSettingsModel {
	if m == nil {
		return nil
	}
	return m.AlertSettings
}

// SetAlertSettingsModel sets the embedded AlertSettingsModel with the provided model.
func (m *WebTestResourceModel) SetAlertSettingsModel(model *resource.AlertSettingsModel) {
	m.AlertSettings = model
}

// GetRequestSettingsModel returns a pointer to the embedded RequestSettingsModel,
// which contains request configuration for the Web test.
func (m *WebTestResourceModel) GetRequestSettingsModel() *resource.RequestSettingsModel {
	if m == nil {
		return nil
	}
	return &m.RequestSettingsModel
}

// SetRequestSettingsModel sets the embedded RequestSettingsModel with the provided model.
func (m *WebTestResourceModel) SetRequestSettingsModel(model *resource.RequestSettingsModel) {
	m.RequestSettingsModel = *model
}

// GetScheduleSettingsModel returns a pointer to the embedded ScheduleSettingsModel,
// which contains scheduling configuration for the Web test.
func (m *WebTestResourceModel) GetScheduleSettingsModel() *resource.ScheduleSettingsModel {
	if m == nil {
		return nil
	}
	return &m.ScheduleSettingsModel
}

// SetScheduleSettingsModel sets the embedded ScheduleSettingsModel with the provided model.
func (m *WebTestResourceModel) SetScheduleSettingsModel(model *resource.ScheduleSettingsModel) {
	m.ScheduleSettingsModel = *model
}

// GetInsightSettingsModel returns a pointer to the embedded InsightSettingsModel,
// which contains insight configuration for the Web test.
func (m *WebTestResourceModel) GetInsightSettingsModel() *resource.InsightSettingsModel {
	if m == nil {
		return nil
	}
	return &m.InsightSettingsModel
}

// SetInsightSettingsModel sets the embedded InsightSettingsModel with the provided model.
func (m *WebTestResourceModel) SetInsightSettingsModel(model *resource.InsightSettingsModel) {
	m.InsightSettingsModel = *model
}

// GetGatewayAddressOrHostModel returns a pointer to the embedded GatewayAddressOrHostModel,
// which contains gateway address or host configuration for the Web test.
func (m *WebTestResourceModel) GetGatewayAddressOrHostModel() *resource.GatewayAddressOrHostModel {
	if m == nil {
		return nil
	}
	return &m.GatewayAddressOrHostModel
}

// SetGatewayAddressOrHostModel sets the embedded GatewayAddressOrHostModel with the provided model.
func (m *WebTestResourceModel) SetGatewayAddressOrHostModel(model *resource.GatewayAddressOrHostModel) {
	m.GatewayAddressOrHostModel = *model
}

// GetURLField for WebTestResourceModel returns the test_url as the URL field.
func (m *WebTestResourceModel) GetURLField() (types.String, bool) {
	if m.TestUrl.IsNull() || m.TestUrl.IsUnknown() {
		return m.TestUrl, false
	}
	return m.TestUrl, true
}

// SetURLField for WebTestResourceModel sets the test_url as the URL field.
func (m *WebTestResourceModel) SetURLField(url types.String) {
	m.TestUrl = url
}

// Interface for models that can provide ChromeVersion settings (Web/Transaction Chrome).
func (m *WebTestResourceModel) GetChromeVersionField() (types.String, bool) {
	if m.ChromeVersion.IsNull() || m.ChromeVersion.IsUnknown() {
		return m.ChromeVersion, false
	}
	return m.ChromeVersion, true
}

func (m *WebTestResourceModel) SetChromeVersionField(version types.String) {
	m.ChromeVersion = version
}

// Interface for models that can provide Simulate settings (Web/Transaction Mobile).
func (m *WebTestResourceModel) GetSimulateField() (types.String, bool) {
	if m.Simulate.IsNull() || m.Simulate.IsUnknown() {
		return m.Simulate, false
	}
	return m.Simulate, true
}

func (m *WebTestResourceModel) SetSimulateField(simulate types.String) {
	m.Simulate = simulate
}
