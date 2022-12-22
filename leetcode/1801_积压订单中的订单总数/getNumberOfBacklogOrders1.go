package main

import (
	"fmt"
	"math"
	"sort"
)

func getNumberOfBacklogOrders1(orders [][]int) int {
	buyOrders := []int{}
	sellOrders := []int{}

	for _, order := range orders {
		orderPrice := order[0]
		orderAmount := order[1]
		orderType := order[2]

		fmt.Println(orderPrice, orderAmount, orderType)

		for i := 0; i < orderAmount; i++ {
			if orderType == 0 {
				if len(sellOrders) == 0 {
					buyOrders = append(buyOrders, orderPrice)
				} else {
					if orderPrice > sellOrders[0] {
						sellOrders = sellOrders[1:len(sellOrders)]
					} else {
						buyOrders = append(buyOrders, orderPrice)
					}
				}
			} else if orderType == 1 {
				if len(buyOrders) == 0 {
					sellOrders = append(sellOrders, orderPrice)
				} else {
					if orderPrice < buyOrders[0] {
						buyOrders = buyOrders[1:len(buyOrders)]
					} else {
						sellOrders = append(sellOrders, orderPrice)
					}
				}
			}
		}
		if orderType == 0 {
			sort.Slice(buyOrders, func(i, j int) bool {
				return buyOrders[i] > buyOrders[j]
			})
		} else {
			sort.Slice(sellOrders, func(i, j int) bool {
				return sellOrders[i] < sellOrders[j]
			})
		}
	}

	backlogOrders := (len(sellOrders) + len(buyOrders)) % (int(math.Pow(10, 9)) + 7)

	return backlogOrders
}

func main() {
	orders := [][]int{{10, 5, 0}, {15, 2, 1}, {25, 1, 1}, {30, 4, 0}}
	fmt.Println(orders)
	fmt.Println(getNumberOfBacklogOrders(orders))
}
