package alerttypes

import (
	"testing"

	"catchpoint-provider/internal/testutil"
	"catchpoint-provider/internal/types"
)

// The difference between alert subtypes for monitors is only in the Timing (Response Time) alert type.
// However, none of these types were available in the previous version of the provider, so we will
// need to add them later.
// While Transaction is very similar to Web, Web also has DaysToExpiration which Transaction does not have.
var transactionMobileValidSubtypes = []string{
	/* TODO: Not defined yet...
	types.CumulativeLayoutShift,
	types.FirstContentfulPaint,
	types.FirstPaint,
	types.LargestContentfulPaint,*/
}

var transactionEmulatedValidSubtypes = []string{
	// TODO: DataTransferTime not fully supported yet.,
}

var transactionChromeValidSubtypes = []string{
	/* TODO: Not defined yet...
	types.CumulativeLayoutShift,
	types.FirstContentfulPaint,
	types.FirstPaint,
	types.LargestContentfulPaint,
	types.DataTransferTime, */
}

func TestGetMonitorAlertTypesForTransaction(t *testing.T) {
	test := []struct {
		testMonitor string
	}{
		{types.ChromeString},
		{types.EmulatedString},
		{types.MobileString},
	}
	for _, tt := range test {
		testTransactionAlertMatrix := GetMonitorAlertTypes(types.TransactionType, tt.testMonitor)

		// The overall alert types for all transaction monitors are the same.
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

		testutil.AssertNotNil(t, "TransactionAlertTypeSchema", testTransactionAlertMatrix)
		testutil.AssertElementsMatch(t, "TransactionAlertTypeSchema", testTransactionAlertMatrix.ListAllValidAlertTypes(), expectedAlertTypes)
	}
}

func TestGetTestCompatibilityMatrixForTransaction(t *testing.T) {
	testTransactionAlertMatrix := GetTestCompatibilityMatrix(types.TransactionType)

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

	testutil.AssertNotNil(t, "TransactionAlertTypeSchema", testTransactionAlertMatrix)
	testutil.AssertElementsMatch(t, "TransactionAlertTypeSchema", testTransactionAlertMatrix.ListAllValidAlertTypes(), expectedAlertTypes)
}

func TestTransactionByteLengthSubTypes(t *testing.T) {
	test := []struct {
		testMonitor string
	}{
		{types.ChromeString},
		{types.EmulatedString},
		{types.MobileString},
	}
	for _, tt := range test {
		testTransactionAlertMatrix := GetMonitorAlertTypes(types.TransactionType, tt.testMonitor)

		testutil.AssertEqual(t, "TransactionByteLengthSubTypes", testTransactionAlertMatrix.IsAlertTypeValid(types.ByteLength), true)
		testutil.AssertElementsMatch(t, "TransactionByteLengthSubTypes", testTransactionAlertMatrix.AlertTypes[types.ByteLength].SubTypes, byteLengthValidSubtypes)
	}
}

func TestTransactionAddressSubTypes(t *testing.T) {
	test := []struct {
		testMonitor string
	}{
		{types.ChromeString},
		{types.EmulatedString},
		{types.MobileString},
	}
	for _, tt := range test {
		testTransactionAlertMatrix := GetMonitorAlertTypes(types.TransactionType, tt.testMonitor)

		testutil.AssertEqual(t, "TransactionAddressSubTypes", testTransactionAlertMatrix.IsAlertTypeValid(types.Address), true)
		testutil.AssertElementsMatch(t, "TransactionAddressSubTypes", testTransactionAlertMatrix.AlertTypes[types.Address].SubTypes, addressValidSubtypes)
	}
}

func TestTransactionInsightSubTypes(t *testing.T) {
	test := []struct {
		testMonitor string
	}{
		{types.ChromeString},
		{types.EmulatedString},
		{types.MobileString},
	}
	for _, tt := range test {
		testTransactionAlertMatrix := GetMonitorAlertTypes(types.TransactionType, tt.testMonitor)

		testutil.AssertEqual(t, "TransactionInsightSubTypes", testTransactionAlertMatrix.IsAlertTypeValid(types.Insight), true)
		testutil.AssertElementsMatch(t, "TransactionInsightSubTypes", testTransactionAlertMatrix.AlertTypes[types.Insight].SubTypes, insightValidSubtypes)
	}
}

func TestTransactionPingSubTypes(t *testing.T) {
	test := []struct {
		testMonitor string
	}{
		{types.ChromeString},
		{types.EmulatedString},
		{types.MobileString},
	}
	for _, tt := range test {
		testTransactionAlertMatrix := GetMonitorAlertTypes(types.TransactionType, tt.testMonitor)

		testutil.AssertEqual(t, "TransactionPingSubTypes", testTransactionAlertMatrix.IsAlertTypeValid(types.Ping), true)
		testutil.AssertElementsMatch(t, "TransactionPingSubTypes", testTransactionAlertMatrix.AlertTypes[types.Ping].SubTypes, pingValidSubtypes)
	}
}

func TestTransactionTimingSubTypes(t *testing.T) {
	test := []struct {
		testMonitor     string
		additionalTypes []string
	}{
		{types.ChromeString, transactionChromeValidSubtypes},
		{types.EmulatedString, transactionEmulatedValidSubtypes},
		{types.MobileString, transactionMobileValidSubtypes},
	}
	for _, tt := range test {
		testTransactionAlertMatrix := GetMonitorAlertTypes(types.TransactionType, tt.testMonitor)

		TransactionTimingSubTypes := []string{
			types.Connect,
			types.ContentLoad,
			types.DNS,
			types.DocumentComplete,
			types.DOMLoad,
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

		testutil.AssertEqual(t, "TransactionTimingSubTypes", testTransactionAlertMatrix.IsAlertTypeValid(types.Timing), true)
		testutil.AssertElementsMatch(t, "TransactionTimingSubTypes", testTransactionAlertMatrix.AlertTypes[types.Timing].SubTypes, TransactionTimingSubTypes)
	}
}

func TestTransactionAvailabilitySubTypes(t *testing.T) {
	test := []struct {
		testMonitor string
	}{
		{types.ChromeString},
		{types.EmulatedString},
		{types.MobileString},
	}
	for _, tt := range test {
		testTransactionAlertMatrix := GetMonitorAlertTypes(types.TransactionType, tt.testMonitor)

		testutil.AssertEqual(t, "TransactionAvailabilitySubTypes", testTransactionAlertMatrix.IsAlertTypeValid(types.Availability), true)
		testutil.AssertElementsMatch(t, "TransactionAvailabilitySubTypes", testTransactionAlertMatrix.AlertTypes[types.Availability].SubTypes, webAvailabilityValidSubtypes)
	}
}

func TestTransactionContentMatchSubTypes(t *testing.T) {
	test := []struct {
		testMonitor string
	}{
		{types.ChromeString},
		{types.EmulatedString},
		{types.MobileString},
	}
	for _, tt := range test {
		testTransactionAlertMatrix := GetMonitorAlertTypes(types.TransactionType, tt.testMonitor)

		testutil.AssertEqual(t, "TransactionContentMatchSubTypes", testTransactionAlertMatrix.IsAlertTypeValid(types.ContentMatch), true)
		testutil.AssertElementsMatch(t, "TransactionContentMatchSubTypes", testTransactionAlertMatrix.AlertTypes[types.ContentMatch].SubTypes, basicContentMatchValidSubtypes)
	}
}

func TestTransactionRequestsSubTypes(t *testing.T) {
	test := []struct {
		testMonitor string
	}{
		{types.ChromeString},
		{types.EmulatedString},
		{types.MobileString},
	}
	for _, tt := range test {
		testTransactionAlertMatrix := GetMonitorAlertTypes(types.TransactionType, tt.testMonitor)

		testutil.AssertEqual(t, "TransactionRequestsSubTypes", testTransactionAlertMatrix.IsAlertTypeValid(types.Requests), true)
		testutil.AssertElementsMatch(t, "TransactionRequestsSubTypes", testTransactionAlertMatrix.AlertTypes[types.Requests].SubTypes, requestsValidSubtypes)
	}
}
