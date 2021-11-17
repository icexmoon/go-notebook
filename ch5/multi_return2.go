package main

import (
	"fmt"
	"strconv"
)

//将给定的字符串形式的数字转换为整形数字
//strNumber 字符串形式的数字
//return int 转换后的整形数字
//return error 转换时出现的错误
func myFunc4(strNumber string) (int, error) {
	intNumber, err := strconv.Atoi(strNumber)
	return intNumber, err
}

func callMyFunc4(strNumber string) {
	intNumber, err := myFunc4(strNumber)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%s after convered is %d\n", strNumber, intNumber)
}

func main() {
	callMyFunc4("123")
	// 123 after convered is 123
	callMyFunc4("12.5")
	// strconv.Atoi: parsing "12.5": invalid syntax
	callMyFunc4("abc")
	// strconv.Atoi: parsing "abc": invalid syntax

}
