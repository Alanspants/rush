package main

import "fmt"

func main() {

	for a := 0; a < 10; a++ {

		if a == 1 {
			fmt.Println("1x1=", 1*a)
		} else if a == 2 {
			fmt.Println("1x2=", 1*a, "2x2=", 2*a)
		} else if a == 3 {
			fmt.Println("1x3=", 1*a, "2x3=", 2*a, "3x3=", 3*a)
		} else if a == 4 {
			fmt.Println("1x4=", 1*a, "2x4=", 2*a, "3x4=", 3*a, "4x4=", 4*a)
		} else if a == 5 {
			fmt.Println("1x5=", 1*a, "2x5=", 2*a, "3x5=", 3*a, "4x5=", 4*a, "5x5=", 5*a)
		} else if a == 6 {
			fmt.Println("1x6=", 1*a, "2x6=", 2*a, "3x6=", 3*a, "4x6=", 4*a, "5x6=", 5*a, "6x6=", 6*a)
		} else if a == 7 {
			fmt.Println("1x7=", 1*a, "2x7=", 2*a, "3x7=", 3*a, "4x7=", 4*a, "5x7=", 5*a, "6x7=", 6*a, "7x7=", 7*a)
		} else if a == 8 {
			fmt.Println("1x8=", 1*a, "2x8=", 2*a, "3x8=", 3*a, "4x8=", 4*a, "5x8=", 5*a, "6x8=", 6*a, "7x8=", 7*a, "8x8=", 8*a)
		} else if a == 9 {
			fmt.Println("1x9=", 1*a, "2x9=", 2*a, "3x9=", 3*a, "4x9=", 4*a, "5x9=", 5*a, "6x9=", 6*a, "7x9=", 7*a, "8x9=", 8*a, "9x9=", 9*a)
		}
	}

}

// 这个题目考察的是循环的知识，在复习以后再尝试写一下。
