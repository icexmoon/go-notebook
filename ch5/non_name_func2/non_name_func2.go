package main

import "fmt"

func main() {
	func() {
		fmt.Println("this is a non named function")
	}()
	// this is a non named function
}
