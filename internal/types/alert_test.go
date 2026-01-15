package types

import (
	"catchpoint-provider/internal/testutil"
	"testing"
)

func TestGetNotificationTypeIDAndName(t *testing.T) {
	test := []struct {
		name string
		id   int
	}{
		{DefaultContacts, 0},
	}
	for _, tt := range test {
		id := testutil.GetOrFail(t, func() (int, bool) { return GetNotificationTypeID(tt.name) }, "GetNotificationTypeID", tt.name)
		name := testutil.GetOrFail(t, func() (string, bool) { return GetNotificationTypeName(id) }, "GetNotificationTypeName", id)

		testutil.AssertEqual(t, "notificationTypeID", id, tt.id)
		testutil.AssertEqual(t, "notificationTypeName", name, tt.name)
	}
}

func TestGetAlertTypeIDAndName(t *testing.T) {
	test := []struct {
		name string
		id   int
	}{
		{ByteLength, 2},
		{ContentMatch, 3},
		{HostFailure, 4},
		{TestFailure, 9},
		{Timing, 7},
		{Ping, 12},
		{Requests, 13},
		{Availability, 15},
		{DNS, 17},
		{Path, 20},
		{ASN, 23},
	}
	for _, tt := range test {
		id := testutil.GetOrFail(t, func() (int, bool) { return GetAlertTypeID(tt.name) }, "GetAlertTypeID", tt.name)
		name := testutil.GetOrFail(t, func() (string, bool) { return GetAlertTypeName(id) }, "GetAlertTypeName", id)

		testutil.AssertEqual(t, "alertTypeID", id, tt.id)
		testutil.AssertEqual(t, "alertTypeName", name, tt.name)
	}
}

func TestGetAlertSubTypeIDAndName(t *testing.T) {
	test := []struct {
		name string
		id   int
	}{
		{ByteLength, 1},
		{Page, 2},
		{FileSize, 3},
		{RegularExpression, 10},
		{ResponseCode, 14},
		{ResponseHeaders, 15},
		{DNS, 50},
		{Connect, 51},
		{Send, 52},
		{Wait, 53},
		{Load, 54},
		{TTFB, 55},
		{ContentLoad, 57},
		{Response, 58},
		{TestTime, 59},
		{DOMLoad, 61},
		{TestTimeWithSuspect, 63},
		{ServerResponse, 64},
		{DocumentComplete, 66},
		{Redirect, 67},
		{PingRTT, 100},
		{PingPacketLoss, 101},
		{RequestsNum, 110},
		{HostsNum, 111},
		{ConnectionsNum, 112},
		{RedirectsNum, 113},
		{OtherNum, 114},
		{ImagesNum, 115},
		{ScriptsNum, 116},
		{HTMLNum, 117},
		{CSSNum, 118},
		{XMLNum, 119},
		{FlashNum, 120},
		{MediaNum, 121},
		{Test, 140},
		{Content, 141},
		{DowntimePercent, 142},
		{DNSAnswer, 161},
		{CitiesNum, 190},
		{ASNsNum, 191},
		{CountriesNum, 193},
		{HopsNum, 194},
		{HandshakeTime, 195},
		{DaysToExpiration, 196},
		{OriginAS, 210},
		{PathAS, 211},
		{OriginNeighbor, 212},
		{PrefixMismatch, 213},
	}
	for _, tt := range test {
		id := testutil.GetOrFail(t, func() (int, bool) { return GetAlertSubTypeID(tt.name) }, "GetAlertSubTypeID", tt.name)
		name := testutil.GetOrFail(t, func() (string, bool) { return GetAlertSubTypeName(id) }, "GetAlertSubTypeName", id)

		testutil.AssertEqual(t, "alertSubTypeID", id, tt.id)
		testutil.AssertEqual(t, "alertSubTypeName", name, tt.name)
	}
}

func TestGetAlertSettingTypeIDAndName(t *testing.T) {
	test := []struct {
		name string
		id   int
	}{
		{Inherit, 0},
		{Override, 1},
		{InheritAndAdd, 2},
	}
	for _, tt := range test {
		id := testutil.GetOrFail(t, func() (int, bool) { return GetAlertSettingTypeID(tt.name) }, "GetAlertSettingTypeID", tt.name)
		name := testutil.GetOrFail(t, func() (string, bool) { return GetAlertSettingTypeName(id) }, "GetAlertSettingTypeName", id)

		testutil.AssertEqual(t, "alertSettingTypeID", id, tt.id)
		testutil.AssertEqual(t, "alertSettingTypeName", name, tt.name)
	}
}

func TestGetOperationTypeIDAndName(t *testing.T) {
	test := []struct {
		name string
		id   int
	}{
		{Equals, 1},
		{NotEquals, 0},
		{GreaterThan, 2},
		{GreaterThanOrEquals, 3},
		{LessThan, 4},
		{LessThanOrEquals, 5},
	}
	for _, tt := range test {
		id := testutil.GetOrFail(t, func() (int, bool) { return GetOperationTypeID(tt.name) }, "GetOperationTypeID", tt.name)
		name := testutil.GetOrFail(t, func() (string, bool) { return GetOperationTypeName(id) }, "GetOperationTypeName", id)

		testutil.AssertEqual(t, "operationTypeID", id, tt.id)
		testutil.AssertEqual(t, "operationTypeName", name, tt.name)
	}
}

func TestGetTriggerTypeIDAndName(t *testing.T) {
	test := []struct {
		name string
		id   int
	}{
		{SpecificValue, 1},
		{TrailingValue, 2},
		{TrendShift, 3},
	}
	for _, tt := range test {
		id := testutil.GetOrFail(t, func() (int, bool) { return GetTriggerTypeID(tt.name) }, "GetTriggerTypeID", tt.name)
		name := testutil.GetOrFail(t, func() (string, bool) { return GetTriggerTypeName(id) }, "GetTriggerTypeName", id)

		testutil.AssertEqual(t, "triggerTypeID", id, tt.id)
		testutil.AssertEqual(t, "triggerTypeName", name, tt.name)
	}
}

func TestGetReminderIDAndName(t *testing.T) {
	test := []struct {
		name string
		id   int
	}{
		{None, 0},
		{string(OneMinute), 1},
		{string(FiveMinutes), 5},
		{string(TenMinutes), 10},
		{string(FifteenMinutes), 15},
		{string(ThirtyMinutes), 30},
		{string(OneHour), 60},
		{string(Daily), 1440},
	}
	for _, tt := range test {
		id := testutil.GetOrFail(t, func() (int, bool) { return GetReminderID(tt.name) }, "GetReminderID", tt.name)
		name := testutil.GetOrFail(t, func() (string, bool) { return GetReminderName(id) }, "GetReminderName", id)

		testutil.AssertEqual(t, "reminderID", id, tt.id)
		testutil.AssertEqual(t, "reminderName", name, tt.name)
	}
}

func TestGetThresholdIntervalIDAndName(t *testing.T) {
	test := []struct {
		name string
		id   int
	}{
		{Default, 0},
		{string(FiveMinutes), 5},
		{string(TenMinutes), 10},
		{string(FifteenMinutes), 15},
		{string(ThirtyMinutes), 30},
		{string(OneHour), 60},
		{string(TwoHours), 120},
		{string(SixHours), 360},
		{string(TwelveHours), 720},
	}
	for _, tt := range test {
		id := testutil.GetOrFail(t, func() (int, bool) { return GetThresholdIntervalID(tt.name) }, "GetThresholdIntervalID", tt.name)
		name := testutil.GetOrFail(t, func() (string, bool) { return GetThresholdIntervalName(id) }, "GetThresholdIntervalName", id)

		testutil.AssertEqual(t, "thresholdIntervalID", id, tt.id)
		testutil.AssertEqual(t, "thresholdIntervalName", name, tt.name)
	}
}

func TestGetHistoricalIntervalIDAndName(t *testing.T) {
	test := []struct {
		name string
		id   int
	}{
		{string(FiveMinutes), 5},
		{string(TenMinutes), 10},
		{string(FifteenMinutes), 15},
		{string(ThirtyMinutes), 30},
		{string(OneHour), 60},
		{string(TwoHours), 120},
		{string(SixHours), 360},
		{string(TwelveHours), 720},
		{string(OneDay), 1440},
		{string(OneWeek), 10080},
	}
	for _, tt := range test {
		id := testutil.GetOrFail(t, func() (int, bool) { return GetHistoricalIntervalID(tt.name) }, "GetHistoricalIntervalID", tt.name)
		name := testutil.GetOrFail(t, func() (string, bool) { return GetHistoricalIntervalName(id) }, "GetHistoricalIntervalName", id)

		testutil.AssertEqual(t, "historicalIntervalID", id, tt.id)
		testutil.AssertEqual(t, "historicalIntervalName", name, tt.name)
	}
}

func TestGetStatisticalTypeIDAndName(t *testing.T) {
	test := []struct {
		name string
		id   int
	}{
		{Average, 1},
	}
	for _, tt := range test {
		id := testutil.GetOrFail(t, func() (int, bool) { return GetStatisticalTypeID(tt.name) }, "GetStatisticalTypeID", tt.name)
		name := testutil.GetOrFail(t, func() (string, bool) { return GetStatisticalTypeName(id) }, "GetStatisticalTypeName", id)

		testutil.AssertEqual(t, "statisticalTypeID", id, tt.id)
		testutil.AssertEqual(t, "statisticalTypeName", name, tt.name)
	}
}

func TestGetFilterTypeID(t *testing.T) {
	test := []struct {
		name string
		id   int
	}{
		{Index, 1},
		{Name, 2},
		{Address, 3},
	}
	for _, tt := range test {
		id := testutil.GetOrFail(t, func() (int, bool) { return GetFilterTypeID(tt.name) }, "GetFilterTypeID", tt.name)

		testutil.AssertEqual(t, "filterTypeID", id, tt.id)
	}
}
