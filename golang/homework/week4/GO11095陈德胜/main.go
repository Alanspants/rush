package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func main() {
	/*
		作业1
		s1 := make([]int, 3, 4)
		s2 := append(s1, 1)

		1、请问s1、s2内各有什么元素？
			s1: [0 0 0]
			s2: [0 0 0 1]
		2、s1修改一个元素会影响s2吗？s2修改一个元素会影响s1吗？
			会影响， 因为s1、s2共用底层数组
		3、s2再增加一个元素会怎么样？
			s2增加元素就会超过原来数组cap，会重新分配一个容量为8的新的底层数组， 与s1没有关系了
	*/

	s1 := make([]int, 3, 4)
	s2 := append(s1, 1)
	fmt.Println("s1:", s1)
	fmt.Println("s2:", s2)
	fmt.Printf("s1 的指针： %p %p   s2的指针: %p %p \n", &s1, &s1[0], s2, &s2[0])
	s1[0] = 2
	fmt.Println("s1:", s1)
	fmt.Println("s2:", s2)
	s3 := append(s2, 3)
	fmt.Printf("s1 的指针： %p %p len: %d, cap:%d\n", &s1, &s1[0], len(s1), cap(s1))
	fmt.Printf("s2 的指针： %p %p len: %d, cap:%d\n", &s2, &s2[0], len(s2), cap(s2))
	fmt.Printf("s3 的指针： %p %p len: %d, cap:%d\n", &s3, &s3[0], len(s3), cap(s3))
	fmt.Println("s1:", s1)
	fmt.Println("s2:", s2)
	fmt.Println("s3:", s3)

	/*
		作业二
		 有一个数组 [1,4,9,16,2,5,10,15]，生成一个新切片，要求新切片元素是数组相邻2项的和
	*/
	s4 := [...]int{1, 4, 9, 16, 2, 5, 10, 15}
	var s5 []int
	for i := 0; i < len(s4); i += 2 {
		// fmt.Println(i, s4[i], s4[i+1])
		s5 = append(s5, s4[i]+s4[i+1])
	}
	fmt.Println(s5)

	/*
		作业三：
		随机产生100个整数
		数字的范围[-200, 200]
		升序输出这些生成的数字并打印其重复的次数
	*/
	arr := make([]int, 0, 100)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < 100; i++ {
		arr = append(arr, r.Intn(400)-200)
	}
	fmt.Println("随机生成数组:", arr)
	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~")
	sort.Sort(sort.IntSlice(arr))
	fmt.Println("升序数组： ", arr)
	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~")
	arrMap := make(map[int]int, len(arr))
	for _, v := range arr {
		arrMap[v]++
	}
	fmt.Println("数组中的数据重复出现的次数：", arrMap)
}

// 批改意见
// 1. s2修改最后一个元素不会影响s1
// 2. Intn的范围是[0, n)，因此Intn(400)-200会漏掉最后一个值200
// 3. 代码中记得写注释，方便后面自己查看
