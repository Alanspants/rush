package main

import (
	"fmt"
)

func main() {
	for i := 1; i < 10; i++ {
		for j := 1; j <= i; j++ {
			if i == j {
				fmt.Printf("%d*%d=%d\n", j, i, j*i)
			} else if j == 1 {
				fmt.Printf("%d*%d=%-2d", j, i, j*i)
			} else {
				fmt.Printf("%d*%d=%-3d", j, i, j*i)
			}
		}
	}
}

// 输出时可以统一把格式调宽，不需要针对几类元素进行调整。这样代码就更加简洁。可以尝试一下
