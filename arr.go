package main

import (
	"fmt"
)

func main() {
	mySlice := make([]float32, 100)
	mySlice[0] = 12.
	mySlice[1] = 15.
	mySlice[2] = 18.
	// mySlice := []float32{12., 15., 18.}
	mapType()
}

func mapType() {
	myMap := make(map[int]string)
	myMap[32] = "Foo"

	fmt.Println(myMap)
	fmt.Println(myMap[99])
	fmt.Println(myMap[32])
}
