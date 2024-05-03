package main

import (
	"fmt"

	"github.com/tmichov/leetcode/easy"
	"github.com/tmichov/leetcode/util"
)

func main() {
	nums := []int{-10, -3, 0, 5, 9}
	root := easy.ConvertSorterArrayToBST(nums)
	util.PrintBinarySearchTree(root)
	
	fmt.Println()
}

