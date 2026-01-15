package alerttypes

import (
	"testing"

	"catchpoint-provider/internal/testutil"
	"catchpoint-provider/internal/types"
)

// The difference between alert subtypes for monitors is only in the Timing (Response Time) alert type.
// However, none of these types were available in the previous version of the provider, so we will
// need to add them later.
var webMobileValidSubtypes = []string{
	/* TODO: Not defined yet...
	types.CumulativeLayoutShift,
	types.FirstContentfulPaint,
	types.FirstPaint,
	types.LargestContentfulPaint,*/
}

var webEmulatedHTTPValidSubtypes = []string{
	// TODO: DataTransferTime not fully supported yet.,
}

var webChromePlaybackValidSubtypes = []string{
	/* TODO: Not defined yet...
	types.CumulativeLayoutShift,
	types.FirstContentfulPaint,
	types.FirstPaint,
	types.LargestContentfulPaint,
	types.DataTransferTime, */
}

func TestGetMonitorAlertTypesForWeb(t *testing.T) {
	test := []struct {
		testMonitor string
	}{
		{types.HTTPString},
		{types.ChromeString},
		{types.EmulatedString},
		{types.PlaybackString},
		{types.MobilePlaybackString},
		{types.MobileString},
	}
	for _, tt := range test {
		testWebAlertMatrix := GetMonitorAlertTypes(types.WebType, tt.testMonitor)

		// The overall alert types for all web monitors are the same.
		// The subtypes vary by monitor, which is tested in other tests.
		expectedAlertTypes := []string{
			types.Availability,
			types.ByteLength,
			types.ContentMatch,
			types.ExperienceScore, // No subtypes
			types.HostFailure,     // No subtypes
			types.Address,
			types.Insight,
			// TODO: this was not in the previous version but is available on the portal: types.JavaScriptError, // No subtypes
			types.Ping,
			types.Requests,
			types.Timing,
			types.TestFailure, // No subtypes
			// TODO: types.Zone - add this.
		}

		testutil.AssertNotNil(t, "WebAlertTypeSchema", testWebAlertMatrix)
		testutil.AssertElementsMatch(t, "WebAlertTypeSchema", testWebAlertMatrix.ListAllValidAlertTypes(), expectedAlertTypes)
	}
}

func TestGetTestCompatibilityMatrixForWeb(t *testing.T) {
	testWebAlertMatrix := GetTestCompatibilityMatrix(types.WebType)

	expectedAlertTypes := []string{
		types.Availability,
		types.ByteLength,
		types.ContentMatch,
		types.ExperienceScore, // No subtypes
		types.HostFailure,     // No subtypes
		types.Address,
		types.Insight,
		// TODO: this was not in the previous version but is available on the portal: types.JavaScriptError, // No subtypes
		types.Ping,
		types.Requests,
		types.Timing,
		types.TestFailure, // No subtypes
		// TODO: types.Zone - add this.
	}

	testutil.AssertNotNil(t, "WebAlertTypeSchema", testWebAlertMatrix)
	testutil.AssertElementsMatch(t, "WebAlertTypeSchema", testWebAlertMatrix.ListAllValidAlertTypes(), expectedAlertTypes)
}

func TestWebByteLengthSubTypes(t *testing.T) {
	test := []struct {
		testMonitor string
	}{
		{types.HTTPString},
		{types.ChromeString},
		{types.EmulatedString},
		{types.PlaybackString},
		{types.MobilePlaybackString},
		{types.MobileString},
	}
	for _, tt := range test {
		testWebAlertMatrix := GetMonitorAlertTypes(types.WebType, tt.testMonitor)

		testutil.AssertEqual(t, "WebByteLengthSubTypes", testWebAlertMatrix.IsAlertTypeValid(types.ByteLength), true)
		testutil.AssertElementsMatch(t, "WebByteLengthSubTypes", testWebAlertMatrix.AlertTypes[types.ByteLength].SubTypes, byteLengthValidSubtypes)
	}
}

func TestWebAddressSubTypes(t *testing.T) {
	test := []struct {
		testMonitor string
	}{
		{types.HTTPString},
		{types.ChromeString},
		{types.EmulatedString},
		{types.PlaybackString},
		{types.MobilePlaybackString},
		{types.MobileString},
	}
	for _, tt := range test {
		testWebAlertMatrix := GetMonitorAlertTypes(types.WebType, tt.testMonitor)

		testutil.AssertEqual(t, "WebAddressSubTypes", testWebAlertMatrix.IsAlertTypeValid(types.Address), true)
		testutil.AssertElementsMatch(t, "WebAddressSubTypes", testWebAlertMatrix.AlertTypes[types.Address].SubTypes, addressValidSubtypes)
	}
}

func TestWebInsightSubTypes(t *testing.T) {
	test := []struct {
		testMonitor string
	}{
		{types.HTTPString},
		{types.ChromeString},
		{types.EmulatedString},
		{types.PlaybackString},
		{types.MobilePlaybackString},
		{types.MobileString},
	}
	for _, tt := range test {
		testWebAlertMatrix := GetMonitorAlertTypes(types.WebType, tt.testMonitor)

		testutil.AssertEqual(t, "WebInsightSubTypes", testWebAlertMatrix.IsAlertTypeValid(types.Insight), true)
		testutil.AssertElementsMatch(t, "WebInsightSubTypes", testWebAlertMatrix.AlertTypes[types.Insight].SubTypes, insightValidSubtypes)
	}
}

func TestWebPingSubTypes(t *testing.T) {
	test := []struct {
		testMonitor string
	}{
		{types.HTTPString},
		{types.ChromeString},
		{types.EmulatedString},
		{types.PlaybackString},
		{types.MobilePlaybackString},
		{types.MobileString},
	}
	for _, tt := range test {
		testWebAlertMatrix := GetMonitorAlertTypes(types.WebType, tt.testMonitor)

		testutil.AssertEqual(t, "WebPingSubTypes", testWebAlertMatrix.IsAlertTypeValid(types.Ping), true)
		testutil.AssertElementsMatch(t, "WebPingSubTypes", testWebAlertMatrix.AlertTypes[types.Ping].SubTypes, pingValidSubtypes)
	}
}

func TestWebTimingSubTypes(t *testing.T) {
	test := []struct {
		testMonitor     string
		additionalTypes []string
	}{
		{types.HTTPString, webEmulatedHTTPValidSubtypes},
		{types.ChromeString, webChromePlaybackValidSubtypes},
		{types.EmulatedString, webEmulatedHTTPValidSubtypes},
		{types.PlaybackString, webChromePlaybackValidSubtypes},
		{types.MobilePlaybackString, webChromePlaybackValidSubtypes},
		{types.MobileString, webMobileValidSubtypes},
	}
	for _, tt := range test {
		testWebAlertMatrix := GetMonitorAlertTypes(types.WebType, tt.testMonitor)

		WebTimingSubTypes := []string{
			types.Connect,
			types.ContentLoad,
			types.DNS,
			types.DocumentComplete,
			types.DOMLoad,
			types.DaysToExpiration,
			types.Load,
			types.Redirect,
			types.Response,
			types.Send,
			types.ServerResponse,
			types.TestTime,
			types.TestTimeWithSuspect,
			// TODO: types.TimeToFirstByte not defined yet.
			types.Wait, // Extended
		}

		testutil.AssertEqual(t, "WebTimingSubTypes", testWebAlertMatrix.IsAlertTypeValid(types.Timing), true)
		testutil.AssertElementsMatch(t, "WebTimingSubTypes", testWebAlertMatrix.AlertTypes[types.Timing].SubTypes, WebTimingSubTypes)
	}
}

func TestWebAvailabilitySubTypes(t *testing.T) {
	test := []struct {
		testMonitor string
	}{
		{types.HTTPString},
		{types.ChromeString},
		{types.EmulatedString},
		{types.PlaybackString},
		{types.MobilePlaybackString},
		{types.MobileString},
	}
	for _, tt := range test {
		testWebAlertMatrix := GetMonitorAlertTypes(types.WebType, tt.testMonitor)

		testutil.AssertEqual(t, "WebAvailabilitySubTypes", testWebAlertMatrix.IsAlertTypeValid(types.Availability), true)
		testutil.AssertElementsMatch(t, "WebAvailabilitySubTypes", testWebAlertMatrix.AlertTypes[types.Availability].SubTypes, webAvailabilityValidSubtypes)
	}
}

func TestWebContentMatchSubTypes(t *testing.T) {
	test := []struct {
		testMonitor string
	}{
		{types.HTTPString},
		{types.ChromeString},
		{types.EmulatedString},
		{types.PlaybackString},
		{types.MobilePlaybackString},
		{types.MobileString},
	}
	for _, tt := range test {
		testWebAlertMatrix := GetMonitorAlertTypes(types.WebType, tt.testMonitor)

		testutil.AssertEqual(t, "WebContentMatchSubTypes", testWebAlertMatrix.IsAlertTypeValid(types.ContentMatch), true)
		testutil.AssertElementsMatch(t, "WebContentMatchSubTypes", testWebAlertMatrix.AlertTypes[types.ContentMatch].SubTypes, extendedContentMatchValidSubtypes)
	}
}

func TestWebRequestsSubTypes(t *testing.T) {
	test := []struct {
		testMonitor string
	}{
		{types.HTTPString},
		{types.ChromeString},
		{types.EmulatedString},
		{types.PlaybackString},
		{types.MobilePlaybackString},
		{types.MobileString},
	}
	for _, tt := range test {
		testWebAlertMatrix := GetMonitorAlertTypes(types.WebType, tt.testMonitor)

		testutil.AssertEqual(t, "WebRequestsSubTypes", testWebAlertMatrix.IsAlertTypeValid(types.Requests), true)
		testutil.AssertElementsMatch(t, "WebRequestsSubTypes", testWebAlertMatrix.AlertTypes[types.Requests].SubTypes, requestsValidSubtypes)
	}
}
