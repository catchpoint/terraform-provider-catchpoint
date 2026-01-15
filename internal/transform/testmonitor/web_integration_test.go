package testmonitor

import (
	"context"
	"testing"

	"catchpoint-provider/internal/fields"
	"catchpoint-provider/internal/models"
	testmonitor "catchpoint-provider/internal/models/resource/testmonitor"
	"catchpoint-provider/internal/testutil"
)

func TestJSONToTerraformWebTest(t *testing.T) {
	test := []struct {
		monitor   string
		monitorID int
	}{
		{"http", 2},
		{"emulated", 3},
		{"chrome", 18},
		{"mobile", 26},
		{"playback", 19},
		{"mobile playback", 20},
	}
	for _, tt := range test {
		// Create test data for Web test with script
		testJSON := &models.TestJSON{
			ID:                    456,
			DivisionID:            789,
			ProductID:             012,
			Name:                  "Test Web Monitor",
			AlertsPaused:          false,
			EnableTestDataWebhook: true,
			StartTime:             "2025-01-01T12:00:00Z",
			EndTime:               testutil.ToStringPtr("2027-01-01T12:00:00Z"),
			URL:                   testutil.ToStringPtr("https://test.example.com"),
			Status: models.GenericIDNameJSON{
				ID:   0,
				Name: "active",
			},
			Monitor: models.GenericIDNameJSON{
				ID:   tt.monitorID,
				Name: tt.monitor,
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
			GatewayAddressOrHost: testutil.ToStringPtr("192.168.0.11"),
			AlertGroup: &models.AlertGroupJSON{
				AlertSettingType: models.GenericIDNameJSON{
					ID:   1,
					Name: "override",
				},
				NotificationGroup: &models.NotificationGroupJSON{
					NotifyOnWarning:  false,
					NotifyOnCritical: true,
					NotifyOnImproved: true,
					Subject:          "Web Test Alert",
					Recipients:       []models.RecipientJSON{{Email: "Web@example.com"}},
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

		WebTestModel := testmonitor.WebTestResourceModel{}

		// Execute the function
		diags := JSONToTerraformTest(context.TODO(), &WebTestModel, testJSON, &WebTestModel)

		// Verify no errors occurred
		testutil.AssertDiagsHasNoErrors(t, diags)

		// Verify main test fields
		testutil.AssertEqual(t, fields.ID, WebTestModel.ID.ValueInt64(), int64(456))
		testutil.AssertEqual(t, fields.DivisionID, WebTestModel.DivisionID.ValueInt64(), int64(789))
		testutil.AssertEqual(t, fields.ProductID, WebTestModel.ProductID.ValueInt64(), int64(012))
		testutil.AssertEqual(t, fields.TestName, WebTestModel.Name.ValueString(), "Test Web Monitor")
		testutil.AssertEqual(t, fields.AlertsPaused, WebTestModel.AlertsPaused.ValueBool(), false)
		testutil.AssertEqual(t, fields.EnableTestDataWebhook, WebTestModel.EnableTestDataWebhook.ValueBool(), true)
		testutil.AssertEqual(t, fields.Monitor, WebTestModel.Monitor.ValueString(), tt.monitor)

		// Verify optional fields are not set (since they weren't in JSON)
		testutil.AssertEqual(t, fields.FolderID, WebTestModel.FolderID.IsNull(), true)
		testutil.AssertEqual(t, fields.TestDescription, WebTestModel.Description.IsNull(), true)

		// Sometimes-optional EndTime field.
		testutil.AssertEqual(t, fields.EndTime, WebTestModel.EndTime.ValueString(), "2027-01-01T12:00:00Z")

		// Verify GatewayAddressOrHost field.
		testutil.AssertEqual(t, fields.GatewayAddressOrHost, WebTestModel.GatewayAddressOrHost.ValueString(), "192.168.0.11")
	}
}

func TestJSONToTerraformWebTestChromeField(t *testing.T) {
	test := []struct {
		chromeVersion              string
		chromeVersionID            int
		chromeApplicationVersionID *int
	}{
		{"stable", 1, nil},
		{"preview", 2, nil},
		{"53", 3, testutil.ToIntPtr(1)},
		{"59", 3, testutil.ToIntPtr(3)},
		{"63", 3, testutil.ToIntPtr(4)},
		{"66", 3, testutil.ToIntPtr(5)},
		{"71", 3, testutil.ToIntPtr(8)},
		{"75", 3, testutil.ToIntPtr(7)},
		{"85", 3, testutil.ToIntPtr(12)},
		{"87", 3, testutil.ToIntPtr(13)},
		{"89", 3, testutil.ToIntPtr(14)},
		{"97", 3, testutil.ToIntPtr(28357)},
		{"108", 3, testutil.ToIntPtr(28558)},
		{"120", 3, testutil.ToIntPtr(31965)},
	}
	for _, tt := range test {
		// Create test data for Web test with script
		testJSON := &models.TestJSON{
			ID:                    456,
			DivisionID:            789,
			ProductID:             012,
			Name:                  "Test Web Monitor",
			AlertsPaused:          false,
			EnableTestDataWebhook: true,
			StartTime:             "2025-01-01T12:00:00Z",
			EndTime:               testutil.ToStringPtr("2027-01-01T12:00:00Z"),
			URL:                   testutil.ToStringPtr("https://test.example.com"),
			ChromeMonitorVersion: &models.ChromeMonitorVersionStructJSON{
				ApplicationVersionType: &models.GenericIDNameOmitEmptyJSON{
					ID:   testutil.ToIntPtr(tt.chromeVersionID),
					Name: testutil.ToStringPtr(tt.chromeVersion),
				},
				ApplicationVersionID: tt.chromeApplicationVersionID,
			},
			Monitor: models.GenericIDNameJSON{
				ID:   18,
				Name: "chrome",
			},
		}

		WebTestModel := testmonitor.WebTestResourceModel{}

		// Execute the function
		diags := JSONToTerraformTest(context.TODO(), &WebTestModel, testJSON, &WebTestModel)

		// Verify no errors occurred
		testutil.AssertDiagsHasNoErrors(t, diags)

		// Verify ChromeVersion field.
		testutil.AssertEqual(t, fields.ChromeVersion, WebTestModel.ChromeVersion.ValueString(), tt.chromeVersion)

	}
}

func TestJSONToTerraformWebTestSimulateField(t *testing.T) {
	test := []struct {
		userAgent   string
		userAgentID int
	}{
		{"ie", 1},
		{"chrome", 2},
		{"android", 3},
		{"iphone", 4},
		{"ipad 2", 5},
		{"kindle fire", 6},
		{"galaxy tab", 7},
		{"iphone 5", 8},
		{"ipad mini", 9},
		{"galaxy note", 10},
		{"nexus 7", 11},
		{"nexus 4", 12},
		{"nokia lumia920", 13},
		{"iphone 6", 14},
		{"blackberry z30", 15},
		{"galaxy s4", 16},
		{"htc onex", 17},
		{"lg optimusg", 18},
		{"droid razr hd", 19},
		{"nexus 6", 20},
		{"iphone 6s", 21},
		{"galaxy s6", 22},
		{"iphone 7", 23},
		{"google pixel", 24},
		{"galaxy s8", 25},
	}
	for _, tt := range test {
		// Create test data for Web test with script
		testJSON := &models.TestJSON{
			ID:                    456,
			DivisionID:            789,
			ProductID:             012,
			Name:                  "Test Web Monitor",
			AlertsPaused:          false,
			EnableTestDataWebhook: true,
			StartTime:             "2025-01-01T12:00:00Z",
			EndTime:               testutil.ToStringPtr("2027-01-01T12:00:00Z"),
			URL:                   testutil.ToStringPtr("https://test.example.com"),
			Monitor: models.GenericIDNameJSON{
				ID:   26,
				Name: "mobile",
			},
			UserAgentType: &models.GenericIDNameOmitEmptyJSON{
				ID:   testutil.ToIntPtr(tt.userAgentID),
				Name: testutil.ToStringPtr(tt.userAgent),
			},
		}

		WebTestModel := testmonitor.WebTestResourceModel{}

		// Execute the function
		diags := JSONToTerraformTest(context.TODO(), &WebTestModel, testJSON, &WebTestModel)

		// Verify no errors occurred
		testutil.AssertDiagsHasNoErrors(t, diags)

		// Verify Simulate field.
		testutil.AssertEqual(t, fields.Simulate, WebTestModel.Simulate.ValueString(), tt.userAgent)
	}
}
