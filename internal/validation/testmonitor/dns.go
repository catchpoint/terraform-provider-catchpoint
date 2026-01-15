package testmonitor

import (
	"context"

	"catchpoint-provider/internal/models/resource/testmonitor"
	"catchpoint-provider/internal/types"

	"github.com/hashicorp/terraform-plugin-framework/resource"
)

func ValidateDNSTestResource(ctx context.Context, req resource.ValidateConfigRequest, resp *resource.ValidateConfigResponse) {
	// Get the full configuration
	var config testmonitor.DNSTestResourceModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// If monitor is "DNS Experience", dns_server must not be set.
	if config.Monitor.String() == types.DNSExperienceString {
		if !config.DNSServer.IsNull() || !config.DNSServer.IsUnknown() {
			resp.Diagnostics.AddError(
				"Invalid Configuration",
				"When 'monitor' is set to 'DNS Experience', the 'dns_server' field must not be specified.",
			)
		}
	}
	// TODO: add any custom validation here...
}
