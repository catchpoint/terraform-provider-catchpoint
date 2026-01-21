package folder

import (
	"context"

	"catchpoint-provider/internal/fields"
	cpschema "catchpoint-provider/internal/schema"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
)

// Build the Schema for a Folder resource (manage_folder).
func BuildFolderSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Description: "Resource for managing Folders in Catchpoint. Note: Folders cannot be deleted via the API. Use 'terraform state rm' to remove from state without deleting the actual folder.",
		Attributes: cpschema.MergeAttributes(
			buildFolderAttributes(),
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

func buildFolderAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		fields.ID: schema.Int64Attribute{
			Computed:    true,
			Required:    false,
			Optional:    false,
			Description: "The internal ID of the Folder.",
		},
		fields.DivisionID: schema.Int64Attribute{
			Required:    true,
			Description: "The Division where the Folder will be created",
		},
		fields.ProductID: schema.Int64Attribute{
			Required:    true,
			Description: "The parent Product under which the Folder will be created",
			PlanModifiers: []planmodifier.Int64{
				int64planmodifier.UseStateForUnknown(),
			},
		},
		fields.FolderName: schema.StringAttribute{
			Required:    true,
			Description: "The name of the Folder",
		},
		// Optional fields:
		fields.ParentID: schema.Int64Attribute{
			Optional:    true,
			Description: "The Folder under which this Folder will be created",
			PlanModifiers: []planmodifier.Int64{
				int64planmodifier.UseStateForUnknown(),
			},
		},
	}
}
