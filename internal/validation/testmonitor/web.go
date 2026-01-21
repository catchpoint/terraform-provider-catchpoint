package testmonitor

import (
	"context"

	"catchpoint-provider/internal/models/resource/testmonitor"
	cptypes "catchpoint-provider/internal/types"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
)

func ValidateWebTestResource(ctx context.Context, req resource.ValidateConfigRequest, resp *resource.ValidateConfigResponse) {
	// Get the full configuration
	var config testmonitor.WebTestResourceModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}

	switch config.Monitor.ValueString() {
	case cptypes.ChromeString:
		validateWebChromeMonitor(config, resp)
	case cptypes.MobileString:
	case cptypes.MobilePlaybackString:
	case cptypes.PlaybackString:
		validateWebMobileMonitor(config, resp)
	case cptypes.EmulatedString:
	case cptypes.HTTPString:
		validateWebEmulatedMonitor(config, resp)
	}

	/* TODO: Validate that "simulate" uses only the correct values.
	The "simulate" field is actually userAgentTypeId which presents differently based on monitor type.

	For Playback, it shows as "Playback Source" in the UI and has values: Chrome, IE.
	For Mobile, it shows as "Simulate" and has many values: Android, BlackBerryZ30, etc...
	For Mobile Playback, it shows as "Playback Source" and has a smaller selection of the Mobile values: Android, iPhone, etc...
	For Emulated, HTTP, and Chrome, it's not available and should not be set.*/

	// TODO: Validate RequestType field should only be used with HTTP, Emulated, and Chrome.
	// "Request Type" of "GET" or "POST" are valid fields for HTTP, Emulated, and Chrome monitor types.
	// However, this field was not available in the previous version of the provider, so we will add it later.

	// TODO: add any custom validation here...
}

func validateWebChromeMonitor(config testmonitor.WebTestResourceModel, resp *resource.ValidateConfigResponse) {
	if config.ChromeVersion.IsNull() {
		resp.Diagnostics.AddAttributeError(
			path.Root("chrome_version"),
			missingRequiredFieldErrorSummary,
			"Web tests with a 'chrome' monitor type require a chrome_version field",
		)
	}
	if !config.Simulate.IsNull() {
		resp.Diagnostics.AddAttributeError(
			path.Root("simulate"),
			additionalFieldErrorSummary,
			"Web tests with a 'chrome' monitor type should not provide a simulate field",
		)
	}
}

func validateWebMobileMonitor(config testmonitor.WebTestResourceModel, resp *resource.ValidateConfigResponse) {
	if config.Simulate.IsNull() {
		resp.Diagnostics.AddAttributeError(
			path.Root("simulate"),
			missingRequiredFieldErrorSummary,
			"Web tests with a 'mobile', 'playback', or 'mobile playback' monitor type require a simulate field",
		)
	}
	if !config.ChromeVersion.IsNull() {
		resp.Diagnostics.AddAttributeError(
			path.Root("chrome_version"),
			additionalFieldErrorSummary,
			"Web tests with a 'mobile', 'playback', or 'mobile playback' monitor type should not provide a chrome_version field",
		)
	}
}

func validateWebEmulatedMonitor(config testmonitor.WebTestResourceModel, resp *resource.ValidateConfigResponse) {
	if !config.ChromeVersion.IsNull() {
		resp.Diagnostics.AddAttributeError(
			path.Root("chrome_version"),
			additionalFieldErrorSummary,
			"Web tests with an 'emulated' or 'http' monitor type should not provide a chrome_version field",
		)
	}
	if !config.Simulate.IsNull() {
		resp.Diagnostics.AddAttributeError(
			path.Root("simulate"),
			additionalFieldErrorSummary,
			"Web tests with an 'emulated' or 'http' monitor type should not provide a simulate field",
		)
	}
}
