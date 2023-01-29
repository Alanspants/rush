package main

import "fmt"

func main() {
	a := make([]int, 5, 9)
	fmt.Println(a)
	fmt.Println(a[4])
	a = a[1:7] //0 ~ 7
	fmt.Println(a)
	fmt.Println(len(a), cap(a))

	//{0, 0, 0, 0, 0, null, null, null,}
}
