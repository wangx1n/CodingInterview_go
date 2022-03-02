package array

func twoSum(nums []int, target int) []int {
	m := map[int]int{}
	for i, x := range nums {
		if res, ok := m[target - x]; ok {
			return []int{res, i}
		}
		m[x] = i
	}
	return []int{}
}
