package main

import (
	"fmt"
	"time"
)

func sayHello() {
	fmt.Println("Hello, World!")
	// time.Sleep(2000 * time.Millisecond)
	fmt.Println("Say hello function ended")
}

func sayHi() {
	fmt.Println("Hi Swapnil")
	time.Sleep(1000 * time.Millisecond)
	fmt.Println("Hi Swapnil function ended")
}
func main() {
	fmt.Println("Learning goroutines")

	go sayHello()
	go sayHi()

	time.Sleep(900 * time.Millisecond)

}
