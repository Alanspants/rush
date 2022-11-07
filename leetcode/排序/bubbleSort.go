package main

import "fmt"

func bubbleSort(num []int) []int {
	isSorted := false
	for j := 0; j < len(num); j++ {
		isSorted = true
		for i := 0; i < len(num)-1; i++ {
			if num[i] > num[i+1] {
				num[i], num[i+1] = num[i+1], num[i]
				isSorted = false
			}
		}
		if isSorted {
			return num
		}
	}
	return num
}

func main() {
	unsortNum := []int{2, 4, 1, 5, 8, 2, 3, 8, 0, 9}
	unsortNum = bubbleSort(unsortNum)
	fmt.Println(unsortNum)
}
