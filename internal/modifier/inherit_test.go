package modifier

import (
	"catchpoint-provider/internal/testutil"
	"context"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func TestInheritBlockPlanModifier(t *testing.T) {
	mod := InheritBlockPlanModifier("type")
	ctx := context.Background()

	attrTypes := map[string]attr.Type{
		"type": types.StringType,
		"foo":  types.StringType,
		"bar":  types.Int64Type,
	}

	tests := []struct {
		name     string
		attrs    map[string]attr.Value
		expected string
	}{
		{
			name: "No attributes set",
			attrs: map[string]attr.Value{
				"type": types.StringNull(),
				"foo":  types.StringNull(),
				"bar":  types.Int64Null(),
			},
			expected: "inherit",
		},
		{
			name: "One attribute set",
			attrs: map[string]attr.Value{
				"type": types.StringNull(),
				"foo":  types.StringValue("hello"),
				"bar":  types.Int64Null(),
			},
			expected: "override",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			objVal := types.ObjectValueMust(attrTypes, tt.attrs)
			req := planmodifier.ObjectRequest{
				PlanValue: objVal,
			}
			var resp planmodifier.ObjectResponse
			mod.PlanModifyObject(ctx, req, &resp)

			result := resp.PlanValue
			actual := result.Attributes()["type"].(types.String).ValueString()

			testutil.AssertEqual(t, "Type", actual, tt.expected)
		})
	}
}

func TestGenericTypePlanModifier(t *testing.T) {
	mod := InheritBlockPlanModifier("type")
	ctx := context.Background()

	attrTypes := map[string]attr.Type{
		"type": types.StringType,
		"foo":  types.StringType,
		"bar":  types.Int64Type,
	}

	tests := []struct {
		name     string
		plan     attr.Value
		expected string
	}{
		{
			name: "All nulls",
			plan: types.ObjectValueMust(attrTypes, map[string]attr.Value{
				"type": types.StringNull(),
				"foo":  types.StringNull(),
				"bar":  types.Int64Null(),
			}),
			expected: "inherit",
		},
		{
			name: "One attribute set",
			plan: types.ObjectValueMust(attrTypes, map[string]attr.Value{
				"type": types.StringNull(),
				"foo":  types.StringValue("hello"),
				"bar":  types.Int64Null(),
			}),
			expected: "override",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := planmodifier.ObjectRequest{
				PlanValue: tt.plan.(types.Object),
			}
			var resp planmodifier.ObjectResponse
			mod.PlanModifyObject(ctx, req, &resp)

			result := resp.PlanValue.Attributes()["type"].(types.String).ValueString()
			testutil.AssertEqual(t, "Type", tt.expected, result)
		})
	}
}
