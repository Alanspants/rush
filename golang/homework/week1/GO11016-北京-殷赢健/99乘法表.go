package main

import "fmt"

func main() {
	for i := 1; i <= 9; i++ {
		for j := 1; j <= 9; j++ {
			fmt.Printf("%dx%d=%-2d  ", j, i, i*j)
			if i == j {
				break
			}
		}
		fmt.Println("")
	}

}
