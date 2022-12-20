package main

import (
	"fmt"
	"temp1/buyOrder"
	"temp1/buyer"
	"temp1/sellOrder"
	"temp1/seller"
)

func main() {
	seller0 := &seller.Seller{
		ID:    0,
		Money: 1000,
	}

	//seller1 := &seller.Seller{
	//	ID:    1,
	//	Money: 1000,
	//}

	buyer0 := &buyer.Buyer{
		ID:    0,
		Money: 1000,
	}

	//buyer1 := &buyer.Buyer{
	//	ID:    1,
	//	Money: 1000,
	//}

	sellOrder0 := &sellOrder.SellOrder{
		ID:          0,
		ContainerID: 0,
		Num:         2,
		SellPrice:   10,
		Seller:      *seller0,
	}

	sellOrder1 := &sellOrder.SellOrder{
		ID:          1,
		ContainerID: 0,
		Num:         2,
		SellPrice:   11,
		Seller:      *seller0,
	}

	buyOrder0 := &buyOrder.BuyOrder{
		ID:          0,
		ContainerID: 0,
		Num:         2,
		BuyPrice:    12,
		Buyer:       *buyer0,
	}

	sellOrders := []sellOrder.SellOrder{*sellOrder0, *sellOrder1}

	bestOrder := buyOrder0.FindBestOrder(sellOrders)

	fmt.Println(bestOrder.ID)
}
