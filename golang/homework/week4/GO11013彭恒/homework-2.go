package main

import (
	"fmt"
)

func main() {
	a1 := [8]int{1, 4, 9, 16, 2, 5, 10, 15}
	//a1 := [9]int{1, 4, 9, 16, 2, 5, 10, 15, 99}
	// 判断长度奇偶
	n := len(a1) & 1
	// 长度一半，因为Go语言是静态语言，在定义变量时就指定了变量类型为int，所以编译器推导出来的运算结果也为int
	// 只保留结果的整数部分
	newLen := len(a1) / 2
	// 定义新切片s1
	var s1 = make([]int, 0, newLen+1)
	// 定义数组索引，初始值为0
	index := 0
	if n == 0 {
		for i := 0; i < newLen; i++ {
			s1 = append(s1, a1[index]+a1[index+1])
			index += 2
		}
	} else {
		for i := 0; i < newLen; i++ {
			s1 = append(s1, a1[index]+a1[index+1])
			index += 2
		}
		// 数组元素为单数，剩余最后一个元素，直接追加到切片中
		s1 = append(s1, a1[len(a1)-1])
	}
	fmt.Println(s1)
}

// 批改意见
// 1. 数组相邻两项的和指的是第1和第2，第2和第3，第3和第4个相加，依次类推
