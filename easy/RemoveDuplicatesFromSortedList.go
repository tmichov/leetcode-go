package easy

import "fmt"

type ListNode struct {
		Val int
		Next *ListNode
}

//1 -> 1 -> 2 -> 2
func RemoveDuplicatesFromSortedList(head *ListNode) *ListNode {
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
