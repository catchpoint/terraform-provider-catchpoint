package testmonitor

import (
	"context"

	"catchpoint-provider/internal/models/resource/testmonitor"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
)

func ValidateAPITestResource(ctx context.Context, req resource.ValidateConfigRequest, resp *resource.ValidateConfigResponse) {
	// Get the full configuration
	var config testmonitor.APITestResourceModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Validate API test requirements
	if config.Script.IsNull() {
		resp.Diagnostics.AddAttributeError(
			path.Root("test_script"),
			"Missing Required Field",
			"API tests require a test_script field",
		)
	}

	if config.ScriptType.IsNull() {
		resp.Diagnostics.AddAttributeError(
			path.Root("test_script_type"),
			"Missing Required Field",
			"API tests require a test_script_type field",
		)
	}

}
