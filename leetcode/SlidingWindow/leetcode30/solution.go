package leetcode30

func findSubstring(s string, words []string) []int {
	ls, wordLen, wordsLen := len(s), len(words[0]), len(words)
	ans := []int{}

	for i := 0; i < wordLen && i+wordsLen*wordLen <= ls; i++ {
		diff := make(map[string]int, wordsLen)

		for j := 0; j < wordsLen; j++ {
			diff[s[i+j*wordLen:i+(j+1)*wordLen]]++
		}

		for _, word := range words {
			diff[word]--
			if diff[word] == 0 {
				delete(diff, word)
			}
		}

		// 这里是每次移动wordLen个字符，相比于i挨个向右移动，要更快
		for k := i; k < ls-wordLen*wordsLen+1; k += wordLen {
			if k != i {
				word := s[k+wordLen*(wordsLen-1) : k+wordLen*wordsLen]
				diff[word]++
				if diff[word] == 0 {
					delete(diff, word)
				}
				word = s[k-wordLen : k]
				diff[word]--
				if diff[word] == 0 {
					delete(diff, word)
				}
			}
			if len(diff) == 0 {
				ans = append(ans, k)
			}

		}
	}
	return ans
}
