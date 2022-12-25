package main

import "fmt"

//func dfs(currentM int, currentN int, m int, n int) int {
//	if currentM == (m-1) && currentN == (n-1) {
//		return 1
//	}
//
//	pathNum := 0
//
//	if currentM < (m - 1) {
//		pathNum += dfs(currentM+1, currentN, m, n)
//	}
//	if currentN < (n - 1) {
//		pathNum += dfs(currentM, currentN+1, m, n)
//	}
//
//	return pathNum
//}
//
//func uniquePaths(m int, n int) int {
//	return dfs(0, 0, m, n)
//}

func uniquePaths(m int, n int) int {
	dp := make([][]int, m)

	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
		dp[i][0] = 1
	}

	for j := 0; j < n; j++ {
		dp[0][j] = 1
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}

	return dp[m-1][n-1]
}

func main() {
	fmt.Println(uniquePaths(3, 7))
}
