package main

import "fmt"

func myFunc9(a int, b int) (result int) {
	defer func() {
		result = result * 2
	}()
	return a + b
}

func main() {
	fmt.Println(myFunc9(1, 2))
	//6
}
