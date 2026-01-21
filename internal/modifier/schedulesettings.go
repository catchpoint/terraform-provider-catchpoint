package modifier

import (
	"context"

	"catchpoint-provider/internal/fields"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Create a plan modifier for schedule validation
func ScheduleConfigurationModifier() planmodifier.Object {
	return &scheduleConfigValidator{}
}

type scheduleConfigValidator struct{}

func (v *scheduleConfigValidator) Description(ctx context.Context) string {
	return "Validates that either node_ids or node_group_ids is specified"
}

func (v *scheduleConfigValidator) MarkdownDescription(ctx context.Context) string {
	return "Validates that either node_ids or node_group_ids is specified"
}

func (v *scheduleConfigValidator) PlanModifyObject(ctx context.Context, req planmodifier.ObjectRequest, resp *planmodifier.ObjectResponse) {
	// This runs during planning when variables are resolved
	if req.PlanValue.IsNull() || req.PlanValue.IsUnknown() {
		return
	}

	scheduleAttrs := req.PlanValue.Attributes()

	settingType := scheduleAttrs[fields.ScheduleSettingType]
	if !settingType.IsNull() && !settingType.IsUnknown() && settingType.(types.String).ValueString() == "inherit" {
		return
	}

	nodeIDs := scheduleAttrs[fields.NodeIDs]
	nodeGroupIDs := scheduleAttrs[fields.NodeGroupIDs]

	// Now variables should be resolved
	nodeIDsEmpty := isListEmpty(nodeIDs)
	nodeGroupIDsEmpty := isListEmpty(nodeGroupIDs)

	if nodeIDsEmpty && nodeGroupIDsEmpty {
		resp.Diagnostics.AddAttributeError(
			req.Path,
			"Invalid Schedule Settings Configuration",
			"When schedule_settings is configured, either node_ids or node_group_ids must be specified.",
		)
	}
}

func isListEmpty(attr any) bool {
	if list, ok := attr.(types.List); ok {
		if list.IsNull() || list.IsUnknown() {
			return true
		}
		return len(list.Elements()) == 0
	}
	return false
}
