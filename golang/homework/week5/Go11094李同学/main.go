//1 2 3 4 5

package main

import (
	"fmt"
)

func fn1(n int) int {
	c := 1
	if n <= 0 {
		panic("n is negative")
	}
	if n == 1 {
		return 1
	}
	for i := 2; i <= n; i++ {
		c *= i
	}
	return c
}

func fn2(n int) int {
	switch {
	case n <= 0:
		panic("n is negative")
	case n == 1:
		return 1
	case n == 2:
		return 2
	}
	return n * fn2(n-1)
	//1*2*3*4*5*6  1 2 3 2 4
}

//a代表阶乘数，b给定初始值1
func fn3(a, b int) int {
	switch {
	case a <= 0:
		panic("n is negative")
	case a == 1:
		return b
	}
	return fn3((a - 1), a*b)
}

// func fnn1(n int) {
// 	c := 0
// 	for i := 1; i <= n; i++ {
// 		if i > 1 {
// 			c += 2
// 			// fmt.Println(c, 1111111)
// 			// fmt.Println(2*n-2-c, 22222222)
// 			fmt.Printf("%s%d", strings.Repeat(" ", 2*n-2-c), i)
// 		} else {
// 			fmt.Println(strings.Repeat(" ", 2*n-3), i)
// 		}
// 		for j := i; j > 1; j-- {
// 			fmt.Print(strings.Repeat(" ", 1), j-1)
// 		}
// 		if i == 1 {
// 			continue
// 		}
// 		fmt.Println()
// 	}
// }

func fnn2(n int) {

	r := 3*n - 1
	for i := 1; i <= n; i++ {
		fmt.Printf("%[1]*d", r, i)
		r -= 3
		for j := i; j > 1; j-- {
			fmt.Printf("%3d", j-1)
		}
		fmt.Println()
	}
}

func main() {
	//作业一
	//方法1
	// fmt.Println(fn1(10))
	// //方法2
	// fmt.Println(fn2(10))
	// //方法3 a代表阶乘数，b给定初始值1
	// fmt.Println(fn3(10, 1))

	//作业二
	fnn2(12)
}
