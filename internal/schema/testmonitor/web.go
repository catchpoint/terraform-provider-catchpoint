package testmonitor

import (
	"context"
	"sync"

	cpschema "catchpoint-provider/internal/schema"
	cptypes "catchpoint-provider/internal/types"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

var (
	testWebSchemaOnce sync.Once
	testWebSchema     schema.Schema
)

func BuildWebTestSchema(ctx context.Context) schema.Schema {
	monitorTypes := []string{cptypes.HTTPString, cptypes.ChromeString, cptypes.EmulatedString, cptypes.MobileString, cptypes.PlaybackString, cptypes.MobilePlaybackString}

	testWebSchemaOnce.Do(func() {
		testWebSchema = schema.Schema{
			Description: "Resource for managing Web Tests in Catchpoint.",
			Attributes: cpschema.MergeAttributes(
				buildCommonTestAttributes(ctx),
				// Web monitors are either HTTP, Chrome, Emulated, Mobile, Playback or Mobile Playback.
				buildMonitorAttribute(ctx, true, cptypes.EmptyString, monitorTypes...),
				// Web requests use the 'test_url' field which maps to the backend 'url'.
				buildURLField(ctx, cptypes.WebType),
				// Web tests have the optional GatewayAddressOrHost field.
				buildGatewayAddressOrHostBlock(),
				// Web tests have the optional ChromeVersion and Simulate fields.
				buildChromeSimulateAttributes(),
				// TODO: "Request Type" of "GET" or "POST" are valid fields for HTTP, Emulated, and Chrome monitor types.
				// However, this field was not available in the previous version of the provider, so we will add it later.
			),
			Blocks: cpschema.MergeBlocks(
				buildLabelsBlock(),
				buildThresholdsBlock(),
				cpschema.BuildRequestSettingsBlock(ctx),
				cpschema.BuildAlertSettingsBlockForTest(ctx, cptypes.WebType),
				cpschema.BuildInsightsBlock(ctx),
				cpschema.BuildScheduleSettingsBlock(ctx),
				cpschema.BuildAdvancedSettingsBlockForTest(ctx, cptypes.WebType),
			),
		}
	})
	return testWebSchema
}
