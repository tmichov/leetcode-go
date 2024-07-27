package medium

func LongestPalindromicSubstring(s string) string {
	if len(s) == 0 {
		return ""
	}

	start, maxLength := 0, 0

	for i := 0; i < len(s); i++ {
		len1 := expandAroundCenter(s, i, i)
		len2 := expandAroundCenter(s, i, i+1)

		length := max(len1, len2)
		if length > maxLength {
			maxLength = length
			start = i - (length-1)/2
		}
	}

	return s[start : start+maxLength]
}

func max(i, j int) int {
	if i > j {
		return i
	}

	return j
}

func expandAroundCenter(s string, left, right int) int {
	for left >= 0 && right < len(s) && s[left] == s[right] {
		left--
		right++
	}

	return right - left - 1
}
