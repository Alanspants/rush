package main

import "fmt"

var i int = 0
var j int = 1

func fibonaci() {
	fmt.Println(i)
	fmt.Println(j)
	for {
		i = j
		j = i + j
		if j > 100 {
			break
		} else {
			fmt.Println(i)
		}
	}
	fmt.Println(i)
}

func main() {
	fibonaci()
}