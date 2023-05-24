package main

import "fmt"

func main() {
	arr := [8]int{1, 4, 9, 16, 2, 5, 10, 15}
	sli := make([]int, 7, 100)

	for i := 0; i < len(arr)-1; i++ {
		sli[i] = arr[i] + arr[i+1]
	}
	fmt.Printf("arr: %v \t sli: %v\n", arr, sli)
}
