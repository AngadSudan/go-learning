## Time Module 

GO provides a powerful `time` package for handling time and date related operations. In GO's time package, the reference date and time for formatting are fixed as "2006-01-02 15:04:05" MST (Mountain standard time) 

GO follows this format because its easy to follow and time when go was announced.

Inorder to `get current time`, we use - `currentTime := time.Now()` which yields an output like `2026-06-14 16:38:26.751343084 +0530 IST m=+0.000023975`
Inorder to `format current time` we use - `formatted := currentTime.Format("02-01-2006")` and not dd-yyyy-mm
Inorder to convert a string to date we can use - `time.Parse(layout_str,date_str)`
Inorder to increase or decrease the time by 2 days we can use the function `currentTime.add(48*time.Hour)`

## Defer keyword

In GO defer statement is used to ensure that a function is called later in the programs execution.
The function called with the defer keyword will be executed when the surrounding function exits in the below example main.

```go
package main
import "fmt"

func main(){
    defer fmt.Println("lets write it down")
}
```

Incase of multiple defer statements it follows LIFO order of execution i.e. Last In First Out.


