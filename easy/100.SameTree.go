package easy

import (
	"fmt"

	"github.com/tmichov/leetcode/util"
)

//       root = 1
//      /       \
//      2        3
//		/   \    /    \
//	 4     5  6      7
//   \
//    8
func SameTree() {
	root := &util.TreeNode{Val: 1}

	root.Left = &util.TreeNode{Val: 2}
	root.Right = &util.TreeNode{Val: 3}

	root1 := &util.TreeNode{Val: 1}

	root1.Left = &util.TreeNode{Val: 2}
	root1.Right = &util.TreeNode{Val: 3}
	root1.Right = &util.TreeNode{Val: 3}
	
	// found at 96.Inorder Traversing tree
	a := inorderCompare(root, root1)

	fmt.Println(a)
}


func inorderCompare(p *util.TreeNode, q *util.TreeNode) bool {
	if p == nil && q == nil {
		return true
	}

	if (p == nil && q != nil) || (p != nil && q == nil) {
		return false
	}

	if p.Val != q.Val {
		return false
	}

	check := inorderCompare(p.Left, q.Left)
	if(!check) {
		return false
	}


	check2 := inorderCompare(p.Right, q.Right)
	if !check2{
		return false
	}

	return true
}
