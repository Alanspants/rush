package main

import (
	"fmt"
	"unicode"
)

func dfs(result *[]string, str []byte, path string, visited map[int]int, index int) {
	if len(path) == len(str) {
		*result = append(*result, path)
		return
	}

	val := str[index]
	if visited[index] == 2 {
		return
	}

	if unicode.IsDigit(rune(val)) {
		path += string(val)
		dfs(result, str, path, visited, index+1)
		return
	}

	if _, ok := visited[index]; ok {
		visited[index] += 1
	} else {
		visited[index] = 1
	}
	//fmt.Println(visited)

	path += string(val)
	dfs(result, str, path, visited, index+1)

	path = path[:len(path)-1]
	//fmt.Printf("%c\n", str[index])
	if unicode.IsUpper(rune(val)) {
		str[index] = byte(unicode.ToLower(rune(val)))
	} else {
		str[index] = byte(unicode.ToUpper(rune(val)))
	}
	//fmt.Printf("%c\n", str[index])
	//fmt.Println(str)
	dfs(result, str, path, visited, index)
	visited[index] = 0
}

func letterCasePermutation(s string) []string {
	result := []string{}
	str := []byte(s)
	visited := map[int]int{}
	dfs(&result, str, "", visited, 0)
	return result
}

func main() {
	s := "a1b2"
	fmt.Println(letterCasePermutation(s))
}
