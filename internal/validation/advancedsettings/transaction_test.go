package advancedsettings

import (
	"testing"

	"catchpoint-provider/internal/testutil"
	"catchpoint-provider/internal/types"
)

// All Transaction tests have *at least* these advanced settings available.-
var validBaseTransactionAdvancedSettings = []string{
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
var validTransactionMobileAdvancedSettings = append([]string{
	"bandwidth_throttling",
	"capture_filmstrip",
	"capture_screenshot",
	"viewport_width",
	"viewport_height",
}, validBaseTransactionAdvancedSettings...)

// Emulated has all the base settings plus these.
var validTransactionEmulatedHTTPAdvancedSettings = append([]string{
	"t30x_redirects_do_not_follow",
}, validBaseTransactionAdvancedSettings...)

// Chrome has all the base settings plus these.
var validTransactionChromeAdvancedSettings = append([]string{
	"bandwidth_throttling",
	"disable_cross_origin_iframe_access",
	"capture_filmstrip",
	"capture_screenshot",
	"viewport_width",
	"viewport_height",
}, validBaseTransactionAdvancedSettings...)

func TestGetAdvancedSettingsForTestTypeTransaction(t *testing.T) {
	advancedSettings := GetAdvancedSettingsForTestType(types.TransactionType)

	expectedKeys := []string{
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
		"bandwidth_throttling",
		"disable_cross_origin_iframe_access",
		"capture_filmstrip",
		"capture_screenshot",
		"viewport_width",
		"viewport_height",
		"t30x_redirects_do_not_follow",
	}

	testutil.AssertElementsMatch(t, "AdvancedSettingsTransaction", advancedSettings, expectedKeys)
}

func TestGetAdvancedSettingsForTestAndMonitorCombinationTransaction(t *testing.T) {
	test := []struct {
		monitor  string
		settings []string
	}{
		{types.ChromeString, validTransactionChromeAdvancedSettings},
		{types.EmulatedString, validTransactionEmulatedHTTPAdvancedSettings},
		{types.MobileString, validTransactionMobileAdvancedSettings},
	}
	for _, tt := range test {
		// Transaction tests have three monitor types: Chrome, Emulated, and Mobile.
		advancedSettings := GetAdvancedSettingsForTestAndMonitorCombination(types.TransactionType, tt.monitor)

		testutil.AssertElementsMatch(t, "AdvancedSettingsTransaction_"+tt.monitor, advancedSettings, tt.settings)
	}
}
