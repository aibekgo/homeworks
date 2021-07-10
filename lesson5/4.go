package main

import "fmt"

func main() {
	arr := []int{1, 1, 2, 3, 2, 3, 4, 4}
	unique := []int{}
	total := 0
	for i := 0; i < len(arr); i++ {
		n := 0
		for j := 0; j < len(arr); j++ {
			if arr[i] == arr[j] {
				unique = append(unique, arr[i])
				n = n + 1
			}
		}
		if n >= 2 {
			isNotExist := true
			for l := 0; l < len(unique); l++ {
				if arr[i] == unique[i] {
					isNotExist = false
				}
			}
			if isNotExist {
				total += 1
			}

		}
	}
	fmt.Println(total)
	fmt.Println(unique)
}
