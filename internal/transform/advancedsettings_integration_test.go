package transform

import (
	"testing"

	"catchpoint-provider/internal/fields"
	"catchpoint-provider/internal/models"
	"catchpoint-provider/internal/testutil"
	cptypes "catchpoint-provider/internal/types"

	"github.com/hashicorp/terraform-plugin-framework/types"
)

func TestJSONToTerraformAdvancedSettings(t *testing.T) {
	// Create test data with common advanced settings
	advancedSettingsJSON := &models.AdvancedSettingsJSON{
		AdvancedSettingType: models.GenericIDNameJSON{
			ID:   1,
			Name: "override",
		},
		ViewportHeight:    testutil.ToIntPtr(1080),
		ViewportWidth:     testutil.ToIntPtr(1920),
		WaitForNoActivity: testutil.ToIntPtr(5000),
		EDNSSubnet:        testutil.ToStringPtr("192.168.1.0/24"),
		AdditionalMonitor: &models.GenericIDNameOmitEmptyJSON{
			ID:   testutil.ToIntPtr(8),
			Name: testutil.ToStringPtr("ping icmp"),
		},
		TestBandwidthThrottling: &models.GenericIDNameOmitEmptyJSON{
			ID:   testutil.ToIntPtr(2),
			Name: testutil.ToStringPtr("3g"),
		},
		AppliedTestFlags: []models.GenericIDNameOmitEmptyJSON{
			{ID: testutil.ToIntPtr(2), Name: testutil.ToStringPtr("flag1")},
			{ID: testutil.ToIntPtr(8), Name: testutil.ToStringPtr("flag2")},
		},
	}

	// Execute the function
	result, diags := JSONToTerraformAdvancedSettingsForProductAndFolder(advancedSettingsJSON)

	// Verify no errors occurred
	testutil.AssertDiagsHasNoErrors(t, diags)

	// Verify the result is not null
	testutil.AssertEqual(t, "result should not be null", result.AdvancedSettings.IsNull(), false)

	// Get the outer attributes
	outerAttrs := result.AdvancedSettings.Attributes()

	// Verify advanced setting type
	if advancedSettingType, exists := outerAttrs[fields.AdvancedSettingType]; exists {
		advancedSettingTypeStr := advancedSettingType.(types.String)
		testutil.AssertEqual(t, "advanced setting type", advancedSettingTypeStr.ValueString(), "override")
	}

	// Verify viewport dimensions
	if viewportHeight, exists := outerAttrs[fields.ViewportHeight]; exists {
		viewportHeightInt := viewportHeight.(types.Int64)
		testutil.AssertEqual(t, "viewport height", viewportHeightInt.ValueInt64(), int64(1080))
	}

	if viewportWidth, exists := outerAttrs[fields.ViewportWidth]; exists {
		viewportWidthInt := viewportWidth.(types.Int64)
		testutil.AssertEqual(t, "viewport width", viewportWidthInt.ValueInt64(), int64(1920))
	}

	// Verify wait for no activity
	if waitForNoActivity, exists := outerAttrs[fields.WaitForNoActivity]; exists {
		waitForNoActivityInt := waitForNoActivity.(types.Int64)
		testutil.AssertEqual(t, "wait for no activity", waitForNoActivityInt.ValueInt64(), int64(5000))
	}

	// Verify EDNS subnet
	if ednsSubnet, exists := outerAttrs[fields.EDNSSubnet]; exists {
		ednsSubnetStr := ednsSubnet.(types.String)
		testutil.AssertEqual(t, "edns subnet", ednsSubnetStr.ValueString(), "192.168.1.0/24")
	}

	// Verify additional monitor
	if additionalMonitor, exists := outerAttrs[fields.AdditionalMonitor]; exists {
		additionalMonitorStr := additionalMonitor.(types.String)
		testutil.AssertEqual(t, "additional monitor should not be null", additionalMonitorStr.IsNull(), false)
	}

	// Verify bandwidth throttling
	if bandwidthThrottling, exists := outerAttrs[fields.BandwidthThrottling]; exists {
		bandwidthThrottlingStr := bandwidthThrottling.(types.String)
		testutil.AssertEqual(t, "bandwidth throttling should not be null", bandwidthThrottlingStr.IsNull(), false)
	}

	// Verify that we have attributes (structure should be non-empty)
	testutil.AssertGreaterThan(t, "should have multiple nested attributes", len(outerAttrs), 3)
}

func TestJSONToTerraformAdvancedSettingsWithNilInput(t *testing.T) {
	testType := cptypes.WebType

	// Execute the function with nil input
	result, diags := JSONToTerraformAdvancedSettingsForTest(nil, testType)

	// Verify no errors occurred
	testutil.AssertDiagsHasNoErrors(t, diags)

	// Verify the result is null (expected behavior for nil input)
	testutil.AssertEqual(t, "result should be null for nil input", result.AdvancedSettings.IsNull(), true)
}

func TestJSONToTerraformAdvancedSettingsMinimalData(t *testing.T) {
	testType := cptypes.PingType

	// Create minimal test data (only required fields)
	advancedSettingsJSON := &models.AdvancedSettingsJSON{
		AdvancedSettingType: models.GenericIDNameJSON{
			ID:   0,
			Name: "inherit",
		},
	}

	// Execute the function
	result, diags := JSONToTerraformAdvancedSettingsForTest(advancedSettingsJSON, testType)

	// Verify no errors occurred
	testutil.AssertDiagsHasNoErrors(t, diags)

	// Verify the result is not null
	testutil.AssertNotNil(t, "result should not be null", result)

	// Get the advanced settings
	outerAttrs := result.AdvancedSettings.Attributes()

	// Verify advanced setting type is set
	if advancedSettingType, exists := outerAttrs[fields.AdvancedSettingType]; exists {
		advancedSettingTypeStr := advancedSettingType.(types.String)
		testutil.AssertEqual(t, "advanced setting type", advancedSettingTypeStr.ValueString(), "inherit")
	}

	// Verify structure exists even with minimal data
	testutil.AssertGreaterThan(t, "should have at least one attribute", len(outerAttrs), 0)
}

func TestJSONToTerraformAdvancedSettingsDifferentTestTypes(t *testing.T) {
	// Test different test types to ensure schema switching works
	// Note: BGP does not get advanced settings.
	testTypes := []cptypes.TestType{
		cptypes.APIType,
		cptypes.DNSType,
		cptypes.PingType,
		cptypes.PlaywrightType,
		cptypes.PuppeteerType,
		cptypes.SSLType,
		cptypes.TracerouteType,
		cptypes.WebType,
	}

	advancedSettingsJSON := &models.AdvancedSettingsJSON{
		AdvancedSettingType: models.GenericIDNameJSON{
			ID:   1,
			Name: "override",
		},
		AppliedTestFlags: []models.GenericIDNameOmitEmptyJSON{},
	}

	for _, testType := range testTypes {
		testName := cptypes.GetTestTypeName(testType)
		t.Run("TestType_"+testName, func(t *testing.T) {
			// Execute the function
			result, diags := JSONToTerraformAdvancedSettingsForTest(advancedSettingsJSON, testType)

			// Verify no errors occurred
			testutil.AssertDiagsHasNoErrors(t, diags)

			// Verify the result is not null
			testutil.AssertNotNil(t, "result should not be null for "+testName, result)

			// Verify structure exists
			outerAttrs := result.AdvancedSettings.Attributes()
			testutil.AssertGreaterThan(t, "should have AdvancedSettings fields", len(outerAttrs), 0)

			_, exists := outerAttrs[fields.AdvancedSettingType]
			testutil.AssertEqual(t, "should contain AdvancedSettingType field", exists, true)
		})
	}
}
