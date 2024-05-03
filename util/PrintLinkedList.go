package util

import "github.com/tmichov/leetcode/easy"

func PrintBinarySearchTree(head *easy.TreeNode) {
	if head == nil {
		return
	}
	PrintBinarySearchTree(head.Left)
	println(head.Val)
	PrintBinarySearchTree(head.Right)
}
