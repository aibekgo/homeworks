package main

import (
	"fmt"
	"math"
)
func main() {
	a := 3.0
	b := 4.0
	c := 5.0
	r := 3.0
	pi := 3.14
	//Периметр треугольника:
	p_trgl := a + b + c
	//Площадь треугольника:
	p := p_trgl/2.0
	s_trgl := math.Sqrt(p*(p-a)*(p-b)*(p-c))
	//Периметр прямоугольника:
	p_rtgl := 2.0*(a + b)
	//Площадь прямоугольника:
	s_rtgl := a*b
	//Периметр круга:
	p_crcl := 2.0*pi*r
	//Площадь круга:
	s_crcl := pi*(math.Pow(r,2))
	fmt.Println("Периметр треугольника=",p_trgl,"Площадь треугольника=",s_trgl, "Периметр прямоугольника=",p_rtgl,
		"Площадь прямоугольника=",s_rtgl, "Периметр круга=", p_crcl, "Площадь круга=", s_crcl)
}
