package main

import "fmt"

func sum(nums []int) int {
	sum := 0
	for _, v := range nums {
		sum += v
	}
	return sum
}

func dfs(candidates []int, target int, res *[][]int, path []int, index int) {
	if sum(path) == target {
		temp := make([]int, len(path))
		copy(temp, path)
		*res = append(*res, temp)
	}

	//for _, v := range candidates {
	//	if sum(path) > target {
	//		break
	//	}
	//	path = append(path, v)
	//	dfs(candidates, target, res, path)
	//	path = path[:len(path)-1]
	//}

	// 保证下一次递归是从比自己大的数字开始，避免了重复的情况
	for i := index; i < len(candidates); i++ {
		if sum(path) > target {
			break
		}
		path = append(path, candidates[i])
		dfs(candidates, target, res, path, i)
		path = path[:len(path)-1]
	}
}

func combinationSum(candidates []int, target int) [][]int {
	res := [][]int{}
	dfs(candidates, target, &res, []int{}, 0)
	return res
}

func main() {
	fmt.Println(combinationSum([]int{2, 3, 6, 7}, 7))
}
