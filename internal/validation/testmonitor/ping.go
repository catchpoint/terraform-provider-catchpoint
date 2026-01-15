package testmonitor

import (
	"context"

	"catchpoint-provider/internal/models/resource/testmonitor"

	"github.com/hashicorp/terraform-plugin-framework/resource"
)

func ValidatePingTestResource(ctx context.Context, req resource.ValidateConfigRequest, resp *resource.ValidateConfigResponse) {
	// Get the full configuration
	var config testmonitor.PingTestResourceModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// TODO: add any custom validation here...
}
