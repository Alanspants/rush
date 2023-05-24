package main

import "fmt"

//作业第-题，n的阶乘（递归）
func Factorial(n, y int) int {
	if n == 1 {
		return y
	}
	return Factorial(n-1, y*n)
}

//优化版本的第二题，但是for循环嵌套了好多层
//判断最高位并得出长度
func PdWidth(a int) int {
	w := 2
	if n := 10; a < 10 {
		return 2
	} else {
		for l := a; l-n >= 0; n = n * 10 {
			w += 1
		}
		return w
	}
}

func test(num int) {
	var width int
	for x := 1; x <= num; x++ {
		for i := num; i >= 1; i-- {
			//判断最高位
			width = PdWidth(i)
			if i <= x {
				fmt.Printf("%*d", width, i)
			} else {
				fmt.Printf("%*s", width, "")
			}
		}
		fmt.Println()
	}
}
func main() {
	test(60)
	//fmt.Printf("%d", Factorial(5, 1))

}
