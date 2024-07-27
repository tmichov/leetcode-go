package medium

func GenerateParentheses(n int) []string {
	combinations := []string{}

	var backtrack func(currentCombo string, opened int, closed int)

	backtrack = func(currentCombo string, opened int, closed int){
		if len(currentCombo) == n*2{
			combinations = append(combinations, currentCombo)
			return
		}

		if opened < n{
			backtrack(currentCombo + "(", opened+1, closed)
		}
		if opened > closed && closed < n{
			backtrack(currentCombo + ")", opened, closed+1)
		}
	}
	backtrack("", 0, 0)

	return combinations
}


