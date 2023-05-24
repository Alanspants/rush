package main

import (
	"errors"
	"fmt"
	"time"
)

//1.求n的阶乘 递归方法实现

func jieCheng1(n int) int {
	//漂亮公式

	if n < 0 {
		panic(errors.New("方法一，负数错误"))
	}
	if n == 1 {
		return 1
	}
	result := jieCheng1(n-1) * n
	return result
}

//阶乘for循环版
func jieCheng2(n int, res int) int {
	//res = 4*3 = 12
	if n < 0 {
		panic(errors.New("方法二，负数错误"))
	}
	for i := 1; i < n || n >= 1; i++ {
		//fmt.Println("第", i, "次", "n=", n, "res=", res)
		res = n * res
		n--
	}
	return res
}

// UpperS 编写函数 接收一个参数n下，n为正整数 上三角打印方式 要求数字必须对齐
/*   上三角  1.先易后难2.先死后活
第1行 2个   1个数 1个空 先打印 3*2-2*1个空  4个空 再打印1个数字
第2行 4个   2个数 2个空 先打印 3*2-2*2个空  2个空 再打印2个数字
第3行 6个   3个数 3个空 先打印 3*2-2*3个空  0个空 再打印3个数字
第i行 2*i个 i个数 i个空 先打印 n*i-2*i个空格      再打印i个数字
-----------------
即每行 先打印n*i-2*i个空格 再打印i个数字 每个数字后都价格空格即可
*/
func UpperS(n int) {
	fmt.Println("上三角")
	for i := 1; i <= n; i++ {
		//每行先打印 n-i个空格
		for k := 1; k <= n*2-2*i; k++ { //3要替换成n
			fmt.Print(" ")
		}
		//每行再打印i个数 i代表所在层数 [1，n]
		for j := i; j >= 1; j-- { //先打印空格 在打印星星
			fmt.Print(j)
			fmt.Print(" ")

		}
		fmt.Println()
	}
}

func DownS(n int) {
	fmt.Println("下三角")
	for i := 1; i <= n; i++ {
		//每行先打印 n-i个空格
		//for k := 1; k <= n*2-2*i; k++ { //3要替换成n
		//	fmt.Print(" ")
		//}
		//每行再打印i个数 i代表所在层数 [1，n]
		for j := i; j >= 1; j-- { //先打印空格 在打印星星
			fmt.Print(j)
			fmt.Print(" ")
		}
		fmt.Println()
	}
}

func main() {
	start := time.Now()
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("捕获异常", err)
		}
	}() //为什么输入负数 阶乘公式报错后 不会接着执行 上三角公式 而是直接退出了?
	var n int //要求的阶乘
	fmt.Println("输入要求的阶乘数")
	fmt.Scan(&n)
	fmt.Println("方法一", jieCheng1(n))
	fmt.Println("方法二", jieCheng2(n, 1))
	fmt.Println("-------")
	UpperS(9) //超过10不会做
	DownS(9)
	cost := time.Since(start)
	fmt.Println("程序运行时间", cost)

}
