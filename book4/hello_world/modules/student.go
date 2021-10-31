// student.go
package modules

import "fmt"

type Student struct {
	name string
	age  int
}

func (s *Student) Print() {
	fmt.Println(s.name, "is", s.age, "years old")
}

func (s *Student) GetName() string {
	return s.name
}

func NewStudent(name string, age int) *Student {
	return &Student{
		name: name,
		age:  age,
	}
}
