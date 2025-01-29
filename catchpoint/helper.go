package catchpoint

import (
	"math/rand"
	"regexp"
	"time"
)

type Config struct {
	ApiToken    string
	LogJson     bool
	Environment string
}

func newConfig(apiToken string, logJson bool, cpEnvironment string) *Config {
	return &Config{
		ApiToken:    apiToken,
		LogJson:     logJson,
		Environment: cpEnvironment,
	}
}

func getTime() string {
	t := time.Now()
	timeCurrent := t.Format(time.RFC3339)
	return string(timeCurrent)
}

func randHexString() string {
	const hexLetters = "abcdef0123456789"
	const numOfLetters = 6
	b := make([]byte, numOfLetters)
	for i := range b {
		b[i] = hexLetters[rand.Intn(len(hexLetters))]
	}
	return "#" + string(b)
}

func getTestStatusTypeId(testStatus string) int {
	id, ok := getTestStatusTypeIds()[testStatus]
	if ok {
		return id
	}
	return 0
}

func getMonitorId(monitor string) int {
	id, ok := getMonitorTypeIds()[monitor]
	if ok {
		return id
	}
	return -1
}

func getMonitorName(monitor int) string {
	name, ok := getMonitorTypeNames()[monitor]
	if ok {
		return name
	}
	return ""
}

func getApiScriptTypeId(scriptType string) int {
	id, ok := getApiScriptTypeIds()[scriptType]
	if ok {
		return id
	}
	return -1
}

func getUserAgentTypeId(userAgentType string) int {
	id, ok := getUserAgentTypeIds()[userAgentType]
	if ok {
		return id
	}
	return 0
}

func getUserAgentTypeName(userAgentType int) string {
	name, ok := getUserAgentTypeNames()[userAgentType]
	if ok {
		return name
	}
	return ""
}

func getChromeVersionId(chromeVersion string) (int, string) {
	for id, chromeVer := range getChromeVersionIds() {
		for _, specifcVersion := range chromeVer {
			if specifcVersion == chromeVersion {
				return id, specifcVersion
			}
		}
	}
	return 0, ""
}

func getChromeApplicationVersionId(chromeApplicationVersion string) (int, string) {
	id, ok := getChromeApplicationVersionIds()[chromeApplicationVersion]
	if ok {
		return id, chromeApplicationVersion
	}
	return 0, ""
}

func getDnsQueryTypeId(queryType string) (int, string) {
	id, ok := getDnsQueryTypeIds()[queryType]
	if ok {
		return id, queryType
	}
	return 0, ""
}

func getDnsQueryName(queryType int) string {
	name, ok := getDnsQueryTypeNames()[queryType]
	if ok {
		return name
	}
	return ""
}

func getFrequencyId(frequency string) (int, string) {
	id, ok := getFrequencyIds()[frequency]
	if ok {
		return id, frequency
	}
	return -1, ""
}

func getFrequencyName(frequency int) string {
	name, ok := getFrequencyNames()[frequency]
	if ok {
		return name
	}
	return ""
}

func getNodeDistributionId(nodeDistribution string) (int, string) {
	id, ok := getNodeDistributionIds()[nodeDistribution]
	if ok {
		return id, nodeDistribution
	}
	return -1, ""
}

func getNodeDistributionName(nodeDistribution int) string {
	name, ok := getNodeDistributionNames()[nodeDistribution]
	if ok {
		return name
	}
	return ""
}

func getNodeThresholdTypeId(nodeThresholdType string) (int, string) {
	id, ok := getNodeThresholdTypeIds()[nodeThresholdType]
	if ok {
		return id, nodeThresholdType
	}
	return -1, ""
}

func getNodeThresholdTypeName(nodeThresholdType int) string {
	name, ok := getNodeThresholdTypeNames()[nodeThresholdType]
	if ok {
		return name
	}
	return ""
}

func getOperationTypeId(operationType string) (int, string) {
	id, ok := getOperationTypeIds()[operationType]
	if ok {
		return id, operationType
	}
	return -1, ""
}

func getOperationTypeName(operationType int) string {
	name, ok := getOperationTypeNames()[operationType]
	if ok {
		return name
	}
	return ""
}

func getTriggerTypeId(triggerType string) (int, string) {
	id, ok := getTriggerTypeIds()[triggerType]
	if ok {
		return id, triggerType
	}
	return 1, "specific value"
}

func getTriggerTypeName(triggerType int) string {
	name, ok := getTriggerTypeNames()[triggerType]
	if ok {
		return name
	}
	return "specific value"
}

func getReminderId(reminder string) (int, string) {
	id, ok := getReminderIds()[reminder]
	if ok {
		return id, reminder
	}
	return 0, "none"
}

func getReminderName(reminder int) string {
	name, ok := getReminderNames()[reminder]
	if ok {
		return name
	}
	return "none"
}

func getThresholdIntervalId(thresholdInterval string) (int, string) {
	id, ok := getThresholdIntervalIds()[thresholdInterval]
	if ok {
		return id, thresholdInterval
	}
	return 0, "default"
}

func getThresholdIntervalName(thresholdInterval int) string {
	name, ok := getThresholdIntervalNames()[thresholdInterval]
	if ok {
		return name
	}
	return "default"
}

func getHistoricalIntervalId(historicalInterval string) (int, string) {
	id, ok := getHistoricalIntervalIds()[historicalInterval]
	if ok {
		return id, historicalInterval
	}
	return 5, "5 minutes"
}

func getHistoricalIntervalName(historicalInterval int) string {
	name, ok := getHistoricalIntervalNames()[historicalInterval]
	if ok {
		return name
	}
	return "5 minutes"
}

func getNotificationTypeId(notificationType string) int {
	id, ok := getNotificationTypeIds()[notificationType]
	if ok {
		return id
	}
	return 0
}

func getAlertTypeId(alertType string) (int, string) {
	id, ok := getAlertTypeIds()[alertType]
	if ok {
		return id, alertType
	}
	return -1, ""
}

func getAlertTypeName(alertType int) string {
	name, ok := getAlertTypeNames()[alertType]
	if ok {
		return name
	}
	return ""
}

func getAlertSubTypeId(alertSubType string) (int, string) {
	id, ok := getAlertSubTypeIds()[alertSubType]
	if ok {
		return id, alertSubType
	}
	return -1, ""
}

func getAlertSubTypeName(alertSubType int) string {
	name, ok := getAlertSubTypeNames()[alertSubType]
	if ok {
		return name
	}
	return ""
}

func getStatisticalTypeId(statisticalType string) (int, string) {
	id, ok := getStatisticalTypeIds()[statisticalType]
	if ok {
		return id, statisticalType
	}
	return 1, "average"
}

func getTestFlagId(testFlag string) int {
	id, ok := getTestFlagIds()[testFlag]
	if ok {
		return id
	}
	return 0

}

func getTestFlagName(testFlag int) string {
	name, ok := getTestFlagNames()[testFlag]
	if ok {
		return name
	}
	return ""
}

func getAdditionalMonitorTypeId(additionalMonitorType string) (int, string) {
	id, ok := getAdditionalMonitorTypeIds()[additionalMonitorType]
	if ok {
		return id, additionalMonitorType
	}
	return 0, ""
}

func getAdditionalMonitorTypeName(additionalMonitorType int) string {
	name, ok := getAdditionalMonitorTypeNames()[additionalMonitorType]
	if ok {
		return name
	}
	return ""
}

func getBandwidthThrottlingTypeId(bandwidthThrottlingType string) (int, string) {
	id, ok := getBandwidthThrottlingTypeIds()[bandwidthThrottlingType]
	if ok {
		return id, bandwidthThrottlingType
	}
	return 0, ""
}

func getBandwidthThrottlingTypeName(bandwidthThrottlingType int) string {
	name, ok := getBandwidthThrottlingTypeNames()[bandwidthThrottlingType]
	if ok {
		return name
	}
	return ""
}

func getReqHeaderTypeId(requestHeader string) (int, string) {
	id, ok := getReqHeaderTypeIds()[requestHeader]
	if ok {
		return id, requestHeader
	}
	return 0, ""
}

func getReqHeaderTypeName(requestHeader int) string {
	name, ok := getReqHeaderTypeNames()[requestHeader]
	if ok {
		return name
	}
	return ""
}

func getAuthenticationTypeId(authenticationType string) (int, string) {
	id, ok := getAuthenticationTypeIds()[authenticationType]
	if ok {
		return id, authenticationType
	}
	return -1, ""
}

func isValidEmail(email string) bool {
	// Regular expression pattern for validating email addresses
	pattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`

	// Compile the pattern
	regex := regexp.MustCompile(pattern)

	// Use the MatchString method to check if the email matches the pattern
	return regex.MatchString(email)
}

func getDnsTypeId(dnsRecordType string) (int, string) {
	id, ok := getDnsTypeIds()[dnsRecordType]
	if ok {
		return id, dnsRecordType
	}
	return 0, ""
}

func getFilterTypeid(filterType string) (int, string) {
	id, ok := getFilterTypeIds()[filterType]
	if ok {
		return id, filterType
	}
	return 0, ""
}

func getAlertSettingTypeId(alertSettingType string) (int, string) {
	id, ok := getAlertSettingTypeIds()[alertSettingType]
	if ok {
		return id, alertSettingType
	}
	return 1, "override"
}

func getAlertSettingTypeName(alertSettingType int) string {
	name, ok := getAlertSettingTypeNames()[alertSettingType]
	if ok {
		return name
	}
	return ""
}
