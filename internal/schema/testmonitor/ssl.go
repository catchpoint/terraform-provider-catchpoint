package testmonitor

import (
	"context"
	"sync"

	"catchpoint-provider/internal/fields"
	cpschema "catchpoint-provider/internal/schema"
	cptypes "catchpoint-provider/internal/types"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var (
	testSSLSchemaOnce sync.Once
	testSSLSchema     schema.Schema
)

func BuildSSLTestSchema(ctx context.Context) schema.Schema {

	testSSLSchemaOnce.Do(func() {
		testSSLSchema = schema.Schema{
			Description: "Resource for managing SSL Tests in Catchpoint.",
			Attributes: cpschema.MergeAttributes(
				buildCommonTestAttributes(ctx),
				// Monitor is not required for SSL because the only valid monitor type is SSL.
				buildMonitorAttribute(ctx, false, cptypes.SSLString, cptypes.SSLString),
				// SSL requests the Location field which maps to the backend url.
				buildURLField(ctx, cptypes.SSLType),
				buildGatewayAddressOrHostBlock(),
				buildSSLTestAttributes(ctx),
			),
			Blocks: cpschema.MergeBlocks(
				buildLabelsBlock(),
				buildThresholdsBlock(),
				cpschema.BuildAdvancedSettingsBlockForTest(ctx, cptypes.SSLType),
				cpschema.BuildScheduleSettingsBlock(ctx),
				cpschema.BuildAlertSettingsBlockForTest(ctx, cptypes.SSLType),
			),
		}
	})
	return testSSLSchema
}

func buildSSLTestAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		fields.EnforceCertificatePinning: schema.BoolAttribute{
			Optional:    true,
			Computed:    true,
			Default:     booldefault.StaticBool(false),
			Description: "Switch for enabling Certificate Pinning feature",
		},
		fields.EnforceCertificateKeyPinning: schema.BoolAttribute{
			Optional:    true,
			Computed:    true,
			Default:     booldefault.StaticBool(false),
			Description: "Switch for enabling Certificate Key Pinning feature",
		},
		fields.FileData: schema.StringAttribute{
			Optional:    true,
			Description: "File data for certificate.",
			Validators: []validator.String{
				stringvalidator.LengthAtLeast(1),
			},
		},
		fields.PassPhrase: schema.StringAttribute{
			Optional:    true,
			Sensitive:   true,
			Description: "Passphrase for certificate.",
			Validators: []validator.String{
				stringvalidator.LengthAtLeast(1),
			},
		},
	}
}
