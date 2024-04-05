package easy

import "fmt"

func MakeGood(s string) string {
		stack := []byte{}

		for i := 0; i < len(s); i++ {
				if len(stack) > 0 && (stack[len(stack)-1] == s[i]+32 || stack[len(stack)-1] == s[i]-32) {
						stack = stack[:len(stack)-1]
				} else {
						stack = append(stack, s[i])
				}
		}

		return string(stack)
}
