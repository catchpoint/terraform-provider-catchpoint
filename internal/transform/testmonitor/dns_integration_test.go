package testmonitor

import (
	"context"
	"testing"

	"catchpoint-provider/internal/fields"
	"catchpoint-provider/internal/models"
	testmonitor "catchpoint-provider/internal/models/resource/testmonitor"
	"catchpoint-provider/internal/testutil"
)

func TestJSONToTerraformDNSDirectTest(t *testing.T) {
	// Create test data for DNS Direct test
	testJSON := &models.TestJSON{
		ID:                    456,
		DivisionID:            789,
		ProductID:             012,
		Name:                  "Test DNS Direct Monitor",
		AlertsPaused:          false,
		EnableTestDataWebhook: true,
		StartTime:             "2025-01-01T12:00:00Z",
		EndTime:               testutil.ToStringPtr("2028-01-01T12:00:00Z"),
		URL:                   testutil.ToStringPtr("dns.example.com"),
		DNSQueryType: &models.GenericIDNameOmitEmptyJSON{
			ID:   testutil.ToIntPtr(1),
			Name: testutil.ToStringPtr("a"),
		},
		DNSServer: testutil.ToStringPtr("8.8.8.8"),
		Status: models.GenericIDNameJSON{
			ID:   0,
			Name: "active",
		},
		Monitor: models.GenericIDNameJSON{
			ID:   13,
			Name: "dns direct",
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
				Subject:          "DNS Test Alert",
				Recipients:       []models.RecipientJSON{{Email: "dns@example.com"}},
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

	dnsTestModel := testmonitor.DNSTestResourceModel{}

	// Execute the function
	diags := JSONToTerraformTest(context.TODO(), &dnsTestModel, testJSON, &dnsTestModel)

	// Verify no errors occurred
	testutil.AssertDiagsHasNoErrors(t, diags)

	// Verify main test fields
	testutil.AssertEqual(t, fields.ID, dnsTestModel.ID.ValueInt64(), int64(456))
	testutil.AssertEqual(t, fields.DivisionID, dnsTestModel.DivisionID.ValueInt64(), int64(789))
	testutil.AssertEqual(t, fields.ProductID, dnsTestModel.ProductID.ValueInt64(), int64(012))
	testutil.AssertEqual(t, fields.TestName, dnsTestModel.Name.ValueString(), "Test DNS Direct Monitor")
	testutil.AssertEqual(t, fields.AlertsPaused, dnsTestModel.AlertsPaused.ValueBool(), false)
	testutil.AssertEqual(t, fields.EnableTestDataWebhook, dnsTestModel.EnableTestDataWebhook.ValueBool(), true)
	testutil.AssertEqual(t, fields.Monitor, dnsTestModel.Monitor.ValueString(), "dns direct")

	// Verify optional fields are not set (since they weren't in JSON)
	testutil.AssertEqual(t, fields.FolderID, dnsTestModel.FolderID.IsNull(), true)
	testutil.AssertEqual(t, fields.TestDescription, dnsTestModel.Description.IsNull(), true)

	// Sometimes-optional EndTime field.
	testutil.AssertEqual(t, fields.EndTime, dnsTestModel.EndTime.ValueString(), "2028-01-01T12:00:00Z")

	// Verify DNS-specific fields
	testutil.AssertEqual(t, "DNS Server", dnsTestModel.DNSServer.ValueString(), "8.8.8.8")
	testutil.AssertEqual(t, "Query Type", dnsTestModel.QueryType.ValueString(), "a")
}

func TestJSONToTerraformDNSExperienceTest(t *testing.T) {
	// Create test data for DNS test with script
	testJSON := &models.TestJSON{
		ID:                    456,
		DivisionID:            789,
		ProductID:             012,
		Name:                  "Test DNS Monitor",
		AlertsPaused:          false,
		EnableTestDataWebhook: true,
		StartTime:             "2025-01-01T12:00:00Z",
		EndTime:               testutil.ToStringPtr("2027-01-01T12:00:00Z"),
		URL:                   testutil.ToStringPtr("dns.example.com"),
		DNSQueryType: &models.GenericIDNameOmitEmptyJSON{
			ID:   testutil.ToIntPtr(5),
			Name: testutil.ToStringPtr("cname"),
		},
		DNSServer: testutil.ToStringPtr("1.1.1.1"),
		Status: models.GenericIDNameJSON{
			ID:   0,
			Name: "active",
		},
		Monitor: models.GenericIDNameJSON{
			ID:   12,
			Name: "dns experience",
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
				Subject:          "DNS Test Alert",
				Recipients:       []models.RecipientJSON{{Email: "dns@example.com"}},
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

	dnsTestModel := testmonitor.DNSTestResourceModel{}

	// Execute the function
	diags := JSONToTerraformTest(context.TODO(), &dnsTestModel, testJSON, &dnsTestModel)

	// Verify no errors occurred
	testutil.AssertDiagsHasNoErrors(t, diags)

	// Verify main test fields
	testutil.AssertEqual(t, fields.ID, dnsTestModel.ID.ValueInt64(), int64(456))
	testutil.AssertEqual(t, fields.DivisionID, dnsTestModel.DivisionID.ValueInt64(), int64(789))
	testutil.AssertEqual(t, fields.ProductID, dnsTestModel.ProductID.ValueInt64(), int64(012))
	testutil.AssertEqual(t, fields.TestName, dnsTestModel.Name.ValueString(), "Test DNS Monitor")
	testutil.AssertEqual(t, fields.AlertsPaused, dnsTestModel.AlertsPaused.ValueBool(), false)
	testutil.AssertEqual(t, fields.EnableTestDataWebhook, dnsTestModel.EnableTestDataWebhook.ValueBool(), true)
	testutil.AssertEqual(t, fields.Monitor, dnsTestModel.Monitor.ValueString(), "dns experience")

	// Verify optional fields are not set (since they weren't in JSON)
	testutil.AssertEqual(t, fields.FolderID, dnsTestModel.FolderID.IsNull(), true)
	testutil.AssertEqual(t, fields.TestDescription, dnsTestModel.Description.IsNull(), true)

	// Sometimes-optional EndTime field.
	testutil.AssertEqual(t, fields.EndTime, dnsTestModel.EndTime.ValueString(), "2027-01-01T12:00:00Z")

	// Verify DNS-specific fields
	testutil.AssertEqual(t, "DNS Server", dnsTestModel.DNSServer.ValueString(), "1.1.1.1")
	testutil.AssertEqual(t, "Query Type", dnsTestModel.QueryType.ValueString(), "cname")
}
