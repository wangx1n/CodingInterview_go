package _02112

import "sort"

func isNStraightHand(hand []int, groupSize int) bool {
	if len(hand) % groupSize != 0 {
		return false
	}
	cntMap := map[int]int{}
	sort.Ints(hand)
	for _, num := range hand {
		cntMap[num]++
	}
	for _, x := range hand {
		if cntMap[x] == 0 {
			continue
		}
		for i := x; i < x + groupSize; i++ {
			if cntMap[i] == 0 {
				return false
			}
			cntMap[i]--
		}
	}
	return true
}