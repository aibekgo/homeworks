package main

import "fmt"

type Student struct {
	FirstName string
	LastName  string
	Marks     Mark
}
type Mark struct {
	Name  []string
	Value []int
}

func (m *Mark) GetMark() string {
	fmt.Println(m.Name, m.Value)
	o := fmt.Sprintf("%s", m.Name)
	return o
}
func (m *Mark) GetAverageMark() {
	avg := 0
	for i := 0; i < len(m.Value); i++ {
		avg += m.Value[i]
	}
	fmt.Println(avg/len(m.Value))
}

func (s *Student) GetFullInfo() {
	fmt.Printf("FirstName: %s LastName:%s Marks:%s",
		s.FirstName, s.LastName, s.Marks.GetMark())
}

func main() {
	m1 := Mark{
		Name:  []string{"A", "B", "C", "D", "C", "F"},
		Value: []int{100, 90, 80, 70, 60, 50}}

	s1 := Student{
		FirstName: "Slava",
		LastName:  "Osipov",
		Marks:     m1,
	}

	s1.GetFullInfo()
	m1.GetAverageMark()
}

/*
   1) создать структуру в которой будет храниться информация об ученике и его оценках
+Mark:
    Name:string
    Value:int

+GetMark():string -> "Name:...., Value:...."

example:
Name:A,B,C,D,E,F
Value:100,90,80,60,50,40

+Ученик:
    FirstName:string
    LastName:string
    Marks:[]Mark

GetFullInfo():string -> "FirstName:..., LastName:...., Marks:[A,B,C,D,E,F]"
GetAverageMark():int ->*/
