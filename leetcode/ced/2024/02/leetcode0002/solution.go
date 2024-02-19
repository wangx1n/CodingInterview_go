package leetcode0002

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var carry, val int
	var tt *ListNode
	ans := &ListNode{}

	ll1 := l1
	ll2 := l2

	if ll1 == nil {
		return ll2
	}
	if ll2 == nil {
		return ll1
	}

	for {
		ll1Val, ll2Val := 0, 0
		if ll1 != nil {
			ll1Val = ll1.Val
		}
		if ll2 != nil {
			ll2Val = ll2.Val
		}

		t := ll1Val + ll2Val + carry
		val = t % 10
		carry = t / 10

		if ll1 == nil && ll2 == nil && val == 0 && carry == 0 {
			break
		}

		if tt == nil {
			tt = &ListNode{Val: val}
			ans.Next = tt
		} else {
			tt.Next = &ListNode{Val: val}
			tt = tt.Next
		}

		if ll1 != nil {
			ll1 = ll1.Next
		}
		if ll2 != nil {
			ll2 = ll2.Next
		}
	}

	return ans.Next
}
