package _8

func mergeKLists(lists []*ListNode) *ListNode {
	return merge(0, len(lists)-1, lists)
}

func merge(list1Idx, list2Idx int, lists []*ListNode) *ListNode {
	if list1Idx == list2Idx {
		return lists[list1Idx]
	}
	if list1Idx > list2Idx {
		return nil
	}
	mid := (list1Idx + list2Idx) / 2
	return mergeTwoLists_(merge(list1Idx, mid, lists), merge(mid+1, list2Idx, lists))
}

func mergeTwoLists_(a, b *ListNode) *ListNode {
	head := &ListNode{}
	tail := head
	for a != nil && b != nil {
		if a.Val < b.Val {
			tail.Next = a
			tail = tail.Next
			a = a.Next
		} else {
			tail.Next = b
			tail = tail.Next
			b = b.Next
		}
	}
	if a == nil {
		tail.Next = b
	} else {
		tail.Next = a
	}
	return head.Next
}
