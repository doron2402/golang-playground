package main

func main() {
	addFunc := func(terms ...int) (numTerms int, sum int) {
		for _, term := range terms {
			sum += term
		}

		numTerms = len(terms)

		return
	}

	numTerms, sum := addFunc(1, 2, 3, 4)
	println("Added: ", numTerms, " Terms for a total of ", sum)

	add3 := func(a int) (result int) {
		result = a + 3
		return
	}
	println(add3(77743))

}
