package main

import (
	"fmt"
	"math"
	"sort"
)

func getNumberOrBacklogOrders(orders [][]int) int {
	// [order_price, order_number]
	buyOrders := [][]int{}
	sellOrders := [][]int{}

	for _, order := range orders {
		orderPrice := order[0]
		orderAmount := order[1]
		orderType := order[2]

		tempBuyOrderNum := 0
		tempSellOrderNum := 0

		for i := 0; i < orderAmount; i++ {
			if orderType == 0 {
				if len(sellOrders) == 0 {
					tempBuyOrderNum += orderAmount
					break
				} else {
					//if orderPrice >= sellOrders[0][0] {
					//	sellOrders[0][1] -= 1
					//	if sellOrders[0][1] == 0 {
					//		sellOrders = sellOrders[1:len(sellOrders)]
					//	}
					//} else {
					//	tempBuyOrderNum++
					//}
					if orderPrice >= sellOrders[0][0] {
						if sellOrders[0][1] >= orderAmount {
							sellOrders[0][1] -= orderAmount
						} else {
							orderAmount -= sellOrders[0][1]
							//i += sellOrders[0][1]
							sellOrders = sellOrders[1:len(sellOrders)]
						}
					} else {
						tempBuyOrderNum += orderAmount
						break
					}
				}
			} else if orderType == 1 {
				if len(buyOrders) == 0 {
					tempSellOrderNum++
				} else {
					if orderPrice <= buyOrders[0][0] {
						buyOrders[0][1] -= 1
						if buyOrders[0][1] == 0 {
							buyOrders = buyOrders[1:len(buyOrders)]
						}
					} else {
						tempSellOrderNum++
					}
				}
			}
		}

		fmt.Println("tempBuyOrderNum: ", tempBuyOrderNum)
		fmt.Println("tempSellOrderNum: ", tempSellOrderNum)
		fmt.Println("-----")

		if orderType == 0 {
			buyOrder := []int{orderPrice, tempBuyOrderNum}
			if tempBuyOrderNum == 0 {
				continue
			}
			buyOrders = append(buyOrders, buyOrder)
			sort.Slice(buyOrders, func(i, j int) bool {
				return buyOrders[i][0] > buyOrders[j][0]
			})
		} else if orderType == 1 {
			sellOrder := []int{orderPrice, tempSellOrderNum}
			if tempSellOrderNum == 0 {
				continue
			}
			sellOrders = append(sellOrders, sellOrder)
			sort.Slice(sellOrders, func(i, j int) bool {
				return sellOrders[i][0] < sellOrders[j][0]
			})
		}

	}

	sellOrderNum := 0
	for _, order := range sellOrders {
		sellOrderNum += order[1]
	}

	buyOrderNum := 0
	for _, order := range buyOrders {
		buyOrderNum += order[1]
	}

	ans := (sellOrderNum + buyOrderNum) % (int(math.Pow(10, 9)) + 7)
	return ans
}

func main() {
	//orders := [][]int{{10, 5, 0}, {15, 2, 1}, {25, 1, 1}, {30, 4, 0}}
	orders := [][]int{{7, 1000000000, 1}, {15, 3, 0}, {5, 999999995, 0}, {5, 1, 1}}
	fmt.Println(orders)
	fmt.Println(getNumberOrBacklogOrders(orders))
}
