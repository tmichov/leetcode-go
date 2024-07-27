package medium

import (
	"github.com/tmichov/leetcode/util"
)

func AddTwoNumbers(l1, l2 *util.ListNode) *util.ListNode {
	root := &util.ListNode{
		Val:  0,
		Next: nil,
	}

	generateNodes(nil, root, l1, l2)

	return root
}

func generateNodes(prev, node, l1, l2 *util.ListNode) {
	if node == nil && l1 == nil && l2 == nil {
		return
	}

	if node == nil {
		node = &util.ListNode{Val: 0}
		prev.Next = node
	}

	var next, next1, next2 *util.ListNode

	if l1 != nil {
		next1 = l1.Next
	}
	if l2 != nil {
		next2 = l2.Next
	}

	sum, remainder := addNumbers(node, l1, l2)

	node.Val = sum

	if remainder > 0 {
		next = &util.ListNode{Val: remainder}
		node.Next = next
	}

	generateNodes(node, node.Next, next1, next2)
}

func addNumbers(root, l1, l2 *util.ListNode) (int, int) {
	val1, val2 := 0, 0

	if l1 != nil {
		val1 = l1.Val
	}

	if l2 != nil {
		val2 = l2.Val
	}

	result := root.Val + val1 + val2

	return result % 10, int(result / 10)
}
