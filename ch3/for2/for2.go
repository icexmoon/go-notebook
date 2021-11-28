package main

import "fmt"

func main() {
	i := 0
	for i < 10 {
		fmt.Printf("%d ", i)
		i++
	}
	fmt.Println()
}

// 0 1 2 3 4 5 6 7 8 9
