package main

import "fmt"

func main() {
	// declaration of an array
	var arr [5]int
	// initialization of values
	arr[0] = 1
	arr[1] = 2
	arr[2] = 3
	arr[3] = 4
	arr[4] = 5

	// declare an array and directly initialize the values
	var arr2 = [5]int{1, 2, 3, 4, 5}
	fmt.Println(arr2)

	//length of array
	var n int = len(arr2)
	fmt.Println("length is :", n)
}
