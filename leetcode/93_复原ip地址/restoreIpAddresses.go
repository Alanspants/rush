package main

import (
	"fmt"
	"strconv"
	"strings"
)

func dfs(s string, path string, ans *[]string, index int) {

	//fmt.Println("dfs: ", path, *ans)

	if len(path) == (len(s) + 3) {
		flag := true

		splitPath := strings.Split(path, ".")
		for _, s := range splitPath {
			num, _ := strconv.Atoi(s)
			if num > 255 {
				flag = false
			} else if len(s) > 1 && s[0] == '0' {
				flag = false
			}
		}
		if flag {
			*ans = append(*ans, path)
		}
		return
	}

	for i := index; i < index+3 && i < (len(path)); i++ {
		currentDotNum := len(path) - len(s) + 1

		//fmt.Print("origin: ", path, " i: ", i, " [")

		//fmt.Printf("origin: %s, i: %d, length: %d [", path, i, len(path))
		if currentDotNum <= 3 {
			if (len(path) - i) > (3-currentDotNum+1)*3 {
				//fmt.Println("kill]")
				continue
			}
		}

		//fmt.Println("ok]")

		temp := path
		path := path[:i] + "." + path[i:]

		//fmt.Println("after: ", path)
		dfs(s, path, ans, i+2)
		path = temp
	}

}

func restoreIpAddresses(s string) []string {
	ans := []string{}
	dfs(s, s, &ans, 1)
	return ans
}

func main() {
	s := "25525511135"
	fmt.Println(restoreIpAddresses(s))
}
