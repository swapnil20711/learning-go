package main

import (
	"fmt"
	"os"
)

func main() {
	/*file, err := os.Create("example.txt")

	if err != nil {
		fmt.Println("Error while creating file!", err)
		return
	}

	defer file.Close()

	content := "Hello world by swapnil!"
	byte, err1 := io.WriteString(file, content+"\n")
	fmt.Println("Byte writtern is: ", byte)
	if err1 != nil {
		fmt.Println("Error while writing file!", err)
		return
	}
	fmt.Println("Successfully created a file")
	*/

	/*file, err := os.Open("example.txt")

	if err != nil {
		fmt.Println("Error while opening file!", err)
		return
	}

	defer file.Close()

	buffer := make([]byte, 1024)

	// Read the file content into the buffer

	for {
		n, err := file.Read(buffer)

		if err == io.EOF {
			break
		}

		if err != nil {
			fmt.Println("Error while reading file!", err)
			return
		}

		fmt.Println(string(buffer[:n]))
	}
	*/

	content, err := os.ReadFile("example.txt")

	if err != nil {
		fmt.Println("Error while reading file!", err)
		return
	}
	fmt.Println(string(content))

}
