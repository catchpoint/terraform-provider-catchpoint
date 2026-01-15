package helpers

import (
	"math/rand"
	"regexp"
	"strings"
	"time"
)

// Allow mocking by setting GetTimeFunc to a different function in tests.
var GetTimeFunc = GetTime

// GetTime returns the current time in RFC3339 format.
func GetTime() string {
	t := time.Now()
	timeCurrent := t.Format(time.RFC3339)
	return string(timeCurrent)
}

// RandomHexString generates a random hex string of 6 characters prefixed with '#'.
func RandomHexString() string {
	const hexLetters = "abcdef0123456789"
	const numOfLetters = 6
	b := make([]byte, numOfLetters)
	for i := range b {
		b[i] = hexLetters[rand.Intn(len(hexLetters))]
	}
	return "#" + string(b)
}

// IsValidEmail checks if the provided string is a valid email address.
func IsValidEmail(email string) bool {
	// Regular expression pattern for validating email addresses
	pattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`

	// Compile the pattern
	regex := regexp.MustCompile(pattern)

	// Use the MatchString method to check if the email matches the pattern
	return regex.MatchString(email)
}

// Flatten maps for simple comparison.
// Example:
//
//	ids := helpers.FlattenToIDs(config.SomeArrayOfObjectsWithIDs, func(x models.GenericIDName) int { return x.ID })
func FlattenToIDs[T any](input []T, getID func(T) int) []int {
	ids := make([]int, len(input))
	for i, v := range input {
		ids[i] = getID(v)
	}
	return ids
}

// NormalizeScript removes leading/trailing whitespace and normalizes line endings in a script.
func NormalizeScript(s string) string {
	// Remove leading/trailing whitespace and normalize line endings.
	s = strings.ReplaceAll(s, "\r\n", "\n")
	s = strings.TrimSpace(s)
	// Ensure exactly one trailing newline to satisfy TF formatting.
	if !strings.HasSuffix(s, "\n") {
		s += "\n"
	}
	return s
}

// FlatCombineStringSlices combines multiple string slices into one.
func FlatCombineStringSlices(slices ...[]string) []string {
	var result []string
	for _, s := range slices {
		result = append(result, s...)
	}
	return result
}

// FlatCombineStringSlicesUnique combines multiple string slices into one, ensuring uniqueness.
func FlatCombineStringSlicesUnique(slices ...[]string) []string {
	set := make(map[string]struct{})
	for _, slice := range slices {
		for _, key := range slice {
			set[key] = struct{}{}
		}
	}
	result := make([]string, 0, len(set))
	for key := range set {
		result = append(result, key)
	}
	return result
}
