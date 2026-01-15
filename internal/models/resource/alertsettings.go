package resource

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// The AlertSettingsModel defines the alertGroup values for all resources - Product, Folder, Test.
// This is available for all supported test types.
type AlertSettingsModel struct {
	AlertSettingType  types.String            `tfsdk:"alert_setting_type"`
	AlertRule         []AlertRuleModel        `tfsdk:"alert_rule"`         // SetNestedBlock
	NotificationGroup *NotificationGroupModel `tfsdk:"notification_group"` // SingleNestedBlock
}

type NotificationGroupModel struct {
	NotifyOnWarning  types.Bool   `tfsdk:"notify_on_warning"`
	NotifyOnCritical types.Bool   `tfsdk:"notify_on_critical"`
	NotifyOnImproved types.Bool   `tfsdk:"notify_on_improved"`
	Subject          types.String `tfsdk:"subject"`
	AlertWebhookIDs  types.List   `tfsdk:"alert_webhook_ids"`
	Emails           types.List   `tfsdk:"recipient_emails"`
	ContactGroupIDs  types.List   `tfsdk:"contact_group_ids"`
}

type AlertRuleModel struct {
	AlertSubType       types.String             `tfsdk:"alert_sub_type"`
	AlertType          types.String             `tfsdk:"alert_type"`
	NotificationType   types.String             `tfsdk:"notification_type"`
	EnforceTestFailure types.Bool               `tfsdk:"enforce_test_failure"`
	OmitScatterplot    types.Bool               `tfsdk:"omit_scatterplot"`
	NotificationGroup  []NotificationGroupModel `tfsdk:"notification_group"` // SetNestedBlock
	AllMatchRecords    types.Bool               `tfsdk:"all_match_records"`
	NodeThresholdModel
	TriggerModel
}

type NodeThresholdModel struct {
	ConsecutiveNumberOfRuns   types.Int64   `tfsdk:"consecutive_number_of_runs"`
	EnableConsecutive         types.Bool    `tfsdk:"enable_consecutive"`
	NodeThresholdType         types.String  `tfsdk:"node_threshold_type"`
	ThresholdNumberOfRuns     types.Int64   `tfsdk:"threshold_number_of_runs"`
	ThresholdPercentageOfRuns types.Float64 `tfsdk:"threshold_percentage_of_runs"`
	NumberOfFailingNodes      types.Int64   `tfsdk:"number_of_failing_nodes"`
	// TODO: we were setting this to false always in sdkv2. We should expose it in the schema and set it here.
	//UseHistoricalAverage      types.Bool    `tfsdk:"use_historical_average"`
}

type TriggerModel struct {
	UseRollingWindow   types.Bool    `tfsdk:"use_rolling_window"`
	Expression         types.String  `tfsdk:"expression"`
	TriggerType        types.String  `tfsdk:"trigger_type"`
	OperationType      types.String  `tfsdk:"operation_type"`
	StatisticalType    types.String  `tfsdk:"statistical_type"`
	HistoricalInterval types.String  `tfsdk:"historical_interval"`
	WarningTrigger     types.Float64 `tfsdk:"warning_trigger"`
	CriticalTrigger    types.Float64 `tfsdk:"critical_trigger"`
	WarningReminder    types.String  `tfsdk:"warning_reminder"`
	CriticalReminder   types.String  `tfsdk:"critical_reminder"`
	ThresholdInterval  types.String  `tfsdk:"threshold_interval"`
	DNSResolvedName    types.String  `tfsdk:"dns_resolved_name"`
	DNSTTL             types.Int64   `tfsdk:"dns_ttl"`
	DNSRecordType      types.String  `tfsdk:"dns_record_type"`
	Level              *LevelModel   `tfsdk:"level"` // SingleNestedBlock
}

type LevelModel struct {
	FilterType  types.String `tfsdk:"filter_type"`
	FilterValue types.String `tfsdk:"filter_value"`
}
