package main

import "fmt"

/*

dp[amount] = dp[amount - coin0] + .. dp[amount - coinn-1]

*/

func change(amount int, coins []int) int {
	dp := make([]int, amount+1)
	dp[0] = 1

	for i := 1; i <= amount; i++ {
		dp[i] = 0
		for j := 0; j < len(coins); j++ {
			if i-coins[j] < 0 {
				continue
			}
			dp[i] += dp[i-coins[j]]
		}
	}
	return dp[amount]
}

func main() {
	fmt.Println(change(5, []int{1, 2, 5}))
}
