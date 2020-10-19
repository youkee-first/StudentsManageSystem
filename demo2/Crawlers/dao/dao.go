package dao

import "demo2/Crawlers/domain"

var students = []domain.Student{domain.Student{1,"aaa",1,1},
	domain.Student{2,"bbb",1,1},
	domain.Student{3,"ccc",1,1},
	domain.Student{4,"ddd",1,2},
	domain.Student{5,"eee",1,2},
	domain.Student{6,"fff",1,2},
	domain.Student{7,"ggg",1,3},
	domain.Student{8,"hhh",1,3},
	domain.Student{9,"iii",1,3},
	domain.Student{10,"jjj",1,4},
}

var m map[int]domain.Date = map[int]domain.Date{}
func SetDutyDate(id int, year int, month int, day int) {
	if id!=0 && month!= 0 && day!=0 {
		date := domain.Date{year,month,day}
		m[id] = date
	}

}
func FindDutyStudent(year int, month int, day int) []int {
	sl := make([]int,0)
	for k,v := range m{
		if v.Year == year && v.Month == month && v.Day == day{
			sl = append(sl, k)
		}
	}
	return sl
}
func FindRoommateById(Id int) []domain.Student{
	roommates := make([]domain.Student,0)
	var index int
	flag := false
	for i,v := range students {
		if Id == v.Id{
			index = i
			flag = true
			break
		}
	}
	if flag{
		for _,v := range students {
			if students[index].FloorNum == v.FloorNum && students[index].RoomNum == v.RoomNum{
				roommates = append(roommates, v)
			}
		}
	}

	return roommates
}

func FindRoommateByName(name string)  []domain.Student{
	roommates := make([]domain.Student,0)
	var index int
	flag := false
	for i,v := range students {
		if name == v.Name{
			index = i
			flag = true
			break
		}
	}
	if flag{
		for _,v := range students {
			if students[index].FloorNum == v.FloorNum && students[index].RoomNum == v.RoomNum{
				roommates = append(roommates, v)
			}
		}
	}

	return roommates
}
func FindRoommateByNum(floor int,room int) []domain.Student {
	roommates := make([]domain.Student,0)
	for _,v := range students {
		if floor == v.FloorNum && room == v.RoomNum{
			roommates = append(roommates, v)
		}
	}
	return roommates
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