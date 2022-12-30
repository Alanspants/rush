package main

type board struct {
	row    int
	column int
	grid   [][]int
}

func (board board) getAllNextStep(isNextPlayer bool) [][]int {

	result := [][]int{}

	if isNextPlayer {
		// white, 2, from right corner
		for i := 0; i < board.row; i++ {
			for j := 0; j < board.column; j++ {
				if board.grid[i][j] == 2 {
					// path-1, left-up
					if i-1 >= 0 && j-1 >= 0 && board.grid[i-1][j-1] == 0 {
						result = append(result, []int{i - 1, j - 1})
					}
					// path-2, right-up
					if i-1 >= 0 && j+1 < board.column && board.grid[i-1][j+1] == 0 {
						result = append(result, []int{i - 1, j + 1})
					}
				}
			}
		}
	} else {
		// black, 1, from left corner, only going down
		for i := 0; i < board.row; i++ {
			for j := 0; j < board.column; j++ {
				if board.grid[i][j] == 1 {
					// path-1, left-down
					if i+1 < board.row && j-1 >= 0 && board.grid[i+1][j-1] == 0 {
						result = append(result, []int{i + 1, j - 1})
					}
					// path-2, right-up
					if i+1 < board.row && j+1 < board.column && board.grid[i+1][j+1] == 0 {
						result = append(result, []int{i + 1, j + 1})
					}
				}
			}
		}
	}
	return result
}

func canMove(current []int, target []int, isNextPlayer bool) bool {
	/*
		  0 1 2 3 4 5
		0 1 0 0 0 0 0
		1 0 1 0 0 0 0
		2 0 0 1 0 0 0
		3 0 0 0 2 0 0
		4 0 0 0 0 2 0
		5 0 0 0 0 0 2
	*/
	if isNextPlayer {
		// white, 1, right-down corner
		rowGap := current[0] - target[0]
		// same row or downside row [false]
		if rowGap <= 0 {
			return false
		}
		// gap 1, 5, 7 ... N row
		if rowGap%2 == 0 {
			columnGap := current[1] - target[1]
			if columnGap%2 != 0 {
				return false
			}
		} else if rowGap%2 != 0 {
			columnGap := current[1] - target[1]
			if columnGap%2 == 0 {
				return false
			}
		}
	}
	return true
}
