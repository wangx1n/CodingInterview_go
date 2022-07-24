package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	return merge(lists, 0, len(lists)-1)
}

func merge(lists []*ListNode, l, r int) *ListNode {
	if l == r {
		return lists[l]
	}
	if l > r {
		return nil
	}
	for l < r {
		mid := (l + r) >> 1
		return mergeList(merge(lists, l, mid), merge(lists, mid+1, r))
	}
	return nil
}

func mergeList(a, b *ListNode) *ListNode {
	h := &ListNode{}
	tail := h
	for a != nil && b != nil {
		if a.Val < b.Val {
			tail.Next = a
			a = a.Next
		} else {
			tail.Next = b
			b = b.Next
		}
		tail = tail.Next
	}
	if a == nil {
		tail.Next = b
	} else {
		tail.Next = a
	}
	return h.Next
}
