package alerttypes

import (
	"testing"

	"catchpoint-provider/internal/testutil"
	"catchpoint-provider/internal/types"
)

func TestGetMonitorAlertTypesForICMPPing(t *testing.T) {
	testPingAlertMatrix := GetMonitorAlertTypes(types.PingType, types.PingICMPString)

	expectedAlertTypes := []string{
		types.Availability,
		types.ExperienceScore,
		types.Address,
		types.Ping,
		types.Timing,
		types.TestFailure,
	}

	testutil.AssertNotNil(t, "PingAlertTypeSchema", testPingAlertMatrix)
	testutil.AssertElementsMatch(t, "PingAlertTypeSchema", testPingAlertMatrix.ListAllValidAlertTypes(), expectedAlertTypes)
}

func TestGetMonitorAlertTypesForUDPPing(t *testing.T) {
	testPingAlertMatrix := GetMonitorAlertTypes(types.PingType, types.PingUDPString)

	expectedAlertTypes := []string{
		types.Availability,
		types.ExperienceScore,
		types.Address,
		types.Ping,
		types.Timing,
		types.TestFailure,
	}

	testutil.AssertNotNil(t, "PingAlertTypeSchema", testPingAlertMatrix)
	testutil.AssertElementsMatch(t, "PingAlertTypeSchema", testPingAlertMatrix.ListAllValidAlertTypes(), expectedAlertTypes)
}

func TestGetMonitorAlertTypesForTCPPing(t *testing.T) {
	testPingAlertMatrix := GetMonitorAlertTypes(types.PingType, types.PingTCPString)

	expectedAlertTypes := []string{
		types.Availability,
		types.ExperienceScore,
		types.Address,
		types.Ping,
		types.Timing,
		types.TestFailure,
	}

	testutil.AssertNotNil(t, "PingAlertTypeSchema", testPingAlertMatrix)
	testutil.AssertElementsMatch(t, "PingAlertTypeSchema", testPingAlertMatrix.ListAllValidAlertTypes(), expectedAlertTypes)
}

func TestGetTestCompatibilityMatrixForPing(t *testing.T) {
	testPingAlertMatrix := GetTestCompatibilityMatrix(types.PingType)

	expectedAlertTypes := []string{
		types.Availability,
		types.ExperienceScore,
		types.Address,
		types.Ping,
		types.Timing,
		types.TestFailure,
	}

	testutil.AssertNotNil(t, "PingAlertTypeSchema", testPingAlertMatrix)
	testutil.AssertElementsMatch(t, "PingAlertTypeSchema", testPingAlertMatrix.ListAllValidAlertTypes(), expectedAlertTypes)
}

func TestPingAvailabilitySubTypes(t *testing.T) {
	testPingAlertMatrix := GetMonitorAlertTypes(types.PingType, types.PingICMPString)

	PingAvailabilitySubTypes := []string{
		types.DowntimePercent,
		types.Test,
	}

	testutil.AssertEqual(t, "PingAvailabilitySubTypes", testPingAlertMatrix.IsAlertTypeValid(types.Availability), true)
	testutil.AssertElementsMatch(t, "PingAvailabilitySubTypes", testPingAlertMatrix.AlertTypes[types.Availability].SubTypes, PingAvailabilitySubTypes)
}

func TestPingWithPingSubTypes(t *testing.T) {
	testPingAlertMatrix := GetMonitorAlertTypes(types.PingType, types.PingICMPString)

	PingSubTypes := []string{
		types.PingPacketLoss,
		types.PingRTT,
	}

	testutil.AssertEqual(t, "PingWithPingSubTypes", testPingAlertMatrix.IsAlertTypeValid(types.Ping), true)
	testutil.AssertElementsMatch(t, "PingWithPingSubTypes", testPingAlertMatrix.AlertTypes[types.Ping].SubTypes, PingSubTypes)
}
