package main

import "fmt"

func main() {
	for i := 1; i < 10; i++ {
		for j := 1; j <= i; j++ {
			if j == 1 {
				fmt.Printf("%v*%v=%-2v", j, i, j*i)
			} else {
				fmt.Printf("%v*%v=%-3v", j, i, j*i)
			}
		}
		fmt.Println()
	}
}
