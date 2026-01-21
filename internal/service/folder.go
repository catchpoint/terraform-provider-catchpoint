package service

import (
	"encoding/json"

	"catchpoint-provider/internal/labels"
	"catchpoint-provider/internal/models"
)

func CreateJSONFolderPatchDocument(config *models.FolderConfigUpdate, path string, isFolderMetaData bool) string {
	return createJSONPatchDocument(config, path, isFolderMetaData)
}

func CreateFolderJSON(config models.FolderConfig) string {
	testFolderTypeId := models.GenericIDNameJSON{ID: 1, Name: labels.Synthetic}

	alertGroup := buildAlertSettings(&config.CommonConfig.AlertSettingsConfig)
	insightData := buildInsightSettings(&config.CommonConfig.InsightSettingsConfig)
	scheduleSettings := buildScheduleSettings(&config.CommonConfig.ScheduleSettingsConfig)
	requestSettings := buildRequestSettings(&config.CommonConfig.RequestSettingsConfig)
	advancedSettings := buildAdvancedSettings(&config.CommonConfig.AdvancedSettingsConfig)

	folderId := 0

	var folder = models.FolderJSON{
		ID:               folderId,
		DivisionID:       config.DivisionID,
		ProductID:        config.ProductID,
		Name:             config.FolderName,
		TestFolderTypeID: testFolderTypeId,
		ScheduleSettings: scheduleSettings,
		AlertGroup:       alertGroup,
		RequestSettings:  requestSettings,
		InsightData:      insightData,
		AdvancedSettings: advancedSettings,
	}

	// Only set if not 0.
	if config.ParentID != 0 {
		folder.ParentID = &config.ParentID
	}

	folderJSON, _ := json.Marshal(folder)
	return string(folderJSON)
}

func RequestSettingFromFolderConfig(config *models.FolderConfig) models.RequestSettingsJSON {
	return buildRequestSettings(&config.CommonConfig.RequestSettingsConfig)
}

func AdvancedSettingFromFolderConfig(config *models.FolderConfig) models.AdvancedSettingsJSON {
	return buildAdvancedSettings(&config.CommonConfig.AdvancedSettingsConfig)
}

func AlertSettingsFromFolderConfig(config *models.FolderConfig) models.AlertGroupJSON {
	return buildAlertSettings(&config.CommonConfig.AlertSettingsConfig)
}

func InsightDataFromFolderConfig(config *models.FolderConfig) models.InsightDataJSON {
	return buildInsightSettings(&config.CommonConfig.InsightSettingsConfig)
}

func ScheduleSettingsFromFolderConfig(config *models.FolderConfig) models.ScheduleSettingsJSON {
	return buildScheduleSettings(&config.CommonConfig.ScheduleSettingsConfig)
}
