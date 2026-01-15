package transform

import (
	"testing"

	"catchpoint-provider/internal/testutil"

	"github.com/hashicorp/terraform-plugin-framework/types"
)

func TestTransformOptionalIntListLargeNumbers(t *testing.T) {
	intSlice := []int{999999, 1000000, 2147483647}

	list, diags := IntSliceToTerraformList(intSlice)

	testutil.AssertDiagsHasNoErrors(t, diags)
	testutil.AssertEqual(t, ErrorListNull, list.IsNull(), false)
	testutil.AssertEqual(t, "list length", len(list.Elements()), 3)
}

func TestTransformOptionalIntListNil(t *testing.T) {
	list, diags := IntSliceToTerraformList(nil)

	testutil.AssertDiagsHasNoErrors(t, diags)
	testutil.AssertEqual(t, ErrorListNull, list.IsNull(), true)
}

func TestTransformOptionalIntListEmpty(t *testing.T) {
	list, diags := IntSliceToTerraformList([]int{})

	testutil.AssertDiagsHasNoErrors(t, diags)
	testutil.AssertEqual(t, ErrorListNull, list.IsNull(), true)
}

func TestTransformOptionalIntListValid(t *testing.T) {
	intSlice := []int{1, 2, 3, 4, 5}

	list, diags := IntSliceToTerraformList(intSlice)

	testutil.AssertDiagsHasNoErrors(t, diags)
	testutil.AssertEqual(t, ErrorListNull, list.IsNull(), false)
	testutil.AssertEqual(t, "list length", len(list.Elements()), 5)

	// Check individual elements
	elements := list.Elements()
	for i, elem := range elements {
		int64Val, ok := elem.(types.Int64)
		if !ok {
			t.Fatalf("Expected element %d to be Int64 type", i)
		}
		testutil.AssertEqual(t, "element value", int64Val.ValueInt64(), int64(intSlice[i]))
	}
}
