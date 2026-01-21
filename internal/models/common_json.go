package models

import (
	"encoding/json"
	"errors"
	"log"
	"regexp"
	"strings"
)

// This is a generic and re-usable struct for all JSON objects that only have an ID and Name field.
type GenericIDNameJSON struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// This is a generic and re-usable struct for all JSON objects that only have an ID and Name field
// where the ID or Name may be omitted.
type GenericIDNameOmitEmptyJSON struct {
	ID   *int    `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

// The AlertWebhookJSON struct is used to represent a webhook in the alerting system.
type AlertWebhookJSON struct {
	ID *int `json:"id,omitempty"`
}

// The RecipientJSON is the JSON representation of the Recipient field.
type RecipientJSON struct {
	Email         string            `json:"email"`
	ID            *int              `json:"id,omitempty"`
	Name          string            `json:"name"`
	RecipientType GenericIDNameJSON `json:"recipientType"`
}

// The NodeJSON is the JSON representation of a Node. See also NodeGroupJSON.
type NodeJSON struct {
	ID          *int              `json:"id,omitempty"`
	Name        string            `json:"name"`
	NetworkType GenericIDNameJSON `json:"networkType"`
}

// The NodeGroupJSON is the JSON representation of a group of Nodes. See also NodeJSON.
type NodeGroupJSON struct {
	ID          *int   `json:"id,omitempty"`
	Name        string `json:"name"`
	NodeGroupID *int   `json:"nodeGroupId,omitempty"`
}

// A generic JSON object with only a required ID field. Typically used by Create payloads
// where we only need the ID of what was created in the response.
type DataJSON struct {
	ID json.Number `json:"id"`
}

// A JSON representation of an Authentication object which is part of the RequestSettings.
type AuthenticationJSON struct {
	AuthenticationMethodType *GenericIDNameOmitEmptyJSON `json:"authenticationMethodType,omitempty"`
	ID                       *int                        `json:"id,omitempty"`
	PasswordIDs              *[]int                      `json:"passwordIds,omitempty"`
}

// A JSON representation of a Notification group which can both be assigned directly to an AlertGroup as
// well as nested within the AlertGroupItems.
type NotificationGroupJSON struct {
	AlertWebhooks    []AlertWebhookJSON `json:"alertWebhooks"`
	NotifyOnCritical bool               `json:"notifyOnCritical"`
	NotifyOnImproved bool               `json:"notifyOnImproved"`
	NotifyOnWarning  bool               `json:"notifyOnWarning"`
	Recipients       []RecipientJSON    `json:"recipients"`
	Subject          string             `json:"subject"`
}

// The NodeThresholdJSON struct is used to represent a threshold for alert settings.
type NodeThresholdJSON struct {
	ConsecutiveRunsEnabled          bool              `json:"consecutiveRunsEnabled"`
	ID                              int               `json:"id"`
	Name                            string            `json:"name"`
	NodeThresholdType               GenericIDNameJSON `json:"nodeThresholdType"`
	NumberOfConsecutiveRuns         *int              `json:"consecutiveRuns,omitempty"`
	NumberOfFailingUnits            *int              `json:"numberOfFailingUnits,omitempty"`
	NumberOfUnits                   *int              `json:"numberOfUnits,omitempty"`
	PercentageOfUnits               *float64          `json:"percentageOfUnits,omitempty"`
	UtilizePerNodeHistoricalAverage bool              `json:"utilizePerNodeHistoricalAverage"`
}

// The TriggerJSON struct represents an alert trigger. Depending on the AlertType and AlertSubType
// the alert may be triggered by multiple things.
type TriggerJSON struct {
	CriticalReminderFrequency GenericIDNameJSON           `json:"criticalReminderFrequency"`
	CriticalTrigger           *float64                    `json:"criticalTrigger,omitempty"`
	DNSRecordType             *GenericIDNameOmitEmptyJSON `json:"dnsRecordType,omitempty"`
	DNSResolvedName           *string                     `json:"dnsResolvedName,omitempty"`
	DNSTTL                    *int                        `json:"dnsTTL,omitempty"`
	Expression                *string                     `json:"expression,omitempty"`
	FilterType                *GenericIDNameOmitEmptyJSON `json:"filterType,omitempty"`
	FilterValue               *string                     `json:"filterValue,omitempty"`
	HistoricalInterval        *GenericIDNameOmitEmptyJSON `json:"historicalInterval,omitempty"`
	ID                        int                         `json:"id"`
	Monitor                   *string                     `json:"monitor,omitempty"`
	OperationType             GenericIDNameJSON           `json:"operationType"`
	StatisticalType           *GenericIDNameOmitEmptyJSON `json:"statisticalType,omitempty"`
	ThresholdInterval         GenericIDNameJSON           `json:"thresholdInterval"`
	TriggerType               GenericIDNameJSON           `json:"triggerType"`
	UseIntervalRollingWindow  bool                        `json:"useIntervalRollingWindow"`
	WarningReminderFrequency  GenericIDNameJSON           `json:"warningReminderFrequency"`
	WarningTrigger            *float64                    `json:"warningTrigger,omitempty"`
}

// An AlertGroupItemJSON is the JSON representation of an AlertGroupItem which is a collection
// within an AlertGroup. Each one of these items represents an alert condition.
type AlertGroupItemJSON struct {
	AlertSubType       *GenericIDNameOmitEmptyJSON `json:"alertSubType,omitempty"`
	AlertType          GenericIDNameJSON           `json:"alertType"`
	EnforceTestFailure bool                        `json:"enforceTestFailure"`
	MatchAllRecords    bool                        `json:"matchAllRecords"`
	NodeThreshold      NodeThresholdJSON           `json:"nodeThreshold"`
	NotificationGroups []NotificationGroupJSON     `json:"notificationGroups"`
	NotificationType   GenericIDNameJSON           `json:"notificationType"`
	OmitScatterplot    bool                        `json:"omitScatterplot"`
	Trigger            TriggerJSON                 `json:"trigger"`
}

// AlertGroupJSON is the JSON representation of an AlertGroup which contains AlertGroupItems, a
// top-level NotificationGroup (for all alerts), and the AlertSettingsType (inherit, override, inherit & add).
type AlertGroupJSON struct {
	AlertGroupItems   []AlertGroupItemJSON   `json:"alertGroupItems,omitempty"`
	AlertSettingType  GenericIDNameJSON      `json:"alertSettingType"`
	NotificationGroup *NotificationGroupJSON `json:"notificationGroup,omitempty"`
}

// InsightDataJSON represents the JSON object for "insights" and "insightsData" fields.
type InsightDataJSON struct {
	Indicators         []GenericIDNameJSON `json:"indicators"`
	InsightSettingType GenericIDNameJSON   `json:"insightSettingType"`
	Tracepoints        []GenericIDNameJSON `json:"tracepoints"`
}

// HTTPHeaderRequestJSON represents the JSON object for HTTP header requests as part of the AdvancedSettings
// for applicable test types, folders and products.
//
// Note: RequestValue should exist but may be an empty string for some header types.
type HTTPHeaderRequestJSON struct {
	ChildHostPattern  *string           `json:"childHostPattern,omitempty"`
	HeaderName        *string           `json:"headerName,omitempty"`
	RequestHeaderType GenericIDNameJSON `json:"requestHeaderType"`
	RequestValue      string            `json:"requestValue"`
}

// The JSON representation of a requestSettings object for folders, products, and applicable test types.
//
// Note: even for the tests that do accept a requestSettings object, not all of them accept LibraryCertificateIDs
// or HTTPHeaders.
type RequestSettingsJSON struct {
	Authentication        *AuthenticationJSON      `json:"authentication,omitempty"`
	HTTPHeaderRequests    *[]HTTPHeaderRequestJSON `json:"httpHeaderRequests,omitempty"`
	LibraryCertificateIDs *[]int                   `json:"libraryCertificateIds,omitempty"`
	RequestSettingType    GenericIDNameJSON        `json:"requestSettingType"`
	TokenIDs              *[]int                   `json:"tokenIds,omitempty"`
}

// The JSON representation of a scheduleSetting or scheduleSettings object for folders, products, and applicable test types.
// Not all tests may override the schedule for some reason (e.g. BGP).
type ScheduleSettingsJSON struct {
	Frequency             GenericIDNameJSON `json:"frequency"`
	ID                    int               `json:"id"`
	MaintenanceScheduleID *int              `json:"maintenanceScheduleId,omitempty"`
	NetworkType           GenericIDNameJSON `json:"networkType"`
	NodeGroups            []NodeGroupJSON   `json:"nodeGroups"`
	Nodes                 []NodeJSON        `json:"nodes"`
	NoOfSubsetNodes       *int              `json:"roundRobinAmount,omitempty"`
	RunScheduleID         *int              `json:"runScheduleId,omitempty"`
	ScheduleSettingType   GenericIDNameJSON `json:"scheduleSettingType"`
	TestNodeDistribution  GenericIDNameJSON `json:"testNodeDistribution"`
}

// The JSON representation of the advancedSettings or advancedSettingsModel object for folders, products, and applicable tests.
//
// Note: the flags are a list of ID+Name pairs where only some flags are applicable to a given test type.
type AdvancedSettingsJSON struct {
	AdditionalMonitor         *GenericIDNameOmitEmptyJSON  `json:"additionalMonitor,omitempty"`
	AdvancedSettingType       GenericIDNameJSON            `json:"advancedSettingType"`
	AppliedTestFlags          []GenericIDNameOmitEmptyJSON `json:"appliedTestFlags"`
	EDNSSubnet                *string                      `json:"ednsSubnet,omitempty"`
	FailureHopCount           *int                         `json:"failureHopCount,omitempty"`
	ID                        int                          `json:"id"`
	MaxStepRuntimeSecOverride *int                         `json:"maxStepRuntimeSecOverride,omitempty"`
	PingCount                 *int                         `json:"pingCount,omitempty"`
	TestBandwidthThrottling   *GenericIDNameOmitEmptyJSON  `json:"testBandwidthThrottling,omitempty"`
	ViewportHeight            *int                         `json:"viewportHeight,omitempty"`
	ViewportWidth             *int                         `json:"viewportWidth,omitempty"`
	WaitForNoActivity         *int                         `json:"waitForNoActivity,omitempty"` // Enabled by Stop On Doc Complete
}

// The JSON representation of an Error from the API. The errors are used in a collection of errors, if any.
type ErrorJSON struct {
	ID      json.Number `json:"id"`
	Message string      `json:"message"`
}

// A JSON representation of a generic response. This is used for Create and Update calls where we don't
// expect a complex object, just an ID, completed bool, messages, traceID, and possibly errors.
type Response struct {
	ResponseData DataJSON `json:"data"`
	CommonResponse
}

// The JSON representation of a deleted resource. This is different than DataJSON because the ID is called 'deleted'.
type DeleteDataJSON struct {
	ID string `json:"deleted"`
}

// The common response structure for all API responses.
type CommonResponse struct {
	Completed bool        `json:"completed"`
	Errors    []ErrorJSON `json:"errors"`
	Messages  []string    `json:"messages"`
	TraceID   string      `json:"traceId"`
}

// An error indicating that a requested object was not found. Our API does not return 404 for missing objects so we
// need to derive from a regex search of the response messages that the object is missing. In such a case, we
// return this error.
type ObjectNotFoundError struct {
	Message string
}

// Error implements the error interface for ObjectNotFoundError.
func (e *ObjectNotFoundError) Error() string {
	return e.Message
}

// Returns an error message if Completed == false with the message containing a flattened representation
// of *all* the errors that occurred.
func (r *CommonResponse) ErrorIfIncomplete() error {
	if r.Completed {
		return nil
	}
	var errs []string
	notFoundRegex := regexp.MustCompile(`No valid .* found for input`)

	for _, e := range r.Errors {
		if e.Message != "" {
			// Since the API returns 400 for everything, but we want to know if an object we manage has been deleted,
			// check the error message for a specific pattern.
			if notFoundRegex.MatchString(e.Message) {
				log.Printf("[ERROR] Object not found: %v", e.Message)
				return &ObjectNotFoundError{Message: e.Message}
			}
			errs = append(errs, e.Message)
		}
	}
	log.Printf("[ERROR] API response incomplete, Errors: %v", errs)

	return errors.New("API response incomplete: " + strings.Join(errs, ", "))
}
