package easy

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//       root = 1
//      /       \
//      2        3
//		/   \    /    \
//	 4     5  6      7
//   \
//    8
func BinaryTreeInorderTraversal() {
	root := &TreeNode{Val: 1}

	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 3}

	root.Left.Left = &TreeNode{Val: 4}
	root.Left.Right = &TreeNode{Val: 5}

	root.Right.Left = &TreeNode{Val: 6}
	root.Right.Right = &TreeNode{Val: 7}

	root.Left.Left.Right = &TreeNode{Val: 8}
	
	a := RecurringBinaryTreeInorderTraversal(root)
//	a := IterativeBinaryTreeInorderTraversal(root)

	fmt.Println(a)
}


func inorder(node *TreeNode, nums []int) []int {
	if node == nil {
		return nums
	}

	nums = inorder(node.Left, nums)

	nums = append(nums, node.Val)

	nums = inorder(node.Right, nums)

	return nums
}

// Solution 1 Recursive
func RecurringBinaryTreeInorderTraversal(root *TreeNode) []int {
	a := inorder(root, []int{})

	return a
}

// Solution 2 Iterative
func IterativeBinaryTreeInorderTraversal(root *TreeNode) []int {
	nums := []int{}
	list := []*TreeNode{}

	curr := root

	for len(list) > 0 || curr != nil {
		if curr != nil {
			list = append(list, curr)
			curr = curr.Left
		} else {
			curr = list[len(list)-1]
			nums = append(nums, curr.Val)
			list = list[:len(list)-1]	

			curr = curr.Right
		}
	}

	return nums
}

