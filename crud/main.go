package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Todo struct {
	UserId    int    `json:"userId"`
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func main() {
	fmt.Println("Learning CRUD....")

	res, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")

	if err != nil {
		fmt.Println("Error getting : ", err)
		return
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		fmt.Println("Error in getting res : ", res.Status)
		return
	}

	// data, error := io.ReadAll(res.Body)

	// if error != nil {
	// 	fmt.Println("Error reading : ", error)
	// 	return
	// }

	// fmt.Println("Data is : ", string(data))

	var todo Todo

	error := json.NewDecoder(res.Body).Decode(&todo)
	if error != nil {
		fmt.Println("Error decoding data : ", error)
		return
	}

	fmt.Println("Todo: ", todo)

	fmt.Println("title is : ", todo.Title)
	fmt.Println("completed is : ", todo.Completed)
}
