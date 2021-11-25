package main

import "sync"

type student struct {
	name string
	age  int
}

type students struct {
	stdsMutex sync.RWMutex
	stds      map[string]*student
}

func (s *students) getStudent(name string) *student {
	s.stdsMutex.RLock()
	student, ok := s.stds[name]
	s.stdsMutex.RUnlock()
	if !ok {
		s.stdsMutex.Lock()
		student, ok = s.stds[name]
		if !ok {
			student := initStudent(name)
			s.stds[name] = student
		}
		s.stdsMutex.Unlock()
	}
	return student
}

func initStudent(name string) *student {
	return &student{}
}

func main() {
	var ss students
	go func() {
		ss.getStudent("std1")
	}()
	go func() {
		ss.getStudent("std2")
	}()
}
