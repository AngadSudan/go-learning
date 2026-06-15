package main

import (
	"fmt"
	"time"
)

func main() {
	curr := time.Now()
	fmt.Println(curr)

	formatted := curr.Format("02-2006-01")
	fmt.Println(formatted)
}
