package orderSys

import (
	"sort"
	"temp4/buyOrder"
	"temp4/sellOrder"
)

type OrderSys struct {
	BuyOrders  []buyOrder.BuyOrder
	SellOrders []sellOrder.SellOrder
}

func (orderSys OrderSys) FindBestBuyOrder(sellOrder sellOrder.SellOrder) buyOrder.BuyOrder {
	availableBuyOrder := []buyOrder.BuyOrder{}
	for _, order := range orderSys.BuyOrders {
		if sellOrder.CargoID == order.CargoID && sellOrder.Price < order.Price {
			availableBuyOrder = append(availableBuyOrder, order)
		}
	}

	if len(availableBuyOrder) == 0 {
		return buyOrder.BuyOrder{
			ID: -1,
		}
	} else {
		sort.Slice(availableBuyOrder, func(i, j int) bool {
			return availableBuyOrder[i].Price > availableBuyOrder[j].Price
		})
	}

	return availableBuyOrder[0]
}
