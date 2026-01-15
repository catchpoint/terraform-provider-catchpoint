package models

type ProductDataJSON struct {
	Products []ProductJSON `json:"products"`
}

type ProductJSON struct {
	ID                int                  `json:"id"`
	DivisionID        int                  `json:"divisionId"`
	Name              string               `json:"name"`
	Status            GenericIDNameJSON    `json:"status"`
	AlertGroupID      int                  `json:"alertGroupId,omitempty"`
	TestDataWebhookID int                  `json:"testDataWebhookId,omitempty"`
	RequestSettings   RequestSettingsJSON  `json:"requestSettings"`
	AlertGroup        AlertGroupJSON       `json:"alertGroup"`
	InsightData       InsightDataJSON      `json:"insightsData"`
	ScheduleSettings  ScheduleSettingsJSON `json:"scheduleSettings"`
	AdvancedSettings  AdvancedSettingsJSON `json:"advancedSettingsModel"`
}

type ProductResponse struct {
	ResponseData ProductDataJSON `json:"data"`
	CommonResponse
}

type DeleteProductResponse struct {
	ResponseData DeleteDataJSON `json:"data"`
	CommonResponse
}
