package main

import (
	"github.com/tmichov/leetcode/medium"
	"github.com/tmichov/leetcode/util"
)


func main() {
	root := &util.ListNode{}
	root.Val = 5
	root.Next = &util.ListNode{}
	root.Next.Val = 4
	root.Next.Next = &util.ListNode{}
	root.Next.Next.Val = 3
	root.Next.Next.Next = &util.ListNode{}
	root.Next.Next.Next.Val = 2

	medium.DeleteNodeInLinkedList(root.Next)
	util.PrintLinkedList(root)
}

