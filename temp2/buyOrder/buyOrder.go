package buyOrder

import (
	"sort"
	"temp2/order"
)

type BuyOrder struct {
	OrderID     int
	ContainerID int
	Num         int
	Price       int
}

func (buyOrder BuyOrder) GetOrderType() string {
	//TODO implement me
	return "BuyOrder"
}

func (buyOrder BuyOrder) GetOrderID() int {
	return buyOrder.OrderID
}

func (buyOrder BuyOrder) GetOrderContainerID() int {
	return buyOrder.ContainerID
}

func (buyOrder BuyOrder) GetOrderNum() int {
	return buyOrder.Num
}

func (buyOrder BuyOrder) GetOrderPrice() int {
	return buyOrder.Price
}

func (buyOrder BuyOrder) FindBestOrder(orders []order.Order) order.Order {
	availableOrders := []order.Order{}
	for _, order := range orders {
		if order.GetOrderContainerID() == buyOrder.ContainerID && order.GetOrderNum() >= buyOrder.Num {
			availableOrders = append(availableOrders, order)
		}
	}
	if len(availableOrders) == 0 {
		return nil
	}
	sort.Slice(availableOrders, func(i, j int) bool {
		return availableOrders[i].GetOrderPrice() < availableOrders[j].GetOrderPrice()
	})
	return availableOrders[0]
}
