package types

import (
	"testing"

	"catchpoint-provider/internal/testutil"
)

func TestGetStatusTypeIDAndName(t *testing.T) {
	tests := []struct {
		status   string
		statusID int
	}{
		{Active, 0},
		{Inactive, 1},
	}
	for _, tt := range tests {
		id := testutil.GetOrFail(t, func() (int, bool) { return GetStatusTypeID(tt.status) }, "GetStatusTypeID", tt.status)
		name := testutil.GetOrFail(t, func() (string, bool) { return GetStatusTypeName(tt.statusID) }, "GetStatusTypeName", tt.statusID)

		testutil.AssertEqual(t, "statusID", id, tt.statusID)
		testutil.AssertEqual(t, "statusName", name, tt.status)
	}
}

func TestGetGenericSettingTypeIDAndName(t *testing.T) {
	test := []struct {
		name string
		id   int
	}{
		{Inherit, 0},
		{Override, 1},
	}
	for _, tt := range test {
		id := testutil.GetOrFail(t, func() (int, bool) { return GetGenericSettingTypeID(tt.name) }, "GetGenericSettingTypeID", tt.name)
		name := testutil.GetOrFail(t, func() (string, bool) { return GetGenericSettingTypeName(id) }, "GetGenericSettingTypeName", id)

		testutil.AssertEqual(t, "genericSettingTypeID", id, tt.id)
		testutil.AssertEqual(t, "genericSettingTypeName", name, tt.name)
	}
}
