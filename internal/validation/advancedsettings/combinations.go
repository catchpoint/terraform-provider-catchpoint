package advancedsettings

import (
	"catchpoint-provider/internal/fields"
	"catchpoint-provider/internal/helpers"
	"catchpoint-provider/internal/types"
)

// BGP takes no advanced settings. For consistency's sake, we define it and return an empty map.
var bgpAdvancedSettings = []string{}

// Advanced settings for both DNS Direct and Experience tests.
var dnsCommonAdvancedSettings = []string{
	fields.AdditionalMonitor,         // Additional Monitor
	types.DebugPrimaryHostOnFailure,  // Debug Primary Host on Failure
	types.DisableRecursiveResolution, // DNS Recursive Resolution Disabled
	fields.EDNSSubnet,                // EDNS Subnet
	types.EnablePathMTUDiscovery,     // Enable MTU Path Discovery
	types.EnableTCPProtocol,          // Enable TCP for DNS Test Type protocol,
	types.EnableNSID,                 // Nameserver Hostname lookup Mechanism,
	types.TryNextNameserverOnFailure, // Try Next Nameserver on Failure
}

var dnsDirectAdvancedSettings = helpers.FlatCombineStringSlices(
	dnsCommonAdvancedSettings,
	[]string{
		types.EnableDNSSEC, // DNSSEC
	},
)

var dnsExperienceAdvancedSettings = helpers.FlatCombineStringSlices(
	dnsCommonAdvancedSettings,
	[]string{
		// Not yet enabled: fields.CacheTLDQueryLevels,
		types.FavorFastestRoundTripNameserver, // Favor Fastest Round-Trip Nameserver
	},
)

var pingAdvancedSettings = []string{
	fields.AdditionalMonitor,
	types.DebugPrimaryHostOnFailure,
	types.EnablePathMTUDiscovery,
	types.VerifyTestOnFailure,
}

var sslAdvancedSettings = []string{
	fields.AdditionalMonitor,
	types.CertificateRevocationDisabled,
	types.EnablePathMTUDiscovery,
	types.VerifyTestOnFailure,
}

var tracerouteAdvancedSettings = []string{
	// Unsupported? types.EnableDSCPPriorityProtocol,
	types.EnablePathMTUDiscovery,
	fields.FailureHopCount,
	fields.PingCount,
	// Unsupported? types.SetECN,
}

// #region Browser Tests

// All browser-based tests share these. Web (mobile, playback, mobile playback, emulated, http, chrome)
// Transaction (chrome, emulated, mobile), API, and Playwright.
var commonBrowserAdvancedSettings = []string{
	types.F40xOr50xHTTPMarkSuccessful,         // 40x or 50x Mark Successful
	fields.AdditionalMonitor,                  // Additional Monitor
	types.DebugPrimaryHostOnFailure,           // Debug Primary Host on Failure
	types.DebugReferencedHostsOnFailure,       // Debug Referenced Hosts on Failure
	types.EnablePathMTUDiscovery,              // Enable MTU Path Discovery
	fields.EnforceTestFailureIfRunsLongerThan, // Enforce Failure if Test Runs Longer Than
	types.HostDataCollectionEnabled,           // Host Data Collection Enabled
	types.CaptureHTTPHeaders,                  // Http Headers Capture
	types.CaptureResponseContent,              // Response Content or Metadata Capture
	types.IgnoreSSLFailures,                   // SSL Errors Ignored
	types.AllowTestDownloadLimitOverride,      // Test Size Override Enabled
	types.EnableSelfVersusThirdPartyZones,     // ThirdPartyZone
	// Not supported yet? types.Tracing,
	types.VerifyTestOnFailure,       // Verify Test on Failure
	types.ZoneDataCollectionEnabled, // Zone Data
}

// Used by Playwright, Puppeteer, Transaction (chrome, mobile), and Web (chrome, mobile).
var browserDisplaySettings = []string{
	types.CaptureFilmstrip,  // FilmstripCapture
	fields.ViewportHeight,   // Viewport
	fields.ViewportWidth,    // Viewport
	types.CaptureScreenshot, // ScreenshotCapture
}

var apiAdvancedSettings = helpers.FlatCombineStringSlices(
	commonBrowserAdvancedSettings,
	[]string{
		types.T30xRedirectsDoNotFollow, // 30x Redirects Do Not Follow
	},
)

// Playwright and Puppeteer share these advanced settings.
var playwrightAdvancedSettings = helpers.FlatCombineStringSlices(
	commonBrowserAdvancedSettings,
	browserDisplaySettings,
	[]string{
		fields.BandwidthThrottling,
		types.DisableCrossOriginIframeAccess, // Cross-Origin Iframe Do Not Allow
	},
)

// Puppeteer has all the same settings as Playwright but with StopTestOnDocumentComplete too.
var puppeteerAdvancedSettings = helpers.FlatCombineStringSlices(
	playwrightAdvancedSettings,
	[]string{
		fields.StopTestOnDocumentComplete,
		fields.WaitForNoActivity, // Enabled by Stop On Doc Complete
	},
)

var transactionChromeAdvancedSettings = helpers.FlatCombineStringSlices(
	commonBrowserAdvancedSettings,
	browserDisplaySettings,
	[]string{
		fields.BandwidthThrottling,           // Bandwidth Throttling
		types.DisableCrossOriginIframeAccess, // Cross-Origin Iframe Do Not Allow
	},
)

var transactionEmulatedAdvancedSettings = helpers.FlatCombineStringSlices(
	commonBrowserAdvancedSettings,
	[]string{
		types.T30xRedirectsDoNotFollow, // 30x Redirects Do Not Follow
	},
)

var transactionMobileAdvancedSettings = helpers.FlatCombineStringSlices(
	commonBrowserAdvancedSettings,
	browserDisplaySettings,
	[]string{
		fields.BandwidthThrottling, // Bandwidth Throttling
	},
)

var webChromeAdvancedSettings = helpers.FlatCombineStringSlices(
	commonBrowserAdvancedSettings,
	browserDisplaySettings,
	[]string{
		fields.BandwidthThrottling,           // Bandwidth Throttling
		types.DisableCrossOriginIframeAccess, // Cross-Origin Iframe Do Not Allow
		types.StopTestOnDocumentComplete,     // Stop Test on Document Complete
		fields.WaitForNoActivity,             // Enabled by Stop On Doc Complete
	},
)

var webEmulatedHttpAdvancedSettings = helpers.FlatCombineStringSlices(
	commonBrowserAdvancedSettings,
	[]string{
		types.T30xRedirectsDoNotFollow, // 30x Redirects Do Not Follow
	},
)

var webMobileAdvancedSettings = helpers.FlatCombineStringSlices(
	commonBrowserAdvancedSettings,
	browserDisplaySettings,
	[]string{
		fields.BandwidthThrottling,        // Bandwidth Throttling
		fields.StopTestOnDocumentComplete, // Stop Test on Document Complete
		fields.WaitForNoActivity,          // Enabled by Stop On Doc Complete
	},
)

var webPlaybackAdvancedSettings = helpers.FlatCombineStringSlices(
	commonBrowserAdvancedSettings,
)
