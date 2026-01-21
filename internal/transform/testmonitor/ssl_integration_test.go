package testmonitor

import (
	"context"
	"testing"

	"catchpoint-provider/internal/fields"
	"catchpoint-provider/internal/models"
	testmonitor "catchpoint-provider/internal/models/resource/testmonitor"
	"catchpoint-provider/internal/testutil"
)

func TestJSONToTerraformSSLTest(t *testing.T) {
	// Create test data for SSL test
	testJSON := &models.TestJSON{
		ID:                           456,
		DivisionID:                   789,
		ProductID:                    012,
		Name:                         "Test SSL Monitor",
		AlertsPaused:                 false,
		EnableTestDataWebhook:        true,
		StartTime:                    "2025-01-01T12:00:00Z",
		EndTime:                      testutil.ToStringPtr("2027-01-01T12:00:00Z"),
		URL:                          testutil.ToStringPtr("ssl.example2.com"),
		EnforceCertificateKeyPinning: testutil.ToBoolPtr(true),
		EnforceCertificatePinning:    testutil.ToBoolPtr(true),
		FileData:                     testutil.ToStringPtr("filedata"),
		PassPhrase:                   testutil.ToStringPtr("passphrase"),
		Status: models.GenericIDNameJSON{
			ID:   0,
			Name: "active",
		},
		Monitor: models.GenericIDNameJSON{
			ID:   31,
			Name: "ssl",
		},
	}

	sslTestModel := testmonitor.SSLTestResourceModel{}

	// Execute the function
	diags := JSONToTerraformTest(context.TODO(), &sslTestModel, testJSON, &sslTestModel)

	// Verify no errors occurred
	testutil.AssertDiagsHasNoErrors(t, diags)

	// Verify main test fields
	testutil.AssertEqual(t, fields.ID, sslTestModel.ID.ValueInt64(), int64(456))
	testutil.AssertEqual(t, fields.DivisionID, sslTestModel.DivisionID.ValueInt64(), int64(789))
	testutil.AssertEqual(t, fields.ProductID, sslTestModel.ProductID.ValueInt64(), int64(012))
	testutil.AssertEqual(t, fields.TestName, sslTestModel.Name.ValueString(), "Test SSL Monitor")
	testutil.AssertEqual(t, fields.AlertsPaused, sslTestModel.AlertsPaused.ValueBool(), false)
	testutil.AssertEqual(t, fields.EnableTestDataWebhook, sslTestModel.EnableTestDataWebhook.ValueBool(), true)
	testutil.AssertEqual(t, fields.Monitor, sslTestModel.Monitor.ValueString(), "ssl")

	// Verify optional fields are not set (since they weren't in JSON)
	testutil.AssertEqual(t, fields.FolderID, sslTestModel.FolderID.IsNull(), true)
	testutil.AssertEqual(t, fields.TestDescription, sslTestModel.Description.IsNull(), true)

	// Sometimes-optional EndTime field.
	testutil.AssertEqual(t, fields.EndTime, sslTestModel.EndTime.ValueString(), "2027-01-01T12:00:00Z")

	// Verify SSL-specific fields
	testutil.AssertEqual(t, "EnforceCertificateKeyPinning", sslTestModel.EnforceCertificateKeyPinning.ValueBool(), true)
	testutil.AssertEqual(t, "EnforceCertificatePinning", sslTestModel.EnforceCertificatePinning.ValueBool(), true)
	testutil.AssertEqual(t, "FileData", sslTestModel.FileData.ValueString(), "filedata")
	testutil.AssertEqual(t, "Passphrase", sslTestModel.PassPhrase.ValueString(), "passphrase")
	testutil.AssertEqual(t, fields.TestLocation, sslTestModel.TestLocation.ValueString(), "ssl.example2.com")
}

func TestJSONToTerraformSSLTestNoSSLSettings(t *testing.T) {
	// Create test data for SSL test
	testJSON := &models.TestJSON{
		ID:                    456,
		DivisionID:            789,
		ProductID:             012,
		Name:                  "Test SSL Monitor2",
		AlertsPaused:          false,
		EnableTestDataWebhook: true,
		StartTime:             "2025-01-01T12:00:00Z",
		EndTime:               testutil.ToStringPtr("2029-01-01T12:00:00Z"),
		URL:                   testutil.ToStringPtr("ssl.example.com"),
		Status: models.GenericIDNameJSON{
			ID:   0,
			Name: "active",
		},
		Monitor: models.GenericIDNameJSON{
			ID:   31,
			Name: "ssl",
		},
	}

	sslTestModel := testmonitor.SSLTestResourceModel{}

	// Execute the function
	// Note: since the sslTestModel is null, no complex blocks will be set.
	diags := JSONToTerraformTest(context.TODO(), &sslTestModel, testJSON, &sslTestModel)

	// Verify no errors occurred
	testutil.AssertDiagsHasNoErrors(t, diags)

	// Verify main test fields
	testutil.AssertEqual(t, fields.ID, sslTestModel.ID.ValueInt64(), int64(456))
	testutil.AssertEqual(t, fields.DivisionID, sslTestModel.DivisionID.ValueInt64(), int64(789))
	testutil.AssertEqual(t, fields.ProductID, sslTestModel.ProductID.ValueInt64(), int64(012))
	testutil.AssertEqual(t, fields.TestName, sslTestModel.Name.ValueString(), "Test SSL Monitor2")
	testutil.AssertEqual(t, fields.AlertsPaused, sslTestModel.AlertsPaused.ValueBool(), false)
	testutil.AssertEqual(t, fields.EnableTestDataWebhook, sslTestModel.EnableTestDataWebhook.ValueBool(), true)
	testutil.AssertEqual(t, fields.Monitor, sslTestModel.Monitor.ValueString(), "ssl")

	// Verify optional fields are not set (since they weren't in JSON)
	testutil.AssertEqual(t, fields.FolderID, sslTestModel.FolderID.IsNull(), true)
	testutil.AssertEqual(t, fields.TestDescription, sslTestModel.Description.IsNull(), true)

	// Sometimes-optional EndTime field.
	testutil.AssertEqual(t, fields.EndTime, sslTestModel.EndTime.ValueString(), "2029-01-01T12:00:00Z")

	testutil.AssertEqual(t, fields.TestLocation, sslTestModel.TestLocation.ValueString(), "ssl.example.com")

	// Verify SSL-specific fields
	testutil.AssertEqual(t, "EnforceCertificateKeyPinning", sslTestModel.EnforceCertificateKeyPinning.ValueBool(), false)
	testutil.AssertEqual(t, "EnforceCertificatePinning", sslTestModel.EnforceCertificatePinning.ValueBool(), false)
	testutil.AssertEqual(t, "FileData", sslTestModel.FileData.IsNull(), true)
	testutil.AssertEqual(t, "Passphrase", sslTestModel.PassPhrase.IsNull(), true)
}
