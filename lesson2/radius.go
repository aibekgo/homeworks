package main

import (
	"fmt"
	"math"
)

func main() {
	var r float64
	fmt.Scan(&r)
	pi := 3.14
	s_crcl := pi*(math.Pow(r,2))
	fmt.Println("площадь круга :", s_crcl)
}
