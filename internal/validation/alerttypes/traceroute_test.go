package alerttypes

import (
	"testing"

	"catchpoint-provider/internal/testutil"
	"catchpoint-provider/internal/types"
)

func TestGetMonitorAlertTypesForTraceroute(t *testing.T) {
	test := []struct {
		testMonitor string
	}{
		{types.TracerouteICMPString},
		{types.TracerouteUDPString},
		{types.TracerouteTCPString},
		// {types.TracerouteQUICString}, // TODO: Add QUIC when it becomes available.
		// {types.TracerouteInSessionString}, // TODO: Add InSession when it becomes available.
	}
	for _, tt := range test {
		testTracerouteAlertMatrix := GetMonitorAlertTypes(types.TracerouteType, tt.testMonitor)

		expectedAlertTypes := []string{
			types.ASN,
			types.Availability,
			types.ExperienceScore,
			types.Address,
			types.Path,
			types.Ping,
			types.TestFailure,
		}

		testutil.AssertNotNil(t, "TracerouteAlertTypeSchema", testTracerouteAlertMatrix)
		testutil.AssertElementsMatch(t, "TracerouteAlertTypeSchema", testTracerouteAlertMatrix.ListAllValidAlertTypes(), expectedAlertTypes)
	}
}

func TestGetTestCompatibilityMatrixForTraceroute(t *testing.T) {
	testTracerouteAlertMatrix := GetTestCompatibilityMatrix(types.TracerouteType)

	expectedAlertTypes := []string{
		types.ASN,
		types.Availability,
		types.ExperienceScore,
		types.Address,
		types.Path,
		types.Ping,
		types.TestFailure,
	}

	testutil.AssertNotNil(t, "TracerouteAlertTypeSchema", testTracerouteAlertMatrix)
	testutil.AssertElementsMatch(t, "TracerouteAlertTypeSchema", testTracerouteAlertMatrix.ListAllValidAlertTypes(), expectedAlertTypes)
}

func TestTracerouteASNSubTypes(t *testing.T) {
	test := []struct {
		testMonitor string
	}{
		{types.TracerouteICMPString},
		{types.TracerouteUDPString},
		{types.TracerouteTCPString},
		// {types.TracerouteQUICString}, // TODO: Add QUIC when it becomes available.
		// {types.TracerouteInSessionString}, // TODO: Add InSession when it becomes available.
	}
	for _, tt := range test {
		testTracerouteAlertMatrix := GetMonitorAlertTypes(types.TracerouteType, tt.testMonitor)

		TracerouteASNSubTypes := []string{
			types.OriginAS,
			types.OriginNeighbor,
			types.PathAS,
			types.PrefixMismatch,
		}

		testutil.AssertEqual(t, "TracerouteASNSubTypes", testTracerouteAlertMatrix.IsAlertTypeValid(types.ASN), true)
		testutil.AssertElementsMatch(t, "TracerouteASNSubTypes", testTracerouteAlertMatrix.AlertTypes[types.ASN].SubTypes, TracerouteASNSubTypes)
	}
}

func TestTracerouteAvailabilitySubTypes(t *testing.T) {
	test := []struct {
		testMonitor string
	}{
		{types.TracerouteICMPString},
		{types.TracerouteUDPString},
		{types.TracerouteTCPString},
		// {types.TracerouteQUICString}, // TODO: Add QUIC when it becomes available.
		// {types.TracerouteInSessionString}, // TODO: Add InSession when it becomes available.
	}
	for _, tt := range test {
		testTracerouteAlertMatrix := GetMonitorAlertTypes(types.TracerouteType, tt.testMonitor)

		TracerouteAvailabilitySubTypes := []string{
			types.DowntimePercent,
			types.Test,
		}

		testutil.AssertEqual(t, "TracerouteAvailabilitySubTypes", testTracerouteAlertMatrix.IsAlertTypeValid(types.Availability), true)
		testutil.AssertElementsMatch(t, "TracerouteAvailabilitySubTypes", testTracerouteAlertMatrix.AlertTypes[types.Availability].SubTypes, TracerouteAvailabilitySubTypes)
	}
}

func TestTraceroutePathSubTypes(t *testing.T) {
	testTracerouteAlertMatrix := GetMonitorAlertTypes(types.TracerouteType, types.TracerouteICMPString)

	PathSubTypes := []string{
		types.ASNsNum,
		types.CitiesNum,
		types.CountriesNum,
		types.HopsNum,
	}

	testutil.AssertEqual(t, "TraceroutePathSubTypes", testTracerouteAlertMatrix.IsAlertTypeValid(types.Path), true)
	testutil.AssertElementsMatch(t, "TraceroutePathSubTypes", testTracerouteAlertMatrix.AlertTypes[types.Path].SubTypes, PathSubTypes)
}

func TestTraceroutePingSubTypes(t *testing.T) {
	testTracerouteAlertMatrix := GetMonitorAlertTypes(types.TracerouteType, types.TracerouteICMPString)

	PingSubTypes := []string{
		types.PingPacketLoss,
		types.PingRTT,
		// TODO: add Jitter when it becomes available.
	}

	testutil.AssertEqual(t, "TraceroutePingSubTypes", testTracerouteAlertMatrix.IsAlertTypeValid(types.Ping), true)
	testutil.AssertElementsMatch(t, "TraceroutePingSubTypes", testTracerouteAlertMatrix.AlertTypes[types.Ping].SubTypes, PingSubTypes)
}
