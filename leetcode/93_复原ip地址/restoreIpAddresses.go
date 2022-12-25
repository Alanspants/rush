package main

import (
	"fmt"
	"strconv"
	"strings"
)

func validation(str string, dotNum int) bool {
	if dotNum == 0 {
		return true
	}
	splitStr := strings.Split(str, ".")
	for _, s := range splitStr {
		if len(s) > 1 && len(s) <= 3 && s[0] == '0' {
			return false
		}
		if dotNum == 3 {
			if len(s) > 3 || len(s) <= 0 {
				return false
			} else if v, _ := strconv.Atoi(s); v > 255 {
				return false
			}
		} else {
			if (3-dotNum+1)*3 <= len(s) {
				return false
			}
		}
	}
	return true
}

func dfs(s string, path string, ans *[]string, index int) {

	if len(path) == (len(s)+3) && validation(path, 3) {
		*ans = append(*ans, path)
		return
	}

	for i := index; i <= index+3; i++ {
		if (len(path)-len(s)) > 3 || !validation(path, len(path)-len(s)) {
			continue
		}
		//temp := path
		path = path[:i] + "." + path[i:len(path)]
		dfs(s, path, ans, index+4)
		path = path[:i] + path[i+1:len(path)]
	}
}

func restoreIpAddresses(s string) []string {
	ans := []string{}
	dfs(s, s, &ans, 1)
	ansMap := map[string]int{}

	for _, str := range ans {
		ansMap[str] = 1
	}

	ans = []string{}
	for k, _ := range ansMap {
		ans = append(ans, k)
	}
	return ans
}

func main() {
	fmt.Println("\n", restoreIpAddresses("101023"))
	//fmt.Println(validation("255.255.11135", 2))
}
