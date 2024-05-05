package main

import (
	"fmt"

	"github.com/tmichov/leetcode/easy"
	"github.com/tmichov/leetcode/util"
)


func main() {
	root, last := util.GenerateLinkedList([]int{1, 1, 2, 2, 3, 4, 5, 6})

	last.Next = root.Next

	test := easy.LinkedListCycle(root)

	fmt.Println(test)
}

