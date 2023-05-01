package main

import "fmt"

func dfs(nums []int, res *[][]int, visited map[int]bool, path []int) {
	if len(path) == len(nums) {
		temp := make([]int, len(path))
		copy(temp, path)
		*res = append(*res, temp)
		return
	}
	for i, v := range nums {
		if visited[i] {
			// 通过index判断
			continue
		}
		// 判断上一个数是否相同 + 是否visited过
		if i > 0 && nums[i-1] == v && !visited[i-1] {
			continue
		}
		path = append(path, v)
		visited[i] = true
		// 注意：在有重复的全排列中，需要用index做hashmap的key，因为有重复的数字
		dfs(nums, res, visited, path)
		visited[i] = false
		path = path[:len(path)-1]
	}
}

func permuteUnique(nums []int) [][]int {
	var res [][]int
	visited := map[int]bool{}
	dfs(nums, &res, visited, []int{})
	return res
}

func quickSort(nums []int, i int, j int) []int {
	if i >= j {
		return nums
	}

	pivot := nums[i]
	left := i
	right := j

	for i < j {
		for i < j && nums[j] >= pivot {
			j--
		}
		nums[i] = nums[j]

		for i < j && nums[i] <= pivot {
			i++
		}
		nums[j] = nums[i]
		//nums[i], nums[j] = nums[j], nums[i]
	}
	nums[i] = pivot
	quickSort(nums, left, i-1)
	quickSort(nums, j+1, right)
	return nums
}

func main() {
	nums := []int{3, 3, 0, 3}
	nums = quickSort(nums, 0, len(nums)-1)
	fmt.Println(nums)
	var ans [][]int
	ans = permuteUnique(nums)
	fmt.Println(ans)
}
