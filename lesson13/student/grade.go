package student

import "fmt"

type Grade struct {
	Name string
	Students []Student
}

func (g *Grade)  AddStudent(s Student)  {
	g.Students = append(g.Students, s) }

func (g *Grade)  ListStudents() {
	for i := 0; i < len(g.Students); i++ {
		fmt.Println(g.Students[i].getFullInfo())
	}
}

func (g *Grade) PrintStudentWithMaxAverageMark() int {
	maxAverage := 0
	for i := 0; i < len(g.Students); i++ {
		//maxAverage = g.Students[0].AverageMark
		if maxAverage < g.Students[i].AverageMark  {
			maxAverage = g.Students[i].AverageMark
		}
	}
	return maxAverage
}