package main

import "fmt"

func main() {
	arr := [...]int{1, 4, 9, 16, 2, 5, 10, 15}
	count := len(arr) - 1
	out := make([]int, 0, count)
	for i := 0; i < count; i++ {
		out = append(out, arr[i]+arr[i+1]) // 这里需要判断 i+1 有没有超界
	}
	fmt.Println(out)
}
