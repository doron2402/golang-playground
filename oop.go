package main

import (
	"fmt"
)

func main() {
	foo := newMyStruct()
	foo.myMap["bar"] = "baz"
	fmt.Println(foo)

	doron := newPerson()
	doron.friends["rachel"] = "1234"
	fmt.Println(doron)
}

type myStuct struct {
	myMap map[string]string
}

func newMyStruct() *myStuct {
	result := myStuct{}
	result.myMap = map[string]string{}
	return &result
}

type person struct {
	name    string
	friends map[string]string
}

func newPerson() *person {
	result := person{}
	result.friends = map[string]string{}
	return &result
}
