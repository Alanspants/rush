package main

import "fmt"

/*
25, 10, 5, 1

10=10
10=5+5
10=5+1+1+1+1+1
10=1+1+1+1+1+1+1+1+1+1

dp[i] = dp[i-25] + dp[i-10] + dp[i-5] + dp[i-1] i >= 25

dp[0] = 0




*/

func waysToChange(n int) int {
	dp := make([]int, n+1)
	dp[0] = 1

	for i := 1; i < n+1; i++ {
		if i >= 25 {
			dp[i] = dp[i-25] + dp[i-10] + dp[i-5] + dp[i-1]
		} else if i >= 10 {
			dp[i] = dp[i-10] + dp[i-5] + dp[i-1]
		} else if i >= 5 {
			dp[i] = dp[i-5] + dp[i-1]
		} else if i >= 1 {
			dp[i] = dp[i-1]
		}
		fmt.Println(i, ": ", dp[i])
	}

	return dp[n]
}

func main() {
	fmt.Println(waysToChange(10))
}
