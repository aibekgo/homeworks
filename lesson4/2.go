package main

import "fmt"

func main() {
	c:=[]int{}
	a:=[]int{}
	b:=[]int{}
	n:=20
	for i:=1; i<n; i++ {
		c = append(c,i)
	}
	fmt.Println(c)
	for i:=c[0]; i<len(c); i++ {
		if c[i]%2==0 {
			continue }
		a = append(a,i)

	}
	fmt.Println(a)
	for i:=c[0]; i<len(c); i++ {
		if c[i]%2!=0 {
			continue }
		b = append(b,i)
	}
	fmt.Println(b)

}
