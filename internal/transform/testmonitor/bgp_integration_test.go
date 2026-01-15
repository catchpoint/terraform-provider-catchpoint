package testmonitor

import (
	"context"
	"testing"

	"catchpoint-provider/internal/fields"
	"catchpoint-provider/internal/models"
	testmonitor "catchpoint-provider/internal/models/resource/testmonitor"
	"catchpoint-provider/internal/testutil"
)

func TestJSONToTerraformBGPTest(t *testing.T) {
	// Create test data for BGP test with script
	testJSON := &models.TestJSON{
		ID:                    456,
		DivisionID:            789,
		ProductID:             012,
		Name:                  "Test BGP Monitor",
		AlertsPaused:          false,
		EnableTestDataWebhook: true,
		StartTime:             "2025-01-01T12:00:00Z",
		EndTime:               testutil.ToStringPtr("2027-01-01T12:00:00Z"),
		URL:                   testutil.ToStringPtr("192.168.0.0/24"),
		Status: models.GenericIDNameJSON{
			ID:   0,
			Name: "active",
		},
		Monitor: models.GenericIDNameJSON{
			ID:   34,
			Name: "bgp",
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
				Subject:          "BGP Test Alert",
				Recipients:       []models.RecipientJSON{{Email: "bgp@example.com"}},
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

	bgpTestModel := testmonitor.BGPTestResourceModel{}

	// Execute the function
	diags := JSONToTerraformTest(context.TODO(), &bgpTestModel, testJSON, &bgpTestModel)

	// Verify no errors occurred
	testutil.AssertDiagsHasNoErrors(t, diags)

	// Verify main test fields
	testutil.AssertEqual(t, fields.ID, bgpTestModel.ID.ValueInt64(), int64(456))
	testutil.AssertEqual(t, fields.DivisionID, bgpTestModel.DivisionID.ValueInt64(), int64(789))
	testutil.AssertEqual(t, fields.ProductID, bgpTestModel.ProductID.ValueInt64(), int64(012))
	testutil.AssertEqual(t, fields.TestName, bgpTestModel.Name.ValueString(), "Test BGP Monitor")
	testutil.AssertEqual(t, fields.AlertsPaused, bgpTestModel.AlertsPaused.ValueBool(), false)
	testutil.AssertEqual(t, fields.EnableTestDataWebhook, bgpTestModel.EnableTestDataWebhook.ValueBool(), true)
	testutil.AssertEqual(t, fields.Monitor, bgpTestModel.Monitor.ValueString(), "bgp")

	// Verify optional fields are not set (since they weren't in JSON)
	testutil.AssertEqual(t, fields.FolderID, bgpTestModel.FolderID.IsNull(), true)
	testutil.AssertEqual(t, fields.TestDescription, bgpTestModel.Description.IsNull(), true)

	// Sometimes-optional EndTime field.
	testutil.AssertEqual(t, fields.EndTime, bgpTestModel.EndTime.ValueString(), "2027-01-01T12:00:00Z")
}
