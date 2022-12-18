package main

import (
	"fmt"
	"temp/cargo"
	"temp/ship"
)

func main() {
	ship0 := &ship.Ship{
		Id:            0,
		Departure:     1,
		Arrive:        4,
		TargetCity:    "Beijing",
		WeightCurrent: 5,
		WeightTotal:   10,
	}

	ship1 := &ship.Ship{
		Id:            1,
		Departure:     1,
		Arrive:        6,
		TargetCity:    "Beijing",
		WeightCurrent: 7,
		WeightTotal:   10,
	}

	ship2 := &ship.Ship{
		Id:            2,
		Departure:     3,
		Arrive:        7,
		TargetCity:    "Beijing",
		WeightCurrent: 2,
		WeightTotal:   6,
	}

	ship3 := &ship.Ship{
		Id:            3,
		Departure:     1,
		Arrive:        7,
		TargetCity:    "Shanghai",
		WeightCurrent: 1,
		WeightTotal:   6,
	}

	ship4 := &ship.Ship{
		Id:            4,
		Departure:     2,
		Arrive:        7,
		TargetCity:    "Shanghai",
		WeightCurrent: 4,
		WeightTotal:   9,
	}

	ship5 := &ship.Ship{
		Id:            5,
		Departure:     2,
		Arrive:        5,
		TargetCity:    "Shanghai",
		WeightCurrent: 2,
		WeightTotal:   5,
	}

	Ships := []*ship.Ship{ship0, ship1, ship2, ship3, ship4, ship5}

	cargo0 := &cargo.Cargo{
		Id:         0,
		Desc:       "Car",
		Weight:     4,
		TargetCity: "Beijing",
	}

	availableShip := cargo0.PickAvailableShip(Ships)
	for _, ship := range availableShip {
		fmt.Println(ship.Id, ship.TargetCity)
	}
}
