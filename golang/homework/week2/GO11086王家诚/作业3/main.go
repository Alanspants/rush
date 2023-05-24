package main

import "fmt"

//打印斐波那契数列数列100以内
//特点① 从第三项开始每一项都等于前两项之和
//② Fn=F(n-1)+F(n-2)
func main() {
	a1 := 1 //斐波那契数列第一项
	a2 := 1 //斐波那契数列第二项
	an := 0
	for i := 1; i < 100; i++ {
		if i == 1 || i == 2 { //先输出出第一项和第二项  最特殊的两项
			an = 1
			fmt.Println("an=", an)
		} else if i >= 3 && an < 80 { // 当循环进入第三项开始时  第三项的和等于前两项相加   an<80 是为了固定出最后一项的值
			a3 := a1 + a2
			a1 = a2
			a2 = a3
			an = a3
			fmt.Println("an=", an)
		}
	}
}

// func main() {
// 	a1 := 1 //斐波那契数列第一项
// 	a2 := 1 //斐波那契数列第二项
// 	an := 0
// 	for i := 1; i < 100; i++ {
// 		if i == 1 || i == 2 {      //先输出出第一项和第二项  最特殊的两项
// 			an = 1
// 			fmt.Println("an=", an)
// 		} else if i >= 3 && i<100 {     // 当循环进入第三项开始时  第三项的和等于前两项相加  打印出100以内的斐波那契数列
// 			a3 := a1 + a2
// 			a1 = a2
// 			a2 = a3
// 			an = a3
// 			fmt.Println("an=", an)
// 		}
// 	}
// }
