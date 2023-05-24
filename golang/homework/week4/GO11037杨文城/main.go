package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

// 1、看下面一段程序，回答问题
// 请问s1、s2内各有什么元素？
// s1修改一个元素会影响s2吗？s2修改一个元素会影响s1吗？
// s2再增加一个元素会怎么样？

// 2、有一个数组 [1,4,9,16,2,5,10,15]，生成一个新切片，要求新切片元素是数组相邻2项的和。
// 3、数字重复统计
// 随机产生100个整数
// 数字的范围[-200, 200]
// 升序输出这些生成的数字并打印其重复的次数

func main() {
	fmt.Println("1、看下面一段程序，回答问题")
	work1()
	fmt.Println("2、有一个数组 [1,4,9,16,2,5,10,15]，生成一个新切片，要求新切片元素是数组相邻2项的和")
	work2()
	fmt.Println("3、数字重复统计")
	work3()
}

func work3() {
	myRand := rand.New(rand.NewSource(time.Now().UnixNano()))
	s1 := make([]int, 0, 100)
	for i := 0; i < 100; i++ {
		num := myRand.Intn(400) - 200
		s1 = append(s1, num)
	}
	sort.Ints(s1)
	fmt.Println(s1)
	m1 := make(map[int]int)
	for _, v := range s1 {

		if _, ok := m1[v]; !ok {
			m1[v] = 1
		} else {
			m1[v]++
		}
	}
	fmt.Println(m1)
}

func work2() {
	arr := [8]int{1, 4, 9, 16, 2, 5, 10, 15}
	s1 := make([]int, 0, 10)

	for i := 0; i < len(arr)-1; i++ {

		s1 = append(s1, arr[i]+arr[i+1])
	}
	fmt.Println(arr)
	fmt.Println(s1)
}

func work1() {

	s1 := make([]int, 3, 4)
	s2 := append(s1, 1)

	fmt.Printf(" s1 %p %p %v \n", &s1, &s1[0], s1)
	fmt.Printf(" s2 %p %p %v \n", &s2, &s2[0], s2)
	fmt.Println("s1改变-------------")
	s1[0] = 2
	fmt.Printf(" s1 %p %p %v \n", &s1, &s1[0], s1)
	fmt.Printf(" s2 %p %p %v \n", &s2, &s2[0], s2)
	fmt.Println("s2改变-------------")
	s2[0] = 2
	s2[3] = 2
	fmt.Printf(" s1 %p %p %v \n", &s1, &s1[0], s1)
	fmt.Printf(" s2 %p %p %v \n", &s2, &s2[0], s2)
	fmt.Println("s2增加元素-----------")
	s2 = append(s2, 2)
	fmt.Printf(" s1 %p %p %v \n", &s1, &s1[0], s1)
	fmt.Printf(" s2 %p %p %v \n", &s2, &s2[0], s2)
	fmt.Println("s1:[0,0,0];s2:[0,0,0,1]")
	fmt.Println("s1修改一个元素会影响s2; s2修改索引为[0-2]元素会影响s1")
	fmt.Println("s2再增加一个元素底层数组扩容。s1与s2分道扬镳")
}

// 批改意见
// 1. 代码逻辑ok，要记得写代码注释，方便后面自己查看
