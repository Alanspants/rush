package main

import "fmt"

func premute(nums []int) [][]int {
	res := [][]int{}
	visited := map[int]bool{}
	dfs(nums, []int{}, &res, visited)
	return res
}

func dfs(nums []int, path []int, res *[][]int, visited map[int]bool) {
	if len(path) == len(nums) {
		temp := make([]int, len(path))
		copy(temp, path)
		*res = append(*res, temp)
		return
	}
	for _, v := range nums {
		if visited[v] {
			continue
		}
		visited[v] = true
		path = append(path, v)
		dfs(nums, path, res, visited)
		visited[v] = false
		path = path[:len(path)-1]
	}
}

func main() {
	input := []int{1, 2, 3}
	fmt.Print(premute(input))
}
