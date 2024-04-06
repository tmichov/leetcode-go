package medium

import "fmt"

func MinimumRemoveToMakeValidParentheses(s string) string {
	var stack []int
	chars := []rune(s)

	for i, value := range chars {
		if value == '(' {
			stack = append(stack, i)
		} else if value == ')' {
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			} else {
				chars[i] = 0
			}
		}
	}

	for _, index := range stack {
		chars[index] = 0
	}

	result := make([]rune, 0, len(chars))
	for _, char := range chars {
		if char != '0' && char != '\u0000' {
			result = append(result, char)

			fmt.Println(result)
		}
	}

	fmt.Println(result)
	return string(result)
}
