package main

import (
	"fmt"
)

func main() {
	init_list := [3]int{1, 1, 2}
	mem_list := []int{}
	index := 0
	for i, v := range init_list {
		// fmt.Println(i)
		mem_list = append(mem_list, v)
		index = i
	}

	sum := init_list[index] + init_list[index-1]

	mem_list = append(mem_list, sum)

	// fmt.Printf("mem_list: %v\n", mem_list)

	// 这个地方的逻辑有问题，要求值是100以内的，不是输出100个以内的。
	for i := 5; i < 100; i++ {
		fmt.Printf("i: %v\n", i)
		for i, _ := range mem_list {
			index = i
		}
		// fmt.Printf("index: %v\n", index)
		sum = mem_list[index] + mem_list[index-1]
		mem_list = append(mem_list, sum)

	}
	fmt.Printf("mem_list: %v\n", mem_list)
}
