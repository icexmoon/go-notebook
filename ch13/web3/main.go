package main

import (
	"fmt"
	"net/http"
)

type hello struct{}

func (hello) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(rw, "hello world!")
}

func main() {
	http.ListenAndServe(":8080", hello{})
}
