package main

import (
	"fmt"
)

type person struct {
	name   string
	age    int
	isMale bool
}

func Person(name string, age int, gender bool) *person {
	p := new(person)
	p.name = name
	p.age = age
	p.isMale = gender
	return p
}

func (p *person) getAge() int {
	return p.age
}

func (p *person) getGender() string {
	if person.isMale == true {
		return "man"
	}
	return "woman"
}

func main() {
	doron := Person("Doron", 30, true)
	fmt.Println(doron)
	fmt.Println("The gender is " + doron.getGender())
}
