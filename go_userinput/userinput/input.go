package userinput

import (
	"bufio"
	"fmt"
	"os"
)

func ReadUserName(name string) {
	fmt.Print("Enter your name: ")
	fmt.Scan(&name)

	fmt.Println("hello there, ", name, "!")
	fmt.Printf("hello there, %s!\n", name)
}

func ReadUserStatement(statement string) {
	reader := bufio.NewReader(os.Stdin)
	statement, _ = reader.ReadString('\n') // read till you get a \n
	fmt.Println(statement)
}
