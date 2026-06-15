## File Handling in GO

file handling refers to creation, updation or deletion  of files using a programming language.

to create a file - `file,err :=os.Create("filename")` and it returns the file along with an error.
Now for any file thats ever created/opened you need to make sure to close that file otherwise it causes a memory leak inside the containers. So to close a file that's already opened, we do `file.Close()` usually with a defer keyword.

Now, inorder to write the content to the file we need to use another library in GO which is the `io` library. 
To write to a file we do the following command - `data_byte, err :=io.WriteString(file, content)`. Data_byte here is the number of bytes written in the file.

Here's what the complete code looks like 
```go
package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.Create("random.txt")
	if err != nil {
		fmt.Println("error in creating file")
		return
	}

	content := "This is the content of the file created using GO"
	_, errs := io.WriteString(file, content)

	if errs != nil {
		fmt.Println("Error writing to the file")
		return
	}

	defer file.Close()
}

``` 

## Read Data from a file 

Inorder to read the contents of a file, you need to first open the file using the function `os.Open()` which returns you either an err or an file string and once this is done you call the defer file.Close() since we dont want any memory leaks. Then inorder to store the buffer you create a buffer slice using the make keyword with a size of 1024 as default. 

Then you iterate over the files content until you reach an `EOF` character that signifies the end of the file avoiding a for loop. Once this is done you create a string out of the buffer and return it as on output.

here's what the code look like 
```go
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
```


Additionaly we also have an library using which we can easily read the contents of the file. `content,err := io.ReadFile(filepath)` however it is useful in the scenarios where you want to read entire content of a file into the memory. For larger files we usually prefer the buffer method.


