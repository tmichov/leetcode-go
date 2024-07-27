package easy

func MaximumNestingDepthOfParentheses(s string) int {
	maxDepth := 0
	depth := 0

	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			depth++
			maxDepth = max(depth, maxDepth)
		} else if s[i] == ')' {
			depth--
		}
	}

	return maxDepth
}
