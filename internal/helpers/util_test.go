package helpers

import (
	"catchpoint-provider/internal/testutil"
	"strings"
	"testing"
)

func TestGetTime(t *testing.T) {
	timeStr := GetTime()
	if len(timeStr) == 0 {
		t.Error("GetTime should return a non-empty string")
	}
}

func TestNormalizeScript(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"  Hello World  ", "Hello World\n"},
		{"\nHello\nWorld\n", "Hello\nWorld\n"},
		{"\r\nHello\r\nWorld\r\n", "Hello\nWorld\n"},
		{"// Step - 1\nopen(\"https://www.google.com\")\n", "// Step - 1\nopen(\"https://www.google.com\")\n"},
		{"", "\n"},
	}

	for _, test := range tests {
		result := NormalizeScript(test.input)
		testutil.AssertEqual(t, "NormalizeScript("+test.input+")", result, test.expected)
	}
}

func TestRandomHexString(t *testing.T) {
	hex := RandomHexString()
	if !strings.HasPrefix(hex, "#") || len(hex) != 7 {
		t.Errorf("RandomHexString returned invalid value: %s", hex)
	}
}

func TestIsValidEmail(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"test@testdomain.com", true},
		{"user.name+tag@domain.co.uk", true},
		{"foobar", false},
		{"test@", false},
		{"test@foo", false},
		{"userfoo.com", false},
		{"", false},
	}

	for _, test := range tests {
		result := IsValidEmail(test.input)
		testutil.AssertEqual(t, "IsValidEmail("+test.input+")", result, test.expected)
	}
}

func TestFlattenToIDs(t *testing.T) {
	type testStruct struct{ ID int }
	input := []testStruct{{ID: 1}, {ID: 2}}
	ids := FlattenToIDs(input, func(x testStruct) int { return x.ID })
	if len(ids) != 2 || ids[0] != 1 || ids[1] != 2 {
		t.Errorf("FlattenToIDs failed: %v", ids)
	}
}
