package testmonitor

import (
	"context"
	"testing"

	"catchpoint-provider/internal/fields"
	"catchpoint-provider/internal/helpers"
	"catchpoint-provider/internal/models"
	testmonitor "catchpoint-provider/internal/models/resource/testmonitor"
	"catchpoint-provider/internal/testutil"
)

func TestJSONToTerraformAPITest(t *testing.T) {
	script := `// Step - 1 \nname("step")\nname("step 1")\nopen("https://www.amazon.com")`

	// Create test data for API test with script
	testJSON := &models.TestJSON{
		ID:                    456,
		DivisionID:            789,
		ProductID:             012,
		Name:                  "Test API Monitor",
		AlertsPaused:          false,
		EnableTestDataWebhook: true,
		StartTime:             "2025-01-01T12:00:00Z",
		EndTime:               testutil.ToStringPtr("2027-01-01T12:00:00Z"),
		URL:                   testutil.ToStringPtr("https://api.example.com"), // This will be ignored for API tests
		Status: models.GenericIDNameJSON{
			ID:   0,
			Name: "active",
		},
		Monitor: models.GenericIDNameJSON{
			ID:   25,
			Name: "api",
		},
		TestRequestData: &models.TestRequestDataJSON{
			RequestData: &script,
			TransactionScriptType: &models.GenericIDNameOmitEmptyJSON{
				ID:   testutil.ToIntPtr(2),
				Name: testutil.ToStringPtr("javascript"),
			},
		},
		RequestSettings: &models.RequestSettingsJSON{
			HTTPHeaderRequests: &[]models.HTTPHeaderRequestJSON{
				{
					HeaderName:   testutil.ToStringPtr("accept"),
					RequestValue: "application/json",
					RequestHeaderType: models.GenericIDNameJSON{
						ID:   2,
						Name: "accept",
					},
				},
			},
		},
		AlertGroup: &models.AlertGroupJSON{
			AlertSettingType: models.GenericIDNameJSON{
				ID:   1,
				Name: "override",
			},
			NotificationGroup: &models.NotificationGroupJSON{
				NotifyOnWarning:  false,
				NotifyOnCritical: true,
				NotifyOnImproved: true,
				Subject:          "API Test Alert",
				Recipients:       []models.RecipientJSON{{Email: "api@example.com"}},
			},
			AlertGroupItems: []models.AlertGroupItemJSON{
				{
					AlertType:          models.GenericIDNameJSON{ID: 2, Name: "availability"},
					EnforceTestFailure: false,
					Trigger: models.TriggerJSON{
						TriggerType: models.GenericIDNameJSON{ID: 1, Name: "specific value"},
					},
					NodeThreshold: models.NodeThresholdJSON{
						Name:                 "all",
						NodeThresholdType:    models.GenericIDNameJSON{ID: 2, Name: "all"},
						NumberOfUnits:        testutil.ToIntPtr(1),
						NumberOfFailingUnits: testutil.ToIntPtr(1),
					},
				},
			},
		},
	}

	apiTestModel := testmonitor.APITestResourceModel{}

	// Execute the function
	diags := JSONToTerraformTest(context.TODO(), &apiTestModel, testJSON, &apiTestModel)

	// Verify no errors occurred
	testutil.AssertDiagsHasNoErrors(t, diags)

	// Verify main test fields
	testutil.AssertEqual(t, fields.ID, apiTestModel.ID.ValueInt64(), int64(456))
	testutil.AssertEqual(t, fields.DivisionID, apiTestModel.DivisionID.ValueInt64(), int64(789))
	testutil.AssertEqual(t, fields.ProductID, apiTestModel.ProductID.ValueInt64(), int64(012))
	testutil.AssertEqual(t, fields.TestName, apiTestModel.Name.ValueString(), "Test API Monitor")
	testutil.AssertEqual(t, fields.AlertsPaused, apiTestModel.AlertsPaused.ValueBool(), false)
	testutil.AssertEqual(t, fields.EnableTestDataWebhook, apiTestModel.EnableTestDataWebhook.ValueBool(), true)
	testutil.AssertEqual(t, fields.Monitor, apiTestModel.Monitor.ValueString(), "api")
	testutil.AssertEqual(t, fields.TestScript, apiTestModel.TestScriptResourceModel.Script.ValueString(), helpers.NormalizeScript(script))
	testutil.AssertEqual(t, fields.TestScriptType, apiTestModel.TestScriptResourceModel.ScriptType.ValueString(), "javascript")

	// Verify optional fields are not set (since they weren't in JSON)
	testutil.AssertEqual(t, fields.FolderID, apiTestModel.FolderID.IsNull(), true)
	testutil.AssertEqual(t, fields.TestDescription, apiTestModel.Description.IsNull(), true)

	// Sometimes-optional EndTime field.
	testutil.AssertEqual(t, fields.EndTime, apiTestModel.EndTime.ValueString(), "2027-01-01T12:00:00Z")
}
