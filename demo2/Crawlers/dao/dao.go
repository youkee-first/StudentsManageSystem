package dao

import "demo2/Crawlers/domain"

var students = []domain.Student{domain.Student{01,"aaa",1,01},domain.Student{02,"bbb",1,02},domain.Student{03,"ccc",1,03}}

func ()  {
	
}
func FindAll() []domain.Student{
	return students
}

func AddStudent(student domain.Student)  {
	students = append(students, student)
}

func DeleteById(id int) {
	for i,v := range students{
		if v.Id==id{
			students = append(students[:i], students[i+1:]...)
			break
		}
	}

}

func UpdateStudent(student domain.Student,id int)  {
	for i,v := range students{
		if v.Id==id{
			students[i] = student
			break
		}
	}
}