package catchpoint

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/google/go-cmp/cmp"
)

type GenericIdName struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}
type GenericIdNameOmitEmpty struct {
	Id   int    `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type Label struct {
	Color  string   `json:"color"`
	Name   string   `json:"name"`
	Values []string `json:"values"`
}

type Thresholds struct {
	TestTimeApdexThresholdWarning      float64 `json:"testTimeApdexThresholdWarning,omitempty"`
	TestTimeApdexThresholdCritical     float64 `json:"testTimeApdexThresholdCritical,omitempty"`
	AvailabilityApdexThresholdWarning  float64 `json:"availabilityApdexThresholdWarning,omitempty"`
	AvailabilityApdexThresholdCritical float64 `json:"availabilityApdexThresholdCritical,omitempty"`
}

type AuthenticationStruct struct {
	AuthenticationMethodType GenericIdNameOmitEmpty `json:"authenticationMethodType,omitempty"`
	PasswordStatus           GenericIdNameOmitEmpty `json:"passwordStatus,omitempty"`
	UserName                 string                 `json:"userName,omitempty"`
	Password                 string                 `json:"password,omitempty"`
	PasswordIds              []int                  `json:"passwordIds,omitempty"`
	//Id                       int                    `json:"id,omitempty"`
}

type HttpHeaderRequest struct {
	RequestValue      string        `json:"requestValue"`
	RequestHeaderType GenericIdName `json:"requestHeaderType"`
	ChildHostPattern  string        `json:"childHostPattern,omitempty"`
}

type RequestSetting struct {
	RequestSettingType    GenericIdName         `json:"requestSettingType"`
	Authentication        *AuthenticationStruct `json:"authentication,omitempty"`
	LibraryCertificateIds []int                 `json:"libraryCertificateIds,omitempty"`
	TokenIds              []int                 `json:"tokenIds,omitempty"`
	HttpHeaderRequests    []HttpHeaderRequest   `json:"httpHeaderRequests,omitempty"`
}

type NodeThresholdStruct struct {
	Id                              int           `json:"id"`
	Name                            string        `json:"name"`
	NodeThresholdType               GenericIdName `json:"nodeThresholdType"`
	NumberOfUnits                   int           `json:"numberOfUnits,omitempty"`
	PercentageOfUnits               float64       `json:"percentageOfUnits,omitempty"`
	NumberOfFailingUnits            int           `json:"numberOfFailingUnits,omitempty"`
	ConsecutiveRunsEnabled          bool          `json:"consecutiveRunsEnabled"`
	UtilizePerNodeHistoricalAverage bool          `json:"utilizePerNodeHistoricalAverage"`
	NumberOfConsecutiveRuns         int           `json:"consecutiveRuns,omitempty"`
}

type TriggerStruct struct {
	Id                        int                     `json:"id"`
	WarningReminderFrequency  GenericIdName           `json:"warningReminderFrequency"`
	CriticalReminderFrequency GenericIdName           `json:"criticalReminderFrequency"`
	TriggerType               GenericIdName           `json:"triggerType"`
	OperationType             GenericIdName           `json:"operationType"`
	StatisticalType           *GenericIdNameOmitEmpty `json:"statisticalType,omitempty"`
	HistoricalInterval        *GenericIdNameOmitEmpty `json:"historicalInterval,omitempty"`
	ThresholdInterval         GenericIdName           `json:"thresholdInterval"`
	WarningTrigger            float64                 `json:"warningTrigger,omitempty"`
	CriticalTrigger           float64                 `json:"criticalTrigger,omitempty"`
	UseIntervalRollingWindow  bool                    `json:"useIntervalRollingWindow"`
	Expression                string                  `json:"expression,omitempty"`
}

type AlertGroupItem struct {
	NodeThreshold      NodeThresholdStruct       `json:"nodeThreshold"`
	Trigger            TriggerStruct             `json:"trigger"`
	NotificationType   GenericIdName             `json:"notificationType"`
	AlertType          GenericIdName             `json:"alertType"`
	AlertSubType       *GenericIdNameOmitEmpty   `json:"alertSubType,omitempty"`
	EnforceTestFailure bool                      `json:"enforceTestFailure"`
	OmitScatterplot    bool                      `json:"omitScatterplot"`
	MatchAllRecords    bool                      `json:"matchAllRecords"`
	NotificationGroups []NotificationGroupStruct `json:"notificationGroups"`
}

type AlertWebhook struct {
	Id int `json:"id,omitempty"`
}

type Recipient struct {
	Id            int           `json:"id,omitempty"`
	Email         string        `json:"email"`
	RecipientType GenericIdName `json:"recipientType"`
	Name          string        `json:"name"`
}

type NotificationGroupStruct struct {
	Subject          string         `json:"subject"`
	NotifyOnWarning  bool           `json:"notifyOnWarning"`
	NotifyOnCritical bool           `json:"notifyOnCritical"`
	NotifyOnImproved bool           `json:"notifyOnImproved"`
	AlertWebhooks    []AlertWebhook `json:"alertWebhooks"`
	Recipients       []Recipient    `json:"recipients"`
}

type AlertGroupStruct struct {
	AlertSettingType  GenericIdName           `json:"alertSettingType"`
	NotificationGroup NotificationGroupStruct `json:"notificationGroup"`
	AlertGroupItems   []AlertGroupItem        `json:"alertGroupItems"`
}

type InsightDataStruct struct {
	InsightSettingType GenericIdName   `json:"insightSettingType"`
	Indicators         []GenericIdName `json:"indicators,omitempty"`
	Tracepoints        []GenericIdName `json:"tracepoints,omitempty"`
}

type Node struct {
	Id          int           `json:"id,omitempty"`
	Name        string        `json:"name"`
	NetworkType GenericIdName `json:"networkType"`
}

type NodeGroup struct {
	Id                   int           `json:"id,omitempty"`
	Name                 string        `json:"name"`
	Description          string        `json:"description"`
	SyntheticNetworkType GenericIdName `json:"syntheticNetworkType"`
	Nodes                []Node        `json:"nodes"`
}

type ScheduleSetting struct {
	ScheduleSettingType   GenericIdName `json:"scheduleSettingType"`
	RunScheduleId         int           `json:"runScheduleId,omitempty"`
	MaintenanceScheduleId int           `json:"maintenanceScheduleId,omitempty"`
	Frequency             GenericIdName `json:"frequency"`
	TestNodeDistribution  GenericIdName `json:"testNodeDistribution"`
	NetworkType           GenericIdName `json:"networkType"`
	Nodes                 []Node        `json:"nodes"`
	NodeGroups            []NodeGroup   `json:"nodeGroups"`
	NoOfSubsetNodes       int           `json:"roundRobinAmount,omitempty"`
	Id                    int           `json:"id"`
}

type AdvancedSetting struct {
	AdvancedSettingType       GenericIdName            `json:"advancedSettingType"`
	AppliedTestFlags          []GenericIdNameOmitEmpty `json:"appliedTestFlags"`
	MaxStepRuntimeSecOverride int                      `json:"maxStepRuntimeSecOverride,omitempty"`
	WaitForNoActivity         *int                     `json:"waitForNoActivity,omitempty"`
	ViewportHeight            int                      `json:"viewportHeight,omitempty"`
	ViewportWidth             int                      `json:"viewportWidth,omitempty"`
	FailureHopCount           int                      `json:"failureHopCount,omitempty"`
	PingCount                 int                      `json:"pingCount,omitempty"`
	EdnsSubnet                string                   `json:"ednsSubnet,omitempty"`
	AdditionalMonitor         *GenericIdNameOmitEmpty  `json:"additionalMonitor,omitempty"`
	TestBandwidthThrottling   *GenericIdNameOmitEmpty  `json:"testBandwidthThrottling,omitempty"`
	Id                        int                      `json:"id"`
}

type ChromeMonitorVersionStruct struct {
	ApplicationVersionType GenericIdNameOmitEmpty `json:"applicationVersionType,omitempty"`
	ApplicationVersionId   int                    `json:"applicationVersionId,omitempty"`
}

type TestRequestDataStruct struct {
	TestId                int                    `json:"testId,omitempty"`
	RequestData           string                 `json:"requestData,omitempty"`
	TransactionScriptType GenericIdNameOmitEmpty `json:"transactionScriptType,omitempty"`
	TestType              GenericIdNameOmitEmpty `json:"testType,omitempty"`
	Monitor               GenericIdNameOmitEmpty `json:"monitor,omitempty"`
}

type Test struct {
	Id                           int                         `json:"id"`
	DivisionId                   int                         `json:"divisionId"`
	ProductId                    int                         `json:"productId"`
	FolderId                     int                         `json:"folderId,omitempty"`
	Name                         string                      `json:"name"`
	Description                  string                      `json:"description"`
	Url                          string                      `json:"url"`
	GatewayAddressOrHost         string                      `json:"gatewayAddressOrHost,omitempty"`
	Labels                       []Label                     `json:"labels,omitempty"`
	TestThresholds               Thresholds                  `json:"thresholdRestModel,omitempty"`
	EnforceCertificatePinning    bool                        `json:"enforceCertificatePinning"`
	EnforceCertificateKeyPinning bool                        `json:"enforceCertificateKeyPinning"`
	EnableTestDataWebhook        bool                        `json:"enableTestDataWebhook"`
	AlertsPaused                 bool                        `json:"alertsPaused"`
	ChangeDate                   string                      `json:"changeDate"`
	StartTime                    string                      `json:"startTime"`
	EndTime                      string                      `json:"endTime"`
	Status                       GenericIdName               `json:"status"`
	Monitor                      GenericIdName               `json:"monitor"`
	DnsServer                    string                      `json:"dnsServer,omitempty"`
	DnsQueryType                 *GenericIdNameOmitEmpty     `json:"dnsQueryType,omitempty"`
	UserAgentType                *GenericIdNameOmitEmpty     `json:"userAgentTypeId,omitempty"`
	ChromeMonitorVersion         *ChromeMonitorVersionStruct `json:"chromeMonitorVersion,omitempty"`
	TestRequestData              *TestRequestDataStruct      `json:"testRequestData,omitempty"`
	TestType                     GenericIdName               `json:"testType"`
	RequestHttpMethod            GenericIdName               `json:"requestHttpMethod"`
	RequestSettings              RequestSetting              `json:"requestSettings"`
	AlertGroup                   AlertGroupStruct            `json:"alertGroup"`
	InsightData                  InsightDataStruct           `json:"insightData"`
	ScheduleSettings             ScheduleSetting             `json:"scheduleSettings"`
	AdvancedSettings             AdvancedSetting             `json:"advancedSettings"`
}

func createJson(config TestConfig) string {

	//Set properties
	status := GenericIdName{Id: config.TestStatus, Name: "Active"}
	monitor := GenericIdName{Id: config.Monitor, Name: "Object"}
	testType := GenericIdName{Id: config.TestType, Name: "Web"}
	requestHttpMethod := GenericIdName{Id: 0, Name: "Get"}

	alertGroup := setTestAlertSettings(&config)

	insightData := setTestInsightSettings(&config)

	scheduleSettings := setTestScheduleSettings(&config)

	advancedSettings := setTestAdvancedSettings(&config)

	labels := setTestLabels(&config)

	thresholds := setTestThresholds(&config)

	requestSettings := setTestRequestSettings(&config)

	testId := 0
	changeDate := getTime()
	var t = Test{}

	t = Test{Id: testId, DivisionId: config.DivisionId, ProductId: config.ProductId, FolderId: config.FolderId, Name: config.TestName, Description: config.TestDescription, Url: config.TestUrl, GatewayAddressOrHost: config.GatewayAddressOrHost, Labels: labels, TestThresholds: thresholds, EnforceCertificatePinning: config.EnforceCertificatePinning, EnforceCertificateKeyPinning: config.EnforceCertificateKeyPinning, EnableTestDataWebhook: config.EnableTestDataWebhook, AlertsPaused: config.AlertsPaused, ChangeDate: changeDate, StartTime: config.StartTime, EndTime: config.EndTime, Status: status, Monitor: monitor, TestType: testType, RequestHttpMethod: requestHttpMethod, RequestSettings: requestSettings, AlertGroup: alertGroup, InsightData: insightData, ScheduleSettings: scheduleSettings, AdvancedSettings: advancedSettings}

	userAgentType := GenericIdNameOmitEmpty{Id: config.SimulateDevice, Name: ""}
	if userAgentType != (GenericIdNameOmitEmpty{}) {
		t.UserAgentType = &userAgentType
	}
	applicationVersionType := GenericIdNameOmitEmpty{Id: config.ChromeVersion.Id, Name: config.ChromeVersion.Name}
	chromeMonitor := ChromeMonitorVersionStruct{ApplicationVersionType: applicationVersionType, ApplicationVersionId: config.ChromeApplicationVersion.Id}
	if chromeMonitor != (ChromeMonitorVersionStruct{}) {
		t.ChromeMonitorVersion = &chromeMonitor
	}

	if testType.Id == int(TestType(Dns)) {
		dnsQueryType := GenericIdNameOmitEmpty{Id: config.DnsQueryType.Id, Name: config.DnsQueryType.Name}
		if dnsQueryType != (GenericIdNameOmitEmpty{}) {
			t.DnsQueryType = &dnsQueryType
		}
		t.DnsServer = config.DnsServer
	}

	requestData := setTestRequestData(&config)

	if testType.Id == int(TestType(Api)) ||
		testType.Id == int(TestType(Transaction)) ||
		testType.Id == int(TestType(Playwright)) ||
		testType.Id == int(TestType(Puppeteer)) {
		t.TestRequestData = &requestData
	}

	testJson, _ := json.Marshal(t)
	return string(testJson)
}

func getTest(apiToken string, testId string) (*Test, string, error) {

	type Data struct {
		Tests []Test `json:"tests"`
	}
	type ApiError struct {
		Id      json.Number `json:"id"`
		Message string      `json:"message"`
	}
	type Response struct {
		ResponseData Data       `json:"data"`
		Messages     []string   `json:"messages"`
		Errors       []ApiError `json:"errors"`
		Completed    bool       `json:"completed"`
		TraceId      string     `json:"traceId"`
	}

	var response Response
	var responseStatus = ""
	getURL := catchpointTestURI + "/" + testId + "?showInheritedProperties=false"
	req, _ := http.NewRequest("", getURL, nil)
	req.Header.Set("Authorization", "Bearer "+apiToken)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return &Test{}, responseStatus, err
	}
	defer resp.Body.Close()

	responseStatus = strings.ToLower(string(resp.Status))
	body, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal([]byte(body), &response)
	//Test not found
	if !response.Completed {
		return nil, responseStatus, err
	}
	test := response.ResponseData.Tests[0]

	return &test, responseStatus, nil
}

func createTest(apiToken string, jsonPayload string) (string, string, string, error) {

	type Data struct {
		Id json.Number `json:"id"`
	}
	type ApiError struct {
		Id      json.Number `json:"id"`
		Message string      `json:"message"`
	}
	type Response struct {
		ResponseData Data       `json:"data"`
		Messages     []string   `json:"messages"`
		Errors       []ApiError `json:"errors"`
		Completed    bool       `json:"completed"`
		TraceId      string     `json:"traceId"`
	}

	var response Response
	var postBody = []byte(jsonPayload)
	var responseBody = ""
	var responseStatus = ""
	var testId = ""
	req, _ := http.NewRequest("POST", catchpointTestURI, bytes.NewBuffer(postBody))
	req.Header.Set("Authorization", "Bearer "+apiToken)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return responseBody, responseStatus, testId, err
	}
	defer resp.Body.Close()

	responseStatus = strings.ToLower(string(resp.Status))
	body, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal([]byte(body), &response)
	testId = string(response.ResponseData.Id)

	return string(body), responseStatus, testId, nil
}

func deleteTest(apiToken string, testId string) (string, string, bool, error) {

	type Data struct {
		Id string `json:"deleted"`
	}
	type ApiError struct {
		Id      json.Number `json:"id"`
		Message string      `json:"message"`
	}
	type Response struct {
		ResponseData Data       `json:"data"`
		Messages     []string   `json:"messages"`
		Errors       []ApiError `json:"errors"`
		Completed    bool       `json:"completed"`
		TraceId      string     `json:"traceId"`
	}

	deleteURL := catchpointTestURI + "/" + testId
	var response Response
	var responseBody = ""
	var responseStatus = ""
	var completed = false
	req, _ := http.NewRequest("DELETE", deleteURL, nil)
	req.Header.Set("Authorization", "Bearer "+apiToken)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return responseBody, responseStatus, completed, err
	}
	defer resp.Body.Close()

	responseStatus = strings.ToLower(string(resp.Status))
	body, _ := ioutil.ReadAll(resp.Body)

	json.Unmarshal([]byte(body), &response)
	completed = response.Completed

	return string(body), responseStatus, completed, nil
}

func setTestRequestData(config *TestConfig) TestRequestDataStruct {
	monitor := GenericIdName{Id: config.Script.Monitor, Name: "string"}
	testType := GenericIdName{Id: config.Script.TestType, Name: "string"}
	transactionScriptType := GenericIdNameOmitEmpty{Id: config.Script.TransactionScriptType, Name: "string"}
	requestData := TestRequestDataStruct{TestId: config.Script.TestId, RequestData: config.Script.RequestData, TransactionScriptType: transactionScriptType, Monitor: GenericIdNameOmitEmpty(monitor), TestType: GenericIdNameOmitEmpty(testType)}

	return requestData
}

func setTestThresholds(config *TestConfig) Thresholds {
	thresholds := Thresholds{TestTimeApdexThresholdWarning: config.TestTimeThresholdWarning, TestTimeApdexThresholdCritical: config.TestTimeThresholdCritical, AvailabilityApdexThresholdWarning: config.AvailabilityThresholdWarning, AvailabilityApdexThresholdCritical: config.AvailabilityThresholdCritical}

	return thresholds
}

func setTestLabels(config *TestConfig) []Label {
	labels := []Label{}

	if len(config.Labels) > 0 {
		for i := range config.Labels {
			labels = append(labels, Label{Color: randHexString(), Name: config.Labels[i].Name, Values: config.Labels[i].Values})
		}
	}

	return labels
}

func setTestAlertSettings(config *TestConfig) AlertGroupStruct {
	alertGroupItems := []AlertGroupItem{}
	recipients := []Recipient{}
	alertWebhooks := []AlertWebhook{}

	alertSettingType := GenericIdName{Id: config.AlertSettingType, Name: "Inherit"}

	for i := range config.AlertRuleConfigs {
		nodeThresholdType := GenericIdName{Id: config.AlertRuleConfigs[i].AlertNodeThresholdType.Id, Name: config.AlertRuleConfigs[i].AlertNodeThresholdType.Name}
		nodeThreshold := NodeThresholdStruct{Id: 0, Name: "", NodeThresholdType: nodeThresholdType, NumberOfUnits: config.AlertRuleConfigs[i].AlertThresholdNumOfRuns, PercentageOfUnits: config.AlertRuleConfigs[i].AlertThresholdPercentOfRuns, NumberOfFailingUnits: config.AlertRuleConfigs[i].AlertThresholdNumOfFailingNodes, ConsecutiveRunsEnabled: config.AlertRuleConfigs[i].AlertEnableConsecutive, UtilizePerNodeHistoricalAverage: false, NumberOfConsecutiveRuns: config.AlertRuleConfigs[i].AlertConsecutiveNumOfRuns}
		warningTrigger := config.AlertRuleConfigs[i].AlertWarningTrigger
		criticalTrigger := config.AlertRuleConfigs[i].AlertCriticalTrigger
		warningReminder := GenericIdName{Id: config.AlertRuleConfigs[i].AlertWarningReminder.Id, Name: config.AlertRuleConfigs[i].AlertWarningReminder.Name}
		criticalReminder := GenericIdName{Id: config.AlertRuleConfigs[i].AlertCriticalReminder.Id, Name: config.AlertRuleConfigs[i].AlertCriticalReminder.Name}
		triggerType := GenericIdName{Id: config.AlertRuleConfigs[i].TriggerType.Id, Name: config.AlertRuleConfigs[i].TriggerType.Name}
		operationType := GenericIdName{Id: config.AlertRuleConfigs[i].OperationType.Id, Name: config.AlertRuleConfigs[i].OperationType.Name}
		statisticalType := GenericIdNameOmitEmpty{Id: config.AlertRuleConfigs[i].StatisticalType.Id, Name: config.AlertRuleConfigs[i].StatisticalType.Name}
		historicalInterval := GenericIdNameOmitEmpty{Id: config.AlertRuleConfigs[i].TrailingHistoricalInterval.Id, Name: config.AlertRuleConfigs[i].TrailingHistoricalInterval.Name}
		thresholdInterval := GenericIdName{Id: config.AlertRuleConfigs[i].AlertThresholdInterval.Id, Name: config.AlertRuleConfigs[i].AlertThresholdInterval.Name}

		trigger := TriggerStruct{Id: 0, WarningReminderFrequency: warningReminder, CriticalReminderFrequency: criticalReminder, Expression: config.AlertRuleConfigs[i].Expression, TriggerType: triggerType, OperationType: operationType, ThresholdInterval: thresholdInterval, UseIntervalRollingWindow: config.AlertRuleConfigs[i].AlertUseRollingWindow, WarningTrigger: warningTrigger, CriticalTrigger: criticalTrigger}

		if statisticalType != (GenericIdNameOmitEmpty{}) && historicalInterval != (GenericIdNameOmitEmpty{}) {
			trigger.StatisticalType = &statisticalType
			trigger.HistoricalInterval = &historicalInterval
		}

		notificationType := GenericIdName{Id: config.AlertRuleConfigs[i].AlertNotificationType, Name: "DefaultContacts"}
		alertType := GenericIdName{Id: config.AlertRuleConfigs[i].AlertType.Id, Name: config.AlertRuleConfigs[i].AlertType.Name}
		alertSubType := GenericIdNameOmitEmpty{Id: config.AlertRuleConfigs[i].AlertSubType.Id, Name: config.AlertRuleConfigs[i].AlertSubType.Name}
		notificationGroups := config.AlertRuleConfigs[i].NotificationGroups
		if alertSubType != (GenericIdNameOmitEmpty{}) {
			alertGroupItems = append(alertGroupItems, AlertGroupItem{NodeThreshold: nodeThreshold, Trigger: trigger, NotificationType: notificationType, AlertType: alertType, AlertSubType: &alertSubType, EnforceTestFailure: config.AlertRuleConfigs[i].AlertEnforceTestFailure, OmitScatterplot: config.AlertRuleConfigs[i].AlertOmitScatterplot, MatchAllRecords: false, NotificationGroups: notificationGroups})
		} else {
			alertGroupItems = append(alertGroupItems,
				AlertGroupItem{NodeThreshold: nodeThreshold,
					Trigger:            trigger,
					NotificationType:   notificationType,
					AlertType:          alertType,
					EnforceTestFailure: config.AlertRuleConfigs[i].AlertEnforceTestFailure,
					OmitScatterplot:    config.AlertRuleConfigs[i].AlertOmitScatterplot,
					MatchAllRecords:    false,
					NotificationGroups: notificationGroups,
				})
		}
	}

	if len(config.AlertRecipientEmails) > 0 {
		recipientType := GenericIdName{Id: 2, Name: "Email"}
		for i := range config.AlertRecipientEmails {
			recipients = append(recipients, Recipient{Email: config.AlertRecipientEmails[i], RecipientType: recipientType})
		}
	}

	if len(config.AlertContactGroups) > 0 {
		recipientType := GenericIdName{Id: 1, Name: "ContactGroup"}
		for i := range config.AlertContactGroups {
			recipients = append(recipients, Recipient{Id: i + 1, RecipientType: recipientType, Name: config.AlertContactGroups[i]})
		}
	}

	if len(config.AlertWebhookIds) > 0 {
		for i := range config.AlertWebhookIds {
			alertWebhooks = append(alertWebhooks, AlertWebhook{Id: config.AlertWebhookIds[i]})
		}
	}

	notifSubject := "${NotificationLevel}:  test=#${TestId} - ${TestName}, alert=${AlertType}"
	if config.AlertSubject != "" {
		notifSubject = config.AlertSubject
	}
	notificationGroup := NotificationGroupStruct{Subject: notifSubject, NotifyOnWarning: true, NotifyOnCritical: true, NotifyOnImproved: true, AlertWebhooks: alertWebhooks, Recipients: recipients}

	alertGroup := AlertGroupStruct{AlertSettingType: alertSettingType, NotificationGroup: notificationGroup, AlertGroupItems: alertGroupItems}

	return alertGroup
}

func setTestInsightSettings(config *TestConfig) InsightDataStruct {
	tracepoints := []GenericIdName{}
	indicators := []GenericIdName{}

	insightSettingType := GenericIdName{Id: config.InsightSettingType, Name: "Inherit"}

	if len(config.TracepointIds) > 0 {
		for i := range config.TracepointIds {
			tracepoints = append(tracepoints, GenericIdName{Id: config.TracepointIds[i], Name: "Tracepoint"})
		}
	}
	if len(config.IndicatorIds) > 0 {
		for i := range config.IndicatorIds {
			indicators = append(indicators, GenericIdName{Id: config.IndicatorIds[i], Name: "Indicator"})
		}
	}

	insightData := InsightDataStruct{InsightSettingType: insightSettingType, Indicators: indicators, Tracepoints: tracepoints}

	return insightData
}

func setTestScheduleSettings(config *TestConfig) ScheduleSetting {
	nodes := []Node{}
	scheduleSettingType := GenericIdName{Id: config.ScheduleSettingType, Name: "Inherit"}
	frequency := GenericIdName{Id: config.TestFrequency.Id, Name: config.TestFrequency.Name}
	testNodeDistribution := GenericIdName{Id: config.NodeDistribution.Id, Name: config.NodeDistribution.Name}
	networkType := GenericIdName{Id: 0, Name: "Backbone"}
	if len(config.NodeIds) > 0 {
		for i := range config.NodeIds {
			nodes = append(nodes, Node{Id: config.NodeIds[i], Name: "node", NetworkType: networkType})
		}
	}
	// Initialize an empty slice of NodeGroup
	var nodeGroups []NodeGroup
	if len(config.NodeGroupIds) > 0 {
		for i := range config.NodeGroupIds {
			nodeGroup := NodeGroup{
				Id:                   config.NodeGroupIds[i].Id,
				Name:                 "DefaultNodeGroupName",
				Description:          "",
				SyntheticNetworkType: networkType,
				Nodes:                []Node{{Id: 123, Name: "DefaultNodeName", NetworkType: networkType}},
			}
			nodeGroups = append(nodeGroups, nodeGroup)
		}
	}
	scheduleSettingId := 0
	scheduleSettings := ScheduleSetting{ScheduleSettingType: scheduleSettingType, RunScheduleId: config.ScheduleRunScheduleId, MaintenanceScheduleId: config.ScheduleMaintenanceScheduleId, Frequency: frequency, TestNodeDistribution: testNodeDistribution, NetworkType: networkType, Nodes: nodes, NodeGroups: nodeGroups, Id: scheduleSettingId}
	if config.NoOfSubsetNodes > 0 {
		scheduleSettings.NoOfSubsetNodes = config.NoOfSubsetNodes
	}

	return scheduleSettings
}

func setTestRequestSettings(config *TestConfig) RequestSetting {
	httpHeaderRequests := []HttpHeaderRequest{}
	requestSettingType := GenericIdName{Id: config.RequestSettingType, Name: "Inherit"}

	if len(config.TestHttpHeaderRequests) > 0 {
		for i := range config.TestHttpHeaderRequests {
			requestHeaderType := GenericIdName{Id: config.TestHttpHeaderRequests[i].RequestHeaderType.Id, Name: config.TestHttpHeaderRequests[i].RequestHeaderType.Name}
			httpHeaderRequests = append(httpHeaderRequests, HttpHeaderRequest{RequestValue: config.TestHttpHeaderRequests[i].RequestValue, RequestHeaderType: requestHeaderType, ChildHostPattern: config.TestHttpHeaderRequests[i].ChildHostPattern})
		}
	}

	var authentication = AuthenticationStruct{}
	// config.AuthenticationType == 0 indicates no authentication type(id) has been set.
	if config.AuthenticationType.Id != 0 {
		authenticationMethodType := GenericIdNameOmitEmpty{Id: config.AuthenticationType.Id, Name: config.AuthenticationType.Name}
		passwordStatus := GenericIdNameOmitEmpty{Id: 1, Name: "Set"}
		passwordIds := config.AuthenticationPasswordIds
		authentication = AuthenticationStruct{AuthenticationMethodType: authenticationMethodType, PasswordStatus: passwordStatus, UserName: config.UserName, Password: config.Password, PasswordIds: passwordIds}
	}

	requestSetting := RequestSetting{RequestSettingType: requestSettingType, HttpHeaderRequests: httpHeaderRequests, TokenIds: config.AuthenticationTokenIds, LibraryCertificateIds: config.AuthenticationCertificateIds}

	if !cmp.Equal(AuthenticationStruct{}, authentication) {
		requestSetting.Authentication = &authentication
	}
	return requestSetting
}

func setTestAdvancedSettings(config *TestConfig) AdvancedSetting {
	appliedTestFlags := []GenericIdNameOmitEmpty{}
	advancedSettingId := 0
	advancedSettingType := GenericIdName{Id: config.AdvancedSettingType, Name: "Override"}
	if len(config.AppliedTestFlags) > 0 {
		for i := range config.AppliedTestFlags {
			if config.AppliedTestFlags[i] != 0 {
				appliedTestFlags = append(appliedTestFlags, GenericIdNameOmitEmpty{Id: config.AppliedTestFlags[i], Name: "Flag"})
			}
		}
	}

	advancedSettings := AdvancedSetting{}
	advancedSettings = AdvancedSetting{AdvancedSettingType: advancedSettingType, AppliedTestFlags: appliedTestFlags, MaxStepRuntimeSecOverride: config.MaxStepRuntimeSecOverride, WaitForNoActivity: config.WaitForNoActivityOnDocComplete, ViewportHeight: config.ViewportHeight, ViewportWidth: config.ViewportWidth, FailureHopCount: config.TracerouteFailureHopCount, PingCount: config.TraceroutePingCount, EdnsSubnet: config.EdnsSubnet, Id: advancedSettingId}
	additionalMonitor := GenericIdNameOmitEmpty{Id: config.AdditionalMonitorType.Id, Name: config.AdditionalMonitorType.Name}
	if additionalMonitor != (GenericIdNameOmitEmpty{}) {
		advancedSettings.AdditionalMonitor = &additionalMonitor
	}
	bandwidthThrottling := GenericIdNameOmitEmpty{Id: config.BandwidthThrottling.Id, Name: config.BandwidthThrottling.Name}
	if bandwidthThrottling != (GenericIdNameOmitEmpty{}) {
		advancedSettings.TestBandwidthThrottling = &bandwidthThrottling
	}

	return advancedSettings
}

func createJsonPatchDocument(config TestConfigUpdate, path string, isTestMetaData bool) string {
	type JsonPatch struct {
		Value string `json:"value"`
		Path  string `json:"path"`
		Op    string `json:"op"`
	}

	type JsonPatchAdvanced struct {
		AdvancedSettingValue AdvancedSetting `json:"value"`
		Path                 string          `json:"path"`
		Op                   string          `json:"op"`
	}
	type JsonPatchRequest struct {
		RequestSettingValue RequestSetting `json:"value"`
		Path                string         `json:"path"`
		Op                  string         `json:"op"`
	}
	type JsonPatchSchedule struct {
		ScheduleSettingValue ScheduleSetting `json:"value"`
		Path                 string          `json:"path"`
		Op                   string          `json:"op"`
	}
	type JsonPatchInsight struct {
		InsightDataValue InsightDataStruct `json:"value"`
		Path             string            `json:"path"`
		Op               string            `json:"op"`
	}
	type JsonPatchAlert struct {
		AlertSettingValue AlertGroupStruct `json:"value"`
		Path              string           `json:"path"`
		Op                string           `json:"op"`
	}
	type JsonPatchLabel struct {
		LabelValue []Label `json:"value"`
		Path       string  `json:"path"`
		Op         string  `json:"op"`
	}
	type JsonPatchThreshold struct {
		ThresholdValue Thresholds `json:"value"`
		Path           string     `json:"path"`
		Op             string     `json:"op"`
	}
	type JsonPatchRequestData struct {
		TestRequestDataValue TestRequestDataStruct `json:"value"`
		Path                 string                `json:"path"`
		Op                   string                `json:"op"`
	}

	var jsonPatchDoc = []byte{}

	if isTestMetaData {
		jsonPatchObject := JsonPatch{
			Value: config.UpdatedFieldValue,
			Path:  path,
			Op:    "replace",
		}
		jsonPatchDoc, _ = json.Marshal(jsonPatchObject)
	}

	if config.SectionToUpdate == "/labels" {
		jsonPatchObject := JsonPatchLabel{
			LabelValue: config.UpdatedLabels,
			Path:       path,
			Op:         "replace",
		}
		jsonPatchDoc, _ = json.Marshal(jsonPatchObject)
	}
	if config.SectionToUpdate == "/thresholdRestModel" {
		jsonPatchObject := JsonPatchThreshold{
			ThresholdValue: config.UpdatedTestThresholds,
			Path:           path,
			Op:             "replace",
		}
		jsonPatchDoc, _ = json.Marshal(jsonPatchObject)
	}
	if config.SectionToUpdate == "/testRequestData" {
		jsonPatchObject := JsonPatchRequestData{
			TestRequestDataValue: config.UpdatedTestRequestData,
			Path:                 path,
			Op:                   "replace",
		}
		jsonPatchDoc, _ = json.Marshal(jsonPatchObject)
	}
	if config.SectionToUpdate == "/advancedSettings" {
		jsonPatchObject := JsonPatchAdvanced{
			AdvancedSettingValue: config.UpdatedAdvancedSettingsSection,
			Path:                 path,
			Op:                   "replace",
		}
		jsonPatchDoc, _ = json.Marshal(jsonPatchObject)
	}
	if config.SectionToUpdate == "/requestSettings" {
		jsonPatchObject := JsonPatchRequest{
			RequestSettingValue: config.UpdatedRequestSettingsSection,
			Path:                path,
			Op:                  "replace",
		}
		jsonPatchDoc, _ = json.Marshal(jsonPatchObject)
	}
	if config.SectionToUpdate == "/insightData" {
		jsonPatchObject := JsonPatchInsight{
			InsightDataValue: config.UpdatedInsightSettingsSection,
			Path:             path,
			Op:               "replace",
		}
		jsonPatchDoc, _ = json.Marshal(jsonPatchObject)
	}
	if config.SectionToUpdate == "/scheduleSettings" {
		jsonPatchObject := JsonPatchSchedule{
			ScheduleSettingValue: config.UpdatedScheduleSettingsSection,
			Path:                 path,
			Op:                   "replace",
		}
		jsonPatchDoc, _ = json.Marshal(jsonPatchObject)
	}
	if config.SectionToUpdate == "/alertGroup" {
		jsonPatchObject := JsonPatchAlert{
			AlertSettingValue: config.UpdatedAlertSettingsSection,
			Path:              path,
			Op:                "replace",
		}
		jsonPatchDoc, _ = json.Marshal(jsonPatchObject)
	}
	return string(jsonPatchDoc)
}

func updateTest(apiToken string, testId string, jsonPayload string) (string, string, bool, error) {

	type Data struct {
		Id int `json:"id"`
	}
	type ApiError struct {
		Id      json.Number `json:"id"`
		Message string      `json:"message"`
	}
	type Response struct {
		ResponseData Data       `json:"data"`
		Messages     []string   `json:"messages"`
		Errors       []ApiError `json:"errors"`
		Completed    bool       `json:"completed"`
		TraceId      string     `json:"traceId"`
	}

	updateURL := catchpointTestURI + "/" + testId
	var jsonPatchDocument = []byte(jsonPayload)
	var response Response
	var responseBody = ""
	var responseStatus = ""
	var completed = false
	req, _ := http.NewRequest("PATCH", updateURL, bytes.NewBuffer(jsonPatchDocument))
	req.Header.Set("Authorization", "Bearer "+apiToken)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return responseBody, responseStatus, completed, err
	}
	defer resp.Body.Close()

	responseStatus = strings.ToLower(string(resp.Status))
	body, _ := ioutil.ReadAll(resp.Body)

	json.Unmarshal([]byte(body), &response)
	completed = response.Completed

	return string(body), responseStatus, completed, nil
}
