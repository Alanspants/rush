package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func randNumsSort(total, limiter int) {
	// 定义随机数种子
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	nums := make([]int, 0, total)
	counter := make(map[int]int)
	keys := make([]int, 0, total)

	for i := 0; i < total; i++ {
		n := r.Intn(limiter) - limiter/2
		nums = append(nums, n)
		if _, ok := counter[n]; !ok {
			counter[n] = 0
			keys = append(keys, n)
		}
		counter[n]++
	}
	fmt.Println(nums)
	fmt.Println(counter)
	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")

	// 升序
	// 1. 切片 => IntSlice 强制类型转换
	// 2. sort
	// sort.Sort(sort.IntSlice(keys))
	// fmt.Println(keys)
	// fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")

	// 降序
	// 1. 切片 => IntSlice 强制类型转换
	// 2. Reverse
	// 3. sort
	sort.Sort(sort.Reverse(sort.IntSlice(keys)))
	for i, k := range keys {
		fmt.Println(i, k, counter[k])
	}
}

func main() {
	randNumsSort(100, 5)
}
