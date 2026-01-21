package service

import (
	"encoding/json"
	"testing"

	"catchpoint-provider/internal/fields"
	"catchpoint-provider/internal/models"
	"catchpoint-provider/internal/testutil"
	"catchpoint-provider/internal/types"
)

const (
	labelsValue             = `"labels":[{"id":0,"name":"string","labelValues":[{"id":0,"name":"string","labelKeyId":0,"associatedLabelValues":[{"id":0,"objectType":{"id":0,"name":"None"},"objectId":0,"labelValueId":0,"labelKeyId":0,"changeDate":"2025-07-03T20:42:46.920Z","changeContactId":0,"createDate":"2025-07-03T20:42:46.920Z","createContactId":0,"name":"string","changeLogObjectType":{"id":0,"name":"None"}}],"labelName":"string","changeLogObjectType":{"id":0,"name":"None"}}],"color":"string","changeLogObjectType":{"id":0,"name":"None"},"objectId":0}],"labelsForChangeLog":["string"]}`
	thresholdRestModelValue = `"thresholdRestModel":{"testTimeMonitorTypeId":0,"testTimeApdexThresholdWarning":0,"testTimeApdexThresholdCritical":0,"testTimeMetricType":{"id":0,"name":"None"},"availabilityMonitorTypeId":0,"availabilityApdexThresholdWarning":0,"availabilityApdexThresholdCritical":0,"availabilityMetricType":{"id":0,"name":"None"}`
	testRequestDataValue    = `"testRequestData":{"id":0,"testID":0,"requestData":"string","emailMessageFields":{"emailFrom":"string","emailTo":["string"],"emailBcc":["string"],"emailCc":["string"],"emailSubject":"string","emailSearch":"string","emailBody":"string"},"dnsServer":"string","subscribeTopic":"string","publishTopic":"string","publishMessage":"string","transactionScriptType":{"id":1,"name":"Selenium"},"scriptName":"string","customFields":[{"parameterId":0,"name":"string","value":"string"}],"testType":{"id":0,"name":"Web"},"monitor":{"id":0,"name":"IE"}}`
	insightDataValue        = `"insightData":{"id":0,"testID":0,"insightSettingType":{"id":1,"name":"Override"},"tracepointIDs":[1,2,3],"indicatorIDs":[4,5,6],"changeLogObjectType":{"id":0,"name":"None"}}`
	ScheduleSettingSection  = `"scheduleSettings":{"scheduleSettingType":{"id":1,"name":"Override"},"frequency":{"id":0,"name":"None"},"startTime":"2024-01-01T01:00:00Z","endTime":"2024-01-01T02:00:00Z","nodeIDs":[1,2,3],"nodeGroupIDs":[{"id":4},{"id":5},{"id":6}],"noOfSubsetNodes":2,"changeLogObjectType":{"id":0,"name":"None"}}`
)

// #region CreateTests

func TestCreateJSONBasicFields(t *testing.T) {
	config := newTestConfig()

	testObj := unmarshalTestConfigToTest(t, &config)

	compareAllTestFields(t, &testObj, &config)
}

func TestCreateJSONUserAgentType(t *testing.T) {
	config := newTestConfig()

	config.SimulateDevice = models.IDNameConfig{
		ID:   3,
		Name: "Android",
	}

	testObj := unmarshalTestConfigToTest(t, &config)

	compareAllTestFields(t, &testObj, &config)
	testutil.AssertEqual(t, "UserAgentType.ID", *testObj.UserAgentType.ID, config.SimulateDevice.ID)
	testutil.AssertEqual(t, "UserAgentType.Name", *testObj.UserAgentType.Name, config.SimulateDevice.Name)
}

func TestCreateJSONChromeMonitorVersion(t *testing.T) {
	config := newTestConfig()
	config.ChromeApplicationVersion.ApplicationVersionID = 1
	config.ChromeApplicationVersion.ApplicationVersionType = models.IDNameConfig{
		ID:   2,
		Name: "preview",
	}

	testObj := unmarshalTestConfigToTest(t, &config)

	compareAllTestFields(t, &testObj, &config)
	testutil.AssertEqual(t, "ChromeMonitorVersion.ApplicationVersionType.ID", *testObj.ChromeMonitorVersion.ApplicationVersionType.ID, config.ChromeApplicationVersion.ApplicationVersionType.ID)
	testutil.AssertEqual(t, "ChromeMonitorVersion.ApplicationVersionType.Name", *testObj.ChromeMonitorVersion.ApplicationVersionType.Name, config.ChromeApplicationVersion.ApplicationVersionType.Name)
}

func TestCreateJSONWithLabelsAndThresholds(t *testing.T) {
	config := newTestConfig()
	config.Labels = []models.TestLabel{
		{Name: "env", Values: []string{"prod", "staging"}},
		{Name: "team", Values: []string{"devops"}},
	}
	config.TestThresholds.TestTimeThresholdWarning = 1.1
	config.TestThresholds.TestTimeThresholdCritical = 2.2
	config.TestThresholds.AvailabilityThresholdWarning = 3.3
	config.TestThresholds.AvailabilityThresholdCritical = 4.4

	testObj := unmarshalTestConfigToTest(t, &config)

	compareAllTestFields(t, &testObj, &config)
	compareLabels(t, &testObj, &config)
	compareThresholds(t, &testObj, &config)
}

func TestCreateJSONWithAlertSettings(t *testing.T) {
	config := newTestConfig()
	config.CommonConfig = newCommonConfig()
	testObj := unmarshalTestConfigToTest(t, &config)

	compareAllTestFields(t, &testObj, &config)
	checkAlertGroupFields(t, testObj.AlertGroup, &config.CommonConfig)
	checkAlertGroupItemFields(t, &testObj.AlertGroup.AlertGroupItems[0], config.AlertSettingsConfig.AlertRules[0])
}

func TestCreateJSONWithAlertSettingsNoSubType(t *testing.T) {
	config := newTestConfig()
	config.CommonConfig = newCommonConfig()

	// There's a different code path if AlertSubType is nil or empty
	// so we set it to an empty struct to simulate that.
	config.CommonConfig.AlertSettingsConfig.AlertRules[0].AlertSubType = models.IDNameConfig{}
	testObj := unmarshalTestConfigToTest(t, &config)

	compareAllTestFields(t, &testObj, &config)
	checkAlertGroupFields(t, testObj.AlertGroup, &config.CommonConfig)
	checkAlertGroupItemFields(t, &testObj.AlertGroup.AlertGroupItems[0], config.AlertSettingsConfig.AlertRules[0])
}

func TestCreateJSONWithInsightSettings(t *testing.T) {
	config := newTestConfig()

	config.CommonConfig = newCommonConfig()
	config.CommonConfig.InsightSettingsConfig.InsightSettingType = models.IDNameConfig{ID: 1, Name: "Override"}
	config.CommonConfig.InsightSettingsConfig.TracepointIDs = []int{1, 2, 3}
	config.CommonConfig.InsightSettingsConfig.IndicatorIDs = []int{4, 5, 6}

	testObj := unmarshalTestConfigToTest(t, &config)

	compareAllTestFields(t, &testObj, &config)
	compareInsightDataStructFields(t, testObj.InsightData, &config.CommonConfig)
}

func TestCreateJSONDNSFields(t *testing.T) {
	config := newTestConfig()
	config.TestType = models.IDNameConfig{ID: int(types.DNSType), Name: "dns"}
	config.DNSQueryType = models.IDNameConfig{ID: 10, Name: "A"}
	config.DNSServer = "8.8.8.8"

	testObj := unmarshalTestConfigToTest(t, &config)

	compareAllTestFields(t, &testObj, &config)

	testutil.AssertEqual(t, "DNSQueryType.ID", *testObj.DNSQueryType.ID, config.DNSQueryType.ID)
	testutil.AssertEqual(t, "DNSQueryType.Name", *testObj.DNSQueryType.Name, config.DNSQueryType.Name)
	testutil.AssertEqual(t, "DNSServer", *testObj.DNSServer, config.DNSServer)
}

func TestCreateJSONWithScheduleSettings(t *testing.T) {
	config := newTestConfig()
	config.CommonConfig = newCommonConfig()
	config.ScheduleSettingsConfig.NodeIDs = []int{1, 2, 3}
	config.ScheduleSettingsConfig.NodeGroupIDs = []models.IDNameConfig{
		{ID: 4},
		{ID: 5},
		{ID: 6},
	}
	config.ScheduleSettingsConfig.NoOfSubsetNodes = 2

	testObj := unmarshalTestConfigToTest(t, &config)

	compareAllTestFields(t, &testObj, &config)
	compareScheduleSettingSectionFields(t, testObj.ScheduleSettings, &config.CommonConfig)
}

func TestCreateJSONTestRequestData(t *testing.T) {
	config := newTestConfig()
	config.TestType = models.IDNameConfig{ID: int(types.APIType), Name: "api"}
	config.Script = models.TestRequestData{
		TestID:                123,
		RequestData:           "request data",
		TransactionScriptType: models.IDNameConfig{ID: 5, Name: "Transaction"},
		TestType:              models.IDNameConfig{ID: 6, Name: "Test"},
		Monitor:               models.IDNameConfig{ID: 7, Name: "Monitor"},
	}

	testObj := unmarshalTestConfigToTest(t, &config)

	compareAllTestFields(t, &testObj, &config)

	testutil.AssertEqual(t, "Script.TestID", *testObj.TestRequestData.TestID, config.Script.TestID)
	testutil.AssertEqual(t, "Script.RequestData", *testObj.TestRequestData.RequestData, config.Script.RequestData)
}

func TestCreateJSONWithRequestSettings(t *testing.T) {
	config := newTestConfig()
	config.CommonConfig.RequestSettingsConfig.RequestSettingType = models.IDNameConfig{ID: 1, Name: "Override"}
	config.CommonConfig.RequestSettingsConfig.TestHTTPHeaderRequests = newTestHTTPHeaders()
	config.CommonConfig.RequestSettingsConfig.AuthenticationType = models.IDNameConfig{
		ID:   3,
		Name: "Basic",
	}

	testObj := unmarshalTestConfigToTest(t, &config)

	compareAllTestFields(t, &testObj, &config)
	compareRequestSettingFields(t, testObj.RequestSettings, &config.CommonConfig)
}

func TestCreateJSONWithAdvancedSettings(t *testing.T) {
	config := newTestConfig()
	config.CommonConfig = newCommonConfigWithAdvancedSettings()

	testObj := unmarshalTestConfigToTest(t, &config)
	compareAllTestFields(t, &testObj, &config)
	compareAdvancedSettingsFields(t, testObj.AdvancedSettings, &config.CommonConfig)
}

func TestCreateJSONWithMinimalAdvancedSettings(t *testing.T) {
	config := newTestConfig()
	config.CommonConfig.AdvancedSettingsConfig = models.AdvancedSettingsConfig{
		AdvancedSettingType: models.IDNameConfig{
			ID:   0,
			Name: "Inherit",
		},
	}

	testObj := unmarshalTestConfigToTest(t, &config)
	compareAllTestFields(t, &testObj, &config)
	testutil.AssertEqual(t, "AdvancedSettings.AdvancedSettingType.Name", testObj.AdvancedSettings.AdvancedSettingType.Name, "Inherit")
}

// #endregion

// #region PatchTests

func TestCreateJSONTestPatchDocument(t *testing.T) {
	tests := []struct {
		section string
		value   string
	}{
		{fields.LabelsSection, labelsValue},
		{fields.ThresholdRestModelSection, thresholdRestModelValue},
		{fields.TestRequestDataSection, testRequestDataValue},
		{fields.InsightDataSection, insightDataValue},
		{fields.ScheduleSettingSection, ScheduleSettingSection},
	}
	for _, tt := range tests {
		update := &models.TestConfigUpdate{
			UpdatedFieldValue: tt.value,
			SectionToUpdate:   tt.section,
		}
		jsonStr := CreateJSONTestPatchDocument(update, tt.section, true)

		if jsonStr == types.EmptyString {
			t.Error(errorEmptyJSONString)
		}
		var patch map[string]any
		if err := json.Unmarshal([]byte(jsonStr), &patch); err != nil {
			t.Errorf("Expected valid JSON, got error: %v", err)
		}
		if patch["path"] != tt.section {
			t.Errorf("Expected path '%s', got %v", tt.section, patch["path"])
		}
	}
}

// #endregion

// #region Helpers

func compareAllTestFields(t *testing.T, testObj *models.TestJSON, config *models.TestConfig) {
	compareTestAttributes(t, testObj, config)
	compareCertificateFields(t, testObj, config)
	compareStatusMonitorType(t, testObj, config)
	compareThresholds(t, testObj, config)
	compareLabels(t, testObj, config)
}

func compareTestAttributes(t *testing.T, testObj *models.TestJSON, config *models.TestConfig) {
	testutil.AssertEqual(t, "DivisionID", testObj.DivisionID, config.DivisionID)
	testutil.AssertEqual(t, "ProductID", testObj.ProductID, config.ProductID)
	testutil.AssertEqual(t, "FolderID", *testObj.FolderID, config.FolderID)
	testutil.AssertEqual(t, "Name", testObj.Name, config.TestName)
	testutil.AssertEqual(t, "Description", *testObj.Description, config.TestDescription)
	testutil.AssertEqual(t, "URL", *testObj.URL, config.TestURL)
	testutil.AssertEqual(t, "EnableTestDataWebhook", testObj.EnableTestDataWebhook, config.EnableTestDataWebhook)
	testutil.AssertEqual(t, "AlertsPaused", testObj.AlertsPaused, config.AlertsPaused)
	testutil.AssertEqual(t, "GatewayAddressOrHost", *testObj.GatewayAddressOrHost, config.GatewayAddressOrHost)
}

func compareCertificateFields(t *testing.T, testObj *models.TestJSON, config *models.TestConfig) {
	testutil.AssertEqual(t, "EnforceCertificatePinning", *testObj.EnforceCertificatePinning, config.EnforceCertificatePinning)
	testutil.AssertEqual(t, "EnforceCertificateKeyPinning", *testObj.EnforceCertificateKeyPinning, config.EnforceCertificateKeyPinning)
	testutil.AssertEqual(t, "FileData", *testObj.FileData, config.FileData)
	testutil.AssertEqual(t, "PassPhrase", *testObj.PassPhrase, config.Passphrase)
	testutil.AssertEqual(t, "CertificateName", *testObj.CertificateName, config.CertificateName)
}

func compareStatusMonitorType(t *testing.T, testObj *models.TestJSON, config *models.TestConfig) {
	testutil.AssertEqual(t, "Status.ID", testObj.Status.ID, config.Status.ID)
	testutil.AssertEqual(t, "Status.Name", testObj.Status.Name, config.Status.Name)
	testutil.AssertEqual(t, "Monitor.ID", testObj.Monitor.ID, config.Monitor.ID)
	testutil.AssertEqual(t, "Monitor.Name", testObj.Monitor.Name, config.Monitor.Name)
	testutil.AssertEqual(t, "TestType.ID", testObj.TestType.ID, config.TestType.ID)
	testutil.AssertEqual(t, "TestType.Name", testObj.TestType.Name, config.TestType.Name)
}

func compareThresholds(t *testing.T, testObj *models.TestJSON, config *models.TestConfig) {
	if config.TestThresholds != (models.TestThresholds{}) {
		testutil.AssertEqual(t, "TestTimeApdexThresholdWarning", *testObj.TestThresholds.TestTimeApdexThresholdWarning, config.TestThresholds.TestTimeThresholdWarning)
		testutil.AssertEqual(t, "TestTimeApdexThresholdCritical", *testObj.TestThresholds.TestTimeApdexThresholdCritical, config.TestThresholds.TestTimeThresholdCritical)
		testutil.AssertEqual(t, "AvailabilityApdexThresholdWarning", *testObj.TestThresholds.AvailabilityApdexThresholdWarning, config.TestThresholds.AvailabilityThresholdWarning)
		testutil.AssertEqual(t, "AvailabilityApdexThresholdCritical", *testObj.TestThresholds.AvailabilityApdexThresholdCritical, config.TestThresholds.AvailabilityThresholdCritical)
	}
}

func compareLabels(t *testing.T, testObj *models.TestJSON, config *models.TestConfig) {
	if testObj.Labels == nil {
		return
	}

	testutil.AssertEqual(t, "Labels count", len(*testObj.Labels), len(config.Labels))

	labelSet := make(map[string]bool)
	for _, e := range config.Labels {
		labelSet[e.Name] = false
	}

	for _, label := range *testObj.Labels {
		if _, ok := labelSet[label.Name]; ok {
			labelSet[label.Name] = true
		}
	}

	for k, v := range labelSet {
		if !v {
			t.Errorf("Label '%s' not found in test object", k)
		}
	}
}

func newTestConfig() models.TestConfig {
	return models.TestConfig{
		ProductID:                    2,
		FolderID:                     3,
		TestName:                     "Test Name",
		TestDescription:              "Test Description",
		TestURL:                      "https://example.com",
		GatewayAddressOrHost:         "gateway.example.com",
		EnforceCertificatePinning:    true,
		EnforceCertificateKeyPinning: false,
		FileData:                     "filedata",
		Passphrase:                   "passphrase",
		CertificateName:              "certname",
		EnableTestDataWebhook:        true,
		AlertsPaused:                 false,
		StartTime:                    "2024-01-01T01:00:00Z",
		EndTime:                      "2024-01-01T02:00:00Z",
		Status:                       models.IDNameConfig{ID: 0, Name: "Active"},
		Monitor:                      models.IDNameConfig{ID: 1, Name: "Monitor Name"},
		TestType:                     models.IDNameConfig{ID: int(types.WebType), Name: "Web"},
	}
}

func unmarshalTestConfigToTest(t *testing.T, config *models.TestConfig) models.TestJSON {
	jsonStr := CreateTestJSON(*config)
	if jsonStr == types.EmptyString {
		t.Error(errorEmptyJSONString)
	}
	var testObj models.TestJSON
	err := json.Unmarshal([]byte(jsonStr), &testObj)
	if err != nil {
		t.Fatalf(errorUnmarshal, err)
	}
	return testObj
}

//#endregion
