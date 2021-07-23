package main

import "fmt"

func GetOneToDim(arr []int, c, r int) [][]int {
	arrDim := [][]int{}
	for i := 0; i < len(arr); i++ {
		arrEl := [][]int{}
		for j := 0; j < len(arr[i]); j++ {
			 arrEl =arr[[i][j]:r]
			}
		}
		arrDim = append(arrDim, arrEl)/*,*/
	}
	return arrDim
}

func main() {
	test := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	c:=4
	r:=5
	fmt.Println(GetOneToDim(test,c,r))
}
