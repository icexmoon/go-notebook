package main

import (
	"fmt"
	"net/http"
)

type hello struct{}

func (hello) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(rw, "hello world!")
}

type bye struct{}

func (bye) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(rw, "bye~~")
}

func main() {
	serverMux := http.NewServeMux()
	serverMux.Handle("/hello", hello{})
	serverMux.Handle("/bye", bye{})
	http.ListenAndServe(":8080", serverMux)
}
