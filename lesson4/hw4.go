package main

import "fmt"

func main() {
	arr :=[]int{-1,2,3,4,-5,-2,-4}
	arr_p := []int{}
	arr_n := []int{}
	for i:=0; i<len(arr); i++ {
		if arr[i]>=0 {arr_p = append(arr_p, arr[i])
		}		else {arr_n = append(arr_n, arr[i])}
	}
	fmt.Println("положительные",arr_p)
	fmt.Println("отрицательные",arr_n)
}


/*Разложить положительные и отрицательные числа по разным массивам
[-1,2,3,4,-5,-2,-4]*/