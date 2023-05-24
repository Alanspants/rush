package mian

import (
	"fmt"
)

//定义循环次数
var num = 9

func main() {
	//九九乘法口诀实现
	for i := 1; i < 10; i++ {
		for j := 1; j <= i; j++ {
			if j == 1 {
				fmt.Printf("%d * %d = %-2d ", j, i, j*i)
			} else {
				fmt.Printf("%d * %d = %-3d ", j, i, j*i)
			}
		}
		fmt.Println()
	}
}
