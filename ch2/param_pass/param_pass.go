package main

import "fmt"

func main() {
	var varArray = [...]int{1, 2, 3}
	var varSlice = []int{1, 2, 3}
	changeArray(varArray)
	changeSlice(varSlice)
	fmt.Println(varArray)
	fmt.Println(varSlice)
}

func changeArray(array [3]int) {
	array[0] = 99
	fmt.Println("after change:", array)
}

func changeSlice(slice []int) {
	slice[0] = 99
}

// after change: [99 2 3]
// [1 2 3]
// [99 2 3]
