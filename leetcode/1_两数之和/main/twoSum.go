package main

import "fmt"

func twoSum(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		if nums[i] >= target {
			continue
		}
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				ans := []int{i, j}
				return ans
			}
		}
	}
	return []int{}
}

func main() {
	var n int
	fmt.Print("数组长度： ")
	fmt.Scan(&n)

	var input []int
	fmt.Print("数组内容： ")
	for i := 0; i < n; i++ {
		var temp int
		fmt.Scan(&temp)
		input = append(input, temp)
	}

	var target int
	fmt.Print("目标： ")
	fmt.Scan(&target)

	fmt.Println(twoSum(input, target))
}
