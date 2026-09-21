package main

import (
	"fmt"
	"log"

	"swapnil.com/greetings"
)

func main() {
	// Set properties of the predefined Logger, including
	// the log entry prefix and a flag to disable printing
	// the time, source file, and line number.
	log.SetPrefix("greetings: ")
	log.SetFlags(0)
	names := []string{"Gladys", "Samantha", "Darrin"}
	messages, err := greetings.Hellos(names)
	if err != nil {
		// If an error was returned, print it to the console and
		// exit the program.
		log.Fatal(err)
	}
	fmt.Println(messages)
}
