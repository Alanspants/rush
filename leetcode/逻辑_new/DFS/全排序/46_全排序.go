package main

import "fmt"

func premute_1(nums []int) [][]int {
	res := [][]int{}
	visited := map[int]bool{}
	dfs(nums, []int{}, &res, visited)
	return res
}

func dfs_1(nums []int, path []int, ans *[][]int, visited map[int]bool) {
	if len(path) == len(nums) {
		temp := make([]int, len(path))
		copy(temp, path)
		*ans = append(*ans, temp)
		return
	}
	for _, v := range nums {
		if visited[v] {
			continue
		}
		path = append(path, v)
		visited[v] = true
		dfs(nums, path, ans, visited)
		path = path[:len(path)-1]
		visited[v] = false
	}
}

func main() {
	nums := []int{1, 2, 3}
	fmt.Print(premute(nums))

}
