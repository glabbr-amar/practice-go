package main

import (
	"log"
)

func main() {
	rectangle := rectangle{5, 6}
	circle := circle{7}
	square := square{4}
	log.Println("area : ", rectangle.area(), circle.area(), square.area())
	log.Println("perimeter : ", rectangle.perimeter(), circle.perimeter(), square.perimeter())

	shapes := []shape{rectangle, circle, square}
	log.Println("printing shapes(slice) area and perimeter using function")
	for _, data := range shapes {
		log.Println(findArea(data), findPerimeter(data))
	}
}
