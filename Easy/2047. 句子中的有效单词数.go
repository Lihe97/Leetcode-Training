package main

import (
	"strings"
	"unicode"
)

func countValidWords(sentence string) int {
	res := 0

	arr := strings.Fields(sentence)
	for _, word := range arr {
		if valid(word) {
			res++
		}
	}
	return res

}
func valid(s string) bool {
	hasHyphens := false
	for i, ch := range s {
		if unicode.IsDigit(ch) || strings.ContainsRune("!.,", ch) && i < len(s)-1 {
			return false
		}
		if ch == '-' {
			if hasHyphens || i == 0 || i == len(s)-1 || !unicode.IsLower(rune(s[i-1])) || !unicode.IsLower(rune(s[i+1])) {
				return false
			}
			hasHyphens = true
		}
	}
	return true
}
