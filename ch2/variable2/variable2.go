package main

import (
	"fmt"

	"github.com/icexmoon/go-notebook/ch2/formater"
)

func main() {
	var varSlice []int
	formater.PrintVariable(varSlice)
	fmt.Println(varSlice == nil)
	var varMap map[string]bool
	formater.PrintVariable(varMap)
	fmt.Println(varMap == nil)
	var varChannel chan int
	formater.PrintVariable(varChannel)
	fmt.Println(varChannel == nil)
	var varInterface interface{}
	formater.PrintVariable(varInterface)
	fmt.Println(varInterface == nil)
}

// []int(nil) []int []
// true
// map[string]bool(nil) map[string]bool map[]
// true
// (chan int)(nil) chan int <nil>
// true
// <nil> <nil> <nil>
// true
