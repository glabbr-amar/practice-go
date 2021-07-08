package main

import (
	"fmt"
)

type deliveryContext struct {
	packageId    string
	currentState packageState
}

func newDeliveryContext(currentState packageState) *deliveryContext {
	//current state will be set to ordered
	var state ordered
	if currentState == nil {
		state = ordered{}
	}
	fmt.Printf("currentState in deliveryContext  %T",currentState)
	return &deliveryContext{currentState: state}
}

func (d *deliveryContext) update() {
	//will update the currentState of deliveryContext
	d.currentState.updateState(deliveryContext{currentState: d.currentState})
}

func setCurrentState(currentState packageState) *deliveryContext {
	return &deliveryContext{currentState: currentState}
}
