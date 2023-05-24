package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func main() {
	// 第一题
	fmt.Println("~~~~~~~~~~~~~~第一题~~~~~~~~~~~~~")
	s1 := make([]int, 3, 4) // 定义一个切片，长度为3，容量为5
	s2 := append(s1, 1)

	fmt.Printf("%v, %d, %d\n", s1, len(s1), cap(s1)) // [0 0 0], 3, 4
	fmt.Printf("%v, %d, %d\n", s2, len(s2), cap(s2)) // [0 0 0 1], 4, 4

	// 请问s1、s2内各有什么元素？
	// s1 为 [0 0 0], 3, 4
	// s2 为 [0 0 0 1], 4, 4

	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~")

	s1[1] = 300
	fmt.Printf("%v, %p, %p\n", s1, &s1, &s1[1]) // [0 300 0], 0xc00009a060, 0xc0000a4068
	fmt.Printf("%v, %p, %p\n", s2, &s2, &s2[1]) // [0 300 0 1], 0xc00009a078, 0xc0000a4068
	// s1修改一个元素会影响s2吗？s2修改一个元素会影响s1吗？	会

	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~")

	s2 = append(s2, 200)
	fmt.Printf("%v, %p, %p\n", s1, &s1, &s1[1]) // [0 300 0], 0xc000008078, 0xc0000161c8
	fmt.Printf("%v, %p, %p\n", s2, &s2, &s2[1]) // [0 300 0 1 200], 0xc000008090, 0xc000010288

	// s2再增加一个元素会怎么样？	会分道扬镳，各玩各的。

	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~")
	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~")

	// 第二题有一个数组 [1,4,9,16,2,5,10,15]，生成一个新切片，要求新切片元素是数组相邻2项的和。
	fmt.Println("~~~~~~~~~~~~~~第二题~~~~~~~~~~~~~")

	a1 := [...]int{1, 4, 9, 16, 2, 5, 10, 15} // 定义一个数组
	fmt.Printf("%T, %d, %d\n", a1, len(a1), cap(a1))

	// 一个空切片
	a2 := make([]int, 0)
	fmt.Println(a2)

	// a2 = append(a2, a1[1]+a1[2])
	// fmt.Println(a2)

	for i := 0; i < len(a1)-1; i++ {
		// fmt.Println(a1[i])
		a2 = append(a2, a1[i]+a1[i+1])
	}
	fmt.Println(a2)

	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~")
	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~")

	/*
	   第三题
	   	数字重复统计
	   	随机产生100个整数
	   	数字的范围[-200, 200]
	   	升序输出这些生成的数字并打印其重复的次数
	*/
	fmt.Println("~~~~~~~~~~~~~~第三题~~~~~~~~~~~~~")

	// 定义随机数
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	m1 := make([]int, 0, 100) // 定义一个长度为0，容量为100的切片

	for i := 0; i < 100; i++ {
		// fmt.Println(i, r.Intn(401)-200)
		m1 = append(m1, r.Intn(401)-200) // 切片添加[-200,200]随机数

	}

	sort.Ints(m1)

	m2 := make(map[int]int, 100) // 定义一个容量100，长度为0的字典
	fmt.Println(m2, len(m2))

	for _, o := range m1 {
		// fmt.Println(i, o)
		if _, ok := m2[o]; !ok {
			m2[o] = 1

		} else {
			m2[o] += 1
		}

	}
	fmt.Printf("升序输出这些数字:%v", m1)

	for i, v := range m2 {
		// fmt.Println(i, v)
		fmt.Printf("%v出现了%v\n", i, v)
	}

}

// 批改意见
// 1. 基本的思维逻辑非常不错，到了第4周，尝试一下将里面的每个题目都独立成函数来完成。
// 2. s2修改最后一个元素时不会影响s1
