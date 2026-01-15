package schema

import (
	"context"
	"fmt"
	"sync"

	"catchpoint-provider/internal/fields"
	"catchpoint-provider/internal/labels"
	"catchpoint-provider/internal/modifier"
	cptypes "catchpoint-provider/internal/types"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var authenticationTypeValidator = oneOfStringValidator(cptypes.ValidAuthenticationTypeNames)

var (
	requestSettingsSchemaOnce sync.Once
	requestSettingsSchema     schema.SingleNestedBlock
)

// BuildRequestSettingsBlock builds the schema attributes for request settings, including certificates.
func BuildRequestSettingsBlock(ctx context.Context) map[string]schema.Block {
	return buildRequestSettingsBlock(ctx)
}

// GetAuthenticationAttributeTypes returns the attribute types for authentication settings.
func GetAuthenticationAttributeTypes() map[string]attr.Type {
	// Get the request settings schema directly
	requestSchema := getRequestSettingsSchema()

	// Get the authentication nested block from within request_settings
	if authBlock, exists := requestSchema.Blocks[fields.Authentication]; exists {
		if singleNestedBlock, ok := authBlock.(schema.SingleNestedBlock); ok {
			return ExtractAllAttributeTypes(
				singleNestedBlock.Attributes,
				singleNestedBlock.Blocks,
			)
		}
	}
	return make(map[string]attr.Type)
}

// GetHTTPRequestHeadersAttributeTypes returns the attribute types for all HTTP request headers.
func GetHTTPRequestHeadersAttributeTypes() map[string]attr.Type {
	requestSchema := getRequestSettingsSchema()
	// Get the http request headers nested block from within request_settings
	if httpHeadersBlock, exists := requestSchema.Blocks[fields.HTTPRequestHeaders]; exists {
		if singleNestedBlock, ok := httpHeadersBlock.(schema.SingleNestedBlock); ok {
			return ExtractAllAttributeTypes(
				singleNestedBlock.Attributes,
				singleNestedBlock.Blocks,
			)
		}
	}
	return make(map[string]attr.Type)
}

// GetRequestSettingsAttributeTypes returns the attribute types for request settings.
func GetRequestSettingsAttributeTypes() map[string]attr.Type {
	requestSchema := getRequestSettingsSchema()
	return ExtractAllAttributeTypes(
		requestSchema.Attributes,
		requestSchema.Blocks,
	)
}

// GetHTTPHeaderAttributeTypes returns the attribute types for a specific HTTP header type.
func GetHTTPHeaderAttributeTypes(headerName string) map[string]attr.Type {
	httpHeadersTypes := GetHTTPRequestHeadersAttributeTypes()

	headerType, exists := httpHeadersTypes[headerName]
	if !exists {
		return map[string]attr.Type{}
	}

	if objectType, ok := headerType.(types.ObjectType); ok {
		return objectType.AttrTypes
	}

	// This should never happen with current schema design; only if someone introduces a field
	// of an incorrect type.
	panic(fmt.Sprintf("Header field %s exists but is not ObjectType: %T", headerName, headerType))
}

// buildRequestSettingsBlock builds the schema attributes for request settings.
func buildRequestSettingsBlock(ctx context.Context) map[string]schema.Block {
	requestSettings := map[string]schema.Block{
		fields.RequestSettings: schema.SingleNestedBlock{
			Description: `Used for overriding the request settings section. 
			Note: omitting this field entirely means that the request settings for this object will not be tracked.
			If you specifically want the object to inherit the parent request settings then supply an object with only the request_setting_type
			field, e.g. '{ request_setting_type = "Inherit" }'.`,
			Attributes: MergeAttributes(
				buildRequestSettingChildAttributes(ctx),
			),
			Blocks: MergeBlocks(
				buildAuthenticationBlock(ctx),
				buildHTTPRequestHeaderBlock(),
			),
			PlanModifiers: []planmodifier.Object{
				modifier.InheritBlockPlanModifier(fields.RequestSettingType),
			},
		},
	}

	return requestSettings
}

func getRequestSettingsSchema() schema.SingleNestedBlock {
	requestSettingsSchemaOnce.Do(func() {
		attrs := BuildRequestSettingsBlock(context.Background())
		requestSettingsSchema = attrs[fields.RequestSettings].(schema.SingleNestedBlock)
	})
	return requestSettingsSchema
}

func buildRequestSettingChildAttributes(ctx context.Context) map[string]schema.Attribute {
	attributes := map[string]schema.Attribute{
		fields.RequestSettingType: schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "Set the request settings type, " + genericSettingTypeValidator.Description(ctx),
			Validators:  []validator.String{genericSettingTypeValidator},
		},
		fields.TokenIDs: schema.ListAttribute{
			Optional:    true,
			Description: "The list of Token IDs to use for the Test",
			//Sensitive:   true,
			ElementType: types.Int64Type,
		},
	}

	attributes[fields.LibraryCertificateIDs] = schema.ListAttribute{
		Optional:    true,
		Computed:    true,
		ElementType: types.Int64Type,
		Default:     listdefault.StaticValue(types.ListValueMust(types.Int64Type, []attr.Value{})),
		Description: "The list of Library Certificate IDs to use for the Test",
	}

	return attributes
}

func buildAuthenticationBlock(ctx context.Context) map[string]schema.Block {
	return map[string]schema.Block{
		fields.Authentication: schema.SingleNestedBlock{
			Description: "Used for overriding the authentication settings",
			Attributes: map[string]schema.Attribute{
				fields.AuthenticationType: schema.StringAttribute{
					Optional:    true,
					Computed:    true,
					Description: "The type of authentication to use, " + authenticationTypeValidator.Description(ctx),
					Validators:  []validator.String{authenticationTypeValidator},
				},
				fields.PasswordIDs: schema.ListAttribute{
					ElementType: types.Int64Type,
					Optional:    true,
					Sensitive:   true,
					Computed:    true,
					Description: "The list of password IDs to use for the Test",
				},
			},
		},
	}
}

func buildHTTPRequestHeaderBlock() map[string]schema.Block {
	return map[string]schema.Block{
		fields.HTTPRequestHeaders: schema.SingleNestedBlock{
			Description: "The HTTP header overrides to include in the request",
			Blocks: MergeBlocks(
				buildHTTPHeaderBlock(cptypes.UserAgentHeader, false, false),
				buildHTTPHeaderBlock(cptypes.AcceptHeader, false, false),
				buildHTTPHeaderBlock(cptypes.AcceptEncodingHeader, false, false),
				buildHTTPHeaderBlock(cptypes.AcceptLanguageHeader, false, false),
				buildHTTPHeaderBlock(cptypes.AcceptCharsetHeader, false, false),
				buildHTTPHeaderBlock(cptypes.CookieHeader, false, false),
				buildHTTPHeaderBlock(cptypes.CacheControlHeader, false, false),
				buildHTTPHeaderBlock(cptypes.ConnectionHeader, false, false),
				buildHTTPHeaderBlock(cptypes.PragmaHeader, false, false),
				buildHTTPHeaderBlock(cptypes.RefererHeader, false, false),
				buildHTTPHeaderBlock(cptypes.HostHeader, false, false),
				// DNS Override is only for child_host_pattern, so we set childHostPatternRequired to true.
				buildHTTPHeaderBlock(cptypes.DNSOverrideHeader, true, false),
				buildHTTPHeaderBlock(cptypes.RequestOverrideHeader, false, false),
				buildHTTPHeaderBlock(cptypes.RequestBlockHeader, false, false),
				buildHTTPHeaderBlock(cptypes.RequestDelayHeader, false, false),
				buildHTTPHeaderBlock(cptypes.DNSResolverOverrideHeader, false, false),
				// Both SNI and Custom also require a header_name to be set.
				buildHTTPHeaderBlock(cptypes.SNIOverrideHeader, false, true),
				buildHTTPHeaderBlock(cptypes.CustomHeader, false, true),
			),
		},
	}
}

// buildHTTPHeaderBlock builds a single HTTP header block with the specified name.
// Note: We don't *actually* require fields because this is a Block. By enabling 'Required' attribute, it would force
// the user to always specify every single header. However, if we used an Attribute instead of a block, then the
// .tf format would change and every user would have to reformat their existing configurations.
func buildHTTPHeaderBlock(headerName string, childHostPatternRequired, includeHeaderName bool) map[string]schema.Block {
	attributes := map[string]schema.Attribute{
		fields.Value: schema.StringAttribute{
			Optional: true,
			Computed: true,
			// Some headers (e.g. Request Block) allow empty values, so we set the default to empty string.
			Default:     stringdefault.StaticString(cptypes.EmptyString),
			Description: "The value of the HTTP header",
		},
		fields.ChildHostPattern: schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Default:     stringdefault.StaticString(cptypes.EmptyString),
			Description: "The child host pattern to match for the HTTP header",
		},
	}

	if includeHeaderName {
		if fields.HeaderName == cptypes.SNIOverrideHeader {
			attributes[fields.HeaderName] = schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The name of the HTTP header",
				Default:     stringdefault.StaticString(labels.SNIOverride),
			}
		} else {
			attributes[fields.HeaderName] = schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The name of the HTTP header",
			}
		}
	}

	var headerDescription string
	if childHostPatternRequired {
		headerDescription = fmt.Sprintf("Sets the '%s' header for the given child_host_pattern", headerName)
	} else {
		headerDescription = fmt.Sprintf("Sets the '%s' header for the test URL or child_host_pattern (if not omitted)", headerName)
	}

	return map[string]schema.Block{
		headerName: schema.SingleNestedBlock{
			Description: headerDescription,
			Attributes:  attributes,
		},
	}
}
