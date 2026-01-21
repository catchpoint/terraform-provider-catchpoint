package alerttypes

import (
	"slices"
	"testing"

	"catchpoint-provider/internal/testutil"
)

var testAlertMatrix = &MonitorAlertTypeCompatibilityMatrix{
	TestType: fakeTestType,
	MonitorTypes: map[string]*MonitorAlertTypes{
		"someMonitorType": {
			AlertTypes: map[string]*AlertSubTypes{
				"anAlertType":      {SubTypes: []string{"alertSubtype", "anotherAlertSubtype"}},
				"anotherAlertType": {SubTypes: []string{"aThirdAlertSubtype"}},
				"anEmptyType":      {},
			},
		},
		"anotherMonitorType": {
			AlertTypes: map[string]*AlertSubTypes{
				"anAlertType":        {SubTypes: []string{"alertSubtype", "anotherAlertSubtype"}},
				"anotherAlertType":   {SubTypes: []string{"aThirdAlertSubtype"}},
				"aTypeNotInFirstSet": {SubTypes: []string{"subTypeNotInFirstSet"}},
			},
		},
	},
}

func TestMonitorAlertTypeCompatibilityMatrixListAllValidAlertTypes(t *testing.T) {
	types := testAlertMatrix.ListAllValidAlertTypes()

	testutil.AssertElementsMatch(t, "AlertTypes_anAlertType", types, []string{"anAlertType", "anotherAlertType", "anEmptyType", "aTypeNotInFirstSet"})
}

func TestMonitorAlertTypeCompatibilityMatrixListAllValidAlertSubTypes(t *testing.T) {
	subTypes := testAlertMatrix.ListAllValidAlertSubTypes()

	testutil.AssertElementsMatch(t, "AlertSubTypes_anAlertType", subTypes, []string{"alertSubtype", "anotherAlertSubtype", "aThirdAlertSubtype", "subTypeNotInFirstSet"})
}

func TestMonitorAlertTypesListAllValidAlertTypes(t *testing.T) {
	types := testAlertMatrix.MonitorTypes["someMonitorType"].ListAllValidAlertTypes()

	testutil.AssertEqual(t, "AlertTypes_anAlertType", slices.Contains(types, "anAlertType"), true)
	testutil.AssertEqual(t, "AlertTypes_anEmptyType", slices.Contains(types, "anEmptyType"), true)
	testutil.AssertEqual(t, "AlertTypes_anotherAlertType", slices.Contains(types, "anotherAlertType"), true)
}

func TestMonitorAlertTypesIsAlertTypeValid(t *testing.T) {
	mat := testAlertMatrix.MonitorTypes["someMonitorType"]

	testutil.AssertEqual(t, "IsAlertTypeValid_AnAlertType", mat.IsAlertTypeValid("anAlertType"), true)
	testutil.AssertEqual(t, "IsAlertTypeValid_AnotherAlertType", mat.IsAlertTypeValid("anotherAlertType"), true)

	// Valid, even though it has no subtypes.
	testutil.AssertEqual(t, "IsAlertTypeValid_AnotherAlertType", mat.IsAlertTypeValid("anEmptyType"), true)

	// Not valid, does not exist in the map.
	testutil.AssertEqual(t, "IsAlertTypeValid_NonExistentType", mat.IsAlertTypeValid("nonExistentType"), false)
}

func TestAlertSubTypesIsAlertSubTypeValid(t *testing.T) {
	sub := testAlertMatrix.MonitorTypes["someMonitorType"].AlertTypes["anAlertType"]

	testutil.AssertEqual(t, "IsAlertSubTypeValid_alertSubtype", sub.IsAlertSubTypeValid("alertSubtype"), true)
	testutil.AssertEqual(t, "IsAlertSubTypeValid_nonExistentSubtype", sub.IsAlertSubTypeValid("nonExistentSubtype"), false)
}

func TestAlertSubTypesIsAlertSubTypeNotValidForEmpty(t *testing.T) {
	// The empty type should not have any subtypes.
	sub := testAlertMatrix.MonitorTypes["someMonitorType"].AlertTypes["anEmptyType"]

	// There's nothing valid but we also shouldn't error.
	testutil.AssertEqual(t, "IsAlertSubTypeValid_alertSubtype", sub.IsAlertSubTypeValid("alertSubtype"), false)
}
