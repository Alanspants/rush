package main

import (
	"fmt"
	"strings"
)

var phoneButton = map[string]string{
	"2": "abc",
	"3": "def",
	"4": "ghi",
	"5": "jkl",
	"6": "mno",
	"7": "pqrs",
	"8": "tuv",
	"9": "wxyz",
}

func dfs(str []string, path string, res *[]string, index int) {
	if len(path) == len(str) {
		*res = append(*res, path)
		return
	}
	currentNum := str[index]
	currentWord := strings.Split(phoneButton[currentNum], "")

	for _, v := range currentWord {
		path += v
		dfs(str, path, res, index+1)
		path = path[:len(path)-1]
	}
}

func letterCombinations(digits string) []string {
	if digits == "" {
		return []string{""}
	}
	str := strings.Split(digits, "")
	var res []string
	dfs(str, "", &res, 0)
	return res
}

func main() {
	fmt.Println(letterCombinations("23"))
}
