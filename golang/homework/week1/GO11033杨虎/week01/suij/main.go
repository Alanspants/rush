package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	r1 := rand.New(rand.NewSource(time.Now().UnixNano()))
	res_list := []int{}

	for i := 0; i < 20; i++ {
		res := r1.Intn(100)
		if res == 0 {
			res = r1.Intn(100)
		}
		// fmt.Printf("%d\n", r1.Intn(100))

		res_list = append(res_list, res)

	}
	fmt.Printf("res_list: %v\n", res_list)

	sum_list := []int{}
	sheng_list := []int{}

	for k, v := range res_list {
		// fmt.Println(k, v)

		if k%2 == 0 {

			sum_list = append(sum_list, v)
		} else {

			sheng_list = append(sheng_list, v)
		}

	}

	fmt.Printf("sum_list: %v\n", sum_list)
	fmt.Printf("sheng_list: %v\n", sheng_list)

	sum := 0
	for _, v := range sum_list {
		// fmt.Printf("v: %v\n", v)
		sum = sum + v

	}
	fmt.Printf("sum: %v\n", sum)

	sheng := 1
	for _, v := range sheng_list {
		// fmt.Printf("v: %v\n", v)
		sheng = sheng * v
	}
	fmt.Printf("sheng: %v\n", sheng)

}
