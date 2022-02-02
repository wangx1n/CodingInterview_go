package _02112

import (
	"sort"
)

// trie 字典树
type trie struct {
	child [26]*trie
	isEnd bool
}

func (root *trie) dfs(word string) bool {
	if word == "" {
		return true
	}
	node := root
	for i, ch := range word {
		node = node.child[ch-'a']
		if node == nil {
			return false
		}
		if node.isEnd && root.dfs(word[i+1:]) {
			return true
		}
	}
	return false
}

func (root *trie) insert(sentense string) {
	node := root
	for _, ch := range sentense {
		ch -= 'a'
		if node.child[ch] == nil {
			node.child[ch] = &trie{}
		}
		node = node.child[ch]
	}
	node.isEnd = true
}

func findAllConcatenatedWordsInADict(words []string) []string {
	sort.Slice(words, func(i, j int) bool { return len(words[i]) < len(words[j]) })
	root := &trie{}
	ans := []string{}
	for _, word := range words {
		if word == "" {
			continue
		}
		if root.dfs(word) {
			ans = append(ans, word)
		} else {
			root.insert(word)
		}
	}
	return ans
}
