package main

import "fmt"

//打印100以内的斐波那契数列,又称为“兔子数列”，指的是这样一个数列：1、1、2、3、5、8、13、21、34、……在数学上，斐波那契数列以如下被以递推的方法定义：F(0)=0，F(1)=1, F(n)=F(n - 1)+F(n - 2)（n ≥ 2，n ∈ N*）
func f(n int) int {
	if n == 1 || n == 2 {
		return 1
	} else {
		return f(n-1) + f(n-2)
	}

}
func main() {
	for i := 1; i <= 20; i++ {
		if f(i) < 100 {
			fmt.Printf("%d ", f(i))
		} else {
			break
		}

	}

}

// 批改意见
// 逻辑没有问题，注意代码格式，例如空行，注释等
