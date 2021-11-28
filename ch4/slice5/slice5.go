package main

import "fmt"

func main() {
	slice := make([]int, 3, 5)
	slice[0] = 1
	slice[1] = 2
	slice[2] = 3
	changeSlice(slice)
	fmt.Println(slice)
	// [1 2 3]
	slice = slice[:5]
	fmt.Println(slice)
	// [1 2 3 4 5]
}

func changeSlice(slice []int) {
	sliceCopy := slice[:5]
	sliceCopy[3] = 4
	sliceCopy[4] = 5
}
