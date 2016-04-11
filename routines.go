package main

import (
	"runtime"
	"time"
)

func main() {

	runtime.GOMAXPROCS(3)
	go abcGen()

	println("First")

	time.Sleep(100 * time.Millisecond)
}

func abcGen() {
	for l := byte('a'); l <= byte('z'); l++ {
		go println(string(l))
	}
}
