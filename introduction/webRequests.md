## Making Web Requests in GO

In GO, a web requests refers to the HTTP request made to the server. The requests are used to retrieve or send data over the internet to interact with application APIs.

In GO, we do that using a package called `net/http` which provides interface to create and send HTTP requests as well as to handle the responses.

We use the package to make a get request calling the function `response,err := http.GET(url)` which returns you 2 things a response and err. 
Now the res requires a connection in GO you once its task is over you have to close the function manually. via `defer response.Body.Close()`.

Once that is done you get the whole response in your response field so you need to first make it accessible using the ioutil library as follows
`content,err:=ioutil.ReadAll(respnose.Body)`. The content here is still in the form of bytes so you need to then convert it to a string to be finally accessible.
Heres the complete walk around code. 
```go
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	var url string = "https://jsonplaceholder.typicode.com/todos/1"
	response, err := http.Get(url)
	if err != nil {
		fmt.Println("error in making API Call")
		return
	}

	defer response.Body.Close()

	// reading the data
	content, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("error in reading API Call")
		return
	}

	parsedData := string(content)
	fmt.Println(parsedData)
}

```
--- 

Another way of making API Calls in GO is by using `http.NewRequest()` method. it takes in 3 parameters
- the method type which can be PUT,POST,GET or `http.newMethodPUT` similarly for rest of the method
- then url you want to send your data to 
- the ioreader for the data

Once this is done you set the header using `req.Header.Set('header_name','header_value')`
and after this your req is ready to go.
But without a client a req cannot be fired so you make a client using `client:=http.Client` and to fire the request we do `client.do(req)` which returns the res and err. 


## Handling URLs in GO

in GO the `net/url` package provides functionalities to parse, construcut and manipulates URL. Parsing of those URLs can be done using url.
URL Components

- Schema : indicates the protocol used (http/https)
- Host : Specify the domain name additionally the port number 
- Path : represents the path to find the resource on the server
- RawQuery: contains query parameters (key value pair ending with `?` and `&`


## JSON 
in GO we use the package `encoding/json` to encode and decode JSON values. in GO `encoding` of values is termed as `marshalling` and `decoding` of those values are termed as `unmarshalling`.

to create a JSON in GO, we create a struct with the datatypes of the row and what will that JSON values be termed. Here's an example

```go
type Person struct {
    Name String "json:name"
    Age  int    "json:age"
    IsAdult bool "json:is_adult"
}
```

To Convert struct to JSON, we use `jsonData,err := json.Marshal(person)` 
To Convert JSON to struct we use  `err:=json.Unmarshal(person,&data)`


