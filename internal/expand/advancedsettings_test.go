package expand

import (
	"testing"

	"catchpoint-provider/internal/fields"
	"catchpoint-provider/internal/models"
	"catchpoint-provider/internal/testutil"
	cptypes "catchpoint-provider/internal/types"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func TestExpandAdvancedSettingsNullObject(t *testing.T) {
	obj := types.ObjectNull(map[string]attr.Type{})
	config := &models.AdvancedSettingsConfig{}

	diags := ExpandAdvancedSettingsConfig(obj, config)

	testutil.AssertDiagsHasNoErrors(t, diags)
	testutil.AssertEqual(t, "AdvancedSettingTypeID", config.AdvancedSettingType.ID, 0)
	testutil.AssertEqual(t, "AdvancedSettingTypeName", config.AdvancedSettingType.Name, "inherit")
	// Config should remain unchanged when object is null
}

func TestExpandAdvancedSettingsEmptyObject(t *testing.T) {
	attrs := map[string]attr.Value{}

	obj, _ := types.ObjectValue(map[string]attr.Type{}, attrs)
	config := &models.AdvancedSettingsConfig{}

	diags := ExpandAdvancedSettingsConfig(obj, config)

	testutil.AssertDiagsHasNoErrors(t, diags)
	// Should set defaults for empty object
	testutil.AssertEqual(t, "applied test flags length", len(config.AppliedTestFlags), 0)
	// AdvancedSettingType should not be set for empty object
	testutil.AssertEqual(t, "AdvancedSettingTypeID", config.AdvancedSettingType.ID, 0)
	testutil.AssertEqual(t, "AdvancedSettingTypeName", config.AdvancedSettingType.Name, "inherit")
}

func TestExpandAdvancedSettingsNoSettings(t *testing.T) {
	attrs := map[string]attr.Value{
		"advanced_setting_type":     types.StringValue("inherit"),
		"enable_path_mtu_discovery": types.BoolNull(),
		"failure_hop_count":         types.Int64Null(),
		"ping_count":                types.Int64Null(),
	}

	obj, _ := types.ObjectValue(map[string]attr.Type{
		"advanced_setting_type":     types.StringType,
		"enable_path_mtu_discovery": types.BoolType,
		"failure_hop_count":         types.Int64Type,
		"ping_count":                types.Int64Type,
	}, attrs)

	config := &models.AdvancedSettingsConfig{}

	diags := ExpandAdvancedSettingsConfig(obj, config)

	testutil.AssertDiagsHasNoErrors(t, diags)
	// Should set defaults for empty object
	testutil.AssertEqual(t, "applied test flags length", len(config.AppliedTestFlags), 0)
	// AdvancedSettingType should not be set for empty object
	testutil.AssertEqual(t, "AdvancedSettingTypeID", config.AdvancedSettingType.ID, 0)
	testutil.AssertEqual(t, "AdvancedSettingTypeName", config.AdvancedSettingType.Name, "inherit")
}

func TestExpandAdvancedSettingsStringFields(t *testing.T) {
	attrs := map[string]attr.Value{
		fields.AdditionalMonitor:   types.StringValue("ping tcp"),
		fields.BandwidthThrottling: types.StringValue("dsl"),
		fields.EDNSSubnet:          types.StringValue("192.168.1.0/24"),
	}

	obj, _ := types.ObjectValue(map[string]attr.Type{
		fields.AdditionalMonitor:   types.StringType,
		fields.BandwidthThrottling: types.StringType,
		fields.EDNSSubnet:          types.StringType,
	}, attrs)

	config := &models.AdvancedSettingsConfig{}

	diags := ExpandAdvancedSettingsConfig(obj, config)

	testutil.AssertDiagsHasNoErrors(t, diags)
	testutil.AssertEqual(t, "advanced setting type name", config.AdvancedSettingType.Name, cptypes.Override)
	testutil.AssertEqual(t, "additional monitor type name", config.AdditionalMonitorType.Name, "ping tcp")
	testutil.AssertEqual(t, "bandwidth throttling name", config.BandwidthThrottling.Name, "dsl")
	testutil.AssertEqual(t, "edns subnet", config.EDNSSubnet, "192.168.1.0/24")
}

func TestExpandAdvancedSettingsIntegerFields(t *testing.T) {
	attrs := map[string]attr.Value{
		fields.EnforceTestFailureIfRunsLongerThan: types.Int64Value(300),
		fields.PingCount:       types.Int64Value(5),
		fields.FailureHopCount: types.Int64Value(10),
		fields.ViewportHeight:  types.Int64Value(768),
		fields.ViewportWidth:   types.Int64Value(1024),
	}

	obj, _ := types.ObjectValue(map[string]attr.Type{
		fields.EnforceTestFailureIfRunsLongerThan: types.Int64Type,
		fields.PingCount:       types.Int64Type,
		fields.FailureHopCount: types.Int64Type,
		fields.ViewportHeight:  types.Int64Type,
		fields.ViewportWidth:   types.Int64Type,
	}, attrs)

	config := &models.AdvancedSettingsConfig{}

	diags := ExpandAdvancedSettingsConfig(obj, config)

	testutil.AssertDiagsHasNoErrors(t, diags)
	testutil.AssertEqual(t, "advanced setting type name", config.AdvancedSettingType.Name, cptypes.Override)
	testutil.AssertEqual(t, "max step runtime override", config.MaxStepRuntimeSecOverride, 300)
	testutil.AssertEqual(t, "traceroute ping count", config.TraceroutePingCount, 5)
	testutil.AssertEqual(t, "traceroute failure hop count", config.TracerouteFailureHopCount, 10)
	testutil.AssertEqual(t, "viewport height", config.ViewportHeight, 768)
	testutil.AssertEqual(t, "viewport width", config.ViewportWidth, 1024)
}

func TestExpandAdvancedSettingsWaitForNoActivitySpecialCase(t *testing.T) {
	attrs := map[string]attr.Value{
		fields.StopTestOnDocumentComplete: types.BoolValue(true),
		fields.WaitForNoActivity:          types.Int64Value(2000),
	}

	obj, _ := types.ObjectValue(map[string]attr.Type{
		fields.StopTestOnDocumentComplete: types.BoolType,
		fields.WaitForNoActivity:          types.Int64Type,
	}, attrs)

	config := &models.AdvancedSettingsConfig{}

	diags := ExpandAdvancedSettingsConfig(obj, config)

	testutil.AssertDiagsHasNoErrors(t, diags)
	testutil.AssertEqual(t, "advanced setting type name", config.AdvancedSettingType.Name, cptypes.Override)
	testutil.AssertNotNil(t, "wait for no activity should be set", config.WaitForNoActivityOnDocComplete)
	if config.WaitForNoActivityOnDocComplete != nil {
		testutil.AssertEqual(t, "wait for no activity value", *config.WaitForNoActivityOnDocComplete, 2000)
	}
}

func TestExpandAdvancedSettingsWaitForNoActivityOnlyStopTest(t *testing.T) {
	attrs := map[string]attr.Value{
		fields.StopTestOnDocumentComplete: types.BoolValue(true),
		// No WaitForNoActivity field
	}

	obj, _ := types.ObjectValue(map[string]attr.Type{
		fields.StopTestOnDocumentComplete: types.BoolType,
	}, attrs)

	config := &models.AdvancedSettingsConfig{}

	diags := ExpandAdvancedSettingsConfig(obj, config)

	testutil.AssertDiagsHasNoErrors(t, diags)
	testutil.AssertEqual(t, "advanced setting type name", config.AdvancedSettingType.Name, cptypes.Override)
	testutil.AssertNil(t, "wait for no activity should not be set", config.WaitForNoActivityOnDocComplete)
}

func TestExpandAdvancedSettingsWaitForNoActivityOnlyWaitField(t *testing.T) {
	attrs := map[string]attr.Value{
		fields.WaitForNoActivity: types.Int64Value(2000),
		// No StopTestOnDocumentComplete field
	}

	obj, _ := types.ObjectValue(map[string]attr.Type{
		fields.WaitForNoActivity: types.Int64Type,
	}, attrs)

	config := &models.AdvancedSettingsConfig{}

	diags := ExpandAdvancedSettingsConfig(obj, config)

	testutil.AssertDiagsHasNoErrors(t, diags)
	testutil.AssertEqual(t, "advanced setting type name", config.AdvancedSettingType.Name, cptypes.Override)
	testutil.AssertNil(t, "wait for no activity should not be set", config.WaitForNoActivityOnDocComplete)
}

func TestExpandAdvancedSettingsTestFlags(t *testing.T) {
	// Test with some valid test flags
	attrs := map[string]attr.Value{
		cptypes.EnablePathMTUDiscovery:         types.BoolValue(true),
		cptypes.DisableCrossOriginIframeAccess: types.BoolValue(true),
		"disable_javascript":                   types.BoolValue(false), // Should not be included
	}

	obj, _ := types.ObjectValue(map[string]attr.Type{
		cptypes.EnablePathMTUDiscovery:         types.BoolType,
		cptypes.DisableCrossOriginIframeAccess: types.BoolType,
		"disable_javascript":                   types.BoolType,
	}, attrs)

	config := &models.AdvancedSettingsConfig{}

	diags := ExpandAdvancedSettingsConfig(obj, config)

	testutil.AssertDiagsHasNoErrors(t, diags)
	testutil.AssertEqual(t, "advanced setting type name", config.AdvancedSettingType.Name, cptypes.Override)

	// Should have 3 test flags (excluding the false one)
	testutil.AssertEqual(t, "applied test flags length", len(config.AppliedTestFlags), 2)

	// Verify the flags were converted to IDs correctly
	expectedFlags := []string{cptypes.EnablePathMTUDiscovery, cptypes.DisableCrossOriginIframeAccess}
	for _, flagName := range expectedFlags {
		if flagID, ok := cptypes.GetTestFlagID(flagName); ok {
			testutil.AssertContains(t, "applied test flags", config.AppliedTestFlags, flagID)
		}
	}
}

func TestExpandAdvancedSettingsTestFlagsNullValues(t *testing.T) {
	attrs := map[string]attr.Value{
		cptypes.EnablePathMTUDiscovery:         types.BoolValue(true),
		cptypes.DisableCrossOriginIframeAccess: types.BoolNull(), // Null value should be ignored
		cptypes.DebugPrimaryHostOnFailure:      types.BoolValue(true),
	}

	obj, _ := types.ObjectValue(map[string]attr.Type{
		cptypes.EnablePathMTUDiscovery:         types.BoolType,
		cptypes.DisableCrossOriginIframeAccess: types.BoolType,
		cptypes.DebugPrimaryHostOnFailure:      types.BoolType,
	}, attrs)

	config := &models.AdvancedSettingsConfig{}

	diags := ExpandAdvancedSettingsConfig(obj, config)

	testutil.AssertDiagsHasNoErrors(t, diags)
	testutil.AssertEqual(t, "advanced setting type name", config.AdvancedSettingType.Name, cptypes.Override)

	// Should have 2 test flags (excluding the null one)
	testutil.AssertEqual(t, "applied test flags length", len(config.AppliedTestFlags), 2)
}

func TestExpandAdvancedSettingsTestFlagsUnknownValues(t *testing.T) {
	attrs := map[string]attr.Value{
		cptypes.EnablePathMTUDiscovery:         types.BoolValue(true),
		cptypes.EnableBindHostname:             types.BoolUnknown(), // Unknown value should be ignored
		cptypes.DisableCrossOriginIframeAccess: types.BoolValue(true),
	}

	obj, _ := types.ObjectValue(map[string]attr.Type{
		cptypes.EnablePathMTUDiscovery:         types.BoolType,
		cptypes.EnableBindHostname:             types.BoolType,
		cptypes.DisableCrossOriginIframeAccess: types.BoolType,
	}, attrs)

	config := &models.AdvancedSettingsConfig{}

	diags := ExpandAdvancedSettingsConfig(obj, config)

	testutil.AssertDiagsHasNoErrors(t, diags)
	testutil.AssertEqual(t, "advanced setting type name", config.AdvancedSettingType.Name, cptypes.Override)

	// Should have 2 test flags (excluding the unknown one)
	testutil.AssertEqual(t, "applied test flags length", len(config.AppliedTestFlags), 2)
}

func TestExpandAdvancedSettingsNullStringFields(t *testing.T) {
	attrs := map[string]attr.Value{
		fields.AdditionalMonitor:   types.StringNull(),
		fields.BandwidthThrottling: types.StringNull(),
		fields.EDNSSubnet:          types.StringNull(),
	}

	obj, _ := types.ObjectValue(map[string]attr.Type{
		fields.AdditionalMonitor:   types.StringType,
		fields.BandwidthThrottling: types.StringType,
		fields.EDNSSubnet:          types.StringType,
	}, attrs)

	config := &models.AdvancedSettingsConfig{}

	diags := ExpandAdvancedSettingsConfig(obj, config)

	testutil.AssertDiagsHasNoErrors(t, diags)
	testutil.AssertEqual(t, "advanced setting type name", config.AdvancedSettingType.Name, cptypes.Inherit)
	// Null string fields should result in empty/default values
	testutil.AssertEqual(t, "additional monitor type name", config.AdditionalMonitorType.Name, "")
	testutil.AssertEqual(t, "bandwidth throttling name", config.BandwidthThrottling.Name, "")
	testutil.AssertEqual(t, "edns subnet", config.EDNSSubnet, "")
}

func TestExpandAdvancedSettingsNullIntegerFields(t *testing.T) {
	attrs := map[string]attr.Value{
		fields.EnforceTestFailureIfRunsLongerThan: types.Int64Null(),
		fields.PingCount:       types.Int64Null(),
		fields.FailureHopCount: types.Int64Null(),
		fields.ViewportHeight:  types.Int64Null(),
		fields.ViewportWidth:   types.Int64Null(),
	}

	obj, _ := types.ObjectValue(map[string]attr.Type{
		fields.EnforceTestFailureIfRunsLongerThan: types.Int64Type,
		fields.PingCount:       types.Int64Type,
		fields.FailureHopCount: types.Int64Type,
		fields.ViewportHeight:  types.Int64Type,
		fields.ViewportWidth:   types.Int64Type,
	}, attrs)

	config := &models.AdvancedSettingsConfig{}

	diags := ExpandAdvancedSettingsConfig(obj, config)

	testutil.AssertDiagsHasNoErrors(t, diags)
	testutil.AssertEqual(t, "advanced setting type name", config.AdvancedSettingType.Name, cptypes.Inherit)
	// Null integer fields should result in zero values
	testutil.AssertEqual(t, "max step runtime override", config.MaxStepRuntimeSecOverride, 0)
	testutil.AssertEqual(t, "traceroute ping count", config.TraceroutePingCount, 0)
	testutil.AssertEqual(t, "traceroute failure hop count", config.TracerouteFailureHopCount, 0)
	testutil.AssertEqual(t, "viewport height", config.ViewportHeight, 0)
	testutil.AssertEqual(t, "viewport width", config.ViewportWidth, 0)
}

func TestExpandAdvancedSettingsZeroValues(t *testing.T) {
	attrs := map[string]attr.Value{
		fields.EnforceTestFailureIfRunsLongerThan: types.Int64Value(0),
		fields.PingCount:      types.Int64Value(0),
		fields.ViewportHeight: types.Int64Value(0),
		fields.ViewportWidth:  types.Int64Value(0),
	}

	obj, _ := types.ObjectValue(map[string]attr.Type{
		fields.EnforceTestFailureIfRunsLongerThan: types.Int64Type,
		fields.PingCount:      types.Int64Type,
		fields.ViewportHeight: types.Int64Type,
		fields.ViewportWidth:  types.Int64Type,
	}, attrs)

	config := &models.AdvancedSettingsConfig{}

	diags := ExpandAdvancedSettingsConfig(obj, config)

	testutil.AssertDiagsHasNoErrors(t, diags)
	testutil.AssertEqual(t, "advanced setting type name", config.AdvancedSettingType.Name, cptypes.Inherit)
	testutil.AssertEqual(t, "max step runtime override", config.MaxStepRuntimeSecOverride, 0)
	testutil.AssertEqual(t, "traceroute ping count", config.TraceroutePingCount, 0)
	testutil.AssertEqual(t, "viewport height", config.ViewportHeight, 0)
	testutil.AssertEqual(t, "viewport width", config.ViewportWidth, 0)
}

func TestExpandAdvancedSettingsLargeValues(t *testing.T) {
	attrs := map[string]attr.Value{
		fields.EnforceTestFailureIfRunsLongerThan: types.Int64Value(999999),
		fields.PingCount:      types.Int64Value(100),
		fields.ViewportHeight: types.Int64Value(4320),
		fields.ViewportWidth:  types.Int64Value(7680),
	}

	obj, _ := types.ObjectValue(map[string]attr.Type{
		fields.EnforceTestFailureIfRunsLongerThan: types.Int64Type,
		fields.PingCount:      types.Int64Type,
		fields.ViewportHeight: types.Int64Type,
		fields.ViewportWidth:  types.Int64Type,
	}, attrs)

	config := &models.AdvancedSettingsConfig{}

	diags := ExpandAdvancedSettingsConfig(obj, config)

	testutil.AssertDiagsHasNoErrors(t, diags)
	testutil.AssertEqual(t, "advanced setting type name", config.AdvancedSettingType.Name, cptypes.Override)
	testutil.AssertEqual(t, "max step runtime override", config.MaxStepRuntimeSecOverride, 999999)
	testutil.AssertEqual(t, "traceroute ping count", config.TraceroutePingCount, 100)
	testutil.AssertEqual(t, "viewport height", config.ViewportHeight, 4320)
	testutil.AssertEqual(t, "viewport width", config.ViewportWidth, 7680)
}

func TestExpandAdvancedSettingsCompleteConfiguration(t *testing.T) {
	attrs := map[string]attr.Value{
		fields.AdditionalMonitor:                  types.StringValue("traceroute tcp"),
		fields.BandwidthThrottling:                types.StringValue("wifi"),
		fields.EDNSSubnet:                         types.StringValue("10.0.0.0/8"),
		fields.EnforceTestFailureIfRunsLongerThan: types.Int64Value(600),
		fields.PingCount:                          types.Int64Value(10),
		fields.FailureHopCount:                    types.Int64Value(15),
		fields.ViewportHeight:                     types.Int64Value(900),
		fields.ViewportWidth:                      types.Int64Value(1440),
		fields.StopTestOnDocumentComplete:         types.BoolValue(true),  // Flag
		fields.WaitForNoActivity:                  types.Int64Value(3000), // Not actually a flag
		cptypes.DisableCrossOriginIframeAccess:    types.BoolValue(true),  // Flag
		cptypes.CaptureScreenshot:                 types.BoolValue(true),  // Flag
		cptypes.CaptureFilmstrip:                  types.BoolValue(true),  // Flag
	}

	obj, _ := types.ObjectValue(map[string]attr.Type{
		fields.AdditionalMonitor:                  types.StringType,
		fields.BandwidthThrottling:                types.StringType,
		fields.EDNSSubnet:                         types.StringType,
		fields.EnforceTestFailureIfRunsLongerThan: types.Int64Type,
		fields.PingCount:                          types.Int64Type,
		fields.FailureHopCount:                    types.Int64Type,
		fields.ViewportHeight:                     types.Int64Type,
		fields.ViewportWidth:                      types.Int64Type,
		fields.StopTestOnDocumentComplete:         types.BoolType,
		fields.WaitForNoActivity:                  types.Int64Type,
		cptypes.DisableCrossOriginIframeAccess:    types.BoolType,
		cptypes.CaptureScreenshot:                 types.BoolType,
		cptypes.CaptureFilmstrip:                  types.BoolType,
	}, attrs)

	config := &models.AdvancedSettingsConfig{}

	diags := ExpandAdvancedSettingsConfig(obj, config)

	testutil.AssertDiagsHasNoErrors(t, diags)

	// Verify all fields
	testutil.AssertEqual(t, "advanced setting type name", config.AdvancedSettingType.Name, cptypes.Override)
	testutil.AssertEqual(t, "additional monitor type name", config.AdditionalMonitorType.Name, "traceroute tcp")
	testutil.AssertEqual(t, "bandwidth throttling name", config.BandwidthThrottling.Name, "wifi")
	testutil.AssertEqual(t, "edns subnet", config.EDNSSubnet, "10.0.0.0/8")
	testutil.AssertEqual(t, "max step runtime override", config.MaxStepRuntimeSecOverride, 600)
	testutil.AssertEqual(t, "traceroute ping count", config.TraceroutePingCount, 10)
	testutil.AssertEqual(t, "traceroute failure hop count", config.TracerouteFailureHopCount, 15)
	testutil.AssertEqual(t, "viewport height", config.ViewportHeight, 900)
	testutil.AssertEqual(t, "viewport width", config.ViewportWidth, 1440)

	// Verify wait for no activity special case
	testutil.AssertNotNil(t, "wait for no activity should be set", config.WaitForNoActivityOnDocComplete)
	if config.WaitForNoActivityOnDocComplete != nil {
		testutil.AssertEqual(t, "wait for no activity value", *config.WaitForNoActivityOnDocComplete, 3000)
	}

	// Verify test flags
	testutil.AssertEqual(t, "applied test flags length", len(config.AppliedTestFlags), 4)
}
