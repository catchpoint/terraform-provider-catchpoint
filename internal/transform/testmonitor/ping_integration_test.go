package testmonitor

import (
	"context"
	"testing"

	"catchpoint-provider/internal/fields"
	"catchpoint-provider/internal/models"
	testmonitor "catchpoint-provider/internal/models/resource/testmonitor"
	"catchpoint-provider/internal/testutil"
	"catchpoint-provider/internal/types"
)

func TestJSONToTerraformPingICMPTest(t *testing.T) {
	test := []struct {
		monitor   string
		monitorID int
	}{
		{types.PingICMPString, 8},
		{types.PingTCPString, 11},
		{types.PingUDPString, 23},
	}
	for _, tt := range test {
		// Create test data for Ping test
		testJSON := &models.TestJSON{
			ID:                    456,
			DivisionID:            789,
			ProductID:             012,
			Name:                  "Test Ping Monitor",
			AlertsPaused:          false,
			EnableTestDataWebhook: true,
			StartTime:             "2025-01-01T12:00:00Z",
			EndTime:               testutil.ToStringPtr("2028-01-01T12:00:00Z"),
			URL:                   testutil.ToStringPtr("www.example.com"),
			Status: models.GenericIDNameJSON{
				ID:   0,
				Name: "active",
			},
			Monitor: models.GenericIDNameJSON{
				ID:   tt.monitorID,
				Name: tt.monitor,
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
					Subject:          "Ping Test Alert",
					Recipients:       []models.RecipientJSON{{Email: "someone@example.com"}},
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

		PingTestModel := testmonitor.PingTestResourceModel{}

		// Execute the function
		diags := JSONToTerraformTest(context.TODO(), &PingTestModel, testJSON, &PingTestModel)

		// Verify no errors occurred
		testutil.AssertDiagsHasNoErrors(t, diags)

		// Verify main test fields
		testutil.AssertEqual(t, fields.ID, PingTestModel.ID.ValueInt64(), int64(456))
		testutil.AssertEqual(t, fields.DivisionID, PingTestModel.DivisionID.ValueInt64(), int64(789))
		testutil.AssertEqual(t, fields.ProductID, PingTestModel.ProductID.ValueInt64(), int64(012))
		testutil.AssertEqual(t, fields.TestName, PingTestModel.Name.ValueString(), "Test Ping Monitor")
		testutil.AssertEqual(t, fields.AlertsPaused, PingTestModel.AlertsPaused.ValueBool(), false)
		testutil.AssertEqual(t, fields.EnableTestDataWebhook, PingTestModel.EnableTestDataWebhook.ValueBool(), true)
		testutil.AssertEqual(t, fields.Monitor, PingTestModel.Monitor.ValueString(), tt.monitor)

		// Verify optional fields are not set (since they weren't in JSON)
		testutil.AssertEqual(t, fields.FolderID, PingTestModel.FolderID.IsNull(), true)
		testutil.AssertEqual(t, fields.TestDescription, PingTestModel.Description.IsNull(), true)

		// Sometimes-optional EndTime field.
		testutil.AssertEqual(t, fields.EndTime, PingTestModel.EndTime.ValueString(), "2028-01-01T12:00:00Z")
	}
}
