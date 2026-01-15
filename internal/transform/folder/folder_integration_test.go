package folder

import (
	"context"
	"fmt"
	"testing"

	"catchpoint-provider/internal/fields"
	"catchpoint-provider/internal/models"
	"catchpoint-provider/internal/models/resource"
	cpschema "catchpoint-provider/internal/schema"
	"catchpoint-provider/internal/testutil"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func TestJSONToTerraformFolder(t *testing.T) {
	t.Helper()

	tests := []struct {
		name     string
		input    *models.FolderJSON
		expected func(t *testing.T, result *resource.FolderResourceModel)
	}{
		{
			name: "complete folder with all fields",
			input: &models.FolderJSON{
				ID:         123,
				DivisionID: 456,
				ProductID:  789,
				Name:       "Test Folder",
				ParentID:   testutil.ToIntPtr(1011),
				AlertGroup: models.AlertGroupJSON{
					AlertSettingType: models.GenericIDNameJSON{ID: 1, Name: "override"},
				},
				AdvancedSettings: models.AdvancedSettingsJSON{
					AdvancedSettingType: models.GenericIDNameJSON{ID: 1, Name: "override"},
				},
				RequestSettings: models.RequestSettingsJSON{
					RequestSettingType: models.GenericIDNameJSON{ID: 1, Name: "override"},
				},
				ScheduleSettings: models.ScheduleSettingsJSON{
					ScheduleSettingType: models.GenericIDNameJSON{ID: 1, Name: "override"},
				},
				InsightData: models.InsightDataJSON{
					InsightSettingType: models.GenericIDNameJSON{ID: 1, Name: "override"},
				},
			},
			expected: func(t *testing.T, result *resource.FolderResourceModel) {
				t.Helper()

				testutil.AssertEqual(t, "ID", result.ID.ValueInt64(), int64(123))
				testutil.AssertEqual(t, "DivisionID", result.DivisionID.ValueInt64(), int64(456))
				testutil.AssertEqual(t, "FolderName", result.FolderName.ValueString(), "Test Folder")
				testutil.AssertEqual(t, "ProductID", result.ProductID.ValueInt64(), int64(789))
				testutil.AssertEqual(t, "ParentID", result.ParentID.ValueInt64(), int64(1011))

				// Verify nested objects are not null
				testutil.AssertNotNil(t, "AlertSettings", result.AlertSettings)
				testutil.AssertFalse(t, "AdvancedSettings", result.AdvancedSettings.IsNull())
				testutil.AssertFalse(t, "RequestSettings", result.RequestSettings.IsNull())
				testutil.AssertFalse(t, "ScheduleSettings", result.ScheduleSettings.IsNull())
				testutil.AssertFalse(t, "Insights", result.Insights.IsNull())
			},
		},
		{
			name: "minimal folder with only required fields",
			input: &models.FolderJSON{
				ID:               456,
				DivisionID:       789,
				Name:             "Minimal Folder",
				AlertGroup:       models.AlertGroupJSON{},
				AdvancedSettings: models.AdvancedSettingsJSON{},
				RequestSettings:  models.RequestSettingsJSON{},
				ScheduleSettings: models.ScheduleSettingsJSON{},
				InsightData:      models.InsightDataJSON{},
			},
			expected: func(t *testing.T, result *resource.FolderResourceModel) {
				t.Helper()

				testutil.AssertEqual(t, "ID", result.ID.ValueInt64(), int64(456))
				testutil.AssertEqual(t, "DivisionID", result.DivisionID.ValueInt64(), int64(789))
				testutil.AssertEqual(t, "FolderName", result.FolderName.ValueString(), "Minimal Folder")
				testutil.AssertTrue(t, "ParentID", result.ParentID.IsNull())

				// Nested objects should still be initialized
				testutil.AssertNotNil(t, "AlertSettings", result.AlertSettings)
				testutil.AssertFalse(t, "AdvancedSettings", result.AdvancedSettings.IsNull())
				testutil.AssertFalse(t, "RequestSettings", result.RequestSettings.IsNull())
				testutil.AssertFalse(t, "ScheduleSettings", result.ScheduleSettings.IsNull())
				testutil.AssertFalse(t, "Insights", result.Insights.IsNull())
			},
		},
		{
			name: "folder with zero optional IDs",
			input: &models.FolderJSON{
				ID:               123,
				DivisionID:       456,
				Name:             "Zero IDs Folder",
				AlertGroup:       models.AlertGroupJSON{},
				AdvancedSettings: models.AdvancedSettingsJSON{},
				RequestSettings:  models.RequestSettingsJSON{},
				ScheduleSettings: models.ScheduleSettingsJSON{},
				InsightData:      models.InsightDataJSON{},
			},
			expected: func(t *testing.T, result *resource.FolderResourceModel) {
				t.Helper()

				testutil.AssertEqual(t, "ID", result.ID.ValueInt64(), int64(123))
				testutil.AssertEqual(t, "DivisionID", result.DivisionID.ValueInt64(), int64(456))
				testutil.AssertEqual(t, "FolderName", result.FolderName.ValueString(), "Zero IDs Folder")
			},
		},
		{
			name: "folder with unknown status ID",
			input: &models.FolderJSON{
				ID:               789,
				DivisionID:       101112,
				Name:             "Unknown Status Folder",
				AlertGroup:       models.AlertGroupJSON{},
				AdvancedSettings: models.AdvancedSettingsJSON{},
				RequestSettings:  models.RequestSettingsJSON{},
				ScheduleSettings: models.ScheduleSettingsJSON{},
				InsightData:      models.InsightDataJSON{},
			},
			expected: func(t *testing.T, result *resource.FolderResourceModel) {
				t.Helper()

				testutil.AssertEqual(t, "ID", result.ID.ValueInt64(), int64(789))
				testutil.AssertEqual(t, "DivisionID", result.DivisionID.ValueInt64(), int64(101112))
				testutil.AssertEqual(t, "FolderName", result.FolderName.ValueString(), "Unknown Status Folder")
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()

			// This represents the user's plan, it's required for nested objects to be populated.
			state := newTestFolderPlan()

			result, diags := JSONToTerraformFolder(ctx, tt.input, state)

			// Verify no errors occurred
			testutil.AssertDiagsHasNoErrors(t, diags)

			// Run test-specific assertions
			tt.expected(t, result)
		})
	}
}

func TestJSONToTerraformFolderEmptyContext(t *testing.T) {
	input := &models.FolderJSON{
		ID:               123,
		DivisionID:       456,
		Name:             "Context Test Folder",
		AlertGroup:       models.AlertGroupJSON{},
		AdvancedSettings: models.AdvancedSettingsJSON{},
		RequestSettings:  models.RequestSettingsJSON{},
		ScheduleSettings: models.ScheduleSettingsJSON{},
		InsightData:      models.InsightDataJSON{},
	}

	// This represents the user's plan, it's required for nested objects to be populated.
	state := newTestFolderPlan()

	// Test with empty context (should still work)
	result, diags := JSONToTerraformFolder(context.Background(), input, state)

	// Verify no errors occurred
	testutil.AssertDiagsHasNoErrors(t, diags)

	testutil.AssertEqual(t, "ID", result.ID.ValueInt64(), int64(123))
	testutil.AssertEqual(t, "FolderName", result.FolderName.ValueString(), "Context Test Folder")
}

// TestTransformFolderFields focuses on the basic field transformation
func TestTransformFolderFields(t *testing.T) {
	tests := []struct {
		name     string
		input    *models.FolderJSON
		expected func(t *testing.T, result *resource.FolderResourceModel)
	}{
		{
			name: "all fields populated",
			input: &models.FolderJSON{
				ID:         123,
				DivisionID: 456,
				Name:       "Test Folder",
			},
			expected: func(t *testing.T, result *resource.FolderResourceModel) {
				t.Helper()

				testutil.AssertEqual(t, "ID", result.ID.ValueInt64(), int64(123))
				testutil.AssertEqual(t, "DivisionID", result.DivisionID.ValueInt64(), int64(456))
				testutil.AssertEqual(t, "FolderName", result.FolderName.ValueString(), "Test Folder")
			},
		},
		{
			name: "optional fields zero",
			input: &models.FolderJSON{
				ID:         123,
				DivisionID: 456,
				Name:       "Test Folder",
			},
			expected: func(t *testing.T, result *resource.FolderResourceModel) {
				t.Helper()

				testutil.AssertEqual(t, "ID", result.ID.ValueInt64(), int64(123))
				testutil.AssertEqual(t, "DivisionID", result.DivisionID.ValueInt64(), int64(456))
				testutil.AssertEqual(t, "FolderName", result.FolderName.ValueString(), "Test Folder")
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var result resource.FolderResourceModel

			diags := transformFolderFields(&result, tt.input)

			// Verify no errors occurred
			testutil.AssertDiagsHasNoErrors(t, diags)
			tt.expected(t, &result)
		})
	}
}

// TestTransformNestedFolderFields tests the nested object transformations
func TestTransformNestedFolderFields(t *testing.T) {
	input := &models.FolderJSON{
		AlertGroup: models.AlertGroupJSON{
			AlertSettingType: models.GenericIDNameJSON{ID: 1, Name: "override"},
		},
		AdvancedSettings: models.AdvancedSettingsJSON{
			AdvancedSettingType: models.GenericIDNameJSON{ID: 1, Name: "override"},
		},
		RequestSettings: models.RequestSettingsJSON{
			RequestSettingType: models.GenericIDNameJSON{ID: 1, Name: "override"},
		},
		ScheduleSettings: models.ScheduleSettingsJSON{
			ScheduleSettingType: models.GenericIDNameJSON{ID: 1, Name: "override"},
		},
		InsightData: models.InsightDataJSON{
			InsightSettingType: models.GenericIDNameJSON{ID: 1, Name: "override"},
		},
	}

	ctx := context.Background()
	var result resource.FolderResourceModel
	// This represents the user's plan, it's required for nested objects to be populated.
	plan := newTestFolderPlan()

	diags := transformNestedFolderFields(ctx, &result, input, plan)

	// Verify no errors occurred
	testutil.AssertDiagsHasNoErrors(t, diags)

	// Verify all nested objects are set
	testutil.AssertNotNil(t, "AlertSettings", result.AlertSettings)
	testutil.AssertFalse(t, "AdvancedSettings", result.AdvancedSettings.IsNull())
	testutil.AssertFalse(t, "RequestSettings", result.RequestSettings.IsNull())
	testutil.AssertFalse(t, "ScheduleSettings", result.ScheduleSettings.IsNull())
	testutil.AssertFalse(t, "Insights", result.Insights.IsNull())
}

func TestTransformRequestSettings(t *testing.T) {
	certIDs := []int{1, 2, 3}
	input := &models.FolderJSON{
		AlertGroup: models.AlertGroupJSON{
			AlertSettingType: models.GenericIDNameJSON{ID: 1, Name: "override"},
		},
		AdvancedSettings: models.AdvancedSettingsJSON{
			AdvancedSettingType: models.GenericIDNameJSON{ID: 1, Name: "override"},
		},
		RequestSettings: models.RequestSettingsJSON{
			RequestSettingType:    models.GenericIDNameJSON{ID: 1, Name: "override"},
			LibraryCertificateIDs: &certIDs,
		},
		ScheduleSettings: models.ScheduleSettingsJSON{
			ScheduleSettingType: models.GenericIDNameJSON{ID: 1, Name: "override"},
		},
		InsightData: models.InsightDataJSON{
			InsightSettingType: models.GenericIDNameJSON{ID: 1, Name: "override"},
		},
	}

	ctx := context.Background()
	var result resource.FolderResourceModel
	// This represents the user's plan, it's required for nested objects to be populated.
	plan := newTestFolderPlan()

	diags := transformNestedFolderFields(ctx, &result, input, plan)

	// Verify no errors occurred
	testutil.AssertDiagsHasNoErrors(t, diags)

	// Verify all nested objects are set
	testutil.AssertNotNil(t, "AlertSettings", result.AlertSettings)
	testutil.AssertFalse(t, "AdvancedSettings", result.AdvancedSettings.IsNull())
	testutil.AssertFalse(t, "RequestSettings", result.RequestSettings.IsNull())
	testutil.AssertFalse(t, "ScheduleSettings", result.ScheduleSettings.IsNull())
	testutil.AssertFalse(t, "Insights", result.Insights.IsNull())

	testutil.AssertEqual(t, "LibraryCertificateIDs length", len(result.RequestSettings.Attributes()[fields.LibraryCertificateIDs].(types.List).Elements()), 3)
}

func newTestFolderPlan() *resource.FolderResourceModel {
	folder := resource.FolderResourceModel{
		AlertSettings: &resource.AlertSettingsModel{
			AlertSettingType: types.StringValue("override"),
		},
		AdvancedSettingsModel: resource.AdvancedSettingsModel{
			AdvancedSettings: minimalAdvancedSettingsObject(),
		},
		RequestSettingsModel: resource.RequestSettingsModel{
			RequestSettings: minimalRequestSettingsObject(),
		},
		ScheduleSettingsModel: resource.ScheduleSettingsModel{
			ScheduleSettings: minimalScheduleSettingsObject(),
		},
		InsightSettingsModel: resource.InsightSettingsModel{
			Insights: minimalInsightSettingsObject(),
		},
	}

	return &folder
}

func minimalAdvancedSettingsObject() types.Object {
	attrTypes := cpschema.GetAdvancedSettingsAttributeTypesForProductAndFolder()
	attrs := buildNullAttrs(attrTypes)
	attrs[fields.AdvancedSettingType] = types.StringValue("inherit")

	obj, diags := types.ObjectValue(attrTypes, attrs)
	if diags.HasError() {
		panic(fmt.Sprintf("failed to create minimal advanced settings object: %v", diags))
	}
	return obj
}

func minimalRequestSettingsObject() types.Object {
	attrTypes := cpschema.GetRequestSettingsAttributeTypes()
	attrs := buildNullAttrs(attrTypes)
	attrs[fields.RequestSettingType] = types.StringValue("inherit")

	obj, diags := types.ObjectValue(attrTypes, attrs)
	if diags.HasError() {
		panic(fmt.Sprintf("failed to create minimal request settings object: %v", diags))
	}
	return obj
}

func minimalScheduleSettingsObject() types.Object {
	attrTypes := cpschema.GetScheduleSettingsAttributeTypes()
	attrs := buildNullAttrs(attrTypes)
	attrs[fields.ScheduleSettingType] = types.StringValue("inherit")

	obj, diags := types.ObjectValue(attrTypes, attrs)
	if diags.HasError() {
		panic(fmt.Sprintf("failed to create minimal schedule settings object: %v", diags))
	}
	return obj
}

func minimalInsightSettingsObject() types.Object {
	attrTypes := cpschema.GetInsightsAttributeTypes()
	attrs := buildNullAttrs(attrTypes)
	attrs[fields.InsightSettingType] = types.StringValue("inherit")

	obj, diags := types.ObjectValue(attrTypes, attrs)
	if diags.HasError() {
		panic(fmt.Sprintf("failed to create minimal insight settings object: %v", diags))
	}
	return obj
}

func buildNullAttrs(attrTypes map[string]attr.Type) map[string]attr.Value {
	attrs := make(map[string]attr.Value, len(attrTypes))
	for name, aType := range attrTypes {
		attrs[name] = createNullValueForType(aType)
	}
	return attrs
}

func createNullValueForType(attrType attr.Type) attr.Value {
	switch attrType {
	case types.StringType:
		return types.StringNull()
	case types.BoolType:
		return types.BoolNull()
	case types.Int64Type:
		return types.Int64Null()
	case types.Float64Type:
		return types.Float64Null()
	default:
		// For complex types, need to handle differently
		switch t := attrType.(type) {
		case types.ListType:
			return types.ListNull(t.ElemType)
		case types.SetType:
			return types.SetNull(t.ElemType)
		case types.MapType:
			return types.MapNull(t.ElemType)
		case types.ObjectType:
			return types.ObjectNull(t.AttrTypes)
		default:
			return types.StringNull() // Fallback
		}
	}
}
