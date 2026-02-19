package utils

import (
	"regexp"
	"strings"
	"unicode"
)

// Regex patterns to match "Stick" characters adjacent to Cyrillic characters.
// We use Unicode codepoints to be 100% precise about which characters we target.
//
// The "Stick" character set includes:
// 1.  Latin:   I (\u0049), i (\u0069), l (\u006C - Small L, common typo)
// 2.  Turkish: ı (\u0131 - Dotless i), İ (\u0130 - Dotted I)
// 3.  Cyrillic: Ӏ (\u04C0 - Palochka Upper), ӏ (\u04CF - Palochka Lower)
// 4.  Cyrillic I: І (\u0406 - Ukr/Bel Upper), і (\u0456 - Ukr/Bel Lower)
//
// Pattern 1: Cyrillic Letter -> Stick (e.g., "кi" -> "к1")
var cyrillicThenStick = regexp.MustCompile(`(\p{Cyrillic})([\x{0049}\x{0069}\x{006C}\x{0131}\x{0130}\x{04C0}\x{04CF}\x{0406}\x{0456}])`)

// Pattern 2: Stick -> Cyrillic Letter (e.g., "iэ" -> "1э")
var stickThenCyrillic = regexp.MustCompile(`([\x{0049}\x{0069}\x{006C}\x{0131}\x{0130}\x{04C0}\x{04CF}\x{0406}\x{0456}])(\p{Cyrillic})`)

func ConvertAllPolachkaLookingLettersTo1InCircassianWords(str string) string {
	// Pass 1: Handle cases like "кi" -> "к1"
	str = cyrillicThenStick.ReplaceAllString(str, "${1}1")

	// Pass 2: Handle cases like "iэ" -> "1э"
	str = stickThenCyrillic.ReplaceAllString(str, "1${2}")

	return str
}

// IsFullyCapitalized checks if all letter characters in the string are uppercase.
// Non-letter characters (digits, punctuation) are ignored.
// Returns false for empty strings or strings with no letters.
func IsFullyCapitalized(s string) bool {
	s = strings.TrimRight(s, ".,;:-!?\"'()[]")
	if s == "" {
		return false
	}
	hasLetter := false
	for _, r := range s {
		if unicode.IsLetter(r) {
			hasLetter = true
			if !unicode.IsUpper(r) {
				return false
			}
		}
	}
	return hasLetter
}

// StartsWithNumber checks if the string begins with a digit (0-9).
func StartsWithNumber(s string) bool {
	if s == "" {
		return false
	}
	return s[0] >= '0' && s[0] <= '9'
}

// StartsWithSpecialCharacter checks if the string begins with a character
// that is neither a letter nor a digit (e.g., ~, ♦, □, etc.).
func StartsWithSpecialCharacter(s string) bool {
	for _, r := range s {
		return !unicode.IsLetter(r) && !unicode.IsDigit(r)
	}
	return false
}

// RemoveSuffixes strips trailing punctuation characters from a string.
func RemoveSuffixes(s string) string {
	return strings.TrimRight(s, ".,;:-!?\"'()[]")
}

// StripZeroWidthChars removes zero-width Unicode characters from the edges of a string.
// These often appear in text files from OCR or copy-paste operations.
func StripZeroWidthChars(s string) string {
	return strings.Trim(s, "\u200B\uFEFF\u200D\u200C")
}
