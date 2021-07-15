package main

import "log"

type ordered struct {
}

func (o ordered) updateState(context *deliveryContext) {
	log.Println("ordered")
	// will go to shipped
	context.currentState = setCurrentState(shipped{})
	//dctx := &deliveryContext{}
	//dctx.SetCurrentState(shipped{})


}