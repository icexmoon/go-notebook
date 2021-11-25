package main

type student struct {
	name string
	age  int
}

type students struct {
	stds map[string]*student
}

func (s *students) getStudent(name string) *student {
	student, ok := s.stds[name]
	if !ok {
		student := initStudent(name)
		s.stds[name] = student
	}
	return student
}

func initStudent(name string) *student {
	return &student{}
}

func main() {
	var ss students
	ss.getStudent("std1")
	ss.getStudent("std2")
	go func() {
		ss.getStudent("std1")
	}()
	go func() {
		ss.getStudent("std2")
	}()
}
