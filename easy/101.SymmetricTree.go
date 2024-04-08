package easy

import "fmt"

func SymmetricTree () {
	root := &TreeNode{Val: 1}

	root.Left = &TreeNode{Val: 4}
	root.Left.Left= &TreeNode{Val: 1}
	root.Left.Left.Left = &TreeNode{Val: 2}

	root.Right = &TreeNode{Val: 4}
	root.Right.Right = &TreeNode{Val: 1}
	root.Right.Right.Right= &TreeNode{Val: 2}

	answer := inorderZeros(root.Left, root.Right)

	fmt.Println(answer)
}

func inorderZeros(left *TreeNode, right *TreeNode) bool {
	if left == nil && right == nil {
		return true;
	}

	if left == nil || right == nil {
		return false;
	}

	if left.Val != right.Val {
		return false
	}

	check := inorderZeros(left.Left, right.Right)
	check2 := inorderZeros(left.Right, right.Left)

	return check && check2
}
