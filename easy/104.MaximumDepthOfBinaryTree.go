package easy

import (
	"fmt"

	"github.com/tmichov/leetcode/util"
)

func MaxDepth() {
	root := &util.TreeNode{Val: 3}

	root.Left = &util.TreeNode{Val: 9}
	root.Right = &util.TreeNode{Val: 20}

	root.Right.Left = &util.TreeNode{Val: 15}
	root.Right.Right = &util.TreeNode{Val: 15}

	if root == nil {
		fmt.Println(0)
		return
	}

	depth := treeDepth(root) + 1

	fmt.Println(depth)
}

func treeDepth(node *util.TreeNode) int {
	if node == nil || (node.Left == nil && node.Right == nil) {
		return 0
	} else {
		leftInt := treeDepth(node.Left)
		rightInt := treeDepth(node.Right)

		fmt.Println(leftInt, rightInt)

		if leftInt > rightInt {
			return leftInt + 1
		} else {
			return rightInt + 1
		}
	}
}



