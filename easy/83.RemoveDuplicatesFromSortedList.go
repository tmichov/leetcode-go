package easy

import (
	"fmt"

	"github.com/tmichov/leetcode/util"
)

//1 -> 1 -> 2 -> 2
func RemoveDuplicatesFromSortedList(head *util.ListNode) *util.ListNode {
		if head == nil {
				return head
		}

		if head.Next == nil {
				return head
		}

		h := head

		next := head.Next;

		for next != nil {
				if next.Val > head.Val {
						head.Next = next
						head = head.Next
				} else {
						head.Next = nil
				}
				next = next.Next
		}
		
		for h != nil {
				fmt.Println(h)
				h = h.Next
		}

		return h
}
