package schema

import (
	"catchpoint-provider/internal/fields"
	"catchpoint-provider/internal/types"
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var (
	additionalMonitorTypeValidator     = oneOfStringValidator(types.ValidAdditionalMonitorTypeNames)
	bandwidthThrottlingTypeValidator   = oneOfStringValidator(types.ValidBandwidthThrottlingTypeNames)
	failureHopCountMaxValidator        = int64validator.AtMost(20)                            // The portal allows 7 to 20.
	failureHopCountMinValidator        = int64validator.AtLeast(7)                            // The portal allows 7 to 20.
	maxStepRuntimeSecOverrideValidator = int64validator.OneOf(5, 10, 15, 20, 30, 60, 90, 120) // The portal has a drop-down for these values.
	pingCountMaxValidator              = int64validator.AtMost(20)                            // The portal allows 4 to 20.
	pingCountMinValidator              = int64validator.AtLeast(4)                            // The portal allows 4 to 20.
	viewportHeightMaxValidator         = int64validator.AtMost(2000)                          // This is the most the portal allows you to specify.
	viewportHeightMinValidator         = int64validator.AtLeast(300)                          // This is the least that the portal allows you to specify.
	viewportWidthMaxValidator          = int64validator.AtMost(2000)                          // This is the most the portal allows you to specify.
	viewportWidthMinValidator          = int64validator.AtLeast(300)                          // This is the least that the portal allows you to specify.
	waitForNoActivityValidator         = int64validator.OneOf(0, 500, 1000, 1500, 2000, 2500, 3000, 3500, 4000, 4500, 5000)
)

// All advanced settings for everything. This lookup table allows any given advanced setting to be added
// to a schema.
var advancedSettingAttributes = map[string]func(ctx context.Context) schema.Attribute{
	fields.AdditionalMonitor: func(ctx context.Context) schema.Attribute {
		return schema.StringAttribute{
			Optional:    true,
			Description: "Set the additional monitor to run along with the test monitor, " + additionalMonitorTypeValidator.Description(ctx),
			Validators:  []validator.String{additionalMonitorTypeValidator},
		}
	},
	types.AllowTestDownloadLimitOverride: func(ctx context.Context) schema.Attribute {
		return schema.BoolAttribute{
			Optional:    true,
			Description: "True enables test download limit override setting",
		}
	},
	fields.BandwidthThrottling: func(ctx context.Context) schema.Attribute {
		return schema.StringAttribute{
			Optional:    true,
			Description: "Set the bandwidth throttling for chrome, " + bandwidthThrottlingTypeValidator.Description(ctx),
			Validators:  []validator.String{bandwidthThrottlingTypeValidator},
		}
	},
	types.CaptureFilmstrip: func(ctx context.Context) schema.Attribute {
		return schema.BoolAttribute{
			Optional:    true,
			Description: "True enables capture filmstrip setting",
		}
	},
	types.CaptureHTTPHeaders: func(ctx context.Context) schema.Attribute {
		return schema.BoolAttribute{
			Optional:    true,
			Description: "True enables capture http headers setting for all runs",
		}
	},
	types.CaptureResponseContent: func(ctx context.Context) schema.Attribute {
		return schema.BoolAttribute{
			Optional:    true,
			Description: "True enables capture response content setting for all runs",
		}
	},
	types.CaptureScreenshot: func(ctx context.Context) schema.Attribute {
		return schema.BoolAttribute{
			Optional:    true,
			Description: "True enables capture screenshot setting for all runs",
		}
	},
	types.CertificateRevocationDisabled: func(ctx context.Context) schema.Attribute {
		return schema.BoolAttribute{
			Optional:    true,
			Description: "True disables certificate revocation",
		}
	},
	types.DebugPrimaryHostOnFailure: func(ctx context.Context) schema.Attribute {
		return schema.BoolAttribute{
			Optional:    true,
			Description: "True enables debug primary host on failure setting",
		}
	},
	types.DebugReferencedHostsOnFailure: func(ctx context.Context) schema.Attribute {
		return schema.BoolAttribute{
			Optional:    true,
			Description: "True enables debug referenced hosts on failure setting",
		}
	},
	types.DisableCrossOriginIframeAccess: func(ctx context.Context) schema.Attribute {
		return schema.BoolAttribute{
			Optional:    true,
			Description: "True enables disable cross origin iframe access setting for chrome monitor",
		}
	},
	types.DisableRecursiveResolution: func(ctx context.Context) schema.Attribute {
		return schema.BoolAttribute{
			Optional:    true,
			Description: "True enables disable recursive resolution setting",
		}
	},
	fields.EDNSSubnet: func(ctx context.Context) schema.Attribute {
		return schema.StringAttribute{
			Optional:    true,
			Description: "Set the EDNS subnet for the test",
		}
	},
	types.EnableBindHostname: func(ctx context.Context) schema.Attribute {
		return schema.BoolAttribute{
			Optional:    true,
			Description: "True enables bind hostname setting",
		}
	},
	types.EnableDNSSEC: func(ctx context.Context) schema.Attribute {
		return schema.BoolAttribute{
			Optional:    true,
			Description: "True enables DNSSEC setting",
		}
	},
	types.EnableNSID: func(ctx context.Context) schema.Attribute {
		return schema.BoolAttribute{
			Optional:    true,
			Description: "True enables NSID setting",
		}
	},
	types.EnablePathMTUDiscovery: func(ctx context.Context) schema.Attribute {
		return schema.BoolAttribute{
			Optional:    true,
			Description: "True enables Path MTU Discovery",
		}
	},
	types.EnableSelfVersusThirdPartyZones: func(ctx context.Context) schema.Attribute {
		return schema.BoolAttribute{
			Optional:    true,
			Description: "True enables self versus third party zones setting and matches self zone by test URL",
		}
	},
	types.EnableTCPProtocol: func(ctx context.Context) schema.Attribute {
		return schema.BoolAttribute{
			Optional:    true,
			Description: "True enables TCP protocol setting",
		}
	},
	fields.EnforceTestFailureIfRunsLongerThan: func(ctx context.Context) schema.Attribute {
		return schema.Int64Attribute{
			Optional:    true,
			Description: "Set the time value in seconds post which the test will be marked as failure, " + maxStepRuntimeSecOverrideValidator.Description(ctx),
			Validators:  []validator.Int64{maxStepRuntimeSecOverrideValidator},
		}
	},
	types.F40xOr50xHTTPMarkSuccessful: func(ctx context.Context) schema.Attribute {
		return schema.BoolAttribute{
			Optional:    true,
			Description: "True enables 40x or 50x error mark successful setting",
		}
	},
	types.HostDataCollectionEnabled: func(ctx context.Context) schema.Attribute {
		return schema.BoolAttribute{
			Optional:    true,
			Description: "True enables host data collection setting",
		}
	},
	types.IgnoreSSLFailures: func(ctx context.Context) schema.Attribute {
		return schema.BoolAttribute{
			Optional:    true,
			Description: "True enables ignore SSL failures setting",
		}
	},
	fields.PingCount: func(ctx context.Context) schema.Attribute {
		return schema.Int64Attribute{
			Optional:    true,
			Description: "Set the ping count, " + pingCountMinValidator.Description(ctx) + "; " + pingCountMaxValidator.Description(ctx),
			Validators:  []validator.Int64{pingCountMinValidator, pingCountMaxValidator},
		}
	},
	fields.FailureHopCount: func(ctx context.Context) schema.Attribute {
		return schema.Int64Attribute{
			Optional:    true,
			Description: "Set the failure hop count, " + failureHopCountMinValidator.Description(ctx) + "; " + failureHopCountMaxValidator.Description(ctx),
			Validators:  []validator.Int64{failureHopCountMinValidator, failureHopCountMaxValidator},
		}
	},
	types.FavorFastestRoundTripNameserver: func(ctx context.Context) schema.Attribute {
		return schema.BoolAttribute{
			Optional:    true,
			Description: "True enables favor fastest round trip nameserver setting",
		}
	},
	types.StopTestOnDocumentComplete: func(ctx context.Context) schema.Attribute {
		return schema.BoolAttribute{
			Optional:    true,
			Description: "True enables stop test on document complete setting",
		}
	},
	types.StopTestOnDOMContentLoad: func(ctx context.Context) schema.Attribute {
		return schema.BoolAttribute{
			Optional:    true,
			Description: "True enables stop test on DOM content load setting",
		}
	},
	types.T30xRedirectsDoNotFollow: func(ctx context.Context) schema.Attribute {
		return schema.BoolAttribute{
			Optional:    true,
			Description: "True enables 30x redirects do not follow setting",
		}
	},
	types.TryNextNameserverOnFailure: func(ctx context.Context) schema.Attribute {
		return schema.BoolAttribute{
			Optional:    true,
			Description: "True enables try next nameserver on failure setting",
		}
	},
	types.VerifyTestOnFailure: func(ctx context.Context) schema.Attribute {
		return schema.BoolAttribute{
			Optional:    true,
			Description: "True enables verify on test failure setting",
		}
	},
	fields.ViewportHeight: func(ctx context.Context) schema.Attribute {
		return schema.Int64Attribute{
			Optional:    true,
			Description: "Set the viewport height, " + viewportHeightMinValidator.Description(ctx) + "; " + viewportHeightMaxValidator.Description(ctx),
			Validators:  []validator.Int64{viewportHeightMinValidator, viewportHeightMaxValidator},
		}
	},
	fields.ViewportWidth: func(ctx context.Context) schema.Attribute {
		return schema.Int64Attribute{
			Optional:    true,
			Description: "Set the viewport width, " + viewportWidthMinValidator.Description(ctx) + "; " + viewportWidthMaxValidator.Description(ctx),
			Validators:  []validator.Int64{viewportWidthMinValidator, viewportWidthMaxValidator},
		}
	},
	fields.WaitForNoActivity: func(ctx context.Context) schema.Attribute {
		return schema.Int64Attribute{
			Optional:    true,
			Description: "Set the time value in ms to stop the test after no network activity on document complete. Use with 'stop_test_on_document_complete' flag, " + waitForNoActivityValidator.Description(ctx),
			Validators:  []validator.Int64{waitForNoActivityValidator},
		}
	},
	types.ZoneDataCollectionEnabled: func(ctx context.Context) schema.Attribute {
		return schema.BoolAttribute{
			Optional:    true,
			Description: "True enables zone data collection setting",
		}
	},
}
