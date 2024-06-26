package romantointeger

import (
	"testing"
)

func TestRomanToInt(t *testing.T) {
	testCases := []struct {
		roman    string
		expected int
	}{
		{"III", 3},
		{"LVIII", 58},
		{"MCMXCIV", 1994},
		{"IV", 4},
	}

	for _, testCase := range testCases {
		actual := RomanToInt(testCase.roman)
		if actual != testCase.expected {
			t.Errorf("Expected %d for roman numeral %s, got %d", testCase.expected, testCase.roman, actual)
		}
	}
}
