package main

import "fmt"

func dfs(res *[]string, visited map[int]bool, path string, target string) {
	if len(path) == len(target) {
		*res = append(*res, path)
		return
	}

	for i, v := range target {
		if visited[i] {
			continue
		}

		if i > 0 && !visited[i-1] && target[i] == target[i-1] {
			continue
		}

		visited[i] = true
		path = path + string(v)
		dfs(res, visited, path, target)
		visited[i] = false
		path = path[:len(path)-1]
	}
}

func permutation(S string) []string {
	strMap := map[byte]int{}
	for _, v := range S {
		if value, exist := strMap[byte(v)]; exist {
			strMap[byte(v)] = value + 1
		} else {
			strMap[byte(v)] = 1
		}
	}
	finalStr := ""
	for key, value := range strMap {
		for i := 0; i < value; i++ {
			finalStr += string(key)
		}
	}

	var res []string
	visited := map[int]bool{}
	dfs(&res, visited, "", finalStr)

	return res
}

//func permutation(S string) []string {
//	var res []string
//	visited := map[int]bool{}
//	dfs(&res, visited, "", S)
//	return res
//}

func main() {
	str := "qwe"
	fmt.Println(permutation(str))

}

//func main() {
//	var s = "qwe"
//	fmt.Println(permutation(s))
//	//for _, v := range s {
//	//	fmt.Println(v)
//	//}
//}
