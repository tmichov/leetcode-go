package main

import (
	"fmt"

	"github.com/tmichov/leetcode/medium"
)

func main() {
	idx := medium.SearchInRotatedArray([]int{1,3}, 3)

	fmt.Println(idx)
}
