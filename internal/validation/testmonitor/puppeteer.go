package testmonitor

import (
	"context"

	"catchpoint-provider/internal/models/resource/testmonitor"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
)

func ValidatePuppeteerTestResource(ctx context.Context, req resource.ValidateConfigRequest, resp *resource.ValidateConfigResponse) {
	// Get the full configuration
	var config testmonitor.PuppeteerTestResourceModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Validate Puppeteer test requirements
	if config.Script.IsNull() {
		resp.Diagnostics.AddAttributeError(
			path.Root("test_script"),
			"Missing Required Field",
			"Puppeteer tests require a test_script field",
		)
	}

}
