package main

import "fmt"

type student struct {
	id    int
	name  string
	class string
}

func newStudent(id int, name, class string) *student {
	return &student{
		id:    id,
		name:  name,
		class: class,
	}
}

type studentMqr struct {
	allStudents []*student
}

func newStudentMqr() *studentMqr {
	return &studentMqr{
		allStudents: make([]*student, 0, 100),
	}
}

//添加学生
func (s *studentMqr) addStudent(newStu *student) {
	s.allStudents = append(s.allStudents, newStu)
}

//添加学生
func (s *studentMqr) mnodifytudent(newStu *student) {
	for i, v := range s.allStudents {
		if newStu.id == v.id{
			s.allStudents[i] = newStu
			return
		}
	}
	fmt.Println("输入的学生不存在")
}

//展示学生
func (s *studentMqr) showStudent() {
	for _, v := range s.allStudents {
		fmt.Printf("学号:%d 姓名: %s 班级：%s \n", v.id, v.name, v.class)
	}

}
