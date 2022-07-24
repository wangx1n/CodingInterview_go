package main

func swapPairs(head *ListNode) *ListNode {
	reverse := func(head, tail *ListNode) (*ListNode, *ListNode) {
		tHead := head
		nex := tail.Next
		for nex != tail {
			tmp := tHead.Next
			tHead.Next = nex
			nex = tHead
			tHead = tmp
		}
		return tail, head
	}
	hair := &ListNode{Next: head}
	prev := hair
	for head != nil {
		tail := prev
		for i := 0; i < 2; i++ {
			tail = tail.Next
			if tail == nil {
				return hair.Next
			}
		}
		nex := tail.Next
		head, tail = reverse(head, tail)
		prev.Next = head
		tail.Next = nex
		prev = tail
		head = nex
	}
	return hair.Next
}
