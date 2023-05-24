package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func sliceTest() {
	s1 := make([]int, 3, 4)
	fmt.Printf("%p %p; %d %d, %v\n", &s1, &s1[0], len(s1), cap(s1), s1)
	s2 := append(s1, 1)
	fmt.Printf("%p %p; %d %d, %v\n", &s2, &s2[0], len(s2), cap(s2), s2)
	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~")
	for i := 0; i < len(s1); i++ {
		s1[i] = i + 30
	}
	fmt.Println(s1)
	fmt.Println(s2)
	s2[3] = 100
	fmt.Println(s1, s2)
	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~")
	s1 = append(s1, 200)
	fmt.Println(s1, s2)
	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~")
	s2 = append(s2, 300)
	fmt.Printf("%p %p; %d %d, %v\n", &s1, &s1[0], len(s1), cap(s1), s1)
	fmt.Printf("%p %p; %d %d, %v\n", &s2, &s2[0], len(s2), cap(s2), s2)
}

func newSlice() {
	arr := [...]int{1, 4, 9, 16, 2, 5, 10, 15}
	fmt.Println(arr)
	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~")
	count := len(arr) - 1
	out := make([]int, 0, count)
	for i := 0; i < count; i++ {
		out = append(out, arr[i]+arr[i+1]) // i + 1有可能超界
	}
	fmt.Println(out)
}

func StatisticsRepetition(total int, limiter int) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	// limiter := 51
	// total := 100
	// map kv对，元素:出现的次数，不用管容量
	counter := make(map[int]int)
	fmt.Println(len(counter)) // 0 kv一个都没有
	// fmt.Println(counter[100])
	// keys := make([]int, 0, len(counter)) // len(counter) kv对数，key去重，不同的key多少个
	keys := make([]int, 0, total) // 至多100个

	for i := 0; i < total; i++ {
		n := r.Intn(limiter) - limiter/2 // 故意改了题目
		// nums[i] = n // nums := make([]int, total, total)
		// counter[n] = counter[n] + 1 // 不推荐，因为有风险
		if _, ok := counter[n]; !ok { // 推荐
			counter[n] = 0
			keys = append(keys, n)
		}
		counter[n]++ // map是按照key去重的
		// counter[n] += 1
	}
	fmt.Println(counter, len(counter))
	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~")
	// map中数字key直接升序了，为什么？很多map的key算法对于整数处理可以认为是取模 n mod 一个超级大整数
	// 降序
	// sort.Ints(x []int) // 1 3 sort.Sort(sort.IntSlice(keys))
	// 1 切片 => IntSlice 强制类型转换
	// 2 Reverse
	// 3 Sort
	sort.Sort(sort.Reverse(sort.IntSlice(keys)))
	// counter.keys // 自动生成代码 postfix，map数据结构没有提供直接获取所有key或者value的方法
	// keys := make([]int, 0, len(counter))
	// for k := range counter {
	// 	keys = append(keys, k)
	// }
	for i, k := range keys { // 就地修改了，降序的
		fmt.Println(i, k, keys[i], counter[k])
	}
}

// func main() {
// 	// sliceTest() // 2.1
// 	// newSlice() // 2.2
// 	nums(100, 11)

// }
