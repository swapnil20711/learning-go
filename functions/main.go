package main

import "fmt"

func simpleFunction() {
	fmt.Println("simple function")
}

// func add(a int, b int) int {
// 	return a + b
// }

func add(a int, b int) (result int) {
	return a + b
}

func multiply(a, b int) (result int) {
	result = a * b
	return
}

func main() {
	fmt.Println("We are learning functions in golang")
	simpleFunction()
	sum := add(2, 3)
	fmt.Println(sum)

	data := multiply(2, 3)
	fmt.Println("Multiplication of two number is:", data)
}
