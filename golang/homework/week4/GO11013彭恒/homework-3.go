package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

// 参数n是需要的随机数个数，参数ran是随机数的范围，范围为[-ran，ran]
// 返回值为整型切片
func randInt(n int, ran int) []int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	s := make([]int, 0, n)
	for i := 0; i < n; i++ {
		n1 := r.Intn(ran) + 1
		f := r.Intn(2)
		if f == 0 {
			s = append(s, -n1)
		} else {
			s = append(s, n1)
		}
	}
	sort.Ints(s)
	return s
}

// 参数s是需要传递计数的切片，参数length是预估map的容量
// 返回值切片为map中key的升序，map为传递参数切片出现次数的键值对，键为切片元素，值为切片元素出现的次数
func timesCount(s []int, length int) ([]int, map[int]int) {
	m := make(map[int]int, length)
	for i := 0; i < len(s); i++ {
		v := m[s[i]]
		v += 1
		m[s[i]] = v
	}
	s1 := make([]int, 0, 100)
	for k, _ := range m {
		s1 = append(s1, k)
	}
	sort.Ints(s1)
	return s1, m
}

func main() {
	s := randInt(100, 200)
	fmt.Println("random slice:", s)

	s1, m := timesCount(s, 100)
	for i := 0; i < len(s1); i++ {
		// 根据升序key的切片，获取map中对应的值
		fmt.Printf("number %4d appears %d times.\n", s1[i], m[s1[i]])
	}
}
