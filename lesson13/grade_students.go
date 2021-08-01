package main

import (
	"fmt"
	"homeworks/lesson13/student"
)

func main() {
	s1 := student.Student{
		FullName:    "Ivanov Valera",
		AverageMark: 90,
	}
	s2 := student.Student{
		FullName:    "Sidorov Petr",
		AverageMark: 100,
	}
	s3 := student.Student{
		FullName:    "Navalniy Lesha",
		AverageMark: 80,
	}
	grade1 := student.Grade{
		Name:     "g1",
		Students: []student.Student{s1,s2},
	}

	fmt.Println("---Before add---")
	grade1.ListStudents()
	grade1.AddStudent(s3)
	fmt.Println("---After add---")
	grade1.ListStudents()
	fmt.Println(grade1.PrintStudentWithMaxAverageMark())

}
