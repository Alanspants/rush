package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func main() {
	//1、看下面一段程序，回答问题
	// s1 := make([]int, 3, 4)
	// s2 := append(s1, 1)
	//答：s1、s2内的元素分别为 s1=[0 0 0] , s2=[0 0 0 1]。
	//	 s1修改一个元素会影响s2；s2修改下标为3的元素不会影响s1，修改其他元素会影响s1。
	//	 s2再增加一个元素会导致需要扩容s2底层使用的数组，从而不再与s1共用底层数组。以后s1和s2将没有关联。

	// 2、有一个数组 [1,4,9,16,2,5,10,15]，生成一个新切片，要求新切片元素是数组相邻2项的和。
	arr := [...]int{1, 4, 9, 16, 2, 5, 10, 15}
	s := []int{}
	var v int
	for i := 0; i < len(arr)-1; i++ {
		v = arr[i] + arr[i+1]
		s = append(s, v)
	}
	fmt.Println(s)
	fmt.Println("~~~~~~~~~~~~~~~~~~~~~")
	fmt.Println()

	// 3、数字重复统计
	// 随机产生100个整数
	// 数字的范围[-200, 200]
	// 升序输出这些生成的数字并打印其重复的次数
	s2 := make([]int, 100)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < len(s2); i++ {
		s2[i] = r.Intn(401) - 200 //0->400 >> -200->200
	}
	fmt.Println("(1)随机生成的数字为：\n", s2)
	sort.Ints(s2) //排序
	fmt.Println("(2)排序后的随机数字为：\n", s2)

	//统计随机数字出现的次数
	m := map[int]int{}
	for _, v := range s2 {
		m[v] += 1
	}
	fmt.Println("(3)升序输出这些生成的随机数及重复次数如下：")
	fmt.Printf("%4[1]s: %[2]s |%4[1]s: %[2]s |%4[1]s: %[2]s |%4[1]s: %[2]s |%4[1]s: %[2]s |%4[1]s: %[2]s |%4[1]s: %[2]s |%4[1]s: %[2]s |%4[1]s: %[2]s |%4[1]s: %[2]s |", "Num", "Counts")
	for i, v := range s2 {
		if i%10 == 0 {
			fmt.Println()
		}
		fmt.Printf("%4d: %3d    |", v, m[v])
	}

}

// 批改意见
// 1. 基本的代码思维逻辑都非常好，第4周可以尝试将功能用函数来实现
