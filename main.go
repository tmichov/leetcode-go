package main

import (
	"fmt"

	"github.com/tmichov/leetcode/easy"
)


func main() {
	scores := []int{5,4,3,2,1}

	listed := easy.RelativeRanks(scores)

	fmt.Println(listed)
}

