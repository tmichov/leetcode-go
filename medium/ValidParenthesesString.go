package medium

func ValidPerenthesesString(s string) bool {
	leftMin := 0
	leftMax := 0

	for _, val := range s {
		if val == '(' {
			leftMin++
			leftMax++
		} else if val == ')' {
			leftMin--
			leftMax--
		} else {
			leftMax++
			leftMin--
		}

		leftMin = max(0, leftMin)

		if leftMax < 0 {
			return false
		}
	}

	return leftMin == 0
}
