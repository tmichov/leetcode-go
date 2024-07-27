package easy

import (
	"strings"
)

func ValidPalindrome(s string) bool {
	lcs := strings.ToLower(s)
	start := 0
	end := len(s) - 1

	for start < end {
		if !isChar(lcs[start]) {
			start++
			continue
		}

		if !isChar(lcs[end]) {
			end--
			continue
		}

		if lcs[start] != lcs[end] {
			return false
		}

		start++
		end--
	}

	return true
}

func isChar(s byte) bool {
	return (s >= 'a' && s <= 'z') || (s >= '0' && s <= '9')
}
