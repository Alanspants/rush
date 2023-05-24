package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func homework1() {
	// 看下面一段程序，回答问题
	// s1 := make([]int, 3, 4)
	// s2 := append(s1, 1)

	// 请问s1、s2内各有什么元素？
	/*
		s1[0 0 0]   len3 cap4
		s2[0 0 0 1] len4 cap4
	*/

	// s1修改一个元素会影响s2吗？s2修改一个元素会影响s1吗？
	/*
		都会影响，因为共用一个底层数组。
	*/

	// s2再增加一个元素会怎么样？
	/*
		s2会扩容，会新建一个新的底层数组，此时len5，cap8
	*/
}

func homework2() {
	// 有一个数组 [1,4,9,16,2,5,10,15]，生成一个新切片，要求新切片元素是数组相邻2项的和。
	array := [...]int{1, 4, 9, 16, 2, 5, 10, 15}
	s := make([]int, 0, len(array))
	for i := 0; i < len(array)-1; i++ {
		s = append(s, array[i]+array[i+1])
	}
	fmt.Println(s)
}

func homework3() {
	// 数字重复统计
	// 随机产生100个整数
	// 数字的范围[-200, 200]
	// 升序输出这些生成的数字并打印其重复的次数

	r := rand.New(rand.NewSource(time.Now().UnixNano())) // 随机数
	s := make([]int, 0, 100)                             // 存放随机数的切片
	m := make(map[int]int, 100)                          // 统计次数的map

	for i := 0; i < 100; i++ { // 100个随机数放进去
		randow := r.Intn(401) - 200 // [0,401) - 200  ->  [-200,201)  == [-200,200]
		s = append(s, randow)
	}

	// fmt.Println(s)
	sort.Ints(s) // 默认升序排
	fmt.Println(s)

	for _, v := range s {
		m[v] += 1 // 切片中随机数当key，有就加一个
	}
	fmt.Println()
	fmt.Println(len(m), m)
}

func main() {
	homework1()
	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
	homework2()
	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
	homework3()
}

// 批改意见
// 1. s2修改最后添加后的元素不会影响s1
