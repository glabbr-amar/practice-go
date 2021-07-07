package main

import (
	"log"
	"sync"
)

var lock = &sync.Mutex{}

type singleton struct {
}

var instance *singleton

func getInstance() *singleton {
	if instance == nil {
		lock.Lock()
		defer lock.Unlock()
		if instance == nil {
			log.Println("Creating instance")
			instance = &singleton{}
		}
	} else {
		log.Println("instance already created")
	}
	return instance
}