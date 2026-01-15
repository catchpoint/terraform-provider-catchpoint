package alerttypes

import (
	"slices"

	"catchpoint-provider/internal/types"
)

type AlertSubTypes struct {
	SubTypes []string
}

type MonitorAlertTypes struct {
	AlertTypes map[string]*AlertSubTypes
}

type MonitorAlertTypeCompatibilityMatrix struct {
	TestType     types.TestType
	MonitorTypes map[string]*MonitorAlertTypes
}

// For a given schema, we must return *all* valid alert types to present the schema. Even if they won't work
// for the test+monitor combination. That gets validated later.
func (m *MonitorAlertTypeCompatibilityMatrix) ListAllValidAlertTypes() []string {
	unique := make(map[string]struct{})
	for _, monitor := range m.MonitorTypes {
		for alertType := range monitor.AlertTypes {
			unique[alertType] = struct{}{}
		}
	}
	result := make([]string, 0, len(unique))
	for alertType := range unique {
		result = append(result, alertType)
	}
	return result
}

// For a given schema, we must return *all* valid alert subtypes to present the schema. Even if they won't work
// for the test+monitor+alertType combination. That gets validated later.
func (m *MonitorAlertTypeCompatibilityMatrix) ListAllValidAlertSubTypes() []string {
	unique := make(map[string]struct{})
	for _, monitor := range m.MonitorTypes {
		for _, alertSubTypes := range monitor.AlertTypes {
			for _, subType := range alertSubTypes.SubTypes {
				unique[subType] = struct{}{}
			}
		}
	}
	result := make([]string, 0, len(unique))
	for subType := range unique {
		result = append(result, subType)
	}
	return result
}

func (m *MonitorAlertTypes) ListAllValidAlertTypes() []string {
	if m == nil {
		panic("MonitorAlertTypes is nil - this likely indicates that the test+monitor combination is invalid.")
	}

	keys := make([]string, 0, len(m.AlertTypes))
	for k := range m.AlertTypes {
		keys = append(keys, k)
	}
	return keys
}

func (m *MonitorAlertTypes) IsAlertTypeValid(alertType string) bool {
	alertTypes := m.ListAllValidAlertTypes()
	return slices.Contains(alertTypes, alertType)
}

func (m *AlertSubTypes) IsAlertSubTypeValid(alertSubType string) bool {
	subTypes := m.SubTypes
	return slices.Contains(subTypes, alertSubType)
}
