package main

import (
	"fmt"
)

type NamedObj struct {
	Name string
}

type Shape struct {
	NamedObj  //inheritance
	color     int32
	isRegular bool
}

func (r *Rectangle) calculateArea() float64 {
	return r.Height * r.Width
}

func (r *Rectangle) calculatePerimeter() float64 {
	return 2*r.Height + 2*r.Width
}

type Point struct {
	x, y float64
}

type Rectangle struct {
	NamedObj            //multiple inheritance
	Shape               //^^
	center        Point //standard composition
	Width, Height float64
}

func main() {

	var aRect Rectangle = Rectangle{NamedObj{"name1"},
		Shape{NamedObj{"name2"}, 0, true},
		Point{0, 0},
		10, 10}

	fmt.Println(aRect.Name)
	fmt.Println(aRect.Shape)
	fmt.Println(aRect.Shape.Name)
	fmt.Println(aRect.calculateArea())
	fmt.Println(aRect.calculatePerimeter())
}
