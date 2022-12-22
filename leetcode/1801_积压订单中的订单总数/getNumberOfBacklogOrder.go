package main

import (
	"fmt"
	"math"
	"sort"
)

func getNumberOfBacklogOrders(orders [][]int) int {
	// [orderPrice, orderAmount]
	buyBacklog := [][]int{}
	sellBacklog := [][]int{}

	for _, order := range orders {
		orderPrice := order[0]
		orderAmount := order[1]
		orderType := order[2]

		// buy
		if orderType == 0 {
			if len(sellBacklog) == 0 {
				buyBacklog = append(buyBacklog, []int{orderPrice, orderAmount})
			} else {
				for _, sellOrder := range sellBacklog {
					if sellOrder[0] <= orderPrice {
						if sellOrder[1] > orderAmount {
							sellOrder[1] -= orderAmount
							break
						} else {
							sellBacklog = sellBacklog[1:len(sellBacklog)]
							orderAmount -= sellOrder[1]
							if len(sellBacklog) == 0 && orderAmount > 0 {
								buyBacklog = append(buyBacklog, []int{orderPrice, orderAmount})
							}
							continue
						}
					} else {
						buyBacklog = append(buyBacklog, []int{orderPrice, orderAmount})
						break
					}
				}
			}
			sort.Slice(buyBacklog, func(i, j int) bool {
				return buyBacklog[i][0] > buyBacklog[j][0]
			})
		} else if orderType == 1 {
			if len(buyBacklog) == 0 {
				sellBacklog = append(sellBacklog, []int{orderPrice, orderAmount})
			} else {
				for _, buyOrder := range buyBacklog {
					if buyOrder[0] >= orderPrice {
						if buyOrder[1] > orderAmount {
							buyOrder[1] -= orderAmount
							break
						} else {
							buyBacklog = buyBacklog[1:len(buyBacklog)]
							orderAmount -= buyOrder[1]
							if len(buyBacklog) == 0 && orderAmount > 0 {
								sellBacklog = append(sellBacklog, []int{orderPrice, orderAmount})
							}
							continue
						}
					} else {
						sellBacklog = append(sellBacklog, []int{orderPrice, orderAmount})
						break
					}
				}
			}
			sort.Slice(sellBacklog, func(i, j int) bool {
				return sellBacklog[i][0] < sellBacklog[j][0]
			})
		}
	}

	sellOrderNum := 0
	for _, order := range sellBacklog {
		sellOrderNum += order[1]
	}

	buyOrderNum := 0
	for _, order := range buyBacklog {
		buyOrderNum += order[1]
	}

	ans := (sellOrderNum + buyOrderNum) % (int(math.Pow(10, 9)) + 7)
	return ans
}

func main() {
	//orders := [][]int{{10, 5, 0}, {15, 2, 1}, {25, 1, 1}, {30, 4, 0}}
	//orders := [][]int{{7, 1000000000, 1}, {15, 3, 0}, {5, 999999995, 0}, {5, 1, 1}}
	orders := [][]int{{19, 28, 0}, {9, 4, 1}, {25, 15, 1}}
	fmt.Println(orders)
	fmt.Println(getNumberOfBacklogOrders(orders))
}
