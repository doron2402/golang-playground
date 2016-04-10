package main

import (
	"fmt"
)

func main() {
	foo := 1
	if foo == 1 {
		println("bar")
	} else {
		println("buz")
	}

	if bar := 2; bar == 1 {
		println("bar is 1")
	} else {
		println("bar is 2")
	}

	//Switch
	switch {
	case foo == 1:
		println("one")
	case foo > 2:
		println("two")
	}

	for i := 0; i < 5; i++ {
		println(i)

	}

	j := 0
	for {
		j++
		println(j)
		if j > 10 {
			break
		}
	}

	s := []string{"foo", "bar", "baz"}
	fmt.Println(s)
	m := make(map[string]string)
	m["first"] = "foo"
	m["second"] = "bar"

	for k, v := range m {
		println(k, v)
	}

}
