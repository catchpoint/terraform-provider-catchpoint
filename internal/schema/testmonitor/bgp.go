package testmonitor

import (
	"context"
	"sync"

	cpschema "catchpoint-provider/internal/schema"
	cptypes "catchpoint-provider/internal/types"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

var (
	testBGPSchemaOnce sync.Once
	testBGPSchema     schema.Schema
)

func BuildBGPTestSchema(ctx context.Context) schema.Schema {
	monitorTypes := []string{cptypes.BGPString, cptypes.BGPBasicString}

	testBGPSchemaOnce.Do(func() {
		testBGPSchema = schema.Schema{
			Description: "Resource for managing BGP Tests in Catchpoint.",
			Attributes: cpschema.MergeAttributes(
				buildCommonTestAttributes(ctx),
				// Monitor is required for BGP because it may be bgp or bgp basic.
				buildMonitorAttribute(ctx, true, cptypes.EmptyString, monitorTypes...),
				// BGP requests the Prefix field which maps to the backend url.
				buildURLField(ctx, cptypes.BGPType),
			),
			Blocks: cpschema.MergeBlocks(
				buildLabelsBlock(),
				cpschema.BuildAlertSettingsBlockForTest(ctx, cptypes.BGPType),
			),
		}
	})
	return testBGPSchema
}
