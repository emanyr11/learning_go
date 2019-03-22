package main

import (
	"fmt"
	"math"
)
/*
List of Methods you can implement on different structs
*/

// Define  the interface
type Shape interface {
	area() float64
}

type Circle struct {
	x, y, radius float64
}

type Rectangle struct {
	length, height float64
}

//Value reciever function for Circle struct called area that returns a float 64
func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

//Value reciever function for Rectangle struct called area that returns a float 64
func (r Rectangle) area() float64 {
	return r.length * r.height
}

//Take in the interface shape
func getArea(s Shape) float64 {
	return s.area()
}

func main() {
	circle := Circle{x: 0, y: 0, radius: 5}
	rectangle := Rectangle{length: 10, height: 5}

	fmt.Printf("Circle Area: %f \n", getArea(circle))
	fmt.Printf("Rectangle Area: %f \n", getArea(rectangle))
}
