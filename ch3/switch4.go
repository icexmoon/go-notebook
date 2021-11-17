package main

import "fmt"

func main() {
	checkType("hello")
	checkType("1")
	checkType(1)
	checkType(2.5)
}

func checkType(variable interface{}) {
	var varForm string
	switch variable.(type) {
	case int:
		varForm = "int"
	case float64, float32:
		varForm = "float"
	case string:
		varForm = "string"
	default:
		varForm = "not known"
	}
	fmt.Printf("%v 's type is %s\n", variable, varForm)
}

// hello 's type is string
// 1 's type is string
// 1 's type is int
// 2.5 's type is float
