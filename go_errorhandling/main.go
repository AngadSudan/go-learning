package main

import (
	"errorhandling/helper"
	"fmt"
)

func main() {
	var num1 int
	var num2 int

	fmt.Print("Enter first number: ")
	fmt.Scan(&num1)
	fmt.Print("Enter second number: ")
	fmt.Scan(&num2)

	ans, err := helper.Divide(float64(num1), float64(num2))
	if err != nil {
		fmt.Println(ans)
	} else {
		fmt.Println(err)
	}
}
