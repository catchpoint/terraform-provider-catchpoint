package folder

import (
	"context"
	"testing"

	cpschema "catchpoint-provider/internal/schema"
	"catchpoint-provider/internal/testutil"
)

func TestFolderResourceSchemaContainsExpectedFields(t *testing.T) {
	ctx := context.Background()
	schema := BuildFolderSchema(ctx)

	expectedFields := []string{
		"id",
		"product_id",
		"division_id",
		"folder_name",
		"parent_id",
	}

	testutil.AssertElementsMatch(t, "FolderResourceSchema", cpschema.GetAttributeKeys(schema.Attributes), expectedFields)
}

func TestFolderResourceSchemaContainsExpectedAdvancedSettingsFields(t *testing.T) {
	ctx := context.Background()
	schema := BuildFolderSchema(ctx)

	expectedFields := []string{
		"additional_monitor",
		"advanced_setting_type",
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

	singleNested, ok := cpschema.GetSingleNestedBlockAttributes(schema.Blocks, "advanced_settings")
	if !ok {
		t.Errorf("expected advanced_settings to be present for Folder schema")
	}

	testutil.AssertElementsMatch(t, "AdvancedSettings", cpschema.GetAttributeKeys(singleNested), expectedFields)
}

func TestFolderResourceSchemaContainsExpectedAlertSettingsFields(t *testing.T) {
	ctx := context.Background()
	schema := BuildFolderSchema(ctx)

	expectedAttributes := []string{
		"alert_setting_type",
	}

	expectedBlocks := []string{
		"alert_rule",
		"notification_group",
	}

	singleNested, ok := cpschema.GetSingleNestedBlockAttributes(schema.Blocks, "alert_settings")
	if !ok {
		t.Errorf("expected alert_settings to be present for Folder schema")
	}

	testutil.AssertElementsMatch(t, "AlertSettingsAttributes", cpschema.GetAttributeKeys(singleNested), expectedAttributes)

	nestedBlocks, ok := cpschema.GetSingleNestedBlocks(schema.Blocks, "alert_settings")
	if !ok {
		t.Errorf("expected alert_settings to have nested blocks")
	}
	testutil.AssertElementsMatch(t, "AlertSettingsBlocks", cpschema.GetBlockKeys(nestedBlocks), expectedBlocks)
}

func TestFolderResourceSchemaContainsExpectedInsightSettingsFields(t *testing.T) {
	ctx := context.Background()
	schema := BuildFolderSchema(ctx)

	expectedFields := []string{
		"indicator_ids",
		"insight_setting_type",
		"tracepoint_ids",
	}

	singleNested, ok := cpschema.GetSingleNestedBlockAttributes(schema.Blocks, "insights")
	if !ok {
		t.Errorf("expected insights to be present for Folder schema")
	}

	testutil.AssertElementsMatch(t, "Insights", cpschema.GetAttributeKeys(singleNested), expectedFields)
}

func TestFolderResourceSchemaContainsExpectedRequestSettingsFields(t *testing.T) {
	ctx := context.Background()
	schema := BuildFolderSchema(ctx)

	expectedFields := []string{
		"library_certificate_ids",
		"request_setting_type",
		"token_ids",
	}

	singleNested, ok := cpschema.GetSingleNestedBlockAttributes(schema.Blocks, "request_settings")
	if !ok {
		t.Errorf("expected request_settings to be present for Folder schema")
	}

	testutil.AssertElementsMatch(t, "RequestSettings", cpschema.GetAttributeKeys(singleNested), expectedFields)

	expectedBlocks := []string{
		"authentication",
		"http_request_headers",
	}

	nestedBlocks, ok := cpschema.GetSingleNestedBlocks(schema.Blocks, "request_settings")
	if !ok {
		t.Errorf("expected request_settings to have nested blocks")
	}
	testutil.AssertElementsMatch(t, "RequestSettingsBlocks", cpschema.GetBlockKeys(nestedBlocks), expectedBlocks)
}

func TestFolderResourceSchemaContainsExpectedScheduleSettingsFields(t *testing.T) {
	ctx := context.Background()
	schema := BuildFolderSchema(ctx)

	expectedFields := []string{
		"frequency",
		"maintenance_schedule_id",
		"no_of_subset_nodes",
		"node_distribution",
		"node_group_ids",
		"node_ids",
		"run_schedule_id",
		"schedule_setting_type",
	}

	singleNested, ok := cpschema.GetSingleNestedBlockAttributes(schema.Blocks, "schedule_settings")
	if !ok {
		t.Errorf("expected schedule_settings to be present for Folder schema")
	}

	testutil.AssertElementsMatch(t, "ScheduleSettings", cpschema.GetAttributeKeys(singleNested), expectedFields)
}
