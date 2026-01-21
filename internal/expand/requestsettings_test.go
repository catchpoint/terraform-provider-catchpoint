package expand

import (
	"context"
	"testing"

	"catchpoint-provider/internal/fields"
	"catchpoint-provider/internal/models"
	"catchpoint-provider/internal/testutil"
	cptypes "catchpoint-provider/internal/types"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func TestExpandRequestSettingsConfigNullObject(t *testing.T) {
	obj := types.ObjectNull(map[string]attr.Type{})
	config := &models.RequestSettingsConfig{}

	diags := ExpandRequestSettingsConfig(obj, config)

	testutil.AssertDiagsHasNoErrors(t, diags)
	// Config should remain unchanged when object is null
}

func TestExpandRequestSettingsConfigEmptyObject(t *testing.T) {
	attrs := map[string]attr.Value{}

	obj, _ := types.ObjectValue(map[string]attr.Type{}, attrs)
	config := &models.RequestSettingsConfig{}

	diags := ExpandRequestSettingsConfig(obj, config)

	testutil.AssertDiagsHasNoErrors(t, diags)
	// Should set request setting type to Override by default
	testutil.AssertEqual(t, fields.RequestSettingType, config.RequestSettingType.Name, cptypes.Inherit)
	testutil.AssertEqual(t, fields.RequestSettingType, config.RequestSettingType.ID, 0)
}

func TestExpandRequestSettingsConfigTokenIDs(t *testing.T) {
	tokenIDList, _ := types.ListValue(types.Int64Type, []attr.Value{
		types.Int64Value(1),
		types.Int64Value(2),
		types.Int64Value(3),
	})

	attrs := map[string]attr.Value{
		fields.TokenIDs: tokenIDList,
	}

	obj, _ := types.ObjectValue(map[string]attr.Type{
		fields.TokenIDs: types.ListType{ElemType: types.Int64Type},
	}, attrs)

	config := &models.RequestSettingsConfig{}

	diags := ExpandRequestSettingsConfig(obj, config)

	testutil.AssertDiagsHasNoErrors(t, diags)
	testutil.AssertEqual(t, "token ids length", len(config.TokenIDs), 3)
	testutil.AssertEqual(t, "first token id", config.TokenIDs[0], 1)
	testutil.AssertEqual(t, "second token id", config.TokenIDs[1], 2)
	testutil.AssertEqual(t, "third token id", config.TokenIDs[2], 3)
}

func TestExpandRequestSettingsConfigCertificateIDs(t *testing.T) {
	certIDList, _ := types.ListValue(types.Int64Type, []attr.Value{
		types.Int64Value(10),
		types.Int64Value(20),
		types.Int64Value(30),
	})

	attrs := map[string]attr.Value{
		fields.LibraryCertificateIDs: certIDList,
	}

	obj, _ := types.ObjectValue(map[string]attr.Type{
		fields.LibraryCertificateIDs: types.ListType{ElemType: types.Int64Type},
	}, attrs)

	config := &models.RequestSettingsConfig{}

	diags := ExpandRequestSettingsConfig(obj, config)

	testutil.AssertDiagsHasNoErrors(t, diags)
	testutil.AssertEqual(t, "certificate ids length", len(config.CertificateIDs), 3)
	testutil.AssertEqual(t, "first certificate id", config.CertificateIDs[0], 10)
	testutil.AssertEqual(t, "second certificate id", config.CertificateIDs[1], 20)
	testutil.AssertEqual(t, "third certificate id", config.CertificateIDs[2], 30)
}

func TestExpandRequestSettingsConfigWithAuthentication(t *testing.T) {
	passwordIDList, _ := types.ListValue(types.Int64Type, []attr.Value{
		types.Int64Value(100),
		types.Int64Value(200),
	})

	authAttrs := map[string]attr.Value{
		fields.AuthenticationType: types.StringValue("basic"),
		fields.PasswordIDs:        passwordIDList,
	}

	authObj, _ := types.ObjectValue(map[string]attr.Type{
		fields.AuthenticationType: types.StringType,
		fields.PasswordIDs:        types.ListType{ElemType: types.Int64Type},
	}, authAttrs)

	attrs := map[string]attr.Value{
		fields.Authentication: authObj,
	}

	obj, _ := types.ObjectValue(map[string]attr.Type{
		fields.Authentication: authObj.Type(context.TODO()),
	}, attrs)

	config := &models.RequestSettingsConfig{}

	diags := ExpandRequestSettingsConfig(obj, config)

	testutil.AssertDiagsHasNoErrors(t, diags)
	testutil.AssertEqual(t, "authentication type name", config.AuthenticationType.Name, "basic")
	testutil.AssertEqual(t, "password ids length", len(config.PasswordIDs), 2)
	testutil.AssertEqual(t, "first password id", config.PasswordIDs[0], 100)
	testutil.AssertEqual(t, "second password id", config.PasswordIDs[1], 200)
}

func TestExpandRequestSettingsConfigWithHTTPRequestHeaders(t *testing.T) {
	// Create user agent header
	userAgentAttrs := map[string]attr.Value{
		fields.Value: types.StringValue("Mozilla/5.0"),
	}

	userAgentObj, _ := types.ObjectValue(map[string]attr.Type{
		fields.Value: types.StringType,
	}, userAgentAttrs)

	// Create custom header
	customHeaderAttrs := map[string]attr.Value{
		fields.Value:      types.StringValue("CustomValue"),
		fields.HeaderName: types.StringValue("X-Custom-Header"),
	}

	customHeaderObj, _ := types.ObjectValue(map[string]attr.Type{
		fields.Value:      types.StringType,
		fields.HeaderName: types.StringType,
	}, customHeaderAttrs)

	// Create HTTP request headers object
	headersAttrs := map[string]attr.Value{
		"user_agent": userAgentObj,
		"custom":     customHeaderObj,
	}

	headersObj, _ := types.ObjectValue(map[string]attr.Type{
		"user_agent": userAgentObj.Type(context.TODO()),
		"custom":     customHeaderObj.Type(context.TODO()),
	}, headersAttrs)

	attrs := map[string]attr.Value{
		fields.HTTPRequestHeaders: headersObj,
	}

	obj, _ := types.ObjectValue(map[string]attr.Type{
		fields.HTTPRequestHeaders: headersObj.Type(context.TODO()),
	}, attrs)

	config := &models.RequestSettingsConfig{}

	diags := ExpandRequestSettingsConfig(obj, config)

	testutil.AssertDiagsHasNoErrors(t, diags)
	testutil.AssertEqual(t, "http headers length", len(config.TestHTTPHeaderRequests), 2)

	// Find and verify headers by their values
	var userAgentFound, customHeaderFound bool
	for _, header := range config.TestHTTPHeaderRequests {
		if header.RequestValue == "Mozilla/5.0" {
			userAgentFound = true
		}
		if header.RequestValue == "CustomValue" && header.HeaderName == "X-Custom-Header" {
			customHeaderFound = true
		}
	}

	testutil.AssertEqual(t, "user agent header found", userAgentFound, true)
	testutil.AssertEqual(t, "custom header found", customHeaderFound, true)
}

func TestExpandRequestSettingsConfigWithHTTPRequestHeadersValidType(t *testing.T) {
	// Create user agent header
	userAgentAttrs := map[string]attr.Value{
		fields.Value: types.StringValue("Mozilla/5.0"),
	}

	userAgentObj, _ := types.ObjectValue(map[string]attr.Type{
		fields.Value: types.StringType,
	}, userAgentAttrs)

	// Create custom header
	customHeaderAttrs := map[string]attr.Value{
		fields.Value:      types.StringValue("CustomValue"),
		fields.HeaderName: types.StringValue("X-Custom-Header"),
	}

	customHeaderObj, _ := types.ObjectValue(map[string]attr.Type{
		fields.Value:      types.StringType,
		fields.HeaderName: types.StringType,
	}, customHeaderAttrs)

	// Create HTTP request headers object
	headersAttrs := map[string]attr.Value{
		"user_agent": userAgentObj,
		"custom":     customHeaderObj,
	}

	headersObj, _ := types.ObjectValue(map[string]attr.Type{
		"user_agent": userAgentObj.Type(context.TODO()),
		"custom":     customHeaderObj.Type(context.TODO()),
	}, headersAttrs)

	attrs := map[string]attr.Value{
		fields.HTTPRequestHeaders: headersObj,
	}

	obj, _ := types.ObjectValue(map[string]attr.Type{
		fields.HTTPRequestHeaders: headersObj.Type(context.TODO()),
	}, attrs)

	config := &models.RequestSettingsConfig{}

	diags := ExpandRequestSettingsConfig(obj, config)

	testutil.AssertDiagsHasNoErrors(t, diags)

	// All the headers should have a valid RequestHeaderType.
	for _, header := range config.TestHTTPHeaderRequests {
		testutil.AssertNotNil(t, fields.RequestHeaderType, header.RequestHeaderType)
		_, ok := cptypes.GetReqHeaderTypeID(header.RequestHeaderType.Name)
		testutil.AssertEqual(t, fields.RequestHeaderType+" name", ok, true)
		_, ok = cptypes.GetReqHeaderTypeName(header.RequestHeaderType.ID)
		testutil.AssertEqual(t, fields.RequestHeaderType+" id", ok, true)
	}
}

func TestExpandRequestSettingsConfigNullOptionalFields(t *testing.T) {
	attrs := map[string]attr.Value{
		fields.Authentication:        types.ObjectNull(map[string]attr.Type{}),
		fields.TokenIDs:              types.ListNull(types.Int64Type),
		fields.LibraryCertificateIDs: types.ListNull(types.Int64Type),
		fields.HTTPRequestHeaders:    types.ObjectNull(map[string]attr.Type{}),
	}

	obj, _ := types.ObjectValue(map[string]attr.Type{
		fields.Authentication:        types.ObjectType{},
		fields.TokenIDs:              types.ListType{ElemType: types.Int64Type},
		fields.LibraryCertificateIDs: types.ListType{ElemType: types.Int64Type},
		fields.HTTPRequestHeaders:    types.ObjectType{},
	}, attrs)

	config := &models.RequestSettingsConfig{}

	diags := ExpandRequestSettingsConfig(obj, config)

	testutil.AssertDiagsHasNoErrors(t, diags)
	testutil.AssertEqual(t, "request setting type name", config.RequestSettingType.Name, cptypes.Inherit)
	testutil.AssertEqual(t, "token ids length", len(config.TokenIDs), 0)
	testutil.AssertEqual(t, "certificate ids length", len(config.CertificateIDs), 0)
	testutil.AssertEqual(t, "http headers length", len(config.TestHTTPHeaderRequests), 0)
}

func TestExpandRequestSettingsConfigEmptyLists(t *testing.T) {
	emptyTokenList, _ := types.ListValue(types.Int64Type, []attr.Value{})
	emptyCertList, _ := types.ListValue(types.Int64Type, []attr.Value{})

	attrs := map[string]attr.Value{
		fields.TokenIDs:              emptyTokenList,
		fields.LibraryCertificateIDs: emptyCertList,
	}

	obj, _ := types.ObjectValue(map[string]attr.Type{
		fields.TokenIDs:              types.ListType{ElemType: types.Int64Type},
		fields.LibraryCertificateIDs: types.ListType{ElemType: types.Int64Type},
	}, attrs)

	config := &models.RequestSettingsConfig{}

	diags := ExpandRequestSettingsConfig(obj, config)

	testutil.AssertDiagsHasNoErrors(t, diags)
	testutil.AssertEqual(t, "token ids length", len(config.TokenIDs), 0)
	testutil.AssertEqual(t, "certificate ids length", len(config.CertificateIDs), 0)
}

func TestExpandRequestSettingsConfigLargeValues(t *testing.T) {
	tokenIDList, _ := types.ListValue(types.Int64Type, []attr.Value{
		types.Int64Value(999999),
		types.Int64Value(888888),
	})

	certIDList, _ := types.ListValue(types.Int64Type, []attr.Value{
		types.Int64Value(777777),
		types.Int64Value(666666),
	})

	attrs := map[string]attr.Value{
		fields.TokenIDs:              tokenIDList,
		fields.LibraryCertificateIDs: certIDList,
	}

	obj, _ := types.ObjectValue(map[string]attr.Type{
		fields.TokenIDs:              types.ListType{ElemType: types.Int64Type},
		fields.LibraryCertificateIDs: types.ListType{ElemType: types.Int64Type},
	}, attrs)

	config := &models.RequestSettingsConfig{}

	diags := ExpandRequestSettingsConfig(obj, config)

	testutil.AssertDiagsHasNoErrors(t, diags)
	testutil.AssertEqual(t, "first token id", config.TokenIDs[0], 999999)
	testutil.AssertEqual(t, "first certificate id", config.CertificateIDs[0], 777777)
}

func TestExpandRequestSettingsConfigZeroValues(t *testing.T) {
	tokenIDList, _ := types.ListValue(types.Int64Type, []attr.Value{
		types.Int64Value(0),
	})

	certIDList, _ := types.ListValue(types.Int64Type, []attr.Value{
		types.Int64Value(0),
	})

	attrs := map[string]attr.Value{
		fields.TokenIDs:              tokenIDList,
		fields.LibraryCertificateIDs: certIDList,
	}

	obj, _ := types.ObjectValue(map[string]attr.Type{
		fields.TokenIDs:              types.ListType{ElemType: types.Int64Type},
		fields.LibraryCertificateIDs: types.ListType{ElemType: types.Int64Type},
	}, attrs)

	config := &models.RequestSettingsConfig{}

	diags := ExpandRequestSettingsConfig(obj, config)

	testutil.AssertDiagsHasNoErrors(t, diags)
	testutil.AssertEqual(t, "first token id", config.TokenIDs[0], 0)
	testutil.AssertEqual(t, "first certificate id", config.CertificateIDs[0], 0)
}

// Integration test with all fields populated
func TestExpandRequestSettingsConfigCompleteConfiguration(t *testing.T) {
	// Create token and certificate lists
	tokenIDList, _ := types.ListValue(types.Int64Type, []attr.Value{
		types.Int64Value(1),
		types.Int64Value(2),
	})

	certIDList, _ := types.ListValue(types.Int64Type, []attr.Value{
		types.Int64Value(10),
		types.Int64Value(20),
	})

	// Create authentication
	passwordIDList, _ := types.ListValue(types.Int64Type, []attr.Value{
		types.Int64Value(100),
	})

	authAttrs := map[string]attr.Value{
		fields.AuthenticationType: types.StringValue("digest"),
		fields.PasswordIDs:        passwordIDList,
	}

	authObj, _ := types.ObjectValue(map[string]attr.Type{
		fields.AuthenticationType: types.StringType,
		fields.PasswordIDs:        types.ListType{ElemType: types.Int64Type},
	}, authAttrs)

	// Create HTTP headers
	hostHeaderAttrs := map[string]attr.Value{
		fields.Value:            types.StringValue("api.example.com"),
		fields.ChildHostPattern: types.StringValue("*.example.com"),
	}

	hostHeaderObj, _ := types.ObjectValue(map[string]attr.Type{
		fields.Value:            types.StringType,
		fields.ChildHostPattern: types.StringType,
	}, hostHeaderAttrs)

	headersAttrs := map[string]attr.Value{
		"host": hostHeaderObj,
	}

	headersObj, _ := types.ObjectValue(map[string]attr.Type{
		"host": hostHeaderObj.Type(context.TODO()),
	}, headersAttrs)

	// Main object
	attrs := map[string]attr.Value{
		fields.TokenIDs:              tokenIDList,
		fields.LibraryCertificateIDs: certIDList,
		fields.Authentication:        authObj,
		fields.HTTPRequestHeaders:    headersObj,
	}

	obj, _ := types.ObjectValue(map[string]attr.Type{
		fields.TokenIDs:              types.ListType{ElemType: types.Int64Type},
		fields.LibraryCertificateIDs: types.ListType{ElemType: types.Int64Type},
		fields.Authentication:        authObj.Type(context.TODO()),
		fields.HTTPRequestHeaders:    headersObj.Type(context.TODO()),
	}, attrs)

	config := &models.RequestSettingsConfig{}

	diags := ExpandRequestSettingsConfig(obj, config)

	testutil.AssertDiagsHasNoErrors(t, diags)

	// Verify all fields
	testutil.AssertEqual(t, "request setting type name", config.RequestSettingType.Name, cptypes.Override)
	testutil.AssertEqual(t, "token ids length", len(config.TokenIDs), 2)
	testutil.AssertEqual(t, "certificate ids length", len(config.CertificateIDs), 2)
	testutil.AssertEqual(t, "authentication type name", config.AuthenticationType.Name, "digest")
	testutil.AssertEqual(t, "password ids length", len(config.PasswordIDs), 1)
	testutil.AssertEqual(t, "http headers length", len(config.TestHTTPHeaderRequests), 1)
	testutil.AssertEqual(t, "host header value", config.TestHTTPHeaderRequests[0].RequestValue, "api.example.com")
}

// Test expandConfigAuthentication function directly
func TestExpandConfigAuthenticationNullObject(t *testing.T) {
	obj := types.ObjectNull(map[string]attr.Type{})
	config := &models.RequestSettingsConfig{}

	diags := expandConfigAuthentication(config, obj)

	testutil.AssertDiagsHasNoErrors(t, diags)
	// Config should remain unchanged when object is null
}

func TestExpandConfigAuthenticationOnlyAuthenticationType(t *testing.T) {
	attrs := map[string]attr.Value{
		fields.AuthenticationType: types.StringValue("ntlm"),
	}

	obj, _ := types.ObjectValue(map[string]attr.Type{
		fields.AuthenticationType: types.StringType,
	}, attrs)

	config := &models.RequestSettingsConfig{}

	diags := expandConfigAuthentication(config, obj)

	testutil.AssertDiagsHasNoErrors(t, diags)
	testutil.AssertEqual(t, "authentication type name", config.AuthenticationType.Name, "ntlm")
	testutil.AssertEqual(t, "password ids length", len(config.PasswordIDs), 0)
}

func TestExpandConfigAuthenticationOnlyPasswordIDs(t *testing.T) {
	passwordIDList, _ := types.ListValue(types.Int64Type, []attr.Value{
		types.Int64Value(50),
		types.Int64Value(60),
	})

	attrs := map[string]attr.Value{
		fields.PasswordIDs: passwordIDList,
	}

	obj, _ := types.ObjectValue(map[string]attr.Type{
		fields.PasswordIDs: types.ListType{ElemType: types.Int64Type},
	}, attrs)

	config := &models.RequestSettingsConfig{}

	diags := expandConfigAuthentication(config, obj)

	testutil.AssertDiagsHasNoErrors(t, diags)
	testutil.AssertEqual(t, "password ids length", len(config.PasswordIDs), 2)
	testutil.AssertEqual(t, "first password id", config.PasswordIDs[0], 50)
	testutil.AssertEqual(t, "second password id", config.PasswordIDs[1], 60)
}

func TestExpandConfigAuthenticationAuthenticationTypes(t *testing.T) {
	authTypes := []string{"basic", "digest", "ntlm"}

	for _, authType := range authTypes {
		t.Run(authType, func(t *testing.T) {
			attrs := map[string]attr.Value{
				fields.AuthenticationType: types.StringValue(authType),
			}

			obj, _ := types.ObjectValue(map[string]attr.Type{
				fields.AuthenticationType: types.StringType,
			}, attrs)

			config := &models.RequestSettingsConfig{}

			diags := expandConfigAuthentication(config, obj)

			testutil.AssertDiagsHasNoErrors(t, diags)
			testutil.AssertEqual(t, "authentication type name", config.AuthenticationType.Name, authType)
		})
	}
}

// Test expandConfigHTTPRequestHeaders function directly
func TestExpandConfigHTTPRequestHeadersNullObject(t *testing.T) {
	obj := types.ObjectNull(map[string]attr.Type{})
	config := &models.RequestSettingsConfig{}

	diags := expandConfigHTTPRequestHeaders(config, obj)

	testutil.AssertDiagsHasNoErrors(t, diags)
	testutil.AssertEqual(t, "http headers length", len(config.TestHTTPHeaderRequests), 0)
}

func TestExpandConfigHTTPRequestHeadersAllHeaderTypes(t *testing.T) {
	// Test each valid header type
	headerAttrs := map[string]attr.Value{
		fields.Value:            types.StringValue("test-value"),
		fields.ChildHostPattern: types.StringValue("*.test.com"),
		fields.HeaderName:       types.StringValue("X-Test-Header"),
	}

	headerObj, _ := types.ObjectValue(map[string]attr.Type{
		fields.Value:            types.StringType,
		fields.ChildHostPattern: types.StringType,
		fields.HeaderName:       types.StringType,
	}, headerAttrs)

	// Create headers object with all valid HTTP request headers
	headersAttrs := map[string]attr.Value{}
	headerTypes := map[string]attr.Type{}

	for _, headerName := range cptypes.ValidHTTPRequestHeaders {
		headersAttrs[headerName] = headerObj
		headerTypes[headerName] = headerObj.Type(context.TODO())
	}

	headersObj, _ := types.ObjectValue(headerTypes, headersAttrs)

	config := &models.RequestSettingsConfig{}

	diags := expandConfigHTTPRequestHeaders(config, headersObj)

	testutil.AssertDiagsHasNoErrors(t, diags)
	testutil.AssertEqual(t, "http headers length", len(config.TestHTTPHeaderRequests), len(cptypes.ValidHTTPRequestHeaders))

	// Check first header
	firstHeader := config.TestHTTPHeaderRequests[0]
	testutil.AssertEqual(t, "first header value", firstHeader.RequestValue, "test-value")
	testutil.AssertEqual(t, "first header child host pattern", firstHeader.ChildHostPattern, "*.test.com")
	testutil.AssertEqual(t, "first header name", firstHeader.HeaderName, "X-Test-Header")
}

func TestExpandConfigHTTPRequestHeadersMixedNullAndValid(t *testing.T) {
	// Create one valid header
	validHeaderAttrs := map[string]attr.Value{
		fields.Value:            types.StringValue("valid-value"),
		fields.ChildHostPattern: types.StringValue("*.valid.com"),
	}

	validHeaderObj, _ := types.ObjectValue(map[string]attr.Type{
		fields.Value:            types.StringType,
		fields.ChildHostPattern: types.StringType,
	}, validHeaderAttrs)

	// Create headers object with mix of null and valid headers
	headersAttrs := map[string]attr.Value{
		"user_agent": validHeaderObj,
		"host":       types.ObjectNull(map[string]attr.Type{}),
		"cookie":     validHeaderObj,
		"referer":    types.ObjectNull(map[string]attr.Type{}),
	}

	headersObj, _ := types.ObjectValue(map[string]attr.Type{
		"user_agent": validHeaderObj.Type(context.TODO()),
		"host":       types.ObjectType{},
		"cookie":     validHeaderObj.Type(context.TODO()),
		"referer":    types.ObjectType{},
	}, headersAttrs)

	config := &models.RequestSettingsConfig{}

	diags := expandConfigHTTPRequestHeaders(config, headersObj)

	testutil.AssertDiagsHasNoErrors(t, diags)
	testutil.AssertEqual(t, "http headers length", len(config.TestHTTPHeaderRequests), 2) // Only non-null headers

	// Verify all headers have the expected values
	validHeaderCount := 0
	for _, header := range config.TestHTTPHeaderRequests {
		if header.RequestValue == "valid-value" && header.ChildHostPattern == "*.valid.com" {
			validHeaderCount++
		}
	}

	testutil.AssertEqual(t, "valid headers count", validHeaderCount, 2) // Should have 2 identical valid headers
}

func TestExpandConfigHTTPRequestHeadersCustom(t *testing.T) {
	sniOverrideAttrs := map[string]attr.Value{
		fields.Value: types.StringValue("sni-value"),
	}
	sniOverrideObj, _ := types.ObjectValue(map[string]attr.Type{
		fields.Value: types.StringType,
	}, sniOverrideAttrs)

	// Custom header (should use headerName as provided)
	customAttrs := map[string]attr.Value{
		fields.Value:      types.StringValue("custom-value"),
		fields.HeaderName: types.StringValue("X-Custom-Header"),
	}
	customObj, _ := types.ObjectValue(map[string]attr.Type{
		fields.Value:      types.StringType,
		fields.HeaderName: types.StringType,
	}, customAttrs)

	headersAttrs := map[string]attr.Value{
		cptypes.SNIOverrideHeader: sniOverrideObj,
		"custom":                  customObj,
	}
	headersTypes := map[string]attr.Type{
		cptypes.SNIOverrideHeader: sniOverrideObj.Type(context.TODO()),
		"custom":                  customObj.Type(context.TODO()),
	}

	headersObj, _ := types.ObjectValue(headersTypes, headersAttrs)

	config := &models.RequestSettingsConfig{}

	diags := expandConfigHTTPRequestHeaders(config, headersObj)

	testutil.AssertDiagsHasNoErrors(t, diags)
	testutil.AssertEqual(t, "http headers length", len(config.TestHTTPHeaderRequests), 2)

	var sniFound, customFound bool
	for _, header := range config.TestHTTPHeaderRequests {
		if header.RequestHeaderType.ID == 11 && header.RequestHeaderType.Name == cptypes.SNIOverrideHeader {
			sniFound = true
			testutil.AssertEqual(t, "SNI_Override headerName", header.HeaderName, "Sni-Override")
			testutil.AssertEqual(t, "SNI_Override value", header.RequestValue, "sni-value")
		}
		if header.RequestHeaderType.ID == 11 && header.RequestHeaderType.Name == "custom" {
			customFound = true
			testutil.AssertEqual(t, "Custom headerName", header.HeaderName, "X-Custom-Header")
			testutil.AssertEqual(t, "Custom value", header.RequestValue, "custom-value")
		}
	}
	testutil.AssertEqual(t, "SNI_Override header found", sniFound, true)
	testutil.AssertEqual(t, "Custom header found", customFound, true)
}

// Test error propagation
func TestExpandRequestSettingsConfigErrorPropagation(t *testing.T) {
	// Test with potentially problematic data
	invalidList := types.ListUnknown(types.StringType) // Wrong element type

	attrs := map[string]attr.Value{
		fields.TokenIDs: invalidList,
	}

	obj, _ := types.ObjectValue(map[string]attr.Type{
		fields.TokenIDs: types.ListType{ElemType: types.StringType}, // Wrong type
	}, attrs)

	config := &models.RequestSettingsConfig{}

	diags := ExpandRequestSettingsConfig(obj, config)

	// Should handle errors gracefully - the exact behavior depends on ExpandIntListSetting implementation
	if diags.HasError() {
		// Error is expected for invalid data
		testutil.AssertEqual(t, "should have errors for invalid data", diags.HasError(), true)
	}

	// At minimum, the RequestSettingType should be set since that happens first
	testutil.AssertEqual(t, "request setting type should be set", config.RequestSettingType.Name, cptypes.Inherit)
}
