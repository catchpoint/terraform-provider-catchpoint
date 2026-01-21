package advancedsettings

import (
	"testing"

	"catchpoint-provider/internal/testutil"
	"catchpoint-provider/internal/types"
)

func TestGetAdvancedSettingsForTestTypePing(t *testing.T) {
	advancedSettings := GetAdvancedSettingsForTestType(types.PingType)

	expectedKeys := []string{
		"additional_monitor",
		"debug_primary_host_on_failure",
		"enable_path_mtu_discovery",
		"verify_test_on_failure",
	}

	testutil.AssertElementsMatch(t, "AdvancedSettingsPing", advancedSettings, expectedKeys)
}

func TestGetAdvancedSettingsForTestAndMonitorCombinationPing(t *testing.T) {
	test := []struct {
		monitor string
	}{
		{types.PingICMPString},
		{types.PingTCPString},
		{types.PingUDPString},
	}
	for _, tt := range test {
		// Ping tests have three monitor types: ICMP, TCP, and UDP.
		advancedSettings := GetAdvancedSettingsForTestAndMonitorCombination(types.PingType, tt.monitor)

		// All three Ping monitor types share these flags.
		expectedKeys := []string{
			"additional_monitor",
			"debug_primary_host_on_failure",
			"enable_path_mtu_discovery",
			"verify_test_on_failure",
		}

		testutil.AssertElementsMatch(t, "AdvancedSettingsPing", advancedSettings, expectedKeys)
	}
}
