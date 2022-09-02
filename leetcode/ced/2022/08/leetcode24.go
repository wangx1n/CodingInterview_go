package _8

func swapPairs(head *ListNode) *ListNode {
	hair := &ListNode{Next: head}
	tail := hair
	prev := hair
	for tail != nil {
		for i := 0; i < 2; i++ {
			tail = tail.Next
			if tail == nil {
				return hair.Next
			}
		}
		nex := tail.Next
		head, tail = reverseList(head, tail)
		prev.Next = head
		tail.Next = nex
		head = nex
		prev = tail
	}
	return hair.Next
}

func reverseList(head, tail *ListNode) (*ListNode, *ListNode) {
	nex := tail.Next
	tHead := head
	for nex != tail {
		tmp := tHead.Next
		tHead.Next = nex
		nex = tHead
		tHead = tmp
	}
	return tail, head
}
