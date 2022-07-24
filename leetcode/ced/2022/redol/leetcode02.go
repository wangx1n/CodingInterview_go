package redol

// ListNode Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	carry := 0
	var head *ListNode
	tail := head
	for l1 != nil || l2 != nil {
		n1, n2 := 0, 0
		if l1 != nil {
			n1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			n2 = l2.Val
			l2 = l2.Next
		}
		sum := n1 + n2 + carry
		if head == nil {
			head = &ListNode{
				Val: sum % 10,
			}
			tail = head
		} else {
			tail.Next = &ListNode{
				Val: sum % 10,
			}
			tail = tail.Next
		}
		carry = sum / 10
	}
	if carry > 0 {
		tail.Next = &ListNode{
			Val: carry,
		}
	}
	return head
}
