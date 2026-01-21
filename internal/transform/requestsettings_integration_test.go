package transform

import (
	"testing"

	"catchpoint-provider/internal/fields"
	"catchpoint-provider/internal/models"
	"catchpoint-provider/internal/testutil"

	"github.com/hashicorp/terraform-plugin-framework/types"
)

const (
	ErrorDiagnostics = "diagnostics has errors"
	ErrorObjectNull  = "object is null"
	ErrorListNull    = "list is null"
)

func TestTransformRequestSettings(t *testing.T) {
	requestSetting := &models.RequestSettingsJSON{
		RequestSettingType:    models.GenericIDNameJSON{ID: 1, Name: "Override"},
		LibraryCertificateIDs: &[]int{1, 2, 3},
		TokenIDs:              &[]int{4, 5},
		HTTPHeaderRequests:    &[]models.HTTPHeaderRequestJSON{},
	}

	obj, diags := JSONToTerraformRequestSettings(requestSetting)

	testutil.AssertDiagsHasNoErrors(t, diags)
	testutil.AssertEqual(t, ErrorObjectNull, obj.RequestSettings.IsNull(), false)

	objAttrs := obj.RequestSettings.Attributes()

	// Check certificate IDs are included
	testutil.AssertListTypeAndLength(t, objAttrs[fields.LibraryCertificateIDs].(types.List), 3)
	// Check token IDs
	testutil.AssertListTypeAndLength(t, objAttrs[fields.TokenIDs].(types.List), 2)
}

func TestTransformAuthenticationValid(t *testing.T) {
	auth := &models.AuthenticationJSON{
		AuthenticationMethodType: &models.GenericIDNameOmitEmptyJSON{
			ID:   testutil.ToIntPtr(1),
			Name: func() *string { v := "Basic"; return &v }(),
		},
		PasswordIDs: &[]int{1, 2, 3},
	}

	obj, diags := jsonToTerraformAuthentication(auth)

	testutil.AssertDiagsHasNoErrors(t, diags)
	testutil.AssertEqual(t, ErrorObjectNull, obj.IsNull(), false)

	objAttrs := obj.Attributes()

	// Check authentication type. It should be "Basic" normalized to "basic".
	testutil.AssertStringTypeAndValue(t, fields.AuthenticationType, objAttrs, "basic")

	// Check password IDs
	testutil.AssertListTypeAndLength(t, objAttrs[fields.PasswordIDs].(types.List), 3)
}

func TestTransformAuthenticationEmptyPasswordIDs(t *testing.T) {
	auth := &models.AuthenticationJSON{
		AuthenticationMethodType: &models.GenericIDNameOmitEmptyJSON{Name: func() *string { v := "None"; return &v }()},
		PasswordIDs:              nil,
	}

	obj, diags := jsonToTerraformAuthentication(auth)

	testutil.AssertDiagsHasNoErrors(t, diags)

	objAttrs := obj.Attributes()

	// Password IDs should be null list
	testutil.AssertNullListTypeForAttribute(t, fields.PasswordIDs, objAttrs)
}

func TestTransformHTTPRequestHeadersValid(t *testing.T) {
	tests := []struct {
		headerTypeID           int
		headerTypeName         string
		expectedHeaderTypeName string
		value                  string
		childHostPattern       string
		headerName             string
		headerNameExpected     bool
	}{
		{
			headerTypeName:         "UserAgent",
			headerTypeID:           1,
			expectedHeaderTypeName: "user_agent",
			value:                  "Mozilla/5.0",
			childHostPattern:       "*.example.com",
			headerNameExpected:     false,
		},
		{
			headerTypeName:         "Custom",
			headerTypeID:           11,
			expectedHeaderTypeName: "custom",
			value:                  "Some Value",
			childHostPattern:       "",
			headerName:             "X-Custom-Header",
			headerNameExpected:     true,
		},
		{
			headerTypeName:         "Custom",
			headerTypeID:           11,
			expectedHeaderTypeName: "sni_override",
			value:                  "Some Value",
			childHostPattern:       "",
			headerName:             "Sni-Override",
			headerNameExpected:     true,
		},
	}
	for _, tt := range tests {
		headers := []models.HTTPHeaderRequestJSON{
			{
				RequestHeaderType: models.GenericIDNameJSON{ID: tt.headerTypeID, Name: tt.headerTypeName},
				RequestValue:      tt.value,
				ChildHostPattern:  &tt.childHostPattern,
				HeaderName:        &tt.headerName,
			},
		}

		obj, diags := jsonToTerraformHTTPRequestHeaders(&headers)

		testutil.AssertDiagsHasNoErrors(t, diags)
		testutil.AssertEqual(t, ErrorObjectNull, obj.IsNull(), false)

		// Check that the header is present
		objAttrs := obj.Attributes()
		headerObj := testutil.AssertObjectTypeAndNotNull(t, tt.expectedHeaderTypeName, objAttrs)

		if tt.headerNameExpected {
			testutil.AssertStringTypeAndValue(t, fields.HeaderName, headerObj.Attributes(), tt.headerName)
		}
	}
}

// Integration test to verify the complete flow
func TestTransformRequestSettingsCompleteFlow(t *testing.T) {
	requestSetting := &models.RequestSettingsJSON{
		RequestSettingType: models.GenericIDNameJSON{ID: 1, Name: "override"},
		Authentication: &models.AuthenticationJSON{
			AuthenticationMethodType: &models.GenericIDNameOmitEmptyJSON{Name: func() *string { v := "None"; return &v }()},
			PasswordIDs:              func() *[]int { v := []int{100, 200}; return &v }(),
		},
		TokenIDs:              &[]int{300, 400},
		LibraryCertificateIDs: &[]int{500, 600},
		HTTPHeaderRequests: &[]models.HTTPHeaderRequestJSON{
			{
				RequestHeaderType: models.GenericIDNameJSON{ID: 1, Name: "user_agent"},
				RequestValue:      "Test-Agent/1.0",
				ChildHostPattern:  func() *string { v := "*.test.com"; return &v }(),
			},
			{
				RequestHeaderType: models.GenericIDNameJSON{ID: 11, Name: "custom"},
				RequestValue:      "test-value",
				HeaderName:        func() *string { v := "X-Test-Header"; return &v }(),
			},
		},
	}

	obj, diags := JSONToTerraformRequestSettings(requestSetting)

	testutil.AssertDiagsHasNoErrors(t, diags)
	testutil.AssertEqual(t, ErrorObjectNull, obj.RequestSettings.IsNull(), false)

	objAttrs := obj.RequestSettings.Attributes()

	// Verify all top-level fields are present
	testutil.AssertNotNil(t, "request setting type", objAttrs[fields.RequestSettingType])
	testutil.AssertNotNil(t, "authentication", objAttrs[fields.Authentication])
	testutil.AssertNotNil(t, "token ids", objAttrs[fields.TokenIDs])
	testutil.AssertNotNil(t, "library certificate ids", objAttrs[fields.LibraryCertificateIDs])
	testutil.AssertNotNil(t, "http request headers", objAttrs[fields.HTTPRequestHeaders])

	// Verify request setting type
	testutil.AssertStringTypeAndValue(t, fields.RequestSettingType, objAttrs, "override")

	// Verify authentication object
	testutil.AssertObjectTypeAndNotNull(t, fields.Authentication, objAttrs)

	// Verify token IDs list
	testutil.AssertListTypeAndLength(t, objAttrs[fields.TokenIDs].(types.List), 2)

	// Verify certificate IDs list
	testutil.AssertListTypeAndLength(t, objAttrs[fields.LibraryCertificateIDs].(types.List), 2)

	// Verify HTTP headers list
	headersList, ok := objAttrs[fields.HTTPRequestHeaders].(types.Object)
	if !ok {
		t.Fatalf("Expected HTTPRequestHeaders to be Object type")
	}
	testutil.AssertEqual(t, "headers list length", len(headersList.Attributes()), 18)

	// Verify specific headers
	userAgentHeader, ok := headersList.Attributes()["user_agent"].(types.Object)
	if !ok {
		t.Fatalf("Expected user_agent header to be Object type")
	}
	testutil.AssertEqual(t, "user_agent header value", userAgentHeader.Attributes()["value"].(types.String).ValueString(), "Test-Agent/1.0")
	testutil.AssertEqual(t, "user_agent child host pattern", userAgentHeader.Attributes()["child_host_pattern"].(types.String).ValueString(), "*.test.com")
}

// Test error propagation
func TestTransformRequestSettingsErrorPropagation(t *testing.T) {
	// Test with malformed data that might cause errors
	requestSetting := &models.RequestSettingsJSON{
		RequestSettingType: models.GenericIDNameJSON{Name: "Override"},
		HTTPHeaderRequests: &[]models.HTTPHeaderRequestJSON{
			{
				RequestHeaderType: models.GenericIDNameJSON{Name: "invalid_header_type"},
				RequestValue:      "some-value",
			},
		},
	}

	obj, diags := JSONToTerraformRequestSettings(requestSetting)

	// The function should handle errors gracefully
	// In case of errors, obj should be null and diagnostics should contain the error
	if diags.HasError() {
		testutil.AssertEqual(t, "object is null on error", obj.RequestSettings.IsNull(), true)
	}
}
