package main

import "fmt"

func main() {
	var product string

	for i := 1; i < 10; i++ {
		for j := 1; j <= i; j++ {
			if i*j < 10 {
				product = fmt.Sprintf("%v*%v=%v  ", j, i, i*j)
			} else {
				product = fmt.Sprintf("%v*%v=%v ", j, i, i*j)
			}
			fmt.Print(product)

		}
		fmt.Printf("\n")
	}

}
