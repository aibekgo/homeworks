package main

import "fmt"

func main() {
	arr := []int{-1, 2, 3, 4, -5, -2, -4}
	arr1 := []int{}
	l := len(arr)
	a := 0
	l_p := 0
	for i := 0; i < l; i++ {
		if arr[i] >= 0 {
			arr1 = append(arr1, arr[i])
		}
	}
	fmt.Println(arr1)
	l_p = len(arr1)
	for i :=0; i<l_p; i++ {
	    if arr1[i]%2 == 0 {a = a + arr1[i]}
	}
	fmt.Println(a)
}

/*Сумма четных положительных элементов массива
[-1,2,3,4,-5,-2,-4]*/