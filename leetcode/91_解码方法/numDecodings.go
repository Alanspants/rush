package main

import "fmt"

/*
	dp[len(s)]: s的解法个数

	a, b

	[case1]
	a = 0, b != 0
	dp[i] = dp[i-1]

	[case2]
	a = 0, b = 0
	return 0

	[case3]
	a != 0, b != 0
	dp[i] = dp[i - 1] + dp[i - 2]

	[case4]
	a != 0, b = 0
	dp[i] = dp[i-1]
*/

func numDecodings(s string) int {
	dp := make([]int, len(s)+1)
	dp[0] = 1

	if s[0] == '0' {
		return 0
	}

	s = " " + s
	for i := 1; i <= len(s)-1; i++ {
		current := s[i] - '0'
		combo := (s[i-1]-'0')*10 + current

		if current != 0 {
			dp[i] = dp[i-1]
		}

		if i > 1 && combo >= 10 && combo <= 26 {
			dp[i] += dp[i-2]
		}
	}
	return dp[len(s)-1]
}

func main() {
	fmt.Println(numDecodings("12"))
}
