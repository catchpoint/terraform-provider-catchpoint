package alerttypes

import (
	"catchpoint-provider/internal/types"
	"fmt"
)

func GetMonitorAlertTypes(testType types.TestType, monitorType string) *MonitorAlertTypes {
	switch testType {
	case types.APIType:
		return apiAlertMatrix.MonitorTypes[monitorType]
	case types.BGPType:
		return bgpAlertMatrix.MonitorTypes[monitorType]
	case types.DNSType:
		return dnsAlertMatrix.MonitorTypes[monitorType]
	case types.SSLType:
		return sslAlertMatrix.MonitorTypes[monitorType]
	case types.PingType:
		return pingAlertMatrix.MonitorTypes[monitorType]
	case types.PlaywrightType:
		return playwrightAlertMatrix.MonitorTypes[monitorType]
	case types.PuppeteerType:
		return puppeteerAlertMatrix.MonitorTypes[monitorType]
	case types.TracerouteType:
		return tracerouteAlertMatrix.MonitorTypes[monitorType]
	case types.WebType:
		return webAlertMatrix.MonitorTypes[monitorType]
	case types.TransactionType:
		return transactionAlertMatrix.MonitorTypes[monitorType]
	default:
		// This will only happen if someone didn't create the type mapping for a given test.
		panic(fmt.Sprintf("Unsupported test type and monitor combination for alert types: %d, %s", testType, monitorType))
	}
}

func GetTestCompatibilityMatrix(testType types.TestType) *MonitorAlertTypeCompatibilityMatrix {
	switch testType {
	case types.APIType:
		return apiAlertMatrix
	case types.BGPType:
		return bgpAlertMatrix
	case types.DNSType:
		return dnsAlertMatrix
	case types.SSLType:
		return sslAlertMatrix
	case types.PingType:
		return pingAlertMatrix
	case types.PlaywrightType:
		return playwrightAlertMatrix
	case types.PuppeteerType:
		return puppeteerAlertMatrix
	case types.TracerouteType:
		return tracerouteAlertMatrix
	case types.TransactionType:
		return transactionAlertMatrix
	case types.WebType:
		return webAlertMatrix
	default:
		// This will only happen if someone didn't create the type mapping for a given test.
		panic(fmt.Sprintf("Unsupported test type for alert compatibility matrix: %d", testType))
	}
}
