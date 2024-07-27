package medium

func LongestSubscringWithoutRepeatingCharacters(s string) int {
	charIndexMap := make(map[byte]int)
	maxLength := 0
	start := 0

	for end := 0; end < len(s); end++ {
		if index, found := charIndexMap[s[end]]; found && index >= start {
			start = index + 1
		}
		charIndexMap[s[end]] = end
		currentLength := end - start + 1
		if currentLength > maxLength {
			maxLength = currentLength
		}
	}

	return maxLength
}
