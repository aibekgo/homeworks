package main

import "fmt"

func main() {
	arr := [][]int{
		{25, 1},
		{70, 1},
		{100, 0},
		{3, 1},
	}
	maxi := 0
	maxiIndex := 0
	for i := 0; i < len(arr); i++ {
		element := arr[i]
		if element[0] > maxi && element[1] == 1 {
			maxi = element[0]
			maxiIndex = i + 1
		}
	}
	fmt.Println(maxiIndex)
}