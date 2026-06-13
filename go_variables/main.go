package main

import "fmt"

func main() {
	// strict type declaration
	var name string = "variable_data"
	var version float64 = 1.0
	var money int = 75000
	var isAdult bool = true
	// variable type declaration
	var version_latest = "1.0"
	// declare constant values
	const pi = 3.14
	// pi = 3.16  - throws error

	//automatic type inference
	person := "me"
	fmt.Println(name)
	fmt.Println(version)
	fmt.Println(version_latest)
	fmt.Println(isAdult)
	fmt.Println("account balance: ", money)
	fmt.Println(person)
}
