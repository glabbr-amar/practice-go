package main

import "log"

type singleton struct {
}

var instance *singleton

func getInstance() *singleton {

	instance = &singleton{}
	log.Println("instance created")
	return instance
}
