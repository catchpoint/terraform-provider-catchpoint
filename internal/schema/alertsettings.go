package schema

import (
	"context"
	"fmt"
	"sync"

	"catchpoint-provider/internal/fields"
	"catchpoint-provider/internal/modifier"
	cptypes "catchpoint-provider/internal/types"
	"catchpoint-provider/internal/validation"
	"catchpoint-provider/internal/validation/alerttypes"

	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var (
	alertSettingTypeValidator   = oneOfStringValidator(cptypes.ValidAlertSettingTypeNames)
	historicalIntervalValidator = oneOfStringValidator(cptypes.ValidHistoricalIntervalNames)
	nodeThresholdTypeValidator  = oneOfStringValidator(cptypes.ValidNodeThresholdTypes)
	notificationTypeValidator   = oneOfStringValidator(cptypes.ValidNotificationTypeNames)
	operationTypeValidator      = oneOfStringValidator(cptypes.ValidOperationTypeNames)
	reminderValidator           = oneOfStringValidator(cptypes.ValidReminderNames)
	statisticalTypeValidator    = oneOfStringValidator(cptypes.ValidStatisticalTypeNames)
	thresholdIntervalValidator  = oneOfStringValidator(cptypes.ValidThresholdIntervalNames)
	triggerTypeValidator        = oneOfStringValidator(cptypes.ValidTriggerTypeNames)
	dnsRecordTypeValidator      = oneOfStringValidator(cptypes.ValidDNSRecordTypes)
	filterTypeValidator         = oneOfStringValidator(cptypes.ValidFilterTypes)

	alertSettingsProductFolderSchemaOnce sync.Once
	alertSettingsProductFolderBlocks     map[string]schema.Block
)

// Folder and Product alert settings always allow for all possible alert types and subtypes.
func BuildAlertSettingsBlockForProductAndFolder(ctx context.Context) map[string]schema.Block {
	alertSettingsProductFolderSchemaOnce.Do(func() {
		alertTypeValidator := oneOfStringValidator(cptypes.ValidAlertTypeNames)
		alertSubTypeValidator := oneOfStringValidator(cptypes.ValidAlertSubTypeNames)

		alertSettingsProductFolderBlocks = buildAlertSettingsBlock(ctx, alertTypeValidator, alertSubTypeValidator)
	})

	return alertSettingsProductFolderBlocks
}

func BuildAlertSettingsBlockForTest(ctx context.Context, testType cptypes.TestType) map[string]schema.Block {
	// Get all possible combinations of alert types and subtypes for the given test type.
	matrix := alerttypes.GetTestCompatibilityMatrix(testType)

	alertTypeValidator := stringvalidator.OneOf(matrix.ListAllValidAlertTypes()...)
	alertSubTypeValidator := stringvalidator.OneOf(matrix.ListAllValidAlertSubTypes()...)

	return buildAlertSettingsBlock(ctx, alertTypeValidator, alertSubTypeValidator)
}

func GetAlertSettingsAttributeTypes() map[string]attr.Type {
	// Use the cached schema instead of rebuilding
	schemaBlocks := BuildAlertSettingsBlockForProductAndFolder(context.Background())

	// Extract the inner attribute types from the cached schema
	if alertSettingsBlock, exists := schemaBlocks[fields.AlertSettings]; exists {
		if singleNested, ok := alertSettingsBlock.(schema.SingleNestedBlock); ok {
			return ExtractAllAttributeTypes(singleNested.Attributes, singleNested.Blocks)
		}
	}
	return make(map[string]attr.Type)
}

func GetAlertRulesAttributeTypes() map[string]attr.Type {
	// Use the cached schema instead of rebuilding
	schemaBlocks := BuildAlertSettingsBlockForProductAndFolder(context.Background())

	// Get the alert_settings block
	if alertSettingsBlock, exists := schemaBlocks[fields.AlertSettings]; exists {
		if singleNested, ok := alertSettingsBlock.(schema.SingleNestedBlock); ok {
			// Get the alert_rule nested block from within alert_settings
			if alertRuleBlock, exists := singleNested.Blocks[fields.AlertRule]; exists {
				if setNestedBlock, ok := alertRuleBlock.(schema.SetNestedBlock); ok {
					return ExtractAllAttributeTypes(
						setNestedBlock.NestedObject.Attributes,
						setNestedBlock.NestedObject.Blocks,
					)
				}
			}
		}
	}
	return make(map[string]attr.Type)
}

func GetNotificationGroupAttributeTypes() map[string]attr.Type {
	// Use the cached schema instead of rebuilding
	schemaBlocks := BuildAlertSettingsBlockForProductAndFolder(context.Background())

	// Get the alert_settings block
	if alertSettingsBlock, exists := schemaBlocks[fields.AlertSettings]; exists {
		if singleNested, ok := alertSettingsBlock.(schema.SingleNestedBlock); ok {
			// Get the notification_group nested block from within alert_settings
			if notificationGroupBlock, exists := singleNested.Blocks[fields.NotificationGroup]; exists {
				if singleNestedBlock, ok := notificationGroupBlock.(schema.SingleNestedBlock); ok {
					return ExtractAllAttributeTypes(
						singleNestedBlock.Attributes,
						singleNestedBlock.Blocks,
					)
				}
			}
		}
	}
	return make(map[string]attr.Type)
}

func GetAlertRulesElementType() attr.Type {
	return types.ObjectType{
		AttrTypes: GetAlertRulesAttributeTypes(),
	}
}

func GetNotificationGroupElementType() attr.Type {
	return types.ObjectType{
		AttrTypes: GetNotificationGroupAttributeTypes(),
	}
}

func buildAlertSettingsBlock(ctx context.Context, alertTypeValidator, alertSubTypeValidator validator.String) map[string]schema.Block {
	return map[string]schema.Block{
		fields.AlertSettings: schema.SingleNestedBlock{
			Description: `Used for overriding the alert settings section. 
			Note: omitting this field entirely means that the alert settings for this object will not be tracked.
			If you specifically want the object to inherit the parent alert settings then supply an object with only the alert_setting_type 
			field, e.g. '{ alert_setting_type = "Inherit" }'.`,
			Attributes: map[string]schema.Attribute{
				fields.AlertSettingType: schema.StringAttribute{
					Optional:    true,
					Computed:    true,
					Description: "Set the alert setting type, " + alertSettingTypeValidator.Description(ctx),
					Validators:  []validator.String{alertSettingTypeValidator},
				},
			},
			Blocks: map[string]schema.Block{
				fields.AlertRule: schema.ListNestedBlock{
					Description: "Sets the alert rule with attributes such as threshold, trigger type, warning, critical trigger and more",
					NestedObject: schema.NestedBlockObject{
						Attributes: map[string]schema.Attribute{
							fields.NodeThresholdType: schema.StringAttribute{
								//Required
								Required:    true,
								Description: "Sets the alert node threshold type, " + nodeThresholdTypeValidator.Description(ctx),
								Validators:  []validator.String{nodeThresholdTypeValidator},
							},
							fields.ThresholdNumberOfRuns: schema.Int64Attribute{
								Optional:    true,
								Description: "Sets the threshold for the number of runs or nodes the alert should trigger",
							},
							fields.ThresholdPercentageOfRuns: schema.Float64Attribute{
								Optional:    true,
								Description: "Sets the threshold for the percentage of runs the alert should trigger",
							},
							fields.NumberOfFailingNodes: schema.Int64Attribute{
								Optional:    true,
								Description: "Sets the number of failed nodes the alert should trigger if node_threshold_type is 'average across nodes'",
							},
							fields.TriggerType: schema.StringAttribute{
								Optional:    true,
								Computed:    true,
								Description: "Sets the alert trigger type, " + triggerTypeValidator.Description(ctx),
								Validators:  []validator.String{triggerTypeValidator},
							},
							fields.OperationType: schema.StringAttribute{
								Optional:    true,
								Computed:    true, // This can be set by the user but defaults to NotEquals for some alert types.
								Description: "Sets the alert operation type, " + operationTypeValidator.Description(ctx),
								Validators:  []validator.String{operationTypeValidator},
							},
							fields.StatisticalType: schema.StringAttribute{
								Optional:    true,
								Computed:    true,
								Description: "Sets the alert statistical type, " + statisticalTypeValidator.Description(ctx),
								Validators:  []validator.String{statisticalTypeValidator},
							},
							fields.HistoricalInterval: schema.StringAttribute{
								Optional:    true,
								Computed:    true,
								Description: "Sets the alert historical interval, " + historicalIntervalValidator.Description(ctx),
								Validators:  []validator.String{historicalIntervalValidator},
							},
							fields.WarningTrigger: schema.Float64Attribute{
								Optional:    true,
								Description: "Warning trigger value for 'specific value' and 'trailing value' trigger types.",
							},
							fields.CriticalTrigger: schema.Float64Attribute{
								Optional:    true,
								Description: "Critical trigger value for 'specific value' and 'trailing value' trigger types.",
							},
							fields.EnableConsecutive: schema.BoolAttribute{
								Optional:    true,
								Computed:    true,
								Description: "Checks consecutive number of runs or nodes for triggering alerts.",
							},
							fields.ConsecutiveNumberOfRuns: schema.Int64Attribute{
								Optional:    true,
								Description: "Sets the number of consecutive runs only if enable_consecutive field is true and node_threshold_type is node",
							},
							fields.Expression: schema.StringAttribute{
								Optional:    true,
								Description: "Sets trigger expression for content match alert type ",
								Validators:  []validator.String{stringvalidator.LengthAtLeast(1)},
							},
							fields.WarningReminder: schema.StringAttribute{
								Optional:    true,
								Computed:    true,
								Description: "Sets the alert warning reminder interval, " + reminderValidator.Description(ctx),
								Validators:  []validator.String{reminderValidator},
							},
							fields.CriticalReminder: schema.StringAttribute{
								Optional:    true,
								Computed:    true,
								Description: "Sets the alert critical reminder interval, " + reminderValidator.Description(ctx),
								Validators:  []validator.String{reminderValidator},
							},
							fields.ThresholdInterval: schema.StringAttribute{
								Optional:    true,
								Computed:    true,
								Default:     stringdefault.StaticString(cptypes.Default),
								Description: "Sets the alert threshold interval, " + thresholdIntervalValidator.Description(ctx),
								Validators:  []validator.String{thresholdIntervalValidator},
							},
							fields.UseRollingWindow: schema.BoolAttribute{
								Optional:    true,
								Computed:    true,
								Description: "Set to true for using rolling window instead of schedule time threshold",
							},
							fields.NotificationType: schema.StringAttribute{
								Optional:    true,
								Computed:    true,
								Description: "Sets the alert notification type, " + notificationTypeValidator.Description(ctx),
								Validators:  []validator.String{notificationTypeValidator},
							},
							fields.AlertType: schema.StringAttribute{
								Required: true,
								// Do not use the .Description() method because it randomly sorts the values differently each time.
								// This looks terrible in the documentation and causes unnecessary diffs.
								// This only happens with alertType and alertSubType though, (likely because they're dynamically generated).
								Description: fmt.Sprintf("Sets the alert type, %v", cptypes.ValidAlertTypeNames),
								Validators:  []validator.String{alertTypeValidator},
							},
							fields.AlertSubType: schema.StringAttribute{
								Optional: true,
								// Do not use the .Description() method because it randomly sorts the values differently each time.
								// This looks terrible in the documentation and causes unnecessary diffs.
								Description: fmt.Sprintf("Sets the alert sub type, %v", cptypes.ValidAlertSubTypeNames),
								Validators:  []validator.String{alertSubTypeValidator},
							},
							fields.EnforceTestFailure: schema.BoolAttribute{
								Optional:    true,
								Computed:    true,
								Description: "Sets enforce test failure property for an alert",
							},
							fields.OmitScatterplot: schema.BoolAttribute{
								Optional:    true,
								Computed:    true,
								Description: "Omits scatterplot image from alert emails if set to true",
							},
							fields.DNSResolvedName: schema.StringAttribute{
								Optional:    true,
								Description: "Sets the DNS resolved name for DNS alert types",
							},
							fields.DNSttl: schema.Int64Attribute{
								Optional:    true,
								Description: "Sets the DNS TTL (time to live) for DNS alert types",
							},
							fields.DNSRecordType: schema.StringAttribute{
								Optional:    true,
								Description: "Sets the DNS record type for DNS alert types, " + dnsRecordTypeValidator.Description(ctx),
								Validators:  []validator.String{dnsRecordTypeValidator},
							},
							fields.AllMatchRecords: schema.BoolAttribute{
								Optional:    true,
								Computed:    true,
								Description: "Set to true to match all DNS records when evaluating DNS alerts",
							},
						},
						Blocks: map[string]schema.Block{
							fields.Level: schema.SingleNestedBlock{
								Description: "Filter configuration for the alert level",
								Attributes: map[string]schema.Attribute{
									"filter_type": schema.StringAttribute{
										Optional:    true,
										Computed:    true,
										Description: "The type of filter to apply, " + filterTypeValidator.Description(ctx),
										Validators:  []validator.String{filterTypeValidator},
									},
									"filter_value": schema.StringAttribute{
										Optional:    true,
										Computed:    true,
										Description: "The value for the filter",
									},
								},
							},
							fields.NotificationGroup: schema.ListNestedBlock{
								Description: "List of Notification groups for configuring alert notifications, including recipients' email addresses and alert settings.",
								Validators: []validator.List{
									listvalidator.SizeAtMost(5),
								},
								NestedObject: schema.NestedBlockObject{
									Attributes: map[string]schema.Attribute{
										fields.NotifyOnWarning: schema.BoolAttribute{
											Optional:    true,
											Computed:    true,
											Default:     booldefault.StaticBool(false),
											Description: "Set to true to include warning alerts in notifications.",
										},
										fields.NotifyOnCritical: schema.BoolAttribute{
											Optional:    true,
											Computed:    true,
											Default:     booldefault.StaticBool(false),
											Description: "Set to true to include critical alerts in notifications.",
										},
										fields.NotifyOnImproved: schema.BoolAttribute{
											Optional:    true,
											Computed:    true,
											Default:     booldefault.StaticBool(false),
											Description: "Set to true to include improved alerts in notifications.",
										},
										fields.Subject: schema.StringAttribute{
											Optional:    true,
											Description: "Email subject for the alert notifications.",
										},
										fields.AlertWebhookIDs: schema.ListAttribute{
											ElementType: types.Int64Type,
											// This exists in the API but is not settable by the user at this level.
											Optional:    false,
											Required:    false,
											Computed:    true,
											Description: "Alert webhook ids for the webhook endpoints to associate with this alert setting.",
										},
										fields.RecipientEmails: schema.ListAttribute{
											ElementType: types.StringType,
											Optional:    true,
											Description: "List of email addresses to receive alert notifications.",
										},
										fields.ContactGroupIDs: schema.ListAttribute{
											ElementType: types.Int64Type,
											Optional:    true,
											Description: "List of contact groups to receive alert notifications.",
										},
									},
								},
							},
						},
					},
					Validators: []validator.List{
						validation.AlertRuleThresholdsValidator(),
					},
				},
				fields.NotificationGroup: schema.SingleNestedBlock{
					Description: "Notification group for setting up alert recipients, adding alert webhook ids.",
					Attributes: map[string]schema.Attribute{
						fields.NotifyOnWarning: schema.BoolAttribute{
							// This exists in the API but is not settable by the user at this level.
							Optional: false,
							Required: false,
							Computed: true,
							Default:  booldefault.StaticBool(true),
						},
						fields.NotifyOnCritical: schema.BoolAttribute{
							// This exists in the API but is not settable by the user at this level.
							Optional: false,
							Required: false,
							Computed: true,
							Default:  booldefault.StaticBool(true),
						},
						fields.NotifyOnImproved: schema.BoolAttribute{
							// This exists in the API but is not settable by the user at this level.
							Optional: false,
							Required: false,
							Computed: true,
							Default:  booldefault.StaticBool(true),
						},
						fields.Subject: schema.StringAttribute{
							Optional:    true,
							Description: "Email subject for the alert notifications.",
						},
						fields.AlertWebhookIDs: schema.ListAttribute{
							ElementType: types.Int64Type,
							Optional:    true,
							Computed:    true,
							Default:     listdefault.StaticValue(types.ListValueMust(types.Int64Type, []attr.Value{})),
							Description: "Alert webhook ids for the webhook endpoints to associate with this alert setting.",
						},
						fields.RecipientEmails: schema.ListAttribute{
							ElementType: types.StringType,
							Optional:    true,
							Description: "List of emails to alert.",
						},
						fields.ContactGroupIDs: schema.ListAttribute{
							ElementType: types.Int64Type,
							Optional:    true,
							Description: "A set of contact groups to receive alert notifications.",
						},
					},
				},
			},
			Validators: []validator.Object{
				validation.RequireNotificationGroupValidator(),
			},
			PlanModifiers: []planmodifier.Object{
				modifier.InheritBlockPlanModifier(fields.AlertSettingType),
			},
		},
	}
}
