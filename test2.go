// Person Class
package main

import "fmt"

func main() {
	a := Person{name: "Doron", age: 40}
	a.sayHi()
	a.printAge()

	b := Employee{
		Person{name: "Asdf", age: 20},
		empId:  1,
		title:  "manager",
		salary: 13000.12}
	fmt.Println(b)
}

type Person struct {
	name string
	age  int
}

type Employee struct {
	Person
	empId  int
	title  string
	salary float64
}

func (p *Person) sayHi() {
	fmt.Println("Hi " + p.name)
}

func (p *Person) printAge() {
	fmt.Printf("%d years.", p.getAge())
}

func (p *Person) getAge() int {
	return p.age
}
