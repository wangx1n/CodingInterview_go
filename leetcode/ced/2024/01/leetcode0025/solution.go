package leetcode0025

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	hair := &ListNode{Next: head}
	prev := hair

	for head != nil {
		tail := prev
		for i := 0; i < k; i++ {
			tail = tail.Next
			if tail == nil {
				return hair.Next
			}
		}

		next := tail.Next

		head, tail = reverse(head, tail)
		prev.Next = head
		tail.Next = next
		prev = tail
		head = next
	}
	return hair.Next
}

func reverse(head, tail *ListNode) (*ListNode, *ListNode) {
	prev := &ListNode{
		Next: head,
	}
	p := head

	for prev != tail {
		next := p.Next
		p.Next = prev
		prev = p
		p = next
	}

	return tail, head
}
