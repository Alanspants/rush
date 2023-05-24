package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func homework1() {
	s1 := make([]int, 3, 4)
	s2 := append(s1, 1)

	// 请问s1、s2内各有什么元素？
	fmt.Println(s1) //[0 0 0]
	fmt.Println(s2) //[0 0 0 1]
	// s1修改一个元素会影响s2吗？s2修改一个元素会影响s1吗？
	fmt.Println("s1修改一个元素会影响s2;s2修改前3个元素会影响s1")
	// s2再增加一个元素会怎么样？
	fmt.Println("s2再增加一个元素会扩容")
}
func homework2() {
	arr := [8]int{1, 4, 9, 16, 2, 5, 10, 15}
	s1 := make([]int, 0, 8)
	for i := 0; i < (len(arr) - 1); i++ {
		s1 = append(s1, arr[i]+arr[i+1])
	}
	fmt.Println(s1)
}

func homework3() {
	var num int
	s2 := make([]int, 0, 100)
	s3 := make([]int, 0, 100)
	m1 := make(map[int]int, 100)

	r := rand.New(rand.NewSource(time.Now().UnixNano())) //随机数生成器
	for i := 0; i < 100; i++ {
		num = r.Intn(401) - 200
		s2 = append(s2, num)
	}

	for _, v := range s2 {
		m1[v] = m1[v] + 1
	}
	for k := range m1 {
		s3 = append(s3, k)
	}
	sort.Ints(s3)
	for j := 0; j < len(s3); j++ {
		fmt.Println(s3[j], m1[s3[j]])
	}
}

func main() {
	homework1()
	homework2()
	homework3()
}

// 批改意见
// 1. 代码逻辑都ok，代码中记得写注释，方便自己后面查看
