package main

import (
	"fmt"
	"log"

	"github.com/icexmoon/go-notebook/ch11/input"
)

func main() {
	fmt.Print("Enter your name:")
	name, err := input.GetLine()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Welcome!", name)
}
