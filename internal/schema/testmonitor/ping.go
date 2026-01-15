package testmonitor

import (
	"context"
	"sync"

	cpschema "catchpoint-provider/internal/schema"
	cptypes "catchpoint-provider/internal/types"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

var (
	testPingSchemaOnce sync.Once
	testPingSchema     schema.Schema
)

func BuildPingTestSchema(ctx context.Context) schema.Schema {
	monitorTypes := []string{cptypes.PingICMPString, cptypes.PingUDPString, cptypes.PingTCPString}
	testPingSchemaOnce.Do(func() {
		testPingSchema = schema.Schema{
			Description: "Resource for managing Ping Tests in Catchpoint.",
			Attributes: cpschema.MergeAttributes(
				buildCommonTestAttributes(ctx),
				// Monitor is required and must be one of the 3 Ping types.
				buildMonitorAttribute(ctx, true, cptypes.EmptyString, monitorTypes...),
				// Ping requests the Location field which maps to the backend url.
				buildURLField(ctx, cptypes.PingType),
			),
			Blocks: cpschema.MergeBlocks(
				buildLabelsBlock(),
				buildThresholdsBlock(),
				cpschema.BuildAlertSettingsBlockForTest(ctx, cptypes.PingType),
				cpschema.BuildScheduleSettingsBlock(ctx),
				cpschema.BuildAdvancedSettingsBlockForTest(ctx, cptypes.PingType),
			),
		}
	})
	return testPingSchema
}
