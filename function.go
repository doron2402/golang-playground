package main

func main() {
	result := add(1, 2, 3, 4, 5)
	println(result)

	println(add2(8))
}

func add(terms ...int) int {
	result := 0
	for _, term := range terms {
		result += term
	}
	return result
}

func add2(num int) int {
	return num + 2
}
