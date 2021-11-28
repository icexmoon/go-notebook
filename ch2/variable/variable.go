package main

import (
	"fmt"

	"github.com/icexmoon/go-notebook/ch2/formater"
)

func main() {
	var varInt int
	var varFloat float64
	var varString string
	var varBool bool
	var varArray [3]int
	var varStruct struct{}
	formater.PrintVariable(varInt)
	fmt.Println(varInt)
	fmt.Printf("%v\n", varInt)
	fmt.Println(varFloat)
	fmt.Println(varString)
	fmt.Println(varBool)
	fmt.Println(varArray)
	fmt.Println(varStruct)
}

// 	0
// 0

// false
// [0 0 0]
// {}
