package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	var url string = "https://jsonplaceholder.typicode.com/todos/1"
	response, err := http.Get(url)
	if err != nil {
		fmt.Println("error in making API Call")
		return
	}

	defer response.Body.Close()

	// reading the data
	content, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("error in reading API Call")
		return
	}

	parsedData := string(content)
	fmt.Println(parsedData)
}
