package main

import "fmt"

func myFunc10(mesgs ...string) {
	for _, message := range mesgs {
		fmt.Printf("%s ", message)
	}
	fmt.Println()
}

func main() {
	myFunc10("hello", "world", "!")
	// hello world !
	messages := []string{"hello", "world", "!"}
	myFunc10(messages...)
	// hello world !
}
