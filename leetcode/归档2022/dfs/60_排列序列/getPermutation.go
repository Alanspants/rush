package main

import (
	"fmt"
	"strconv"
)

func dfs(nums []int, visited map[int]bool, path string, n int, k int, currentIndex int, fac []int) string {

	if currentIndex == n {
		return path
	}

	for _, v := range nums {
		fmt.Printf("本子树count： %d\n", fac[len(nums)-currentIndex-1])
		if visited[v] {
			fmt.Println("visited skip")
			continue
		}
		if fac[n-currentIndex-1] < k {
			fmt.Println("cnt skip")
			k -= fac[n-currentIndex-1]
			continue
		}
		path += strconv.Itoa(v)
		visited[v] = true
		fmt.Printf("v: %d\n", v)
		fmt.Printf("path: %s\n", path)
		path = dfs(nums, visited, path, n, k, currentIndex+1, fac)
	}

	return path
}

func getPermutation(n int, k int) string {
	visited := map[int]bool{}
	nums := make([]int, n)
	for i := 1; i <= n; i++ {
		nums[i-1] = i
	}
	fmt.Println(factorial(n))
	fmt.Println(nums)
	res := dfs(nums, visited, "", n, k, 0, factorial(n))
	return res
}

func factorial(n int) []int {
	facRes := make([]int, n+1)
	for i := 0; i <= n; i++ {
		facTemp := 1
		for j := i; j > 0; j-- {
			if i == 0 {
				facTemp = 1
				break
			}
			facTemp = facTemp * j
		}
		facRes[i] = facTemp
	}
	return facRes
}

func main() {
	n := 3
	k := 2
	res := getPermutation(n, k)
	fmt.Println(res)
}
