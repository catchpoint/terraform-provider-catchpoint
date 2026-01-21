package schema

import (
	"catchpoint-provider/internal/testutil"
	"catchpoint-provider/internal/types"
	"context"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

func TestBuildAdvancedSettingsAttributesForProductContainsExpectedAttributes(t *testing.T) {
	ctx := context.Background()
	attrs := BuildAdvancedSettingsAttributesForProductAndFolder(ctx)
	keys := getAttributeKeys(t, attrs)

	// These should match the merged settings in buildAdvancedSettingsAttributesForProduct
	expectedKeys := []string{
		"advanced_setting_type",
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

	testutil.AssertElementsMatch(t, "ProductAdvancedSettings", keys, expectedKeys)
}

func TestBuildAdvancedSettingsAttributesForAPITestContainsExpectedAttributes(t *testing.T) {
	ctx := context.Background()
	attrs := BuildAdvancedSettingsBlockForTest(ctx, types.APIType)
	keys := getAttributeKeys(t, attrs)

	expectedKeys := []string{
		"advanced_setting_type",
		"f40x_or_50x_http_mark_successful",
		"additional_monitor",
		"debug_primary_host_on_failure",
		"debug_referenced_hosts_on_failure",
		"enable_path_mtu_discovery",
		"enforce_test_failure_if_runs_longer_than",
		"capture_http_headers",
		"capture_response_content",
		"host_data_collection_enabled",
		"ignore_ssl_failures",
		"allow_test_download_limit_override",
		"enable_self_versus_third_party_zones",
		"verify_test_on_failure",
		"zone_data_collection_enabled",
		"t30x_redirects_do_not_follow",
	}

	testutil.AssertElementsMatch(t, "APITestAdvancedSettings", keys, expectedKeys)
}

func TestBuildAdvancedSettingsAttributesForSSLTestContainsExpectedAttributes(t *testing.T) {
	ctx := context.Background()
	attrs := BuildAdvancedSettingsBlockForTest(ctx, types.SSLType)
	keys := getAttributeKeys(t, attrs)

	// These should match the merged settings in buildAdvancedSettingsAttributesForSSLTest
	expectedKeys := []string{
		"advanced_setting_type",
		"additional_monitor",
		"certificate_revocation_disabled",
		"enable_path_mtu_discovery",
		"verify_test_on_failure",
	}

	testutil.AssertElementsMatch(t, "SSLTestAdvancedSettings", keys, expectedKeys)
}

func TestBuildAdvancedSettingsAttributesForTracerouteTestContainsExpectedAttributes(t *testing.T) {
	ctx := context.Background()
	attrs := BuildAdvancedSettingsBlockForTest(ctx, types.TracerouteType)
	keys := getAttributeKeys(t, attrs)

	// These should match the tracerouteAdvancedSettings
	expectedKeys := []string{
		"advanced_setting_type",
		// TODO: not yet available "enable_dscp_priority_protocol",
		"enable_path_mtu_discovery",
		"ping_count",
		"failure_hop_count",
	}

	testutil.AssertElementsMatch(t, "TracerouteTestAdvancedSettings", keys, expectedKeys)
}

func TestBuildAdvancedSettingsAttributesForBGPTestContainsExpectedAttributes(t *testing.T) {
	ctx := context.Background()
	attrs := BuildAdvancedSettingsBlockForTest(ctx, types.BGPType)
	keys := getAttributeKeys(t, attrs)

	// BGP has no advanced settings currently.
	expectedKeys := []string{}

	testutil.AssertElementsMatch(t, "BGPTestAdvancedSettings", keys, expectedKeys)
}

func TestBuildAdvancedSettingsAttributesForWebTestContainsExpectedAttributes(t *testing.T) {
	ctx := context.Background()
	attrs := BuildAdvancedSettingsBlockForTest(ctx, types.WebType)
	keys := getAttributeKeys(t, attrs)

	// These should match the merged settings in buildAdvancedSettingsAttributesForWebTest
	expectedKeys := []string{
		"advanced_setting_type",
		// webChromeAdvancedSettings
		"f40x_or_50x_http_mark_successful",
		"additional_monitor",
		"debug_primary_host_on_failure",
		"debug_referenced_hosts_on_failure",
		"enable_path_mtu_discovery",
		"enforce_test_failure_if_runs_longer_than",
		"capture_http_headers",
		"capture_response_content",
		"host_data_collection_enabled",
		"ignore_ssl_failures",
		"allow_test_download_limit_override",
		"enable_self_versus_third_party_zones",
		"verify_test_on_failure",
		"zone_data_collection_enabled",
		"capture_filmstrip",
		"viewport_height",
		"viewport_width",
		"capture_screenshot",
		"bandwidth_throttling",
		"disable_cross_origin_iframe_access",
		"stop_test_on_document_complete",
		// webEmulatedHttpAdvancedSettings
		"t30x_redirects_do_not_follow",
		"wait_for_no_activity",
		// webMobileAdvancedSettings
		// (already included above, but bandwidth_throttling and stop_test_on_document_complete should be present)
		// webPlaybackAdvancedSettings
		// (no additional keys, just commonBrowserAdvancedSettings)
	}

	testutil.AssertElementsMatch(t, "WebTestAdvancedSettings", keys, expectedKeys)
}

// Helper to get attribute keys from the returned map
func getAttributeKeys(t *testing.T, attrs map[string]schema.Block) []string {
	advancedSettingsAttr, ok := attrs["advanced_settings"].(schema.SingleNestedBlock)
	if !ok {
		t.Fatalf("advanced_settings is not a SingleNestedBlock")
	}

	keys := make([]string, 0, len(advancedSettingsAttr.Attributes))
	for k := range advancedSettingsAttr.Attributes {
		keys = append(keys, k)
	}
	return keys
}
