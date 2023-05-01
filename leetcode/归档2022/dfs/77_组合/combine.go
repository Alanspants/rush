package main

import "fmt"

func dfs(n int, k int, path []int, res *[][]int, index int) {
	if len(path) == k {
		temp := make([]int, len(path))
		copy(temp, path)
		*res = append(*res, temp)
	}

	for i := index; i <= n; i++ {
		path = append(path, i)
		dfs(n, k, path, res, i+1)
		path = path[:len(path)-1]
	}
}

func combine(n int, k int) [][]int {
	path := []int{}
	res := [][]int{}
	dfs(n, k, path, &res, 1)
	return res
}

func main() {
	fmt.Println(combine(4, 2))
}
