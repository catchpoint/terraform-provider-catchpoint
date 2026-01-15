package testutil

import (
	"fmt"
	"math"
	"reflect"
	"slices"
	"sort"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

const (
	errorMismatch  = "%s mismatch: got %v, want %v"
	floatTolerance = 1e-9
)

// AssertDiagsHasNoErrors checks if the diagnostics contain any errors.
// If there are errors, it reports them in the test output.
func AssertDiagsHasNoErrors(t *testing.T, diags diag.Diagnostics) {
	if diags.HasError() {
		for _, d := range diags {
			t.Errorf("Unexpected error in diagnostics: %s", d.Detail())
		}
	}
}

// AssertListTypeAndLength asserts that a list is of the expected type and has the expected length.
func AssertListTypeAndLength(t *testing.T, list types.List, expectedLength int) {
	if list.IsNull() {
		t.Fatalf("Expected list to be not null, but it is null")
	}
	if len(list.Elements()) != expectedLength {
		t.Errorf("Expected list length %d, got %d", expectedLength, len(list.Elements()))
	}
}

// AssertStringTypeAndValue asserts that a field is of String type and has the expected value.
func AssertStringTypeAndValue(t *testing.T, field string, attrTypes map[string]attr.Value, expectedValue string) {
	stringVal, ok := attrTypes[field].(types.String)
	if !ok {
		t.Errorf("Expected %s to be StringType, got %T", field, attrTypes[field])
	}

	AssertEqual(t, field+" value", stringVal.ValueString(), expectedValue)
}

// AssertNullListTypeForAttribute asserts that a field is a List type and is null.
func AssertNullListTypeForAttribute(t *testing.T, field string, attrs map[string]attr.Value) {
	list, ok := attrs[field].(types.List)
	if !ok {
		t.Fatalf("Expected %s to be List type", field)
	}
	AssertEqual(t, fmt.Sprintf("%s list is null", field), list.IsNull(), true)
}

// AssertObjectTypeAndNotNull asserts that a field is of Object type and not null.
func AssertObjectTypeAndNotNull(t *testing.T, field string, attrs map[string]attr.Value) types.Object {
	obj, ok := attrs[field].(types.Object)
	if !ok {
		t.Fatalf("Expected %s to be Object type, got %T", field, attrs[field])
	}
	AssertEqual(t, fmt.Sprintf("%s is null", field), obj.IsNull(), false)
	return obj
}

// Assert equality of two comparable types.
// Param 2 is the field name for error reporting.
func AssertEqual[T comparable](t *testing.T, fieldName string, got, want T) {
	if got != want {
		t.Errorf(errorMismatch, fieldName, got, want)
	}
}

func AssertTrue(t *testing.T, fieldName string, condition bool) {
	if !condition {
		t.Errorf("Expected %s to be true but was false", fieldName)
	}
}

func AssertFalse(t *testing.T, fieldName string, condition bool) {
	if condition {
		t.Errorf("Expected %s to be false but was true", fieldName)
	}
}

// Assert that a field is greater than a specific value.
func AssertGreaterThan(t *testing.T, fieldName string, got, want int) {
	if got <= want {
		t.Errorf("%s should be greater than %v, but got %v", fieldName, want, got)
	}
}

// Assert that a field contains a specific value in a slice.
// Param 2 is the field name for error reporting.
func AssertContains[T comparable](t *testing.T, fieldName string, slice []T, expected T) {
	t.Helper()
	if slices.Contains(slice, expected) {
		return // Found it
	}
	t.Errorf("%s should contain %v, but it doesn't. Actual slice: %v", fieldName, expected, slice)
}

// Assert that a string contains a specific substring.
func AssertStringContains(t *testing.T, fieldName string, str, substr string) {
	if !strings.Contains(str, substr) {
		t.Errorf("should contain %q in %q: got %q", substr, fieldName, str)
	}
}

// Assert that two values are not equal.
// Param 2 is the field name for error reporting.
func AssertNotEqual[T comparable](t *testing.T, fieldName string, got, want T) {
	if got == want {
		t.Errorf(errorMismatch, fieldName, got, want)
	}
}

// Assert that a field is not nil.
// Param 2 is the field name for error reporting.
func AssertNotNil(t *testing.T, fieldName string, got any) {
	if isReallyNil(got) {
		t.Errorf("Expected %s to be not nil, but it is nil", fieldName)
	}
}

// Assert that a field is nil.
// Param 2 is the field name for error reporting.
func AssertNil(t *testing.T, fieldName string, got any) {
	if !isReallyNil(got) {
		t.Errorf("Expected %s to be nil, but it is not: %v", fieldName, got)
	}
}

// AssertElementsMatch asserts that two slices contain the same elements, regardless of order.
// It sorts both slices before comparing them and reports missing elements.
func AssertElementsMatch(t *testing.T, fieldName string, got, want []string) {
	gotCopy := append([]string(nil), got...)
	wantCopy := append([]string(nil), want...)

	// Sort both slices
	sort.Strings(gotCopy)
	sort.Strings(wantCopy)

	if !reflect.DeepEqual(gotCopy, wantCopy) {
		// Find elements in 'want' but not in 'got'
		missing := findMissing(wantCopy, gotCopy)

		// Find elements in 'got' but not in 'want'
		extra := findMissing(gotCopy, wantCopy)

		errorMsg := fmt.Sprintf("%s mismatch:\n  got: %v\n  want: %v", fieldName, got, want)

		if len(missing) > 0 {
			errorMsg += fmt.Sprintf("\n  missing elements: %v", missing)
		}

		if len(extra) > 0 {
			errorMsg += fmt.Sprintf("\n  extra elements: %v", extra)
		}

		t.Error(errorMsg)
	}
}

// Assert equality of two comparable types using reflect.DeepEqual.
// Param 2 is the field name for error reporting.
func AssertDeepEqual[T any](t *testing.T, fieldName string, got, want []T) {
	if !reflect.DeepEqual(got, want) {
		t.Errorf(errorMismatch, fieldName, got, want)
	}
}

// Assert that a Type DOES NOT exist in the attributes map.
func AssertTypeNotExists(t *testing.T, field string, attrs map[string]attr.Type) {
	if _, exists := attrs[field]; exists {
		t.Errorf("Expected field %s to not exist in attributes", field)
	}
}

// Assert that an Attribute DOES NOT exist in the attributes map.
func AssertAttribNotExists(t *testing.T, field string, attrs map[string]attr.Value) {
	if _, exists := attrs[field]; exists {
		t.Errorf("Expected field %s to not exist in attributes", field)
	}
}

// Assert that a field is a List type of Object and not null.
func AssertListTypeOfType(t *testing.T, field string, attrTypes map[string]attr.Type, expectedType attr.Type) {
	if listType, ok := attrTypes[field].(types.ListType); ok {
		AssertEqual(t, field+" element type", listType.ElemType, expectedType)
	} else {
		t.Errorf("Expected %s to be ListType, got %T", field, attrTypes[field])
	}
}

// AssertFloat64Equal compares two types.Float64 values with floating-point tolerance
func AssertFloat64Equal(t *testing.T, name string, got, want types.Float64) {
	t.Helper()

	// Handle null/unknown cases
	if got.IsNull() != want.IsNull() {
		t.Errorf("%s null state mismatch: got IsNull()=%v, want IsNull()=%v", name, got.IsNull(), want.IsNull())
		return
	}
	if got.IsUnknown() != want.IsUnknown() {
		t.Errorf("%s unknown state mismatch: got IsUnknown()=%v, want IsUnknown()=%v", name, got.IsUnknown(), want.IsUnknown())
		return
	}

	// If both are null or unknown, they're equal
	if got.IsNull() || got.IsUnknown() {
		return
	}

	// Compare values with tolerance
	gotVal := got.ValueFloat64()
	wantVal := want.ValueFloat64()

	if math.Abs(gotVal-wantVal) > floatTolerance {
		t.Errorf("%s mismatch: got %f, want %f", name, gotVal, wantVal)
	}
}
