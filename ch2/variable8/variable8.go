package main

import "fmt"

func main() {
	var a, b, c int = 1, 2, 3
	fmt.Println(a, b, c)
	a, b, c = 4, 5, 6
	fmt.Println(a, b, c)
}

// 1 2 3
// 4 5 6
