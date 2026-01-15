package testmonitor

import (
	"context"
	"testing"

	cpschema "catchpoint-provider/internal/schema"
	"catchpoint-provider/internal/testutil"
)

func TestTracerouteTestResourceSchemaContainsExpectedFields(t *testing.T) {
	ctx := context.Background()
	schema := BuildTracerouteTestSchema(ctx)

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

	testutil.AssertElementsMatch(t, "TracerouteTestResourceSchemaAttributes", cpschema.GetAttributeKeys(schema.Attributes), expectedFields)

	expectedBlocks := []string{
		"label",
		"advanced_settings",
		"alert_settings",
		"schedule_settings",
		"thresholds",
	}
	testutil.AssertElementsMatch(t, "TracerouteTestResourceSchemaBlocks", cpschema.GetBlockKeys(schema.Blocks), expectedBlocks)
}

func TestTracerouteTestResourceSchemaContainsExpectedAdvancedSettingsFields(t *testing.T) {
	ctx := context.Background()
	schema := BuildTracerouteTestSchema(ctx)

	expectedFields := []string{
		"advanced_setting_type",
		"enable_path_mtu_discovery",
		"failure_hop_count",
		"ping_count",
	}

	singleNested, ok := cpschema.GetSingleNestedBlockAttributes(schema.Blocks, "advanced_settings")
	if !ok {
		t.Errorf("expected advanced_settings to be present for Traceroute test schema")
	}

	testutil.AssertElementsMatch(t, "AdvancedSettings", cpschema.GetAttributeKeys(singleNested), expectedFields)
}

func TestTracerouteTestResourceSchemaContainsExpectedAlertSettingsFields(t *testing.T) {
	ctx := context.Background()
	schema := BuildTracerouteTestSchema(ctx)

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

func TestTracerouteTestResourceSchemaContainsExpectedScheduleSettingsFields(t *testing.T) {
	ctx := context.Background()
	schema := BuildTracerouteTestSchema(ctx)

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
		t.Errorf("expected schedule_settings to be present for Traceroute test schema")
	}

	testutil.AssertElementsMatch(t, "ScheduleSettings", cpschema.GetAttributeKeys(singleNested), expectedFields)
}
