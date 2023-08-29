package stringcomparison

import (
	"unicode"

	"github.com/TechMDW/go-utils/division"
)

type CompareStringsResult struct {
	SameCharCount       int
	MatchPercent        float64
	SameCharCountNoCase int
	MatchPercentNoCase  float64
}

type CompareStringsCharsResult struct {
	SameChars       []rune
	SameCharsNoCase []rune
}

// CompareStrings takes two strings as input and returns a ComparisonResult struct
// containing various metrics that quantify their similarity.
//
// Metrics included are:
//
// - SameCharCount: The number of characters that are exactly the same at the
// corresponding positions in both strings.
//
// - MatchPercent: 0 to 1. The percentage of characters that are exactly the same,
// calculated based on the shorter string's length.
//
// - SameCharCountNoCase: The number of characters that are the same when case is ignored,
// at the corresponding positions in both strings.
//
// - MatchPercentNoCase: 0 to 1. The percentage of characters that are the same when case is ignored,
// calculated based on the shorter string's length.
//
// The function uses rune slices for Unicode support. It uses bitwise OR for fast ASCII character comparison,
// and falls back to unicode.SimpleFold for full Unicode case-insensitive comparison.
func CompareStrings(str1, str2 string) CompareStringsResult {
	var sameCharCount, sameCharCountNoCase int
	str1Runes := []rune(str1)
	str2Runes := []rune(str2)

	minLen := len(str1Runes)
	if len(str2Runes) < minLen {
		minLen = len(str2Runes)
	}

	minLenFloat := float64(minLen)

	for i := 0; i < minLen; i++ {
		str1Rune := str1Runes[i]
		str2Rune := str2Runes[i]
		if str1Rune == str2Rune {
			sameCharCount++
		}

		// If both runes are ASCII, we can do a simple case-insensitive comparison.
		// This is much faster than unicode.SimpleFold. About 1.8x faster during benchmarks.
		if str1Rune < 128 && str2Rune < 128 {
			if str1Rune|32 == str2Rune|32 {
				sameCharCountNoCase++
			}
		} else if unicode.SimpleFold(str1Rune) == unicode.SimpleFold(str2Rune) {
			sameCharCountNoCase++
		}
	}

	return CompareStringsResult{
		SameCharCount:       sameCharCount,
		MatchPercent:        division.DivideFloat64(float64(sameCharCount), minLenFloat),
		SameCharCountNoCase: sameCharCountNoCase,
		MatchPercentNoCase:  division.DivideFloat64(float64(sameCharCountNoCase), minLenFloat),
	}
}
