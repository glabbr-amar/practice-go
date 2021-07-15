package main

import "log"

type shipped struct {
}

func (s shipped) updateState(context *deliveryContext) {
	log.Println("shipped")
	// will go to out for delivery
	context.currentState = setCurrentState(outForDelivery{})
	//dctx := &deliveryContext{}
	//dctx.SetCurrentState(outForDelivery{})
}


