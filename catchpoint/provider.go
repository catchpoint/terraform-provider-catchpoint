package catchpoint

import (
	"context"
	"os"
	"regexp"
	"strconv"

	"catchpoint-provider/internal"
	"catchpoint-provider/internal/fields"
	cptypes "catchpoint-provider/internal/types"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure CatchpointProvider satisfies various provider interfaces.
var _ provider.Provider = &CatchpointProvider{}

// CatchpointProvider defines the provider implementation.
type CatchpointProvider struct {
	// version is set to the provider version on release, "dev" when the
	// provider is built and ran locally, and "test" when running acceptance
	// tests.
	version string
}

// CatchpointProviderModel describes the provider data model.
type CatchpointProviderModel struct {
	APIToken    types.String `tfsdk:"api_token"`
	APIVersion  types.String `tfsdk:"api_version"`
	LogJSON     types.String `tfsdk:"log_json"`
	Environment types.String `tfsdk:"catchpoint_environment"`
}

func (p *CatchpointProvider) Metadata(ctx context.Context, req provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "catchpoint"
	resp.Version = p.version
}

func (p *CatchpointProvider) Schema(ctx context.Context, req provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "The Catchpoint provider allows you to manage Catchpoint resources.",
		Attributes: map[string]schema.Attribute{
			fields.APIToken: schema.StringAttribute{
				Description: "The Catchpoint API token. Can also be set via the `CATCHPOINT_API_TOKEN` environment variable.",
				// Optional to allow env var usage. 'Configure' will validate that one is set one way or another.
				Optional:  true,
				Sensitive: true,
			},
			fields.APIVersion: schema.StringAttribute{
				Description: "The Catchpoint API version to use. Defaults to v4.0. Can also be set via the `CATCHPOINT_API_VERSION` environment variable.",
				Optional:    true,
				Validators: []validator.String{
					&apiVersionValidator{},
				},
			},
			fields.LogJSON: schema.StringAttribute{
				Description: "Enable or disable test json payload logging for debugging. Accepts string and converts to bool using ParseBool function. Can also be set via the `LOG_JSON` environment variable.",
				Optional:    true,
			},
			fields.Environment: schema.StringAttribute{
				Description: "Set the environment to stage, qa or prod. This is for internal use. Defaults to prod if empty. Can also be set via the `CATCHPOINT_ENVIRONMENT` environment variable.",
				Optional:    true,
			},
		},
	}
}

func (p *CatchpointProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	var config CatchpointProviderModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)

	if resp.Diagnostics.HasError() {
		return
	}

	apiToken := resolveAPIToken(config.APIToken.ValueString())
	if apiToken == cptypes.EmptyString {
		resp.Diagnostics.AddError(
			"Missing API Token",
			"The provider cannot create the Catchpoint API client as there is a missing or empty value for the Catchpoint API token. "+
				"Set the api_token value in the configuration or use the CATCHPOINT_API_TOKEN environment variable. "+
				"If either is already set, ensure the value is not empty.",
		)
		return
	}

	apiVersion := resolveAPIVersion(config.APIVersion.ValueString())
	logJSON := resolveLogJSON(config.LogJSON.ValueString())
	isLogJSON, err := strconv.ParseBool(logJSON)
	if err != nil || logJSON == cptypes.EmptyString {
		isLogJSON = false
	}
	environment := resolveEnvironment(config.Environment.ValueString())

	// Configure the internal client
	internal.SetEnvironment(environment, apiVersion)
	client := internal.NewConfig(apiToken, isLogJSON, environment)

	// Make the client available during DataSource and Resource type Configure methods.
	resp.DataSourceData = client
	resp.ResourceData = client
}

func resolveAPIToken(configValue string) string {
	if configValue != cptypes.EmptyString {
		return configValue
	}
	if token := os.Getenv("CATCHPOINT_API_TOKEN"); token != cptypes.EmptyString {
		return token
	}
	return cptypes.EmptyString
}

func resolveAPIVersion(configValue string) string {
	if configValue != cptypes.EmptyString {
		return configValue
	}
	if version := os.Getenv("CATCHPOINT_API_VERSION"); version != cptypes.EmptyString {
		return version
	}
	return "v4.0"
}

func resolveLogJSON(configValue string) string {
	if configValue != cptypes.EmptyString {
		return configValue
	}
	if log := os.Getenv("LOG_JSON"); log != cptypes.EmptyString {
		return log
	}
	return cptypes.EmptyString
}

func resolveEnvironment(configValue string) string {
	if configValue != cptypes.EmptyString {
		return configValue
	}
	if env := os.Getenv("CATCHPOINT_ENVIRONMENT"); env != cptypes.EmptyString {
		return env
	}
	return cptypes.EmptyString
}

func (p *CatchpointProvider) Resources(ctx context.Context) []func() resource.Resource {
	return listResources()
}

func (p *CatchpointProvider) DataSources(ctx context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{
		// Add data sources here when implemented
	}
}

func New(version string) func() provider.Provider {
	return func() provider.Provider {
		return &CatchpointProvider{
			version: version,
		}
	}
}

// apiVersionValidator validates the API version format
type apiVersionValidator struct{}

func (v *apiVersionValidator) Description(ctx context.Context) string {
	return "API version must be in the format 'vX.Y' where X and Y are integers"
}

func (v *apiVersionValidator) MarkdownDescription(ctx context.Context) string {
	return "API version must be in the format `vX.Y` where X and Y are integers"
}

func (v *apiVersionValidator) ValidateString(ctx context.Context, req validator.StringRequest, resp *validator.StringResponse) {
	if req.ConfigValue.IsNull() || req.ConfigValue.IsUnknown() {
		return
	}

	version := req.ConfigValue.ValueString()
	regex := regexp.MustCompile(`^v[0-9]+\.[0-9]+$`)
	if !regex.MatchString(version) {
		resp.Diagnostics.AddAttributeError(
			req.Path,
			"Invalid API Version Format",
			"API version must be in the format 'vX.Y' where X and Y are integers. "+
				"For example: 'v1.0', 'v2.1', 'v4.0'.",
		)
	}
}
