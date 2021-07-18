package main

import "fmt"

func getMaxFromNumbers(a, b, c int) int {
m:=0
if a>=b && a>=c {m=a
}else if  b>=a && b>=c {m=b
} else if  c>=a && c>=b {m=c
}
return m
}

func main() {
	k := getMaxFromNumbers(10, 7, 5)
	fmt.Println(k)

}