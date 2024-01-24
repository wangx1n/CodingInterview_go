package leetcode128

func longestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	m := make(map[int]bool)
	for _, v := range nums {
		m[v] = true
	}

	res := 1

	for num := range m {
		if !m[num-1] {
			temRes := 1
			currentNum := num + 1
			for m[currentNum] {
				currentNum = currentNum + 1
				temRes++
			}
			if temRes > res {
				res = temRes
			}
		}
	}
	return res
}
