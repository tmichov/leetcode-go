package util

import "fmt"

func PrintLinkedList(root *ListNode) {
	if root == nil {
		return
	}

	fmt.Println(root.Val)
	PrintLinkedList(root.Next)
}

func PrintBinarySearchTree(head *TreeNode) {
	if head == nil {
		return
	}
	PrintBinarySearchTree(head.Left)
	println(head.Val)
	PrintBinarySearchTree(head.Right)
}
