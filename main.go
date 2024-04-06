package main

import (
	"fmt"

	"github.com/tmichov/leetcode/medium"
)

func main() {
	string := "lee(t(c)o)de)"

	s := medium.MinimumRemoveToMakeValidParentheses(string)

	fmt.Println(s)
}

