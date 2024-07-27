package easy

import "github.com/tmichov/leetcode/util"

func getIntersectionNode(headA, headB *util.ListNode) *util.ListNode {
	lengthA, lastA := getLastElement(1, headA)
	lengthB, lastB := getLastElement(1, headB)

	if lastA != lastB {
		return nil
	}

	var rootA, rootB *util.ListNode
	if lengthA < lengthB {
		rootA, rootB = getRootElements(headB, headA, lengthB-lengthA)
	} else {
		rootA, rootB = getRootElements(headA, headB, lengthA-lengthB)
	}

	for rootA != rootB {
		rootA = rootA.Next
		rootB = rootB.Next
	}

	return rootA
}

func getRootElements(rootA, rootB *util.ListNode, count int) (*util.ListNode, *util.ListNode) {
	if count == 0 {
		return rootA, rootB
	}

	rootA = rootA.Next
	return getRootElements(rootA, rootB, count-1)
}

func getLastElement(count int, root *util.ListNode) (int, *util.ListNode) {
	if root.Next == nil {
		return count, root
	}

	return getLastElement(count+1, root.Next)
}
