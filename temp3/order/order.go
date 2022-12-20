package order

import (
	"sort"
	"temp3/voyage"
)

type Order struct {
	OrderID     int
	WeightTotal int
	StartTime   int
	EndTime     int
	StartCity   string
	EndCity     string
}

func (order Order) FindBestVoyage(voyages []voyage.Voyage) voyage.Voyage {
	availableVoyages := []voyage.Voyage{}
	for _, voyage := range voyages {
		flag := true
		if order.WeightTotal > voyage.Capacity {
			flag = false
		}
		if order.StartTime > voyage.DepartureTime {
			flag = false
		}
		if order.EndTime < voyage.ArrivalTime {
			flag = false
		}
		if order.StartCity != voyage.DepartureCity && order.EndCity != voyage.ArrivalCity {
			flag = false
		}
		if flag {
			availableVoyages = append(availableVoyages, voyage)
		}
	}

	if len(availableVoyages) == 0 {
		return voyage.Voyage{
			VoyageID: -1,
		}
	}

	sort.Slice(availableVoyages, func(i, j int) bool {
		return availableVoyages[i].ArrivalTime < availableVoyages[j].ArrivalTime
	})

	return availableVoyages[0]

}
