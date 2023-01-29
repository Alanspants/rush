package main

import "fmt"

var b = 200

func outer() {
	c := 99
	var inner = func() {
		c = 100
		fmt.Println("inner: ", c)
	}
	inner()
	fmt.Println(c)
}

func main() {
	outer()
}
