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
