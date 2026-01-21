package testmonitor

import (
	"context"
	"regexp"
	"slices"

	"catchpoint-provider/internal/fields"
	cpschema "catchpoint-provider/internal/schema"
	cptypes "catchpoint-provider/internal/types"

	"github.com/hashicorp/terraform-plugin-framework-validators/float64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var (
	chromeVersionValidator = oneOfStringValidator(cptypes.ValidChromeVersions)
	simulateValidator      = oneOfStringValidator(cptypes.ValidUserAgentTypes)
)

func buildMonitorAttribute(ctx context.Context, required bool, defaultValue string, possibleValues ...string) map[string]schema.Attribute {
	var MonitorValidator = stringvalidator.OneOf(possibleValues...)

	attribute := schema.StringAttribute{
		Optional:    !required,
		Required:    required,
		Computed:    !required,
		Description: "Set the monitor type, " + MonitorValidator.Description(ctx),
		Validators:  []validator.String{MonitorValidator},
	}

	// Only set the default value if the value is not empty.
	if defaultValue != cptypes.EmptyString {
		attribute.Default = stringdefault.StaticString(defaultValue)
	}

	return map[string]schema.Attribute{fields.Monitor: attribute}
}

func buildLabelsBlock() map[string]schema.Block {
	return map[string]schema.Block{
		fields.Label: schema.ListNestedBlock{
			Description: "Labels with key,[values] pairs",
			NestedObject: schema.NestedBlockObject{
				Attributes: map[string]schema.Attribute{
					"key": schema.StringAttribute{
						Required:    true,
						Description: "Label key",
					},
					"values": schema.ListAttribute{
						ElementType: types.StringType,
						Required:    true,
						Description: "Label values",
					},
				},
			},
		},
	}
}

func buildThresholdsBlock() map[string]schema.Block {
	return map[string]schema.Block{
		fields.Thresholds: schema.SingleNestedBlock{
			Description: "Test thresholds for test time and availability percentage",
			Attributes: map[string]schema.Attribute{
				fields.TestTimeWarning: schema.Float64Attribute{
					Optional:    true,
					Description: "The test time warning threshold",
					Computed:    true,
				},
				fields.TestTimeCritical: schema.Float64Attribute{
					Optional:    true,
					Description: "The test time critical threshold",
					Computed:    true,
				},
				fields.AvailabilityWarning: schema.Float64Attribute{
					Optional:    true,
					Description: "The availability warning threshold",
					Computed:    true,
					Validators:  []validator.Float64{float64validator.Between(0, 100)},
				},
				fields.AvailabilityCritical: schema.Float64Attribute{
					Optional:    true,
					Description: "The availability critical threshold",
					Computed:    true,
					Validators:  []validator.Float64{float64validator.Between(0, 100)},
				},
			},
		},
	}
}

func buildGatewayAddressOrHostBlock() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		fields.GatewayAddressOrHost: schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Default:     stringdefault.StaticString(cptypes.EmptyString),
			Description: "Host/IP to use for network troubleshooting and monitoring",
		},
	}
}

func buildTestRequestDataAttributes(ctx context.Context, testType cptypes.TestType) map[string]schema.Attribute {
	attributes := map[string]schema.Attribute{
		fields.TestScript: schema.StringAttribute{
			Required:    true,
			Description: "The Script that will simulate user workflow",
		},
	}

	// For API tests, the user must specify either selenium or javascript.
	if testType == cptypes.APIType {
		stringvalidator := stringvalidator.OneOf("selenium", "javascript")
		attributes[fields.TestScriptType] = schema.StringAttribute{
			Required:    true,
			Description: "The type of script, " + stringvalidator.Description(ctx),
			Validators:  []validator.String{stringvalidator},
		}
	}

	// For PW/Puppeteer, there is only one valid value so we compute and add it ourselves.
	if testType == cptypes.PlaywrightType || testType == cptypes.PuppeteerType {
		attributes[fields.TestScriptType] = schema.StringAttribute{
			Computed:    true,
			Optional:    false,
			Required:    false,
			Description: "The type of script.",
			Default:     stringdefault.StaticString(cptypes.GetTestTypeName(testType)),
		}
	}

	// For Transaction tests, the only acceptable value is selenium, so we compute and add it ourselves.
	if testType == cptypes.TransactionType {
		attributes[fields.TestScriptType] = schema.StringAttribute{
			Computed:    true,
			Optional:    false,
			Required:    false,
			Description: "The type of script.",
			Default:     stringdefault.StaticString("selenium"),
		}
	}

	return attributes
}

func buildURLField(ctx context.Context, testType cptypes.TestType) map[string]schema.Attribute {
	switch testType {
	case cptypes.BGPType:
		return map[string]schema.Attribute{
			fields.Prefix: schema.StringAttribute{
				Required:    true,
				Description: "IPV4 address with a netmask range from 8 to 24 or IPV6 address with a netmask range from 28 to 128",
			},
		}
	case cptypes.DNSType:
		return map[string]schema.Attribute{
			fields.TestDomain: schema.StringAttribute{
				Required:    true,
				Description: "The domain to be tested. Example: www.catchpoint.com",
			},
		}
	case cptypes.SSLType, cptypes.PingType, cptypes.TracerouteType:
		desc := "The domain or IP to be tested. Example: catchpoint.com"
		validators := []validator.String{
			stringvalidator.LengthAtLeast(1),
		}

		if testType == cptypes.SSLType {
			regex := regexp.MustCompile(`^ssl:\/\/[a-zA-Z0-9.-]+(:[0-9]+)?`)
			validator := stringvalidator.RegexMatches(regex, "Must start with 'ssl://' followed by a valid domain or IP address, and an optional port number")
			validators = append(validators, validator)
			desc = "The SSL domain to be tested. Example: ssl://www.catchpoint.com. " + validator.Description(ctx)
		}

		return map[string]schema.Attribute{
			fields.TestLocation: schema.StringAttribute{
				Required:    true,
				Description: desc,
				Validators:  validators,
			},
		}
	case cptypes.WebType:
		return map[string]schema.Attribute{
			fields.TestURL: schema.StringAttribute{
				Required:    true,
				Description: "The URL to be tested. Example: https://www.catchpoint.com",
			},
		}
	default:
		return nil
	}
}

// Builds the ChromeVersion and Simulate attributes for Web and Transaction tests.
// Both are validated elsewhere as their presence depends on the monitor specified.
func buildChromeSimulateAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		fields.ChromeVersion: schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "Set the Chrome version to simulate, " + chromeVersionValidator.Description(context.Background()),
			Validators:  []validator.String{chromeVersionValidator},
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.UseStateForUnknown(),
			},
		},
		fields.Simulate: schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "The device to simulate for mobile, mobile playback(playback source) monitors, " + simulateValidator.Description(context.Background()),
			Validators:  []validator.String{simulateValidator},
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.UseStateForUnknown(),
			},
		},
	}
}

// Shared attributes for all test resources
func buildCommonTestAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// Required fields:
		fields.DivisionID: schema.Int64Attribute{
			Required:    true,
			Description: "The Division where the Test will be created",
		},
		fields.ProductID: schema.Int64Attribute{
			Required:    true,
			Description: "The parent Product under which the Test will be created",
			PlanModifiers: []planmodifier.Int64{
				int64planmodifier.UseStateForUnknown(),
			},
		},
		fields.TestName: schema.StringAttribute{
			Required:    true,
			Description: "The name of the Test",
		},
		// Optional fields:
		fields.FolderID: schema.Int64Attribute{
			Optional:    true,
			Description: "The Folder under which the Test will be created",
			PlanModifiers: []planmodifier.Int64{
				int64planmodifier.UseStateForUnknown(),
			},
		},
		fields.TestDescription: schema.StringAttribute{
			Optional:    true,
			Description: "The Test description",
		},
		// Optional fields with defaults:
		fields.Status: schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Default:     stringdefault.StaticString(cptypes.Active),
			Description: "Set the Test status, " + cpschema.StatusValidator.Description(ctx),
			Validators:  []validator.String{cpschema.StatusValidator},
		},
		fields.AlertsPaused: schema.BoolAttribute{
			Optional:    true,
			Computed:    true,
			Default:     booldefault.StaticBool(false),
			Description: "Switch for pausing Test alerts. Default: false",
		},
		fields.StartTime: schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "Start time for the Test in ISO format like 2027-12-30T04:59:00Z",
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.UseStateForUnknown(),
			},
		},
		fields.EndTime: schema.StringAttribute{
			// Whether this is required or not depends on the Client settings. But we can't check those
			// so we need to make it optional for everyone and let the API return an error if it's required.
			Optional:    true,
			Computed:    true,
			Default:     stringdefault.StaticString(cptypes.EmptyString),
			Description: "End time for the Test in ISO format like 2027-12-30T04:59:00Z",
		},
		fields.EnableTestDataWebhook: schema.BoolAttribute{
			Optional:    true,
			Computed:    true,
			Default:     booldefault.StaticBool(true),
			Description: "Switch for enabling test data webhook feature. Default: true",
		},
		// Optional fields that can not be set:
		fields.ID: schema.Int64Attribute{
			Required:    false,
			Optional:    false,
			Computed:    true,
			Description: "The unique identifier for the Test. This is only set by the API and should be 0 for a new test.",
		},
	}
}

// OneOfStringValidator returns a string validator built from the provided values.
// It copies and sorts the slice so the produced description is deterministic.
func oneOfStringValidator(values []string) validator.String {
	copied := slices.Clone(values)
	slices.Sort(copied)
	return stringvalidator.OneOf(copied...)
}
