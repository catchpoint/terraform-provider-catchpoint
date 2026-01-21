package resource

import (
	cptypes "catchpoint-provider/internal/types"

	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Interface that all test resource models must implement.
type TestResourceModelProvider interface {
	New() TestResourceModelProvider // An instance-factory instead of a package-level factory to satisfy generic programming patterns.
	SetID(types.Int64)
	GetID() types.Int64
	GetTestType() cptypes.TestType
	GetTestResourceModel() *BaseTestResourceModel
	IsNull() bool
}

func (m *BaseTestResourceModel) IsNull() bool {
	return m == nil
}

// Interface for models that can provide a URL field.
type URLProvider interface {
	GetURLField() (types.String, bool)
	SetURLField(url types.String)
}

// Interface for models that can provide a test script resource.
type TestScriptResourceProvider interface {
	GetTestScriptResourceModel() *TestScriptResourceModel
	SetTestScriptResourceModel(script *TestScriptResourceModel)
}

func (m *TestScriptResourceModel) IsNull() bool {
	return m == nil
}

// Interface for models that can provide a ThresholdModel.
type ThresholdModelProvider interface {
	GetTestThresholdModel() *ThresholdModel
	SetTestThresholdModel(threshold *ThresholdModel)
}

func (m *ThresholdModel) IsNull() bool {
	return m == nil
}

// Interface for models that can provide an AdvancedSettingsModel.
type AdvancedSettingsModelProvider interface {
	GetAdvancedSettingsModel() *AdvancedSettingsModel
	SetAdvancedSettingsModel(model *AdvancedSettingsModel)
}

func (m *AdvancedSettingsModel) IsNull() bool {
	return m == nil || m.AdvancedSettings.IsNull()
}

// Interface for models that can provide a AlertSettingsModel.
type AlertSettingsModelProvider interface {
	GetAlertSettingsModel() *AlertSettingsModel
	SetAlertSettingsModel(model *AlertSettingsModel)
}

func (m *AlertSettingsModel) IsNull() bool {
	return m == nil
}

func (m *AlertSettingsModel) IsUnknown() bool {
	return len(m.AlertRule) == 0 && m.NotificationGroup.IsNull()
}

func (m *NotificationGroupModel) IsNull() bool {
	return m == nil
}

// Interface for models that can provide a RequestSettingsModel.
type RequestSettingsModelProvider interface {
	GetRequestSettingsModel() *RequestSettingsModel
	SetRequestSettingsModel(model *RequestSettingsModel)
}

func (m *RequestSettingsModel) IsNull() bool {
	return m == nil || m.RequestSettings.IsNull()
}

// Interface for models that can provide a ScheduleSettingsModel.
type ScheduleSettingsModelProvider interface {
	GetScheduleSettingsModel() *ScheduleSettingsModel
	SetScheduleSettingsModel(model *ScheduleSettingsModel)
}

func (m *ScheduleSettingsModel) IsNull() bool {
	return m == nil || m.ScheduleSettings.IsNull()
}

// Interface for models that can provide a InsightSettingsModel.
type InsightSettingsModelProvider interface {
	GetInsightSettingsModel() *InsightSettingsModel
	SetInsightSettingsModel(model *InsightSettingsModel)
}

func (m *InsightSettingsModel) IsNull() bool {
	return m == nil || m.Insights.IsNull()
}

// Interface for models that can provide a GatewayAddressOrHostModel.
type GatewayAddressOrHostModelProvider interface {
	GetGatewayAddressOrHostModel() *GatewayAddressOrHostModel
	SetGatewayAddressOrHostModel(model *GatewayAddressOrHostModel)
}

func (m *GatewayAddressOrHostModel) IsNull() bool {
	return m == nil || m.GatewayAddressOrHost.IsNull()
}

// Interface for DNS test models that can provide QueryType and DNSServer fields.
type DNSTestModelProvider interface {
	GetQueryTypeField() (types.String, bool)
	SetQueryTypeField(queryType types.String)

	GetDNSServerField() (types.String, bool)
	SetDNSServerField(dnsServer types.String)
}

// Interface for models that can provide a SSLSettingsModel (SSL Test).
type SSLSettingsModelProvider interface {
	GetSSLSettingsModel() *SSLSettingsModel
	SetSSLSettingsModel(model *SSLSettingsModel)
}

func (m *SSLSettingsModel) IsNull() bool {
	return m == nil
}

// Interface for models that can provide ChromeVersion settings (Web/Transaction Chrome).
type ChromeVersionProvider interface {
	GetChromeVersionField() (types.String, bool)
	SetChromeVersionField(version types.String)
}

// Interface for models that can provide Simulate settings (Web/Transaction Mobile).
type SimulateProvider interface {
	GetSimulateField() (types.String, bool)
	SetSimulateField(simulate types.String)
}
