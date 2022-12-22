package main

import (
	"fmt"
	"sort"
)

func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	ans := [][]int{}
	for _, interval := range intervals {
		if len(ans) == 0 {
			ans = append(ans, interval)
			continue
		}

		flag := false
		for i := range ans {
			if ans[i][0] <= interval[0] && ans[i][1] >= interval[0] && interval[1] > ans[i][1] {

				/*
					ans[i][0] = [1, 5]
					interval = [3, 7]
					-> [1, 7]
				*/
				ans[i][1] = interval[1]
				flag = true
				break
			} else if ans[i][0] >= interval[0] && ans[i][0] <= interval[1] && ans[i][1] >= interval[1] {
				/*
					ans[i][0] = [3, 7]
					interval = [1, 5]
					-> [1, 7]
				*/
				ans[i][0] = interval[0]
				flag = true
				break
			} else if ans[i][0] >= interval[0] && ans[i][1] <= interval[1] {
				/*
					ans[i][0] = [3, 5]
					interval = [1, 7]
					-> [1, 7]
				*/
				ans[i][0] = interval[0]
				ans[i][1] = interval[1]
				flag = true
				break
			} else if ans[i][0] <= interval[0] && ans[i][1] >= interval[1] {
				/*
					ans[i][0] = [1, 7]
					interval = [3, 5]
					-> [1, 7]
				*/
				flag = true
				break
			}
		}
		if flag == false {
			ans = append(ans, interval)
		}
	}
	return ans
}

func main() {
	//intervals := [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}
	intervals := [][]int{{1, 4}, {2, 3}}
	fmt.Println(merge(intervals))
}
