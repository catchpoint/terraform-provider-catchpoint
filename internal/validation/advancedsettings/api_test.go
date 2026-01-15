package advancedsettings

import (
	"testing"

	"catchpoint-provider/internal/testutil"
	"catchpoint-provider/internal/types"
)

func TestGetAdvancedSettingsForTestTypeAPI(t *testing.T) {
	advancedSettings := GetAdvancedSettingsForTestType(types.APIType)

	expectedKeys := []string{
		"additional_monitor",
		"allow_test_download_limit_override",
		"capture_http_headers",
		"capture_response_content",
		"debug_primary_host_on_failure",
		"debug_referenced_hosts_on_failure",
		"enable_path_mtu_discovery",
		"enable_self_versus_third_party_zones",
		"enforce_test_failure_if_runs_longer_than",
		"f40x_or_50x_http_mark_successful",
		"host_data_collection_enabled",
		"ignore_ssl_failures",
		"t30x_redirects_do_not_follow",
		"verify_test_on_failure",
		"zone_data_collection_enabled",
	}

	testutil.AssertElementsMatch(t, "AdvancedSettingsAPI", advancedSettings, expectedKeys)
}

func TestGetAdvancedSettingsForTestAndMonitorCombinationAPI(t *testing.T) {
	// API tests only have one monitor type, which is API. Selenium/Javascript is set as test_script_type and
	// doesn't affect the advanced settings.
	advancedSettings := GetAdvancedSettingsForTestAndMonitorCombination(types.APIType, "api")

	expectedKeys := []string{
		"additional_monitor",
		"allow_test_download_limit_override",
		"capture_http_headers",
		"capture_response_content",
		"debug_primary_host_on_failure",
		"debug_referenced_hosts_on_failure",
		"enable_path_mtu_discovery",
		"enable_self_versus_third_party_zones",
		"enforce_test_failure_if_runs_longer_than",
		"f40x_or_50x_http_mark_successful",
		"host_data_collection_enabled",
		"ignore_ssl_failures",
		"t30x_redirects_do_not_follow",
		"verify_test_on_failure",
		"zone_data_collection_enabled",
	}

	testutil.AssertElementsMatch(t, "AdvancedSettingsAPI", advancedSettings, expectedKeys)
}
