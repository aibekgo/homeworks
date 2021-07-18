package main

import "fmt"

func GetMaxOfArr(a []int) int {
	m := a[0]
	for i := 0; i < len(a); i++ {
		if a[i] >= m {
			m = a[i]
		}
	}
	return m
}

func main() {
	a := []int{2, 8, 7, 4, 5, 10, 5}
	k := GetMaxOfArr(a)
    fmt.Println(k)
}
