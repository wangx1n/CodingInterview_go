package _9

func strStr(haystack string, needle string) int {
	isVaild := false
	for i := 0; i <= len(haystack)-len(needle); i++ {
		for j, k := i, 0; k < len(needle); j++ {
			if haystack[j] != needle[k] {
				break
			}
			if k == len(needle)-1 {
				isVaild = !isVaild
			}
			k++
		}
		if isVaild {
			return i
		}
	}
	return -1
}
