package main

import (
	"fmt"
)

type Stack struct {
	inner []byte
}

func (s *Stack) Push(v byte) {
	s.inner = append(s.inner, v)
}

func (s *Stack) Pop() byte {
	n := len(s.inner) - 1
	v := s.inner[n]
	s.inner = s.inner[:n]
	return v
}

func (s *Stack) Peek() byte {
	var ans byte
	if len(s.inner) == 0 {
		return ans
	}
	n := len(s.inner) - 1
	return s.inner[n]
}

func isValid(s string) bool {
	var stack Stack
	for i := 0; i < len(s); i++ {
		tmp := string(stack.Peek())
		if tmp == "(" && string(s[i]) == ")" {
			stack.Pop()
		} else if tmp == "{" && string(s[i]) == "}" {
			stack.Pop()
		} else if tmp == "[" && string(s[i]) == "]" {
			stack.Pop()
		} else {
			stack.Push(s[i])
		}
	}
	if len(stack.inner) != 0 {
		return false
	}

	return true
}

func main() {
	var str string
	fmt.Scan(&str)
	fmt.Println(isValid(str))
}
