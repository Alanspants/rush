package main

import "fmt"

func romanToInt(s string) int {
	ans := 0
	for index := 0; index < len(s); index++ {
		if index != len(s)-1 {
			thisValue := symbolValue(s[index])
			nextValue := symbolValue(s[index+1])
			if thisValue < nextValue {
				ans += nextValue - thisValue
				index++
			} else {
				ans += thisValue
			}
		} else {
			ans += symbolValue(s[index])
		}
	}
	return ans
}

func symbolValue(c byte) int {
	switch c {
	case 'I':
		return 1
	case 'V':
		return 5
	case 'X':
		return 10
	case 'L':
		return 50
	case 'C':
		return 100
	case 'D':
		return 500
	case 'M':
		return 1000
	}
	return 0
}

func main() {
	var s string
	fmt.Scan(&s)

	fmt.Println(romanToInt(s))
}
