package medium

import (
	"fmt"

	"github.com/tmichov/leetcode/util"
)

func RemoveNthNodeFromEndOfLinkedList(head *util.ListNode, n int) *util.ListNode {
	if head.Next == nil {
		return nil
	}

	var prev *util.ListNode = nil

	curr := head
	nth := head
	count := 1

	for curr.Next != nil {
		curr = curr.Next

		if count >= n {
			prev = nth
			nth = nth.Next
		}

		count++
	}

	fmt.Println("prev", prev, "nth", nth, "curr", curr)

	if prev != nil {
		prev.Next = nth.Next
	} else {
		head = head.Next
	}
	return head
}
