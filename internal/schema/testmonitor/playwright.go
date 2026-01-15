package testmonitor

import (
	"context"
	"sync"

	cpschema "catchpoint-provider/internal/schema"
	cptypes "catchpoint-provider/internal/types"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

var (
	testPlaywrightSchemaOnce sync.Once
	testPlaywrightSchema     schema.Schema
)

func BuildPlaywrightTestSchema(ctx context.Context) schema.Schema {
	monitorTypes := []string{cptypes.EdgeString, cptypes.ChromeString}

	testPlaywrightSchemaOnce.Do(func() {
		testPlaywrightSchema = schema.Schema{
			Description: "Resource for managing Playwright Tests in Catchpoint.",
			Attributes: cpschema.MergeAttributes(
				buildCommonTestAttributes(ctx),
				// Playwright monitor is either Edge or Chrome, no default.
				buildMonitorAttribute(ctx, true, cptypes.EmptyString, monitorTypes...),
				// Playwright requests the testRequestData info instead of using an URL of some sort.
				buildTestRequestDataAttributes(ctx, cptypes.PlaywrightType),
				// Playwright tests have the optional GatewayAddressOrHost field.
				buildGatewayAddressOrHostBlock(),
			),
			Blocks: cpschema.MergeBlocks(
				buildLabelsBlock(),
				buildThresholdsBlock(),
				cpschema.BuildRequestSettingsBlock(ctx),
				cpschema.BuildAlertSettingsBlockForTest(ctx, cptypes.PlaywrightType),
				cpschema.BuildInsightsBlock(ctx),
				cpschema.BuildScheduleSettingsBlock(ctx),
				cpschema.BuildAdvancedSettingsBlockForTest(ctx, cptypes.PlaywrightType),
			),
		}
	})
	return testPlaywrightSchema
}
