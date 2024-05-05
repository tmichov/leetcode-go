package easy

import "github.com/tmichov/leetcode/util"

func PathSum(root *util.TreeNode, targetSum int) bool {

	return depthFirstSearch(root, 0, targetSum)
}

func depthFirstSearch(node *util.TreeNode, sum int, targetSum int) bool {
	if node == nil {
		return false
	}

	if node.Left == nil && node.Right == nil {
		return sum + node.Val == targetSum
	}

	return depthFirstSearch(node.Left, sum+node.Val, targetSum) || depthFirstSearch(node.Right, sum+node.Val, targetSum)
}
