package main

import "fmt"

func modifyValueByReference(num *int) {
	*num = *num + 20
}

func main() {
	// var num int
	// num = 2

	// var ptr *int
	// ptr = &num

	num := "swapnil"
	ptr := &num

	// fmt.Println("Num has value:", num)
	fmt.Println("Ptr contains:", ptr)
	fmt.Println("Data contains through pointer:", *ptr)

	var pointer *int

	if pointer == nil {
		fmt.Println("Pointer is not assigned")
	}

	value := 10
	modifyValueByReference(&value)

	fmt.Println("Value contains :", value)

}
