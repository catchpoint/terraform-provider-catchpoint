package advancedsettings

import (
	"testing"

	"catchpoint-provider/internal/testutil"
	"catchpoint-provider/internal/types"
)

func TestGetAdvancedSettingsForTestTypePuppeteer(t *testing.T) {
	advancedSettings := GetAdvancedSettingsForTestType(types.PuppeteerType)

	expectedKeys := []string{
		"verify_test_on_failure",
		"debug_primary_host_on_failure",
		"debug_referenced_hosts_on_failure",
		"capture_http_headers",
		"capture_response_content",
		"capture_filmstrip",
		"capture_screenshot",
		"ignore_ssl_failures",
		"host_data_collection_enabled",
		"zone_data_collection_enabled",
		"stop_test_on_document_complete",
		"f40x_or_50x_http_mark_successful",
		"enable_self_versus_third_party_zones",
		"allow_test_download_limit_override",
		"disable_cross_origin_iframe_access",
		"enable_path_mtu_discovery",
		"enforce_test_failure_if_runs_longer_than",
		"additional_monitor",
		"viewport_height",
		"viewport_width",
		"bandwidth_throttling",
		"wait_for_no_activity",
	}

	testutil.AssertElementsMatch(t, "AdvancedSettingsPuppeteer", advancedSettings, expectedKeys)
}

func TestGetAdvancedSettingsForTestAndMonitorCombinationPuppeteer(t *testing.T) {
	test := []struct {
		testMonitor string
	}{
		{types.EdgeString},
		{types.ChromeString},
	}
	for _, tt := range test {
		// Puppeteer tests have either Edge or Chrome monitors, and both accept the same AdvancedSettings.
		advancedSettings := GetAdvancedSettingsForTestAndMonitorCombination(types.PuppeteerType, tt.testMonitor)

		expectedKeys := []string{
			"verify_test_on_failure",
			"debug_primary_host_on_failure",
			"debug_referenced_hosts_on_failure",
			"capture_http_headers",
			"capture_response_content",
			"capture_filmstrip",
			"capture_screenshot",
			"ignore_ssl_failures",
			"host_data_collection_enabled",
			"zone_data_collection_enabled",
			"stop_test_on_document_complete",
			"f40x_or_50x_http_mark_successful",
			"enable_self_versus_third_party_zones",
			"allow_test_download_limit_override",
			"disable_cross_origin_iframe_access",
			"enable_path_mtu_discovery",
			"enforce_test_failure_if_runs_longer_than",
			"additional_monitor",
			"viewport_height",
			"viewport_width",
			"bandwidth_throttling",
			"wait_for_no_activity",
		}

		testutil.AssertElementsMatch(t, "AdvancedSettingsPuppeteer_"+tt.testMonitor, advancedSettings, expectedKeys)
	}
}
