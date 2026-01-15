package testmonitor

import (
	"context"
	"sync"

	cpschema "catchpoint-provider/internal/schema"
	cptypes "catchpoint-provider/internal/types"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

var (
	testTransactionSchemaOnce sync.Once
	testTransactionSchema     schema.Schema
)

func BuildTransactionTestSchema(ctx context.Context) schema.Schema {
	monitorTypes := []string{cptypes.ChromeString, cptypes.EmulatedString, cptypes.MobileString}

	testTransactionSchemaOnce.Do(func() {
		testTransactionSchema = schema.Schema{
			Description: "Resource for managing Transaction Tests in Catchpoint.",
			Attributes: cpschema.MergeAttributes(
				buildCommonTestAttributes(ctx),
				// Monitor is not required for Transaction tests because the only valid monitor type is Transaction.
				buildMonitorAttribute(ctx, true, cptypes.EmptyString, monitorTypes...),
				// Transaction requests the testRequestData info instead of using an URL of some sort.
				buildTestRequestDataAttributes(ctx, cptypes.TransactionType),
				// Transaction tests have the optional ChromeVersion and Simulate fields.
				buildChromeSimulateAttributes(),
				// Transaction tests have the optional GatewayAddressOrHost field.
				buildGatewayAddressOrHostBlock(),
			),
			Blocks: cpschema.MergeBlocks(
				buildLabelsBlock(),
				buildThresholdsBlock(),
				cpschema.BuildRequestSettingsBlock(ctx),
				cpschema.BuildAlertSettingsBlockForTest(ctx, cptypes.TransactionType),
				cpschema.BuildInsightsBlock(ctx),
				cpschema.BuildScheduleSettingsBlock(ctx),
				cpschema.BuildAdvancedSettingsBlockForTest(ctx, cptypes.TransactionType),
			),
		}
	})
	return testTransactionSchema
}
