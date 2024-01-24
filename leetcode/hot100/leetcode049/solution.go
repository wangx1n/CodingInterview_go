package leetcode049

import "sort"

func groupAnagrams(strs []string) [][]string {
	res := make(map[string][]string)

	for _, str := range strs {
		s := []byte(str)
		sort.Slice(s, func(i, j int) bool {
			return s[i] < s[j]
		})
		res[string(s)] = append(res[string(s)], str)
	}

	ans := make([][]string, 0)
	for _, v := range res {
		ans = append(ans, v)
	}
	return ans
}
