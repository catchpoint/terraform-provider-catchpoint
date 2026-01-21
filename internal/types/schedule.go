package types

var (
	ValidNodeDistributions []string
	ValidFrequencyNames    []string
	frequencyNames         map[int]string
	nodeDistributionNames  map[int]string
	nodeThresholdTypeNames map[int]string
)

// #region FrequencyTypes

var FrequencyIDs = map[string]int{
	None:                    0,
	string(OneMinute):       1,
	string(FiveMinutes):     2,
	string(TenMinutes):      3,
	string(FifteenMinutes):  4,
	string(TwentyMinutes):   5,
	string(ThirtyMinutes):   6,
	string(SixtyMinutes):    7,
	string(TwoHours):        8,
	string(ThreeHours):      9,
	string(FourHours):       10,
	string(SixHours):        11,
	string(EightHours):      12,
	string(TwelveHours):     13,
	string(TwentyFourHours): 14,
	string(FourMinutes):     15,
	string(TwoMinutes):      16,
}

// ValidFrequencyValues contains all valid frequency values. This was manually created because
// even if you sort the map by ID, TwoMinutes and FourMinutes end up at the end.
// This is a better presentation for error messages.
var ValidFrequencyValues = []string{
	None,
	string(OneMinute),
	string(TwoMinutes),
	string(FourMinutes),
	string(FiveMinutes),
	string(TenMinutes),
	string(FifteenMinutes),
	string(TwentyMinutes),
	string(ThirtyMinutes),
	string(SixtyMinutes),
	string(TwoHours),
	string(ThreeHours),
	string(FourHours),
	string(SixHours),
	string(EightHours),
	string(TwelveHours),
	string(TwentyFourHours),
}

func GetFrequencyID(name string) (int, bool) {
	id, ok := FrequencyIDs[name]
	return id, ok
}

func GetFrequencyName(id int) (string, bool) {
	name, ok := frequencyNames[id]
	return name, ok
}

// #endregion

// #region NodeDistributionTypes

var NodeDistributionIDs = map[string]int{
	Random:     0,
	Concurrent: 1,
}

func GetNodeDistributionID(name string) (int, bool) {
	id, ok := NodeDistributionIDs[name]
	return id, ok
}

func GetNodeDistributionName(id int) (string, bool) {
	name, ok := nodeDistributionNames[id]
	return name, ok
}

// #endregion

func init() {
	frequencyNames = populateReverseMap(FrequencyIDs)
	ValidFrequencyNames = populateValidNames(FrequencyIDs)

	nodeDistributionNames = populateReverseMap(NodeDistributionIDs)
	ValidNodeDistributions = populateValidNames(NodeDistributionIDs)
}
