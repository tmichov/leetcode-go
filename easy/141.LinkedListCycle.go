package easy

import "github.com/tmichov/leetcode/util"

func LinkedListCycle(head *util.ListNode) bool {
	if head == nil {
		return false
	}

	fast := head.Next
	slow := head

	for fast != nil && fast.Next != nil {
		if slow == fast {
			return true
		}

		slow = slow.Next
		fast = fast.Next.Next
	}

	return false
}
