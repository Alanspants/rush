package main

import "fmt"

func dfs(board [][]byte, word string, visited *[][]bool, x int, y int, index int) bool {

	if index == len(word) {
		return true
	}
	if (*visited)[x][y] {
		return false
	}
	if board[x][y] == word[index] {
		if index == (len(word) - 1) {
			return true
		}
		(*visited)[x][y] = true
		if (x+1) < len(board) && board[x+1][y] == word[index+1] {
			if dfs(board, word, visited, x+1, y, index+1) {
				return true
			}
		}
		if (x-1) >= 0 && board[x-1][y] == word[index+1] {
			if dfs(board, word, visited, x-1, y, index+1) {
				return true
			}
		}
		if (y+1) < len(board[x]) && board[x][y+1] == word[index+1] {
			fmt.Println("向上")
			if dfs(board, word, visited, x, y+1, index+1) {
				return true
			}
		}
		fmt.Println("向下")
		if (y-1) >= 0 && board[x][y-1] == word[index+1] {
			if dfs(board, word, visited, x, y-1, index+1) {
				return true
			}
		}
		(*visited)[x][y] = false
	} else {
		return false
	}

	return false
}

func exist(board [][]byte, word string) bool {
	visited := make([][]bool, len(board))
	for i := range board {
		visited[i] = make([]bool, len(board[i]))
	}

	for i := range board {
		for j := range board[i] {
			if board[i][j] == word[0] {
				fmt.Println("--------")
				if dfs(board, word, &visited, i, j, 0) {
					return true
				}
			}
		}
	}
	return false
	//TODO：先找到开头
}

func main() {
	board := [][]byte{
		{'A', 'B', 'C', 'E'},
		{'S', 'F', 'C', 'S'},
		{'A', 'D', 'E', 'E'},
	}
	fmt.Println(exist(board, "SEE"))
}
