package main

import "fmt"

func divisorGame(n int) bool {
	if n%2 == 0 {
		return true
	} else {
		return false
	}
}

func main() {
	var n int
	fmt.Scanf("n=%d\n", &n)
	fmt.Println(divisorGame(n))
}
