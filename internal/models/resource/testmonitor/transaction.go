package testmonitor

import (
	"catchpoint-provider/internal/models/resource"
	cptypes "catchpoint-provider/internal/types"

	"github.com/hashicorp/terraform-plugin-framework/types"
)

type TransactionTestResourceModel struct {
	resource.BaseTestResourceModel
	resource.AdvancedSettingsModel
	resource.RequestSettingsModel
	resource.ScheduleSettingsModel
	resource.InsightSettingsModel
	resource.GatewayAddressOrHostModel
	resource.TestScriptResourceModel
	AlertSettings *resource.AlertSettingsModel `tfsdk:"alert_settings"`
	Thresholds    *resource.ThresholdModel     `tfsdk:"thresholds"`
	Simulate      types.String                 `tfsdk:"simulate"`
	ChromeVersion types.String                 `tfsdk:"chrome_version"`
}

// New creates a new instance of TransactionTestResourceModel.
// Note: generally a New() method is not added to the interface but is a "package-level factory" function.
// However, in this case, we need it to create a new instance of the specific type T for our generic programming
// pattern in the TestResource struct.
// This allows us to create a new instance of the same type as the plan, which is necessary for merging state.
func (m *TransactionTestResourceModel) New() resource.TestResourceModelProvider {
	return &TransactionTestResourceModel{}
}

func (m *TransactionTestResourceModel) IsNull() bool {
	return m == nil
}

// GetTestType returns the test type for the Transaction test resource, which is cptypes.TransactionType.
func (m *TransactionTestResourceModel) GetTestType() cptypes.TestType {
	return cptypes.TransactionType
}

// GetID returns the unique identifier (ID) for the TransactionTestResourceModel.
func (m *TransactionTestResourceModel) GetID() types.Int64 {
	return m.ID
}

// SetID sets the unique identifier (ID) for the TransactionTestResourceModel.
func (m *TransactionTestResourceModel) SetID(id types.Int64) {
	m.ID = id
}

// GetTestResourceModel returns a pointer to the embedded BaseTestResourceModel,
// which contains the core test resource fields.
func (m *TransactionTestResourceModel) GetTestResourceModel() *resource.BaseTestResourceModel {
	if m == nil {
		return nil
	}
	return &m.BaseTestResourceModel
}

// GetTestScriptResourceModel returns a pointer to the embedded TestScriptResourceModel,
// which contains the script configuration for the Transaction test.
func (m *TransactionTestResourceModel) GetTestScriptResourceModel() *resource.TestScriptResourceModel {
	if m == nil {
		return nil
	}
	return &m.TestScriptResourceModel
}

// SetTestScriptResourceModel sets the embedded TestScriptResourceModel with the provided script.
func (m *TransactionTestResourceModel) SetTestScriptResourceModel(script *resource.TestScriptResourceModel) {
	m.TestScriptResourceModel = *script
}

// GetTestThresholdModel returns a pointer to the ThresholdModel for the Transaction test,
// which contains threshold configuration for test metrics.
func (m *TransactionTestResourceModel) GetTestThresholdModel() *resource.ThresholdModel {
	if m == nil {
		return nil
	}
	return m.Thresholds
}

// SetTestThresholdModel sets the ThresholdModel for the Transaction test with the provided threshold.
func (m *TransactionTestResourceModel) SetTestThresholdModel(threshold *resource.ThresholdModel) {
	m.Thresholds = threshold
}

// GetAdvancedSettingsModel returns a pointer to the embedded AdvancedSettingsModel,
// which contains advanced configuration options for the Transaction test.
func (m *TransactionTestResourceModel) GetAdvancedSettingsModel() *resource.AdvancedSettingsModel {
	if m == nil {
		return nil
	}
	return &m.AdvancedSettingsModel
}

// SetAdvancedSettingsModel sets the embedded AdvancedSettingsModel with the provided model.
func (m *TransactionTestResourceModel) SetAdvancedSettingsModel(model *resource.AdvancedSettingsModel) {
	m.AdvancedSettingsModel = *model
}

// GetAlertSettingsModel returns a pointer to the embedded AlertSettingsModel,
// which contains alert configuration for the Transaction test.
func (m *TransactionTestResourceModel) GetAlertSettingsModel() *resource.AlertSettingsModel {
	if m == nil {
		return nil
	}
	return m.AlertSettings
}

// SetAlertSettingsModel sets the embedded AlertSettingsModel with the provided model.
func (m *TransactionTestResourceModel) SetAlertSettingsModel(model *resource.AlertSettingsModel) {
	m.AlertSettings = model
}

// GetRequestSettingsModel returns a pointer to the embedded RequestSettingsModel,
// which contains request configuration for the Transaction test.
func (m *TransactionTestResourceModel) GetRequestSettingsModel() *resource.RequestSettingsModel {
	if m == nil {
		return nil
	}
	return &m.RequestSettingsModel
}

// SetRequestSettingsModel sets the embedded RequestSettingsModel with the provided model.
func (m *TransactionTestResourceModel) SetRequestSettingsModel(model *resource.RequestSettingsModel) {
	m.RequestSettingsModel = *model
}

// GetScheduleSettingsModel returns a pointer to the embedded ScheduleSettingsModel,
// which contains scheduling configuration for the Transaction test.
func (m *TransactionTestResourceModel) GetScheduleSettingsModel() *resource.ScheduleSettingsModel {
	if m == nil {
		return nil
	}
	return &m.ScheduleSettingsModel
}

// SetScheduleSettingsModel sets the embedded ScheduleSettingsModel with the provided model.
func (m *TransactionTestResourceModel) SetScheduleSettingsModel(model *resource.ScheduleSettingsModel) {
	m.ScheduleSettingsModel = *model
}

// GetInsightSettingsModel returns a pointer to the embedded InsightSettingsModel,
// which contains insight configuration for the Transaction test.
func (m *TransactionTestResourceModel) GetInsightSettingsModel() *resource.InsightSettingsModel {
	if m == nil {
		return nil
	}
	return &m.InsightSettingsModel
}

// SetInsightSettingsModel sets the embedded InsightSettingsModel with the provided model.
func (m *TransactionTestResourceModel) SetInsightSettingsModel(model *resource.InsightSettingsModel) {
	m.InsightSettingsModel = *model
}

// GetGatewayAddressOrHostModel returns a pointer to the embedded GatewayAddressOrHostModel,
// which contains gateway address or host configuration for the Transaction test.
func (m *TransactionTestResourceModel) GetGatewayAddressOrHostModel() *resource.GatewayAddressOrHostModel {
	if m == nil {
		return nil
	}
	return &m.GatewayAddressOrHostModel
}

// SetGatewayAddressOrHostModel sets the embedded GatewayAddressOrHostModel with the provided model.
func (m *TransactionTestResourceModel) SetGatewayAddressOrHostModel(model *resource.GatewayAddressOrHostModel) {
	m.GatewayAddressOrHostModel = *model
}

// Interface for models that can provide ChromeVersion settings (Web/Transaction Chrome).
func (m *TransactionTestResourceModel) GetChromeVersionField() (types.String, bool) {
	if m.ChromeVersion.IsNull() || m.ChromeVersion.IsUnknown() {
		return m.ChromeVersion, false
	}
	return m.ChromeVersion, true
}

func (m *TransactionTestResourceModel) SetChromeVersionField(version types.String) {
	m.ChromeVersion = version
}

// Interface for models that can provide Simulate settings (Web/Transaction Mobile).
func (m *TransactionTestResourceModel) GetSimulateField() (types.String, bool) {
	if m.Simulate.IsNull() || m.Simulate.IsUnknown() {
		return m.Simulate, false
	}
	return m.Simulate, true
}

func (m *TransactionTestResourceModel) SetSimulateField(simulate types.String) {
	m.Simulate = simulate
}
