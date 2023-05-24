package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

// 2、有一个数组 [1,4,9,16,2,5,10,15]，生成一个新切片，要求新切片元素是数组相邻2项的和。
func Transform() {
	arr := [...]int{1, 4, 9, 16, 2, 5, 10, 15}
	var s = []int{}
	count := len(arr) - 1        //要加的次数
	for i := 0; i < count; i++ { //index 8 out of bounds [0:8] 防止索引i+1越界，所以循环条件要小于length-1即<7循环中i最大为6,6+1=7不会越界
		s = append(s, arr[i]+arr[i+1])
	}
	fmt.Printf("arr:%#v\n", arr)
	fmt.Printf("slice: %#v\n", s)
}

// 3、数字重复统计
// 随机产生100个整数
// 数字的范围[-200, 200]
// 升序输出这些生成的数字并打印其重复的次数
func StatisticsRepetition(total, start, end int) map[int]int {
	r := rand.New(rand.NewSource(time.Now().UnixNano())) //新的随机数生成器
	var random int
	m := make(map[int]int, 0)
	keys := make([]int, 0, total)
	for i := 0; i < total; i++ {
		if start <= 0 { //判断随机数起始数是正数还是负数
			random = r.Intn(end+1) + start //传负数，直接加
		} else {
			random = r.Intn(end+1) - start //正数就减
		}
		if _, ok := m[random]; !ok {
			m[random] = 0
			keys = append(keys, random)
		} else {
			m[random]++
		}
	}
	sort.Sort(sort.Reverse(sort.IntSlice(keys)))
	for i, k := range keys {
		fmt.Printf("%d：%d,\t %d:%d\n", i, k, keys[i], m[k])
	}
	fmt.Println(keys)
	return m
}
func main() {
	// 请问s1、s2内各有什么元素？答：s1[0 0 0]、s2[0 0 0 1]
	// s1修改一个元素会影响s2吗？会影响。s2修改一个元素会影响s1吗？不一定，修改s2[3]不会影响。
	// s2再增加一个元素会怎么样？会扩容，子切片与母切片不再共用一个底层数组，s2与s1分道扬镳。
	s1 := make([]int, 3, 4)
	s2 := append(s1, 1)
	fmt.Printf("s1 %+v\t 长度:%10d\t 容量:%d\t %p\t %p\n", s1, len(s1), cap(s1), &s1, &s1[0])
	fmt.Printf("s2 %+v\t 长度:%10d\t 容量:%d\t %p\t %p\n", s2, len(s2), cap(s2), &s2, &s2[0])
	fmt.Println("——————————————————————————————————————————————————————————————————————————————————————————————————————————")
	for i := 0; i < len(s1); i++ {
		s1[i] = i + 30 //修改s1里的每一个值
	}
	fmt.Println(s1)
	fmt.Println(s2)
	s2[3] = 100
	fmt.Println(s1, s2)  //改变s2的第三个元素，s1没变
	s1 = append(s1, 200) //s1追加一个s1[3]:200就把s2[3]的值100覆盖了
	fmt.Println(s1, s2)
	s2 = append(s2, 300)
	fmt.Printf("s1 %+v\t 长度:%10d\t 容量:%d\t %p\t %p\n", s1, len(s1), cap(s1), &s1, &s1[0])
	fmt.Printf("s2 %+v\t 长度:%10d\t 容量:%d\t %p\t %p\n", s2, len(s2), cap(s2), &s2, &s2[0])
	fmt.Println("-------------------------------------作业2开始--------------------------------------------------")
	Transform() // []int{5, 13, 25, 18, 7, 15, 25}
	fmt.Println("==============================================作业3开始========================================================")
	fmt.Println(StatisticsRepetition(100, 200, 100))
}

// 批改意见
// 1. 随机数的范围不要搞的太复杂，太复杂也没有注释，以后自己回头看的时候都看不懂
