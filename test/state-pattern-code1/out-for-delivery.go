package main

import "log"

type outForDelivery struct {
}

func (o outForDelivery) updateState(context *deliveryContext) {
	log.Println("out for delivery")
	// will go to delivered
	context.currentState = setCurrentState(delivered{})
	//dctx := &deliveryContext{}
	//dctx.SetCurrentState(delivered{})
}
