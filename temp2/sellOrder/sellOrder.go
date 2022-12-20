package sellOrder

import (
	"sort"
	"temp2/order"
)

type SellOrder struct {
	OrderID      int
	ContainerID  int
	AvailableNum int
	Price        int
}

func (sellOrder SellOrder) GetOrderType() string {
	//TODO implement me
	return "SellOrder"
}

func (sellOrder SellOrder) GetOrderID() int {
	return sellOrder.OrderID
}

func (sellOrder SellOrder) GetOrderContainerID() int {
	return sellOrder.ContainerID
}

func (sellOrder SellOrder) GetOrderNum() int {
	return sellOrder.AvailableNum
}

func (sellOrder SellOrder) GetOrderPrice() int {
	return sellOrder.Price
}

func (sellOrder SellOrder) FindBestOrder(orders []order.Order) order.Order {
	availableOrders := []order.Order{}
	for _, order := range orders {
		if order.GetOrderContainerID() == sellOrder.ContainerID && order.GetOrderNum() <= sellOrder.AvailableNum {
			availableOrders = append(availableOrders, order)
		}
	}
	if len(availableOrders) == 0 {
		return nil
	}
	sort.Slice(availableOrders, func(i, j int) bool {
		return availableOrders[i].GetOrderPrice() > availableOrders[j].GetOrderPrice()
	})
	return availableOrders[0]
}
