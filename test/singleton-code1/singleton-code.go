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

func GetInstance() *singleton {
	once.Do(func() {
		instance = &singleton{}
		log.Println("instance created")
	})
	return instance
}
