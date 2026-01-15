package transform

import (
	"fmt"

	"catchpoint-provider/internal/fields"
	"catchpoint-provider/internal/helpers"
	"catchpoint-provider/internal/models"
	cpresource "catchpoint-provider/internal/models/resource"
	cptypes "catchpoint-provider/internal/types"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func JSONToTerraformAlertGroup(alertSetting *models.AlertGroupJSON, config *cpresource.AlertSettingsModel) (resource *cpresource.AlertSettingsModel, diags diag.Diagnostics) {
	if alertSetting == nil || config.IsNull() {
		return nil, nil
	}
	resource = &cpresource.AlertSettingsModel{}

	// Note: we cannot use the GenericSettingTypeName here because AlertSettingType also offers "Inherit & Add."
	if !config.AlertSettingType.IsNull() && !config.AlertSettingType.IsUnknown() {
		alertSettingTypeName, ok := cptypes.GetAlertSettingTypeName(alertSetting.AlertSettingType.ID)
		if !ok {
			diags.Append(diag.NewWarningDiagnostic("Invalid AlertSettingType ID", fmt.Sprintf("The provided AlertSettingType ID is not recognized: %d", alertSetting.AlertSettingType.ID)))
		}
		resource.AlertSettingType = types.StringValue(alertSettingTypeName)
	}

	alertItems, alertItemsDiags := jsonToTerraformAlertGroupItems(alertSetting.AlertGroupItems)
	diags.Append(alertItemsDiags...)
	if diags.HasError() {
		return
	}
	resource.AlertRule = alertItems

	if alertSetting.NotificationGroup != nil && !config.NotificationGroup.IsNull() {
		notificationGroup, notificationGroupDiags := jsonToTerraformNotificationGroup(*alertSetting.NotificationGroup)
		diags.Append(notificationGroupDiags...)
		if diags.HasError() {
			return
		}
		resource.NotificationGroup = &notificationGroup
	}

	return
}

func jsonToTerraformAlertGroupItems(alertItems []models.AlertGroupItemJSON) (resource []cpresource.AlertRuleModel, diags diag.Diagnostics) {
	for _, alertItem := range alertItems {
		// Convert each alert item to an object
		itemObj, itemDiags := jsonToTerraformAlertRule(alertItem)
		diags.Append(itemDiags...)
		if diags.HasError() {
			return
		}
		resource = append(resource, itemObj)
	}

	return
}

func jsonToTerraformAlertRule(alertItem models.AlertGroupItemJSON) (alertRuleModel cpresource.AlertRuleModel, diags diag.Diagnostics) {
	diags = append(diags, stringValueFromIDWithLookup(alertItem.NotificationType.ID, &alertRuleModel.NotificationType, cptypes.GetNotificationTypeName, fields.NotificationType)...)
	if diags.HasError() {
		return
	}

	diags = append(diags, stringValueFromIDWithLookup(alertItem.AlertType.ID, &alertRuleModel.AlertType, cptypes.GetAlertTypeName, fields.AlertType)...)
	if diags.HasError() {
		return
	}

	diags = append(diags, stringValueFromOptionalGenericWithLookup(alertItem.AlertSubType, &alertRuleModel.AlertSubType, cptypes.GetAlertSubTypeName, fields.AlertSubType)...)
	if diags.HasError() {
		return
	}

	notificationGroupObj, notificationGroupDiags := jsonToTerraformNotificationGroups(alertItem.NotificationGroups)
	diags.Append(notificationGroupDiags...)
	if diags.HasError() {
		return
	}
	alertRuleModel.NotificationGroup = notificationGroupObj

	alertRuleModel.EnforceTestFailure = types.BoolValue(alertItem.EnforceTestFailure)
	alertRuleModel.OmitScatterplot = types.BoolValue(alertItem.OmitScatterplot)
	alertRuleModel.AllMatchRecords = types.BoolValue(alertItem.MatchAllRecords)

	nodeThreshold, nodeThresholdDiags := jsonToTerraformNodeThreshold(alertItem.NodeThreshold)
	diags.Append(nodeThresholdDiags...)
	if diags.HasError() {
		return
	}
	alertRuleModel.NodeThresholdModel = nodeThreshold

	trigger, triggerDiags := jsonToTerraformTrigger(alertItem.Trigger)
	diags.Append(triggerDiags...)
	if diags.HasError() {
		return
	}
	alertRuleModel.TriggerModel = trigger

	return
}

func jsonToTerraformNodeThreshold(threshold models.NodeThresholdJSON) (resource cpresource.NodeThresholdModel, diags diag.Diagnostics) {
	setInt64ValueIfNotNil(threshold.NumberOfConsecutiveRuns, &resource.ConsecutiveNumberOfRuns)

	resource.EnableConsecutive = types.BoolValue(threshold.ConsecutiveRunsEnabled)

	diags = append(diags, stringValueFromIDWithLookup(threshold.NodeThresholdType.ID, &resource.NodeThresholdType, cptypes.GetNodeThresholdTypeName, fields.NodeThresholdType)...)
	if diags.HasError() {
		return
	}

	setInt64ValueIfNotNil(threshold.NumberOfFailingUnits, &resource.NumberOfFailingNodes)
	setInt64ValueIfNotNil(threshold.NumberOfUnits, &resource.ThresholdNumberOfRuns)
	setFloat64ValueIfNotNil(threshold.PercentageOfUnits, &resource.ThresholdPercentageOfRuns)

	return
}

func jsonToTerraformTrigger(trigger models.TriggerJSON) (triggerModel cpresource.TriggerModel, diags diag.Diagnostics) {
	setStringValueIfNotNil(trigger.Expression, &triggerModel.Expression)

	triggerModel.UseRollingWindow = types.BoolValue(trigger.UseIntervalRollingWindow)

	// Not optional.
	diags = append(diags, stringValueFromIDWithLookup(trigger.TriggerType.ID, &triggerModel.TriggerType, cptypes.GetTriggerTypeName, fields.TriggerType)...)
	if diags.HasError() {
		return
	}
	diags = append(diags, stringValueFromIDWithLookup(trigger.OperationType.ID, &triggerModel.OperationType, cptypes.GetOperationTypeName, fields.OperationType)...)
	if diags.HasError() {
		return
	}
	diags = append(diags, stringValueFromIDWithLookup(trigger.WarningReminderFrequency.ID, &triggerModel.WarningReminder, cptypes.GetReminderName, fields.WarningReminder)...)
	if diags.HasError() {
		return
	}
	diags = append(diags, stringValueFromIDWithLookup(trigger.CriticalReminderFrequency.ID, &triggerModel.CriticalReminder, cptypes.GetReminderName, fields.CriticalReminder)...)
	if diags.HasError() {
		return
	}
	diags = append(diags, stringValueFromIDWithLookup(trigger.ThresholdInterval.ID, &triggerModel.ThresholdInterval, cptypes.GetThresholdIntervalName, fields.ThresholdInterval)...)
	if diags.HasError() {
		return
	}

	// Optional.
	diags = append(diags, stringValueFromOptionalGenericWithLookup(trigger.HistoricalInterval, &triggerModel.HistoricalInterval, cptypes.GetHistoricalIntervalName, fields.HistoricalInterval)...)
	if diags.HasError() {
		return
	}
	if triggerModel.HistoricalInterval.IsNull() {
		triggerModel.HistoricalInterval = types.StringValue(cptypes.EmptyString)
	}

	diags = append(diags, stringValueFromOptionalGenericWithLookup(trigger.StatisticalType, &triggerModel.StatisticalType, cptypes.GetStatisticalTypeName, fields.StatisticalType)...)
	if diags.HasError() {
		return
	}
	if triggerModel.StatisticalType.IsNull() {
		triggerModel.StatisticalType = types.StringValue(cptypes.EmptyString)
	}

	setFloat64ValueIfNotNil(trigger.WarningTrigger, &triggerModel.WarningTrigger)
	setFloat64ValueIfNotNil(trigger.CriticalTrigger, &triggerModel.CriticalTrigger)

	// DNS settings.
	setStringValueIfNotNil(trigger.DNSResolvedName, &triggerModel.DNSResolvedName)
	setInt64ValueIfNotNil(trigger.DNSTTL, &triggerModel.DNSTTL)
	diags = append(diags, stringValueFromOptionalGenericWithLookup(trigger.DNSRecordType, &triggerModel.DNSRecordType, cptypes.GetDNSRecordTypeName, fields.DNSRecordType)...)
	if diags.HasError() {
		return
	}

	// "Level" settings.
	if trigger.FilterType != nil || trigger.FilterValue != nil {
		triggerModel.Level = &cpresource.LevelModel{}
		diags = append(diags, stringValueFromOptionalGenericWithLookup(trigger.FilterType, &triggerModel.Level.FilterType, cptypes.GetFilterTypeName, fields.FilterType)...)
		if diags.HasError() {
			return
		}
		setStringValueIfNotNil(trigger.FilterValue, &triggerModel.Level.FilterValue)
	}

	return
}

func jsonToTerraformNotificationGroups(notificationGroup []models.NotificationGroupJSON) (resource []cpresource.NotificationGroupModel, diags diag.Diagnostics) {
	if len(notificationGroup) == 0 {
		return
	}

	for _, group := range notificationGroup {
		notificationGroupObj, groupDiags := jsonToTerraformNotificationGroup(group)
		diags.Append(groupDiags...)
		if diags.HasError() {
			return
		}
		resource = append(resource, notificationGroupObj)
	}

	return
}

func jsonToTerraformNotificationGroup(notificationGroup models.NotificationGroupJSON) (notifGroupModel cpresource.NotificationGroupModel, diags diag.Diagnostics) {
	ids := helpers.FlattenToIDs(notificationGroup.AlertWebhooks, func(x models.AlertWebhookJSON) int { return *x.ID })
	notifGroupModel.AlertWebhookIDs, diags = intSliceToInt64List(ids)

	emailList, contactGroups, emailDiags := jsonRecipientsToLists(notificationGroup.Recipients)
	diags.Append(emailDiags...)
	if diags.HasError() {
		return
	}
	notifGroupModel.Emails = emailList
	notifGroupModel.ContactGroupIDs = contactGroups
	notifGroupModel.Subject = types.StringValue(notificationGroup.Subject)

	notifGroupModel.NotifyOnWarning = types.BoolValue(notificationGroup.NotifyOnWarning)
	notifGroupModel.NotifyOnCritical = types.BoolValue(notificationGroup.NotifyOnCritical)
	notifGroupModel.NotifyOnImproved = types.BoolValue(notificationGroup.NotifyOnImproved)

	return
}

func jsonRecipientsToLists(emails []models.RecipientJSON) (emailList, contactGroups types.List, diags diag.Diagnostics) {
	emailValues := extractEmailValues(emails)
	contactGroupValues := extractContactGroupValues(emails)

	// Create email list
	if len(emailValues) == 0 {
		emailList = types.ListNull(types.StringType)
	} else {
		var emailListDiags diag.Diagnostics
		emailList, emailListDiags = types.ListValue(types.StringType, emailValues)
		diags.Append(emailListDiags...)
		if diags.HasError() {
			return
		}
	}

	// Create contact groups list
	if len(contactGroupValues) == 0 {
		contactGroups = types.ListNull(types.Int64Type)
	} else {
		var contactGroupsDiags diag.Diagnostics
		contactGroups, contactGroupsDiags = types.ListValue(types.Int64Type, contactGroupValues)
		diags.Append(contactGroupsDiags...)
		if diags.HasError() {
			return
		}
	}

	return
}

// Emails and ContactGroups are in the json together but presented in terraform seperately.
// Extract only the Emails here based on RecipientType.ID != 1 (not ContactGroup).
// Emails could be 2(Email) or 0 (Contact).
func extractEmailValues(recipients []models.RecipientJSON) (emails []attr.Value) {
	for _, recipient := range recipients {
		if recipient.RecipientType.ID != 1 && recipient.Email != cptypes.EmptyString {
			emails = append(emails, types.StringValue(recipient.Email))
		}
	}
	return
}

// Emails and ContactGroups are in the json together but presented in terraform seperately.
// Extract only the ContactGroup IDs here based on RecipientType.ID == 1 (ContactGroup).
func extractContactGroupValues(recipients []models.RecipientJSON) (ids []attr.Value) {
	for _, recipient := range recipients {
		if recipient.RecipientType.ID == 1 && recipient.ID != nil {
			ids = append(ids, types.Int64Value(int64(*recipient.ID)))
		}
	}
	return
}

func intSliceToInt64List(intSlice []int) (intList types.List, diags diag.Diagnostics) {
	int64Slice := make([]types.Int64, len(intSlice))
	for i, v := range intSlice {
		int64Slice[i] = types.Int64Value(int64(v))
	}

	attrValues := make([]attr.Value, len(int64Slice))
	for i, v := range int64Slice {
		attrValues[i] = v
	}

	intList, listDiags := types.ListValue(types.Int64Type, attrValues)
	diags.Append(listDiags...)
	if diags.HasError() {
		return types.ListNull(types.Int64Type), diags
	}

	return
}

func stringValueFromIDWithLookup(source int, destination *types.String, lookupFunc func(key int) (string, bool), fieldName string) (diags diag.Diagnostics) {
	lookupValue, found := lookupFunc(source)
	if found {
		*destination = types.StringValue(lookupValue)
	} else {
		diags.Append(diag.NewWarningDiagnostic(
			"Invalid Value",
			fmt.Sprintf("Could not find value for field '%s': %d is not recognized.", fieldName, source),
		))
	}

	return
}

func stringValueFromOptionalGenericWithLookup(source *models.GenericIDNameOmitEmptyJSON, destination *types.String, lookupFunc func(key int) (string, bool), fieldName string) (diags diag.Diagnostics) {
	if source != nil && source.ID != nil {
		return stringValueFromIDWithLookup(*source.ID, destination, lookupFunc, fieldName)
	}

	return
}

func setStringValueIfNotNil(source *string, destination *types.String) {
	if source != nil {
		*destination = types.StringValue(*source)
	}
}

func setFloat64ValueIfNotNil(source *float64, destination *types.Float64) {
	if source != nil {
		*destination = types.Float64Value(*source)
	}
}

func setInt64ValueIfNotNil(source *int, destination *types.Int64) {
	if source != nil {
		*destination = types.Int64Value(int64(*source))
	}
}

func setBoolValueIfNotNil(source *bool, destination *types.Bool) {
	if source != nil {
		*destination = types.BoolValue(*source)
	}
}
