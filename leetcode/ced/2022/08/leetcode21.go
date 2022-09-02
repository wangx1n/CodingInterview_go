package _8

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil && list2 == nil {
		return nil
	}
	var head *ListNode
	var rear *ListNode

	for list1 != nil || list2 != nil {
		if list1 == nil {
			if head == nil {
				head = &ListNode{0, list2}
				break
			}
			rear.Next = list2
			break
		}
		if list2 == nil {
			if head == nil {
				head = &ListNode{0, list1}
				break
			}
			rear.Next = list1
			break
		}
		val1, val2 := list1.Val, list2.Val
		if val1 < val2 {
			if head == nil {
				head = &ListNode{0, list1}
				rear = head.Next
				list1 = list1.Next
			} else {
				rear.Next = list1
				rear = rear.Next
				list1 = list1.Next
			}
		} else {
			if head == nil {
				head = &ListNode{0, list2}
				rear = head.Next
				list2 = list2.Next
			} else {
				rear.Next = list2
				rear = rear.Next
				list2 = list2.Next
			}
		}
	}
	return head.Next
}
