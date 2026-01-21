package testmonitor

import (
	"context"

	"catchpoint-provider/internal/models/resource/testmonitor"
	cptypes "catchpoint-provider/internal/types"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
)

const (
	additionalFieldErrorSummary      = "Additional Field"
	missingRequiredFieldErrorSummary = "Missing Required Field"
)

func ValidateTransactionTestResource(ctx context.Context, req resource.ValidateConfigRequest, resp *resource.ValidateConfigResponse) {
	// Get the full configuration
	var config testmonitor.TransactionTestResourceModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}

	switch config.Monitor.ValueString() {
	case cptypes.ChromeString:
		validateChromeMonitor(config, resp)
	case cptypes.MobileString:
		validateMobileMonitor(config, resp)
	case cptypes.EmulatedString:
		validateEmulatedMonitor(config, resp)
	}
}

func validateChromeMonitor(config testmonitor.TransactionTestResourceModel, resp *resource.ValidateConfigResponse) {
	if config.ChromeVersion.IsNull() {
		resp.Diagnostics.AddAttributeError(
			path.Root("chrome_version"),
			missingRequiredFieldErrorSummary,
			"Transaction tests with a 'chrome' monitor type require a chrome_version field",
		)
	}
	if !config.Simulate.IsNull() {
		resp.Diagnostics.AddAttributeError(
			path.Root("simulate"),
			additionalFieldErrorSummary,
			"Transaction tests with a 'chrome' monitor type should not provide a simulate field",
		)
	}
}

func validateMobileMonitor(config testmonitor.TransactionTestResourceModel, resp *resource.ValidateConfigResponse) {
	if config.Simulate.IsNull() {
		resp.Diagnostics.AddAttributeError(
			path.Root("simulate"),
			missingRequiredFieldErrorSummary,
			"Transaction tests with a 'mobile' monitor type require a simulate field",
		)
	}
	if !config.ChromeVersion.IsNull() {
		resp.Diagnostics.AddAttributeError(
			path.Root("chrome_version"),
			additionalFieldErrorSummary,
			"Transaction tests with a 'mobile' monitor type should not provide a chrome_version field",
		)
	}
}

func validateEmulatedMonitor(config testmonitor.TransactionTestResourceModel, resp *resource.ValidateConfigResponse) {
	if !config.ChromeVersion.IsNull() {
		resp.Diagnostics.AddAttributeError(
			path.Root("chrome_version"),
			additionalFieldErrorSummary,
			"Transaction tests with an 'emulated' monitor type should not provide a chrome_version field",
		)
	}
	if !config.Simulate.IsNull() {
		resp.Diagnostics.AddAttributeError(
			path.Root("simulate"),
			additionalFieldErrorSummary,
			"Transaction tests with an 'emulated' monitor type should not provide a simulate field",
		)
	}
}
