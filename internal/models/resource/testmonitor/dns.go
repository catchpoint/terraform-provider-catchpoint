package testmonitor

import (
	"catchpoint-provider/internal/models/resource"
	cptypes "catchpoint-provider/internal/types"

	"github.com/hashicorp/terraform-plugin-framework/types"
)

type DNSTestResourceModel struct {
	resource.BaseTestResourceModel
	resource.AdvancedSettingsModel
	resource.ScheduleSettingsModel
	AlertSettings *resource.AlertSettingsModel `tfsdk:"alert_settings"`
	Thresholds    *resource.ThresholdModel     `tfsdk:"thresholds"`
	TestDomain    types.String                 `tfsdk:"test_domain"` // URL field.
	QueryType     types.String                 `tfsdk:"query_type"`
	DNSServer     types.String                 `tfsdk:"dns_server"`
}

// New creates a new instance of DNSTestResourceModel.
// Note: generally a New() method is not added to the interface but is a "package-level factory" function.
// However, in this case, we need it to create a new instance of the specific type T for our generic programming
// pattern in the TestResource struct.
// This allows us to create a new instance of the same type as the plan, which is necessary for merging state.
func (m *DNSTestResourceModel) New() resource.TestResourceModelProvider {
	return &DNSTestResourceModel{}
}

func (m *DNSTestResourceModel) IsNull() bool {
	return m == nil
}

// GetTestType returns the test type for the DNS test resource, which is cptypes.DNSType.
func (m *DNSTestResourceModel) GetTestType() cptypes.TestType {
	return cptypes.DNSType
}

// GetID returns the unique identifier (ID) for the DNSTestResourceModel.
func (m *DNSTestResourceModel) GetID() types.Int64 {
	return m.ID
}

// SetID sets the unique identifier (ID) for the DNSTestResourceModel.
func (m *DNSTestResourceModel) SetID(id types.Int64) {
	m.ID = id
}

// GetTestResourceModel returns a pointer to the embedded BaseTestResourceModel,
// which contains the core test resource fields.
func (m *DNSTestResourceModel) GetTestResourceModel() *resource.BaseTestResourceModel {
	if m == nil {
		return nil
	}
	return &m.BaseTestResourceModel
}

// GetURLField for DNSTestResourceModel returns the TestDomain as the URL field.
func (m *DNSTestResourceModel) GetURLField() (types.String, bool) {
	if m.TestDomain.IsNull() || m.TestDomain.IsUnknown() {
		return m.TestDomain, false
	}
	return m.TestDomain, true
}

// SetURLField for DNSTestResourceModel sets the TestDomain as the URL field.
func (m *DNSTestResourceModel) SetURLField(url types.String) {
	m.TestDomain = url
}

// GetTestThresholdModel returns a pointer to the ThresholdModel for the DNS test,
// which contains threshold configuration for test metrics.
func (m *DNSTestResourceModel) GetTestThresholdModel() *resource.ThresholdModel {
	if m == nil {
		return nil
	}
	return m.Thresholds
}

// SetTestThresholdModel sets the ThresholdModel for the DNS test with the provided threshold.
func (m *DNSTestResourceModel) SetTestThresholdModel(threshold *resource.ThresholdModel) {
	m.Thresholds = threshold
}

// GetAdvancedSettingsModel returns a pointer to the embedded AdvancedSettingsModel,
// which contains advanced configuration options for the DNS test.
func (m *DNSTestResourceModel) GetAdvancedSettingsModel() *resource.AdvancedSettingsModel {
	if m == nil {
		return nil
	}
	return &m.AdvancedSettingsModel
}

// SetAdvancedSettingsModel sets the embedded AdvancedSettingsModel with the provided model.
func (m *DNSTestResourceModel) SetAdvancedSettingsModel(model *resource.AdvancedSettingsModel) {
	m.AdvancedSettingsModel = *model
}

// GetAlertSettingsModel returns a pointer to the embedded AlertSettingsModel,
// which contains alert configuration for the DNS test.
func (m *DNSTestResourceModel) GetAlertSettingsModel() *resource.AlertSettingsModel {
	if m == nil {
		return nil
	}
	return m.AlertSettings
}

// SetAlertSettingsModel sets the embedded AlertSettingsModel with the provided model.
func (m *DNSTestResourceModel) SetAlertSettingsModel(model *resource.AlertSettingsModel) {
	m.AlertSettings = model
}

// GetScheduleSettingsModel returns a pointer to the embedded ScheduleSettingsModel,
// which contains scheduling configuration for the DNS test.
func (m *DNSTestResourceModel) GetScheduleSettingsModel() *resource.ScheduleSettingsModel {
	if m == nil {
		return nil
	}
	return &m.ScheduleSettingsModel
}

// SetScheduleSettingsModel sets the embedded ScheduleSettingsModel with the provided model.
func (m *DNSTestResourceModel) SetScheduleSettingsModel(model *resource.ScheduleSettingsModel) {
	m.ScheduleSettingsModel = *model
}

// GetQueryTypeField returns the QueryType field for the DNS test.
func (m *DNSTestResourceModel) GetQueryTypeField() (types.String, bool) {
	if m.QueryType.IsNull() || m.QueryType.IsUnknown() {
		return m.QueryType, false
	}
	return m.QueryType, true
}

// SetQueryTypeField sets the QueryType field for the DNS test.
func (m *DNSTestResourceModel) SetQueryTypeField(queryType types.String) {
	m.QueryType = queryType
}

// GetDNSServerField returns the DNSServer field for the DNS test. DNS Direct only.
func (m *DNSTestResourceModel) GetDNSServerField() (types.String, bool) {
	if m.DNSServer.IsNull() || m.DNSServer.IsUnknown() {
		return m.DNSServer, false
	}
	return m.DNSServer, true
}

// SetDNSServerField sets the DNSServer field for the DNS test. DNS Direct only.
func (m *DNSTestResourceModel) SetDNSServerField(dnsServer types.String) {
	m.DNSServer = dnsServer
}
