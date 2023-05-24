/*
作业1：打印99乘法表
1*1=1
1*2=2	2*2=4
1*3=3	2*3=6	3*3=9
1*4=4	2*4=8	3*4=12	4*4=16
1*5=5	2*5=10	3*5=15	4*5=20	5*5=25
1*6=6	2*6=12	3*6=18	4*6=24	5*6=30	6*6=36
1*7=7	2*7=14	3*7=21	4*7=28	5*7=35	6*7=42	7*7=49
1*8=8	2*8=16	3*8=24	4*8=32	5*8=40	6*8=48	7*8=56	8*8=64
1*9=9	2*9=18	3*9=27	4*9=36	5*9=45	6*9=54	7*9=63	8*9=72	9*9=81
*/

package main

import (
	"fmt"
	"time"
)

// 获取程序运行的时长
func runProgCost(n string, p func()) {
	t1 := time.Now()
	p()
	t2 := time.Since(t1)
	fmt.Printf("program %s run cost time: %v\n", n, t2)
}

// 使用一次for循环打印99乘法表
func test1() {
	j := 1
	count := 10
	// 控制被乘数=行数
	for i := 1; i < count; i++ {
		//控制乘数=列数，每一行乘数小于等于被乘数
		if j != i {
			// 除了最后一个乘法表达式都是以制表符结尾
			fmt.Printf("%d*%d=%d\t", j, i, i*j)
			// 每一行乘数一直自增到等于被乘数
			j++
			// 为了保证每一行被乘数不变,这里被乘数先减1，for循环被乘数再加1等于被乘数没有变，直到换行即乘数=被乘数
			i--
		} else {
			fmt.Printf("%d*%d=%d\n", j, i, i*j)
			//每行第一个乘法表达式乘数都是1
			j = 1
		}
	}
	// time.Sleep(3 * time.Second)
}

//使用两次循环打印99乘法表
func test2() {
	count := 10
	// 控制被乘数=行数
	for i := 1; i < count; i++ {
		//控制乘数=列数，每一行乘数小于等于被乘数
		for j := 1; j <= i; j++ {
			// 除了最后一个乘法表达式都是以制表符结尾
			if j != i {
				fmt.Printf("%d*%d=%-3d", j, i, i*j)
			} else {
				fmt.Printf("%d*%d=%d\n", j, i, i*j)
			}
		}
	}
}

// 执行入口
func main() {
	runProgCost("test1", test1)
	runProgCost("test2", test2)
}
