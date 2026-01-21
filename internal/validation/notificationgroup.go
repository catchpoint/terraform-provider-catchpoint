package validation

import (
	"context"

	"catchpoint-provider/internal/fields"
	"catchpoint-provider/internal/types"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

var _ validator.Object = requireNotificationGroupValidator{}

// RequireNotificationGroupValidator returns a validator that enforces that when the
// alert_settings block is declared in configuration, a notification_group block
// is also declared with a non-empty subject and at least one of
// recipient_emails or contact_group_ids populated.
func RequireNotificationGroupValidator() validator.Object {
	return requireNotificationGroupValidator{}
}

type requireNotificationGroupValidator struct{}

func (v requireNotificationGroupValidator) Description(_ context.Context) string {
	return "Ensures notification_group is set with subject and either recipient_emails or contact_group_ids when alert_settings is provided."
}

func (v requireNotificationGroupValidator) MarkdownDescription(ctx context.Context) string {
	return v.Description(ctx)
}

func (v requireNotificationGroupValidator) ValidateObject(ctx context.Context, req validator.ObjectRequest, resp *validator.ObjectResponse) {
	// If the user did not configure alert_settings, skip validation.
	if req.ConfigValue.IsNull() || req.ConfigValue.IsUnknown() {
		return
	}

	ov := req.ConfigValue
	attrs := ov.Attributes()

	ngAttr, ngOk := attrs[fields.NotificationGroup]
	ngPath := req.Path.AtName(fields.NotificationGroup)

	if alertSettingType, ok := attrs[fields.AlertSettingType]; ok {
		if sv, ok := alertSettingType.(basetypes.StringValue); ok && sv.ValueString() == types.Inherit {
			// If alert_settings.type is "inherit", skip validation.
			return
		}
	}

	validateNotificationGroup(ngAttr, ngOk, ngPath, resp)
}

func validateNotificationGroup(ngAttr attr.Value, ok bool, ngPath path.Path, resp *validator.ObjectResponse) {
	// notification_group block must exist
	if !ok || ngAttr.IsNull() || ngAttr.IsUnknown() {
		resp.Diagnostics.Append(diag.NewAttributeErrorDiagnostic(
			ngPath,
			"Missing notification_group block",
			"The alert_settings block was provided, so a notification_group block is required with a subject and either recipient_emails or contact_group_ids.",
		))
		return
	}

	ngObj, ok := ngAttr.(basetypes.ObjectValue)
	if !ok {
		// Malformed; framework should not allow this state, but guard anyway.
		return
	}

	ngAttrs := ngObj.Attributes()

	subjectAttr, hasSubject := ngAttrs[fields.Subject]
	subjectPath := ngPath.AtName(fields.Subject)

	if !hasSubject || subjectAttr.IsNull() || subjectAttr.IsUnknown() {
		resp.Diagnostics.Append(diag.NewAttributeErrorDiagnostic(
			subjectPath,
			"Missing notification_group.subject",
			"A non-empty subject is required when defining alert_settings.",
		))
	} else {
		if sv, ok := subjectAttr.(basetypes.StringValue); ok && sv.ValueString() == "" {
			resp.Diagnostics.Append(diag.NewAttributeErrorDiagnostic(
				subjectPath,
				"Empty notification_group.subject",
				"A non-empty subject is required when defining alert_settings.",
			))
		}
	}

	recipsLen := listLenAttr(ngAttrs[fields.RecipientEmails])
	cgLen := listLenAttr(ngAttrs[fields.ContactGroupIDs])

	if recipsLen == 0 && cgLen == 0 {
		resp.Diagnostics.Append(diag.NewAttributeErrorDiagnostic(
			ngPath,
			"Missing notification recipients",
			"At least one of notification_group.recipient_emails or notification_group.contact_group_ids must be provided.",
		))
	}
}

func listLenAttr(a attr.Value) int {
	if a == nil || a.IsNull() || a.IsUnknown() {
		return 0
	}
	lv, ok := a.(basetypes.ListValue)
	if !ok {
		return 0
	}
	return len(lv.Elements())
}
