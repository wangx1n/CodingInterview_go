package graph

func findOrder(numCourses int, prerequisites [][]int) []int {
	var (
		edges = make([][]int, numCourses)
		visited = make([]int, numCourses)
		result []int // 这里不能写成make，因为make的[]int再append会在原来初始0后面append
		valid = true
		dfs func(u int)
	)
	dfs = func(u int) {
		visited[u] = 1
		for _, v := range edges[u] {
			if visited[v] == 0 {
				dfs(v)
				if !valid {
					return
				}
			} else if visited[v] == 1 {
				valid = false
				return
			}
		}
		visited[u] = 2
		result = append(result, u)
	}

	for _, info := range prerequisites {
		edges[info[1]] = append(edges[info[1]], info[0])
	}

	for i := 0; i < numCourses && valid; i++ {
		if visited[i] == 0 { // 这句如果不加的话会走过已经遍历过的点，即，只需要遍历起始节点就可以了，中间节点没有必要遍历
			dfs(i)
		}
	}
	if !valid {
		return []int{}
	}
	for i := 0; i < len(result)/2; i ++ { // dfs 最后需要逆序输出
		result[i], result[numCourses-i-1] = result[numCourses-i-1], result[i]
	}
	return result
}

func main() {
	findOrder(2, [][]int{{1,0}})
}

