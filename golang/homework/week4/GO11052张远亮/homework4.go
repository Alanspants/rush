package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

/**
s1 := make([]int,3,4)
s2 := append(s1,1)
(1)请问s1、s2内各有什么元素？
s1 [0,0,0]
s2 [0,0,0,1]
(2)s1修改一个元素会影响s2吗？s2修改一个元素会影响s1吗？
s1修改一个元素会影响s2,因为s1和s2共用一个底层数组。只是pointer不一样
s2修改元素会影响s1，以为共用一个底层数组
(3)s2再增加一个元素会怎么样？
设置的初始容量是4，现在append以后已经达到4个了，如果在增加一个元素会导致扩容，s1和s2会分道扬镳，就不能共有一个底层数组。
**/

// 2、有一个数组 [1,4,9,16,2,5,10,15]，生成一个新切片，要求新切片元素是数组相邻2项的和。
func arraySum() {
	a1 := [8]int{1, 4, 9, 16, 2, 5, 10, 15}
	var s1 []int

	for i := 0; i < len(a1)-1; i++ {
		s1 = append(s1, a1[i]+a1[i+1])
	}
	fmt.Println(s1)
}

/*
3、数字重复统计
随机产生100个整数
数字的范围[-200, 200]
升序输出这些生成的数字并打印其重复的次数
*/
func sortRand() {
	sum := make([]int, 0, 100)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < 100; i++ {
		randNum := r.Intn(401) - 200
		sum = append(sum, randNum)
	}
	sort.Ints(sum)
	fmt.Println(sum)
	m1 := make(map[int]int)
	for _, v := range sum {
		if m1[v] != 0 {
			m1[v]++
		} else {
			m1[v] = 1
		}
	}
	//for k, v := range m1 {
	//	fmt.Printf("数字: %d重复的次数: %d\n", k, v)
	//}
	fmt.Println(m1)
}

/*
	3、数字重复统计
	随机产生100个整数
	数字的范围[-200, 200]
	升序输出这些生成的数字并打印其重复的次数

	第二种方式
*/

func sortRand2() {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	s1 := make([]int, 0, 100)
	m1 := make(map[int]int)
	keys := make([]int, 0, 100)
	for i := 0; i < 100; i++ {
		n := r.Intn(100) - 50
		s1 = append(s1, n)
		if _, ok := m1[n]; !ok {
			m1[n] = 0
			keys = append(keys, n)
		}
		m1[n]++
	}
	//降序排序
	sort.Sort(sort.Reverse(sort.IntSlice(keys)))
	for k, v := range keys {
		fmt.Println(k, v, keys[k], m1[v])
	}
}

func main() {
	arraySum()
	//sortRand()
	sortRand2()

}

// 批改意见
// 1. 代码中记得写注释，方便自己以后查看
// 2. s2 修改最后一个元素不影响s1
