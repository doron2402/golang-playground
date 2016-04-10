package main

func main() {
	message := "First"

	changeMsgByRef(&message)
	println("Should be second: ", message)

	sayHello("Hello", "World")
}

func changeMsgByRef(message *string) {
	println("Should be First: ", *message)
	*message = "Second"
}

func sayHello(messages ...string) {
	for _, msg := range messages {
		println(msg)
	}
}
