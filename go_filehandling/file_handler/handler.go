package filehandler

import (
	"fmt"
	"io"
	"os"
)

func WriteToFile(name string) bool {
	file, err := os.Create(name)
	if err != nil {
		fmt.Println("error in creating file")
		return false
	}

	content := "This is the content of the file created using GO"
	_, errs := io.WriteString(file, content)

	if errs != nil {
		fmt.Println("Error writing to the file")
		return false
	}

	defer file.Close()
	return true
}

func ReadFromFile(name string) (string, error) {
	file, err := os.Open(name)
	if err != nil {
		fmt.Println("couldnot open file")
		return "", fmt.Errorf("error in opening file")
	}
	defer file.Close()

	buffer := make([]byte, 1024)
	var content string
	for {
		n, e := file.Read(buffer)
		if e == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("couldnot read file")
			return "", fmt.Errorf("error in reading file")
		}

		content = string(buffer[:n])
	}

	return content, nil
}
