package main

import "fmt"

// 打印100以内的斐波那契数列
//1, 1,  2, 3, 	5, 	8, 	13, 21, 34, 55, 89
//a  b  c
//   a  b   c
//      a   b   c
//          a   b  c
// 		        a  b  c
//
//  a+b=c=0+1=1; a=b=1; b=1 ;a没用了
//
//  a+b=c=1+1=2; a=b=1; b=2
//
//	a+b=c=1+2=3; a=b=2; b=3
//
//  a+b=c=2+3=5; a=b=3; b=5

//  a+b=c=3+5=8; a=b=5; b=8
func main() {
	a := 0
	b := 1
	for i := 1; ; i++ {
		c := a + b
		a = b
		b = c
		if c > 100 {
			break
		} else {
			fmt.Println(c)
		}
	}
}

// func main() {
// 	a := 0
// 	b := 1
// 	// var n int
// 	for i := 0; i < 100; i++ {
// 		a = a + b
// 		b = b + a
// 		if b >= 100 {
// 			break
// 		}
// 		fmt.Println(a)
// 		fmt.Println(b)

// 	}
// }
