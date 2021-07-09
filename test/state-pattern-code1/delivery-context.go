package main

import (
	"fmt"
)

type deliveryContext struct {
	packageId    string
	currentState packageState
}

func (d *deliveryContext) updateState(context *deliveryContext) {
	//var state ordered
	if context.currentState == nil{
		d.currentState = ordered{}
	}
}

/*func newDeliveryContext(currentState packageState) *deliveryContext {
	//current state will be set to ordered
	var state ordered
	if currentState == nil {
		state = ordered{}
	}
	fmt.Printf("currentState in deliveryContext  %T\n",state)
	return &deliveryContext{currentState: state}
}
*/
func (d *deliveryContext) update() {
	//will update the currentState of deliveryContext
	fmt.Printf("current state in update : %T\n",d.currentState)
	d.currentState.updateState(&deliveryContext{currentState: d.currentState})
}

func setCurrentState(currentState packageState) *deliveryContext {
	fmt.Printf("current state in setCurrentState : %T\n",currentState)
	return &deliveryContext{currentState: currentState}
}
