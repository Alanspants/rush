package main

import "fmt"

//var jump = []struct {
//	i int
//	j int
//}{
//	{-2, -1}, {-2, 1}, {2, -1}, {2, 1}, {-1, -2}, {-1, 2}, {1, -2}, {1, 2}
//}

type Jump struct {
	x int
	y int
}

var jump = []Jump{{-2, -1}, {-2, 1}, {2, -1}, {2, 1}, {-1, -2}, {-1, 2}, {1, -2}, {1, 2}}

// dp[step][x][y]
func knightProbability(n int, k int, row int, column int) float64 {
	var dp = make([][][]float64, k+1)
	for step, _ := range dp {
		//fmt.Println(step)
		dp[step] = make([][]float64, n)
		for x, _ := range dp[step] {
			//fmt.Println(x)
			dp[step][x] = make([]float64, n)
			for y, _ := range dp[step][x] {
				if step == 0 {
					dp[step][x][y] = 1
				} else {
					for _, j := range jump {
						tempx := x + j.x
						tempy := y + j.y
						if tempx < n && tempx >= 0 && tempy < n && tempy >= 0 {
							dp[step][x][y] += dp[step-1][tempx][tempy] / 8
						}
					}
				}
			}
		}
	}
	return dp[k][row][column]

}

func main() {
	fmt.Printf("%.5f\n", knightProbability(3, 2, 0, 0))
}
