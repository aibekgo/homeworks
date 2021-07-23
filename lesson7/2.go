package main

import "fmt"

func GetMaxOfTwoDim(arr [][]int) []int {
	sums := []int{}
	for i := 0; i < len(arr); i++ {
		maxi := arr[0][0]
		for j := 0; j < len(arr[i]); j++ {
			if maxi < arr[i][j] {
				maxi = arr[i][j]
			}
		}
		sums = append(sums, maxi)
	}
	return sums
}

func main() {
	test := [][]int{
		{1, 2, 3, 4, 5},
		{2, 8, 8, 1, 10},
		{10, 10, 3, 4},
	}
	fmt.Println(GetMaxOfTwoDim(test))
}