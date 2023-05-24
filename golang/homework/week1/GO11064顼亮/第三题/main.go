package main

import (
	"fmt"
)

func main() {
	var s []int
	s = append(s, 1)
	s = append(s, 1)
	for i := 1; s[i] < 100; i++ {
		if s[i]+s[i-1] > 100 {
			break
		}
		s = append(s, s[i]+s[i-1])
	}
	fmt.Println(s)
}
