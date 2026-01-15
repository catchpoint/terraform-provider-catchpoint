package service

import (
	"catchpoint-provider/internal/models"
	"catchpoint-provider/internal/testutil"
	"catchpoint-provider/internal/types"
	"encoding/json"
	"testing"
)

const (
	folderNameSection = "/folderName"
)

// #region CreateTests

func TestCreateFolderJSONMinimal(t *testing.T) {
	config := newFolderConfig()
	testObj := unmarshalFolderConfigToTest(t, &config)

	compareAllFolderFields(t, &testObj, &config)
}

func TestCreateFolderJSONFull(t *testing.T) {
	config := newFolderConfig()
	config.CommonConfig = newCommonConfig()
	testObj := unmarshalFolderConfigToTest(t, &config)

	compareAllFolderFields(t, &testObj, &config)
	checkAlertGroupFields(t, &testObj.AlertGroup, &config.CommonConfig)
}

func TestCreateFolderJSONNoNodeGroupNotNil(t *testing.T) {
	config := newFolderConfig()
	config.CommonConfig = newCommonConfig()
	config.ScheduleSettingsConfig.NodeGroupIDs = []models.IDNameConfig{}
	testObj := unmarshalFolderConfigToTest(t, &config)

	compareAllFolderFields(t, &testObj, &config)
	checkAlertGroupFields(t, &testObj.AlertGroup, &config.CommonConfig)
	testutil.AssertNotNil(t, "NodeGroupIDs", testObj.ScheduleSettings.NodeGroups)
}

func TestCreateFolderJSONNoRecipientsNotNil(t *testing.T) {
	config := newFolderConfig()
	config.CommonConfig = newCommonConfig()
	config.CommonConfig.AlertSettingsConfig.NotificationGroup.Emails = []string{}
	config.CommonConfig.AlertSettingsConfig.NotificationGroup.ContactGroupIDs = []int{}
	testObj := unmarshalFolderConfigToTest(t, &config)

	compareAllFolderFields(t, &testObj, &config)
	checkAlertGroupFields(t, &testObj.AlertGroup, &config.CommonConfig)
	testutil.AssertNotNil(t, "TopRecipients", testObj.AlertGroup.NotificationGroup.Recipients)
	testutil.AssertNotNil(t, "NestedRecipients", &testObj.AlertGroup.AlertGroupItems[0].NotificationGroups[0].Recipients)
}

func TestCreateJSONFolderPatchDocument(t *testing.T) {
	update := &models.FolderConfigUpdate{
		UpdatedFieldValue: "foo",
		SectionToUpdate:   folderNameSection,
	}
	jsonStr := CreateJSONFolderPatchDocument(update, folderNameSection, true)
	if jsonStr == types.EmptyString {
		t.Error(errorEmptyJSONString)
	}
	var patch map[string]any
	if err := json.Unmarshal([]byte(jsonStr), &patch); err != nil {
		t.Errorf(errorUnmarshal, err)
	}
	if patch["path"] != folderNameSection {
		t.Errorf("Expected path '%s', got %v", folderNameSection, patch["path"])
	}
}

func TestCreateFolderJSONWithInsightSettings(t *testing.T) {
	config := newFolderConfig()

	config.CommonConfig = newCommonConfig()
	config.CommonConfig.InsightSettingsConfig.InsightSettingType = models.IDNameConfig{ID: 1, Name: "Override"}
	config.CommonConfig.InsightSettingsConfig.TracepointIDs = []int{1, 2, 3}
	config.CommonConfig.InsightSettingsConfig.IndicatorIDs = []int{4, 5, 6}

	testObj := unmarshalFolderConfigToTest(t, &config)

	compareAllFolderFields(t, &testObj, &config)
	compareInsightDataStructFields(t, &testObj.InsightData, &config.CommonConfig)
}

func TestCreateFolderJSONWithScheduleSettings(t *testing.T) {
	config := newFolderConfig()
	config.CommonConfig = newCommonConfig()
	config.ScheduleSettingsConfig.NodeIDs = []int{1, 2, 3}
	config.ScheduleSettingsConfig.NodeGroupIDs = []models.IDNameConfig{
		{ID: 4},
		{ID: 5},
		{ID: 6},
	}
	config.ScheduleSettingsConfig.NoOfSubsetNodes = 2

	testObj := unmarshalFolderConfigToTest(t, &config)

	compareAllFolderFields(t, &testObj, &config)
	compareScheduleSettingSectionFields(t, &testObj.ScheduleSettings, &config.CommonConfig)
}

func TestCreateFolderJSONWithRequestSettings(t *testing.T) {
	config := newFolderConfig()
	config.CommonConfig.RequestSettingsConfig.RequestSettingType.ID = 1
	config.CommonConfig.RequestSettingsConfig.TestHTTPHeaderRequests = newTestHTTPHeaders()
	config.CommonConfig.RequestSettingsConfig.AuthenticationType = models.IDNameConfig{
		ID:   3,
		Name: "Basic",
	}

	testObj := unmarshalFolderConfigToTest(t, &config)

	compareAllFolderFields(t, &testObj, &config)
	compareRequestSettingFields(t, &testObj.RequestSettings, &config.CommonConfig)
}

func TestCreateFolderJSONWithAdvancedSettings(t *testing.T) {
	config := newFolderConfig()
	config.CommonConfig = newCommonConfigWithAdvancedSettings()

	testObj := unmarshalFolderConfigToTest(t, &config)
	compareAllFolderFields(t, &testObj, &config)
	compareAdvancedSettingsFields(t, &testObj.AdvancedSettings, &config.CommonConfig)
}

// #endregion

// #region Helpers

func compareAllFolderFields(t *testing.T, testObj *models.FolderJSON, config *models.FolderConfig) {
	testutil.AssertEqual(t, "FolderName", testObj.Name, config.FolderName)
	testutil.AssertEqual(t, "DivisionID", testObj.DivisionID, config.DivisionID)
	testutil.AssertEqual(t, "ProductID", testObj.ProductID, config.ProductID)

	if testObj.ParentID != nil {
		testutil.AssertEqual(t, "ParentID", *testObj.ParentID, config.ParentID)
	} else {
		testutil.AssertNil(t, "ParentID", config.ParentID)
	}
}

func unmarshalFolderConfigToTest(t *testing.T, config *models.FolderConfig) models.FolderJSON {
	jsonStr := CreateFolderJSON(*config)
	if jsonStr == types.EmptyString {
		t.Error(errorEmptyJSONString)
	}
	var testObj models.FolderJSON
	err := json.Unmarshal([]byte(jsonStr), &testObj)
	if err != nil {
		t.Fatalf(errorUnmarshal, err)
	}
	return testObj
}

func newFolderConfig() models.FolderConfig {
	return models.FolderConfig{
		FolderName: "Test Folder",
		ParentID:   1,
		ProductID:  3,
	}
}

//#endregion
