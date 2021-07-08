package main

import (
	"log"
	"sync"
)

type singleton struct {
}

var (
	instance *singleton
	once     sync.Once
)

func getInstance() *singleton {
	once.Do(func() {
		instance = &singleton{}
		log.Println("instance created")
	})
	return instance
}
