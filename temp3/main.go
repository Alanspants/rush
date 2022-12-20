package main

import (
	"fmt"
	"temp3/order"
	"temp3/voyage"
)

func main() {
	var voyage0 voyage.Voyage
	voyage0 = voyage.Voyage{
		0,
		10,
		5,
		10,
		"Beijing",
		"Shanghai",
	}

	var voyage1 voyage.Voyage
	voyage0 = voyage.Voyage{
		1,
		8,
		4,
		9,
		"Beijing",
		"Shanghai",
	}

	var order0 order.Order
	order0 = order.Order{
		0,
		5,
		3,
		6,
		"Beijing",
		"Shanghai",
	}

	bestVoyage := order0.FindBestVoyage([]voyage.Voyage{voyage0, voyage1})
	if bestVoyage.VoyageID == -1 {
		fmt.Println("There is no suitable voyage.")
	} else {
		fmt.Println("The most suitable voyage is voyageID:", bestVoyage.VoyageID)
	}
}
