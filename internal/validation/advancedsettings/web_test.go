package advancedsettings

import (
	"testing"

	"catchpoint-provider/internal/testutil"
	"catchpoint-provider/internal/types"
)

// All Web tests have *at least* these advanced settings available. Playback/Mobile-Playback have *only* these.
var validBaseWebAdvancedSettings = []string{
	"f40x_or_50x_http_mark_successful",
	"additional_monitor",
	"debug_primary_host_on_failure",
	"debug_referenced_hosts_on_failure",
	"enable_path_mtu_discovery",
	"enforce_test_failure_if_runs_longer_than",
	"host_data_collection_enabled",
	"capture_http_headers",
	"capture_response_content",
	"ignore_ssl_failures",
	"allow_test_download_limit_override",
	"enable_self_versus_third_party_zones",
	"verify_test_on_failure",
	"zone_data_collection_enabled",
}

// Mobile has all the base settings plus these.
var validWebMobileAdvancedSettings = append([]string{
	"bandwidth_throttling",
	"capture_filmstrip",
	"capture_screenshot",
	"viewport_width",
	"viewport_height",
	"stop_test_on_document_complete",
	"wait_for_no_activity",
}, validBaseWebAdvancedSettings...)

// Emulated and HTTP have all the base settings plus these.
var validWebEmulatedHTTPAdvancedSettings = append([]string{
	"t30x_redirects_do_not_follow",
	// "http_version", // TODO: add when supported
}, validBaseWebAdvancedSettings...)

// Chrome has all the base settings plus these.
var validWebChromeAdvancedSettings = append([]string{
	"bandwidth_throttling",
	"disable_cross_origin_iframe_access",
	"capture_filmstrip",
	"capture_screenshot",
	"viewport_width",
	"viewport_height",
	"stop_test_on_document_complete",
	"wait_for_no_activity",
}, validBaseWebAdvancedSettings...)

func TestGetAdvancedSettingsForTestTypeWeb(t *testing.T) {
	advancedSettings := GetAdvancedSettingsForTestType(types.WebType)

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
		"capture_filmstrip",
		"capture_screenshot",
		"viewport_width",
		"viewport_height",
		"stop_test_on_document_complete",
		"wait_for_no_activity",
		"bandwidth_throttling",
		"disable_cross_origin_iframe_access",
	}

	testutil.AssertElementsMatch(t, "AdvancedSettingsWeb", advancedSettings, expectedKeys)
}

func TestGetAdvancedSettingsForTestAndMonitorCombinationWeb(t *testing.T) {
	test := []struct {
		monitor  string
		settings []string
	}{
		{types.ChromeString, validWebChromeAdvancedSettings},
		{types.EmulatedString, validWebEmulatedHTTPAdvancedSettings},
		{types.HTTPString, validWebEmulatedHTTPAdvancedSettings},
		{types.MobileString, validWebMobileAdvancedSettings},
		{types.MobilePlaybackString, validBaseWebAdvancedSettings},
		{types.PlaybackString, validBaseWebAdvancedSettings},
	}
	for _, tt := range test {
		// Web tests have three monitor types: Chrome, Emulated, Http, Mobile, Mobile Playback, and Playback.
		advancedSettings := GetAdvancedSettingsForTestAndMonitorCombination(types.WebType, tt.monitor)

		testutil.AssertElementsMatch(t, "AdvancedSettingsWeb_"+tt.monitor, advancedSettings, tt.settings)
	}
}
