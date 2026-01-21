package alerttypes

import (
	"catchpoint-provider/internal/testutil"
	"catchpoint-provider/internal/types"
	"testing"
)

const (
	fakeTestType types.TestType = 0
)

var (
	// Availability subtypes can be lumped into 3 types, standard(Traceroute, DNS, Ping), web, and BGP.
	availabilityValidSubtypes    = []string{types.DowntimePercent, types.Test}
	webAvailabilityValidSubtypes = []string{types.DowntimePercent, types.Content, types.Test}
	bgpAvailabilityValidSubtypes = []string{types.DowntimePercent, types.Reachability, types.Test}
	// ASN always has these four subtypes.
	asnValidSubtypes = []string{types.OriginAS, types.OriginNeighbor, types.PathAS, types.PrefixMismatch}
	// Where ByteLength is used, it always has the same subtypes.
	byteLengthValidSubtypes = []string{types.FileSize, types.Page, types.ByteLength}
	// Address always either has no subtypes or these three.
	addressValidSubtypes = []string{types.AnyURLs, types.ChildURLs, types.TestURLs}
	// Ping usually has PacketLoss and RTT but for Traceroute, also has Jitter.
	pingValidSubtypes = []string{types.PingPacketLoss, types.PingRTT}
	// Insight always has Indicators and Tracepoints.
	insightValidSubtypes = []string{types.Indicators, types.Tracepoints}

	// Requests always has these subtypes.
	requestsValidSubtypes = []string{
		types.ConnectionsNum,
		types.CSSNum,
		types.FlashNum,
		types.HostsNum,
		types.HTMLNum,
		types.ImagesNum,
		types.MediaNum,
		types.OtherNum,
		types.RedirectsNum,
		types.RequestsNum,
		types.ScriptsNum,
		types.XMLNum,
	}

	// ContentMatch either has one set of subtypes or the other.
	basicContentMatchValidSubtypes = []string{
		types.ResponseCode,
		types.ResponseHeaders,
		types.RegularExpression,
	}

	// This is only used for Web tests (minus API).
	extendedContentMatchValidSubtypes = []string{
		types.CompareToPrevious,
		types.CompareToFile,
		types.ResponseCode,
		types.ResponseHeaders,
		types.RegularExpression,
	}
)

var fakeAlertMatrix = &MonitorAlertTypeCompatibilityMatrix{
	TestType: fakeTestType,
	MonitorTypes: map[string]*MonitorAlertTypes{
		"fakeMonitorType1": {
			AlertTypes: map[string]*AlertSubTypes{
				"fakeAlertType":        {SubTypes: []string{"fakeAlertSubType"}},
				"anotherFakeAlertType": {SubTypes: []string{"anotherFakeAlertSubType"}},
			},
		},
	},
}

func TestValidateAlertSettings(t *testing.T) {
	orig := GetMonitorAlertTypesFunc
	defer func() { GetMonitorAlertTypesFunc = orig }()

	GetMonitorAlertTypesFunc = func(testType types.TestType, monitorType string) *MonitorAlertTypes {
		return fakeAlertMatrix.MonitorTypes[monitorType]
	}

	alertSettings := []map[string]any{
		{"alert_type": "fakeAlertType", "alert_sub_type": "fakeAlertSubType"},
	}

	err := ValidateTestAlertTypeCombination(alertSettings, fakeTestType, "fakeMonitorType1")
	testutil.AssertNil(t, "ValidateAlertSettings", err)
}

func TestValidateAlertSettingsInvalidMonitor(t *testing.T) {
	orig := GetMonitorAlertTypesFunc
	defer func() { GetMonitorAlertTypesFunc = orig }()

	GetMonitorAlertTypesFunc = func(testType types.TestType, monitorType string) *MonitorAlertTypes {
		return fakeAlertMatrix.MonitorTypes[monitorType]
	}

	alertSettings := []map[string]any{
		{"alert_type": "fakeAlertType", "alert_sub_type": "fakeAlertSubType"},
	}

	err := ValidateTestAlertTypeCombination(alertSettings, fakeTestType, "notAValidMonitor")
	testutil.AssertNotNil(t, "ValidateAlertSettings", err)
}

func TestValidateAlertSettingsInvalidAlertType(t *testing.T) {
	orig := GetMonitorAlertTypesFunc
	defer func() { GetMonitorAlertTypesFunc = orig }()

	GetMonitorAlertTypesFunc = func(testType types.TestType, monitorType string) *MonitorAlertTypes {
		return fakeAlertMatrix.MonitorTypes[monitorType]
	}

	alertSettings := []map[string]any{
		{"alert_type": "notAValidAlertType", "alert_sub_type": "fakeAlertSubType"},
	}

	err := ValidateTestAlertTypeCombination(alertSettings, fakeTestType, "fakeMonitorType1")
	testutil.AssertNotNil(t, "ValidateAlertSettings", err)
}

func TestValidateAlertSettingsInvalidAlertSubType(t *testing.T) {
	orig := GetMonitorAlertTypesFunc
	defer func() { GetMonitorAlertTypesFunc = orig }()

	GetMonitorAlertTypesFunc = func(testType types.TestType, monitorType string) *MonitorAlertTypes {
		return fakeAlertMatrix.MonitorTypes[monitorType]
	}

	alertSettings := []map[string]any{
		{"alert_type": "fakeAlertType", "alert_sub_type": "notAValidAlertSubType"},
	}

	err := ValidateTestAlertTypeCombination(alertSettings, fakeTestType, "fakeMonitorType1")
	testutil.AssertNotNil(t, "ValidateAlertSettings", err)
}

func TestValidateAlertSettingsInvalidAlertSubTypeForType(t *testing.T) {
	orig := GetMonitorAlertTypesFunc
	defer func() { GetMonitorAlertTypesFunc = orig }()

	GetMonitorAlertTypesFunc = func(testType types.TestType, monitorType string) *MonitorAlertTypes {
		return fakeAlertMatrix.MonitorTypes[monitorType]
	}

	alertSettings := []map[string]any{
		{"alert_type": "fakeAlertType", "alert_sub_type": "anotherFakeAlertSubType"},
	}

	err := ValidateTestAlertTypeCombination(alertSettings, fakeTestType, "fakeMonitorType1")
	testutil.AssertNotNil(t, "ValidateAlertSettings", err)
}
