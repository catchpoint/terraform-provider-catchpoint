package merge

import (
	"fmt"

	cpresource "catchpoint-provider/internal/models/resource"
	cptypes "catchpoint-provider/internal/types"

	"github.com/hashicorp/terraform-plugin-framework/types"
)

// AlertRuleSorter implements sort.Interface for []cpresource.AlertRuleModel
type AlertRuleSorter []cpresource.AlertRuleModel

func (a AlertRuleSorter) Len() int      { return len(a) }
func (a AlertRuleSorter) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a AlertRuleSorter) Less(i, j int) bool {
	// Sort by AlertType first
	if a[i].AlertType.ValueString() != a[j].AlertType.ValueString() {
		return a[i].AlertType.ValueString() < a[j].AlertType.ValueString()
	}

	// Then AlertSubType
	if a[i].AlertSubType.ValueString() != a[j].AlertSubType.ValueString() {
		return a[i].AlertSubType.ValueString() < a[j].AlertSubType.ValueString()
	}

	// Then TriggerType
	if a[i].TriggerType.ValueString() != a[j].TriggerType.ValueString() {
		return a[i].TriggerType.ValueString() < a[j].TriggerType.ValueString()
	}

	// Then NodeThresholdType
	if a[i].NodeThresholdType.ValueString() != a[j].NodeThresholdType.ValueString() {
		return a[i].NodeThresholdType.ValueString() < a[j].NodeThresholdType.ValueString()
	}

	// Finally, use a numeric field as final tie-breaker for stable sort
	return a[i].ThresholdNumberOfRuns.ValueInt64() < a[j].ThresholdNumberOfRuns.ValueInt64()
}

// buildSimpleAlertRuleKey creates a stable identity for a rule.
func buildSimpleAlertRuleKey(rule cpresource.AlertRuleModel) string {
	alertType := rule.AlertType.ValueString()
	if alertType == cptypes.EmptyString {
		alertType = cptypes.None
	}
	subType := rule.AlertSubType.ValueString()
	if subType == cptypes.EmptyString {
		subType = cptypes.None
	}
	trigger := rule.TriggerType.ValueString()
	if trigger == cptypes.EmptyString {
		trigger = cptypes.None
	}
	nodeThresh := rule.NodeThresholdType.ValueString()
	if nodeThresh == cptypes.EmptyString {
		nodeThresh = cptypes.None
	}

	var threshPart string
	switch {
	case !rule.ThresholdNumberOfRuns.IsNull() && !rule.ThresholdNumberOfRuns.IsUnknown():
		threshPart = fmt.Sprintf("runs:%d", rule.ThresholdNumberOfRuns.ValueInt64())
	case !rule.ThresholdPercentageOfRuns.IsNull() && !rule.ThresholdPercentageOfRuns.IsUnknown():
		// limit float format to avoid spurious key differences (whole numbers -> no decimals)
		p := rule.ThresholdPercentageOfRuns.ValueFloat64()
		if p == float64(int64(p)) {
			threshPart = fmt.Sprintf("pct:%d", int64(p))
		} else {
			threshPart = fmt.Sprintf("pct:%g", p)
		}
	default:
		// Should not happen if validator ran, but keep deterministic fallback
		threshPart = "no-threshold"
	}

	key := fmt.Sprintf("%s|%s|%s|%s|%s", alertType, subType, trigger, nodeThresh, threshPart)

	// Include trigger numeric discriminators only if they exist (optional)
	if !rule.WarningTrigger.IsNull() && !rule.WarningTrigger.IsUnknown() {
		key += fmt.Sprintf("|w:%g", rule.WarningTrigger.ValueFloat64())
	}
	if !rule.CriticalTrigger.IsNull() && !rule.CriticalTrigger.IsUnknown() {
		key += fmt.Sprintf("|c:%g", rule.CriticalTrigger.ValueFloat64())
	}
	return key
}

// AlertSettings merges plan and API state for alert settings.
// When alert_setting_type is "inherit" in the API, it indicates drift that needs correction.
func AlertSettings(plan, apiState *cpresource.AlertSettingsModel) *cpresource.AlertSettingsModel {
	if plan == nil {
		return nil
	}

	if apiState == nil {
		// Creation case: no prior API state
		return &cpresource.AlertSettingsModel{
			AlertSettingType:  plan.AlertSettingType,
			AlertRule:         normalizeAlertRuleList(plan.AlertRule, nil),
			NotificationGroup: mergeTopLevelNotificationGroup(plan.NotificationGroup, nil),
		}
	}

	result := &cpresource.AlertSettingsModel{}

	if !plan.AlertSettingType.IsNull() && !plan.AlertSettingType.IsUnknown() {
		result.AlertSettingType = plan.AlertSettingType
	} else {
		result.AlertSettingType = apiState.AlertSettingType
	}

	// Normal merge for alert rules and notification groups
	// The API still returns these even when alert_setting_type is "inherit"
	result.AlertRule = normalizeAlertRuleList(plan.AlertRule, apiState.AlertRule)
	result.NotificationGroup = mergeTopLevelNotificationGroup(plan.NotificationGroup, getAPINotificationGroup(apiState))

	return result
}

// normalizeAlertRuleList merges plan rules with api rules, preserving plan order.
func normalizeAlertRuleList(planRules, apiRules []cpresource.AlertRuleModel) []cpresource.AlertRuleModel {
	if len(planRules) == 0 {
		return nil
	}

	apiMap := make(map[string]cpresource.AlertRuleModel, len(apiRules))
	for _, ar := range apiRules {
		apiMap[buildSimpleAlertRuleKey(ar)] = ar
	}

	// Create a new slice to hold the final, merged rules.
	out := make([]cpresource.AlertRuleModel, len(planRules))
	for i, planRule := range planRules {
		// Start with a copy of the plan rule for this iteration.
		mergedRule := planRule
		key := buildSimpleAlertRuleKey(planRule)

		// Find the matching API rule.
		if apiRule, ok := apiMap[key]; ok {
			// Merge the plan and API rule, modifying mergedRule in place.
			mergeAlertRule(&mergedRule, &apiRule)
		}

		// Finalize all computed values to ensure they are known.
		finalizeComputedValues(&mergedRule)
		out[i] = mergedRule
	}
	return out
}

// mergeAlertRule now accepts pointers and modifies the 'plan' rule directly.
func mergeAlertRule(plan *cpresource.AlertRuleModel, api *cpresource.AlertRuleModel) {
	if api == nil {
		return
	}

	mergeAlertRuleComputedStrings(plan, api)
	mergeAlertRuleComputedBools(plan, api)

	// Merge nested notification groups and assign the result back to the plan object.
	if len(plan.NotificationGroup) > 0 {
		plan.NotificationGroup = mergeNestedNotificationGroups(plan.NotificationGroup, api.NotificationGroup)
	}

	// Merge the nested level block.
	if plan.Level != nil {
		plan.Level = mergeLevelModel(plan.Level, api.Level)
	}
}

func mergeAlertRuleComputedStrings(plan, api *cpresource.AlertRuleModel) {
	stringFields := []struct {
		planField *types.String
		apiField  types.String
	}{
		{&plan.OperationType, api.OperationType},
		{&plan.NotificationType, api.NotificationType},
		{&plan.WarningReminder, api.WarningReminder},
		{&plan.CriticalReminder, api.CriticalReminder},
		{&plan.StatisticalType, api.StatisticalType},
		{&plan.HistoricalInterval, api.HistoricalInterval},
		{&plan.TriggerType, api.TriggerType},
	}
	for _, f := range stringFields {
		if f.planField.IsNull() || f.planField.IsUnknown() {
			*f.planField = f.apiField
		}
	}
}

func mergeAlertRuleComputedBools(plan, api *cpresource.AlertRuleModel) {
	boolFields := []struct {
		planField *types.Bool
		apiField  types.Bool
	}{
		{&plan.UseRollingWindow, api.UseRollingWindow},
		{&plan.EnforceTestFailure, api.EnforceTestFailure},
		{&plan.AllMatchRecords, api.AllMatchRecords},
		{&plan.OmitScatterplot, api.OmitScatterplot},
		{&plan.EnableConsecutive, api.EnableConsecutive},
	}
	for _, f := range boolFields {
		if f.planField.IsNull() || f.planField.IsUnknown() {
			*f.planField = f.apiField
		}
	}
}

// finalizeComputedValues ensures all Optional+Computed fields have a known value.
func finalizeComputedValues(r *cpresource.AlertRuleModel) {
	// Ensure all computed bools are known (default to false if null).
	forceBool(&r.UseRollingWindow)
	forceBool(&r.EnforceTestFailure)
	forceBool(&r.AllMatchRecords)
	forceBool(&r.OmitScatterplot)
	forceBool(&r.EnableConsecutive)

	// Ensure computed strings are known (default to "" if null).
	if r.OperationType.IsNull() || r.OperationType.IsUnknown() {
		r.OperationType = types.StringValue("")
	}

	if r.NotificationType.IsNull() || r.NotificationType.IsUnknown() {
		r.NotificationType = types.StringValue("")
	}

	if r.WarningReminder.IsNull() || r.WarningReminder.IsUnknown() {
		r.WarningReminder = types.StringValue("")
	}

	if r.CriticalReminder.IsNull() || r.CriticalReminder.IsUnknown() {
		r.CriticalReminder = types.StringValue("")
	}
	if r.StatisticalType.IsNull() || r.StatisticalType.IsUnknown() {
		r.StatisticalType = types.StringValue("")
	}
	if r.HistoricalInterval.IsNull() || r.HistoricalInterval.IsUnknown() {
		r.HistoricalInterval = types.StringValue("")
	}
	if r.TriggerType.IsNull() || r.TriggerType.IsUnknown() {
		r.TriggerType = types.StringValue("")
	}

	// Finalize computed values in nested notification groups
	for i := range r.NotificationGroup {
		finalizeNotificationGroupComputedValues(&r.NotificationGroup[i])
	}
}

// finalizeNotificationGroupComputedValues ensures all computed values in a
// notification group are known after apply.
func finalizeNotificationGroupComputedValues(ng *cpresource.NotificationGroupModel) {
	// Force computed bools to known values
	forceBool(&ng.NotifyOnWarning)
	forceBool(&ng.NotifyOnCritical)
	forceBool(&ng.NotifyOnImproved)

	// Force computed list to known value
	if ng.AlertWebhookIDs.IsNull() || ng.AlertWebhookIDs.IsUnknown() {
		ng.AlertWebhookIDs = types.ListNull(types.Int64Type)
	}
}

// buildNotificationGroupKey creates a stable identity for a nested notification group.
// It uses fields that uniquely identify a group, like its recipients.
func buildNotificationGroupKey(ng cpresource.NotificationGroupModel) string {
	// A key based on subject is often sufficient if subjects are unique per rule.
	if !ng.Subject.IsNull() && !ng.Subject.IsUnknown() && ng.Subject.ValueString() != "" {
		return ng.Subject.ValueString()
	}

	// If subject is not a reliable key, create one from recipient counts.
	// This is a fallback and assumes the combination of counts is unique.
	emailCount := 0
	if !ng.Emails.IsNull() {
		emailCount = len(ng.Emails.Elements())
	}
	groupCount := 0
	if !ng.ContactGroupIDs.IsNull() {
		groupCount = len(ng.ContactGroupIDs.Elements())
	}
	return fmt.Sprintf("emails:%d,groups:%d", emailCount, groupCount)
}

// mergeTopLevelNotificationGroup handles the single notification group under alert_settings.
// Here, AlertWebhookIDs is NOT computed and should only come from the plan.
func mergeTopLevelNotificationGroup(plan, api *cpresource.NotificationGroupModel) *cpresource.NotificationGroupModel {
	if plan == nil {
		return nil // User did not configure this block.
	}

	result := &cpresource.NotificationGroupModel{}

	// The three NotifyOn* bools are Optional+Computed.
	result.NotifyOnWarning = pickComputedBool(plan.NotifyOnWarning, api, func(ng *cpresource.NotificationGroupModel) types.Bool { return ng.NotifyOnWarning })
	result.NotifyOnCritical = pickComputedBool(plan.NotifyOnCritical, api, func(ng *cpresource.NotificationGroupModel) types.Bool { return ng.NotifyOnCritical })
	result.NotifyOnImproved = pickComputedBool(plan.NotifyOnImproved, api, func(ng *cpresource.NotificationGroupModel) types.Bool { return ng.NotifyOnImproved })

	if plan.AlertWebhookIDs.IsNull() || plan.AlertWebhookIDs.IsUnknown() {
		result.AlertWebhookIDs = types.ListNull(types.Int64Type)
	} else {
		result.AlertWebhookIDs = plan.AlertWebhookIDs
	}

	// Purely optional fields: take the user's value from the plan, even if null.
	// Do NOT fall back to the API value if the user omitted them.
	result.Subject = plan.Subject
	result.Emails = plan.Emails
	result.ContactGroupIDs = plan.ContactGroupIDs

	return result
}

// mergeNestedNotificationGroups handles the list of notification groups under an alert_rule.
// It now uses key-based matching to correlate plan and API objects correctly.
func mergeNestedNotificationGroups(planGroups, apiGroups []cpresource.NotificationGroupModel) []cpresource.NotificationGroupModel {
	if len(planGroups) == 0 {
		return nil
	}

	apiMap := make(map[string]cpresource.NotificationGroupModel, len(apiGroups))
	for _, apiGroup := range apiGroups {
		key := buildNotificationGroupKey(apiGroup)
		apiMap[key] = apiGroup
	}

	// The result must be a new slice to avoid modifying the original plan state directly.
	mergedGroups := make([]cpresource.NotificationGroupModel, len(planGroups))

	// Iterate by index to correctly assign the merged result to the new slice.
	for i := range planGroups {
		planGroup := &planGroups[i] // Get a pointer to the current plan group.
		key := buildNotificationGroupKey(*planGroup)

		// Look for a matching API group.
		if apiGroup, ok := apiMap[key]; ok {
			// Match found, merge them.
			mergedGroups[i] = *mergeNestedNotificationGroup(planGroup, &apiGroup)
		} else {
			// No matching API group, merge with nil to apply defaults to the plan object.
			mergedGroups[i] = *mergeNestedNotificationGroup(planGroup, nil)
		}
	}

	return mergedGroups
}

// mergeNestedNotificationGroup handles a single notification group from a nested list.
func mergeNestedNotificationGroup(plan, api *cpresource.NotificationGroupModel) *cpresource.NotificationGroupModel {
	if plan == nil {
		return nil
	}

	// Start with a copy of the plan to avoid modifying the original object.
	result := *plan

	// The three NotifyOn* bools are Optional+Computed.
	result.NotifyOnWarning = pickComputedBool(plan.NotifyOnWarning, api, func(ng *cpresource.NotificationGroupModel) types.Bool { return ng.NotifyOnWarning })
	result.NotifyOnCritical = pickComputedBool(plan.NotifyOnCritical, api, func(ng *cpresource.NotificationGroupModel) types.Bool { return ng.NotifyOnCritical })
	result.NotifyOnImproved = pickComputedBool(plan.NotifyOnImproved, api, func(ng *cpresource.NotificationGroupModel) types.Bool { return ng.NotifyOnImproved })

	// AlertWebhookIDs is Optional+Computed in the nested block.
	if plan.AlertWebhookIDs.IsNull() || plan.AlertWebhookIDs.IsUnknown() {
		if api != nil && !api.AlertWebhookIDs.IsNull() && !api.AlertWebhookIDs.IsUnknown() {
			result.AlertWebhookIDs = api.AlertWebhookIDs
		}
		// Note: We don't set the empty list here because finalizeComputedValues
		// will handle it later. This prevents duplicate work and keeps the logic clean.
	}

	return &result
}

// pickComputedBool is a helper to resolve Optional+Computed boolean values.
// It prefers the plan value, falls back to the API value, and finally defaults to false.
func pickComputedBool(planVal types.Bool, api *cpresource.NotificationGroupModel, getAPIVal func(*cpresource.NotificationGroupModel) types.Bool) types.Bool {
	// If the user provided a value, use it.
	if !planVal.IsNull() && !planVal.IsUnknown() {
		return planVal
	}
	// If the user did not, try the API value.
	if api != nil {
		apiVal := getAPIVal(api)
		if !apiVal.IsNull() && !apiVal.IsUnknown() {
			return apiVal
		}
	}
	// If both are missing, default to false as it's a computed value.
	return types.BoolValue(false)
}

func mergeLevelModel(plan, apiState *cpresource.LevelModel) *cpresource.LevelModel {
	if plan == nil {
		return apiState
	}

	result := &cpresource.LevelModel{
		FilterType:  plan.FilterType,
		FilterValue: plan.FilterValue,
	}

	// Merge computed fields from API if plan values are null/unknown
	if plan.FilterType.IsNull() || plan.FilterType.IsUnknown() {
		if apiState != nil {
			result.FilterType = apiState.FilterType
		}
	}

	return result
}

func getAPINotificationGroup(apiState *cpresource.AlertSettingsModel) *cpresource.NotificationGroupModel {
	if apiState == nil {
		return nil
	}
	return apiState.NotificationGroup
}

// forceBool is a simplified helper to ensure a bool is not null/unknown.
func forceBool(v *types.Bool) {
	if v.IsNull() || v.IsUnknown() {
		*v = types.BoolValue(false)
	}
}
