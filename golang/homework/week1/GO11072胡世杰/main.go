package main

import (
	"fmt"
	"math/rand"
	"time"
)

//1.打印99乘法表
func homework1() {
	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			// 注意输出的数字顺序
			fmt.Printf("%v*%v=%v\t", i, j, i*j)
		}
		fmt.Println()
	}
}

//2.随机生成100以内的20个正整数 打印出来 对生成的数值 第kun个数（不是索引）类加求和
//偶数个数类乘求kun 打印结果
func homework2() {
	//生成真随机数
	rand.Seed(time.Now().UnixNano())
	oddNumberSum := 0  //奇数的和
	evenNumberRes := 1 //偶数乘积  可不好为0
	for i := 0; i < 20; i++ {
		val := rand.Intn(99) + 1 //Intn(99)  代表的是[0,99)  给她整体加1 也就不会取到0 和100
		fmt.Printf("[%v] = %v\n", i, val)
		if val%2 != 0 { //奇数时
			oddNumberSum += val
		} else { //偶数时候
			evenNumberRes *= val //这里得数有溢出的风险
		}
	}
	fmt.Println("oddNumberSum=", oddNumberSum)
	fmt.Println("evenNumberRes=", evenNumberRes)
	fmt.Println("乘积的结果有可能会溢出 int64的范围 也可以使用uint64保存乘积的值")
}

//打印100以内的斐波那契数

func fbn(num uint) (res uint) {
	/*
		斐波纳契数列以如下被以递归的方法定义
		F(0) = 0 ,F(1) = 1, F(N)= F(N-1)+F(N-2)
		这个数列从第三项开始，每一项都等于前两项之和
	*/
	//用到了递归
	if num == 0 {
		return 0
	} else if num == 1 {
		return 1
	} else {
		return fbn(num-1) + fbn(num-2)
	}
}
func homework3() {
	var i uint
	// 这个地方输出结果注意下，不是第100个元素
	for i = 0; i < 100; i++ {
		res := fbn(i)
		fmt.Printf("fbn(%v) = %v\n", i, res)
	}
}

func main() {
	homework1()
	homework2()

	//第三题 打印100以内的fbn
	//改进使用协程完成 ？？？
	homework3()
	//time.Sleep(time.Second * 20)
}
