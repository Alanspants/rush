package main

import (
	"fmt"
)

func printTriangle(n int) {
	last := ""
	for i := n; i >= 1; i-- {
		if i == 1 {
			last += fmt.Sprint(i)
		} else {
			last += fmt.Sprintf("%d ", i)
		}
	}

	width := len(last)
	for i := width - 1; i >= 1; i-- {
		if last[i] == 32 {
			fmt.Printf("%*s\n", width, last[i+1:])
		}
	}
	fmt.Print(last)
}

func main() {
	printTriangle(12)
}
