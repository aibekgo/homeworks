package main

import "fmt"

func main() {
	var a,b,c int
	fmt.Scan(&a)
	fmt.Scan(&b)
	fmt.Scan(&c)
	maxi := 0
	if a >= b && a >= c { maxi=a} else if  b >= a && b >= c  {maxi=b} else if  c >= a && c >= b {
		maxi=c}
	fmt.Println("maximum",maxi)
}

