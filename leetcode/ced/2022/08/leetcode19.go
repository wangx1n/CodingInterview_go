package _8

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	hair := &ListNode{0, head}
	front, rear := hair, head

	for i := 0; i < n; i++ {
		rear = rear.Next
	}

	for ; rear != nil; rear = rear.Next {
		front = front.Next
	}

	front.Next = front.Next.Next

	return hair.Next
}
