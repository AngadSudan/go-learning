# User Input in GO

to take the user input in go we need to utilize the `fmt` package in go. 
The fmt package and use pointer in a way to store the value inside a variable similar to C.

The following coded gives an explanation for the user input.

```go
package main

import "fmt"

func main() {
	var name string
	fmt.Print("Enter your name: ")
	fmt.Scan(&name)

	fmt.Println("hello there, ", name, "!")
	fmt.Printf("hello there, %s!\n", name)
}
```

Here we get out key difference between Println & Printf as well. Println adds an additional space after every , seperated variable name or string value inside of it.

The `Scan` method however only reads till whitespace character so any string beyond a spacebar is discarded from the input. Incase we want to read a whole line then we need to have another package termed as `buffio`. We will use NewScanner or ReadString function for complex scenarios.

then the method turns out something like this

```go
func ReadUserStatement(statement string) {
	reader := bufio.NewReader(os.Stdin)
	statement, _ = reader.ReadString('\n') // read till you get a \n
	fmt.Println(statement)
}
```


