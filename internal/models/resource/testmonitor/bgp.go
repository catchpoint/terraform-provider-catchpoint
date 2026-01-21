package testmonitor

import (
	"catchpoint-provider/internal/models/resource"
	cptypes "catchpoint-provider/internal/types"

	"github.com/hashicorp/terraform-plugin-framework/types"
)

type BGPTestResourceModel struct {
	resource.BaseTestResourceModel
	AlertSettings *resource.AlertSettingsModel `tfsdk:"alert_settings"`
	Prefix        types.String                 `tfsdk:"prefix"`
}

func (m *BGPTestResourceModel) New() resource.TestResourceModelProvider {
	return &BGPTestResourceModel{}
}

func (m *BGPTestResourceModel) IsNull() bool {
	return m == nil
}

func (m *BGPTestResourceModel) GetTestType() cptypes.TestType {
	return cptypes.BGPType
}

// GetID returns the ID for the BGPTestResourceModel.
func (m *BGPTestResourceModel) GetID() types.Int64 {
	return m.ID
}

// SetID sets the ID for the BGPTestResourceModel.
func (m *BGPTestResourceModel) SetID(id types.Int64) {
	m.ID = id
}

// GetTestResourceModel returns the TestResourceModel for BGP tests.
func (m *BGPTestResourceModel) GetTestResourceModel() *resource.BaseTestResourceModel {
	if m == nil {
		return nil
	}
	return &m.BaseTestResourceModel
}

// GetURLField for BGPTestResourceModel returns the Prefix as the URL field.
func (m *BGPTestResourceModel) GetURLField() (types.String, bool) {
	if m.Prefix.IsNull() || m.Prefix.IsUnknown() {
		return m.Prefix, false
	}
	return m.Prefix, true
}

// SetURLField for BGPTestResourceModel sets the Prefix as the URL field.
func (m *BGPTestResourceModel) SetURLField(url types.String) {
	m.Prefix = url
}

// GetAlertSettingsModel returns the AlertSettingsModel for BGP tests.
func (m *BGPTestResourceModel) GetAlertSettingsModel() *resource.AlertSettingsModel {
	if m == nil {
		return nil
	}
	return m.AlertSettings
}

// SetAlertSettingsModel sets the AlertSettingsModel for BGP tests.
func (m *BGPTestResourceModel) SetAlertSettingsModel(model *resource.AlertSettingsModel) {
	m.AlertSettings = model
}
