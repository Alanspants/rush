package main

import (
	"container/heap"
	"fmt"
)

type order struct {
	price  int
	amount int
}

type orderHeap struct {
	value    []order
	heapType string
}

func (h orderHeap) Len() int {
	return len(h.value)
}

func (h orderHeap) Swap(i, j int) {
	h.value[i], h.value[j] = h.value[j], h.value[i]
}

func (h orderHeap) Less(i, j int) bool {
	if h.heapType == "small" {
		return h.value[i].price < h.value[j].price
	} else {
		return h.value[j].price < h.value[i].price
	}
}

func (h *orderHeap) Push(x interface{}) {
	h.value = append(h.value, x.(order))
}

func (h *orderHeap) Pop() interface{} {
	x := h.value[len(h.value)-1]
	h.value = h.value[:len(h.value)-1]
	return x
}

// -------------------------------------

func getNumberOfBacklogOrders(orders [][]int) int {
	if len(orders) == 0 {
		return 0
	}

	buyHeap := &orderHeap{
		heapType: "big",
	}
	sellHeap := &orderHeap{
		heapType: "small",
	}

	for _, o := range orders {
		currentOrder := order{
			o[0],
			o[1],
		}
		if o[2] == 0 {
			// buy order
			for len(sellHeap.value) > 0 && currentOrder.price >= sellHeap.value[0].price && currentOrder.amount != 0 {
				if sellHeap.value[0].amount > currentOrder.amount {
					sellHeap.value[0].amount -= currentOrder.amount
					currentOrder.amount = 0
				} else {
					currentOrder.amount -= sellHeap.value[0].amount
					sellHeap.value[0].amount = 0
					heap.Pop(sellHeap)
				}
			}
		} else if o[2] == 1 {
			// sell order
			for len(buyHeap.value) > 0 && currentOrder.price <= buyHeap.value[0].price && currentOrder.amount != 0 {
				if buyHeap.value[0].amount > currentOrder.amount {
					buyHeap.value[0].amount -= currentOrder.amount
					currentOrder.amount = 0
				} else {
					currentOrder.amount -= buyHeap.value[0].amount
					buyHeap.value[0].amount = 0
					heap.Pop(buyHeap)
				}
			}
		}
		if currentOrder.amount != 0 {
			if o[2] == 0 {
				heap.Push(buyHeap, currentOrder)
			} else {
				heap.Push(sellHeap, currentOrder)
			}
		}
		fmt.Println(buyHeap)
	}

	res := 0
	for _, o := range buyHeap.value {
		res += o.price
		res %= 1000000007
	}

	for _, o := range sellHeap.value {
		res += o.price
		res %= 1000000007
	}

	return res
}

func main() {
	order := [][]int{
		{10, 5, 0},
		{15, 2, 1},
		{25, 1, 1},
		{30, 4, 0},
	}
	fmt.Println(getNumberOfBacklogOrders(order))
}
