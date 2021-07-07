package main

import (
	"log"
)

func main() {
	sum := add(1, 2, 3)
	log.Println("sum of three numbers 1,2 and 3 is ", sum)
	aSlice := []int{1, 2, 3, 4, 5}
	sum = add(aSlice...)
	log.Println("sum of all number in slices ", sum)
}
