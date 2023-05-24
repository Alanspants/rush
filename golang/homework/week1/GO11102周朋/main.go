package main

import (
	"fmt"
	"log"
	"math"
	"math/rand"
	"time"
)

func main() {
	// fmt.Println("hello world !")
	// TODO 1,打印乘法口诀表
	for i := 1; i < 10; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%4d * %d = %d", j, i, i*j)
		}
		fmt.Println()
	}
	// const (
	// 	a = iota //0
	// 	b        //1
	// 	c        //2
	// 	_        //3
	// 	d        //4
	// 	e = 10   //5
	// 	f        //6
	// 	g = iota //7
	// 	h        //8
	// )
	// fmt.Println(a, b, c, d, e, f, g, h)

	// TODO 第二题
	seconds := time.Now().Unix()
	rand.Seed(seconds)
	a := 0
	b := 1.0
	for i := 0; i < 20; i++ {
		target := rand.Intn(99) + 1
		// fmt.Println(target)
		if i&1 == 0 {
			b *= float64(target)
		} else {
			a += target
		}
	}
	fmt.Println("--------第二题--------")
	fmt.Println(a)
	fmt.Printf("%.2f\n", b)

	fmt.Println("--------第三题--------")
	// // TODO 第三题
	for i := 0; i < 100; i++ {
		fibSeqValue, err := fibonacciSequence(i)
		if err != nil {
			log.Fatal(err)
		}
		if fibSeqValue > 100 {
			break
		}
		fmt.Println(fibSeqValue)

	}

	// err := errors.New("height can't he negative")
	// fmt.Println(err)
	// fmt.Println(err.Error())
	// root, err := squareRoot(-9.3)
	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Printf("%0.3f", root)
	// }

	// TODO pointer
	// var myInt int
	// fmt.Println(reflect.TypeOf(&myInt))
	// var myFloat float64
	// fmt.Println(reflect.TypeOf(&myFloat))
	// var myBool bool
	// fmt.Println(reflect.TypeOf(&myBool))
	// var myIntPointer *int
	// myIntPointer = &myInt
	// fmt.Println(myIntPointer)
	// var myFloatPointer *float64
	// myFloatPointer = &myFloat
	// fmt.Println(myFloatPointer)

	// var myFloatPointer *float64 = createPointer()
	// fmt.Println(*myFloatPointer)

	// var myBool bool = true
	// printPointer(&myBool)
	// amount := 6
	// double(&amount)
	// fmt.Println(amount)
	// arr2 := [5]int{15, 23, 8, 10, 7}
	// arr1 := bubbleSort(arr2)
	// fmt.Println(arr1)
	// fmt.Println(15&5, 15&^5)
}

func bubbleSort(arr [5]int) [5]int {
	legth := len(arr)
	for i := 0; i < legth; i++ {
		for j := 0; j < legth-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}
func double(number *int) {
	*number *= 2
}

func printPointer(myBoolPointer *bool) {
	fmt.Println(*myBoolPointer)
}
func createPointer() *float64 {
	var myFloat = 98.6
	return &myFloat
}
func squareRoot(number float64) (float64, error) {
	if number < 0 {
		return 0, fmt.Errorf("can't get square root of negative number")
	}
	return math.Sqrt(number), nil

}
func fibonacciSequence(n int) (int, error) {
	if n < 0 {
		return 0, fmt.Errorf("a n of %d is invalid", n)
	} else if n == 0 || n == 1 {
		return n, nil
	} else {
		fb1, err := fibonacciSequence(n - 1)
		if err != nil {
			log.Fatal(err)
		}
		fb2, err := fibonacciSequence(n - 2)
		if err != nil {
			log.Fatal(err)
		}
		return fb1 + fb2, nil
	}

}
