package testmonitor

import (
	"catchpoint-provider/internal/models/resource"
	cptypes "catchpoint-provider/internal/types"

	"github.com/hashicorp/terraform-plugin-framework/types"
)

type APITestResourceModel struct {
	resource.BaseTestResourceModel
	resource.AdvancedSettingsModel
	resource.RequestSettingsModel
	resource.ScheduleSettingsModel
	resource.InsightSettingsModel
	resource.GatewayAddressOrHostModel
	resource.TestScriptResourceModel
	AlertSettings *resource.AlertSettingsModel `tfsdk:"alert_settings"`
	Thresholds    *resource.ThresholdModel     `tfsdk:"thresholds"`
}

// New creates a new instance of APITestResourceModel.
// Note: generally a New() method is not added to the interface but is a "package-level factory" function.
// However, in this case, we need it to create a new instance of the specific type T for our generic programming
// pattern in the TestResource struct.
// This allows us to create a new instance of the same type as the plan, which is necessary for merging state.
func (m *APITestResourceModel) New() resource.TestResourceModelProvider {
	return &APITestResourceModel{}
}

func (m *APITestResourceModel) IsNull() bool {
	return m == nil
}

// GetTestType returns the test type for the API test resource, which is cptypes.APIType.
func (m *APITestResourceModel) GetTestType() cptypes.TestType {
	return cptypes.APIType
}

// GetID returns the unique identifier (ID) for the APITestResourceModel.
func (m *APITestResourceModel) GetID() types.Int64 {
	return m.ID
}

// SetID sets the unique identifier (ID) for the APITestResourceModel.
func (m *APITestResourceModel) SetID(id types.Int64) {
	m.ID = id
}

// GetTestResourceModel returns a pointer to the embedded BaseTestResourceModel,
// which contains the core test resource fields.
func (m *APITestResourceModel) GetTestResourceModel() *resource.BaseTestResourceModel {
	if m == nil {
		return nil
	}
	return &m.BaseTestResourceModel
}

// GetTestScriptResourceModel returns a pointer to the embedded TestScriptResourceModel,
// which contains the script configuration for the API test.
func (m *APITestResourceModel) GetTestScriptResourceModel() *resource.TestScriptResourceModel {
	if m == nil {
		return nil
	}
	return &m.TestScriptResourceModel
}

// SetTestScriptResourceModel sets the embedded TestScriptResourceModel with the provided script.
func (m *APITestResourceModel) SetTestScriptResourceModel(script *resource.TestScriptResourceModel) {
	m.TestScriptResourceModel = *script
}

// GetTestThresholdModel returns a pointer to the ThresholdModel for the API test,
// which contains threshold configuration for test metrics.
func (m *APITestResourceModel) GetTestThresholdModel() *resource.ThresholdModel {
	if m == nil {
		return nil
	}
	return m.Thresholds
}

// SetTestThresholdModel sets the ThresholdModel for the API test with the provided threshold.
func (m *APITestResourceModel) SetTestThresholdModel(threshold *resource.ThresholdModel) {
	m.Thresholds = threshold
}

// GetAdvancedSettingsModel returns a pointer to the embedded AdvancedSettingsModel,
// which contains advanced configuration options for the API test.
func (m *APITestResourceModel) GetAdvancedSettingsModel() *resource.AdvancedSettingsModel {
	if m == nil {
		return nil
	}
	return &m.AdvancedSettingsModel
}

// SetAdvancedSettingsModel sets the embedded AdvancedSettingsModel with the provided model.
func (m *APITestResourceModel) SetAdvancedSettingsModel(model *resource.AdvancedSettingsModel) {
	m.AdvancedSettingsModel = *model
}

// GetAlertSettingsModel returns a pointer to the embedded AlertSettingsModel,
// which contains alert configuration for the API test.
func (m *APITestResourceModel) GetAlertSettingsModel() *resource.AlertSettingsModel {
	if m == nil {
		return nil
	}
	return m.AlertSettings
}

// SetAlertSettingsModel sets the embedded AlertSettingsModel with the provided model.
func (m *APITestResourceModel) SetAlertSettingsModel(model *resource.AlertSettingsModel) {
	m.AlertSettings = model
}

// GetRequestSettingsModel returns a pointer to the embedded RequestSettingsModel,
// which contains request configuration for the API test.
func (m *APITestResourceModel) GetRequestSettingsModel() *resource.RequestSettingsModel {
	if m == nil {
		return nil
	}
	return &m.RequestSettingsModel
}

// SetRequestSettingsModel sets the embedded RequestSettingsModel with the provided model.
func (m *APITestResourceModel) SetRequestSettingsModel(model *resource.RequestSettingsModel) {
	m.RequestSettingsModel = *model
}

// GetScheduleSettingsModel returns a pointer to the embedded ScheduleSettingsModel,
// which contains scheduling configuration for the API test.
func (m *APITestResourceModel) GetScheduleSettingsModel() *resource.ScheduleSettingsModel {
	if m == nil {
		return nil
	}
	return &m.ScheduleSettingsModel
}

// SetScheduleSettingsModel sets the embedded ScheduleSettingsModel with the provided model.
func (m *APITestResourceModel) SetScheduleSettingsModel(model *resource.ScheduleSettingsModel) {
	m.ScheduleSettingsModel = *model
}

// GetInsightSettingsModel returns a pointer to the embedded InsightSettingsModel,
// which contains insight configuration for the API test.
func (m *APITestResourceModel) GetInsightSettingsModel() *resource.InsightSettingsModel {
	if m == nil {
		return nil
	}
	return &m.InsightSettingsModel
}

// SetInsightSettingsModel sets the embedded InsightSettingsModel with the provided model.
func (m *APITestResourceModel) SetInsightSettingsModel(model *resource.InsightSettingsModel) {
	m.InsightSettingsModel = *model
}

// GetGatewayAddressOrHostModel returns a pointer to the embedded GatewayAddressOrHostModel,
// which contains gateway address or host configuration for the API test.
func (m *APITestResourceModel) GetGatewayAddressOrHostModel() *resource.GatewayAddressOrHostModel {
	if m == nil {
		return nil
	}
	return &m.GatewayAddressOrHostModel
}

// SetGatewayAddressOrHostModel sets the embedded GatewayAddressOrHostModel with the provided model.
func (m *APITestResourceModel) SetGatewayAddressOrHostModel(model *resource.GatewayAddressOrHostModel) {
	m.GatewayAddressOrHostModel = *model
}
