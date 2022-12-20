package order

type Order interface {
	FindBestOrder(order Order) (bestOrder Order)
}
