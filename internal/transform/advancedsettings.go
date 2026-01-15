package transform

import (
	"fmt"

	"catchpoint-provider/internal/fields"
	"catchpoint-provider/internal/logger"
	"catchpoint-provider/internal/models"
	"catchpoint-provider/internal/models/resource"
	cpschema "catchpoint-provider/internal/schema"
	cptypes "catchpoint-provider/internal/types"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// JSONToTerraformAdvancedSettingsForTest converts AdvancedSettingsJSON to a Terraform object for tests.
func JSONToTerraformAdvancedSettingsForTest(advancedSettings *models.AdvancedSettingsJSON, testType cptypes.TestType) (resource resource.AdvancedSettingsModel, diags diag.Diagnostics) {
	attrTypes := cpschema.GetAdvancedSettingsAttributeTypesForTest(testType)
	resource.AdvancedSettings, diags = jsonToTerraformAdvancedSettings(advancedSettings, attrTypes)
	return
}

// JSONToTerraformAdvancedSettingsForProductAndFolder converts AdvancedSettingsJSON to a Terraform object for products and folders.
func JSONToTerraformAdvancedSettingsForProductAndFolder(advancedSettings *models.AdvancedSettingsJSON) (resource resource.AdvancedSettingsModel, diags diag.Diagnostics) {
	attrTypes := cpschema.GetAdvancedSettingsAttributeTypesForProductAndFolder()
	resource.AdvancedSettings, diags = jsonToTerraformAdvancedSettings(advancedSettings, attrTypes)
	return
}

// jsonToTerraformAdvancedSettings is a shared helper for converting AdvancedSettingsJSON to a Terraform object.
func jsonToTerraformAdvancedSettings(advancedSettings *models.AdvancedSettingsJSON, attrTypes map[string]attr.Type) (obj types.Object, diags diag.Diagnostics) {
	obj = types.ObjectNull(attrTypes)

	if advancedSettings == nil {
		logger.WarnBG("AdvancedSettings is nil, returning null object")
		return
	}

	advancedSettingTypeName, ok := cptypes.GetGenericSettingTypeName(advancedSettings.AdvancedSettingType.ID)
	if !ok {
		logger.WarnBG("Invalid AdvancedSettingType ID: %d", advancedSettings.AdvancedSettingType.ID)
		diags.Append(diag.NewWarningDiagnostic("Invalid AdvancedSettingType ID", "The provided AdvancedSettingType ID is not recognized."))
	}

	// Build the nested attributes map.
	attrs, buildDiags := buildAdvancedSettingsAttrs(advancedSettings, attrTypes)
	if buildDiags.HasError() {
		diags.Append(buildDiags...)
		return
	}

	attrs[fields.AdvancedSettingType] = types.StringValue(advancedSettingTypeName)

	// Create the object directly.
	obj, objDiags := types.ObjectValue(attrTypes, attrs)
	diags.Append(objDiags...)
	return
}

func buildAdvancedSettingsAttrs(advancedSettings *models.AdvancedSettingsJSON, attrTypes map[string]attr.Type) (attrs map[string]attr.Value, diags diag.Diagnostics) {
	attrs = make(map[string]attr.Value)

	handleEDNSSubnet(attrs, attrTypes, advancedSettings)

	additionalMonitorDiags := handleAdditionalMonitor(attrs, attrTypes, advancedSettings)
	diags.Append(additionalMonitorDiags...)

	SetIntOrNull(attrs, attrTypes, fields.FailureHopCount, advancedSettings.FailureHopCount)
	SetIntOrNull(attrs, attrTypes, fields.EnforceTestFailureIfRunsLongerThan, advancedSettings.MaxStepRuntimeSecOverride)
	SetIntOrNull(attrs, attrTypes, fields.PingCount, advancedSettings.PingCount)
	SetIntOrNull(attrs, attrTypes, fields.ViewportHeight, advancedSettings.ViewportHeight)
	SetIntOrNull(attrs, attrTypes, fields.ViewportWidth, advancedSettings.ViewportWidth)
	SetIntOrNull(attrs, attrTypes, fields.WaitForNoActivity, advancedSettings.WaitForNoActivity)

	handleTestFlagsDiag := handleTestFlags(attrs, attrTypes, advancedSettings)
	diags.Append(handleTestFlagsDiag...)

	handleBandwidthThrottlingDiags := handleBandwidthThrottling(attrs, attrTypes, advancedSettings)
	diags.Append(handleBandwidthThrottlingDiags...)

	SetNullValuesForMissingFields(attrs, attrTypes)

	return
}

func handleEDNSSubnet(attrs map[string]attr.Value, attrTypes map[string]attr.Type, advancedSettings *models.AdvancedSettingsJSON) {
	if _, exists := attrTypes[fields.EDNSSubnet]; exists {
		if advancedSettings.EDNSSubnet != nil {
			attrs[fields.EDNSSubnet] = types.StringValue(*advancedSettings.EDNSSubnet)
		} else {
			attrs[fields.EDNSSubnet] = types.StringNull()
		}
	}
}

func handleAdditionalMonitor(attrs map[string]attr.Value, attrTypes map[string]attr.Type, advancedSettings *models.AdvancedSettingsJSON) (diags diag.Diagnostics) {
	if _, exists := attrTypes[fields.AdditionalMonitor]; exists {
		if advancedSettings.AdditionalMonitor != nil {
			additionalMonitorName, ok := cptypes.GetAdditionalMonitorTypeName(*advancedSettings.AdditionalMonitor.ID)
			if ok {
				attrs[fields.AdditionalMonitor] = types.StringValue(additionalMonitorName)
			} else {
				logger.WarnBG("Unknown additional monitor type %d", *advancedSettings.AdditionalMonitor.ID)
				diags.Append(diag.NewWarningDiagnostic("Invalid AdditionalMonitor ID", "The provided AdditionalMonitor ID is not recognized."))
				attrs[fields.AdditionalMonitor] = types.StringNull()
			}
		} else {
			attrs[fields.AdditionalMonitor] = types.StringNull()
		}
	}
	return
}

func handleBandwidthThrottling(attrs map[string]attr.Value, attrTypes map[string]attr.Type, advancedSettings *models.AdvancedSettingsJSON) (diags diag.Diagnostics) {
	if _, exists := attrTypes[fields.BandwidthThrottling]; exists {
		if advancedSettings.TestBandwidthThrottling != nil {
			bandwidthThrottlingName, ok := cptypes.GetBandwidthThrottlingTypeName(*advancedSettings.TestBandwidthThrottling.ID)
			if ok {
				attrs[fields.BandwidthThrottling] = types.StringValue(bandwidthThrottlingName)
			} else {
				logger.WarnBG("Unknown bandwidth throttling type %d", *advancedSettings.TestBandwidthThrottling.ID)
				diags.Append(diag.NewWarningDiagnostic("Invalid BandwidthThrottling ID", fmt.Sprintf("The provided BandwidthThrottling ID '%d' is not recognized in the schema.", *advancedSettings.TestBandwidthThrottling.ID)))
				attrs[fields.BandwidthThrottling] = types.StringNull()
			}
		} else {
			attrs[fields.BandwidthThrottling] = types.StringNull()
		}
	}

	return
}

func handleTestFlags(attrs map[string]attr.Value, attrTypes map[string]attr.Type, advancedSettings *models.AdvancedSettingsJSON) (diags diag.Diagnostics) {
	// For every test flag received from the backend, check if it exists in the attribute types.
	for _, testFlag := range advancedSettings.AppliedTestFlags {
		// The name from the backend does not match the name in our schema so do a lookup.
		testFlagName, ok := cptypes.GetTestFlagName(*testFlag.ID)

		// If what we got from the backend is known in our types, attempt to set it in the schema.
		if ok {
			if _, exists := attrTypes[testFlagName]; exists {
				// If we got it from the backend, and it exists in our schema, set it to true.
				attrs[testFlagName] = types.BoolValue(true)
			} else {
				// If we got it from the backend but it's not in the schema, the schema needs to be updated.
				// This is a warning because it means the backend has a test flag that we don't know about.
				// It could be a new test flag that hasn't been added to the provider yet.
				logger.WarnBG("Test flag '%s' ID '%d' is not available in the provider schema.", testFlagName, *testFlag.ID)
				diags.Append(diag.NewWarningDiagnostic("Invalid TestFlag ID", fmt.Sprintf("The provided TestFlag ID '%d' is not recognized in the schema.", *testFlag.ID)))
			}
		} else {
			// If we got it from the backend but it's not in our types, the types need to be updated.
			// This is a warning because it means the backend has a test flag that we don't know about.
			// It could be a new test flag that hasn't been added to the provider yet.
			logger.WarnBG("Test flag ID '%d' not known to provider types. Will not map.", *testFlag.ID)
			diags.Append(diag.NewWarningDiagnostic("Invalid TestFlag ID", fmt.Sprintf("The provided TestFlag ID '%d' is not recognized in the type mapping.", *testFlag.ID)))
		}
	}

	return
}
