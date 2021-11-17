package main

import "fmt"

var test = "package var"

func main() {
	fmt.Println(test)
	var test = "main function var"
	if test := "if syntax var"; true {
		fmt.Println(test)
		test := "if body block var"
		fmt.Println(test)
	}
	fmt.Println(test)
}

// package var
// if syntax var
// if body block var
// main function var
