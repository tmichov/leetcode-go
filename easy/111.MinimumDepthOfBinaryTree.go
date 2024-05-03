package easy

func MinimumDepthOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}

	leftDepth := MinimumDepthOfBinaryTree(root.Left)
	rightDepth := MinimumDepthOfBinaryTree(root.Right)

	return 1+min(leftDepth, rightDepth)
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

