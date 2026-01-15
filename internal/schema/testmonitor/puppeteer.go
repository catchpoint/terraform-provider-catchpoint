package testmonitor

import (
	"context"
	"sync"

	cpschema "catchpoint-provider/internal/schema"
	cptypes "catchpoint-provider/internal/types"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

var (
	testPuppeteerSchemaOnce sync.Once
	testPuppeteerSchema     schema.Schema
)

func BuildPuppeteerTestSchema(ctx context.Context) schema.Schema {
	testPuppeteerSchemaOnce.Do(func() {
		testPuppeteerSchema = schema.Schema{
			Description: "Resource for managing Puppeteer Tests in Catchpoint.",
			Attributes: cpschema.MergeAttributes(
				buildCommonTestAttributes(ctx),
				// Puppeteer monitor only accepts Chrome.
				buildMonitorAttribute(ctx, false, cptypes.ChromeString, cptypes.ChromeString),
				// Puppeteer requests the testRequestData info instead of using an URL of some sort.
				buildTestRequestDataAttributes(ctx, cptypes.PuppeteerType),
				// Puppeteer tests have the optional GatewayAddressOrHost field.
				buildGatewayAddressOrHostBlock(),
			),
			Blocks: cpschema.MergeBlocks(
				buildLabelsBlock(),
				buildThresholdsBlock(),
				cpschema.BuildRequestSettingsBlock(ctx),
				cpschema.BuildAlertSettingsBlockForTest(ctx, cptypes.PuppeteerType),
				cpschema.BuildInsightsBlock(ctx),
				cpschema.BuildScheduleSettingsBlock(ctx),
				cpschema.BuildAdvancedSettingsBlockForTest(ctx, cptypes.PuppeteerType),
			),
		}
	})
	return testPuppeteerSchema
}
