package stringcomparison_test

import (
	"testing"

	"github.com/TechMDW/go-utils/stringcomparison"
)

func TestCompareStrings(t *testing.T) {
	tests := []struct {
		str1, str2     string
		expectedResult stringcomparison.CompareStringsResult
	}{
		{"hello", "hello", stringcomparison.CompareStringsResult{5, 1, 5, 1}},
		{"hello", "world", stringcomparison.CompareStringsResult{1, 0.2, 1, 0.2}},
		{"hello", "hEllo", stringcomparison.CompareStringsResult{4, 0.8, 5, 1}},
		{"world", "W0rLd", stringcomparison.CompareStringsResult{2, 0.4, 4, 0.8}},
		{"", "", stringcomparison.CompareStringsResult{0, 0, 0, 0}},
		{" ", "", stringcomparison.CompareStringsResult{0, 0, 0, 0}},
		{" ", " ", stringcomparison.CompareStringsResult{1, 1, 1, 1}},
	}

	for _, test := range tests {
		result := stringcomparison.CompareStrings(test.str1, test.str2)
		if result != test.expectedResult {
			t.Errorf("For %s, %s expected %v but got %v", test.str1, test.str2, test.expectedResult, result)
		}
	}
}

func BenchmarkCompareStrings(b *testing.B) {
	for i := 0; i < b.N; i++ {
		stringcomparison.CompareStrings("hello", "hEllo")
	}
}
