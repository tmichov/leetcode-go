package easy

import (
	"fmt"

	"github.com/tmichov/leetcode/util"
)

func SymmetricTree () {
	root := &util.TreeNode{Val: 1}

	root.Left = &util.TreeNode{Val: 4}
	root.Left.Left= &util.TreeNode{Val: 1}
	root.Left.Left.Left = &util.TreeNode{Val: 2}

	root.Right = &util.TreeNode{Val: 4}
	root.Right.Right = &util.TreeNode{Val: 1}
	root.Right.Right.Right= &util.TreeNode{Val: 2}

	answer := inorderZeros(root.Left, root.Right)

	fmt.Println(answer)
}

func inorderZeros(left *util.TreeNode, right *util.TreeNode) bool {
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
