package util

import (
	"fmt"
)

func PrintLinkedList(root *ListNode) {
	if root == nil {
		return
	}

	fmt.Println(root.Val)
	PrintLinkedList(root.Next)
}

func GenerateLinkedList(nums []int) (*ListNode, *ListNode) {
	var root *ListNode
	var curr *ListNode

	for i := 0; i < len(nums); i++ {
		if i == 0 {
			root = &ListNode{Val: nums[i]}
			curr = root
			continue
		}

		curr.Next = &ListNode{Val: nums[i]}

		curr = curr.Next
	}

	return root, curr
}

func PrintBinarySearchTree(head *TreeNode) {
	if head == nil {
		return
	}
	PrintBinarySearchTree(head.Left)
	println(head.Val)
	PrintBinarySearchTree(head.Right)
}

func GenerateBinaryTree() *TreeNode {
	root := &TreeNode{
		Val: 5,
	}

	root.Left = &TreeNode{Val: 4}
	root.Right = &TreeNode{Val: 6}

	root.Left.Left = &TreeNode{Val: 2}
	root.Right.Left = &TreeNode{Val: 5}
	root.Right.Right = &TreeNode{Val: 9}

	root.Left.Left.Left = &TreeNode{Val: 1}
	root.Left.Left.Right = &TreeNode{Val: 3}
	root.Right.Right.Right = &TreeNode{Val: 10}
	root.Right.Right.Left = &TreeNode{Val: 9}

	return root
}

