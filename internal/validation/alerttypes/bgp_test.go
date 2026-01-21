package alerttypes

import (
	"testing"

	"catchpoint-provider/internal/testutil"
	"catchpoint-provider/internal/types"
)

func TestGetMonitorAlertTypesForBGP(t *testing.T) {
	// Both BGP and BGP Basic have the same alert types.
	testBGPAlertMatrix := GetMonitorAlertTypes(types.BGPType, types.BGPString)

	expectedAlertTypes := []string{
		types.ASN,
		types.Availability,
		types.TestFailure, // No subtypes
	}

	testutil.AssertNotNil(t, "BGPAlertTypeSchema", testBGPAlertMatrix)
	testutil.AssertElementsMatch(t, "BGPAlertTypeSchema", testBGPAlertMatrix.ListAllValidAlertTypes(), expectedAlertTypes)
}

func TestGetTestCompatibilityMatrixForBGP(t *testing.T) {
	testBGPAlertMatrix := GetTestCompatibilityMatrix(types.BGPType)

	expectedAlertTypes := []string{
		types.ASN,
		types.Availability,
		types.TestFailure, // No subtypes
	}

	testutil.AssertNotNil(t, "BGPAlertTypeSchema", testBGPAlertMatrix)
	testutil.AssertElementsMatch(t, "BGPAlertTypeSchema", testBGPAlertMatrix.ListAllValidAlertTypes(), expectedAlertTypes)
}

func TestBGPASNSubTypes(t *testing.T) {
	testBGPAlertMatrix := GetMonitorAlertTypes(types.BGPType, types.BGPString)

	testutil.AssertEqual(t, "BGPASNSubTypes", testBGPAlertMatrix.IsAlertTypeValid(types.ASN), true)
	testutil.AssertElementsMatch(t, "BGPASNSubTypes", testBGPAlertMatrix.AlertTypes[types.ASN].SubTypes, asnValidSubtypes)
}

func TestBGPAvailabilitySubTypes(t *testing.T) {
	testBGPAlertMatrix := GetMonitorAlertTypes(types.BGPType, types.BGPString)

	testutil.AssertEqual(t, "BGPAvailabilitySubTypes", testBGPAlertMatrix.IsAlertTypeValid(types.Availability), true)
	testutil.AssertElementsMatch(t, "BGPAvailabilitySubTypes", testBGPAlertMatrix.AlertTypes[types.Availability].SubTypes, bgpAvailabilityValidSubtypes)
}
