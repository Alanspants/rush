package main

import (
	"fmt"
	"magedu.com/week9/link"
)

func main() {
	// homework-1
	//timecalc.TimeOps()

	// one way linked list
	//owl := link.NewOneWayLinkedList()
	//fmt.Printf("%T %+[1]v\n", owl)
	//
	//owl.Append(10).Append(20).Append(30)
	//fmt.Printf("%T %+[1]v\n", owl)
	//owl.IterOneWayNodes()

	// two way linked list
	twl := link.NewTwoWayLinkedList()
	fmt.Printf("%T %+[1]v\n", twl)

	twl.Append(40).Append(50).Append(60).Append(70)
	fmt.Printf("%T %+[1]v\n", twl)

	// two way linked list iter nodes
	fmt.Println(twl.IterTwoWayNodes(false))
	fmt.Println(twl.IterTwoWayNodes(true))

	// two way linked list remove
	//err := twl.Remove(4)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//err = twl.Remove(3)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//err = twl.Remove(0)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//err = twl.Remove(0)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//err = twl.Remove(0)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//err = twl.Remove(0)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//fmt.Println(twl.IterTwoWayNodes())

	// two way linked list insert
	//err := twl.Insert(-1, 1)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//err = twl.Insert(30, 300)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//err = twl.Insert(0, 100)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//err = twl.Insert(4, 400)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//fmt.Println(twl.IterTwoWayNodes())

	// two way linked list pop
	//twl.Pop()
	//twl.Pop()
	//twl.Pop()
	//twl.Pop()

	//pop1, err := twl.Pop()
	//if err != nil {
	//	fmt.Println(err)
	//}
	//pop2, err := twl.Pop()
	//if err != nil {
	//	fmt.Println(err)
	//}
	//pop3, err := twl.Pop()
	//if err != nil {
	//	fmt.Println(err)
	//}
	//pop4, err := twl.Pop()
	//if err != nil {
	//	fmt.Println(err)
	//}
	//fmt.Println(pop1, pop2, pop3, pop4)
	//pop5, err := twl.Pop()
	//if err != nil {
	//	fmt.Printf("error: %T %[1]s\n", err)
	//}
	//fmt.Println(pop5)
}
