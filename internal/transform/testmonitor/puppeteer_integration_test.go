package testmonitor

import (
	"context"
	"testing"

	"catchpoint-provider/internal/fields"
	"catchpoint-provider/internal/models"
	testmonitor "catchpoint-provider/internal/models/resource/testmonitor"
	"catchpoint-provider/internal/testutil"
)

func TestJSONToTerraformPuppeteerTest(t *testing.T) {
	script := `// Step - 1 \nname("step")\nname("step 1")\nopen("https://www.amazon.com")`

	// Create test data for Puppeteer test with script
	testJSON := &models.TestJSON{
		ID:                    456,
		DivisionID:            789,
		ProductID:             012,
		Name:                  "Test Puppeteer Monitor",
		AlertsPaused:          false,
		EnableTestDataWebhook: true,
		StartTime:             "2025-01-01T12:00:00Z",
		EndTime:               testutil.ToStringPtr("2027-01-01T12:00:00Z"),
		URL:                   testutil.ToStringPtr("https://something.example.com"), // This will be ignored for Puppeteer tests
		Status: models.GenericIDNameJSON{
			ID:   0,
			Name: "active",
		},
		Monitor: models.GenericIDNameJSON{
			ID:   18,
			Name: "chrome",
		},
		TestRequestData: &models.TestRequestDataJSON{
			RequestData: &script,
			TransactionScriptType: &models.GenericIDNameOmitEmptyJSON{
				ID:   testutil.ToIntPtr(4),
				Name: testutil.ToStringPtr("puppeteer"),
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
				Subject:          "Puppeteer Test Alert",
				Recipients:       []models.RecipientJSON{{Email: "Puppeteer@example.com"}},
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

	PuppeteerTestModel := testmonitor.PuppeteerTestResourceModel{}

	// Execute the function
	diags := JSONToTerraformTest(context.TODO(), &PuppeteerTestModel, testJSON, &PuppeteerTestModel)

	// Verify no errors occurred
	testutil.AssertDiagsHasNoErrors(t, diags)

	// Verify main test fields
	testutil.AssertEqual(t, fields.ID, PuppeteerTestModel.ID.ValueInt64(), int64(456))
	testutil.AssertEqual(t, fields.DivisionID, PuppeteerTestModel.DivisionID.ValueInt64(), int64(789))
	testutil.AssertEqual(t, fields.ProductID, PuppeteerTestModel.ProductID.ValueInt64(), int64(012))
	testutil.AssertEqual(t, fields.TestName, PuppeteerTestModel.Name.ValueString(), "Test Puppeteer Monitor")
	testutil.AssertEqual(t, fields.AlertsPaused, PuppeteerTestModel.AlertsPaused.ValueBool(), false)
	testutil.AssertEqual(t, fields.EnableTestDataWebhook, PuppeteerTestModel.EnableTestDataWebhook.ValueBool(), true)
	testutil.AssertEqual(t, fields.Monitor, PuppeteerTestModel.Monitor.ValueString(), "chrome")

	// Verify optional fields are not set (since they weren't in JSON)
	testutil.AssertEqual(t, fields.FolderID, PuppeteerTestModel.FolderID.IsNull(), true)
	testutil.AssertEqual(t, fields.TestDescription, PuppeteerTestModel.Description.IsNull(), true)

	// Sometimes-optional EndTime field.
	testutil.AssertEqual(t, fields.EndTime, PuppeteerTestModel.EndTime.ValueString(), "2027-01-01T12:00:00Z")
}
