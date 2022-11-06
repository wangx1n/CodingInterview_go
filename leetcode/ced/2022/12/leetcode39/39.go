package leetcode39

func combinationSum(candidates []int, target int) [][]int {
	var ans [][]int
	var t_ans []int
	dfs(candidates, t_ans, &ans, target, 0)
	return ans
}

func dfs(candidates, t_ans []int, ans *[][]int, target, idx int) {
	if idx == len(candidates) {
		return
	}
	if target == 0 {
		*ans = append(*ans, append([]int{}, t_ans...))
		return
	}

	dfs(candidates, t_ans, ans, target, idx+1)

	if target-candidates[idx] >= 0 {
		t_ans = append(t_ans, candidates[idx])
		dfs(candidates, t_ans, ans, target-candidates[idx], idx)
		t_ans = t_ans[:len(t_ans)-1]
	}
}
