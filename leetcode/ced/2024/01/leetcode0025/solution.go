package leetcode0025

import "sort"

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func totalNQueens(n int) int {
	column := map[int]bool{}
	LToR, RToL := map[int]bool{}, map[int]bool{}
	queen := make([]int, n)
	for i := 0; i < n; i++ {
		queen[i] = -1
	}

	ans := make([][]string, 0, n)

	dfs(n, 0, column, LToR, RToL, queen, &ans)

	return len(ans)
}

func dfs(n, row int, column map[int]bool, LToR, RToL map[int]bool, queen []int, ans *[][]string) {
	if row == n {
		rr := buildAns(queen, n)
		*ans = append(*ans, rr)
		return
	}

	for i := 0; i < n; i++ {
		if column[i] {
			continue
		}
		ltor := row - i
		if LToR[ltor] {
			continue
		}
		rtol := row + i
		if RToL[rtol] {
			continue
		}

		queen[row] = i
		column[i] = true
		LToR[ltor] = true
		RToL[ltor] = true
		dfs(n, row+1, column, LToR, RToL, queen, ans)
		queen[row] = -1
		column[i] = false
		LToR[ltor] = false
		RToL[ltor] = false
	}
}

func buildAns(queen []int, n int) []string {
	ret := []string{}
	for i := 0; i < n; i++ {
		ans := make([]byte, n)
		for j := 0; j < n; j++ {
			ans[j] = '.'
		}
		ans[i] = 'Q'
		ret = append(ret, string(ans))
	}
	return ret
}

func FindSubstring(s string, words []string) []int {
	landmine := map[string]int{}
	n := len(words) * len(words[0])

	for i := 0; i < len(words); i++ {
		landmine[words[i]]++
	}

	left, right := 0, n

	ans := []int{}

	for right <= len(words) {
		subString := s[left:right]
		copyed := copyLandMine(landmine)
		for i := 0; i < n; i += len(words[0]) {
			ss := subString[i : i+len(words[0])]
			copyed[ss]--
			if copyed[ss] < 0 {
				break
			}
			if copyed[ss] == 0 {
				delete(copyed, ss)
			}
		}
		if len(copyed) == 0 {
			ans = append(ans, left)
		}
		left++
		right++
	}
	return ans
}

func copyLandMine(landmine map[string]int) map[string]int {
	ret := make(map[string]int, len(landmine))
	for k, v := range landmine {
		ret[k] = v
	}
	return ret
}

func minWindow(s string, t string) string {
	if len(s) > len(t) {
		return ""
	}

	letterWithIndex := map[byte]int{}
	uniqLetter := map[byte]bool{}

	for i := 0; i < len(t); i++ {
		uniqLetter[t[i]] = true
	}

	for i := 0; i < len(s); i++ {
		if uniqLetter[s[i]] {
			letterWithIndex[s[i]] = i
		}
	}

	ans := []int{}
	for _, v := range letterWithIndex {
		ans = append(ans, v)
	}

	sort.Ints(ans)

	ret := s[ans[0]:ans[len(ans)-1]]
	return ret
}

func rotateRight(head *ListNode, k int) *ListNode {
	hair := &ListNode{Next: head}
	tHead := hair
	tail := hair

	n := -1
	for tHead != nil {
		n++
		tHead = tHead.Next
	}

	newTailIndex := n - (k % n)
	if newTailIndex == 0 {
		return head
	}

	tail.Next = head
	target := hair
	for i := 0; i < newTailIndex; i++ {
		target = target.Next
	}

	head = target.Next
	target.Next = nil

	return hair.Next
}

type LRUCache struct {
	Head, Tail *dualLinkList
	IndexMap   map[int]*dualLinkList
	Size, Cap  int
}

type dualLinkList struct {
	Key, Val   int
	Prev, Next *dualLinkList
}

func constructDualLinkList(key, val int) *dualLinkList {
	return &dualLinkList{
		Key: key,
		Val: val,
	}
}

func Constructor(capacity int) LRUCache {
	head, tail := constructDualLinkList(0, 0), constructDualLinkList(0, 0)
	head.Next = tail
	tail.Prev = head
	return LRUCache{
		Head:     head,
		Tail:     tail,
		IndexMap: map[int]*dualLinkList{},
		Size:     0,
		Cap:      capacity,
	}
}

func (this *LRUCache) moveToHead(key int) {
	node := this.IndexMap[key]
	node.Prev.Next = node.Next
	node.Next.Prev = node.Prev
	this.addToHead(key)
}

func (this *LRUCache) addToHead(key int) {
	node := this.IndexMap[key]
	next := this.Head.Next

	this.Head.Next = node
	next.Prev = node

	node.Next = next
	node.Prev = this.Head
}

// 不存在：返回负一
// 存在：移动至链表头结点
//
//	先删除节点
//	再添加至头结点
func (this *LRUCache) Get(key int) int {
	// 不存在，返回-1
	dualLinkList, ok := this.IndexMap[key]
	if !ok {
		return -1
	}
	this.moveToHead(dualLinkList.Key)
	return dualLinkList.Val
}

func (this *LRUCache) removeTail() {
	prev := this.Tail.Prev
	tail := this.Tail

	prev.Prev.Next = tail
	this.Tail.Prev = prev.Prev
}

// 存在，更新值
// 移动至头结点
// 不存在，添加到头结点
func (this *LRUCache) Put(key int, value int) {
	// 存在，更新值
	dd := this.IndexMap[key]

	if dd != nil {
		dd.Val = value
		this.moveToHead(dd.Key)
		return
	}

	node := &dualLinkList{
		Key: key,
		Val: value,
	}
	this.IndexMap[key] = node
	this.Size += 1
	this.addToHead(node.Key)

	if this.Size > this.Cap {
		this.removeTail()
		this.Size -= 1
	}

	return
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	queue1 := make([]*TreeNode, 0, 10)
	queue2 := make([]*TreeNode, 0, 10)

	queue1 = append(queue1, p)
	queue2 = append(queue2, q)

	for len(queue1) != 0 || len(queue2) != 0 {
		if len(queue1) != len(queue2) {
			return false
		}

		newQueue1 := make([]*TreeNode, 0, 10)
		newQueue2 := make([]*TreeNode, 0, 10)

		for i := 0; i < len(queue1); i++ {
			node := queue1[i]
			if node == nil {
				newQueue1 = append(newQueue1, nil)
				newQueue1 = append(newQueue1, nil)
				continue
			}
			newQueue1 = append(newQueue1, node.Left)
			newQueue1 = append(newQueue1, node.Right)
		}

		for i := 0; i < len(queue2); i++ {
			node := queue2[i]
			if node == nil {
				newQueue1 = append(newQueue1, nil)
				newQueue1 = append(newQueue1, nil)
				continue
			}
			newQueue2 = append(newQueue2, node.Left)
			newQueue2 = append(newQueue2, node.Right)
		}

		for i := 0; i < len(queue1); i++ {
			if queue1[i] == nil {

			}
			if queue1[i] != queue2[i] {
				return false
			}
		}

		queue1 = make([]*TreeNode, len(newQueue1))
		queue2 = make([]*TreeNode, len(newQueue2))

		copy(queue1, newQueue1)
		copy(queue2, newQueue2)
	}

	if len(queue1) != len(queue2) {
		return false
	}

	return true
}

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	queue := []*Node{}
	queue = append(queue, root)

	for len(queue) > 0 {
		newQueue := make([]*Node, 0, len(queue)*2)
		n := len(queue)
		for i := 0; i < n; i++ {
			if queue[i] == nil {
				continue
			}
			if queue[i].Left != nil {
				newQueue = append(newQueue, queue[i].Left)
			}
			if queue[i].Right != nil {
				newQueue = append(newQueue, queue[i].Right)
			}
			if i+1 < n {
				queue[i].Right = queue[i+1]
				continue
			}
			queue[i].Right = nil
		}
		queue = make([]*Node, len(newQueue))
		copy(queue, newQueue)
	}

	return root
}
