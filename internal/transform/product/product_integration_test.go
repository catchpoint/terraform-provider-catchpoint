package product

import (
	"context"
	"testing"

	"catchpoint-provider/internal/fields"
	"catchpoint-provider/internal/models"
	"catchpoint-provider/internal/models/resource"
	"catchpoint-provider/internal/testutil"
	cptypes "catchpoint-provider/internal/types"

	"github.com/hashicorp/terraform-plugin-framework/types"
)

func TestJSONToTerraformProduct(t *testing.T) {
	t.Helper()

	tests := []struct {
		name     string
		input    *models.ProductJSON
		expected func(t *testing.T, result *resource.ProductResourceModel)
	}{
		{
			name: "complete product with all fields",
			input: &models.ProductJSON{
				ID:                123,
				DivisionID:        456,
				Name:              "Test Product",
				Status:            models.GenericIDNameJSON{ID: 0, Name: "active"},
				AlertGroupID:      789,
				TestDataWebhookID: 101112,
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
			expected: func(t *testing.T, result *resource.ProductResourceModel) {
				t.Helper()

				testutil.AssertEqual(t, "ID", result.ID.ValueInt64(), int64(123))
				testutil.AssertEqual(t, "DivisionID", result.DivisionID.ValueInt64(), int64(456))
				testutil.AssertEqual(t, "ProductName", result.ProductName.ValueString(), "Test Product")
				testutil.AssertEqual(t, "Status", result.Status.ValueString(), cptypes.Active)
				testutil.AssertEqual(t, "AlertGroupID", result.AlertGroupID.ValueInt64(), int64(789))
				testutil.AssertEqual(t, "TestDataWebhookID", result.TestDataWebhookID.ValueInt64(), int64(101112))

				// Verify nested objects are not null
				testutil.AssertNotNil(t, "AlertSettings", result.AlertSettings)
				testutil.AssertFalse(t, "AdvancedSettings", result.AdvancedSettings.IsNull())
				testutil.AssertFalse(t, "RequestSettings", result.RequestSettings.IsNull())
				testutil.AssertFalse(t, "ScheduleSettings", result.ScheduleSettings.IsNull())
				testutil.AssertFalse(t, "Insights", result.Insights.IsNull())
			},
		},
		{
			name: "minimal product with only required fields",
			input: &models.ProductJSON{
				ID:               456,
				DivisionID:       789,
				Name:             "Minimal Product",
				Status:           models.GenericIDNameJSON{ID: 0, Name: "active"},
				AlertGroup:       models.AlertGroupJSON{},
				AdvancedSettings: models.AdvancedSettingsJSON{},
				RequestSettings:  models.RequestSettingsJSON{},
				ScheduleSettings: models.ScheduleSettingsJSON{},
				InsightData:      models.InsightDataJSON{},
			},
			expected: func(t *testing.T, result *resource.ProductResourceModel) {
				t.Helper()

				testutil.AssertEqual(t, "ID", result.ID.ValueInt64(), int64(456))
				testutil.AssertEqual(t, "DivisionID", result.DivisionID.ValueInt64(), int64(789))
				testutil.AssertEqual(t, "ProductName", result.ProductName.ValueString(), "Minimal Product")
				testutil.AssertEqual(t, "Status", result.Status.ValueString(), cptypes.Active)

				// Optional fields should be null
				testutil.AssertTrue(t, "AlertGroupID", result.AlertGroupID.IsNull())
				testutil.AssertTrue(t, "TestDataWebhookID", result.TestDataWebhookID.IsNull())

				// Nested objects should still be initialized
				testutil.AssertNotNil(t, "AlertSettings", result.AlertSettings)
				testutil.AssertFalse(t, "AdvancedSettings", result.AdvancedSettings.IsNull())
				testutil.AssertFalse(t, "RequestSettings", result.RequestSettings.IsNull())
				testutil.AssertFalse(t, "ScheduleSettings", result.ScheduleSettings.IsNull())
				testutil.AssertFalse(t, "Insights", result.Insights.IsNull())
			},
		},
		{
			name: "product with zero optional IDs",
			input: &models.ProductJSON{
				ID:                123,
				DivisionID:        456,
				Name:              "Zero IDs Product",
				Status:            models.GenericIDNameJSON{ID: 0, Name: "active"},
				AlertGroupID:      0, // Should result in null
				TestDataWebhookID: 0, // Should result in null
				AlertGroup:        models.AlertGroupJSON{},
				AdvancedSettings:  models.AdvancedSettingsJSON{},
				RequestSettings:   models.RequestSettingsJSON{},
				ScheduleSettings:  models.ScheduleSettingsJSON{},
				InsightData:       models.InsightDataJSON{},
			},
			expected: func(t *testing.T, result *resource.ProductResourceModel) {
				t.Helper()

				testutil.AssertEqual(t, "ID", result.ID.ValueInt64(), int64(123))
				testutil.AssertEqual(t, "DivisionID", result.DivisionID.ValueInt64(), int64(456))
				testutil.AssertEqual(t, "ProductName", result.ProductName.ValueString(), "Zero IDs Product")

				// Zero IDs should be null
				testutil.AssertTrue(t, "AlertGroupID", result.AlertGroupID.IsNull())
				testutil.AssertTrue(t, "TestDataWebhookID", result.TestDataWebhookID.IsNull())
			},
		},
		{
			name: "product with unknown status ID",
			input: &models.ProductJSON{
				ID:               789,
				DivisionID:       101112,
				Name:             "Unknown Status Product",
				Status:           models.GenericIDNameJSON{ID: 99999}, // Unknown status ID
				AlertGroup:       models.AlertGroupJSON{},
				AdvancedSettings: models.AdvancedSettingsJSON{},
				RequestSettings:  models.RequestSettingsJSON{},
				ScheduleSettings: models.ScheduleSettingsJSON{},
				InsightData:      models.InsightDataJSON{},
			},
			expected: func(t *testing.T, result *resource.ProductResourceModel) {
				t.Helper()

				testutil.AssertEqual(t, "ID", result.ID.ValueInt64(), int64(789))
				testutil.AssertEqual(t, "DivisionID", result.DivisionID.ValueInt64(), int64(101112))
				testutil.AssertEqual(t, "ProductName", result.ProductName.ValueString(), "Unknown Status Product")

				// Unknown status should default to Active
				testutil.AssertEqual(t, "Status", result.Status.ValueString(), cptypes.Active)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()

			// This represents the user's plan, it's required for nested objects to be populated.
			state := &resource.ProductResourceModel{
				AlertSettings: &resource.AlertSettingsModel{
					AlertSettingType: types.StringValue("override"),
				},
			}

			result, diags := JSONToTerraformProduct(ctx, tt.input, state)

			// Verify no errors occurred
			testutil.AssertDiagsHasNoErrors(t, diags)

			// Run test-specific assertions
			tt.expected(t, result)
		})
	}
}

func TestJSONToTerraformProductEmptyContext(t *testing.T) {
	input := &models.ProductJSON{
		ID:               123,
		DivisionID:       456,
		Name:             "Context Test Product",
		Status:           models.GenericIDNameJSON{ID: 0, Name: "active"},
		AlertGroup:       models.AlertGroupJSON{},
		AdvancedSettings: models.AdvancedSettingsJSON{},
		RequestSettings:  models.RequestSettingsJSON{},
		ScheduleSettings: models.ScheduleSettingsJSON{},
		InsightData:      models.InsightDataJSON{},
	}

	// This represents the user's plan, it's required for nested objects to be populated.
	state := &resource.ProductResourceModel{
		AlertSettings: &resource.AlertSettingsModel{
			AlertSettingType: types.StringValue("override"),
		},
	}

	// Test with empty context (should still work)
	result, diags := JSONToTerraformProduct(context.Background(), input, state)

	// Verify no errors occurred
	testutil.AssertDiagsHasNoErrors(t, diags)

	testutil.AssertEqual(t, "ID", result.ID.ValueInt64(), int64(123))
	testutil.AssertEqual(t, "ProductName", result.ProductName.ValueString(), "Context Test Product")
}

// TestTransformProductFields focuses on the basic field transformation
func TestTransformProductFields(t *testing.T) {
	tests := []struct {
		name     string
		input    *models.ProductJSON
		expected func(t *testing.T, result *resource.ProductResourceModel)
	}{
		{
			name: "all fields populated",
			input: &models.ProductJSON{
				ID:                123,
				DivisionID:        456,
				Name:              "Test Product",
				Status:            models.GenericIDNameJSON{ID: 0, Name: "active"},
				AlertGroupID:      789,
				TestDataWebhookID: 101112,
			},
			expected: func(t *testing.T, result *resource.ProductResourceModel) {
				t.Helper()

				testutil.AssertEqual(t, "ID", result.ID.ValueInt64(), int64(123))
				testutil.AssertEqual(t, "DivisionID", result.DivisionID.ValueInt64(), int64(456))
				testutil.AssertEqual(t, "ProductName", result.ProductName.ValueString(), "Test Product")
				testutil.AssertEqual(t, "Status", result.Status.ValueString(), cptypes.Active)
				testutil.AssertEqual(t, "AlertGroupID", result.AlertGroupID.ValueInt64(), int64(789))
				testutil.AssertEqual(t, "TestDataWebhookID", result.TestDataWebhookID.ValueInt64(), int64(101112))
			},
		},
		{
			name: "optional fields zero",
			input: &models.ProductJSON{
				ID:                123,
				DivisionID:        456,
				Name:              "Test Product",
				Status:            models.GenericIDNameJSON{ID: 0, Name: "active"},
				AlertGroupID:      0,
				TestDataWebhookID: 0,
			},
			expected: func(t *testing.T, result *resource.ProductResourceModel) {
				t.Helper()

				testutil.AssertEqual(t, "ID", result.ID.ValueInt64(), int64(123))
				testutil.AssertEqual(t, "DivisionID", result.DivisionID.ValueInt64(), int64(456))
				testutil.AssertEqual(t, "ProductName", result.ProductName.ValueString(), "Test Product")
				testutil.AssertTrue(t, "AlertGroupID", result.AlertGroupID.IsNull())
				testutil.AssertTrue(t, "TestDataWebhookID", result.TestDataWebhookID.IsNull())
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var result resource.ProductResourceModel

			diags := transformProductFields(&result, tt.input)

			// Verify no errors occurred
			testutil.AssertDiagsHasNoErrors(t, diags)
			tt.expected(t, &result)
		})
	}
}

// TestTransformNestedProductFields tests the nested object transformations
func TestTransformNestedProductFields(t *testing.T) {
	input := &models.ProductJSON{
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
	var result resource.ProductResourceModel
	// This represents the user's plan, it's required for nested objects to be populated.
	plan := &resource.ProductResourceModel{
		AlertSettings: &resource.AlertSettingsModel{
			AlertSettingType: types.StringValue("override"),
		},
	}

	diags := transformNestedProductFields(ctx, &result, input, plan)

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
	input := &models.ProductJSON{
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
	var result resource.ProductResourceModel
	// This represents the user's plan, it's required for nested objects to be populated.
	plan := &resource.ProductResourceModel{
		AlertSettings: &resource.AlertSettingsModel{
			AlertSettingType: types.StringValue("override"),
		},
	}

	diags := transformNestedProductFields(ctx, &result, input, plan)

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
