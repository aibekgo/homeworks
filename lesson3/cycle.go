package main

import "fmt"

func main() {
 	n:=0
 	u:=1
 	fmt.Scanf("%d", &n)
 	for i:=1; i<=n; i++ {
 		u = u * i
	}
fmt.Println(u)

 }
