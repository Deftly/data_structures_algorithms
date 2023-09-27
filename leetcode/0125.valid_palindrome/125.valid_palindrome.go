package data_structures_algorithms

import (
	"strings"
	"unicode"
)

func isPalindrome(s string) bool {
	mapFunc := func(r rune) rune {
		if !unicode.IsLetter(r) && !unicode.IsDigit(r) {
			return -1
		}

		return unicode.ToLower(r)
	}

	s = strings.Map(mapFunc, s)

	i, j := 0, len(s)-1
	for i < j {
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}
	return true
}
