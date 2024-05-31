package main

import (
	"io"
	"net/http"
	"os"
)

// Server has Front Controller Pattern, Back controller pattern. Request received by one object in Front controller, it decides based on request and routes the request to back controller which then sends it to the client.

// the net/http package does all this for us

func main() {
	http.HandleFunc("/menu", Handler) // handler invoked with every request (url path), back controller (registering our back controller so that web service is aware of it)
	http.ListenAndServe(":3000", nil) // listen for requests and serve. Usually provide Ip address or network address that server is listing on, go assumes local host by default, we also have to provide tcp port.
	// when I open localhost on port 3000, it sends a get request to the server listening which triggers the handler function.
}

// go provides us a default front controller, implementing a basic back controller function below
// response writer used to write responses back out to the request, and second arg used to read the request in.
func Handler(w http.ResponseWriter, r *http.Request) {
	f, _ := os.Open("./menu.txt") // returns file object and error
	io.Copy(w, f)                 // allows copy from read source to a write source (file to request)
}
