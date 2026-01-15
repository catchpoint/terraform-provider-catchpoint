package alerttypes

import (
	"fmt"
	maps0 "maps"

	"catchpoint-provider/internal/fields"
	"catchpoint-provider/internal/helpers"
	"catchpoint-provider/internal/types"
)

var GetMonitorAlertTypesFunc = GetMonitorAlertTypes

// #region ASN

var validASNMap = map[string]*AlertSubTypes{
	types.ASN: {
		SubTypes: []string{
			types.OriginAS,
			types.OriginNeighbor,
			types.PathAS,
			types.PrefixMismatch,
		},
	},
}

// #endregion

// #region Availability

// Ping, Traceroute, DNS Availability alertType allows 'DowntimePercent', and 'Test'.
var validCommonAvailabilityMap = map[string]*AlertSubTypes{
	types.Availability: {
		SubTypes: []string{types.DowntimePercent, types.Test},
	},
}

// Only BGP. Availability alertType allows 'DowntimePercent', 'Test', and 'Reachability'.
var validBGPAvailabilityMap = extendAlertTypeSubTypes(validCommonAvailabilityMap, types.Reachability)

// For Browser-based tests (Playwright, Puppeteer, API, Transaction, Web),
// the availability alertType allows 'DowntimePercent', 'Test', and 'Content'.
var validBrowserAvailabilityMap = extendAlertTypeSubTypes(validCommonAvailabilityMap, types.Content)

// #endregion

// #region ByteLength

// ByteLength type always has the same subtypes when used. File Size, Page, Byte Length (Test URL).
var validByteLengthMap = map[string]*AlertSubTypes{
	types.ByteLength: {
		SubTypes: []string{
			types.FileSize,
			types.Page,
			types.ByteLength, // This is displayed on the portal as "Test Url"
		},
	},
}

// #endregion

// #region ContentMatch

// ContentMatch has no subtypes when used with SSL tests.
var validContentMatchMap = map[string]*AlertSubTypes{
	types.ContentMatch: {},
}

// For PW/Puppeteer, API, and Transaction tests, the ContentMatch alertType
// allows 'RegularExpression', 'ResponseCode', and 'ResponseHeaders'.
var validBasicContentMatchMap = map[string]*AlertSubTypes{
	types.ContentMatch: {
		SubTypes: []string{
			types.RegularExpression,
			types.ResponseCode,
			types.ResponseHeaders,
		},
	},
}

// For Web tests, the ContentMatch alertType allows for the same as Basic but also with 'Compare To File' and
// 'Compare To Previous'.
var validExpandedContentMatchMap = extendAlertTypeSubTypes(
	validBasicContentMatchMap,
	types.CompareToFile,
	types.CompareToPrevious,
)

// #endregion

// #region DNS

var validDNSMatchMap = map[string]*AlertSubTypes{
	types.DNS: {
		SubTypes: []string{
			types.DNSAdditional,
			types.DNSAnswer,
			types.DNSAuthority,
			types.DNSGeneral,
		},
	},
}

// #endregion

// #region Experience Score

var validExperienceScoreMatchMap = map[string]*AlertSubTypes{
	types.ExperienceScore: {},
}

// #endregion

// #region Host Failure

// 'Host Failure' never has a subtype.
var validHostFailureScoreMatchMap = map[string]*AlertSubTypes{types.HostFailure: {}}

// #endregion

// #region IPAddress

// DNS, Traceroute, and Ping tests have the 'IP Address' alertType with no subtypes.
var validBaseIPAddressMatchMap = map[string]*AlertSubTypes{
	types.Address: {},
}

// PW/Puppeteer, API, Transaction and Web tests have an 'IP Address' alertType with 'Any URLs', 'Child URLs', and 'Test URLs'.
var validExtendedIPAddressMatchMap = map[string]*AlertSubTypes{
	types.Address: {
		SubTypes: []string{
			types.AnyURLs,
			types.ChildURLs,
			types.TestURLs,
		},
	},
}

// #endregion

// #region Insight

// Insight has the same subtypes for all tests.
var validInsightMatchMap = map[string]*AlertSubTypes{
	types.Insight: {
		SubTypes: []string{types.Indicators, types.Tracepoints},
	},
}

// #endregion

// #region Path (Traceroute only)

var validPathMatchMap = map[string]*AlertSubTypes{
	types.Path: {
		SubTypes: []string{
			types.ASNsNum,
			types.CitiesNum,
			types.CountriesNum,
			types.HopsNum,
		},
	},
}

// #endregion

// #region Ping

// Ping has the same subtypes for most tests except traceroute (PingPacketLoss, PingRTT).
var validBasePingMatchMap = map[string]*AlertSubTypes{
	types.Ping: {
		SubTypes: []string{types.PingPacketLoss, types.PingRTT},
	},
}

// Traceroute has the same as base but also with 'Ping Jitter' as a subtype.
var validExtendedPingMatchMap = extendAlertTypeSubTypes(
	validBasePingMatchMap,
	// TODO: this does not appear to be defined yet: types.PingJitter, but it is an acceptable value for Traceroute tests.
)

// #endregion

// #region Requests

// Requests has the same subtypes for all tests.
var validRequestsMatchMap = map[string]*AlertSubTypes{
	types.Requests: {
		SubTypes: []string{
			types.ConnectionsNum,
			types.CSSNum,
			types.FlashNum,
			types.HostsNum,
			types.HTMLNum,
			types.ImagesNum,
			types.MediaNum,
			types.OtherNum,
			types.RedirectsNum,
			types.RequestsNum,
			types.ScriptsNum,
			types.XMLNum,
		},
	},
}

// #endregion

// #region Test Failure

// 'Test Failure' never has a subtype.
var validTestFailureMap = map[string]*AlertSubTypes{types.TestFailure: {}}

// #endregion

// #region Monitor AlertTypeCompatibilityMatrices

// An API test has 'Availability', 'ByteLength', 'ContentMatch', 'Experience Score', 'Host Failure', 'IP Address',
// 'Insight', 'Ping', 'Requests', 'Response Time', and 'Test Failure' alertTypes.
var apiAlertMatrix = &MonitorAlertTypeCompatibilityMatrix{
	TestType: types.APIType,
	MonitorTypes: map[string]*MonitorAlertTypes{
		types.APIString: {
			AlertTypes: mergeMaps(
				validBrowserAvailabilityMap,
				validByteLengthMap,
				validBasicContentMatchMap,
				validExperienceScoreMatchMap,
				validHostFailureScoreMatchMap,
				validExtendedIPAddressMatchMap,
				validInsightMatchMap,
				validBasePingMatchMap,
				validRequestsMatchMap,
				validAPIResponseTimeMatchMap,
				validTestFailureMap,
			),
		},
	},
}

var validBGPAlertTypes = mergeMaps(
	validASNMap,
	validBGPAvailabilityMap,
	validTestFailureMap,
)

var bgpAlertMatrix = &MonitorAlertTypeCompatibilityMatrix{
	TestType: types.BGPType,
	MonitorTypes: map[string]*MonitorAlertTypes{
		types.BGPString:      {AlertTypes: validBGPAlertTypes},
		types.BGPBasicString: {AlertTypes: validBGPAlertTypes},
	},
}

var validDNSAlertTypes = mergeMaps(
	validCommonAvailabilityMap,
	validDNSMatchMap,
	validExperienceScoreMatchMap,
	validBaseIPAddressMatchMap,
	validBasePingMatchMap,
	validEmptyResponseTimeMatchMap,
	validTestFailureMap,
)

var dnsAlertMatrix = &MonitorAlertTypeCompatibilityMatrix{
	TestType: types.DNSType,
	MonitorTypes: map[string]*MonitorAlertTypes{
		types.DNSDirectString:     {AlertTypes: validDNSAlertTypes},
		types.DNSExperienceString: {AlertTypes: validDNSAlertTypes},
	},
}

var sslAlertMatrix = &MonitorAlertTypeCompatibilityMatrix{
	TestType: types.SSLType,
	MonitorTypes: map[string]*MonitorAlertTypes{
		types.SSLString: {
			AlertTypes: mergeMaps(
				validContentMatchMap,
				validExperienceScoreMatchMap,
				validSSLResponseTimeMatchMap,
				validTestFailureMap,
			),
		},
	},
}

// All 3 Ping monitor types have the same alert types and subtypes.
var validPingAlertTypes = mergeMaps(
	validCommonAvailabilityMap,
	validExperienceScoreMatchMap,
	validBaseIPAddressMatchMap,
	validBasePingMatchMap,
	validEmptyResponseTimeMatchMap,
	validTestFailureMap,
)

var pingAlertMatrix = &MonitorAlertTypeCompatibilityMatrix{
	TestType: types.PingType,
	MonitorTypes: map[string]*MonitorAlertTypes{
		types.PingICMPString: {AlertTypes: validPingAlertTypes},
		types.PingUDPString:  {AlertTypes: validPingAlertTypes},
		types.PingTCPString:  {AlertTypes: validPingAlertTypes},
	},
}

var validTracerouteAlertTypes = mergeMaps(
	validASNMap,
	validCommonAvailabilityMap,
	validExperienceScoreMatchMap,
	validBaseIPAddressMatchMap,
	validPathMatchMap,
	validExtendedPingMatchMap,
	validTestFailureMap,
)

var tracerouteAlertMatrix = &MonitorAlertTypeCompatibilityMatrix{
	TestType: types.TracerouteType,
	MonitorTypes: map[string]*MonitorAlertTypes{
		types.TracerouteICMPString: {AlertTypes: validTracerouteAlertTypes},
		types.TracerouteUDPString:  {AlertTypes: validTracerouteAlertTypes},
		types.TracerouteTCPString:  {AlertTypes: validTracerouteAlertTypes},
	},
}

var validBaseTransactionAlertTypes = mergeMaps(
	validBrowserAvailabilityMap,
	validByteLengthMap,
	validBasicContentMatchMap,
	validExperienceScoreMatchMap,
	validHostFailureScoreMatchMap,
	validExtendedIPAddressMatchMap,
	// TODO: add Javascript Failure
	validInsightMatchMap,
	validBasePingMatchMap,
	validRequestsMatchMap,
	validTestFailureMap,
	// validZoneMap, // TODO: Zone alert type needs to be added here.
)

var validTransactionChromeAlertTypes = mergeMaps(
	validBaseTransactionAlertTypes,
	validTransactionChromeResponseTimeMatchMap,
)

var validTransactionEmulatedAlertTypes = mergeMaps(
	validBaseTransactionAlertTypes,
	validTransactionEmulatedResponseTimeMatchMap,
)

var validTransactionMobileAlertTypes = mergeMaps(
	validBaseTransactionAlertTypes,
	validTransactionMobileResponseTimeMatchMap,
)

var transactionAlertMatrix = &MonitorAlertTypeCompatibilityMatrix{
	TestType: types.TransactionType,
	MonitorTypes: map[string]*MonitorAlertTypes{
		types.ChromeString:   {AlertTypes: validTransactionChromeAlertTypes},
		types.EmulatedString: {AlertTypes: validTransactionEmulatedAlertTypes},
		types.MobileString:   {AlertTypes: validTransactionMobileAlertTypes},
	},
}

var validBaseWebAlertTypes = mergeMaps(
	validBrowserAvailabilityMap,
	validByteLengthMap,
	validExpandedContentMatchMap,
	validExperienceScoreMatchMap,
	validHostFailureScoreMatchMap,
	validExtendedIPAddressMatchMap,
	// TODO: add Javascript Failure
	validInsightMatchMap,
	validBasePingMatchMap,
	validRequestsMatchMap,
	validTestFailureMap,
	// TODO: add Zone
)

var validWebEmulatedHTTPAlertTypes = mergeMaps(
	validBaseWebAlertTypes,
	validWebEmulatedResponseTimeMatchMap,
)

var valiWebChromePlaybackAlertTypes = mergeMaps(
	validBaseWebAlertTypes,
	validWebChromePlaybackResponseTimeMatchMap,
)

var validWebMobileAlertTypes = mergeMaps(
	validBaseWebAlertTypes,
	validWebMobileResponseTimeMatchMap,
)

var webAlertMatrix = &MonitorAlertTypeCompatibilityMatrix{
	TestType: types.WebType,
	MonitorTypes: map[string]*MonitorAlertTypes{
		types.HTTPString:           {AlertTypes: validWebEmulatedHTTPAlertTypes},
		types.EmulatedString:       {AlertTypes: validWebEmulatedHTTPAlertTypes},
		types.MobileString:         {AlertTypes: validWebMobileAlertTypes},
		types.MobilePlaybackString: {AlertTypes: valiWebChromePlaybackAlertTypes},
		types.PlaybackString:       {AlertTypes: valiWebChromePlaybackAlertTypes},
		types.ChromeString:         {AlertTypes: valiWebChromePlaybackAlertTypes},
	},
}

var validPlaywrightAlertTypes = mergeMaps(
	validBrowserAvailabilityMap,
	validByteLengthMap,
	validBasicContentMatchMap,
	validExperienceScoreMatchMap,
	validHostFailureScoreMatchMap,
	validExtendedIPAddressMatchMap,
	validInsightMatchMap,
	validBasePingMatchMap,
	validRequestsMatchMap,
	validPlaywrightPuppeteerResponseTimeMatchMap,
	validTestFailureMap,
	// validZoneMap, // TODO: Zone alert type needs to be added here.
)

var playwrightAlertMatrix = &MonitorAlertTypeCompatibilityMatrix{
	TestType: types.PlaywrightType,
	MonitorTypes: map[string]*MonitorAlertTypes{
		types.EdgeString:   {AlertTypes: validPlaywrightAlertTypes},
		types.ChromeString: {AlertTypes: validPlaywrightAlertTypes},
	},
}

// Currently, puppeteer gets the same alert types as Playwright.
var puppeteerAlertMatrix = &MonitorAlertTypeCompatibilityMatrix{
	TestType: types.PuppeteerType,
	MonitorTypes: map[string]*MonitorAlertTypes{
		types.ChromeString: {AlertTypes: validPlaywrightAlertTypes},
	},
}

// #endregion

func ValidateTestAlertTypeCombination(alertSettingsList []map[string]any, test types.TestType, monitor string) error {
	monitorAlertTypes := GetMonitorAlertTypesFunc(test, monitor)
	if monitorAlertTypes == nil {
		return fmt.Errorf("invalid monitor type '%s' for test type '%v'", monitor, test)
	}

	for _, alertSetting := range alertSettingsList {
		alertType, _ := alertSetting[fields.AlertType].(string)
		alertSubType, _ := alertSetting[fields.AlertSubType].(string)

		if !monitorAlertTypes.IsAlertTypeValid(alertType) {
			return fmt.Errorf("invalid alert type '%s' for monitor '%s'", alertType, monitor)
		}

		alertSubTypes := monitorAlertTypes.AlertTypes[alertType]

		if alertSubType != types.EmptyString && !alertSubTypes.IsAlertSubTypeValid(alertSubType) {
			return fmt.Errorf("invalid alert sub type '%s' for alert type '%s'", alertSubType, alertType)
		}
	}
	return nil
}

// #region Helper functions

func mergeMaps(maps ...map[string]*AlertSubTypes) map[string]*AlertSubTypes {
	result := make(map[string]*AlertSubTypes)
	for _, m := range maps {
		maps0.Copy(result, m)
	}
	return result
}

func extendAlertTypeSubTypes(base map[string]*AlertSubTypes, extra ...string) map[string]*AlertSubTypes {
	// Assumes base has only one key
	for k, v := range base {
		return map[string]*AlertSubTypes{
			k: {
				SubTypes: helpers.FlatCombineStringSlicesUnique(v.SubTypes, extra),
			},
		}
	}
	// This can only be reached on developer error. The extendAlertTypeSubTypes needs a base with the same key
	// as the one it is extending.
	panic("base map must have exactly one key")
}

// #endregion
