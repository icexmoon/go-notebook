package main

import "sync"

type student struct {
	name string
	age  int
}

type students struct {
	stdsMutex sync.Once
	stds      map[string]*student
}

func (s *students) GetStudent(name string) *student {
	s.stdsMutex.Do(s.initAll)
	return s.stds[name]
}

func (ss *students) initAll() {
	ss.stds["std1"] = initStudent("std1")
	ss.stds["std2"] = initStudent("std2")
}

func initStudent(name string) *student {
	return &student{}
}

func main() {
	var ss students
	go func() {
		ss.GetStudent("std1")
	}()
	go func() {
		ss.GetStudent("std2")
	}()
}
