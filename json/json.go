package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Email string `json:"email"`
}

type Person struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	IsAdult bool   `json:"isAdult"`
}

func main() {
	// u := User{Name: "Alice", Age: 30, Email: "alice@example.com"}
	// jsonData, err := json.Marshal(u)

	// if err != nil {
	// 	fmt.Println("Error marshaling:", err)
	// 	return
	// }

	// fmt.Println(string(jsonData))

	// jsonString := `{"name":"Bob","age":25}`
	// var u1 User
	// err1 := json.Unmarshal([]byte(jsonString), &u1)

	// if err1 != nil {
	// 	fmt.Println("Error unmarshaling:", err)
	// 	return
	// }

	// fmt.Printf("Parsed Struct: %+v\n", u1)

	person := Person{Name: "Alice", Age: 30, IsAdult: true}

	fmt.Println("Person data is : ", person)

	personJson, err := json.Marshal(person)
	if err != nil {
		fmt.Println("Error marshaling:", err)
	}
	fmt.Println("Person json data is : ", string(personJson))

	// Unmarshalling data
	var personData Person
	err = json.Unmarshal(personJson, &personData)

	if err != nil {
		fmt.Println("Error unmarshaling:", err)
		return
	}

	fmt.Println("Person data is : ", personData)

}
