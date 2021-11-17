package main

import "fmt"

type binaryOperationFunc func(float64, float64) float64

func main() {
	var addFunc binaryOperationFunc
	var redFunc binaryOperationFunc
	if addFunc == nil {
		addFunc = func(a float64, b float64) float64 { return a + b }
	}
	if redFunc == nil {
		redFunc = func(a float64, b float64) float64 { return a - b }
	}
	fmt.Println(addFunc(1, 2))
	// 3
	fmt.Println(redFunc(1, 2))
	// -1
}
