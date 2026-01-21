package testmonitor

import (
	"context"

	cptypes "catchpoint-provider/internal/types"

	"github.com/hashicorp/terraform-plugin-framework/resource"
)

type testConfigValidator struct {
	testType cptypes.TestType
}

func NewTestConfigValidator(testType cptypes.TestType) resource.ConfigValidator {
	return testConfigValidator{testType: testType}
}

func (v testConfigValidator) Description(ctx context.Context) string {
	return "Validates test configurations"
}

func (v testConfigValidator) MarkdownDescription(ctx context.Context) string {
	return "Validates test configurations"
}

func (v testConfigValidator) ValidateResource(ctx context.Context, req resource.ValidateConfigRequest, resp *resource.ValidateConfigResponse) {
	switch v.testType {
	case cptypes.APIType:
		ValidateAPITestResource(ctx, req, resp)
	case cptypes.BGPType:
		ValidateBGPTestResource(ctx, req, resp)
	case cptypes.DNSType:
		ValidateDNSTestResource(ctx, req, resp)
	case cptypes.SSLType:
		ValidateSSLTestResource(ctx, req, resp)
	case cptypes.PingType:
		ValidatePingTestResource(ctx, req, resp)
	case cptypes.PlaywrightType:
		ValidatePlaywrightTestResource(ctx, req, resp)
	case cptypes.PuppeteerType:
		ValidatePuppeteerTestResource(ctx, req, resp)
	case cptypes.TracerouteType:
		ValidateTracerouteTestResource(ctx, req, resp)
	case cptypes.TransactionType:
		ValidateTransactionTestResource(ctx, req, resp)
	case cptypes.WebType:
		ValidateWebTestResource(ctx, req, resp)
	default:
		resp.Diagnostics.AddError(
			"Unsupported Test Type",
			"The Test type validator requested is missing from the ValidateResources function.",
		)
	}
}
