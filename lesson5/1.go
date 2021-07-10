package main

import "fmt"

func main() {
	a := []int{1, 2, 4, 9, 0}
	min_e := 0
	for i := 0; i < len(a); i++ {
		if min_e < a[i] {
			min_e = a[i]
		}
	}
	fmt.Println(min_e)
}
