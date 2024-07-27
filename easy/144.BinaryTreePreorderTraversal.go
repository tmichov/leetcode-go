package easy

import "github.com/tmichov/leetcode/util"

func BinaryTreePreorderTraversal(root *util.TreeNode) []int {
	list := make([]int, 0)
	if root == nil {
		return list
	}

	list = append(list, root.Val)
	list = append(list, BinaryTreePreorderTraversal(root.Left)...)
	list = append(list, BinaryTreePreorderTraversal(root.Right)...)

	return list
}
