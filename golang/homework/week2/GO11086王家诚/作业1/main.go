package main

import "fmt"

//打印九九乘法表  要求间隔均匀
func main() {
	for i := 1; i <= 9; i++ {
		for x := 1; x <= i; x++ {
			fmt.Printf("%d*%d=%d\t", i, x, i*x)
		} //   \t 间隔一个制表符
		fmt.Println() //  当循环完成一次后换行继续第二次循环
	}

}

// i累加 当i=1时 x<=i  输出 1 x++ x=2 x<=i不成立 返回第一层循环  此时i=2
//x=1 x<=i  输出1  x++ x=2 x<=i  输出2  以此类推
/*
i,x
1 1
2 1
2 2
3 1
3 2            用x*i 完成计算结果
3 3
4 1
4 2
4 3
4 4
............
*/
