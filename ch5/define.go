package main

import "fmt"

func myFunc(message string) bool {
	fmt.Println(message)
	// hello
	return true
}

func main() {
	myFunc("hello")
}
