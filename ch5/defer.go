package main

import "fmt"

func myFunc7() {
	defer fmt.Println("defer in myFunc is called")
	fmt.Println("myFunc is returned")
	return
}

func main() {
	myFunc7()
	// myFunc is returned
	// defer in myFunc is called
}
