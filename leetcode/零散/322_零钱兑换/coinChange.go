package main

import (
	"fmt"
	"math"
)

/*
coins[x, y, z]
dp[i] = min(dp[i - x], dp[i - y], dp[i - z]) + 1

dp[0] = 0

*/

func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	dp[0] = 0

	for i := 1; i <= amount; i++ {
		dp[i] = math.MaxInt
		for j := 0; j < len(coins); j++ {
			if i >= coins[j] && dp[i-coins[j]] != math.MaxInt {
				if dp[i] >= dp[i-coins[j]]+1 {
					dp[i] = dp[i-coins[j]] + 1
				}
			}
		}
	}

	if dp[amount] == math.MaxInt {
		dp[amount] = -1
	}

	return dp[amount]
}

func main() {
	fmt.Println(coinChange([]int{1, 2, 5}, 11))
}
