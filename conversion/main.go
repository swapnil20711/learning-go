package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome to our pizza app")

	fmt.Println("Please rate our pizza between 1 and 5")

	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')

	fmt.Println("Thanks for rating,", input)

	numRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Added 1 to your rating: ", numRating+1)
	}

	var num int = 42
	fmt.Println("Number is:", num)
	fmt.Printf("Type of number is %T\n", num)

	var data float64 = float64(num)
	data = data + 1.23
	fmt.Println("Data is:", data)
	fmt.Printf("Type of data is %T\n", data)

	num = 123
	str := strconv.Itoa(num)

	fmt.Println("Str is:", str)
	fmt.Printf("Type of Str is %T\n", str)

	number_string := "1234"
	number_int, _ := strconv.Atoi(number_string)

	number_int = number_int + 1231233

	fmt.Println("number_int is:", number_int)
	fmt.Printf("Type of number_int is %T\n", number_int)

	num_string := "3.14"
	number_float, _ := strconv.ParseFloat(num_string, 64)
	fmt.Println("number_float is ", number_float)
	fmt.Printf("Type of number_float is %T\n", number_float)
}
