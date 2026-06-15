## Creating Web Servers in GO

Inorder to create an http server, we are going to use the `net/http` package in GO. and the function that is going to be involved is the `http.ListenAndServe()` function.

the method takes in two input parameters
- the port that it is binded to 
- the http handlers

How the function works is it basically listens up for TCP Requests on the specified addresses and then serves those requests using the HTTP handler. 

Here's a sample example of the code for just binding your PORT using go code.

```go
package main

import "net/http"

type server struct {
	addr string
}

func (s *server) ServeHTTP(http.ResponseWriter, *http.Request) {}

func main() {
	s := &server{
		addr: ":8000",
	}
	http.ListenAndServe(s.addr, s)
}
```

\`ServeHTTP\` is the single method required by the \`http.Handler\` interface.

| With ServeHTTP | Without ServeHTTP |
|----------------|-------------------|
| \`*server\` implements \`http.Handler\` | \`*server\` does not implement \`http.Handler\` |
| Can pass \`s\` to \`ListenAndServe\` | Compile-time error |
| Handles incoming requests | No request handling capability |

So in your code, the method exists primarily to make \`*server\` satisfy \`http.Handler\` and become the HTTP server's request handler.

When your server is hit, this is the method that is invoked via which you return a sample response. If you want to send a hello message to the user when they hit the backend url you can run the following code.

```go
package main

import (
	"net/http"
)

type server struct {
	addr string
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello from the server"))
}

func main() {
	s := &server{
		addr: ":8000",
	}
	http.ListenAndServe(s.addr, s)
}

```


NOw in this way we have a request handler that returns us the same thing on any request that you recieve. But in more production codebase we'd want routing and segregation of resources which can be done via 

```go
package main

import (
	"net/http"
)

type server struct {
	addr string
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		switch r.URL.Path {
		case "/":
			w.Write([]byte("index page"))
			return
		}
	default:
		w.Write([]byte("404 not found"))
		return
	}
}

func main() {
	s := &server{
		addr: ":8000",
	}
	http.ListenAndServe(s.addr, s)
}
```

Inorder for us to handle routing the better way we use something called a `mux`.

--- 

### MUX 

A mux in Go stands for Multiplexer.

Its job is to examine an incoming HTTP request and decide which handler should process it based on things like the URL path, HTTP method, host, etc.

The way in which we define a mux is `mux := http.NewServeMux()`.
Once we have our mux ready, the next thing that we do is create an instance of the http server. 
```go
srv := &http.Server{
		Addr:    s.addr,
		Handler: mux,
	}
```

now inorder for the request handling to happen we need to register the routes which can be done in the following way.
```go
    mux.HandleFunc("GET /users", s.handleUserMethods)
```

The first is the route and the request type after go `1.22` and userMethod is the handler that will be handling the further request which goes something like this 

```go
func (s *api) handleUserMethods(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("welcome user"))
```


