package types

import (
	"catchpoint-provider/internal/testutil"
	"testing"
)

func TestGetFrequencyIDAndName(t *testing.T) {
	tests := []struct {
		name string
		id   int
	}{
		{None, 0},
		{string(OneMinute), 1},
		{string(FiveMinutes), 2},
		{string(TenMinutes), 3},
		{string(FifteenMinutes), 4},
		{string(TwentyMinutes), 5},
		{string(ThirtyMinutes), 6},
		{string(SixtyMinutes), 7},
		{string(TwoHours), 8},
		{string(ThreeHours), 9},
		{string(FourHours), 10},
		{string(SixHours), 11},
		{string(EightHours), 12},
		{string(TwelveHours), 13},
		{string(TwentyFourHours), 14},
		{string(FourMinutes), 15},
		{string(TwoMinutes), 16},
	}

	for _, tt := range tests {
		id := testutil.GetOrFail(t, func() (int, bool) { return GetFrequencyID(tt.name) }, "GetFrequencyID", tt.name)
		name := testutil.GetOrFail(t, func() (string, bool) { return GetFrequencyName(id) }, "GetFrequencyName", id)

		testutil.AssertEqual(t, "frequencyID", id, tt.id)
		testutil.AssertEqual(t, "frequencyName", name, tt.name)
	}
}

func TestGetNodeDistributionIDAndName(t *testing.T) {
	tests := []struct {
		name string
		id   int
	}{
		{Random, 0},
		{Concurrent, 1},
	}

	for _, tt := range tests {
		id := testutil.GetOrFail(t, func() (int, bool) { return GetNodeDistributionID(tt.name) }, "GetNodeDistributionID", tt.name)
		name := testutil.GetOrFail(t, func() (string, bool) { return GetNodeDistributionName(id) }, "GetNodeDistributionName", id)

		testutil.AssertEqual(t, "nodeDistributionID", id, tt.id)
		testutil.AssertEqual(t, "nodeDistributionName", name, tt.name)
	}
}

func TestGetNodeThresholdTypeIDAndName(t *testing.T) {
	tests := []struct {
		name string
		id   int
	}{
		{Runs, 0},
		{AverageAcrossNodes, 1},
		{Node, 2},
	}

	for _, tt := range tests {
		id := testutil.GetOrFail(t, func() (int, bool) { return GetNodeThresholdTypeID(tt.name) }, "GetNodeThresholdTypeID", tt.name)
		name := testutil.GetOrFail(t, func() (string, bool) { return GetNodeThresholdTypeName(id) }, "GetNodeThresholdTypeName", id)
		testutil.AssertEqual(t, "nodeThresholdTypeID", id, tt.id)
		testutil.AssertEqual(t, "nodeThresholdTypeName", name, tt.name)
	}
}
