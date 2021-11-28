package main

import "fmt"

func main() {
	i := 0
	for {
		if i >= 10 {
			break
		}
		fmt.Printf("%d ", i)
		i++
	}
	fmt.Println()
}

// 0 1 2 3 4 5 6 7 8 9
