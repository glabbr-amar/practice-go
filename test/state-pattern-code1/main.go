package main

/*ordered -> shipped -> out for delivery -> delivered*/
type packageState interface {
	updateState(context deliveryContext)
}

func main() {
	context := deliveryContext{"amar123", nil}
	newDeliveryContext(context.currentState)
	context.update() // will set the state to ordered
	context.update() // will set the state to shipped
	context.update() // will set the state to out for delivery
	context.update() // will set the state to delivered
}
