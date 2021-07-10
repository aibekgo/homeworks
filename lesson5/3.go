package main

import "fmt"

func main() {
	arr := []int{1, 1, 2, 3, 2, 3, 4}
	n := 0
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr); j++ {
			if arr[i] == arr[j] {n = n+1}
			//fmt.Println(n)
		}
	}
fmt.Println(n)
}
