package main

import "fmt"

//返回字符串的长度和首个字符
//message 字符串
//return int 字符串长度
//return rune 字符串首个字符
func myFunc3(message string) (int, rune) {
	length := len(message)
	runeString := []rune(message)
	return length, runeString[0]
}

func main() {
	length, firstRune := myFunc3("hello")
	fmt.Printf("the length is %d the first rune is %s\n", length, string(firstRune))
	// 	the length is 5 the first rune is h
	length, firstRune = myFunc3("你好")
	fmt.Printf("the length is %d the first rune is %s\n", length, string(firstRune))
	// the length is 6 the first rune is 你
}
