package main

import (
	"fmt"
	"sort"
)

func sum(nums []int) int {
	ans := 0
	for _, v := range nums {
		ans += v
	}
	return ans
}

func dfs(res *[][]int, path []int, target int, candidates []int, index int) {
	if sum(path) == target {
		temp := make([]int, len(path))
		copy(temp, path)
		*res = append(*res, temp)
	}
	for i := index; i < len(candidates); i++ {
		if sum(path) > target {
			continue
		}
		// 小剪枝：同一层相同数值的结点，从第 2 个开始，候选数更少，结果一定发生重复，因此跳过，用 continue
		if i > index && candidates[i] == candidates[i-1] {
			continue
		}

		path = append(path, candidates[i])
		dfs(res, path, target, candidates, i+1)
		path = path[:len(path)-1]
	}

}

func combinationSum2(candidates []int, target int) [][]int {
	res := [][]int{}
	sort.Slice(candidates, func(i, j int) bool {
		return candidates[i] < candidates[j]
	})
	dfs(&res, []int{}, target, candidates, 0)
	return res
}

func main() {
	fmt.Println(combinationSum2([]int{10, 1, 2, 7, 6, 1, 5}, 8))
}
