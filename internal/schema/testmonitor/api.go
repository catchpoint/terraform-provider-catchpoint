package testmonitor

import (
	"context"
	"sync"

	cpschema "catchpoint-provider/internal/schema"
	cptypes "catchpoint-provider/internal/types"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

var (
	testAPISchemaOnce sync.Once
	testAPISchema     schema.Schema
)

func BuildAPITestSchema(ctx context.Context) schema.Schema {
	testAPISchemaOnce.Do(func() {
		testAPISchema = schema.Schema{
			Description: "Resource for managing API Tests in Catchpoint.",
			Attributes: cpschema.MergeAttributes(
				buildCommonTestAttributes(ctx),
				// Monitor is not required for API tests because the only valid monitor type is API.
				buildMonitorAttribute(ctx, false, cptypes.APIString, cptypes.APIString),
				// API requests the testRequestData info instead of using an URL of some sort.
				buildTestRequestDataAttributes(ctx, cptypes.APIType),
				// API tests have the optional GatewayAddressOrHost field.
				buildGatewayAddressOrHostBlock(),
			),
			Blocks: cpschema.MergeBlocks(
				buildLabelsBlock(),
				buildThresholdsBlock(),
				cpschema.BuildRequestSettingsBlock(ctx),
				cpschema.BuildAlertSettingsBlockForTest(ctx, cptypes.APIType),
				cpschema.BuildInsightsBlock(ctx),
				cpschema.BuildScheduleSettingsBlock(ctx),
				cpschema.BuildAdvancedSettingsBlockForTest(ctx, cptypes.APIType),
			),
		}
	})
	return testAPISchema
}
