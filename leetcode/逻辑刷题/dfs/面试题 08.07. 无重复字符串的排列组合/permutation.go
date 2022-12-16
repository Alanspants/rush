package main

import "fmt"

func dfs(res *[]string, visited map[int]bool, path string, target string) {
	if len(path) == len(target) {
		*res = append(*res, path)
		return
	}

	for _, v := range target {
		if visited[int(v)] {
			continue
		}

		visited[int(v)] = true
		path = path + string(v)
		dfs(res, visited, path, target)
		visited[int(v)] = false
		path = path[:len(path)-1]
	}
}

func permutation(S string) []string {
	var res []string
	visited := map[int]bool{}
	dfs(&res, visited, "", S)
	return res
}

func main() {
	var s = "qwe"
	fmt.Println(permutation(s))
	//for _, v := range s {
	//	fmt.Println(v)
	//}
}
