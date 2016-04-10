package main

import (
	"fmt"
)

func main() {
	// using the parent class
	mp := messagePrinter{"foo"}
	mp.printMessage()
	//Using enhancedMessagePrinter - composition
	emp := enhancedMessagePrinter{}
	emp.message = "bar"
	emp.printMessage()
	//composition
	emp2 := enhancedMessagePrinter{messagePrinter{"bar2"}}
	emp2.printMessage()
}

type messagePrinter struct {
	message string
}

func (mp *messagePrinter) printMessage() {
	fmt.Println(mp.message)
}

type enhancedMessagePrinter struct {
	messagePrinter
}
