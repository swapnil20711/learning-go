package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type Todo struct {
	UserId    int    `json:"userId"`
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func performGetRequest() {

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

func performPostRequest() {
	todo := Todo{
		UserId:    23,
		Title:     "Swapnil Bhojwani",
		Completed: true,
	}

	// convert todo data in json
	jsonData, err := json.Marshal(todo)

	if err != nil {
		fmt.Println("Error marshalling", err)
		return
	}

	// convert data to string
	jsonString := string(jsonData)

	//covert string to json reader

	jsonReader := strings.NewReader(jsonString)

	myUrl := "https://jsonplaceholder.typicode.com/todos/"

	//send post req
	res, err := http.Post(myUrl, "application/json", jsonReader)

	if err != nil {
		fmt.Println("Error sending req:", err)
		return
	}

	defer res.Body.Close()

	data, _ := io.ReadAll(res.Body)
	fmt.Println("Data is : ", string(data))
}

func performUpdateRequest() {
	todo := Todo{
		UserId:    2323424,
		Title:     "Swapnil Bhojwani Go Lang",
		Completed: false,
	}

	// convert todo data in json
	jsonData, err := json.Marshal(todo)

	if err != nil {
		fmt.Println("Error marshalling", err)
		return
	}

	// convert data to string
	jsonString := string(jsonData)

	//covert string to json reader

	jsonReader := strings.NewReader(jsonString)

	myUrl := "https://jsonplaceholder.typicode.com/todos/1"

	//send post req
	req, err := http.NewRequest(http.MethodPut, myUrl, jsonReader)

	if err != nil {
		fmt.Println("Error sending req:", err)
		return
	}

	req.Header.Set("Content-type", "application/json")

	//send the req

	client := http.Client{}
	res, err := client.Do(req)

	if err != nil {
		fmt.Println("Error sending req: ", err)
	}
	defer res.Body.Close()

	data, _ := io.ReadAll(res.Body)
	fmt.Println("Data is : ", string(data))
	fmt.Println("Status is : ", res.Status)
}

func performDeleteRequest() {
	myUrl := "https://jsonplaceholder.typicode.com/todos/1"

	//create deleted request
	req, err := http.NewRequest(http.MethodDelete, myUrl, nil)

	if err != nil {
		fmt.Println("Error sending req:", err)
		return
	}

	client := http.Client{}
	res, err := client.Do(req)

	if err != nil {
		fmt.Println("Error sending req: ", err)
	}
	defer res.Body.Close()
	fmt.Println("Status is : ", res.Status)

}

func main() {
	fmt.Println("Learning CRUD....")
	// performGetRequest()
	performDeleteRequest()
}
