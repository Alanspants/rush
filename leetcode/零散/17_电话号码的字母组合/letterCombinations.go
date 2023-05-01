package main

import "fmt"

var letters = map[string]string{
	"2": "abc",
	"3": "def",
	"4": "ghi",
	"5": "jkl",
	"6": "mno",
	"7": "pqrs",
	"8": "tuv",
	"9": "wxyz",
}

func dfs(digits string, ans *[]string, path string, index int) {
	if len(digits) == len(path) {
		*ans = append(*ans, path)
		return
	}

	digit := string(digits[index])
	for i := 0; i < len(letters[digit]); i++ {
		letter := string(letters[digit][i])
		path += letter
		dfs(digits, ans, path, index+1)
		path = path[:len(path)-1]
	}
}

func letterCombinations(digits string) []string {
	if digits == "" {
		return []string{}
	}
	ans := []string{}
	dfs(digits, &ans, "", 0)
	return ans
}

func main() {
	fmt.Println(letterCombinations("23"))
}
