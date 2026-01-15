package service

import (
	"encoding/json"

	"catchpoint-provider/internal/helpers"
	"catchpoint-provider/internal/labels"
	"catchpoint-provider/internal/models"
	"catchpoint-provider/internal/types"
)

func CreateJSONTestPatchDocument(config *models.TestConfigUpdate, path string, isTestMetaData bool) string {
	return createJSONPatchDocument(config, path, isTestMetaData)
}

func CreateTestJSON(config models.TestConfig) string {
	test := newTestJSON(config)

	setJSONBlockFields(&config, &test)
	setJSONOptionalFields(&config, &test)

	testJSON, _ := json.Marshal(test)
	return string(testJSON)
}

func RequestSettingFromTestConfig(config *models.TestConfig) models.RequestSettingsJSON {
	return buildRequestSettings(&config.CommonConfig.RequestSettingsConfig)
}

func AdvancedSettingFromTestConfig(config *models.TestConfig) models.AdvancedSettingsJSON {
	return buildAdvancedSettings(&config.CommonConfig.AdvancedSettingsConfig)
}

func AlertSettingsFromTestConfig(config *models.TestConfig) models.AlertGroupJSON {
	return buildAlertSettings(&config.CommonConfig.AlertSettingsConfig)
}

func InsightDataFromTestConfig(config *models.TestConfig) models.InsightDataJSON {
	return buildInsightSettings(&config.CommonConfig.InsightSettingsConfig)
}

func ScheduleSettingsFromTestConfig(config *models.TestConfig) models.ScheduleSettingsJSON {
	return buildScheduleSettings(&config.CommonConfig.ScheduleSettingsConfig)
}

func TestRequestDataFromTestConfig(config *models.TestConfig) models.TestRequestDataJSON {
	monitor := models.GenericIDNameOmitEmptyJSON{ID: &config.Script.Monitor.ID, Name: &config.Script.Monitor.Name}
	testType := models.GenericIDNameOmitEmptyJSON{ID: &config.Script.TestType.ID, Name: &config.Script.TestType.Name}
	transactionScriptType := models.GenericIDNameOmitEmptyJSON{ID: &config.Script.TransactionScriptType.ID, Name: &config.Script.TransactionScriptType.Name}
	requestData := models.TestRequestDataJSON{
		TestID:                &config.Script.TestID,
		RequestData:           &config.Script.RequestData,
		TransactionScriptType: &transactionScriptType,
		Monitor:               &monitor,
		TestType:              &testType,
	}

	return requestData
}

func ThresholdsFromTestConfig(config *models.TestConfig) (thresholds models.TestThresholdsJSON) {
	thresholds = models.TestThresholdsJSON{}

	// Only set thresholds that have values.
	if config.TestThresholds.TestTimeThresholdWarning > 0 {
		thresholds.TestTimeApdexThresholdWarning = &config.TestThresholds.TestTimeThresholdWarning
	}
	if config.TestThresholds.TestTimeThresholdCritical > 0 {
		thresholds.TestTimeApdexThresholdCritical = &config.TestThresholds.TestTimeThresholdCritical
	}
	if config.TestThresholds.AvailabilityThresholdWarning > 0 {
		thresholds.AvailabilityApdexThresholdWarning = &config.TestThresholds.AvailabilityThresholdWarning
	}
	if config.TestThresholds.AvailabilityThresholdCritical > 0 {
		thresholds.AvailabilityApdexThresholdCritical = &config.TestThresholds.AvailabilityThresholdCritical
	}

	return
}

func LabelsFromTestConfig(config *models.TestConfig) []models.LabelsJSON {
	labels := []models.LabelsJSON{}

	if len(config.Labels) > 0 {
		for i := range config.Labels {
			labels = append(labels, models.LabelsJSON{Color: helpers.RandomHexString(), Name: config.Labels[i].Name, Values: config.Labels[i].Values})
		}
	}

	return labels
}

func newTestJSON(config models.TestConfig) models.TestJSON {
	//Set properties
	status := models.GenericIDNameJSON{ID: config.Status.ID, Name: config.Status.Name}
	monitor := models.GenericIDNameJSON{ID: config.Monitor.ID, Name: config.Monitor.Name}
	testType := models.GenericIDNameJSON{ID: config.TestType.ID, Name: config.TestType.Name}

	testID := 0
	changeDate := helpers.GetTime()

	var t = models.TestJSON{
		ID:                    testID,
		DivisionID:            config.DivisionID,
		ProductID:             config.ProductID,
		Name:                  config.TestName,
		EnableTestDataWebhook: config.EnableTestDataWebhook,
		AlertsPaused:          config.AlertsPaused,
		ChangeDate:            changeDate,
		StartTime:             config.StartTime,
		Status:                status,
		Monitor:               monitor,
		TestType:              testType,
	}

	return t
}

func setJSONBlockFields(config *models.TestConfig, testJSON *models.TestJSON) {
	requestHTTPMethod := models.GenericIDNameJSON{ID: 0, Name: labels.Get}

	alertGroup := buildAlertSettings(&config.CommonConfig.AlertSettingsConfig)
	requestSettings := buildRequestSettings(&config.CommonConfig.RequestSettingsConfig)
	scheduleSettings := buildScheduleSettings(&config.CommonConfig.ScheduleSettingsConfig)
	advancedSettings := buildAdvancedSettings(&config.CommonConfig.AdvancedSettingsConfig)

	insightData := InsightDataFromTestConfig(config)
	labels := LabelsFromTestConfig(config)

	if requestHTTPMethod != (models.GenericIDNameJSON{}) {
		testJSON.RequestHTTPMethod = &requestHTTPMethod
	}

	if insightData.Indicators != nil || insightData.Tracepoints != nil {
		testJSON.InsightData = &insightData
	}

	if len(labels) > 0 {
		testJSON.Labels = &labels
	}

	if requestSettings != (models.RequestSettingsJSON{}) {
		testJSON.RequestSettings = &requestSettings
	}

	if advancedSettings.AdvancedSettingType != (models.GenericIDNameJSON{}) {
		testJSON.AdvancedSettings = &advancedSettings
	}

	if scheduleSettings.ScheduleSettingType != (models.GenericIDNameJSON{}) {
		testJSON.ScheduleSettings = &scheduleSettings
	}

	if alertGroup.AlertSettingType != (models.GenericIDNameJSON{}) {
		testJSON.AlertGroup = &alertGroup
	}
}

func setJSONOptionalFields(config *models.TestConfig, test *models.TestJSON) {
	if config.FolderID != 0 {
		test.FolderID = &config.FolderID
	}

	assignStringJSON(&test.URL, config.TestURL)
	assignStringJSON(&test.Description, config.TestDescription)
	assignStringJSON(&test.GatewayAddressOrHost, config.GatewayAddressOrHost)
	assignStringJSON(&test.CertificateName, config.CertificateName)

	test.EnforceCertificatePinning = &config.EnforceCertificatePinning
	test.EnforceCertificateKeyPinning = &config.EnforceCertificateKeyPinning

	assignStringJSON(&test.FileData, config.FileData)
	assignStringJSON(&test.PassPhrase, config.Passphrase)
	assignStringJSON(&test.EndTime, config.EndTime)

	if config.SimulateDevice != (models.IDNameConfig{}) {
		userAgentType := models.GenericIDNameOmitEmptyJSON{ID: &config.SimulateDevice.ID, Name: &config.SimulateDevice.Name}
		if userAgentType != (models.GenericIDNameOmitEmptyJSON{}) {
			test.UserAgentType = &userAgentType
		}
	}

	if config.ChromeApplicationVersion != (models.ChromeMonitorVersion{}) {
		applicationVersionType := models.GenericIDNameOmitEmptyJSON{ID: &config.ChromeApplicationVersion.ApplicationVersionType.ID, Name: &config.ChromeApplicationVersion.ApplicationVersionType.Name}
		chromeMonitor := models.ChromeMonitorVersionStructJSON{ApplicationVersionType: &applicationVersionType, ApplicationVersionID: &config.ChromeApplicationVersion.ApplicationVersionID}
		if chromeMonitor != (models.ChromeMonitorVersionStructJSON{}) {
			test.ChromeMonitorVersion = &chromeMonitor
		}
	}

	if config.DNSQueryType != (models.IDNameConfig{}) {
		dnsQueryType := models.GenericIDNameOmitEmptyJSON{ID: &config.DNSQueryType.ID, Name: &config.DNSQueryType.Name}
		test.DNSQueryType = &dnsQueryType
	}

	if config.DNSServer != types.EmptyString {
		test.DNSServer = &config.DNSServer
	}

	if config.Script != (models.TestRequestData{}) {
		requestData := TestRequestDataFromTestConfig(config)
		test.TestRequestData = &requestData
	}

	thresholds := ThresholdsFromTestConfig(config)
	if thresholds != (models.TestThresholdsJSON{}) {
		test.TestThresholds = &thresholds
	}
}
