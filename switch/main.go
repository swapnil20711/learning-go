package main

import "fmt"

func main() {
	day := 10
	switch day % 7 {
	case 0:
		fmt.Println("Sunday")
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	case 4:
		fmt.Println("Thrusday")
	case 5:
		fmt.Println("Friday")
	case 6:
		fmt.Println("Saturday")
	default:
		fmt.Println("Unknown day")
	}

	month := "February"

	switch month {
	case "January", "February", "March":
		fmt.Println("Winter")
	case "April", "May", "June":
		fmt.Println("Spring")
	default:
		fmt.Println("Other season")
	}

	temperature := -10

	switch {
	case temperature < 0:
		fmt.Println("Freezing")

	case temperature >= 0 && temperature < 10:
		fmt.Println("Cold")

	case temperature >= 10 && temperature < 20:
		fmt.Println("Cool")

	case temperature >= 20 && temperature < 30:
		fmt.Println("Warm")

	default:
		fmt.Println("Hot")
	}

}
