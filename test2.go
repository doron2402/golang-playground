// Person Class
package main

import "fmt"

func main() {
	a := Person{name: "Doron", age: 40}
	a.sayHi()
	a.printAge()
}

type Person struct {
	name string
	age  int
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
