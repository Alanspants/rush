package main

import "fmt"

func fi(n int) (result int) {
	if n == 0 {
		result = 1
	} else if n == 1 {
		result = 1
	} else {
		result = fi(n-1) + fi(n-2)
	}
	return result
}

func main() {

	var result int
	for i := 0; i < 100; i++ {
		result = fi(i)
		if result <= 100 {
			fmt.Println(result)
		} else {
			break
		}
	}

}
