package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func main() {
	/*
		作业1
	*/
	s1 := make([]int, 3, 4)
	s2 := append(s1, 1)
	// s1 []int{0,0,0}
	// s2 []int{0,0,0,1}
	// s1或s2修改都会影响另外一个切片，共用底层数组
	// s2增加元素就会超过cap，会分配一个容量为8的新的底层数组，与s1就没有关系了
	// 验证
	fmt.Printf("s1:%v,%p\n", s1, &s1[0])
	fmt.Printf("s2:%v,%p\n", s2, &s2[0])
	s1[0] = 1
	fmt.Printf("s1:%v,%p\n", s1, &s1[0])
	fmt.Printf("s2:%v,%p\n", s2, &s2[0])
	s2[1] = 2
	fmt.Printf("s1:%v,%p\n", s1, &s1[0])
	fmt.Printf("s2:%v,%p\n", s2, &s2[0])
	s2 = append(s2, 5)
	s1[1] = 3
	fmt.Printf("s1:%v,%p\n", s1, &s1[0])
	fmt.Printf("s2:%v,%p,%v\n", s2, &s2[0], cap(s2))
	fmt.Println(strings.Repeat("-", 30))
	/*
		作业2
	*/
	s := [...]int{1, 4, 9, 16, 2, 5, 10, 15}
	sCap := len(s) - 1
	s3 := make([]int, 0, sCap)
	for i, v := range s {
		s3 = append(s3, v+s[i+1])
		if i == sCap-1 {
			break
		}
	}
	fmt.Println(s3)
	fmt.Println(strings.Repeat("-", 30))

	/*
		作业3
	*/
	rand.Seed(time.Now().UnixMicro())
	randMap := make(map[int]int, 100)
	for i := 0; i < 100; i++ {
		randMap[rand.Intn(401)-200] += 1
	}
	fmt.Println(randMap)
	sum := 0
	for _, v := range randMap {
		sum += v
	}
	fmt.Println(len(randMap))
	fmt.Printf("SUM:%v\n", sum)
}

// 批改意见
// 1. 代码需要写注释，方便后面自己查看
// 2. 第四周了，可以尝试用函数来完成
