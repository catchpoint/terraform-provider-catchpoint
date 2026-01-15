package types

import "slices"

var (
	ValidStatusTypes             []string
	ValidGenericSettingTypeNames []string
	ValidInsightSettingTypeNames []string
	statusTypeNames              map[int]string
	genericSettingTypeNames      map[int]string
	insightStatusTypeNames       map[int]string
)

// #region StatusTypes

var StatusTypeIDs = map[string]int{
	Active:   0,
	Inactive: 1,
}

func GetStatusTypeID(name string) (int, bool) {
	id, ok := StatusTypeIDs[name]
	return id, ok
}

func GetStatusTypeName(id int) (string, bool) {
	name, ok := statusTypeNames[id]
	return name, ok
}

// #endregion
// #region SettingTypes

// These are generic setting types used in various places, such as schedulesSettingType and requestSettingType.
// Note: AlertSettingType has its own because it also has the 'Include & Add' option.
var GenericSettingTypeIDs = map[string]int{
	Inherit:  0,
	Override: 1,
}

func GetGenericSettingTypeID(genericSettingType string) (int, bool) {
	id, ok := GenericSettingTypeIDs[genericSettingType]
	return id, ok
}

func GetGenericSettingTypeName(id int) (string, bool) {
	name, ok := genericSettingTypeNames[id]
	return name, ok
}

var InsightStatusTypeIDs = map[string]int{
	Inherit:    0,
	Override:   1,
	NoSettings: 3,
}

func GetInsightSettingTypeID(name string) (int, bool) {
	id, ok := InsightStatusTypeIDs[name]
	return id, ok
}

func GetInsightSettingTypeName(id int) (string, bool) {
	name, ok := insightStatusTypeNames[id]
	return name, ok
}

// #endregion
// #region Helper functions

func populateReverseMap(src map[string]int) map[int]string {
	dst := make(map[int]string, len(src))
	for name, id := range src {
		dst[id] = name
	}
	return dst
}

func populateValidNames(reverseMap map[string]int) []string {
	validNames := make([]string, 0)
	for name := range reverseMap {
		validNames = append(validNames, name)
	}
	slices.Sort(validNames)

	return validNames
}

// #endregion

func init() {
	statusTypeNames = populateReverseMap(StatusTypeIDs)
	ValidStatusTypes = populateValidNames(StatusTypeIDs)

	genericSettingTypeNames = populateReverseMap(GenericSettingTypeIDs)
	ValidGenericSettingTypeNames = populateValidNames(GenericSettingTypeIDs)

	insightStatusTypeNames = populateReverseMap(InsightStatusTypeIDs)
	ValidInsightSettingTypeNames = populateValidNames(InsightStatusTypeIDs)
}
