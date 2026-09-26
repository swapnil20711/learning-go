package main

import "fmt"

func main() {
	fmt.Println("We are learning array in golang")

	var names [5]string
	names[0] = "Swapnil"
	names[1] = "Prince"

	fmt.Println("Names of persons is :", names)

	var numbers = [5]int{1, 2, 3, 4, 5}

	fmt.Println("Number is : ", numbers)
	fmt.Println("Length of numbers array is : ", len(numbers))

	fmt.Println("Value of name at 1st index is:", names[1])

	var name [5]string
	name[2] = "prince"
	name[0] = "swapnil"
	fmt.Println("Value of prices is :", name)
	fmt.Printf("Value of prices is : %q\n", name)
}
