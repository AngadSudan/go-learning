# Functions in GO

In go functions are declared using the `func` keyword. followed by the name of the function and parameters if any, return type if any and the function body.

The usual syntax for a function with no return type is as follows
```go
func function_name(){}
```

input parameters are given in the following way 
```go
func function_name(name string){
    // function body
}
```

output params or the return types are given as follows
```go
func function_1(age int) bool {}
```

## Error handling inside Go 
Usually in go you can return multiple datatypes and that's how you usually take care of the edge cases. Lets assume a function given by divide. 
If the second number becomes zero then the output must return to error. 

Inorder to tackle that our function returns 2 values 
```go
func Divide(num1, num2 float64) (float64, error) {
	if num2 == 0 {
		return 0, fmt.Errorf("Demoninator can not be zero")
	}

	return num1 / num2, nil
}
```

`error` is a keyword in go used to determine an error and when returning multiple values you can return values seperated by , or nil values.

Here now when you are going to use it, it can be done in the following way.
```go
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
```

But its not guranteed that you will always use all the values returned by the function. As per GO, you must use all variables declared in the code file. So GO introduces a special character `_` underscore as a blank identifier.

It serves as a write-only variable that allows you to discard values returned by a function or to ignore specific values when you are not interested to use them. 

## Arrays in GO

Arrays in GO provide a simple and elegant way to work with a fixed collection of elements.

Arrays are declared using the following syntax - `var name [size]type`.
They have fixed size meaning you must specify the number of elements array hold. 

Inorder to `initialize` arrays with given values you can use `var numbers = [5]int{1,2,3,4,5}`
Inorder to `access` values for an array we can use `numbers[<index>] = <value>`
Inorder to determine the `length` of the array we can use `len(arr)`.


```
In Go, arrays or slices, the elements are intialized to their zero value. The zero value is the default value assigned to a variables of a certain type when no explicit value is provided.
```


## Slices in GO

in GO slice are a more flexible and dynamic data structure(under the hood an array) that provides a more powerful alternatiev to arrays.
Slices are dynamic in length and their value can be increased during runtime as per need. 

declaration of slice are as follow
```go
number := []int{1,2,3,4,5}
```

`make` function in slices is what we can use to create a slice with specific length and capacity.
`append` function is another function that can be used to add elements to slice.

In slices we get 2 terms one of them is `length` and another is `capacity`.
`Length` - the total number of values that are currently present in a slice.
`Capacity` - the maximum number of values that can be stores in the slice. For capacity we have the function ``` cap(<slice_variable>) ```


 
Creating a slice using make function ``` number:= make([]int, <length>, <capacity>) ``` 
Appending a value into the slice ``` number = append(number,<value>) ```

if at a condition where `length == capacity` another value is entered, then it increases the capacity of the slice to almost double and then increases the length of the slice.

 
