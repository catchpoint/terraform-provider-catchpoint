package testutil

import (
	"math/rand"
	"reflect"
	"strconv"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var rnd = rand.New(rand.NewSource(time.Now().UnixNano()))

// GenerateRandomEmail generates a random email address for testing purposes.
func GenerateRandomEmail() string {
	user := "user" + strconv.Itoa(rnd.Intn(100000))
	domain := "example" + strconv.Itoa(rnd.Intn(1000)) + ".com"
	return user + "@" + domain
}

// GenerateRandomIntN returns a random integer in the range [0, n).
// If n is less than or equal to 0, it returns 0.
func GenerateRandomIntN(n int) int {
	if n <= 0 {
		return 0
	}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.Intn(n)
}

// GetOrFail calls fn, checks ok, and returns the value or fails the test.
func GetOrFail[T any](t *testing.T, fn func() (T, bool), msg string, param any) T {
	val, ok := fn()
	if !ok {
		t.Fatalf("%s: lookup failed for %v", msg, param)
	}
	return val
}

// GetTypeOrAssert checks if the field exists in attrTypes and returns its Type or Asserts.
func GetTypeOrAssert(t *testing.T, field string, attrTypes map[string]attr.Type) attr.Type {
	if gotType, exists := attrTypes[field]; !exists {
		t.Errorf("Expected field %s not found in attribute types", field)
		return nil
	} else {
		return gotType
	}
}

// GetObjectOrAssert checks if the field exists in attrTypes and returns its ObjectType or Asserts.
func GetObjectOrAssert(t *testing.T, field string, attrTypes map[string]attr.Type) types.ObjectType {
	attrType, exists := attrTypes[field]
	if !exists {
		t.Errorf("Expected field %s not found in attribute types", field)
		return types.ObjectType{}
	}

	objType, ok := attrType.(types.ObjectType)
	if !ok {
		t.Errorf("Expected %s to be ObjectType, got %T", field, attrType)
		return types.ObjectType{}
	}

	return objType
}

// GetAttributeOrAssert checks if the field exists in attrs as an Attribute and returns it or asserts.
func GetAttributeOrAssert(t *testing.T, field string, attrs map[string]schema.Attribute) schema.Attribute {
	if gotField, ok := attrs[field]; !ok {
		t.Errorf("expected %s field to be present when includeCertificates is true", field)
	} else {
		return gotField
	}
	return nil
}

// GetBlockOrAssert checks if the field exists in attrs as a Block and returns it or asserts.
func GetBlockOrAssert(t *testing.T, field string, attrs map[string]schema.Block) schema.Block {
	if gotBlock, ok := attrs[field]; !ok {
		t.Errorf("expected %s block to be present", field)
	} else {
		return gotBlock
	}
	return nil
}

// GetSingleNestedAttributeOrAssert checks if the field exists in attrs as a SingleNestedAttribute and returns it or asserts.
func GetSingleNestedAttributeOrAssert(t *testing.T, field string, attrs map[string]schema.Attribute) schema.SingleNestedAttribute {
	attr, ok := attrs[field]
	if !ok {
		t.Fatalf("expected %s field to be present", field)
	}
	singleNested, ok := attr.(schema.SingleNestedAttribute)
	if !ok {
		t.Fatalf("expected %s to be SingleNestedAttribute", field)
	}
	return singleNested
}

// isReallyNil checks if the value is nil or a nil pointer.
func isReallyNil(i any) bool {
	if i == nil {
		return true
	}
	v := reflect.ValueOf(i)
	switch v.Kind() {
	case reflect.Chan, reflect.Func, reflect.Interface, reflect.Map, reflect.Ptr, reflect.Slice:
		return v.IsNil()
	}
	return false
}

// findMissing returns elements that are in 'want' but not in 'got'
// Both slices should be sorted
func findMissing(want, got []string) []string {
	var missing []string
	gotMap := make(map[string]int)

	// Count occurrences in 'got'
	for _, item := range got {
		gotMap[item]++
	}

	// Find missing items from 'want'
	wantMap := make(map[string]int)
	for _, item := range want {
		wantMap[item]++
	}

	for item, wantCount := range wantMap {
		gotCount := gotMap[item]
		if gotCount < wantCount {
			// Add the missing occurrences
			for i := 0; i < wantCount-gotCount; i++ {
				missing = append(missing, item)
			}
		}
	}

	return missing
}

// ToIntPtr converts an int to a pointer to int.
func ToIntPtr(v int) *int {
	return &v
}

// ToFloat64Ptr converts a float64 to a pointer to float64.
func ToFloat64Ptr(v float64) *float64 {
	return &v
}

// ToStringPtr converts a string to a pointer to string.
func ToStringPtr(v string) *string {
	return &v
}

// ToBoolPtr converts a bool to a pointer to bool.
func ToBoolPtr(v bool) *bool {
	return &v
}
