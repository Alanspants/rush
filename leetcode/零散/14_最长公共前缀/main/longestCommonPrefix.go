package main

import "fmt"

func longestCommonPrefix(strs []string) string {
	var shareStr string = strs[0]
	for index := 1; index < len(strs); index++ {
		var temp string = ""
		for i, j := 0, 0; i < len(shareStr) && j < len(strs[index]); i, j = i+1, j+1 {
			if shareStr[i] == strs[index][j] {
				temp += string(shareStr[i])
			} else {
				break
			}
		}
		shareStr = temp
	}
	return shareStr
}

func main() {
	var n int
	fmt.Print("字符串个数: ")
	fmt.Scan(&n)

	var str []string
	for i := 0; i < n; i++ {
		var temp string
		fmt.Scan(&temp)
		str = append(str, temp)
	}

	fmt.Println(longestCommonPrefix(str))
}
