package _1

import (
	"math/rand"
	common "gitee.com/wangx1n9347/common_go/List"
)

type Solution struct {
	head *common.ListNode
}

func Constructor(head *common.ListNode) Solution {
	return Solution{head}
}

// GetRandom 水塘抽样
func (s *Solution) GetRandom() (ans int) {
	k := 1 // 随机选几个k就是几，这道题k是1
	for node, i := s.head, k; node != nil; node = node.Next {
		if rand.Intn(i) < k { // 1/i 的概率选中（替换为答案）
			ans = node.Val
		}
		i++
	}
	return
}
