package alerttypes

import (
	"testing"

	"catchpoint-provider/internal/testutil"
	"catchpoint-provider/internal/types"
)

func TestGetMonitorAlertTypesForSSL(t *testing.T) {
	// Both SSL Direct and SSL Experience have the same alert types.
	testSSLAlertMatrix := GetMonitorAlertTypes(types.SSLType, types.SSLString)

	expectedAlertTypes := []string{
		types.ContentMatch,
		types.ExperienceScore,
		types.Timing,
		types.TestFailure,
	}

	testutil.AssertNotNil(t, "SSLAlertTypeSchema", testSSLAlertMatrix)
	testutil.AssertElementsMatch(t, "SSLAlertTypeSchema", testSSLAlertMatrix.ListAllValidAlertTypes(), expectedAlertTypes)
}

func TestGetTestCompatibilityMatrixForSSL(t *testing.T) {
	// The SSL test type only has one monitor type, which is SSL.
	testSSLAlertMatrix := GetTestCompatibilityMatrix(types.SSLType)

	expectedAlertTypes := []string{
		types.ContentMatch,
		types.ExperienceScore,
		types.Timing,
		types.TestFailure,
	}

	testutil.AssertNotNil(t, "SSLAlertTypeSchema", testSSLAlertMatrix)
	testutil.AssertElementsMatch(t, "SSLAlertTypeSchema", testSSLAlertMatrix.ListAllValidAlertTypes(), expectedAlertTypes)
}

func TestSSLTimingSubTypes(t *testing.T) {
	testSSLAlertMatrix := GetMonitorAlertTypes(types.SSLType, types.SSLString)

	sslTimingSubTypes := []string{
		types.Connect,
		types.DaysToExpiration,
		types.HandshakeTime,
	}

	testutil.AssertEqual(t, "SSLTimingSubTypes", testSSLAlertMatrix.IsAlertTypeValid(types.Timing), true)
	testutil.AssertElementsMatch(t, "SSLTimingSubTypes", testSSLAlertMatrix.AlertTypes[types.Timing].SubTypes, sslTimingSubTypes)
}
