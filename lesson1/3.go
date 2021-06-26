package main

import "fmt"

func main() {
	n := 10000.0
	y := 10.0
	p := 0.1
	m := (n * p * ((1.0 + p)*y)) / (12.0 * ((1.0 + p)*y)-1.0)
	s := (m * 12.0) * y
	fmt.Println("месячные выплаты:",m, "суммарная выплата:",s)
}

