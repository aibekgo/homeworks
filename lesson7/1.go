package main

import "fmt"

func GetMulOfTwoDim(arr [][]int) []int {
	sums := []int{}
	for i := 0; i < len(arr); i++ {
		sumi := 1
		for j := 0; j < len(arr[i]); j++ {
			sumi *= arr[i][j]
		}
		sums = append(sums, sumi)
	}
	return sums
}

func main() {
	test := [][]int{
		{1, 2, 3, 4, 5},
		{2, 8, 8, 1, 10},
		{10, 10, 3, 4},
	}
	fmt.Println(GetMulOfTwoDim(test))
}