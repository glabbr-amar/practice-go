package main

import (
	"log"
)

func main() {
	defer log.Println("done in main")
	for i := 0; i < 100; i++ {
		go GetInstance()
	}
}
