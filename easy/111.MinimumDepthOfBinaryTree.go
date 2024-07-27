package easy

import "github.com/tmichov/leetcode/util"

func MinimumDepthOfBinaryTree(root *util.TreeNode) int {
	if root == nil {
		return 0
	}

	leftDepth := MinimumDepthOfBinaryTree(root.Left)
	rightDepth := MinimumDepthOfBinaryTree(root.Right)

	return 1 + min(leftDepth, rightDepth)
}

func min(leftDepth, rightDepth int) int {
	if leftDepth == 0 {
		return rightDepth
	}

	if rightDepth == 0 {
		return leftDepth
	}

	if leftDepth < rightDepth {
		return leftDepth
	}

	return rightDepth
}
