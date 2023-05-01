package main

import (
	"fmt"
	"strconv"
)

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	xStr := strconv.Itoa(x)
	for i, j := 0, len(xStr)-1; i <= j; i, j = i+1, j-1 {
		if xStr[i] != xStr[j] {
			return false
		}
	}
	return true
}

func main() {
	var x int
	fmt.Scan(&x)
	fmt.Println(isPalindrome(x))
}
