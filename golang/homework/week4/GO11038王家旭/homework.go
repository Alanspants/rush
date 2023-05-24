package main

import (
	"fmt"
	// "math/rand"
	// "time"
)

func one() {
	// 请问s1、s2内各有什么元素？
	// 答：s1有0 0 0 s2有0 0 0 1
	// s1修改一个元素会影响s2吗？s2修改一个元素会影响s1吗？
	// 会 不会
	// s2再增加一个元素会怎么样？
	// 内存地址发生变更，已部署于同一个切片
	s1 := make([]int, 3, 4)
	s2 := append(s1, 1)
	fmt.Println(s1, s2)
}

func two() {
	// 有一个数组 [1,4,9,16,2,5,10,15]，生成一个新切片，要求新切片元素是数组相邻2项的和
	a1 := []int{1, 4, 9, 16, 2, 5, 10, 15}
	var a2 = []int{}
	for i := 0; i < len(a1); i++ {
		//a2 = append(a2, a1[i]+a1[i+1])
		if i == 7 {
			break
		} else {
			a2 = append(a2, a1[i]+a1[i+1])
		}
	}
	fmt.Println(a2)
}

func main() {
	one()
	two()
}

// 批改意见
// 1. s2修改部分元素会影响s1，修改最后一个不影响
// 2. 第三题也可以尝试写一下
