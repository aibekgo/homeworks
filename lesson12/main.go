package main

import "homeworks/lesson12/person"

func main() {
	p1 := person.Person{"Artur", "Li", 50, ""}
	p2 := person.Person{"", "",   0, ""}
	p1.PrintData()
	p2.PrintData()

}
