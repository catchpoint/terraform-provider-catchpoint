package testmonitor

import (
	"context"
	"testing"

	"catchpoint-provider/internal/fields"
	"catchpoint-provider/internal/models"
	testmonitor "catchpoint-provider/internal/models/resource/testmonitor"
	"catchpoint-provider/internal/testutil"
)

func TestJSONToTerraformTracerouteICMPTest(t *testing.T) {
	test := []struct {
		monitor   string
		monitorID int
	}{
		{"traceroute icmp", 9},
		{"traceroute tcp", 29},
		{"traceroute udp", 14},
	}
	for _, tt := range test {
		// Create test data for Traceroute test
		testJSON := &models.TestJSON{
			ID:                    456,
			DivisionID:            789,
			ProductID:             012,
			Name:                  "Test Traceroute Monitor",
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
					Subject:          "Traceroute Test Alert",
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

		TracerouteTestModel := testmonitor.TracerouteTestResourceModel{}

		// Execute the function
		diags := JSONToTerraformTest(context.TODO(), &TracerouteTestModel, testJSON, &TracerouteTestModel)

		// Verify no errors occurred
		testutil.AssertDiagsHasNoErrors(t, diags)

		// Verify main test fields
		testutil.AssertEqual(t, fields.ID, TracerouteTestModel.ID.ValueInt64(), int64(456))
		testutil.AssertEqual(t, fields.DivisionID, TracerouteTestModel.DivisionID.ValueInt64(), int64(789))
		testutil.AssertEqual(t, fields.ProductID, TracerouteTestModel.ProductID.ValueInt64(), int64(012))
		testutil.AssertEqual(t, fields.TestName, TracerouteTestModel.Name.ValueString(), "Test Traceroute Monitor")
		testutil.AssertEqual(t, fields.AlertsPaused, TracerouteTestModel.AlertsPaused.ValueBool(), false)
		testutil.AssertEqual(t, fields.EnableTestDataWebhook, TracerouteTestModel.EnableTestDataWebhook.ValueBool(), true)
		testutil.AssertEqual(t, fields.Monitor, TracerouteTestModel.Monitor.ValueString(), tt.monitor)

		// Verify optional fields are not set (since they weren't in JSON)
		testutil.AssertEqual(t, fields.FolderID, TracerouteTestModel.FolderID.IsNull(), true)
		testutil.AssertEqual(t, fields.TestDescription, TracerouteTestModel.Description.IsNull(), true)

		// Sometimes-optional EndTime field.
		testutil.AssertEqual(t, fields.EndTime, TracerouteTestModel.EndTime.ValueString(), "2028-01-01T12:00:00Z")
	}
}
