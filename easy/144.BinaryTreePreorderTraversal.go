package easy

import "github.com/tmichov/leetcode/util"

func BinaryTreePreorderTraversal(root *util.TreeNode) []int {
	list := []int{root.Val}

	if root.Left != nil {
		list = append(list, BinaryTreePreorderTraversal(root.Left)...)
	}

	if root.Right != nil {
		list = append(list, BinaryTreePreorderTraversal(root.Right)...)
	}

	return list
}
