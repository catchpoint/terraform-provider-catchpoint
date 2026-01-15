package product

import (
	"context"

	"catchpoint-provider/internal/fields"
	cpschema "catchpoint-provider/internal/schema"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

// Build the Schema for a Product resource (manage_product).
func BuildProductSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Description: "Resource for managing Products in Catchpoint. Note: Products cannot be deleted via the API. Use 'terraform state rm' to remove from state without deleting the actual product.",
		Attributes: cpschema.MergeAttributes(
			buildProductAttributes(ctx),
		),
		Blocks: cpschema.MergeBlocks(
			cpschema.BuildAdvancedSettingsAttributesForProductAndFolder(ctx),
			cpschema.BuildRequestSettingsBlock(ctx),
			cpschema.BuildAlertSettingsBlockForProductAndFolder(ctx),
			cpschema.BuildInsightsBlock(ctx),
			cpschema.BuildScheduleSettingsBlock(ctx),
		),
	}
}

func buildProductAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		fields.ID: schema.Int64Attribute{
			Computed:    true,
			Required:    false,
			Optional:    false,
			Description: "The internal ID of the Product.",
		},
		fields.DivisionID: schema.Int64Attribute{
			Required:    true,
			Description: "The Division where the Product will be created",
		},
		fields.ProductName: schema.StringAttribute{
			Required:    true,
			Description: "The name of the Product",
		},
		fields.Status: schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "Set the Product status, " + cpschema.StatusValidator.Description(ctx),
			Validators:  []validator.String{cpschema.StatusValidator},
		},
		fields.TestDataWebhookID: schema.Int64Attribute{
			Optional:    true,
			Computed:    true,
			Description: "The ID of the Test Data Webhook to use for the Product",
		},
		// Note: This does not appear to be visible from the portal, but is used in the API.
		fields.AlertGroupID: schema.Int64Attribute{
			Optional:    true,
			Computed:    true,
			Description: "The ID of the Alert Group to use for the Product",
		},
	}
}
