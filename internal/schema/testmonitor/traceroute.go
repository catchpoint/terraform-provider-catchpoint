package testmonitor

import (
	"context"
	"sync"

	cpschema "catchpoint-provider/internal/schema"
	cptypes "catchpoint-provider/internal/types"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

var (
	testTracerouteSchemaOnce sync.Once
	testTracerouteSchema     schema.Schema
)

func BuildTracerouteTestSchema(ctx context.Context) schema.Schema {
	monitorTypes := []string{cptypes.TracerouteICMPString, cptypes.TracerouteUDPString, cptypes.TracerouteTCPString}
	testTracerouteSchemaOnce.Do(func() {
		testTracerouteSchema = schema.Schema{
			Description: "Resource for managing Traceroute Tests in Catchpoint.",
			Attributes: cpschema.MergeAttributes(
				buildCommonTestAttributes(ctx),
				// Monitor is required and must be one of the 3 Traceroute types.
				buildMonitorAttribute(ctx, true, cptypes.EmptyString, monitorTypes...),
				// Traceroute requests the Location field which maps to the backend url.
				buildURLField(ctx, cptypes.TracerouteType),
			),
			Blocks: cpschema.MergeBlocks(
				buildLabelsBlock(),
				buildThresholdsBlock(),
				cpschema.BuildAlertSettingsBlockForTest(ctx, cptypes.TracerouteType),
				cpschema.BuildScheduleSettingsBlock(ctx),
				cpschema.BuildAdvancedSettingsBlockForTest(ctx, cptypes.TracerouteType),
			),
		}
	})
	return testTracerouteSchema
}
