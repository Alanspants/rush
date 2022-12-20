package order

type Order interface {
	FindBestOrder([]Order) Order
	GetOrderType() string
	GetOrderID() int
	GetOrderContainerID() int
	GetOrderNum() int
	GetOrderPrice() int
}
