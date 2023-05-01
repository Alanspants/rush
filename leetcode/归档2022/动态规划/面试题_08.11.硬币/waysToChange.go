package main

import "fmt"

func waysToChange(n int) int {
	coins := []int{1, 5, 10, 25}
	dp := make([]int, n+1)
	dp[0] = 1

	for _, coin := range coins {
		for i := coin; i <= n; i++ {
			dp[i] += dp[i-coin]
		}
	}

	return dp[n]
}

func main() {
	fmt.Println(waysToChange(waysToChange(10)))
}
