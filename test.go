package main

import (
	"fmt"
)

const (
	message string = "This is a contstant"
	first   string = "this is first"
	num0
	num1 = iota
	num2
	num3
)

func main() {
	printCollection()
	var myInt int

	myInt = 9
	myInt++

	printString()
	printComplex()
}

func printString() {
	myString := "Hello Go"
	fmt.Println(myString)
}

func printComplex() {
	myComplex := complex(2, 3)
	fmt.Println(myComplex)
	fmt.Println(real(myComplex))
	fmt.Println(imag(myComplex))

}

func printCollection() {
	// Array
	myArray := [10]int{}
	myArray[0] = 42
	myArray[1] = 11
	fmt.Println(myArray[1])

	newArray := [...]int{1, 2, 3, 4, 5}
	mySlice := newArray[:]
	fmt.Println(mySlice)
}
