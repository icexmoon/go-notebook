package main

import (
	"fmt"

	"github.com/icexmoon/go-notebook/ch4/my_append"
)

func main() {
	var numbers = []int{1, 2, 3}
	fmt.Println(numbers)
	// [1 2 3]
	var numbers2 = [...]int{1, 2, 3}
	fmt.Println(numbers2)
	// [1 2 3]
	var numbers3 = make([]int, 3, 5)
	fmt.Println(numbers3)
	// [0 0 0]
	numbers3[0] = 1
	numbers3[1] = 2
	numbers3[2] = 3
	fmt.Println(numbers3)
	// [1 2 3]
	array1 := [...]int{1, 2, 3, 4, 5}
	slice1 := array1[:]
	fmt.Println(slice1)
	// [1 2 3 4 5]
	slice2 := array1[2:4]
	fmt.Println(slice2)
	// [3 4]
	slice3 := slice1[2:]
	fmt.Println(slice3)
	// [3 4 5]
	slice4 := make([]int, 3, 5)
	slice4[0] = 1
	slice4[1] = 2
	slice4[2] = 3
	fmt.Println(len(slice4), cap(slice4))
	// 	3 5
	fmt.Println(slice4)
	// [1 2 3]
	slice4 = slice4[:5]
	slice4[3] = 4
	slice4[4] = 5
	fmt.Println(slice4)
	// [1 2 3 4 5]
	numbers3 = append(numbers3, 4)
	numbers3 = append(numbers3, 5)
	numbers3 = append(numbers3, 6)
	fmt.Println(numbers3)
	// [1 2 3 4 5 6]
	numbers4 := []int{1, 2, 3}
	numbers4 = my_append.MyAppend(numbers4, 6)
	numbers4 = my_append.MyAppend(numbers4, 7)
	fmt.Println(numbers4)
	// [1 2 3 6 7]

}
