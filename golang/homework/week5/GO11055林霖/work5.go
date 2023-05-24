package main

import (
	"fmt"
	"strconv"
	"strings"
)

func multi1(num int) int {
	if num <= 0 {
		panic("num 大于0")
	}
	sum := 1
	for i := 1; i <= num; i++ {
		sum = sum * i
	}
	return sum
}

func multi2(num int) int {
	if num <= 1 {
		return 1
	}
	return num * multi2(num-1)
}

func multi3(sum, num int) int {
	if num <= 1 {
		return 1 * sum
	}
	return multi3(sum*num, num-1)
}

func printNum(num int) {
	size := len(strconv.Itoa(num))
	spaceMap := make(map[int]string, size)
	for i := 1; i <= size; i++ {
		spaceMap[i] = strings.Repeat(" ", i)
	}
	for i := 1; i <= num; i++ {
		var str strings.Builder
		for j := num; j > 0; j-- {
			if j > i {
				s, exit := spaceMap[len(strconv.Itoa(j))]
				if !exit {
					panic("error..")
				}
				str.WriteString(s)
				str.WriteString(" ")
			} else {
				str.WriteString(strconv.Itoa(j))
				str.WriteString(" ")
			}
		}

		fmt.Println(strings.TrimSuffix(str.String(), " "))
	}
}

func main() {
	fmt.Println(multi1(10))
	fmt.Println(multi2(10))
	fmt.Println(multi3(1, 10))

	printNum(20)
}
