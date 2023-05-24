package main

import "fmt"

func main() {
	// 请问s1、s2内各有什么元素？
	s1 := make([]int, 3, 4)
	s2 := append(s1, 1)

	// - make 创建一个切片 s1，长度为 3， 容量为 4，底层数组起始地址
	// - append 底层数组中追加，增加一个元素不用扩容，返回一个 header 给了 s2
	// - 目前公用底层数组
	// - s1[0 0 0]
	// - s2[0 0 0 1]
	fmt.Printf("s1: %p %p %v\ns2: %p %p %v\n", &s1, &s1[0], s1, &s2, &s2[0], s2)
	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")

	// s1修改一个元素会影响s2吗？s2修改一个元素会影响s1吗？
	// - s1 修改会影响 s2
	s1[2] = 8
	fmt.Printf("s1: %p %p %v\ns2: %p %p %v\n", &s1, &s1[0], s1, &s2, &s2[0], s2)
	// - s2 修改不一定影响 s1（本质上影响）
	s2[3] = 9
	fmt.Printf("s1: %p %p %v\ns2: %p %p %v\n", &s1, &s1[0], s1, &s2, &s2[0], s2)
	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")

	// s2再增加一个元素会怎么样？
	// - s2 扩容，分配一个全新的底层数组（分道扬镳）
	s2 = append(s2, 6)
	fmt.Printf("s1: %p %p %v\ns2: %p %p %v\n", &s1, &s1[0], s1, &s2, &s2[0], s2)
}
