// This package provides validation functions for various types used in the Catchpoint provider.
// It includes the functions necessary to retrieve IDs and names for different types with sensible defaults.
package validation

import (
	"catchpoint-provider/internal/models"
	"catchpoint-provider/internal/types"
)

// GetFrequencyOrDefault returns a default frequency configuration based on the provided frequency string.
// If the frequency is not recognized, it returns a default configuration with ID -1 and an empty name.
func GetFrequencyOrDefault(frequency string) models.IDNameConfig {
	if id, ok := types.GetFrequencyID(frequency); ok {
		return models.IDNameConfig{ID: id, Name: frequency}
	}
	return models.IDNameConfig{ID: -1, Name: types.EmptyString}
}

// GetNodeDistributionOrDefault returns a default node distribution configuration based on the provided node distribution string.
// If the node distribution is not recognized, it returns a default configuration with ID -1 and an empty name.
func GetNodeDistributionOrDefault(nodeDistribution string) models.IDNameConfig {
	if id, ok := types.GetNodeDistributionID(nodeDistribution); ok {
		return models.IDNameConfig{ID: id, Name: nodeDistribution}
	}
	return models.IDNameConfig{ID: -1, Name: types.EmptyString}
}

// GetStatusTypeOrDefault returns a default status configuration based on the provided status string.
// If the status is not recognized, it returns a default configuration with ID 0 and the name "Active".
func GetStatusTypeOrDefault(testStatus string) models.IDNameConfig {
	if id, ok := types.GetStatusTypeID(testStatus); ok {
		return models.IDNameConfig{ID: id, Name: testStatus}
	}
	return models.IDNameConfig{ID: 0, Name: types.Active}
}

// GetMonitorTypeOrDefault returns a default monitor configuration based on the provided monitor string.
// If the monitor is not recognized, it returns a default configuration with ID -1 and an empty name.
func GetMonitorTypeOrDefault(monitor string) models.IDNameConfig {
	if id, ok := types.GetMonitorTypeID(monitor); ok {
		return models.IDNameConfig{ID: id, Name: monitor}
	}
	return models.IDNameConfig{ID: -1, Name: types.EmptyString}
}

// GetAPIScriptTypeOrDefault returns a default API script configuration based on the provided script type string.
// If the script type is not recognized, it returns a default configuration with ID -1 and an empty name.
func GetAPIScriptTypeOrDefault(scriptType string) models.IDNameConfig {
	if id, ok := types.GetAPIScriptTypeID(scriptType); ok {
		return models.IDNameConfig{ID: id, Name: scriptType}
	}
	return models.IDNameConfig{ID: -1, Name: types.EmptyString}
}

// GetUserAgentTypeOrDefault returns a default user agent configuration based on the provided user agent type string.
// If the user agent type is not recognized, it returns a default configuration with ID 0 and an empty name.
func GetUserAgentTypeOrDefault(userAgentType string) models.IDNameConfig {
	if id, ok := types.GetUserAgentTypeID(userAgentType); ok {
		return models.IDNameConfig{ID: id, Name: userAgentType}
	}
	return models.IDNameConfig{ID: 0, Name: types.EmptyString}
}

// GetChromeVersionOrDefault returns a default Chrome version configuration based on the provided version string.
// If the version is not recognized, it returns a default configuration with ID 0 and an empty name.
func GetChromeVersionOrDefault(chromeVersion string) models.IDNameConfig {
	for id, chromeVer := range types.GetChromeVersionIDs() {
		for _, specificVersion := range chromeVer {
			if specificVersion == chromeVersion {
				return models.IDNameConfig{ID: id, Name: specificVersion}
			}
		}
	}
	return models.IDNameConfig{ID: 0, Name: types.EmptyString}
}

// GetChromeApplicationVersionOrDefault returns a default Chrome application version configuration based on the provided version string.
// If the version is not recognized, it returns a default configuration with ID 0 and an empty name.
func GetChromeApplicationVersionOrDefault(chromeApplicationVersion string) models.IDNameConfig {
	if id, ok := types.GetChromeApplicationVersionIDs()[chromeApplicationVersion]; ok {
		return models.IDNameConfig{ID: id, Name: chromeApplicationVersion}
	}
	return models.IDNameConfig{ID: 0, Name: types.EmptyString}
}

// GetDNSQueryTypeOrDefault returns a default DNS query type configuration based on the provided query type string.
// If the query type is not recognized, it returns a default configuration with ID 0 and an empty name.
func GetDNSQueryTypeOrDefault(queryType string) models.IDNameConfig {
	if id, ok := types.GetDNSQueryTypeID(queryType); ok {
		return models.IDNameConfig{ID: id, Name: queryType}
	}
	return models.IDNameConfig{ID: 0, Name: types.EmptyString}
}

// GetDNSRecordTypeOrDefault returns a default DNS record type configuration based on the provided record type string.
// If the record type is not recognized, it returns a default configuration with ID 0 and an empty name.
func GetNodeThresholdTypeOrDefault(nodeThresholdType string) models.IDNameConfig {
	if id, ok := types.GetNodeThresholdTypeID(nodeThresholdType); ok {
		return models.IDNameConfig{ID: id, Name: nodeThresholdType}
	}
	return models.IDNameConfig{ID: -1, Name: types.EmptyString}
}

// GetOperationTypeOrDefault returns a default operation type configuration based on the provided operation type string.
// If the operation type is not recognized, it returns a default configuration with ID 0 "not equals".
func GetOperationTypeOrDefault(operationType string) models.IDNameConfig {
	if id, ok := types.GetOperationTypeID(operationType); ok {
		return models.IDNameConfig{ID: id, Name: operationType}
	}
	return models.IDNameConfig{ID: 0, Name: types.NotEquals}
}

// GetTriggerTypeOrDefault returns a default trigger type configuration based on the provided trigger type string.
// If the trigger type is not recognized, it returns a default configuration with ID 1 and the name "Specific".
func GetTriggerTypeOrDefault(triggerType string) models.IDNameConfig {
	if id, ok := types.GetTriggerTypeID(triggerType); ok {
		return models.IDNameConfig{ID: id, Name: triggerType}
	}
	return models.IDNameConfig{ID: 1, Name: types.SpecificValue}
}

// GetReminderTypeOrDefault returns a default reminder type configuration based on the provided reminder string.
// If the reminder is not recognized, it returns a default configuration with ID 0 and the name "None".
func GetReminderTypeOrDefault(reminder string) models.IDNameConfig {
	if id, ok := types.GetReminderID(reminder); ok {
		return models.IDNameConfig{ID: id, Name: reminder}
	}
	return models.IDNameConfig{ID: 0, Name: types.None}
}

// GetThresholdIntervalOrDefault returns a default threshold interval configuration based on the provided interval string.
// If the interval is not recognized, it returns a default configuration with ID 0 and the name "Default".
func GetThresholdIntervalOrDefault(thresholdInterval string) models.IDNameConfig {
	if id, ok := types.GetThresholdIntervalID(thresholdInterval); ok {
		return models.IDNameConfig{ID: id, Name: thresholdInterval}
	}
	return models.IDNameConfig{ID: 0, Name: types.Default}
}

// GetHistoricalIntervalOrDefault returns a default historical interval configuration based on the provided interval string.
// If the interval is not recognized, it returns a default configuration with ID 5 and the name "5 Minutes".
func GetHistoricalIntervalOrDefault(historicalInterval string) models.IDNameConfig {
	if id, ok := types.GetHistoricalIntervalID(historicalInterval); ok {
		return models.IDNameConfig{ID: id, Name: historicalInterval}
	}
	return models.IDNameConfig{ID: 5, Name: string(types.FiveMinutes)}
}

// GetNotificationTypeOrDefault returns a default notification type configuration based on the provided notification type string.
// If the notification type is not recognized, it returns a default configuration with ID 0 and the name "Contacts".
func GetNotificationTypeOrDefault(notificationType string) models.IDNameConfig {
	if id, ok := types.GetNotificationTypeID(notificationType); ok {
		return models.IDNameConfig{ID: id, Name: notificationType}
	}
	return models.IDNameConfig{ID: 0, Name: types.DefaultContacts}
}

// GetAlertTypeOrDefault returns a default alert type configuration based on the provided alert type string.
// If the alert type is not recognized, it returns a default configuration with ID -1 and an empty name.
func GetAlertTypeOrDefault(alertType string) models.IDNameConfig {
	if id, ok := types.GetAlertTypeID(alertType); ok {
		return models.IDNameConfig{ID: id, Name: alertType}
	}
	return models.IDNameConfig{ID: -1, Name: types.EmptyString}
}

// GetAlertSubTypeOrDefault returns a default alert sub-type configuration based on the provided alert sub-type string.
// If the alert sub-type is not recognized, it returns a default configuration with ID -1 and an empty name.
func GetAlertSubTypeOrDefault(alertSubType string) models.IDNameConfig {
	if id, ok := types.GetAlertSubTypeID(alertSubType); ok {
		return models.IDNameConfig{ID: id, Name: alertSubType}
	}
	return models.IDNameConfig{ID: -1, Name: types.EmptyString}
}

// GetStatisticalTypeOrDefault returns a default statistical type configuration based on the provided statistical type string.
// If the statistical type is not recognized, it returns a default configuration with ID 1 and the name "Average".
func GetStatisticalTypeOrDefault(statisticalType string) models.IDNameConfig {
	if id, ok := types.GetStatisticalTypeID(statisticalType); ok {
		return models.IDNameConfig{ID: id, Name: statisticalType}
	}
	return models.IDNameConfig{ID: 1, Name: types.Average}
}

// GetAdditionalMonitorTypeOrDefault returns a default additional monitor type configuration based on the provided additional monitor type string.
// If the additional monitor type is not recognized, it returns a default configuration with ID 0 and an empty name.
func GetAdditionalMonitorTypeOrDefault(additionalMonitorType string) models.IDNameConfig {
	if id, ok := types.GetAdditionalMonitorTypeID(additionalMonitorType); ok {
		return models.IDNameConfig{ID: id, Name: additionalMonitorType}
	}
	return models.IDNameConfig{ID: 0, Name: types.EmptyString}
}

// GetBandwidthThrottlingTypeOrDefault returns a default bandwidth throttling type configuration based on the provided bandwidth throttling type string.
// If the bandwidth throttling type is not recognized, it returns a default configuration with ID 0 and an empty name.
func GetBandwidthThrottlingTypeOrDefault(bandwidthThrottlingType string) models.IDNameConfig {
	if id, ok := types.GetBandwidthThrottlingTypeID(bandwidthThrottlingType); ok {
		return models.IDNameConfig{ID: id, Name: bandwidthThrottlingType}
	}
	return models.IDNameConfig{ID: 0, Name: types.EmptyString}
}

// GetReqHeaderTypeOrDefault returns a default request header type configuration based on the provided request header string.
// If the request header is not recognized, it returns a default configuration with ID 0 and an empty name.
func GetReqHeaderTypeOrDefault(requestHeader string) models.IDNameConfig {
	if id, ok := types.GetReqHeaderTypeID(requestHeader); ok {
		return models.IDNameConfig{ID: id, Name: requestHeader}
	}
	return models.IDNameConfig{ID: 0, Name: types.EmptyString}
}

// GetAuthenticationTypeOrDefault returns a default authentication type configuration based on the provided authentication type string.
// If the authentication type is not recognized, it returns a default configuration with ID 0 and the name "None".
func GetAuthenticationTypeOrDefault(authenticationType string) models.IDNameConfig {
	if id, ok := types.GetAuthenticationTypeID(authenticationType); ok {
		return models.IDNameConfig{ID: id, Name: authenticationType}
	}
	return models.IDNameConfig{ID: 0, Name: types.None}
}

// GetDNSRecordTypeOrDefault returns a default DNS record type configuration based on the provided DNS record type string.
// If the DNS record type is not recognized, it returns a default configuration with ID 0 and an empty name.
func GetDNSRecordTypeOrDefault(dnsRecordType string) models.IDNameConfig {
	if id, ok := types.GetDNSRecordTypeID(dnsRecordType); ok {
		return models.IDNameConfig{ID: id, Name: dnsRecordType}
	}
	return models.IDNameConfig{ID: 0, Name: types.EmptyString}
}

// GetFilterTypeOrDefault returns a default filter type configuration based on the provided filter type string.
// If the filter type is not recognized, it returns a default configuration with ID 0 and an empty name.
func GetFilterTypeOrDefault(filterType string) models.IDNameConfig {
	if id, ok := types.GetFilterTypeID(filterType); ok {
		return models.IDNameConfig{ID: id, Name: filterType}
	}
	return models.IDNameConfig{ID: 0, Name: types.EmptyString}
}

// GetAlertSettingTypeOrDefault returns a default alert setting type configuration based on the provided alert setting type string.
// If the alert setting type is not recognized, it returns a default configuration with ID 1 and the name "Override".
func GetAlertSettingTypeOrDefault(alertSettingType string) models.IDNameConfig {
	if id, ok := types.GetAlertSettingTypeID(alertSettingType); ok {
		return models.IDNameConfig{ID: id, Name: alertSettingType}
	}
	return models.IDNameConfig{ID: 1, Name: types.Override}
}

// GetGenericSettingTypeOrDefault returns a default generic setting type configuration based on the provided setting type string.
// If the setting type is not recognized, it returns a default configuration with ID 0 and the name "Inherit".
func GetGenericSettingTypeOrDefault(scheduleSettingType string) models.IDNameConfig {
	if id, ok := types.GetGenericSettingTypeID(scheduleSettingType); ok {
		return models.IDNameConfig{ID: id, Name: scheduleSettingType}
	}
	return models.IDNameConfig{ID: 0, Name: types.Inherit}
}

// GetInsightSettingTypeOrDefault returns a default insight setting type configuration based on the provided insight setting type string.
// If the insight setting type is not recognized, it returns a default configuration with ID 0 and the name "Inherit".
func GetInsightSettingTypeOrDefault(insightSettingType string) models.IDNameConfig {
	if id, ok := types.GetInsightSettingTypeID(insightSettingType); ok {
		return models.IDNameConfig{ID: id, Name: insightSettingType}
	}
	return models.IDNameConfig{ID: 0, Name: types.Inherit}
}
