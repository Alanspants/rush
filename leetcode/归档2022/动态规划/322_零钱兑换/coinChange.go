package main

import (
	"fmt"
	"math"
)

/*
n: amount
dp[n] = min(dp[n - x0], dp[n - x1] .. dp[n - xn-1]) + 1

dp[0] = 0

*/

func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	dp[0] = 0

	for i := 1; i <= amount; i++ {
		dp[i] = math.MaxInt
		for j := 0; j < len(coins); j++ {
			if i < coins[j] || dp[i-coins[j]] == math.MaxInt {
				continue
			}
			if dp[i] > dp[i-coins[j]] {
				dp[i] = dp[i-coins[j]]
			}
		}
		if dp[i] != math.MaxInt {
			dp[i]++
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
