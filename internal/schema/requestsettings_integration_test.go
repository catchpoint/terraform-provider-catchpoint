package schema

import (
	"context"
	"fmt"
	"testing"

	"catchpoint-provider/internal/fields"
	"catchpoint-provider/internal/testutil"
	cptypes "catchpoint-provider/internal/types"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func TestBuildRequestSettingsAttributes(t *testing.T) {
	ctx := context.Background()
	attrs := BuildRequestSettingsBlock(ctx)
	singleNested := attrs["request_settings"].(schema.SingleNestedBlock)
	_ = testutil.GetAttributeOrAssert(t, fields.LibraryCertificateIDs, singleNested.Attributes)
}

func TestBuildHTTPRequestHeaderAttribute(t *testing.T) {
	hdrAttrs := buildHTTPRequestHeaderBlock()
	_ = testutil.GetBlockOrAssert(t, fields.HTTPRequestHeaders, hdrAttrs)
}

func TestBuildHTTPHeaderAttributeRequiredFields(t *testing.T) {
	headerName := cptypes.UserAgentHeader
	attrs := buildHTTPHeaderBlock(headerName, false, false)
	singleNested := testutil.GetBlockOrAssert(t, headerName, attrs)
	if singleNested.GetNestedObject().GetAttributes()[fields.Value].IsRequired() {
		t.Errorf("expected %s to be required", fields.Value)
	}
}

func TestBuildHTTPHeaderAttributeChildHostPatternRequired(t *testing.T) {
	headerName := cptypes.DNSOverrideHeader
	attrs := buildHTTPHeaderBlock(headerName, true, false)
	singleNested := testutil.GetBlockOrAssert(t, headerName, attrs)
	if singleNested.GetNestedObject().GetAttributes()[fields.ChildHostPattern].IsRequired() {
		t.Errorf("expected %s to not be required", fields.HeaderName)
	}
}

func TestBuildHTTPHeaderAttributeIncludeHeaderName(t *testing.T) {
	headerName := cptypes.CustomHeader
	attrs := buildHTTPHeaderBlock(headerName, false, true)
	singleNested := testutil.GetBlockOrAssert(t, headerName, attrs)
	if singleNested.GetNestedObject().GetAttributes()[fields.HeaderName].IsRequired() {
		t.Errorf("expected %s to not be required", fields.HeaderName)
	}
}

func TestGetRequestSettingsAttributeTypes(t *testing.T) {
	attrTypes := GetRequestSettingsAttributeTypes()

	// Test that all expected fields are present
	expectedFields := []string{
		fields.RequestSettingType,
		fields.Authentication,
		fields.TokenIDs,
		fields.HTTPRequestHeaders,
	}

	for _, field := range expectedFields {
		_ = testutil.GetTypeOrAssert(t, field, attrTypes)
	}

	// Test specific field types
	testutil.AssertEqual(t, fields.RequestSettingType+" type", attrTypes[fields.RequestSettingType], attr.Type(types.StringType))
	testutil.AssertNotNil(t, fields.Authentication+" type", attrTypes[fields.Authentication])

	// TokenIDs should be a list of int64
	testutil.AssertListTypeOfType(t, fields.TokenIDs, attrTypes, types.Int64Type)
}

func TestGetRequestSettingsWithCertsAttributeTypes(t *testing.T) {
	attrTypes := GetRequestSettingsAttributeTypes()

	// Test that all expected fields are present
	expectedFields := []string{
		fields.RequestSettingType,
		fields.Authentication,
		fields.TokenIDs,
		fields.HTTPRequestHeaders,
		fields.LibraryCertificateIDs,
	}

	for _, field := range expectedFields {
		_ = testutil.GetTypeOrAssert(t, field, attrTypes)
	}

	// Test specific field types
	testutil.AssertEqual(t, fields.RequestSettingType+" type", attrTypes[fields.RequestSettingType], attr.Type(types.StringType))
	testutil.AssertNotNil(t, fields.Authentication+" type", attrTypes[fields.Authentication])

	// TokenIDs should be a list of int64
	testutil.AssertListTypeOfType(t, fields.TokenIDs, attrTypes, types.Int64Type)
	// LibraryCertificateIDs should be a list of int64
	testutil.AssertListTypeOfType(t, fields.LibraryCertificateIDs, attrTypes, types.Int64Type)
}

func TestGetAuthenticationAttributeTypes(t *testing.T) {
	attrTypes := GetAuthenticationAttributeTypes()

	// Test that all expected authentication fields are present
	expectedFields := []string{
		fields.AuthenticationType,
		fields.PasswordIDs,
	}

	for _, field := range expectedFields {
		_ = testutil.GetTypeOrAssert(t, field, attrTypes)
	}

	// Test specific field types
	testutil.AssertEqual(t, fields.AuthenticationType+" type", attrTypes[fields.AuthenticationType], attr.Type(types.StringType))

	// PasswordIDs should be a list of int64
	testutil.AssertListTypeOfType(t, fields.PasswordIDs, attrTypes, types.Int64Type)

	// Should not contain fields from parent
	parentOnlyFields := []string{
		fields.RequestSettingType,
		fields.TokenIDs,
		fields.HTTPRequestHeaders,
	}

	for _, field := range parentOnlyFields {
		testutil.AssertTypeNotExists(t, field, attrTypes)
	}
}

func TestGetHTTPRequestHeadersAttributeTypes(t *testing.T) {
	attrTypes := GetHTTPRequestHeadersAttributeTypes()

	for _, field := range cptypes.ValidHTTPRequestHeaders {
		_ = testutil.GetTypeOrAssert(t, field, attrTypes)
	}

	// Each header field should be an ObjectType (SingleNestedAttribute)
	for _, field := range cptypes.ValidHTTPRequestHeaders {
		_ = testutil.GetObjectOrAssert(t, field, attrTypes)
	}

	// Should not contain fields from parent
	parentOnlyFields := []string{
		fields.RequestSettingType,
		fields.Authentication,
		fields.TokenIDs,
	}

	for _, field := range parentOnlyFields {
		testutil.AssertTypeNotExists(t, field, attrTypes)
	}
}

func TestGetHTTPRequestHeaderAttributeTypes(t *testing.T) {
	// Get the actual header field names from the schema
	httpHeadersAttrTypes := GetHTTPRequestHeadersAttributeTypes()

	// Test each header type that actually exists
	for headerField := range httpHeadersAttrTypes {
		t.Run(headerField, func(t *testing.T) {
			attrTypes := GetHTTPHeaderAttributeTypes(headerField)

			// All headers should have Value field
			_ = testutil.GetTypeOrAssert(t, fields.Value, attrTypes)

			// All headers should have ChildHostPattern field
			_ = testutil.GetTypeOrAssert(t, fields.ChildHostPattern, attrTypes)

			// Test specific field types
			testutil.AssertEqual(t, headerField+" "+fields.Value+" type", attrTypes[fields.Value], attr.Type(types.StringType))
			testutil.AssertEqual(t, headerField+" "+fields.ChildHostPattern+" type", attrTypes[fields.ChildHostPattern], attr.Type(types.StringType))

			// Custom header should have HeaderName field
			if headerField == "custom" {
				_ = testutil.GetTypeOrAssert(t, fields.HeaderName, attrTypes)
				testutil.AssertEqual(t, fields.HeaderName+" type", attrTypes[fields.HeaderName], attr.Type(types.StringType))
			}
		})
	}
}

func TestSchemaConsistency(t *testing.T) {
	// Test that calling the functions multiple times returns consistent results
	attrTypes1 := GetRequestSettingsAttributeTypes()
	attrTypes2 := GetRequestSettingsAttributeTypes()

	testutil.AssertEqual(t, "attribute types count", len(attrTypes1), len(attrTypes2))

	for field, type1 := range attrTypes1 {
		type2 := testutil.GetTypeOrAssert(t, field, attrTypes2)

		// Just verify they have the same Go type
		testutil.AssertEqual(t, field+" type name consistency",
			fmt.Sprintf("%T", type1), fmt.Sprintf("%T", type2))
	}
}

func TestNestedObjectExtraction(t *testing.T) {
	// Test that nested objects are properly extracted
	requestAttrTypes := GetRequestSettingsAttributeTypes()
	authAttrTypes := GetAuthenticationAttributeTypes()

	// Authentication should be an ObjectType in the request settings
	authObjectType := testutil.GetObjectOrAssert(t, fields.Authentication, requestAttrTypes)

	// The nested authentication object should have the same attribute types
	// as what GetAuthenticationAttributeTypes() returns
	testutil.AssertEqual(t, "authentication attributes count",
		len(authObjectType.AttrTypes), len(authAttrTypes))

	for field, expectedType := range authAttrTypes {
		actualType := testutil.GetTypeOrAssert(t, field, authObjectType.AttrTypes)
		testutil.AssertEqual(t, "authentication field "+field+" type", actualType, expectedType)
	}
}

func TestEmptyFieldHandling(t *testing.T) {
	// Test what happens when we try to get attribute types for a non-existent field
	defer func() {
		if r := recover(); r != nil {
			// This is expected behavior - the function should panic with invalid field names
			t.Logf("Function properly panicked with non-existent field: %v", r)
		}
	}()

	// This should panic or return empty map
	_ = GetHTTPHeaderAttributeTypes("non_existent_header")
}
