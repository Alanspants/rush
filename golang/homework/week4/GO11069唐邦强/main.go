package main

import (
	"fmt"
	"math/rand"
	"sort"
)

func main() {
	/*
			1.1、看下面一段程序，回答问题
			s1 := make([]int, 3, 4)
			s2 := append(s1, 1)
			请问s1、s2内各有什么元素？ s1[0 0 0] s2[0 0 0 1]
			s1修改一个元素会影响s2吗？s2修改一个元素会影响s1吗？  都会
			s2再增加一个元素会怎么样？    切片扩容, 不再共用底层数组

		s1 := make([]int, 3, 4)
		s2 := append(s1, 1)
		fmt.Printf("s1:%v %p %p len:%d cap:%d\n", s1, &s1, &s1[0], len(s1), cap(s1))
		fmt.Printf("s2:%v %p %p len:%d cap:%d\n", s2, &s2, &s2[0], len(s2), cap(s2))
		s1[0] = 2
		fmt.Printf("s1:%v %p %p len:%d cap:%d\n", s1, &s1, &s1[0], len(s1), cap(s1))
		s2[2] = 5
		fmt.Printf("s2:%v %p %p len:%d cap:%d\n", s2, &s2, &s2[0], len(s2), cap(s2))
		s2 = append(s2, 6)
		fmt.Printf("s2:%v %p %p len:%d cap:%d\n", s2, &s2, &s2[0], len(s2), cap(s2))
	*/
	arr := [8]int{1, 4, 9, 16, 2, 5, 10, 15}
	arrsum := make([]int, len(arr)-1)
	for i := 1; i < 8; i++ {
		arrsum[i-1] = arr[i] + arr[i-1]

	}
	fmt.Printf("arrsum: %v\n", arrsum)
	arr1 := make([]int, 0)
	count := make(map[int]int)
	for j := 0; j < 100; j++ {
		r := rand.Intn(400) - 200
		count[r]++
		arr1 = append(arr1, r)
	}
	sort.Ints(arr1)
	fmt.Printf("arr1: %v\n", arr1)
	for k, v := range count {
		fmt.Printf("munber: %v count: %v\n", k, v)
	}
}

// 批改意见
// 1. Intn的范围是[0, n)，因此Intn(400)-200会漏掉最后一个数字200
// 2. 代码中记得写注释，方便自己查看
