package main

import (
	"go_userinput/userinput"
)

func main() {
	var name string
	var statement string
	userinput.ReadUserName(name)
	userinput.ReadUserStatement(statement)
}
