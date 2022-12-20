package main

import (
	"linkList/linkNode"
)

func main() {
	head := linkNode.CreateList([]int{1, 2, 3, 4, 5})
	linkNode.PrintList(&head)
}
