package main

import "fmt"
import "runtime"

func main() {
	runtime.GOMAXPROCS(4)
	fmt.Println("Start.")

	ch := make(chan string)

	doneCh := make(chan bool)

	go abcGen(ch)
	go print(ch, doneCh)

	<-doneCh
}

func abcGen(ch chan string) {
	for l := byte('a'); l <= byte('z'); l++ {
		ch <- string(l)
	}

	close(ch)
}

func print(ch chan string, doneCh chan bool) {
	for l := range ch {
		fmt.Println(l)
	}

	doneCh <- true
}
