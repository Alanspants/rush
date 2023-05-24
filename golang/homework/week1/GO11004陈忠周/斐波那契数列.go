package main

import "fmt"

// 打印100以内的斐波那契数列
func main() {

	// goto 方式
	// 	i, j, v := 0, 1, 1
	// vart:
	// 	fmt.Println(v)
	// 	v = i + j
	// 	i, j = j, v
	// 	for {
	// 		if v >= 100 {
	// 			break
	// 		}
	// 		goto vart
	// 	}

	// for 方式
	i, j, v := 0, 0, 1
	for ; ; v = i + j {
		if v >= 100 {
			break
		}
		fmt.Println(v)
		i, j = j, v
	}
}
