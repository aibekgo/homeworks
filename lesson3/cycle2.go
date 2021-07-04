package main

import "fmt"

func main() {
	s := 0
	e := 0
	fmt.Scanf("%d %d", &s, &e)
	sum := 0
	mlt := 0
	n := e - s
	n1 := n/2
	for i:=s; i<n1; i++ {
	sum = sum+i
		//fmt.Println("сумма=", sum)
	}
	for i:=n1; i<s; i++ {
		mlt = mlt*i
		//fmt.Println( "умножение=", mlt)
	}
	fmt.Println("сумма=", sum, "умножение=", mlt)
}
