package isogram

import (
	"strings"
	"unicode"
)

func IsIsogram(word string) bool {
	letterCounter := map[rune]bool{}

	for _, r := range strings.ToLower(word) {
		if !unicode.IsLetter(r) {
			continue
		}

		if letterCounter[r] {
			return false
		}

		letterCounter[r] = true
	}

	return true
}
