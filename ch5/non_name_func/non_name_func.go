package main

import "fmt"

func main() {
	nonNamedFunc := func(a int, b int) int {
		return a + b
	}
	fmt.Println(nonNamedFunc(1, 2))
	//3
}
