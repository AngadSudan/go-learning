## Packages and Imports

In go packages are used instead of classes. There is no concept of OOPS. Every package in GO is essentially a directory in your workspace. **Each GO Files must belong to some package**, and must start with package followed by the package name.

The main package is a special package in GO. An executable program must contain a main package. 

Go uses relative imports to bring packages to current file. They usually come from the path `$GOPATH/src` since most packages are stored in pkg directory. 

__Inorder to import a package you must use the underlying syntax__ 
```go
import (
    "fmt"
)
```

main function can only be defined under main package itself.

---
## Creating a new package
lets say we want to create another package which is `my_go_util` so first we will create a file util.go. As you can guess a package name doesnot need to be same as the file name. 

```go
package my_go_util
import "fmt"

func PrintMessage(message string){
    fmt.Println(message)
}
```

- In Go the concept of package is closely ties to directory structure.
- Every file in a same package is considered part of the same package

