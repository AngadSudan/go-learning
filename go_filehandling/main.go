package main

import (
	filehandler "filehandling/file_handler"
	"fmt"
)

func main() {
	fileCreated := filehandler.WriteToFile("write.txt")
	if fileCreated {
		fmt.Println("file creation was successful")
	} else {
		fmt.Println("file creation had an error")
	}

	content, err := filehandler.ReadFromFile("random.txt")
	if err != nil {
		fmt.Println("data reading had an error")
	} else {
		fmt.Println(content)
	}
}
