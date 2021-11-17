package main

import "fmt"

func main() {
	myFunc := func(a int, b int) int { return a + b }
	fmt.Printf("%T", myFunc)
	// func(int, int) int
}
