// 包
package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*
入口函数
*/
func main() {
	test01()
	test02()
	test03()
}

func test01() {
	// 打印九九乘法表
	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%v*%v=%v\t", j, i, j*i)
		}
		fmt.Printf("\n")
	}

}

/*
随机生成100以内的20个非0正整数，打印出来。对生成的数值，第单数个（不是索引）累加求和，
第偶数个累乘求积。
*/
func test02() {
	//计算乘积
	var pro int64 = 1
	var sum int32 = 0
	//随机数生成器
	r := rand.New(rand.NewSource(time.Now().UnixMicro()))
	for i := 1; i <= 20; i++ {
		number := r.Int31n(99) + 1
		if i&1 == 0 {
			fmt.Printf("第偶数个数的下标为%v,数值为%v\n", i, number)
			pro *= int64(number)
		} else {
			fmt.Printf("第奇数个数的下标为%v,数值为%v\n", i, number)
			sum += number
		}
		// if i%2 == 0 {
		// 	fmt.Printf("第偶数个数的下标为%v,数值为%v\n", i, number)
		// 	pro *= int64(number)
		// } else {
		// 	fmt.Printf("第奇数个数的下标为%v,数值为%v\n", i, number)
		// 	sum += number
		// }
	}
	fmt.Printf("第奇数个累加求和的结果为:%v\n", sum)
	fmt.Printf("第偶数个累乘求积的结果为:%v\n", pro)
}

/*
打印100以内的斐波那契数列
斐波那契数列（Fibonacci sequence），又称黄金分割数列，因数学家莱昂纳多·斐波那契（Leonardo Fibonacci）以兔子繁殖为例子而引入，
故又称为“兔子数列”，指的是这样一个数列：1、1、2、3、5、8、13、21、34、……
在数学上，斐波那契数列以如下被以递推的方法定义：F(0)=0，F(1)=1, F(n)=F(n - 1)+F(n - 2)（n ≥ 2，n ∈ N*）
*/
func test03() {
	fmt.Printf("100以内的斐波那契数列为")
	for i := 1; i < 20; i++ {
		number := fibonacci(int16(i))
		if number < 100 {
			fmt.Printf("%v\t", number)
		} else {
			fmt.Printf("\n")
			break
		}
	}
}

func fibonacci(index int16) int32 {
	switch index {
	case 0, 1:
		return int32(index)
	// case 1:
	// 	return 1
	default:
		return fibonacci(index-1) + fibonacci(index-2)
	}
}

// 批改意见
