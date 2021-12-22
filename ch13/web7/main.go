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
	http.Handle("/hello", hello{})
	http.Handle("/bye", bye{})
	http.ListenAndServe(":8080", nil)
}
