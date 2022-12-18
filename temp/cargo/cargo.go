package cargo

import (
	"sort"
	"temp/ship"
)

type Cargo struct {
	Id         int
	Desc       string
	Weight     int
	TargetCity string
}

func (cargo *Cargo) PickAvailableShip(shipList []*ship.Ship) []*ship.Ship {

	var availableShipList []*ship.Ship

	for _, ship := range shipList {
		shipAvailableWeight := ship.WeightTotal - ship.WeightCurrent
		if ship.TargetCity == cargo.TargetCity && shipAvailableWeight >= cargo.Weight {
			availableShipList = append(availableShipList, ship)
		}
	}

	sort.Slice(availableShipList, func(i, j int) bool {
		timeI := availableShipList[i].Arrive - availableShipList[i].Departure
		timeJ := availableShipList[j].Arrive - availableShipList[j].Departure
		return timeI > timeJ
	})

	return availableShipList
}
