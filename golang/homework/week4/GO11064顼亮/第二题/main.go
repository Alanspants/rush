package main

import (
	"fmt"
)

// 3、有一个数组 [1,4,9,16,2,5,10,15]，生成一个新切片，要求新切片元素是数组相邻2项的和。

func main() {
	var arr = [8]int{1, 4, 9, 16, 2, 5, 10, 15}
	fmt.Println("数组: ", arr)
	// [1,2,4,5,9,10,15,16]
	// [9, 10, 15, 16]
	// 4+5=9, 1+9=10, 5+10=15, 1+15=16
	var s []int
	for i := 0; i < len(arr)-1; i++ {
		s = append(s, arr[i]+arr[i+1])
	}
	fmt.Println("切片: ", s)
}
