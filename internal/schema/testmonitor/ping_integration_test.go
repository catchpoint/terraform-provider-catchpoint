package testmonitor

import (
	"context"
	"testing"

	cpschema "catchpoint-provider/internal/schema"
	"catchpoint-provider/internal/testutil"
)

func TestPingTestResourceSchemaContainsExpectedFields(t *testing.T) {
	ctx := context.Background()
	schema := BuildPingTestSchema(ctx)

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
		"test_location",
	}

	testutil.AssertElementsMatch(t, "PingTestResourceSchemaAttributes", cpschema.GetAttributeKeys(schema.Attributes), expectedFields)

	expectedBlocks := []string{
		"label",
		"advanced_settings",
		"alert_settings",
		"schedule_settings",
		"thresholds",
	}
	testutil.AssertElementsMatch(t, "PingTestResourceSchemaBlocks", cpschema.GetBlockKeys(schema.Blocks), expectedBlocks)
}

func TestPingTestResourceSchemaContainsExpectedAdvancedSettingsFields(t *testing.T) {
	ctx := context.Background()
	schema := BuildPingTestSchema(ctx)

	expectedFields := []string{
		"additional_monitor",
		"advanced_setting_type",
		"enable_path_mtu_discovery",
		"debug_primary_host_on_failure",
		"verify_test_on_failure",
	}

	singleNested, ok := cpschema.GetSingleNestedBlockAttributes(schema.Blocks, "advanced_settings")
	if !ok {
		t.Errorf("expected advanced_settings to be present for Ping test schema")
	}

	testutil.AssertElementsMatch(t, "AdvancedSettings", cpschema.GetAttributeKeys(singleNested), expectedFields)
}

func TestPingTestResourceSchemaContainsExpectedAlertSettingsFields(t *testing.T) {
	ctx := context.Background()
	schema := BuildPingTestSchema(ctx)

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

func TestPingTestResourceSchemaContainsExpectedScheduleSettingsFields(t *testing.T) {
	ctx := context.Background()
	schema := BuildPingTestSchema(ctx)

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
		t.Errorf("expected schedule_settings to be present for Ping test schema")
	}

	testutil.AssertElementsMatch(t, "ScheduleSettings", cpschema.GetAttributeKeys(singleNested), expectedFields)
}
