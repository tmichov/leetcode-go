package util

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type ListNode struct {
	Val  int
	Next *ListNode
}

type LinkedList struct {
	length int
	Head *LinkedListNode
	Tail *LinkedListNode
}

type LinkedListNode struct {
	Val int
	Next *LinkedListNode
	Prev *LinkedListNode
}

func InitializeLinkedList() *LinkedList {
	return &LinkedList{}
}

func (ll *LinkedList) Append(val int) {
	ll.length++

	if ll.Head == nil {
		ll.Head = &LinkedListNode{Val: val}
		ll.Tail = ll.Head
		return
	}

	ll.Tail.Next = &LinkedListNode{Val: val}
	ll.Tail = ll.Tail.Next
}

func (ll *LinkedList) Prepend(val int) {
	if ll.Head == nil {
		ll.Append(val)
		return
	}

	ll.length++

	newHead := &LinkedListNode{Val: val}
	newHead.Next = ll.Head;
	ll.Head.Prev = newHead;
	ll.Head = newHead;
}

func (ll *LinkedList) DeleteAt(idx int) {
	if ll.Head == nil {
		return
	}

	if(idx >= ll.length) {
		return
	}

	ll.length--

	if idx == 0 {
		ll.Head = ll.Head.Next
		ll.Head.Prev = nil
		return
	}

	curr := ll.Head
	for i := 0; i < idx - 1; i++ {
		curr = curr.Next
	}
}

func (ll *LinkedList) Print() {
	curr := ll.Head

	for curr != nil {
		println(curr.Val)
		curr = curr.Next
	}
}

