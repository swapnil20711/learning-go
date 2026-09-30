package main

import (
	"fmt"
)

func add(a, b int) int {
	return a + b
}

func main() {
	fmt.Println("Starting of the program")
	data := add(5, 6)
	defer fmt.Println(data)
	defer fmt.Println("Middle of the program")
	fmt.Println("End of the program")

	// maintains in stack
	// defer fmt.Println(data) , defer fmt.Println("Middle of the program")
	// lifo will be followed so Middle of the program will be printed then after that data would be printed
}
