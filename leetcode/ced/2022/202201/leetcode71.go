package _02201

import "strings"

func simplifyPath(path string) string {
	stack := []string{}
	for _, name := range strings.Split(path, "/") {
		if name == ".." { // 注意多if else的时候要注意判断条件
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			}
		} else if name != "" && name != "." {
			stack = append(stack, name)
		}
	}
	return "/" + strings.Join(stack, "/")
}
