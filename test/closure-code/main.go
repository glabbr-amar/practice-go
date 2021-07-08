package main

import "log"

func evenGenerator() func() int {
	i:=0
	log.Println("Hello world")
	return func() int {
		i = i + 2
		return i
	}
}
func main() {
	log.Println("This is Amar Kumar")
	// nextEven is basically storing the function returned by evenGenerator
	nextEven := evenGenerator()
	log.Println(nextEven())
	log.Println(nextEven())
	log.Println(nextEven())
	log.Println(nextEven())
}
