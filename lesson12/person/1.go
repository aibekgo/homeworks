package person

import "fmt"

type Person struct {
	FirstName string
	LastName string
	Balance float64
	FullName string
}

func (p *Person) PrintData() {
	p.setDefault()
	p.setFullName()
	fmt.Println(p.FirstName, p.LastName, p.Balance, p.FullName)
}

func (p *Person) setDefault() {
	if p.FirstName == "" {
		p.FirstName = "Ivan"
	}
	if p.LastName == "" {
		p.LastName = "Ivanov"
	}
	if p.Balance == 0 {
		p.Balance = 100
	}
	if p.FullName == "" {
		p.FullName = p.LastName  + " " + p.FirstName
	}
}
func (p *Person) setFullName() {
	p.FullName =  p.LastName  + " " + p.FirstName
}

func main() {
	p1 := Person{"Artur", "Li", 50, ""}
	p2 := Person{}
	p1.PrintData()
	p2.PrintData()
}