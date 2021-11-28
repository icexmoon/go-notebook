package main

import "fmt"

func myFunc2(param interface{}) {
	switch param.(type) {
	case string:
		fmt.Println("this is a string")
	case int:
		fmt.Println("this is a int")
	case float64, float32:
		fmt.Println("this is a float")
	default:
		fmt.Println("unknown type")
	}
}

func main() {
	myFunc2("123")
	// 	this is a string
	myFunc2(123)
	// this is a int
	myFunc2(123.1)
	// this is a float
	myFunc2(func() {})
	// unknown type
}
