package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Todo struct {
	UserID    int64  `json:"userId"`
	ID        int64  `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func main() {
	fmt.Println("Learning web services...")

	response, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")

	if err != nil {
		fmt.Println("Error getting GET response")
		return
	}

	defer response.Body.Close()

	fmt.Printf("Type of response is : %T\n", response)
	// fmt.Println("Response : ", response)

	data, err := io.ReadAll(response.Body)

	if err != nil {
		fmt.Println("Error while reading res", err)
		return
	}

	fmt.Println("response : ", string(data))

	var todo Todo

	json.Unmarshal(data, &todo)

	fmt.Println(todo.Title)
}
