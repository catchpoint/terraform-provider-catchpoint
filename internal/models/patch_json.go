package models

type JSONPatch struct {
	Value string `json:"value"`
	Path  string `json:"path"`
	Op    string `json:"op"`
}

type JSONPatchChromeVersion struct {
	ChromeVersionValue ChromeMonitorVersionStructJSON `json:"value"`
	Path               string                         `json:"path"`
	Op                 string                         `json:"op"`
}

type JSONPatchAdvanced struct {
	AdvancedSettingValue AdvancedSettingsJSON `json:"value"`
	Path                 string               `json:"path"`
	Op                   string               `json:"op"`
}

type JSONPatchRequest struct {
	RequestSettingValue RequestSettingsJSON `json:"value"`
	Path                string              `json:"path"`
	Op                  string              `json:"op"`
}

type JSONPatchSchedule struct {
	ScheduleSettingValue any    `json:"value"`
	Path                 string `json:"path"`
	Op                   string `json:"op"`
}

type JSONPatchInsight struct {
	InsightDataValue []map[string]int `json:"value"`
	Path             string           `json:"path"`
	Op               string           `json:"op"`
}

type JSONPatchAlert struct {
	AlertSettingValue AlertGroupJSON `json:"value"`
	Path              string         `json:"path"`
	Op                string         `json:"op"`
}

type JSONPatchLabel struct {
	LabelValue []LabelsJSON `json:"value"`
	Path       string       `json:"path"`
	Op         string       `json:"op"`
}

type JSONPatchThreshold struct {
	ThresholdValue TestThresholdsJSON `json:"value"`
	Path           string             `json:"path"`
	Op             string             `json:"op"`
}

type JSONPatchRequestData struct {
	TestRequestDataValue TestRequestDataJSON `json:"value"`
	Path                 string              `json:"path"`
	Op                   string              `json:"op"`
}
