package _8

import "testing"

func TestMergeTwoLists(t *testing.T) {
	l1 := &ListNode{1, &ListNode{2, &ListNode{4, nil}}}
	l2 := &ListNode{1, &ListNode{3, &ListNode{4, nil}}}
	mergeTwoLists(l1, l2)
}
