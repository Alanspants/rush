package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func main() {
	/*
		   1、看下面一段程序，回答问题
		   s1 := make([]int, 3, 4)
		   s2 := append(s1, 1)
		   请问s1、s2内各有什么元素？
		   s1: [0 0 0]  s2: [0 0 0 1]
		   s1修改一个元素会影响s2吗？s2修改一个元素会影响s1吗？
		   s1,s2共用一个底层数组，会互相影响
		    s2再增加一个元素会怎么样？
			s2会扩容len(5) cap(8)，与s1不再共用底层数组
	*/

	// 2、有一个数组 [1,4,9,16,2,5,10,15]，生成一个新切片，要求新切片元素是数组相邻2项的和。
	arr := [...]int{1, 4, 9, 16, 2, 5, 10, 15}
	Newslice := make([]int, 0, len(arr))
	// for i := 0; i < len(arr)-1; i++ {
	// 	Newslice = append(Newslice, arr[i]+arr[i+1])
	// }
	for i, v := range arr {
		if i < len(arr)-1 {
			Newslice = append(Newslice, v+arr[i+1])
		}
	}
	fmt.Println(Newslice)

	/*
	   3、数字重复统计
	   随机产生100个整数
	   数字的范围[-200, 200]
	   升序输出这些生成的数字并打印其重复的次数
	*/
	s := make([]int, 0, 100)
	src := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < 100; i++ {
		r100 := src.Intn(401) - 200
		s = append(s, r100)
	}
	sort.Ints(s)
	fmt.Println(s)
	m := make(map[int]int)
	for _, v := range s {
		if _, ok := m[v]; ok {
			m[v] += 1
		} else {
			m[v] = 1
		}
	}
	for i, v := range m {
		fmt.Printf("数字[%d]出现了%d次\n", i, v)
	}

}

// 批改意见
// 1. s2虽然和s1共用一个底层数组，但是s2长度比s1大，因此s2修改最后一个元素不影响s1
// 2. 输出的时候记得排序
