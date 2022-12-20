package sellOrder

import (
	"sort"
	"temp1/buyOrder"
	"temp1/seller"
)

type SellOrder struct {
	ID          int
	ContainerID int
	Num         int
	SellPrice   int
	Seller      seller.Seller
}

func (sellOrder *SellOrder) FindBestOrder(buyOrders []buyOrder.BuyOrder) *buyOrder.BuyOrder {
	sort.Slice(buyOrders, func(i, j int) bool {
		return buyOrders[i].BuyPrice*buyOrders[i].Num > buyOrders[j].BuyPrice*buyOrders[j].Num
	})
	for _, order := range buyOrders {
		if order.ContainerID == sellOrder.ContainerID && order.Num == sellOrder.Num {
			sellTotalPrice := sellOrder.Num * sellOrder.SellPrice
			orderTotalPrice := order.Num * order.BuyPrice
			if sellTotalPrice < orderTotalPrice {
				return &order
			}
		}
	}
	return &buyOrder.BuyOrder{}
}
