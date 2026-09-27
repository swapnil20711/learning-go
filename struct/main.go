package main

import "fmt"

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

type Contact struct {
	Email string
	Phone string
}

type Address struct {
	House   int
	Area    string
	State   string
	PinCode int
}

type Employee struct {
	Person_Details Person
	Person_Contact Contact
	Person_Address Address
}

func main() {
	var person Person
	// fmt.Println("Person is:", person)
	person.FirstName = "Swapnil"
	person.LastName = "Bhojwani"
	person.Age = 25

	// fmt.Println("Person is:", person)

	// 2nd method
	person1 := Person{
		FirstName: "Akash",
		LastName:  "Sharma",
		Age:       25,
	}

	fmt.Println("Person 1 is:", person1)

	// new keyword
	var person2 = new(Person)

	person2.FirstName = "Virat"
	person2.LastName = "Kohli"
	person2.Age = 37

	// due to pointer & is printed
	// fmt.Println("Person 2 is:", person2)

	// fmt.Println("Age of Swapnil is :", person.Age)

	employee := Employee{
		Person_Details: person,
		Person_Contact: Contact{
			Email: "swapnil.bhojwani@gmail.com",
			Phone: "+91xxxxxxxxx",
		},
		Person_Address: Address{
			House: 306,
			State: "Karnataka",
			Area:  "HSR layout",
		},
	}

	employee.Person_Contact.Email = "swap.bhojwani@gmail.com"
	employee.Person_Address.PinCode = 560102

	fmt.Println("Employee is : ", employee)
}
