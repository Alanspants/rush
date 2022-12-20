package buyOrder

import (
	"sort"
	"temp1/buyer"
	"temp1/sellOrder"
)

type BuyOrder struct {
	ID          int
	ContainerID int
	Num         int
	BuyPrice    int
	Buyer       buyer.Buyer
}

func (buyOrder *BuyOrder) FindBestOrder(sellOrders []sellOrder.SellOrder) *sellOrder.SellOrder {
	sort.Slice(sellOrders, func(i, j int) bool {
		return sellOrders[i].SellPrice*sellOrders[i].Num < sellOrders[j].SellPrice*sellOrders[i].Num
	})
	for _, order := range sellOrders {
		if order.ContainerID == buyOrder.ContainerID && order.Num == buyOrder.Num {
			buyTotalPrice := buyOrder.Num * buyOrder.BuyPrice
			orderTotalPrice := order.Num * order.SellPrice
			if buyTotalPrice >= orderTotalPrice {
				return &order
			}
		}
	}
	return &sellOrder.SellOrder{}
}
