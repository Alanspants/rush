package main

import (
	"fmt"
	"sort"
)

func dfs(nums []int, path []int, res *[][]int, index int) {
	temp := make([]int, len(path))
	copy(temp, path)
	*res = append(*res, temp)

	for i := index; i < len(nums); i++ {
		if i > index && nums[i] == nums[i-1] {
			continue
		}

		path = append(path, nums[i])
		dfs(nums, path, res, i+1)
		path = path[:len(path)-1]
	}
}

func subsets(nums []int) [][]int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] > nums[j]
	})
	res := [][]int{}
	dfs(nums, []int{}, &res, 0)
	return res
}

func main() {
	fmt.Println(subsets([]int{1, 2, 3}))
}
