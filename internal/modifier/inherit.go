package modifier

import (
	"context"

	"catchpoint-provider/internal/logger"
	cptypes "catchpoint-provider/internal/types"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func InheritBlockPlanModifier(settingTypeName string) planmodifier.Object {
	return &inheritBlockPlanModifier{settingTypeName}
}

type inheritBlockPlanModifier struct {
	settingTypeName string
}

func (d *inheritBlockPlanModifier) Description(ctx context.Context) string {
	return "Ensures that the block is set to 'inherit' upstream values if none are set. Otherwise, sets to 'override'"
}

func (d *inheritBlockPlanModifier) MarkdownDescription(ctx context.Context) string {
	return d.Description(ctx)
}

func (d *inheritBlockPlanModifier) PlanModifyObject(ctx context.Context, req planmodifier.ObjectRequest, resp *planmodifier.ObjectResponse) {
	obj := req.PlanValue

	if req.PlanValue.IsNull() || req.PlanValue.IsUnknown() {
		return
	}

	attrs := obj.Attributes()

	if !attrs[d.settingTypeName].IsNull() && !attrs[d.settingTypeName].IsUnknown() {
		// User has explicitly set the type; respect that.
		logger.Debug(ctx, "InheritBlockPlanModifier: %s is already set to %s, respecting user choice", d.settingTypeName, attrs[d.settingTypeName].String())
		return
	}

	// Check if any attribute other than `type` is set.
	for name, val := range attrs {
		if name == d.settingTypeName {
			continue
		}
		if !val.IsNull() && !val.IsUnknown() {
			attrs[d.settingTypeName] = types.StringValue(cptypes.Override)
			value, diags := types.ObjectValue(obj.AttributeTypes(ctx), attrs)
			if diags.HasError() {
				return
			}
			resp.PlanValue = value
			logger.Debug(ctx, "InheritBlockPlanModifier: Setting %s to 'override' because attribute %s is set", d.settingTypeName, name)
			return
		}
	}

	// No other attributes set — treat as "Inherit."
	attrs[d.settingTypeName] = types.StringValue(cptypes.Inherit)
	value, diags := types.ObjectValue(obj.AttributeTypes(ctx), attrs)
	if diags.HasError() {
		return
	}
	resp.PlanValue = value
	logger.Debug(ctx, "InheritBlockPlanModifier: Setting %s to 'inherit' because no other attributes are set", d.settingTypeName)
}
