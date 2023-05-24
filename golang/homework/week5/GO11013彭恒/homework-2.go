package main

import "fmt"

// 打印空格
// 参数n是输出空格的个数
func printBlank(n int) {
	for i := 0; i < n; i++ {
		fmt.Printf(" ")
	}
}

// 输出上三角形
// 参数n是输出上三角形的参数，大于0，小于100
func upperTriangle(n int) {
	if n <= 0 {
		fmt.Println("wrong param:", n, ", please param greater than 0, less than 100")
		return
	} else if n > 99 {
		fmt.Println("wrong param:", n, ", please param greater than 0, less than 100")
		return
	}
	for i := 0; i < n; i++ {
		for j := i + 1; j > 0; j-- {
			if j == 1 && i == 0 {
				if n > 9 && n < 100 {
					printBlank((n - 9) * 3)
					printBlank(8 * 2)
					fmt.Printf("%d\n", j)
				} else if n <= 9 {
					printBlank((n - 1) * 2)
					fmt.Printf("%d\n", j)
				}
			} else if j == 1 {
				fmt.Printf("%d\n", j)
			} else if j == i+1 {
				if n <= 9 {
					printBlank((n - j) * 2)
					fmt.Printf("%d ", j)
				} else if n > 9 && n < 100 {
					if j <= 9 {
						printBlank(n - 9)
						printBlank((n - j) * 2)
						fmt.Printf("%d ", j)
					} else if j > 9 {
						printBlank((n - j) * 3)
						fmt.Printf("%d ", j)
					}
				}
			} else {
				fmt.Printf("%d ", j)
			}
		}
	}
}

func main() {
	upperTriangle(15)
}

// 批改意见
// 代码逻辑都是对的，记得写一些注释，要不代码太长后面自己看的时候都看不懂了。
// 计算宽度的逻辑有点复杂，可以改进一下，例如每行取最后一行宽度即可
