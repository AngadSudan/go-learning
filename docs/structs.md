## Structs in GO

`struct` or short form of structure is a composite datatype in GO that groups together variables under a single name. 
Each field in a struct can have a different datatype. 

Declaring a struct:

```go
type Person struct {
    firstname string
    lastname string
    Age int
}
```

Creating an instance of the struct: 

```go
var person Person
person.firstname = "John"
person.lastname = "Doe"
person.Age = 40

person2:= Person {
    firstname:"John",
    lastname:"Doe",
    Age:40
}

person3 := new(Person)
person3.firstname = "John"
person3.lastname = "Doe"
person3.Age = 40


```


