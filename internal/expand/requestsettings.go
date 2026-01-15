package expand

import (
	"catchpoint-provider/internal/fields"
	"catchpoint-provider/internal/models"
	cptypes "catchpoint-provider/internal/types"
	"catchpoint-provider/internal/validation"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// ExpandRequestSettingsConfig expands the plan Object for requestSettings into the Configuration object.
func ExpandRequestSettingsConfig(requestSetting types.Object, config *models.RequestSettingsConfig) (diags diag.Diagnostics) {
	attrs := requestSetting.Attributes()

	if isBlockEmpty(requestSetting) ||
		(attrExistsAndNotNull(attrs, fields.RequestSettingType) && attrs[fields.RequestSettingType].Equal(types.StringValue(cptypes.Inherit))) {
		config.RequestSettingType = validation.GetGenericSettingTypeOrDefault(cptypes.Inherit)
		config.AuthenticationType = validation.GetAuthenticationTypeOrDefault(cptypes.None)
		config.PasswordIDs = []int{}
		config.TokenIDs = []int{}
		config.CertificateIDs = []int{}
		return
	}

	// If the user has specified these settings, we can assume that they're trying to override the defaults.
	config.RequestSettingType = validation.GetGenericSettingTypeOrDefault(cptypes.Override)

	if attrExistsAndNotNull(attrs, fields.Authentication) {
		authenticationObject := attrs[fields.Authentication].(types.Object)

		setDiag := expandConfigAuthentication(config, authenticationObject)
		diags.Append(setDiag...)
		if diags.HasError() {
			return
		}
	} else {
		// If Authentication is not specified, set it to the default "none" type.
		config.AuthenticationType = validation.GetAuthenticationTypeOrDefault(cptypes.None)
		config.PasswordIDs = []int{} // Default to empty slice
		config.TokenIDs = []int{}    // Default to empty slice
	}

	setDiag := expandIntListSettingFromAttrs(attrs, fields.TokenIDs, func(val []int) {
		config.TokenIDs = val
	})
	diags.Append(setDiag...)
	if diags.HasError() {
		return
	}

	setDiag = expandIntListSettingFromAttrs(attrs, fields.LibraryCertificateIDs, func(val []int) {
		config.CertificateIDs = val
	})
	diags.Append(setDiag...)
	if diags.HasError() {
		return
	}

	if attrExistsAndNotNull(attrs, fields.HTTPRequestHeaders) {
		httpRequestHeadersObject := attrs[fields.HTTPRequestHeaders].(types.Object)
		setDiag = expandConfigHTTPRequestHeaders(config, httpRequestHeadersObject)
		diags.Append(setDiag...)
		if diags.HasError() {
			return
		}
	} else {
		// If HTTPRequestHeaders is not specified, set it to an empty slice.
		config.TestHTTPHeaderRequests = []models.TestHTTPHeaderRequestConfig{}
	}

	return
}

func expandConfigAuthentication(config *models.RequestSettingsConfig, authSettings types.Object) (diags diag.Diagnostics) {
	attrs := authSettings.Attributes()

	if attrExistsAndNotNull(attrs, fields.AuthenticationType) {
		authenticationTypeStr := attrs[fields.AuthenticationType].(types.String)
		config.AuthenticationType = validation.GetAuthenticationTypeOrDefault(authenticationTypeStr.ValueString())
	} else {
		// If AuthenticationType is not specified, set it to the default "none" type.
		config.AuthenticationType = validation.GetAuthenticationTypeOrDefault(cptypes.None)
	}

	setDiags := expandIntListSettingFromAttrs(attrs, fields.PasswordIDs, func(val []int) {
		config.PasswordIDs = val
	})
	diags.Append(setDiags...)
	return
}

func expandConfigHTTPRequestHeaders(config *models.RequestSettingsConfig, requestSetting types.Object) (diags diag.Diagnostics) {
	attrs := requestSetting.Attributes()

	for headerKey, attrVal := range attrs {
		requestHeaderObj, ok := attrVal.(types.Object)
		if !ok || requestHeaderObj.IsNull() || requestHeaderObj.IsUnknown() {
			continue
		}

		headerDiags, header := expandSingleHTTPRequestHeader(headerKey, requestHeaderObj)
		diags.Append(headerDiags...)
		if header != nil {
			config.TestHTTPHeaderRequests = append(config.TestHTTPHeaderRequests, *header)
		}
	}

	return
}

// expandSingleHTTPRequestHeader extracts a single HTTP request header config.
func expandSingleHTTPRequestHeader(requestHeader string, requestHeaderObj types.Object) (diags diag.Diagnostics, header *models.TestHTTPHeaderRequestConfig) {
	requestHeaderAttrs := requestHeaderObj.Attributes()

	headerTypeID, ok := cptypes.GetReqHeaderTypeID(requestHeader)
	if !ok {
		diags.Append(diag.NewErrorDiagnostic(
			"Invalid Request Header Type",
			"Unknown Request Header Type: "+requestHeader,
		))
		return diags, nil
	}

	h := models.TestHTTPHeaderRequestConfig{
		RequestHeaderType: models.IDNameConfig{
			ID:   headerTypeID,
			Name: requestHeader,
		},
	}

	expandStringSettingFromAttrs(requestHeaderAttrs, fields.Value, func(val string) {
		h.RequestValue = val
	})
	expandStringSettingFromAttrs(requestHeaderAttrs, fields.ChildHostPattern, func(val string) {
		h.ChildHostPattern = val
	})

	// Sni-Override is a special case. It's actually a Custom header type (ID=11) in the backend,
	// where the headerName has been set to "Sni-Override". But we present it to terraform users as a type of header.
	// So if the type is Sni-Override, we need to set the header name to "Sni-Override" regardless of what the user put in the header_name field
	// (which is likely empty). For all other types, we take the header_name field as-is.
	if headerTypeID == 11 && requestHeader == cptypes.SNIOverrideHeader {
		h.HeaderName = "Sni-Override"
	} else {
		expandStringSettingFromAttrs(requestHeaderAttrs, fields.HeaderName, func(val string) {
			h.HeaderName = val
		})
	}

	return diags, &h
}
