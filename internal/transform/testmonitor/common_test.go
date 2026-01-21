package testmonitor

import (
	"context"
	"testing"

	"catchpoint-provider/internal/models"
	testmonitor "catchpoint-provider/internal/models/resource/testmonitor"
	"catchpoint-provider/internal/testutil"

	"github.com/hashicorp/terraform-plugin-framework/types"
)

func TestJSONToTerraformTestWithNilInputShouldPanic(t *testing.T) {
	// Verify that calling with nil input causes a panic
	defer func() {
		if r := recover(); r != nil {
			// Expected panic occurred
			panicMsg := r.(string)
			testutil.AssertStringContains(t, "panic message should mention nil input", panicMsg, "nil TestJSON")
		} else {
			t.Error("Expected panic when calling JSONToTerraformTest with nil input")
		}
	}()

	apiTestModel := testmonitor.APITestResourceModel{}

	// This should panic
	_ = JSONToTerraformTest(context.TODO(), &apiTestModel, nil, &apiTestModel)
}

func TestJSONToTerraformTestWithLabels(t *testing.T) {
	// Create test data with multiple labels
	testJSON := &models.TestJSON{
		ID:                    123,
		DivisionID:            456,
		ProductID:             789,
		Name:                  "Test with Labels",
		AlertsPaused:          false,
		EnableTestDataWebhook: true,
		StartTime:             "2025-01-01T12:00:00Z",
		Status: models.GenericIDNameJSON{
			ID:   0,
			Name: "active",
		},
		Monitor: models.GenericIDNameJSON{
			ID:   25,
			Name: "api",
		},
		// Test with multiple labels
		Labels: &[]models.LabelsJSON{
			{
				Name:   "team",
				Values: []string{"backend", "platform"},
			},
			{
				Name:   "environment",
				Values: []string{"production"},
			},
			{
				Name:   "region",
				Values: []string{"us-east-1", "us-west-2", "eu-central-1"},
			},
		},
	}

	apiTestModel := testmonitor.APITestResourceModel{}

	// Execute the function
	diags := JSONToTerraformTest(context.TODO(), &apiTestModel, testJSON, &apiTestModel)

	// Verify no errors occurred
	testutil.AssertDiagsHasNoErrors(t, diags)

	// Verify labels were set correctly
	testutil.AssertEqual(t, "number of labels", len(apiTestModel.Label), 3)

	// Test first label: team
	firstLabel := apiTestModel.Label[0]
	testutil.AssertEqual(t, "first label key", firstLabel.Key.ValueString(), "team")
	testutil.AssertEqual(t, "first label values count", len(firstLabel.Values.Elements()), 2)

	// Convert Values to slice of strings for easier testing
	firstLabelValues := make([]string, len(firstLabel.Values.Elements()))
	for i, v := range firstLabel.Values.Elements() {
		firstLabelValues[i] = v.(types.String).ValueString()
	}
	testutil.AssertContains(t, "first label should contain 'backend'", firstLabelValues, "backend")
	testutil.AssertContains(t, "first label should contain 'platform'", firstLabelValues, "platform")

	// Test second label: environment
	secondLabel := apiTestModel.Label[1]
	testutil.AssertEqual(t, "second label key", secondLabel.Key.ValueString(), "environment")
	testutil.AssertEqual(t, "second label values count", len(secondLabel.Values.Elements()), 1)

	secondLabelValues := make([]string, len(secondLabel.Values.Elements()))
	for i, v := range secondLabel.Values.Elements() {
		secondLabelValues[i] = v.(types.String).ValueString()
	}
	testutil.AssertContains(t, "second label should contain 'production'", secondLabelValues, "production")

	// Test third label: region
	thirdLabel := apiTestModel.Label[2]
	testutil.AssertEqual(t, "third label key", thirdLabel.Key.ValueString(), "region")
	testutil.AssertEqual(t, "third label values count", len(thirdLabel.Values.Elements()), 3)

	thirdLabelValues := make([]string, len(thirdLabel.Values.Elements()))
	for i, v := range thirdLabel.Values.Elements() {
		thirdLabelValues[i] = v.(types.String).ValueString()
	}
	testutil.AssertContains(t, "third label should contain 'us-east-1'", thirdLabelValues, "us-east-1")
	testutil.AssertContains(t, "third label should contain 'us-west-2'", thirdLabelValues, "us-west-2")
	testutil.AssertContains(t, "third label should contain 'eu-central-1'", thirdLabelValues, "eu-central-1")
}
