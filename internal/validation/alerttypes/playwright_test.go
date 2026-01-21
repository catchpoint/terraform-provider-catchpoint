package alerttypes

import (
	"testing"

	"catchpoint-provider/internal/testutil"
	"catchpoint-provider/internal/types"
)

func TestGetMonitorAlertTypesForPlaywright(t *testing.T) {
	test := []struct {
		testMonitor string
	}{
		{types.EdgeString},
		{types.ChromeString},
	}
	for _, tt := range test {
		// The Playwright test type accepts monitors Edge and Chrome which both accept the same alert types.
		testPlaywrightAlertMatrix := GetMonitorAlertTypes(types.PlaywrightType, tt.testMonitor)

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

		testutil.AssertNotNil(t, "PlaywrightAlertTypeSchema", testPlaywrightAlertMatrix)
		testutil.AssertElementsMatch(t, "PlaywrightAlertTypeSchema", testPlaywrightAlertMatrix.ListAllValidAlertTypes(), expectedAlertTypes)
	}
}

func TestGetTestCompatibilityMatrixForPlaywright(t *testing.T) {
	testPlaywrightAlertMatrix := GetTestCompatibilityMatrix(types.PlaywrightType)

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

	testutil.AssertNotNil(t, "PlaywrightAlertTypeSchema", testPlaywrightAlertMatrix)
	testutil.AssertElementsMatch(t, "PlaywrightAlertTypeSchema", testPlaywrightAlertMatrix.ListAllValidAlertTypes(), expectedAlertTypes)
}

func TestPlaywrightAvailabilitySubTypes(t *testing.T) {
	test := []struct {
		testMonitor string
	}{
		{types.EdgeString},
		{types.ChromeString},
	}
	for _, tt := range test {
		testPlaywrightAlertMatrix := GetMonitorAlertTypes(types.PlaywrightType, tt.testMonitor)

		testutil.AssertEqual(t, "PlaywrightAvailabilitySubTypes", testPlaywrightAlertMatrix.IsAlertTypeValid(types.Availability), true)
		testutil.AssertElementsMatch(t, "PlaywrightAvailabilitySubTypes", testPlaywrightAlertMatrix.AlertTypes[types.Availability].SubTypes, webAvailabilityValidSubtypes)
	}
}

func TestPlaywrightByteLengthSubTypes(t *testing.T) {
	test := []struct {
		testMonitor string
	}{
		{types.EdgeString},
		{types.ChromeString},
	}
	for _, tt := range test {
		testPlaywrightAlertMatrix := GetMonitorAlertTypes(types.PlaywrightType, tt.testMonitor)

		testutil.AssertEqual(t, "PlaywrightByteLengthSubTypes", testPlaywrightAlertMatrix.IsAlertTypeValid(types.ByteLength), true)
		testutil.AssertElementsMatch(t, "PlaywrightByteLengthSubTypes", testPlaywrightAlertMatrix.AlertTypes[types.ByteLength].SubTypes, byteLengthValidSubtypes)
	}
}

func TestPlaywrightContentMatchSubTypes(t *testing.T) {
	test := []struct {
		testMonitor string
	}{
		{types.EdgeString},
		{types.ChromeString},
	}
	for _, tt := range test {
		testPlaywrightAlertMatrix := GetMonitorAlertTypes(types.PlaywrightType, tt.testMonitor)

		testutil.AssertEqual(t, "PlaywrightContentMatchSubTypes", testPlaywrightAlertMatrix.IsAlertTypeValid(types.ContentMatch), true)
		testutil.AssertElementsMatch(t, "PlaywrightContentMatchSubTypes", testPlaywrightAlertMatrix.AlertTypes[types.ContentMatch].SubTypes, basicContentMatchValidSubtypes)
	}
}

func TestPlaywrightAddressSubTypes(t *testing.T) {
	test := []struct {
		testMonitor string
	}{
		{types.EdgeString},
		{types.ChromeString},
	}
	for _, tt := range test {
		testPlaywrightAlertMatrix := GetMonitorAlertTypes(types.PlaywrightType, tt.testMonitor)

		testutil.AssertEqual(t, "PlaywrightAddressSubTypes", testPlaywrightAlertMatrix.IsAlertTypeValid(types.Address), true)
		testutil.AssertElementsMatch(t, "PlaywrightAddressSubTypes", testPlaywrightAlertMatrix.AlertTypes[types.Address].SubTypes, addressValidSubtypes)
	}
}

func TestPlaywrightInsightSubTypes(t *testing.T) {
	test := []struct {
		testMonitor string
	}{
		{types.EdgeString},
		{types.ChromeString},
	}
	for _, tt := range test {
		testPlaywrightAlertMatrix := GetMonitorAlertTypes(types.PlaywrightType, tt.testMonitor)

		testutil.AssertEqual(t, "PlaywrightInsightSubTypes", testPlaywrightAlertMatrix.IsAlertTypeValid(types.Insight), true)
		testutil.AssertElementsMatch(t, "PlaywrightInsightSubTypes", testPlaywrightAlertMatrix.AlertTypes[types.Insight].SubTypes, insightValidSubtypes)
	}
}

func TestPlaywrightPingSubTypes(t *testing.T) {
	test := []struct {
		testMonitor string
	}{
		{types.EdgeString},
		{types.ChromeString},
	}
	for _, tt := range test {
		testPlaywrightAlertMatrix := GetMonitorAlertTypes(types.PlaywrightType, tt.testMonitor)

		testutil.AssertEqual(t, "PlaywrightPingSubTypes", testPlaywrightAlertMatrix.IsAlertTypeValid(types.Ping), true)
		testutil.AssertElementsMatch(t, "PlaywrightPingSubTypes", testPlaywrightAlertMatrix.AlertTypes[types.Ping].SubTypes, pingValidSubtypes)
	}
}

func TestPlaywrightRequestsSubTypes(t *testing.T) {
	test := []struct {
		testMonitor string
	}{
		{types.EdgeString},
		{types.ChromeString},
	}
	for _, tt := range test {
		testPlaywrightAlertMatrix := GetMonitorAlertTypes(types.PlaywrightType, tt.testMonitor)

		testutil.AssertEqual(t, "PlaywrightRequestsSubTypes", testPlaywrightAlertMatrix.IsAlertTypeValid(types.Requests), true)
		testutil.AssertElementsMatch(t, "PlaywrightRequestsSubTypes", testPlaywrightAlertMatrix.AlertTypes[types.Requests].SubTypes, requestsValidSubtypes)
	}
}

func TestPlaywrightTimingSubTypes(t *testing.T) {
	test := []struct {
		testMonitor string
	}{
		{types.EdgeString},
		{types.ChromeString},
	}
	for _, tt := range test {
		testPlaywrightAlertMatrix := GetMonitorAlertTypes(types.PlaywrightType, tt.testMonitor)

		PlaywrightTimingSubTypes := []string{
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

		testutil.AssertEqual(t, "PlaywrightTimingSubTypes", testPlaywrightAlertMatrix.IsAlertTypeValid(types.Timing), true)
		testutil.AssertElementsMatch(t, "PlaywrightTimingSubTypes", testPlaywrightAlertMatrix.AlertTypes[types.Timing].SubTypes, PlaywrightTimingSubTypes)
	}
}
