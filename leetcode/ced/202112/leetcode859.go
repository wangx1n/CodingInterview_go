package _02112
/**
859. 亲密字符串
给你两个字符串 s 和 goal ，只要我们可以通过交换 s 中的两个字母得到与 goal 相等的结果，就返回 true ；否则返回 false 。

交换字母的定义是：取两个下标 i 和 j （下标从 0 开始）且满足 i != j ，接着交换 s[i] 和 s[j] 处的字符。

例如，在 "abcd" 中交换下标 0 和下标 2 的元素可以生成 "cbad" 。


示例 1：

输入：s = "ab", goal = "ba"
输出：true
解释：你可以交换 s[0] = 'a' 和 s[1] = 'b' 生成 "ba"，此时 s 和 goal 相等。
示例 2：

输入：s = "ab", goal = "ab"
输出：false
解释：你只能交换 s[0] = 'a' 和 s[1] = 'b' 生成 "ba"，此时 s 和 goal 不相等。
示例 3：

输入：s = "aa", goal = "aa"
输出：true
解释：你可以交换 s[0] = 'a' 和 s[1] = 'a' 生成 "aa"，此时 s 和 goal 相等。
示例 4：

输入：s = "aaaaaaabc", goal = "aaaaaaacb"
输出：true


提示：

1 <= s.length, goal.length <= 2 * 104
s 和 goal 由小写英文字母组成
 */
func buddyStrings(s string, goal string) bool {
	if len(s) != len(goal) {
		return false;
	}
	if s == goal {
		tmp := [26]int{}
		for i := 0; i < len(s); i++ {
			if tmp[s[i]-'a'] == 0 {
				tmp[s[i]-'a'] += 1
			} else {
				return true
			}
		}
		return false
	} else {
		first, second := -1, -1
		for i := 0; i < len(goal); i++ {
			if s[i] != goal[i] {
				if first == -1 {
					first = i
				} else if second == -1 {
					second = i
				} else {
					return false
				}
			}
		}
		if second != -1 && s[first] == goal[second] && s[second] == goal[first] {
			return true
		}
	}
	return false
}
