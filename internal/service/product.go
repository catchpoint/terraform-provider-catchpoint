package service

import (
	"encoding/json"

	"catchpoint-provider/internal/models"
)

func CreateJSONProductPatchDocument(config *models.ProductConfigUpdate, path string, isProductMetaData bool) string {
	return createJSONPatchDocument(config, path, isProductMetaData)
}

func CreateProductJSON(config models.ProductConfig) string {
	status := models.GenericIDNameJSON{ID: config.Status.ID, Name: config.Status.Name}

	alertGroup := buildAlertSettings(&config.CommonConfig.AlertSettingsConfig)
	insightData := buildInsightSettings(&config.CommonConfig.InsightSettingsConfig)
	scheduleSettings := buildScheduleSettings(&config.CommonConfig.ScheduleSettingsConfig)
	requestSettings := buildRequestSettings(&config.CommonConfig.RequestSettingsConfig)
	advancedSettings := buildAdvancedSettings(&config.CommonConfig.AdvancedSettingsConfig)

	productId := 0

	var product = models.ProductJSON{
		ID:                productId,
		DivisionID:        config.DivisionID,
		Name:              config.ProductName,
		Status:            status,
		TestDataWebhookID: config.TestDataWebhookID,
		AlertGroupID:      config.AlertGroupID,
		ScheduleSettings:  scheduleSettings,
		AlertGroup:        alertGroup,
		RequestSettings:   requestSettings,
		InsightData:       insightData,
		AdvancedSettings:  advancedSettings,
	}

	productJSON, _ := json.Marshal(product)
	return string(productJSON)
}

func RequestSettingFromProductConfig(config *models.ProductConfig) models.RequestSettingsJSON {
	return buildRequestSettings(&config.CommonConfig.RequestSettingsConfig)
}

func AdvancedSettingFromProductConfig(config *models.ProductConfig) models.AdvancedSettingsJSON {
	return buildAdvancedSettings(&config.CommonConfig.AdvancedSettingsConfig)
}

func AlertSettingsFromProductConfig(config *models.ProductConfig) models.AlertGroupJSON {
	return buildAlertSettings(&config.CommonConfig.AlertSettingsConfig)
}

func InsightDataFromProductConfig(config *models.ProductConfig) models.InsightDataJSON {
	return buildInsightSettings(&config.CommonConfig.InsightSettingsConfig)
}

func ScheduleSettingsFromProductConfig(config *models.ProductConfig) models.ScheduleSettingsJSON {
	return buildScheduleSettings(&config.CommonConfig.ScheduleSettingsConfig)
}
