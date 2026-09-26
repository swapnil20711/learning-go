package main

import "fmt"

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("denominator must not be zero")
	}
	return a / b, nil
}

// func divide(a, b float64) (float64, string) {
// 	if b == 0 {
// 		return 0, "denominator must not be zero"
// 	}
// 	return a / b, "nil"
// }

func main() {
	fmt.Println("Error handling in the function")
	// if we want to ignore error we can use _
	// ans, _ := divide(10, 2)
	ans, err := divide(10, 2)

	if err == nil {
		fmt.Println("Division of two numbers is  :", ans)
	}

}
