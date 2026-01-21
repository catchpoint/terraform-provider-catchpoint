package transform

import (
	"catchpoint-provider/internal/fields"
	"catchpoint-provider/internal/labels"
	"catchpoint-provider/internal/logger"
	"catchpoint-provider/internal/models"
	"catchpoint-provider/internal/models/resource"
	cpschema "catchpoint-provider/internal/schema"
	cptypes "catchpoint-provider/internal/types"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func JSONToTerraformRequestSettings(requestSetting *models.RequestSettingsJSON) (resource resource.RequestSettingsModel, diags diag.Diagnostics) {
	resource.RequestSettings, diags = jsonToTerraformRequestSettings(requestSetting)
	return
}

func jsonToTerraformRequestSettings(requestSetting *models.RequestSettingsJSON) (obj types.Object, diags diag.Diagnostics) {
	attrsType := cpschema.GetRequestSettingsAttributeTypes()
	obj = types.ObjectNull(attrsType)

	if requestSetting == nil {
		logger.WarnBG("RequestSettings is nil, returning null object")
		return
	}

	requestSettingTypeName, ok := cptypes.GetGenericSettingTypeName(requestSetting.RequestSettingType.ID)
	if !ok {
		logger.WarnBG("Invalid RequestSettingType ID: %d", requestSetting.RequestSettingType.ID)
		diags.Append(diag.NewWarningDiagnostic("Invalid RequestSettingType ID", "The provided RequestSettingType ID is not recognized."))
	}

	authObj, authDiags := jsonToTerraformAuthentication(requestSetting.Authentication)
	diags.Append(authDiags...)
	if diags.HasError() {
		return
	}

	tokenList := types.ListNull(types.Int64Type)
	if requestSetting.TokenIDs != nil {
		var tokenDiags diag.Diagnostics
		tokenList, tokenDiags = IntSliceToTerraformList(*requestSetting.TokenIDs)
		diags.Append(tokenDiags...)
		if diags.HasError() {
			return
		}
	}

	headerObj, headerDiags := jsonToTerraformHTTPRequestHeaders(requestSetting.HTTPHeaderRequests)
	diags.Append(headerDiags...)
	if diags.HasError() {
		return
	}

	attrs := map[string]attr.Value{
		fields.RequestSettingType: types.StringValue(requestSettingTypeName),
		fields.Authentication:     authObj,
		fields.TokenIDs:           tokenList,
		fields.HTTPRequestHeaders: headerObj,
	}

	diags.Append(handleCertificates(attrs, requestSetting)...)
	if diags.HasError() {
		return
	}

	obj, objDiags := types.ObjectValue(attrsType, attrs)
	diags.Append(objDiags...)
	return
}

func jsonToTerraformAuthentication(auth *models.AuthenticationJSON) (obj types.Object, diags diag.Diagnostics) {
	obj = types.ObjectNull(cpschema.GetAuthenticationAttributeTypes())

	if auth == nil {
		return
	}

	var passwordList = types.ListNull(types.Int64Type)
	if auth.PasswordIDs != nil {
		var pwDiags diag.Diagnostics
		passwordList, pwDiags = IntSliceToTerraformList(*auth.PasswordIDs)
		diags.Append(pwDiags...)
		if diags.HasError() {
			return
		}
	}

	var authTypeName string
	if auth.AuthenticationMethodType != nil && auth.AuthenticationMethodType.ID != nil {
		var ok bool
		authTypeName, ok = cptypes.GetAuthenticationTypeName(*auth.AuthenticationMethodType.ID)
		if !ok {
			logger.ErrorBG("Unknown Authentication Type ID: %d, Name: '%s'", *auth.AuthenticationMethodType.ID, *auth.AuthenticationMethodType.Name)
		}
	} else {
		authTypeName = cptypes.None // Default to "none" if not specified
	}

	attrs := map[string]attr.Value{
		fields.AuthenticationType: types.StringValue(authTypeName),
		fields.PasswordIDs:        passwordList,
	}

	obj, objDiags := types.ObjectValue(cpschema.GetAuthenticationAttributeTypes(), attrs)
	diags.Append(objDiags...)

	return
}

func jsonToTerraformHTTPRequestHeaders(headers *[]models.HTTPHeaderRequestJSON) (obj types.Object, diags diag.Diagnostics) {
	attrsType := cpschema.GetHTTPRequestHeadersAttributeTypes()
	obj = types.ObjectNull(attrsType)

	if headers == nil {
		return
	}

	headerMap, mapDiags := mapHeadersByType(headers)
	diags.Append(mapDiags...)
	if diags.HasError() {
		return
	}

	attrs := make(map[string]attr.Value)
	for headerName := range attrsType {
		if header, exists := headerMap[headerName]; exists {
			headerObj, headerDiags := jsonToTerraformHTTPRequestHeader(header, headerName)
			diags.Append(headerDiags...)
			if diags.HasError() {
				return
			}
			attrs[headerName] = headerObj
		} else {
			attrs[headerName] = types.ObjectNull(cpschema.GetHTTPHeaderAttributeTypes(headerName))
		}
	}

	obj, objDiags := types.ObjectValue(attrsType, attrs)
	diags.Append(objDiags...)

	return
}

// mapHeadersByType maps headers to their schema names and returns diagnostics for unknown types.
func mapHeadersByType(headers *[]models.HTTPHeaderRequestJSON) (map[string]models.HTTPHeaderRequestJSON, diag.Diagnostics) {
	headerMap := make(map[string]models.HTTPHeaderRequestJSON)
	var diags diag.Diagnostics

	for _, header := range *headers {
		var headerName string

		// Special handling for "Sni-Override" header. We give the user the ability to set "Sni-Override" but
		// it's actually just a type 11 (Custom) header with a fixed name. So we need to map it here.
		if header.HeaderName != nil && *header.HeaderName == labels.SNIOverride {
			headerMap[cptypes.SNIOverrideHeader] = header
			continue
		}

		headerName, ok := cptypes.GetReqHeaderTypeName(header.RequestHeaderType.ID)
		if !ok {
			diags.AddError("Unknown Request Header Type", fmt.Sprintf("Unknown Request Header Type ID: %d, Name: '%s'", header.RequestHeaderType.ID, header.RequestHeaderType.Name))
			continue
		}

		headerMap[headerName] = header
	}
	return headerMap, diags
}

func jsonToTerraformHTTPRequestHeader(header models.HTTPHeaderRequestJSON, schemaHeaderName string) (obj types.Object, diags diag.Diagnostics) {
	attrTypes := cpschema.GetHTTPHeaderAttributeTypes(schemaHeaderName)
	attrs := make(map[string]attr.Value)

	attrs[fields.Value] = types.StringValue(header.RequestValue)

	SetStringOrNull(attrs, attrTypes, fields.ChildHostPattern, header.ChildHostPattern)

	// Always set header_name for sni_override
	if schemaHeaderName == cptypes.SNIOverrideHeader {
		attrs[fields.HeaderName] = types.StringValue("Sni-Override")
	} else {
		SetStringOrNull(attrs, attrTypes, fields.HeaderName, header.HeaderName)
	}

	obj, objDiags := types.ObjectValue(attrTypes, attrs)
	diags.Append(objDiags...)
	return
}

func handleCertificates(attrs map[string]attr.Value, requestSetting *models.RequestSettingsJSON) (diags diag.Diagnostics) {
	if requestSetting.LibraryCertificateIDs != nil && len(*requestSetting.LibraryCertificateIDs) > 0 {
		certIDs, certDiags := IntSliceToTerraformList(*requestSetting.LibraryCertificateIDs)
		diags.Append(certDiags...)
		if diags.HasError() {
			return diags
		}
		attrs[fields.LibraryCertificateIDs] = certIDs
	} else {
		attrs[fields.LibraryCertificateIDs] = types.ListValueMust(types.Int64Type, []attr.Value{})
	}
	return
}
