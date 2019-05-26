package learn

import (
	"strconv"
)

type Student struct {
	ID     int
	Name   string
	Gender string
}

type Class struct {
	Title    string
	Students []*Student
}

func NewClass(title string, students *[]Student) {
	class := new(Class)
	class.Title = title
	for _, student := range *students {
		class.Students = append(class.Students, &student)
	}
}


func newStudent(ID int,Name,Gender string) *Student{
	s := new(Student)
	s.ID = ID
	s.Gender = Gender
	s.Name = Name
	return s
}

/*
创建Student
*/
func makeStrudentList()*[]Student{

	studentList := make([]Student,0,200)
	

	for i:=0;i<=20;i++{
		s:=newStudent(i,"test"+strconv.Itoa(i),strconv.Itoa(i))
		studentList = append(studentList,*s)
	}
	return &studentList
}




func Test18(){
	
}





