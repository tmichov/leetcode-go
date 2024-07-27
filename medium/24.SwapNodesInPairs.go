package medium

import (
	"github.com/tmichov/leetcode/util"
)

func SwapNodesInPairs(head *util.ListNode) *util.ListNode {
	if head == nil {
		return head
	}
	h := head
	curr := head
	var prev *util.ListNode = nil

	if curr.Next == nil {
		return curr
	}

	for curr != nil && curr.Next != nil {
		next := curr.Next

		swap(curr, next)
		if prev == nil {
			h = next
		} else {
			prev.Next = next
		}

		prev = curr
		curr = curr.Next
	}

	return h
}

func swap(n1, n2 *util.ListNode) (*util.ListNode, *util.ListNode) {
	tmp := n1

	tmp.Next = n2.Next
	n1 = n2

	n2 = tmp
	n1.Next = n2

	return n1, n2
}
