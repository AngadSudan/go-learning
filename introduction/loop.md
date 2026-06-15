## Loops in GO

In GO, `for` loop is the primary way to create loops. The basic structure of for is similar to other languages, but there's only one form of for loop that covers variety of use cases.

Syntax 
```go
for i:=0;i<5;i++ {
    fmt.Println(i);
}
```

Infinite loop using `for`
```go
counter:=0
for {
    fmt.Println("Infinite Loop")
    counter++
    if counter==4 {
        break
    }
}
```


Iterating over a slice/iterable 

```go
numbers := []int{1,2,3,4,5}
for index,value := range numbers {
    fmt.Println(index,value)
}
```


