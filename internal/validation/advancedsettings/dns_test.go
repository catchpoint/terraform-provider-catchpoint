package advancedsettings

import (
	"testing"

	"catchpoint-provider/internal/testutil"
	"catchpoint-provider/internal/types"
)

func TestGetAdvancedSettingsForTestTypeDNS(t *testing.T) {
	advancedSettings := GetAdvancedSettingsForTestType(types.DNSType)

	expectedKeys := []string{
		"additional_monitor",
		"debug_primary_host_on_failure",
		"disable_recursive_resolution",
		"edns_subnet",
		"enable_dnssec",
		"enable_path_mtu_discovery",
		"favor_fastest_round_trip_nameserver",
		"try_next_nameserver_on_failure",
		"enable_nsid",
		"enable_tcp_protocol",
	}

	testutil.AssertElementsMatch(t, "AdvancedSettingsDNS", advancedSettings, expectedKeys)
}

func TestGetAdvancedSettingsForTestAndMonitorCombinationDNS(t *testing.T) {
	test := []struct {
		testMonitor         string
		additionalTestFlags []string
	}{
		{types.DNSDirectString, []string{"enable_dnssec"}},                           // TODO: missing "Set DNS Query Attempts and Timeout"
		{types.DNSExperienceString, []string{"favor_fastest_round_trip_nameserver"}}, //TODO: missing Cache TLD Nameserver Queries.
	}
	for _, tt := range test {
		// DNS tests have either DNS Direct or DNS Experience monitors.
		advancedSettings := GetAdvancedSettingsForTestAndMonitorCombination(types.DNSType, tt.testMonitor)

		// DNS Direct and DNS Experience share these flags.
		expectedKeys := []string{
			"additional_monitor",
			"debug_primary_host_on_failure",
			"disable_recursive_resolution",
			"edns_subnet",
			"enable_path_mtu_discovery",
			"enable_tcp_protocol",
			"enable_nsid",
			"try_next_nameserver_on_failure",
		}

		// Append the unique flag(s) for the monitor type.
		expectedKeys = append(expectedKeys, tt.additionalTestFlags...)

		testutil.AssertElementsMatch(t, "AdvancedSettingsDNS_"+tt.testMonitor, advancedSettings, expectedKeys)
	}
}
