package main

import "fmt"

func main() {
	var numbers []*int
	for i := 0; i < 10; i++ {
		func() {
			numbers = append(numbers, &i)
		}()
	}
	for _, number := range numbers {
		fmt.Printf("%d ", *number)
	}
}

// 10 10 10 10 10 10 10 10 10 10
