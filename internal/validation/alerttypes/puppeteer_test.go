package alerttypes

import (
	"testing"

	"catchpoint-provider/internal/testutil"
	"catchpoint-provider/internal/types"
)

func TestGetMonitorAlertTypesForPuppeteer(t *testing.T) {
	test := []struct {
		testMonitor string
	}{
		{types.ChromeString},
	}
	for _, tt := range test {
		// The Puppeteer test type accepts monitors Edge and Chrome which both accept the same alert types.
		testPuppeteerAlertMatrix := GetMonitorAlertTypes(types.PuppeteerType, tt.testMonitor)

		expectedAlertTypes := []string{
			types.Availability,
			types.ByteLength,
			types.ContentMatch,
			types.ExperienceScore, // No subtypes
			types.HostFailure,     // No subtypes
			types.Address,
			types.Insight,
			types.Ping,
			types.Requests,
			types.Timing,
			types.TestFailure, // No subtypes
			// types.Zone, // TODO: Zone alert type needs to be added here.
		}

		testutil.AssertNotNil(t, "PuppeteerAlertTypeSchema", testPuppeteerAlertMatrix)
		testutil.AssertElementsMatch(t, "PuppeteerAlertTypeSchema", testPuppeteerAlertMatrix.ListAllValidAlertTypes(), expectedAlertTypes)
	}
}

func TestGetTestCompatibilityMatrixForPuppeteer(t *testing.T) {
	testPuppeteerAlertMatrix := GetTestCompatibilityMatrix(types.PuppeteerType)

	expectedAlertTypes := []string{
		types.Availability,
		types.ByteLength,
		types.ContentMatch,
		types.ExperienceScore, // No subtypes
		types.HostFailure,     // No subtypes
		types.Address,
		types.Insight,
		types.Ping,
		types.Requests,
		types.Timing,
		types.TestFailure, // No subtypes
		// types.Zone, // TODO: Zone alert type needs to be added here.
	}

	testutil.AssertNotNil(t, "PuppeteerAlertTypeSchema", testPuppeteerAlertMatrix)
	testutil.AssertElementsMatch(t, "PuppeteerAlertTypeSchema", testPuppeteerAlertMatrix.ListAllValidAlertTypes(), expectedAlertTypes)
}

func TestPuppeteerAvailabilitySubTypes(t *testing.T) {
	test := []struct {
		testMonitor string
	}{
		{types.ChromeString},
	}
	for _, tt := range test {
		testPuppeteerAlertMatrix := GetMonitorAlertTypes(types.PuppeteerType, tt.testMonitor)

		testutil.AssertEqual(t, "PuppeteerAvailabilitySubTypes", testPuppeteerAlertMatrix.IsAlertTypeValid(types.Availability), true)
		testutil.AssertElementsMatch(t, "PuppeteerAvailabilitySubTypes", testPuppeteerAlertMatrix.AlertTypes[types.Availability].SubTypes, webAvailabilityValidSubtypes)
	}
}

func TestPuppeteerByteLengthSubTypes(t *testing.T) {
	test := []struct {
		testMonitor string
	}{
		{types.ChromeString},
	}
	for _, tt := range test {
		testPuppeteerAlertMatrix := GetMonitorAlertTypes(types.PuppeteerType, tt.testMonitor)

		testutil.AssertEqual(t, "PuppeteerByteLengthSubTypes", testPuppeteerAlertMatrix.IsAlertTypeValid(types.ByteLength), true)
		testutil.AssertElementsMatch(t, "PuppeteerByteLengthSubTypes", testPuppeteerAlertMatrix.AlertTypes[types.ByteLength].SubTypes, byteLengthValidSubtypes)
	}
}

func TestPuppeteerContentMatchSubTypes(t *testing.T) {
	test := []struct {
		testMonitor string
	}{
		{types.ChromeString},
	}
	for _, tt := range test {
		testPuppeteerAlertMatrix := GetMonitorAlertTypes(types.PuppeteerType, tt.testMonitor)

		testutil.AssertEqual(t, "PuppeteerContentMatchSubTypes", testPuppeteerAlertMatrix.IsAlertTypeValid(types.ContentMatch), true)
		testutil.AssertElementsMatch(t, "PuppeteerContentMatchSubTypes", testPuppeteerAlertMatrix.AlertTypes[types.ContentMatch].SubTypes, basicContentMatchValidSubtypes)
	}
}

func TestPuppeteerAddressSubTypes(t *testing.T) {
	test := []struct {
		testMonitor string
	}{
		{types.ChromeString},
	}
	for _, tt := range test {
		testPuppeteerAlertMatrix := GetMonitorAlertTypes(types.PuppeteerType, tt.testMonitor)

		testutil.AssertEqual(t, "PuppeteerAddressSubTypes", testPuppeteerAlertMatrix.IsAlertTypeValid(types.Address), true)
		testutil.AssertElementsMatch(t, "PuppeteerAddressSubTypes", testPuppeteerAlertMatrix.AlertTypes[types.Address].SubTypes, addressValidSubtypes)
	}
}

func TestPuppeteerInsightSubTypes(t *testing.T) {
	test := []struct {
		testMonitor string
	}{
		{types.ChromeString},
	}
	for _, tt := range test {
		testPuppeteerAlertMatrix := GetMonitorAlertTypes(types.PuppeteerType, tt.testMonitor)

		testutil.AssertEqual(t, "PuppeteerInsightSubTypes", testPuppeteerAlertMatrix.IsAlertTypeValid(types.Insight), true)
		testutil.AssertElementsMatch(t, "PuppeteerInsightSubTypes", testPuppeteerAlertMatrix.AlertTypes[types.Insight].SubTypes, insightValidSubtypes)
	}
}

func TestPuppeteerPingSubTypes(t *testing.T) {
	test := []struct {
		testMonitor string
	}{
		{types.ChromeString},
	}
	for _, tt := range test {
		testPuppeteerAlertMatrix := GetMonitorAlertTypes(types.PuppeteerType, tt.testMonitor)

		testutil.AssertEqual(t, "PuppeteerPingSubTypes", testPuppeteerAlertMatrix.IsAlertTypeValid(types.Ping), true)
		testutil.AssertElementsMatch(t, "PuppeteerPingSubTypes", testPuppeteerAlertMatrix.AlertTypes[types.Ping].SubTypes, pingValidSubtypes)
	}
}

func TestPuppeteerRequestsSubTypes(t *testing.T) {
	test := []struct {
		testMonitor string
	}{
		{types.ChromeString},
	}
	for _, tt := range test {
		testPuppeteerAlertMatrix := GetMonitorAlertTypes(types.PuppeteerType, tt.testMonitor)

		testutil.AssertEqual(t, "PuppeteerRequestsSubTypes", testPuppeteerAlertMatrix.IsAlertTypeValid(types.Requests), true)
		testutil.AssertElementsMatch(t, "PuppeteerRequestsSubTypes", testPuppeteerAlertMatrix.AlertTypes[types.Requests].SubTypes, requestsValidSubtypes)
	}
}

func TestPuppeteerTimingSubTypes(t *testing.T) {
	test := []struct {
		testMonitor string
	}{
		{types.ChromeString},
	}
	for _, tt := range test {
		testPuppeteerAlertMatrix := GetMonitorAlertTypes(types.PuppeteerType, tt.testMonitor)

		PuppeteerTimingSubTypes := []string{
			types.Connect,
			types.ContentLoad,
			// TODO: CumulativeLayoutShift not fully supported yet.
			types.DNS,
			// TODO: DataTransferTime not fully supported yet.,
			types.DocumentComplete,
			types.DOMLoad,
			// TODO: FirstContentfulPaint not fully supported yet.
			// TODO: FirstPaint not fully supported yet.
			// TODO: LargestContentfulPaint not fully supported yet.
			types.Load,
			types.Redirect,
			types.Response,
			types.Send,
			types.ServerResponse,
			types.TestTime,
			types.TestTimeWithSuspect,
			// TODO: types.TimeToFirstByte not defined yet.
			// TODO: types.TotalBlockingTime not defined yet.
		}

		testutil.AssertEqual(t, "PuppeteerTimingSubTypes", testPuppeteerAlertMatrix.IsAlertTypeValid(types.Timing), true)
		testutil.AssertElementsMatch(t, "PuppeteerTimingSubTypes", testPuppeteerAlertMatrix.AlertTypes[types.Timing].SubTypes, PuppeteerTimingSubTypes)
	}
}
