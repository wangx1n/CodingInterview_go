package leetcode017

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return nil
	}
	letters := map[byte][]byte{
		'2': []byte{'a', 'b', 'c'},
		'3': []byte{'d', 'e', 'f'},
		'4': []byte{'g', 'h', 'i'},
		'5': []byte{'j', 'k', 'l'},
		'6': []byte{'m', 'n', 'o'},
		'7': []byte{'q', 'p', 'r', 's'},
		'8': []byte{'t', 'u', 'v'},
		'9': []byte{'w', 'x', 'y', 'z'},
	}
	var ret []string
	dfs(&ret, letters, 0, "", digits)
	return ret
}

func dfs(ret *[]string, letters map[byte][]byte, index int, ans string, digits string) {
	if len(ans) == len(digits) {
		*ret = append(*ret, ans)
		return
	}
	digit := digits[index]
	letter := letters[digit]
	for i := 0; i < len(letter); i++ {
		dfs(ret, letters, index+1, ans+string(letter[i]), digits)
	}
}
