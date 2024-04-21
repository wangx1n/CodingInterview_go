package _4

import "strings"

func reverseWords(s string) string {
	ss := strings.Split(s, " ")
	withoutSpace := make([]string, 0, len(ss))

	for _, s := range ss {
		if s == string(" ") {
			continue
		}
		withoutSpace = append(withoutSpace, s)
	}

	group := reverseGroup(withoutSpace)
	ret := ""

	for _, v := range group {
		ret += v
		ret += string(" ")
	}
	return ret[:len(ret)-1]
}

func reverseGroup(group []string) []string {
	for i := 0; i < len(group)/2; i++ {
		group[i], group[len(group)-i-1] = group[len(group)-i-1], group[i]
	}
	return group
}
