package validation

import (
	"context"
	"fmt"

	"catchpoint-provider/internal/fields"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// AlertRuleThresholdsValidator ensures each alert_rule object defines at least one
// of threshold_number_of_runs or threshold_percentage_of_runs.
func AlertRuleThresholdsValidator() validator.List {
	return alertRuleThresholdsValidator{}
}

type alertRuleThresholdsValidator struct{}

func (v alertRuleThresholdsValidator) Description(_ context.Context) string {
	return "Each alert_rule must set at least one of threshold_number_of_runs or threshold_percentage_of_runs."
}

func (v alertRuleThresholdsValidator) MarkdownDescription(ctx context.Context) string {
	return v.Description(ctx)
}

func (v alertRuleThresholdsValidator) ValidateList(ctx context.Context, req validator.ListRequest, resp *validator.ListResponse) {
	// If list not configured, nothing to validate.
	if req.ConfigValue.IsNull() || req.ConfigValue.IsUnknown() {
		return
	}

	var lv []attr.Value
	diags := req.ConfigValue.ElementsAs(ctx, &lv, false)
	if diags.HasError() {
		return
	}

	for idx, elem := range lv {
		obj, ok := elem.(basetypes.ObjectValue)
		if !ok {
			continue
		}
		attrs := obj.Attributes()

		runsAttr, hasRuns := attrs[fields.ThresholdNumberOfRuns]
		pctAttr, hasPct := attrs[fields.ThresholdPercentageOfRuns]

		hasRunsVal := hasRuns && !runsAttr.IsNull() && !runsAttr.IsUnknown()
		hasPctVal := hasPct && !pctAttr.IsNull() && !pctAttr.IsUnknown()

		if !hasRunsVal && !hasPctVal {
			resp.Diagnostics.Append(diag.NewAttributeErrorDiagnostic(
				req.Path, // attribute path to the list
				"Missing threshold in alert_rule",
				fmt.Sprintf("alert_rule[%d] must set at least one of threshold_number_of_runs or threshold_percentage_of_runs.", idx),
			))
		}
	}
}

var _ validator.List = alertRuleThresholdsValidator{}
