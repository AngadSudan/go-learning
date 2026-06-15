## GO Routines

Goroutines are a key feature of the GO programming language that allow you to run functions concurrently or in parallel with other parts of the program. This greatly improves the performance of your program,especially when dealing with task that involve IO operations or other blocking operatons.

it is similar to breaking down your task into smaller sub-tasks that can be executed independently instead of waiting for them one after another.

### Concurrency 
Concurrency is when two or more control flows(thread) of execution share one or more CPUs. In such cases the scheduler is responsible for deciding when each thread gets to execute and on which CPU. 

## Parallelism
It is a subset of concurrency in whch two or more threads execute at the same real time on two or more CPUs. 

Inorder to register a go `subroutine` you just add the `go` keyword infront of the function call.

Here's a code for GO Routines.

```go
package main

import (
	"fmt"
	"time"
)

func A() {
	fmt.Println("task A")
}
func B() {
	fmt.Println("task B")
}
func main() {
	go A()
	B()

	time.Sleep(1000 * time.Millisecond)
}

```

In a go subroutine, we cannot determine the order in which the functions are executed. For which we created `sync.waitGroups`
`Sync.waitgroup` are in initiative in GO that are used to wait for a collection of GO subroutines to be finished. It allows you to make sure that a group of go routines are completed before continuing the rest of the execution of the program. 

### How it works?

- you create a `sync.WaitGroup` variable to keep track of number of goroutines you want to wait for.
- for each go routine you start you increment the waitGroup counter using the `Add` method.
- inside each go routine you call the `Done` function that signals the work is done.
- finally you call `Wait` on the waitgroup which halts the main program till all the go routines are done.

Here's a sample code on how to do it
```go

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 0; i < 100; i++ {
		// heavy computation task
	}

	fmt.Println(id)
}
func WithGoRoutine() {
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go worker(i, &wg)
	}

	wg.Wait()
}
```
