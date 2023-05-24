package main

import "fmt"

func factorial(n int) int {
	switch {
	case n < 0:
		panic("输入的数字有误")
	case n == 0 || n == 1:
		return 1
	}
	num := n * factorial(n-1)
	return num
}

func triangle(n int) {
	var width int
	for i := 1; i <= n; i++ {
		for j := n - i; j > 0; j-- {
			if n-j >= 9 {
				width = 3
			} else {
				width = 2
			}
			fmt.Printf("%-[2]*[1]v", (" "), width)
		}
		for k := 0; k < i; k++ {
			if i-k >= 10 {
				width = 3
			} else if i-k < 10 {
				width = 2
			}
			fmt.Printf("%-[2]*[1]d", (i - k), width)
		}
		fmt.Println()
	}
}

func main() {
	// homework1
	// var n int
	// fmt.Println("求n的阶乘,请输入数字:")
	// fmt.Scanf("%d", &n)
	num := factorial(10)
	fmt.Println(num)

	// homework2
	triangle(12)
}

// 批改意见
// 代码逻辑和结果都是对的，记得函数内写注释，代码太长后面自己看的时候都看不懂了。
