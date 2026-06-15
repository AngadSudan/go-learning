package routine

import (
	"fmt"
	"sync"
	"time"
)

func A() {
	fmt.Println("task A")
}
func B() {
	fmt.Println("task B")
}

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

func WithoutGoRoutine() {
	go A()
	B()

	time.Sleep(1000 * time.Millisecond)
}
