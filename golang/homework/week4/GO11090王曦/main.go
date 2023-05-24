package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func func1() {
	fmt.Println("习题1.")
	//s1 := make([]int, 3, 4)
	//s2 := append(s1, 1)
	// ############# s1,s2各有什么元素？
	// s1 [0 0 0]
	// s2 [0 0 0 1]
	// ############# s1修改一个元素会影响s2吗？ s2修改一个元素会影响s1吗？
	// s1与s2共用一底层数组，因s1\s2的len不同，s1修改会影响s2，s2前三元素修改影响s1，s2[3]修改不影响s1
	// ############# s2再增加一个元素会怎样
	// s2与s1共用一个底层数组，s1在make时确定了容量为4，而s2已经容量为4了，当s2再进行append时，s2会发生扩容，底层数组与s1分道扬镳

}

func func2() {
	fmt.Println("习题2.")
	a := [8]int{1, 4, 9, 16, 2, 5, 10, 15} // 有一个数组

	s := make([]int, 0, 7)          // 生成新的切片
	for i := 0; i < len(a)-1; i++ { // 计算相邻元素，故末元素不循环
		s = append(s, a[i]+a[i+1]) // 相邻元素和
	}
	fmt.Println(s) // a数组的相邻元素和的切片
}

func func3() {
	fmt.Println("习题3.")
	// 随机生成100整数
	fmt.Println("随机100整数，区间[-200, 200]")
	s := make([]int, 0, 100)
	for i := 0; i < 100; i++ {
		s = append(s, randInt(-200, 200))
	}
	fmt.Println(s)

	// 升序排列
	fmt.Println("升序")
	sort.Ints(s)
	fmt.Println(s)

	// 打印重复次数
	fmt.Println("打印重复次数")
	m := make(map[int]int)        // 生成一个k\v都是int的map
	for i := 0; i < len(s); i++ { // 外层循环每个元素
		n := s[i]
		m[n] += 1 // 对相同数循环到了+1
	}
	fmt.Println(m)

}

// 传入大小值，以满足生成负数需求
func randInt(lower, upper int) int {
	rand.Seed(time.Now().UnixNano())
	time.Sleep(1 * time.Millisecond) // 睡1ms, 防止seed一致
	rng := (upper + 1) - lower       // intn [0, n) 范围，所以upper + 1 包含最大值
	return rand.Intn(rng) + lower
}

func main() {
	func1()
	func2()
	func3()
}

// 批改意见
// 1. 随机数的逻辑可以尝试着再简化一下
