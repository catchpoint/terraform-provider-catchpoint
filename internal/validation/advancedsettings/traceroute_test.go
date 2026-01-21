package advancedsettings

import (
	"testing"

	"catchpoint-provider/internal/testutil"
	"catchpoint-provider/internal/types"
)

func TestGetAdvancedSettingsForTestTypeTraceroute(t *testing.T) {
	advancedSettings := GetAdvancedSettingsForTestType(types.TracerouteType)

	expectedKeys := []string{
		"enable_path_mtu_discovery",
		"failure_hop_count",
		"ping_count",
		// TODO: missing EnableDSCPPriorityProtocol
		// TODO: missing SetECN
	}

	testutil.AssertElementsMatch(t, "AdvancedSettingsTraceroute", advancedSettings, expectedKeys)
}

func TestGetAdvancedSettingsForTestAndMonitorCombinationTraceroute(t *testing.T) {
	test := []struct {
		monitor string
	}{
		{types.TracerouteICMPString},
		{types.TracerouteTCPString},
		{types.TracerouteUDPString},
	}
	for _, tt := range test {
		// Traceroute tests have three monitor types: ICMP, TCP, and UDP.
		advancedSettings := GetAdvancedSettingsForTestAndMonitorCombination(types.TracerouteType, tt.monitor)

		// All three Traceroute monitor types share these flags.
		expectedKeys := []string{
			"enable_path_mtu_discovery",
			"failure_hop_count",
			"ping_count",
			// TODO: missing EnableDSCPPriorityProtocol
			// TODO: missing SetECN
		}

		testutil.AssertElementsMatch(t, "AdvancedSettingsTraceroute", advancedSettings, expectedKeys)
	}
}
