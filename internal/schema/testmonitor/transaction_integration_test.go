package testmonitor

import (
	"context"
	"testing"

	cpschema "catchpoint-provider/internal/schema"
	"catchpoint-provider/internal/testutil"
)

func TestTransactionTestResourceSchemaContainsExpectedFields(t *testing.T) {
	ctx := context.Background()
	schema := BuildTransactionTestSchema(ctx)

	expectedFields := []string{
		"id",
		"monitor",
		"division_id",
		"product_id",
		"folder_id",
		"test_name",
		"test_description",
		"status",
		"alerts_paused",
		"start_time",
		"end_time",
		"enable_test_data_webhook",
		"test_script",
		"test_script_type",
		"gateway_address_or_host",
		"chrome_version",
		"simulate",
	}

	testutil.AssertElementsMatch(t, "TransactionTestResourceSchemaAttributes", cpschema.GetAttributeKeys(schema.Attributes), expectedFields)

	expectedBlocks := []string{
		"label",
		"advanced_settings",
		"alert_settings",
		"insights",
		"request_settings",
		"schedule_settings",
		"thresholds",
	}
	testutil.AssertElementsMatch(t, "TransactionTestResourceSchemaBlocks", cpschema.GetBlockKeys(schema.Blocks), expectedBlocks)
}

func TestTransactionTestResourceSchemaContainsExpectedAdvancedSettingsFields(t *testing.T) {
	ctx := context.Background()
	schema := BuildTransactionTestSchema(ctx)

	expectedFields := []string{
		"advanced_setting_type",
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

	singleNested, ok := cpschema.GetSingleNestedBlockAttributes(schema.Blocks, "advanced_settings")
	if !ok {
		t.Errorf("expected advanced_settings to be present for Transaction test schema")
	}

	testutil.AssertElementsMatch(t, "AdvancedSettings", cpschema.GetAttributeKeys(singleNested), expectedFields)
}

func TestTransactionTestResourceSchemaContainsExpectedAlertSettingsFields(t *testing.T) {
	ctx := context.Background()
	schema := BuildTransactionTestSchema(ctx)

	expectedAttributes := []string{
		"alert_setting_type",
	}

	expectedBlocks := []string{
		"alert_rule",
		"notification_group",
	}

	singleNested, ok := cpschema.GetSingleNestedBlockAttributes(schema.Blocks, "alert_settings")
	if !ok {
		t.Errorf("expected alert_settings to be present for Product schema")
	}

	testutil.AssertElementsMatch(t, "AlertSettingsAttributes", cpschema.GetAttributeKeys(singleNested), expectedAttributes)

	nestedBlocks, ok := cpschema.GetSingleNestedBlocks(schema.Blocks, "alert_settings")
	if !ok {
		t.Errorf("expected alert_settings to have nested blocks")
	}
	testutil.AssertElementsMatch(t, "AlertSettingsBlocks", cpschema.GetBlockKeys(nestedBlocks), expectedBlocks)
}

func TestTransactionTestResourceSchemaContainsExpectedInsightSettingsFields(t *testing.T) {
	ctx := context.Background()
	schema := BuildTransactionTestSchema(ctx)

	expectedFields := []string{
		"indicator_ids",
		"insight_setting_type",
		"tracepoint_ids",
	}

	singleNested, ok := cpschema.GetSingleNestedBlockAttributes(schema.Blocks, "insights")
	if !ok {
		t.Errorf("expected insights to be present for Transaction test schema")
	}

	testutil.AssertElementsMatch(t, "Insights", cpschema.GetAttributeKeys(singleNested), expectedFields)
}

func TestTransactionTestResourceSchemaContainsExpectedRequestSettingsFields(t *testing.T) {
	ctx := context.Background()
	schema := BuildTransactionTestSchema(ctx)

	expectedFields := []string{
		"library_certificate_ids",
		"request_setting_type",
		"token_ids",
	}

	singleNested, ok := cpschema.GetSingleNestedBlockAttributes(schema.Blocks, "request_settings")
	if !ok {
		t.Errorf("expected request_settings to be present for Transaction test schema")
	}

	testutil.AssertElementsMatch(t, "RequestSettings", cpschema.GetAttributeKeys(singleNested), expectedFields)

	expectedBlocks := []string{
		"http_request_headers",
		"authentication",
	}

	nestedBlocks, ok := cpschema.GetSingleNestedBlocks(schema.Blocks, "request_settings")
	if !ok {
		t.Errorf("expected request_settings to have nested blocks")
	}
	testutil.AssertElementsMatch(t, "RequestSettingsBlocks", cpschema.GetBlockKeys(nestedBlocks), expectedBlocks)
}

func TestTransactionTestResourceSchemaContainsExpectedScheduleSettingsFields(t *testing.T) {
	ctx := context.Background()
	schema := BuildTransactionTestSchema(ctx)

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
		t.Errorf("expected schedule_settings to be present for Transaction test schema")
	}

	testutil.AssertElementsMatch(t, "ScheduleSettings", cpschema.GetAttributeKeys(singleNested), expectedFields)
}
