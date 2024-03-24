package main

import "fmt"
import "math"

func structs() {
	// set values in structure
	rect1 := Rectangle{height:10, width:10}
	fmt.Println(rect1.height)

	// this worked before the interface was added
	// fmt.Println(area(rect1))

	// or if you know the order of the characteristics
	rect2 := Rectangle{2,4}
	fmt.Println("The area of this rectangle is", getArea(rect2))

	circle := Circle{7}
	fmt.Println("The area of this circle is", getArea(circle))
}

// creating interface
type Shape interface {
	area() float64
}


// define structure
type Rectangle struct {
	height float64
	width float64
}

type Circle struct {
	radius float64
}

// writing custom functions for structures
func (rect Rectangle) area() float64 {
	return rect.height*rect.width
}

func (circle Circle) area() float64 {
	return math.Pi * math.Pow(circle.radius, 2)
}

// when using just structs and not interfaces the * works otherwise it causes error?
// func (circle *Circle) area() float64 {
// 	return math.Pi * math.Pow(circle.radius, 2)
// }

func getArea(shape Shape) float64 {
	return shape.area()
}

