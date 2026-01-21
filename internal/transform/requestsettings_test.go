package transform

import (
	"testing"

	"catchpoint-provider/internal/models"
	"catchpoint-provider/internal/testutil"

	"github.com/hashicorp/terraform-plugin-framework/types"
)

func TestTransformRequestSettingsNil(t *testing.T) {
	obj, diags := JSONToTerraformRequestSettings(nil)

	testutil.AssertDiagsHasNoErrors(t, diags)
	testutil.AssertEqual(t, ErrorObjectNull, obj.RequestSettings.IsNull(), true)
}

func TestTransformAuthenticationNil(t *testing.T) {
	obj, diags := jsonToTerraformAuthentication(nil)

	testutil.AssertDiagsHasNoErrors(t, diags)
	testutil.AssertEqual(t, ErrorObjectNull, obj.IsNull(), true)
}

func TestTransformHTTPRequestHeadersNil(t *testing.T) {
	obj, diags := jsonToTerraformHTTPRequestHeaders(nil)

	testutil.AssertDiagsHasNoErrors(t, diags)
	testutil.AssertEqual(t, ErrorObjectNull, obj.IsNull(), true)
}

func TestTransformHTTPRequestHeadersEmpty(t *testing.T) {
	headers := []models.HTTPHeaderRequestJSON{}

	obj, diags := jsonToTerraformHTTPRequestHeaders(&headers)

	testutil.AssertDiagsHasNoErrors(t, diags)
	testutil.AssertEqual(t, ErrorObjectNull, obj.IsNull(), false)
}

func TestTransformHTTPRequestHeadersWithValues(t *testing.T) {
	headers := []models.HTTPHeaderRequestJSON{
		{
			HeaderName:   testutil.ToStringPtr("UserAgent"),
			RequestValue: "CatchpointTestAgent",
			RequestHeaderType: models.GenericIDNameJSON{
				ID:   1,
				Name: "UserAgent",
			},
		},
		{
			HeaderName:   testutil.ToStringPtr("Sni-Override"),
			RequestValue: "somevalue",
			RequestHeaderType: models.GenericIDNameJSON{
				ID:   11,
				Name: "Custom",
			},
		},
		{
			HeaderName:   testutil.ToStringPtr("X-Custom-Header"),
			RequestValue: "A custom value",
			RequestHeaderType: models.GenericIDNameJSON{
				ID:   11,
				Name: "Custom",
			},
		},
	}

	obj, diags := jsonToTerraformHTTPRequestHeaders(&headers)

	testutil.AssertDiagsHasNoErrors(t, diags)
	testutil.AssertEqual(t, ErrorObjectNull, obj.IsNull(), false)

	sniHeader := obj.Attributes()["sni_override"].(types.Object).Attributes()
	testutil.AssertEqual(t, "sni_override header name", sniHeader["header_name"].(types.String).ValueString(), "Sni-Override")
	testutil.AssertEqual(t, "sni_override value", sniHeader["value"].(types.String).ValueString(), "somevalue")

	userAgentHeader := obj.Attributes()["user_agent"].(types.Object).Attributes()
	testutil.AssertEqual(t, "user_agent value", userAgentHeader["value"].(types.String).ValueString(), "CatchpointTestAgent")

	customHeader := obj.Attributes()["custom"].(types.Object).Attributes()
	testutil.AssertEqual(t, "custom header name", customHeader["header_name"].(types.String).ValueString(), "X-Custom-Header")
	testutil.AssertEqual(t, "custom value", customHeader["value"].(types.String).ValueString(), "A custom value")
}
