package _8

func reverseKGroup(head *ListNode, k int) *ListNode {
	hair := &ListNode{Next: head}
	prev := hair
	tail := hair
	for tail.Next != nil {
		for i := 0; i < k; i++ {
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

func reverse(head, tail *ListNode) (*ListNode, *ListNode) {
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
