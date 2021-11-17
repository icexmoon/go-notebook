package main

import "fmt"

type people struct {
	Name string
	Age  int
}

func (*people) string() string {
	return fmt.Sprintf("%s is %d years old", people.Name, people.Age)
}

type stringer interface {
	string() string
}

func main() {
	var varInt int = 1                                                          //整型
	var varFloat float64 = 1.5                                                  //浮点型
	var varArray [3]int = [...]int{1, 2, 3}                                     //整型数组
	var varMap map[string]bool = map[string]bool{"apple": true, "xiaomi": true} //映射
	var varBool bool = true                                                     //布尔
	var varSlice []int = []int{1, 2, 3}                                         //切片
	var varStruct people = people{Name: "lalala", Age: 11}                      //结构体
	var varInterface stringer = &people{Name: "lalala", Age: 11}                //接口
	var varChannel chan int = make(chan int)                                    //通道
}
