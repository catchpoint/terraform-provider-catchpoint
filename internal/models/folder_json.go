package models

type FolderData struct {
	Folders []FolderJSON `json:"folders"`
}

type FolderJSON struct {
	ID               int                  `json:"id"`
	DivisionID       int                  `json:"divisionId"`
	ProductID        int                  `json:"productId"`
	ParentID         *int                 `json:"parentId,omitempty"`
	Name             string               `json:"name"`
	TestFolderTypeID GenericIDNameJSON    `json:"testFolderTypeId"`
	RequestSettings  RequestSettingsJSON  `json:"requestSetting"`
	AlertGroup       AlertGroupJSON       `json:"alertGroup"`
	InsightData      InsightDataJSON      `json:"insights"`
	ScheduleSettings ScheduleSettingsJSON `json:"scheduleSetting"`
	AdvancedSettings AdvancedSettingsJSON `json:"advancedSettings"`
}

type FolderResponse struct {
	ResponseData FolderData `json:"data"`
	CommonResponse
}

type DeleteFolderResponse struct {
	ResponseData DeleteDataJSON `json:"data"`
	CommonResponse
}
