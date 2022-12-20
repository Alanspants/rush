package main

import (
	"fmt"
	"temp2/buyOrder"
	"temp2/order"
	"temp2/sellOrder"
)

func main() {
	var buyOrder0 order.Order
	buyOrder0 = buyOrder.BuyOrder{
		OrderID:     0,
		ContainerID: 1,
		Num:         20,
		Price:       10,
	}

	var sellOrder0 order.Order
	sellOrder0 = sellOrder.SellOrder{
		OrderID:      0,
		ContainerID:  1,
		AvailableNum: 10,
		Price:        8,
	}

	var sellOrder1 order.Order
	sellOrder1 = sellOrder.SellOrder{
		OrderID:      1,
		ContainerID:  1,
		AvailableNum: 15,
		Price:        9,
	}

	order := buyOrder0.FindBestOrder([]order.Order{sellOrder0, sellOrder1})
	if order == nil {
		fmt.Println("there is no suitable order")
	} else {
		fmt.Println(order.GetOrderID())
	}
}
