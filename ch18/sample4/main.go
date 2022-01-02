package main

import "fmt"

func Add(a int, b int) int {
	return a * b
}

func Mul(a int, b int) int {
	//is developing...
	return 0
}

func main() {
	fmt.Printf("%d+%d=%d\n", 1, 2, Add(1, 2))
}
