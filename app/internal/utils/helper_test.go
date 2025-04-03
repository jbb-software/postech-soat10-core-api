package utils

import (
	"database/sql"
	"testing"
)

func TestNullString(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected sql.NullString
	}{
		{
			name:     "Empty string",
			input:    "",
			expected: sql.NullString{},
		},
		{
			name:     "Non-empty string",
			input:    "test",
			expected: sql.NullString{String: "test", Valid: true},
		},
		{
			name:     "String with spaces",
			input:    "test with spaces",
			expected: sql.NullString{String: "test with spaces", Valid: true},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := NullString(tc.input)
			if result != tc.expected {
				t.Errorf("NullString(%q) = %+v, expected %+v", tc.input, result, tc.expected)
			}
		})
	}
}

func TestContains(t *testing.T) {
	testCases := []struct {
		name     string
		slice    []string
		value    string
		expected bool
	}{
		{
			name:     "Value exists",
			slice:    []string{"a", "b", "c"},
			value:    "b",
			expected: true,
		},
		{
			name:     "Value does not exist",
			slice:    []string{"a", "b", "c"},
			value:    "d",
			expected: false,
		},
		{
			name:     "Empty slice",
			slice:    []string{},
			value:    "a",
			expected: false,
		},
		{
			name:     "Slice with one element, element found",
			slice:    []string{"a"},
			value:    "a",
			expected: true,
		},
		{
			name:     "Slice with one element, element not found",
			slice:    []string{"a"},
			value:    "b",
			expected: false,
		},
		{
			name:     "Slice with duplicate values, element found",
			slice:    []string{"a", "b", "b", "c"},
			value:    "b",
			expected: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := Contains(tc.slice, tc.value)
			if result != tc.expected {
				t.Errorf("Contains(%v, %q) = %t, expected %t", tc.slice, tc.value, result, tc.expected)
			}
		})
	}
}
