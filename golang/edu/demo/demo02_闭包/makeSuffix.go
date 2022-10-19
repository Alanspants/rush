package main

import (
	"fmt"
	"strings"
)

func makeSuffix(suffix string) func(string) string {
	return func(filename string) string {
		if strings.HasSuffix(filename, suffix) {
			return filename
		} else {
			return filename + suffix
		}
	}
}

func main() {
	var str string
	fmt.Scan(&str)
	f := makeSuffix(".jpg")
	fmt.Println(f(str))
}
