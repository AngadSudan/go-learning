# Go Lang

Go is a simple, fast and concurrent programming language developed by Google. It is strong statically types programming language which essentially means variables have a type that cannot be changed over time and must be deined at compile time. This would help us to capture errors easily at compilation time.

## Reasons for Go Language

---

- Go is designed for `faster compilation without the dependency checking`.
- Go has a comprehensive set of string functions which uses `garbage collection` to make working with strings more efficient compared to some other languages.
- Go has its own lighweight thread managed by go runtime called `goroutine`.

## Key Features

---

- Concurrency
- Garbage Collection
- Static typing
- Lightweight
- Fast Compilation
- Zero dependencies
- Built in support for testing

Tools written in go are `Kubernetes`,`docker`,`Prometheus`

## Some Standard for GO

---

the folder structure in go usually follows the following folders

- bin/ - the directory contains the executable binaries after your go program is built
- pkg/ - contains compiled package files. When you compile a go package its object files go here
- src/ - this is where your complete codebase exists.

```
NOTE
All your codebase must lie in the src directory so its a part of the go workspace.
```

## GO Modules

if we want to manage go project outside default `GOPATH` then we use Go Modules, which allow us to manage projects outside of `GOPATH`.

You can initialize a new go module by using

```go
go mod init
```

this creates a go.mod file which talks to the `go modules` and makes sure that the project is self contained.

To run the project that you created run the action `go run main.goi

## Variables in Go

variables are location where we use to store data.
Go uses a JS type notation for decalaration of variables. Since it is statically types, a variable comprises of a name and a datatype

```go
var <variable1> <datatype> = <value>
```

datatypes that can be stored in GO are:

| datatype | keyword |
| -------- | ------- |
| integer  | int     |
| float    | float64 |
| string   | string  |
| boolean  | bool    |

```
If you want to use any value again in some other file/module then their variable name must always begin with a capital letter. The same rules applies for functions as well.
```
