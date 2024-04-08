package easy

import "fmt"

func MaxDepth() {
	root := &TreeNode{Val: 3}

	root.Left = &TreeNode{Val: 9}
	root.Right = &TreeNode{Val: 20}

	root.Right.Left = &TreeNode{Val: 15}
	root.Right.Right = &TreeNode{Val: 15}

	if root == nil {
		fmt.Println(0)
		return
	}

	depth := treeDepth(root) + 1

	fmt.Println(depth)
}

func treeDepth(node *TreeNode) int {
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



