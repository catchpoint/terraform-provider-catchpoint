package service

import (
	"catchpoint-provider/internal/fields"
	"catchpoint-provider/internal/models"
	"catchpoint-provider/internal/testutil"
	"catchpoint-provider/internal/types"
	"encoding/json"
	"testing"
)

const (
	advancedSettingsValue   = `"advancedSettingType":{"id":1,"name":"Override"},"appliedTestFlags":[{"id":1,"name":"Flag"},{"id":2,"name":"Flag"}],"maxStepRuntimeSecOverride":60,"waitForNoActivity":1,"viewportHeight":1080,"viewportWidth":1920,"failureHopCount":5,"pingCount":3,"ednsSubnet":"255.255.255.0","additionalMonitor":{"id":1,"name":"Monitor"},"testBandwidthThrottling":{"id":1,"name":"Throttle"},"verifyTestOnFailure":true,"id":0}`
	requestSettingsValue    = `"requestSettingType":{"id":1,"name":"Inherit"},"authentication":{"authenticationMethodType":{"id":1,"name":"Auth"},"passwordIds":[1,2]},"libraryCertificateIds":[5,6],"tokenIds":[3,4],"httpHeaderRequests":[{"requestValue":"val","requestHeaderType":{"id":1,"name":"Header"},"childHostPattern":"pattern","headerName":"X-Test"}`
	alertGroupSettingsValue = `"alertGroup":{"alertSettingType":{"id":0,"name":"Inherit"},"notificationGroup":{"subject":"TestSubject","notifyOnWarning":true,"notifyOnCritical":true,"notifyOnImproved":true,"alertWebhooks":[{"id":101}],"recipients":[{"email":"test@example.com","recipientType":{"id":2,"name":"Email"},"name":""},{"id":1,"email":"","recipientType":{"id":1,"name":"ContactGroup"},"name":"Group1"}]},"alertGroupItems":[{"nodeThreshold":{"id":0,"name":"","nodeThresholdType":{"id":1,"name":"Type1"},"numberOfUnits":5,"percentageOfUnits":50,"numberOfFailingUnits":2,"consecutiveRunsEnabled":true,"utilizePerNodeHistoricalAverage":false,"consecutiveRuns":3},"trigger":{"id":0,"warningReminderFrequency":{"id":1,"name":"Warn"},"criticalReminderFrequency":{"id":2,"name":"Crit"},"triggerType":{"id":1,"name":"Trigger"},"operationType":{"id":1,"name":"Op"},"statisticalType":{"id":1,"name":"Stat"},"historicalInterval":{"id":1,"name":"Hist"},"thresholdInterval":{"id":1,"name":"Interval"},"warningTrigger":1.1,"criticalTrigger":2.2,"useIntervalRollingWindow":true,"expression":"expr"},"notificationType":{"id":1,"name":"DefaultContacts"},"alertType":{"id":1,"name":"AlertType"},"alertSubType":{"id":1,"name":"SubType"},"enforceTestFailure":true,"omitScatterplot":false,"matchAllRecords":false,"notificationGroups":[]}]}`
)

// #region CreateTests

func TestCreateProductJSONMinimal(t *testing.T) {
	config := newProductConfig()
	testObj := unmarshalProductConfigToTest(t, &config)

	compareProductFields(t, &testObj, &config)
}

func TestCreateProductJSONMinimalHasNoAdvancedFlags(t *testing.T) {
	config := newProductConfig()
	testObj := unmarshalProductConfigToTest(t, &config)

	compareProductFields(t, &testObj, &config)

	testutil.AssertEqual(t, "AppliedTestFlags", len(testObj.AdvancedSettings.AppliedTestFlags), 0)
}

func TestCreateProductJSONMinimalHasAuthenticationObject(t *testing.T) {
	config := models.ProductConfig{
		ProductName:       "Test Product",
		AlertGroupID:      123,
		Status:            models.IDNameConfig{ID: 0, Name: "Active"},
		TestDataWebhookID: 456,
	}

	testObj := unmarshalProductConfigToTest(t, &config)

	compareProductFields(t, &testObj, &config)

	testutil.AssertEqual(t, "Authentication.ID", *testObj.RequestSettings.Authentication.ID, 0)
	testutil.AssertEqual(t, "Authentication.AuthenticationMethodType.ID", *testObj.RequestSettings.Authentication.AuthenticationMethodType.ID, 0)
	testutil.AssertEqual(t, "Authentication.AuthenticationMethodType.Name", *testObj.RequestSettings.Authentication.AuthenticationMethodType.Name, "none")
	testutil.AssertDeepEqual(t, "Authentication.PasswordIDs", *testObj.RequestSettings.Authentication.PasswordIDs, []int{})
}

func TestCreateProductJSONNoWebhooksNotNil(t *testing.T) {
	config := newProductConfig()
	config.CommonConfig.AlertSettingsConfig.NotificationGroup.WebhookIDs = []int{} // No webhooks
	testObj := unmarshalProductConfigToTest(t, &config)

	compareProductFields(t, &testObj, &config)
	// Assert that AlertWebhookIDs is an empty slice, not nil (nil breaks the API).
	testutil.AssertDeepEqual(t, "AlertWebhookIDs", testObj.AlertGroup.NotificationGroup.AlertWebhooks, []models.AlertWebhookJSON{})
}

func TestCreateProductJSONNoInsightsNotNil(t *testing.T) {
	config := newProductConfig()
	config.CommonConfig.InsightSettingsConfig.IndicatorIDs = []int{}  // No insights
	config.CommonConfig.InsightSettingsConfig.TracepointIDs = []int{} // No insights
	testObj := unmarshalProductConfigToTest(t, &config)

	compareProductFields(t, &testObj, &config)
	testutil.AssertDeepEqual(t, "InsightDataIndicators", testObj.InsightData.Indicators, []models.GenericIDNameJSON{})
	testutil.AssertNotNil(t, "InsightDataIndicators", testObj.InsightData.Indicators) // Indicators should be an empty slice, not nil
	testutil.AssertDeepEqual(t, "InsightDataTracepoints", testObj.InsightData.Tracepoints, []models.GenericIDNameJSON{})
	testutil.AssertNotNil(t, "InsightDataTracepoints", testObj.InsightData.Tracepoints) // Tracepoints should be an empty slice, not nil
}

func TestCreateProductJSONFull(t *testing.T) {
	config := newProductConfig()
	config.CommonConfig = newCommonConfig()
	testObj := unmarshalProductConfigToTest(t, &config)

	compareProductFields(t, &testObj, &config)
	checkAlertGroupFields(t, &testObj.AlertGroup, &config.CommonConfig)
}

func TestCreateProductJSONNoBandwidthThrottling(t *testing.T) {
	config := newProductConfig()
	config.CommonConfig = newCommonConfig()
	config.CommonConfig.AdvancedSettingsConfig.BandwidthThrottling = models.IDNameConfig{}

	testObj := unmarshalProductConfigToTest(t, &config)

	testutil.AssertNil(t, "TestBandwidthThrottling", testObj.AdvancedSettings.TestBandwidthThrottling)
}

func TestCreateProductJSONWithInsightSettings(t *testing.T) {
	config := newProductConfig()

	config.CommonConfig = newCommonConfig()
	config.CommonConfig.InsightSettingsConfig.InsightSettingType = models.IDNameConfig{ID: 1, Name: "Override"}
	config.CommonConfig.InsightSettingsConfig.TracepointIDs = []int{1, 2, 3}
	config.CommonConfig.InsightSettingsConfig.IndicatorIDs = []int{4, 5, 6}

	testObj := unmarshalProductConfigToTest(t, &config)

	compareProductFields(t, &testObj, &config)
	compareInsightDataStructFields(t, &testObj.InsightData, &config.CommonConfig)
}

func TestCreateProductJSONWithScheduleSettings(t *testing.T) {
	config := newProductConfig()
	config.CommonConfig = newCommonConfig()
	config.ScheduleSettingsConfig.NodeIDs = []int{1, 2, 3}
	config.ScheduleSettingsConfig.NodeGroupIDs = []models.IDNameConfig{
		{ID: 4},
		{ID: 5},
		{ID: 6},
	}
	config.ScheduleSettingsConfig.NoOfSubsetNodes = 2

	testObj := unmarshalProductConfigToTest(t, &config)

	compareProductFields(t, &testObj, &config)
	compareScheduleSettingSectionFields(t, &testObj.ScheduleSettings, &config.CommonConfig)
}

func TestCreateProductJSONWithRequestSettings(t *testing.T) {
	config := newProductConfig()
	config.CommonConfig.RequestSettingsConfig.RequestSettingType = models.IDNameConfig{ID: 1, Name: "Override"}
	config.CommonConfig.RequestSettingsConfig.TestHTTPHeaderRequests = newTestHTTPHeaders()
	config.CommonConfig.RequestSettingsConfig.AuthenticationType = models.IDNameConfig{
		ID:   3,
		Name: "Basic",
	}

	testObj := unmarshalProductConfigToTest(t, &config)

	compareProductFields(t, &testObj, &config)
	compareRequestSettingFields(t, &testObj.RequestSettings, &config.CommonConfig)
}

func TestCreateProductJSONWithAdvancedSettings(t *testing.T) {
	config := newProductConfig()
	config.CommonConfig = newCommonConfigWithAdvancedSettings()

	testObj := unmarshalProductConfigToTest(t, &config)
	compareProductFields(t, &testObj, &config)
	compareAdvancedSettingsFields(t, &testObj.AdvancedSettings, &config.CommonConfig)
}

// #endregion

// #region PatchTests

func TestCreateJSONProductPatchDocument(t *testing.T) {
	tests := []struct {
		section string
		value   string
	}{
		{fields.AdvancedSettingsModelSection, advancedSettingsValue},
		{fields.RequestSettingSection, requestSettingsValue},
		{fields.AlertGroupSection, alertGroupSettingsValue},
	}
	for _, tt := range tests {
		update := &models.ProductConfigUpdate{
			UpdatedFieldValue: tt.value,
			SectionToUpdate:   tt.section,
		}
		jsonStr := CreateJSONProductPatchDocument(update, tt.section, true)
		if jsonStr == types.EmptyString {
			t.Error(errorEmptyJSONString)
		}
		var patch map[string]interface{}
		if err := json.Unmarshal([]byte(jsonStr), &patch); err != nil {
			t.Errorf("Expected valid JSON, got error: %v", err)
		}
		if patch["path"] != tt.section {
			t.Errorf("Expected path '%s', got %v", tt.section, patch["path"])
		}
	}
}

//#endregion

// #region Helpers

func compareProductFields(t *testing.T, testObj *models.ProductJSON, config *models.ProductConfig) {
	testutil.AssertEqual(t, "ProductName", testObj.Name, config.ProductName)
	testutil.AssertEqual(t, "DivisionID", testObj.DivisionID, config.DivisionID)
	testutil.AssertEqual(t, "ProductStatusID", testObj.Status.ID, config.Status.ID)
	testutil.AssertEqual(t, "ProductStatusName", testObj.Status.Name, config.Status.Name)
}

func unmarshalProductConfigToTest(t *testing.T, config *models.ProductConfig) models.ProductJSON {
	jsonStr := CreateProductJSON(*config)
	if jsonStr == types.EmptyString {
		t.Error(errorEmptyJSONString)
	}
	var testObj models.ProductJSON

	err := json.Unmarshal([]byte(jsonStr), &testObj)
	if err != nil {
		t.Fatalf(errorUnmarshal, err)
	}
	return testObj
}

func newProductConfig() models.ProductConfig {
	return models.ProductConfig{
		CommonConfig:      newCommonConfig(),
		ProductName:       "Test Product",
		AlertGroupID:      123,
		Status:            models.IDNameConfig{ID: 0, Name: "Active"},
		TestDataWebhookID: 456,
	}
}

//#endregion
