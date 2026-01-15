package advancedsettings

import (
	"testing"

	"catchpoint-provider/internal/testutil"
	"catchpoint-provider/internal/types"
)

const (
	fakeTestType types.TestType = 0
)

// Fake advanced settings for testing purposes.
var validAdvancedSettings = []string{
	"some_advanced_setting",
	"another_advanced_setting",
	"yet_another_advanced_setting",
}

// Every advanced setting we know about should be listed here for Product and Folder schemas.
// This is more of an integration test.
func TestGetAllAdvancedSettings(t *testing.T) {
	advancedSettings := GetAllAdvancedSettings()

	expectedKeys := []string{
		"additional_monitor",
		"allow_test_download_limit_override",
		"bandwidth_throttling",
		"capture_filmstrip",
		"capture_http_headers",
		"capture_response_content",
		"capture_screenshot",
		"certificate_revocation_disabled",
		"debug_primary_host_on_failure",
		"debug_referenced_hosts_on_failure",
		"disable_cross_origin_iframe_access",
		"disable_recursive_resolution",
		"edns_subnet",
		"enable_dnssec",
		"enable_nsid",
		"enable_path_mtu_discovery",
		"enable_self_versus_third_party_zones",
		"enable_tcp_protocol",
		"enforce_test_failure_if_runs_longer_than",
		"f40x_or_50x_http_mark_successful",
		"failure_hop_count",
		"favor_fastest_round_trip_nameserver",
		"host_data_collection_enabled",
		"ignore_ssl_failures",
		"ping_count",
		"stop_test_on_document_complete",
		"t30x_redirects_do_not_follow",
		"try_next_nameserver_on_failure",
		"verify_test_on_failure",
		"viewport_height",
		"viewport_width",
		"wait_for_no_activity",
		"zone_data_collection_enabled",
	}

	testutil.AssertElementsMatch(t, "AllAdvancedSettings", advancedSettings, expectedKeys)
}

func TestValidateAdvancedSettings(t *testing.T) {
	orig := GetAdvancedSettingsForTestAndMonitorCombinationFunc
	defer func() { GetAdvancedSettingsForTestAndMonitorCombinationFunc = orig }()

	GetAdvancedSettingsForTestAndMonitorCombinationFunc = func(testType types.TestType, monitorType string) []string {
		return validAdvancedSettings
	}

	alertSettings := []map[string]any{
		{"some_advanced_setting": true},
		{"another_advanced_setting": true},
		{"yet_another_advanced_setting": true},
	}

	err := ValidateAdvancedSettingsCombination(alertSettings, fakeTestType, "fakeMonitorType1")
	testutil.AssertNil(t, "ValidateAdvancedSettingsError", err)
}

func TestValidateAdvancedSettingsInvalidMonitor(t *testing.T) {
	orig := GetAdvancedSettingsForTestAndMonitorCombinationFunc
	defer func() { GetAdvancedSettingsForTestAndMonitorCombinationFunc = orig }()

	GetAdvancedSettingsForTestAndMonitorCombinationFunc = func(testType types.TestType, monitorType string) []string {
		return nil
	}

	alertSettings := []map[string]any{
		{"some_advanced_setting": true},
		{"another_advanced_setting": true},
		{"yet_another_advanced_setting": true},
	}

	// This will error because the monitor type is invalid.
	err := ValidateAdvancedSettingsCombination(alertSettings, fakeTestType, "fakeMonitorType1")
	testutil.AssertNotNil(t, "ValidateAdvancedSettingsError", err)
}

func TestValidateAdvancedSettingsInvalidSetting(t *testing.T) {
	orig := GetAdvancedSettingsForTestAndMonitorCombinationFunc
	defer func() { GetAdvancedSettingsForTestAndMonitorCombinationFunc = orig }()

	GetAdvancedSettingsForTestAndMonitorCombinationFunc = func(testType types.TestType, monitorType string) []string {
		return validAdvancedSettings
	}

	alertSettings := []map[string]any{
		{"some_invalid_seting": true},
	}

	// This will error because the advanced setting is not valid.
	err := ValidateAdvancedSettingsCombination(alertSettings, fakeTestType, "fakeMonitorType1")
	testutil.AssertNotNil(t, "ValidateAdvancedSettingsError", err)
}
