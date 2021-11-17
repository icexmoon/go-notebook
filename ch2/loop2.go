package main

import "fmt"

func main() {
	var numbers []*int
	for i := 0; i < 10; i++ {
		i := i
		func() {
			numbers = append(numbers, &i)
		}()
	}
	for _, number := range numbers {
		fmt.Printf("%d ", *number)
	}
}

// 0 1 2 3 4 5 6 7 8 9
