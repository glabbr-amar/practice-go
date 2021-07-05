package main

import (
	"math"
)

// interface implementing same function
type shape interface {
	area() float32
	perimeter() float32
}
// methods for rectangle, circle and square
func (r rectangle) area() float32 {
	return r.length * r.breadth
}
func (r rectangle) perimeter() float32 {
	return (r.length + r.breadth) * 2
}

func (c circle) area() float32{
	return math.Pi * c.radius * c.radius
}
func (c circle) perimeter() float32 {
	return 2 * math.Pi * c.radius
}

func (s square) area() float32 {
	return s.side * s.side
}
func (s square) perimeter() float32{
	return 4*s.side
}

// structure for rectangle, circle and square
type rectangle struct{
	length,breadth float32
}
type circle struct {
	radius float32
}
type square struct {
	side float32
}

func findArea(s shape) float32 {
	return s.area()
}
func findPerimeter(s shape) float32 {
	return s.perimeter()
}
