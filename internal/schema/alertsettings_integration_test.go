package schema

import (
	"context"
	"testing"

	"catchpoint-provider/internal/fields"
	"catchpoint-provider/internal/testutil"
	cptypes "catchpoint-provider/internal/types"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func TestBuildAlertSettingsAttributesForProductAndFolder(t *testing.T) {
	attrs := BuildAlertSettingsBlockForProductAndFolder(context.Background())

	// Should have exactly one attribute
	testutil.AssertEqual(t, "should have one attribute", len(attrs), 1)

	// Should contain AlertSettings field
	alertSettings, exists := attrs[fields.AlertSettings]
	testutil.AssertEqual(t, "should contain AlertSettings field", exists, true)

	// Should be SingleNestedBlock
	singleNested, ok := alertSettings.(schema.SingleNestedBlock)
	testutil.AssertEqual(t, "should be SingleNestedBlock", ok, true)

	// Verify nested attributes exist
	nestedAttrs := singleNested.Attributes
	testutil.AssertEqual(t, "should have 1 nested attributes", len(nestedAttrs), 1)

	nestedBlocks := singleNested.Blocks
	testutil.AssertEqual(t, "should have 2 nested blocks", len(nestedBlocks), 2)

	// Verify main nested fields exist
	_, hasAlertSettingType := nestedAttrs[fields.AlertSettingType]
	testutil.AssertEqual(t, "should have AlertSettingType", hasAlertSettingType, true)

	_, hasAlertRule := nestedBlocks[fields.AlertRule]
	testutil.AssertEqual(t, "should have AlertRule", hasAlertRule, true)

	_, hasNotificationGroup := nestedBlocks[fields.NotificationGroup]
	testutil.AssertEqual(t, "should have NotificationGroup", hasNotificationGroup, true)
}

func TestBuildAlertSettingsAttributesForTest(t *testing.T) {
	ctx := context.Background()
	testType := cptypes.APIType

	attrs := BuildAlertSettingsBlockForTest(ctx, testType)

	// Should have exactly one attribute
	testutil.AssertEqual(t, "should have one attribute", len(attrs), 1)

	// Should contain AlertSettings field
	alertSettings, exists := attrs[fields.AlertSettings]
	testutil.AssertEqual(t, "should contain AlertSettings field", exists, true)

	// Should be SingleNestedBlock
	singleNested, ok := alertSettings.(schema.SingleNestedBlock)
	testutil.AssertEqual(t, "should be SingleNestedBlock", ok, true)

	// Should have same structure as product/folder version
	nestedAttrs := singleNested.Attributes
	testutil.AssertEqual(t, "should have 1 nested attributes", len(nestedAttrs), 1)
}

func TestGetAlertSettingsAttributeTypes(t *testing.T) {
	attrTypes := GetAlertSettingsAttributeTypes()

	// Should return map of attribute types
	testutil.AssertNotNil(t, "attribute types", attrTypes)
	testutil.AssertEqual(t, "should have 3 attributes", len(attrTypes), 3)

	// Verify expected fields exist
	expectedAttributes := []string{
		fields.AlertSettingType,
		fields.AlertRule,
		fields.NotificationGroup,
	}

	for _, field := range expectedAttributes {
		attrType, exists := attrTypes[field]
		testutil.AssertEqual(t, "should contain field "+field, exists, true)
		testutil.AssertNotNil(t, "attribute type"+field, attrType)
	}

	// Verify specific types
	testutil.AssertEqual(t, "alert setting type should be StringType", attrTypes[fields.AlertSettingType], attr.Type(types.StringType))

	// AlertRule should be SetType
	alertRuleType, ok := attrTypes[fields.AlertRule].(types.ListType)
	testutil.AssertEqual(t, "alert rule should be ListType", ok, true)
	testutil.AssertNotNil(t, "alert rule element type", alertRuleType.ElemType)

	// NotificationGroup should be ObjectType
	notificationGroupType, ok := attrTypes[fields.NotificationGroup].(types.ObjectType)
	testutil.AssertEqual(t, "notification group", ok, true)
	testutil.AssertNotNil(t, "notification group attr types", notificationGroupType.AttrTypes)
}

func TestGetAlertRulesElementType(t *testing.T) {
	elementType := GetAlertRulesElementType()

	// Should return ObjectType
	objectType, ok := elementType.(types.ObjectType)
	testutil.AssertEqual(t, "should be ObjectType", ok, true)

	// Should have attribute types
	testutil.AssertNotNil(t, "attr types should not be nil", objectType.AttrTypes)
}

func TestGetNotificationGroupAttributeTypes(t *testing.T) {
	attrTypes := GetNotificationGroupAttributeTypes()

	// Should return map of attribute types for notification group
	testutil.AssertNotNil(t, "attribute types should not be nil", attrTypes)
	testutil.AssertEqual(t, "should have correct number of attributes", len(attrTypes), 7)

	// Verify expected notification group fields exist
	expectedFields := []string{
		fields.NotifyOnWarning,
		fields.NotifyOnCritical,
		fields.NotifyOnImproved,
		fields.Subject,
		fields.AlertWebhookIDs,
		fields.RecipientEmails,
		fields.ContactGroupIDs,
	}

	for _, field := range expectedFields {
		attrType, exists := attrTypes[field]
		testutil.AssertEqual(t, "should contain field "+field, exists, true)
		testutil.AssertNotNil(t, "attribute type should not be nil for "+field, attrType)
	}

	// Verify specific types
	testutil.AssertEqual(t, "notify on warning should be BoolType", attrTypes[fields.NotifyOnWarning], attr.Type(types.BoolType))
	testutil.AssertEqual(t, "subject should be StringType", attrTypes[fields.Subject], attr.Type(types.StringType))

	// Verify list types
	webhookIDsType := types.ListType{ElemType: types.Int64Type}
	testutil.AssertEqual(t, "webhook ids should be ListType with Int64 elements", attrTypes[fields.AlertWebhookIDs], attr.Type(webhookIDsType))

	emailsType := types.ListType{ElemType: types.StringType}
	testutil.AssertEqual(t, "emails should be ListType with String elements", attrTypes[fields.RecipientEmails], attr.Type(emailsType))
}

func TestGetNotificationGroupElementType(t *testing.T) {
	elementType := GetNotificationGroupElementType()

	// Should return ObjectType
	objectType, ok := elementType.(types.ObjectType)
	testutil.AssertEqual(t, "should be ObjectType", ok, true)

	// Should have attribute types
	testutil.AssertNotNil(t, "attr types should not be nil", objectType.AttrTypes)
	testutil.AssertEqual(t, "should have 7 attr types", len(objectType.AttrTypes), 7)
}

func TestGetAlertSettingsProductFolderSchema(t *testing.T) {
	schemaAttrs := GetAlertSettingsAttributeTypes()

	// Should return map of schema attributes
	testutil.AssertNotNil(t, "schema should not be nil", schemaAttrs)
	testutil.AssertEqual(t, "should have 3 attributes", len(schemaAttrs), 3)

	// Verify expected fields exist
	expectedFields := []string{
		fields.AlertSettingType,
		fields.AlertRule,
		fields.NotificationGroup,
	}

	for _, field := range expectedFields {
		attr, exists := schemaAttrs[field]
		testutil.AssertEqual(t, "should contain field "+field, exists, true)
		testutil.AssertNotNil(t, "attribute should not be nil for "+field, attr)
	}
}

func TestGetAlertSettingsProductFolderSchemaConsistency(t *testing.T) {
	// Get schema multiple times to ensure consistency (tests sync.Once)
	schema1 := BuildAlertSettingsBlockForProductAndFolder(context.Background())
	schema2 := BuildAlertSettingsBlockForProductAndFolder(context.Background())

	testutil.AssertEqual(t, "should return same number of attributes", len(schema1), len(schema2))

	// Verify all fields are the same
	for field := range schema1 {
		_, exists := schema2[field]
		testutil.AssertEqual(t, "field should exist in both calls for "+field, exists, true)
	}
}
