package testmonitor

import (
	"context"
	"sync"

	"catchpoint-provider/internal/fields"
	cpschema "catchpoint-provider/internal/schema"
	"catchpoint-provider/internal/types"
	cptypes "catchpoint-provider/internal/types"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var (
	testDNSSchemaOnce  sync.Once
	testDNSSchema      schema.Schema
	queryTypeValidator = oneOfStringValidator(types.ValidDNSQueryTypes)
)

func BuildDNSTestSchema(ctx context.Context) schema.Schema {
	monitorTypes := []string{cptypes.DNSExperienceString, cptypes.DNSDirectString}

	testDNSSchemaOnce.Do(func() {
		testDNSSchema = schema.Schema{
			Description: "Resource for managing DNS Tests in Catchpoint.",
			Attributes: cpschema.MergeAttributes(
				buildCommonTestAttributes(ctx),
				// DNS monitor is either Direct or Experience, no default.
				buildMonitorAttribute(ctx, true, cptypes.EmptyString, monitorTypes...),
				// DNS tests use the TestDomain field.
				buildURLField(ctx, cptypes.DNSType),
				// Query Type/DNS Server.
				buildDNSTestAttributes(ctx),
			),
			Blocks: cpschema.MergeBlocks(
				buildLabelsBlock(),
				buildThresholdsBlock(),
				cpschema.BuildAlertSettingsBlockForTest(ctx, cptypes.DNSType),
				cpschema.BuildScheduleSettingsBlock(ctx),
				cpschema.BuildAdvancedSettingsBlockForTest(ctx, cptypes.DNSType),
			),
		}
	})

	return testDNSSchema
}

func buildDNSTestAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		fields.QueryType: schema.StringAttribute{
			Required:    true,
			Description: "The type of DNS query, " + queryTypeValidator.Description(ctx),
			Validators:  []validator.String{queryTypeValidator},
		},
		// Not a valid field for DNS Experience tests, only for DNS Direct. But we can't conditionally require it.
		// Instead, our DNS validator will ensure it's only set for DNS Direct tests.
		fields.DNSServer: schema.StringAttribute{
			Required:    false,
			Optional:    true,
			Description: "IP address or host name. If empty, uses node's resolver. For DNS Direct monitor only.",
		},
	}
}
