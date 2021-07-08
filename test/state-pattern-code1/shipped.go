package main

import "log"

type shipped struct {
}

func (s shipped) updateState(context deliveryContext) {
	log.Println("shipped")
	// will go to out for delivery
	setCurrentState(outForDelivery{})
}


