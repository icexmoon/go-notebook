package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6}
	for index, value := range numbers {
		fmt.Printf("numbers[%d]=%d\n", index, value)
	}
}

// numbers[0]=1
// numbers[1]=2
// numbers[2]=3
// numbers[3]=4
// numbers[4]=5
// numbers[5]=6
