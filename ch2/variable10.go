package main

import "fmt"

func main() {
	var a = "a in main function"
	if true {
		a, b := "a in if body", "b in if body"
		fmt.Println(a, b)
	}
	fmt.Println(a)
}

// a in if body b in if body
// a in main function
