
package main

import "fmt"

func bubbleSort(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr); j++ {
			if arr[i] > arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
	return arr
}

func GetSumOfTwoDim(arrSum [][]int) []int {
	sums := []int{}
	for i := 0; i < len(arrSum); i++ {
		sumi := 0
		for j := 0; j < len(arrSum[i]); j++ {
			sumi += arrSum[i][j]
		}
		sums = append(sums, sumi)
	}
	return sums
}


func main() {
	arr := [][]int{{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 1, 2, 3}}

	fmt.Println(arr)
	arrSum:= GetSumOfTwoDim(arr)
	arrSort:= bubbleSort(arrSum)
	fmt.Println(arrSum)
	fmt.Println(arrSort)
}


/*
2) создать функцию которая отсортирует по убыванию сумму элементов из двумерного массива
1 2 3 4
5 6 7 8
9 1 2 3

[10, 26, 15]

Ответ:[26,25,10]


*/
