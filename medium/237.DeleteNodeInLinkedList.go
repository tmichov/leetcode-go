package medium

import "github.com/tmichov/leetcode/util"

func DeleteNodeInLinkedList(node *util.ListNode) {
	handleNode(node, nil)
}

func handleNode(node *util.ListNode, prev *util.ListNode) {
	if node.Next == nil {
		prev.Next = nil
		return
	}

	node.Val = node.Next.Val

	handleNode(node.Next, node)
}
