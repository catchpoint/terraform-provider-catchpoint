package testmonitor

import (
	"fmt"

	"catchpoint-provider/internal/expand"
	"catchpoint-provider/internal/helpers"
	"catchpoint-provider/internal/models"
	"catchpoint-provider/internal/models/resource"
	cptypes "catchpoint-provider/internal/types"
	"catchpoint-provider/internal/validation"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// ExpandTestConfig is the generic expand function. It orchestrates the expansion
// of a Terraform plan of any type T into a TestConfig API model.
func ExpandTestConfig[T resource.TestResourceModelProvider](plan T, config *models.TestConfig) (diags diag.Diagnostics) {
	// Expand common fields.
	diags.Append(expandTestConfigFromTestResourceModel(*plan.GetTestResourceModel(), config, plan.GetTestType())...)
	if diags.HasError() {
		return
	}

	// Expand complex nested objects by checking for interface implementation.
	diags.Append(expandComplexBlocks(plan, config)...)
	if diags.HasError() {
		return
	}

	// Expand script data if the model supports it.
	if provider, ok := any(plan).(resource.TestScriptResourceProvider); ok {
		diags.Append(expandTestRequestData(*provider.GetTestScriptResourceModel(), config)...)
	}

	// Check if the model implements the URLProvider interface.
	if urlProvider, ok := any(plan).(resource.URLProvider); ok {
		// If it does, get the URL field.
		if urlValue, hasURL := urlProvider.GetURLField(); hasURL && !urlValue.IsNull() {
			config.TestURL = urlValue.ValueString()
		}
	}

	// Handle GatewayAddressOrHost if present.
	if gatewayProvider, ok := any(plan).(resource.GatewayAddressOrHostModelProvider); ok {
		gatewayModel := gatewayProvider.GetGatewayAddressOrHostModel()
		if !gatewayModel.GatewayAddressOrHost.IsNull() {
			config.GatewayAddressOrHost = gatewayModel.GatewayAddressOrHost.ValueString()
		}
	}

	diags.Append(expandDNSTestConfig(plan, config)...)
	if diags.HasError() {
		return
	}

	diags.Append(expandThresholdConfig(plan, config)...)
	if diags.HasError() {
		return
	}

	diags.Append(expandSSLTestConfig(plan, config)...)
	if diags.HasError() {
		return
	}

	diags.Append(expandChromeVersionConfig(plan, config)...)
	if diags.HasError() {
		return
	}

	diags.Append(expandUserAgentTestConfig(plan, config)...)
	if diags.HasError() {
		return
	}

	return
}

func expandDNSTestConfig[T resource.TestResourceModelProvider](plan T, config *models.TestConfig) (diags diag.Diagnostics) {
	if dnsProvider, ok := any(plan).(resource.DNSTestModelProvider); ok {
		dnsServer, ok := dnsProvider.GetDNSServerField()
		if ok && !dnsServer.IsNull() {
			config.DNSServer = dnsServer.ValueString()
		}
		queryType, ok := dnsProvider.GetQueryTypeField()
		if ok && !queryType.IsNull() {
			// This field should always be set to something valid due to schema validation.
			config.DNSQueryType = validation.GetDNSQueryTypeOrDefault(queryType.ValueString())
		}
	}
	return
}

func expandSSLTestConfig[T resource.TestResourceModelProvider](plan T, config *models.TestConfig) (diags diag.Diagnostics) {
	if sslProvider, ok := any(plan).(resource.SSLSettingsModelProvider); ok {
		sslModel := sslProvider.GetSSLSettingsModel()
		if !sslModel.EnforceCertificatePinning.IsNull() {
			config.EnforceCertificatePinning = sslModel.EnforceCertificatePinning.ValueBool()
		}
		if !sslModel.EnforceCertificateKeyPinning.IsNull() {
			config.EnforceCertificateKeyPinning = sslModel.EnforceCertificateKeyPinning.ValueBool()
		}
		if !sslModel.FileData.IsNull() {
			config.FileData = sslModel.FileData.ValueString()
		}
		if !sslModel.PassPhrase.IsNull() {
			config.Passphrase = sslModel.PassPhrase.ValueString()
		}
	}
	return
}

func expandThresholdConfig[T resource.TestResourceModelProvider](plan T, config *models.TestConfig) (diags diag.Diagnostics) {
	if thresholdProvider, ok := any(plan).(resource.ThresholdModelProvider); ok {
		thresholdModel := thresholdProvider.GetTestThresholdModel()
		if thresholdModel != nil {
			if !thresholdModel.AvailabilityCritical.IsNull() {
				config.TestThresholds.AvailabilityThresholdCritical = thresholdModel.AvailabilityCritical.ValueFloat64()
			}
			if !thresholdModel.AvailabilityWarning.IsNull() {
				config.TestThresholds.AvailabilityThresholdWarning = thresholdModel.AvailabilityWarning.ValueFloat64()
			}
			if !thresholdModel.TestTimeCritical.IsNull() {
				config.TestThresholds.TestTimeThresholdCritical = thresholdModel.TestTimeCritical.ValueFloat64()
			}
			if !thresholdModel.TestTimeWarning.IsNull() {
				config.TestThresholds.TestTimeThresholdWarning = thresholdModel.TestTimeWarning.ValueFloat64()
			}
		}
	}
	return
}

// expandTestConfigFromTestResourceModel expands the Terraform plan into a TestConfig model.
func expandTestConfigFromTestResourceModel(plan resource.BaseTestResourceModel, config *models.TestConfig, testType cptypes.TestType) (diags diag.Diagnostics) {
	// Set Monitor and TestType first as they're used later on.
	monitor, ok := cptypes.GetMonitorTypeID(plan.Monitor.ValueString())
	if !ok {
		diags.AddError("Invalid Monitor Type",
			fmt.Sprintf("Monitor type %s not supported.", plan.Monitor.ValueString()),
		)
		return
	}

	config.Monitor.ID = monitor
	config.Monitor.Name = plan.Monitor.ValueString()

	testTypeName := cptypes.GetTestTypeName(testType)
	config.TestType.ID = int(testType)
	config.TestType.Name = testTypeName

	// Set the fields that are required by the schema.
	requiredDiags := expandRequiredFields(plan, config)
	diags.Append(requiredDiags...)
	if diags.HasError() {
		return
	}

	// Set the optional fields that may be nil or have defaults.
	optionalDiags := expandOptionalFields(plan, config)
	diags.Append(optionalDiags...)
	if diags.HasError() {
		return
	}

	return
}

func expandRequiredFields(plan resource.BaseTestResourceModel, config *models.TestConfig) (diags diag.Diagnostics) {
	// These fields are required by the schema and must be set or the user receives an error.
	config.CommonConfig.DivisionID = int(plan.DivisionID.ValueInt64())
	config.ProductID = int(plan.ProductID.ValueInt64())
	config.TestName = plan.Name.ValueString()

	return
}

func expandOptionalFields(plan resource.BaseTestResourceModel, config *models.TestConfig) (diags diag.Diagnostics) {
	// These 2 have no default, so only set if provided. These 2 fields are omitted by API if empty.
	if !plan.FolderID.IsNull() {
		config.FolderID = int(plan.FolderID.ValueInt64())
	}

	if !plan.Description.IsNull() {
		config.TestDescription = plan.Description.ValueString()
	}

	// Defaults to "active" in the backend and by schema.
	config.Status = expand.GetStatusFromStringOrActive(plan.Status)

	// Defaults to false via schema if not set by user.
	if !plan.AlertsPaused.IsNull() {
		config.AlertsPaused = plan.AlertsPaused.ValueBool()
	}

	// Defaults to the current time via the schema if not set by user.
	if !plan.StartTime.IsNull() && plan.StartTime.ValueString() != cptypes.EmptyString {
		config.StartTime = plan.StartTime.ValueString()
	} else {
		currentTime := helpers.GetTime()
		config.StartTime = currentTime
	}

	// Defaults to empty string via schema if not set by user.
	// Note: this field is odd and is only required for *some* clients.
	if !plan.EndTime.IsNull() {
		config.EndTime = plan.EndTime.ValueString()
	}

	// All tests support labels, so we can always expand them here.
	diags.Append(expandLabels(plan, config)...)
	if diags.HasError() {
		return
	}

	// Defaults to true via schema if not set by user.
	if !plan.EnableTestDataWebhook.IsNull() {
		config.EnableTestDataWebhook = plan.EnableTestDataWebhook.ValueBool()
	}

	return
}

func expandLabels(plan resource.BaseTestResourceModel, config *models.TestConfig) (diags diag.Diagnostics) {
	if len(plan.Label) > 0 {
		// Initialize the Labels slice if not already done
		config.Labels = make([]models.TestLabel, 0, len(plan.Label))

		for i, labelModel := range plan.Label {
			// Validate that each label has both key and values set
			if labelModel.Key.IsNull() || labelModel.Values.IsNull() {
				diags.AddError("Invalid Label Model",
					fmt.Sprintf("Label at index %d must have both key and values set.", i),
				)
				return
			}

			// Convert the values list to a string slice
			labelValues := make([]string, 0, len(labelModel.Values.Elements()))
			for _, value := range labelModel.Values.Elements() {
				if stringValue, ok := value.(types.String); ok {
					labelValues = append(labelValues, stringValue.ValueString())
				} else {
					diags.AddError("Invalid Label Value Type",
						fmt.Sprintf("Label at index %d contains a non-string value.", i),
					)
					return
				}
			}

			// Add the label to the config
			config.Labels = append(config.Labels, models.TestLabel{
				Name:   labelModel.Key.ValueString(),
				Values: labelValues,
			})
		}
	}

	return
}

func expandTestRequestData(testScript resource.TestScriptResourceModel, config *models.TestConfig) (diags diag.Diagnostics) {
	if !testScript.Script.IsNull() {
		if !testScript.Script.IsNull() {
			config.Script.RequestData = helpers.NormalizeScript(testScript.Script.ValueString()) // Set the script as request data.
		} else {
			// Theoretically we shouldn't get here. But this is possible if someone introduces a schema that
			// doesn't require some form of location or script.
			diags.AddError("Test Script Required",
				"Test script must be provided for API, Playwright, Puppeteer, and Transaction tests.",
			)
			return
		}

		if !testScript.ScriptType.IsNull() {
			scriptType, ok := cptypes.GetAPIScriptTypeID(testScript.ScriptType.ValueString())
			// This shouldn't happen due to schema validation, but we check just in case.
			if !ok {
				diags.AddError("Invalid Script Type",
					fmt.Sprintf("Test script type %s not supported.", testScript.ScriptType.ValueString()),
				)
				return
			}

			config.Script.TransactionScriptType.ID = scriptType
			config.Script.TransactionScriptType.Name = testScript.ScriptType.ValueString()
		}

		// For whatever reason, the testRequestData object needs copies of the parent monitor and test type.
		// We set them again here to ensure they're in the payload when this is later sent to the backend.
		config.Script.Monitor.ID = config.Monitor.ID
		config.Script.Monitor.Name = config.Monitor.Name
		config.Script.TestType.ID = config.TestType.ID
		config.Script.TestType.Name = config.TestType.Name
	}

	return
}

func expandUserAgentTestConfig[T resource.TestResourceModelProvider](plan T, config *models.TestConfig) (diags diag.Diagnostics) {
	if userAgentProvider, ok := any(plan).(resource.SimulateProvider); ok {
		if userAgentModel, ok := userAgentProvider.GetSimulateField(); ok && !userAgentModel.IsNull() {
			simulateDevice, ok := cptypes.GetUserAgentTypeID(userAgentModel.ValueString())
			if !ok {
				diags.AddError("Invalid User Agent",
					fmt.Sprintf("User agent %s not supported.", userAgentModel.ValueString()),
				)
				return
			}
			config.SimulateDevice.ID = simulateDevice
			config.SimulateDevice.Name = userAgentModel.ValueString()
		}
	}
	return
}

func expandChromeVersionConfig[T resource.TestResourceModelProvider](plan T, config *models.TestConfig) (diags diag.Diagnostics) {
	if chromeProvider, ok := any(plan).(resource.ChromeVersionProvider); ok {
		if chromeModel, ok := chromeProvider.GetChromeVersionField(); ok && !chromeModel.IsNull() {
			diags.Append(setChromeVersionFields(chromeModel.ValueString(), config)...)
			if diags.HasError() {
				return
			}
		}
	}
	return
}

func setChromeVersionFields(chromeVersionStr string, config *models.TestConfig) (diags diag.Diagnostics) {
	chromeVersion, ok := cptypes.GetChromeVersionID(chromeVersionStr)
	if !ok {
		diags.AddError("Invalid Chrome Version",
			fmt.Sprintf("Chrome version %s not supported.", chromeVersionStr),
		)
		return
	}

	config.ChromeApplicationVersion.ApplicationVersionType.ID = chromeVersion
	config.ChromeApplicationVersion.ApplicationVersionType.Name = chromeVersionStr

	if chromeVersion == 3 {
		if chromeApplicationVersionID, ok := cptypes.GetChromeApplicationVersionID(chromeVersionStr); ok {
			config.ChromeApplicationVersion.ApplicationVersionID = chromeApplicationVersionID
		}
	}
	return
}

// expandComplexBlocks expands all supported nested objects (Alerts, Advanced, etc.)
// by checking if the plan model implements the corresponding provider interface.
func expandComplexBlocks[T resource.TestResourceModelProvider](plan T, config *models.TestConfig) (diags diag.Diagnostics) {
	// Expand Alert Settings
	if provider, ok := any(plan).(resource.AlertSettingsModelProvider); ok {
		diags.Append(expand.ExpandAlertSettingsConfig(provider.GetAlertSettingsModel(), &config.CommonConfig.AlertSettingsConfig)...)
		if diags.HasError() {
			return
		}
	}

	// Expand Advanced Settings
	if provider, ok := any(plan).(resource.AdvancedSettingsModelProvider); ok {
		model := provider.GetAdvancedSettingsModel()
		diags.Append(expand.ExpandAdvancedSettingsConfig(model.AdvancedSettings, &config.CommonConfig.AdvancedSettingsConfig)...)
		if diags.HasError() {
			return
		}
	}

	// Expand Request Settings
	if provider, ok := any(plan).(resource.RequestSettingsModelProvider); ok {
		model := provider.GetRequestSettingsModel()
		diags.Append(expand.ExpandRequestSettingsConfig(model.RequestSettings, &config.CommonConfig.RequestSettingsConfig)...)
		if diags.HasError() {
			return
		}
	}

	// Expand Insight Settings
	if provider, ok := any(plan).(resource.InsightSettingsModelProvider); ok {
		model := provider.GetInsightSettingsModel()
		diags.Append(expand.ExpandInsightSettingsConfig(model.Insights, &config.CommonConfig.InsightSettingsConfig)...)
		if diags.HasError() {
			return
		}
	}

	// Expand Schedule Settings
	if provider, ok := any(plan).(resource.ScheduleSettingsModelProvider); ok {
		model := provider.GetScheduleSettingsModel()
		diags.Append(expand.ExpandScheduleSettingsConfig(model.ScheduleSettings, &config.CommonConfig.ScheduleSettingsConfig)...)
		if diags.HasError() {
			return
		}
	}

	return
}
