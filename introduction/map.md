## Maps in GO

in GO, `map` is a datastructure that provides an unordered collection of key-value pair where each key is essentially unique.
similar to `dictionaries` or `hashmap` in other programming languages. 

They are used to associate value with keys along with their effective retreival.

For creation of map we also use the `make` keyword.
```go
grades:= make(map[string]int)) 
```

the key datatype must be inside the brackets and the returned value datatype must be outside those brackets.

`Inserting` a value inside a map - `grades["student1"] = 78`
`Accessing` a value in the map - `grades["student1"]`
`Deleting` a value in the map - `delete(grades,"student1")`

A map usually returns a default value incase the key doesn't exist. Inorder to tackle this, usually maps in GO returns two values. One of them being the actual value and other being a variable which tells whether the key exists or not. 

`check` whether a key exists in a map we usually do
```go
value,exists := grades["student92"]

if exists {
    fmt.Println(value)
}else {
    fmt.Println("doesn't exist")
}
```


Declaring a map with values directly 
```go
person := map[string]int{
    "a":1,
    "b":2,
    "c":3
}
```

