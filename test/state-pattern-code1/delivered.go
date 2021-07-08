package main

import "log"

type delivered struct {
}

func (d delivered) updateState(context deliveryContext) {
	log.Println("delivered")
}

