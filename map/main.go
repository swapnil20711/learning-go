package main

import "fmt"

func main() {
	//name->grades
	grades := make(map[string]int)

	fmt.Println(grades["Swapnil"])

	grades["Swapnil"] = 34
	grades["Prince"] = 100
	grades["Alice"] = 90
	grades["Bob"] = 85
	grades["Charlie"] = 95

	fmt.Println(grades["Bob"])

	grades["Bob"] = 100

	fmt.Println(grades["Bob"])

	delete(grades, "Bob")

	fmt.Println(grades["Bob"])

	grade, exists := grades["David"]
	fmt.Println("Grade of David is :", grade)
	fmt.Println("David exists :", exists)

	Grade, Exists := grades["Swapnil"]
	fmt.Println("Grade of Swapnil is :", Grade)
	fmt.Println("Swapnil exists :", Exists)

	for index, value := range grades {
		fmt.Printf("Key is : %s and marks is %d\n", index, value)
	}

	persons := map[string]int{
		"Alice":   90,
		"Bob":     23,
		"Charlie": 32,
	}

	fmt.Println(persons)

	for index, value := range persons {
		fmt.Printf("Key is : %s and age is %d\n", index, value)
	}

}
