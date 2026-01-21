package alerttypes

import (
	"testing"

	"catchpoint-provider/internal/testutil"
	"catchpoint-provider/internal/types"
)

func TestGetMonitorAlertTypesForAPI(t *testing.T) {
	// The API test type only has one monitor type, which is API.
	testAPIAlertMatrix := GetMonitorAlertTypes(types.APIType, types.APIString)

	expectedAlertTypes := []string{
		types.ByteLength,
		types.HostFailure, // No subtypes
		types.Address,
		types.Insight,
		types.Ping,
		types.Timing,
		types.Availability,
		types.ContentMatch,
		types.ExperienceScore, // No subtypes
		types.Requests,
		types.TestFailure, // No subtypes
	}

	testutil.AssertNotNil(t, "APIAlertTypeSchema", testAPIAlertMatrix)
	testutil.AssertElementsMatch(t, "APIAlertTypeSchema", testAPIAlertMatrix.ListAllValidAlertTypes(), expectedAlertTypes)
}

func TestGetTestCompatibilityMatrixForAPI(t *testing.T) {
	testAPIAlertMatrix := GetTestCompatibilityMatrix(types.APIType)

	expectedAlertTypes := []string{
		types.ByteLength,
		types.HostFailure, // No subtypes
		types.Address,
		types.Insight,
		types.Ping,
		types.Timing,
		types.Availability,
		types.ContentMatch,
		types.ExperienceScore, // No subtypes
		types.Requests,
		types.TestFailure, // No subtypes
	}

	testutil.AssertNotNil(t, "APIAlertTypeSchema", testAPIAlertMatrix)
	testutil.AssertElementsMatch(t, "APIAlertTypeSchema", testAPIAlertMatrix.ListAllValidAlertTypes(), expectedAlertTypes)
}

func TestAPIByteLengthSubTypes(t *testing.T) {
	testAPIAlertMatrix := GetMonitorAlertTypes(types.APIType, types.APIString)

	testutil.AssertEqual(t, "APIByteLengthSubTypes", testAPIAlertMatrix.IsAlertTypeValid(types.ByteLength), true)
	testutil.AssertElementsMatch(t, "APIByteLengthSubTypes", testAPIAlertMatrix.AlertTypes[types.ByteLength].SubTypes, byteLengthValidSubtypes)
}

func TestAPIAddressSubTypes(t *testing.T) {
	testAPIAlertMatrix := GetMonitorAlertTypes(types.APIType, types.APIString)

	testutil.AssertEqual(t, "APIAddressSubTypes", testAPIAlertMatrix.IsAlertTypeValid(types.Address), true)
	testutil.AssertElementsMatch(t, "APIAddressSubTypes", testAPIAlertMatrix.AlertTypes[types.Address].SubTypes, addressValidSubtypes)
}

func TestAPIInsightSubTypes(t *testing.T) {
	testAPIAlertMatrix := GetMonitorAlertTypes(types.APIType, types.APIString)

	testutil.AssertEqual(t, "APIInsightSubTypes", testAPIAlertMatrix.IsAlertTypeValid(types.Insight), true)
	testutil.AssertElementsMatch(t, "APIInsightSubTypes", testAPIAlertMatrix.AlertTypes[types.Insight].SubTypes, insightValidSubtypes)
}

func TestAPIPingSubTypes(t *testing.T) {
	testAPIAlertMatrix := GetMonitorAlertTypes(types.APIType, types.APIString)

	testutil.AssertEqual(t, "APIPingSubTypes", testAPIAlertMatrix.IsAlertTypeValid(types.Ping), true)
	testutil.AssertElementsMatch(t, "APIPingSubTypes", testAPIAlertMatrix.AlertTypes[types.Ping].SubTypes, pingValidSubtypes)
}

func TestAPITimingSubTypes(t *testing.T) {
	testAPIAlertMatrix := GetMonitorAlertTypes(types.APIType, types.APIString)

	apiTimingSubTypes := []string{
		types.Connect,
		types.ContentLoad,
		types.DNS,
		// TODO: DataTransferTime not fully supported yet.,
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

	testutil.AssertEqual(t, "APITimingSubTypes", testAPIAlertMatrix.IsAlertTypeValid(types.Timing), true)
	testutil.AssertElementsMatch(t, "APITimingSubTypes", testAPIAlertMatrix.AlertTypes[types.Timing].SubTypes, apiTimingSubTypes)
}

func TestAPIAvailabilitySubTypes(t *testing.T) {
	testAPIAlertMatrix := GetMonitorAlertTypes(types.APIType, types.APIString)

	testutil.AssertEqual(t, "APIAvailabilitySubTypes", testAPIAlertMatrix.IsAlertTypeValid(types.Availability), true)
	testutil.AssertElementsMatch(t, "APIAvailabilitySubTypes", testAPIAlertMatrix.AlertTypes[types.Availability].SubTypes, webAvailabilityValidSubtypes)
}

func TestAPIContentMatchSubTypes(t *testing.T) {
	testAPIAlertMatrix := GetMonitorAlertTypes(types.APIType, types.APIString)

	testutil.AssertEqual(t, "APIContentMatchSubTypes", testAPIAlertMatrix.IsAlertTypeValid(types.ContentMatch), true)
	testutil.AssertElementsMatch(t, "APIContentMatchSubTypes", testAPIAlertMatrix.AlertTypes[types.ContentMatch].SubTypes, basicContentMatchValidSubtypes)
}

func TestAPIRequestsSubTypes(t *testing.T) {
	testAPIAlertMatrix := GetMonitorAlertTypes(types.APIType, types.APIString)

	testutil.AssertEqual(t, "APIRequestsSubTypes", testAPIAlertMatrix.IsAlertTypeValid(types.Requests), true)
	testutil.AssertElementsMatch(t, "APIRequestsSubTypes", testAPIAlertMatrix.AlertTypes[types.Requests].SubTypes, requestsValidSubtypes)
}
