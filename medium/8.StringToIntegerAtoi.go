package medium

import (
	"strings"
)

func StringToIntegerAtoi(s string) int {
	s = strings.TrimSpace(s)
	if s == "" {
		return 0
	}
	maxNum := (1 << 31) - 1

	negative := 1
	num := 0

	if s[0] == '-' {
		negative = -1
		s = s[1:]
	} else if s[0] == '+' {
		s = s[1:]
	}

	for _, char := range s {
		if char < '0' || char > '9' {
			break
		}

		num = (num * 10) + int(char-'0')

		if num > maxNum {
			num = maxNum
			if negative == -1 {
				num += 1
			}
			break
		}
	}

	return num * negative
}
