package alerttypes

import (
	"testing"

	"catchpoint-provider/internal/testutil"
	"catchpoint-provider/internal/types"
)

func TestGetMonitorAlertTypesForDNS(t *testing.T) {
	// Both DNS Direct and DNS Experience have the same alert types.
	testDNSAlertMatrix := GetMonitorAlertTypes(types.DNSType, types.DNSDirectString)

	expectedAlertTypes := []string{
		types.Availability,
		types.DNS,
		types.ExperienceScore, // No subtypes
		types.Address,
		types.Ping,
		types.Timing,
		types.TestFailure, // No subtypes
	}

	testutil.AssertNotNil(t, "DNSAlertTypeSchema", testDNSAlertMatrix)
	testutil.AssertElementsMatch(t, "DNSAlertTypeSchema", testDNSAlertMatrix.ListAllValidAlertTypes(), expectedAlertTypes)
}

func TestGetTestCompatibilityMatrixForDNS(t *testing.T) {
	// The DNS test type only has one monitor type, which is DNS.
	testDNSAlertMatrix := GetTestCompatibilityMatrix(types.DNSType)

	expectedAlertTypes := []string{
		types.Availability,
		types.DNS,
		types.ExperienceScore, // No subtypes
		types.Address,         // No subtypes
		types.Ping,
		types.Timing,
		types.TestFailure, // No subtypes
	}

	testutil.AssertNotNil(t, "DNSAlertTypeSchema", testDNSAlertMatrix)
	testutil.AssertElementsMatch(t, "DNSAlertTypeSchema", testDNSAlertMatrix.ListAllValidAlertTypes(), expectedAlertTypes)
}

func TestDNSPingSubTypes(t *testing.T) {
	testDNSAlertMatrix := GetMonitorAlertTypes(types.DNSType, types.DNSDirectString)

	testutil.AssertEqual(t, "DNSPingSubTypes", testDNSAlertMatrix.IsAlertTypeValid(types.Ping), true)
	testutil.AssertElementsMatch(t, "DNSPingSubTypes", testDNSAlertMatrix.AlertTypes[types.Ping].SubTypes, pingValidSubtypes)
}

func TestDNSAvailabilitySubTypes(t *testing.T) {
	testDNSAlertMatrix := GetMonitorAlertTypes(types.DNSType, types.DNSDirectString)

	testutil.AssertEqual(t, "DNSAvailabilitySubTypes", testDNSAlertMatrix.IsAlertTypeValid(types.Availability), true)
	testutil.AssertElementsMatch(t, "DNSAvailabilitySubTypes", testDNSAlertMatrix.AlertTypes[types.Availability].SubTypes, availabilityValidSubtypes)
}

func TestDNSDNSSubTypes(t *testing.T) {
	testDNSAlertMatrix := GetMonitorAlertTypes(types.DNSType, types.DNSDirectString)

	testutil.AssertEqual(t, "DNSDNSSubTypes", testDNSAlertMatrix.IsAlertTypeValid(types.DNS), true)
	testutil.AssertElementsMatch(t, "DNSDNSSubTypes", testDNSAlertMatrix.AlertTypes[types.DNS].SubTypes, []string{types.DNSAdditional, types.DNSAnswer, types.DNSAuthority, types.DNSGeneral})
}
