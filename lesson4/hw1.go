package main

import "fmt"

func main() {
	arr := []int{-1, 2, 3, 4, -5, -2, -4}
	arr_p := []int{}
	l := len(arr)
	a := 0
	avg := 0
	for i := 0; i < l; i++ {
		if arr[i] >= 0 {
			arr_p = append(arr_p, arr[i])}
	}
	l_p := len(arr_p)
	for i :=0; i<l_p; i++ {
		a = a + arr_p[i]
	}
	avg = a / l_p
	fmt.Println(arr_p)
	fmt.Println(avg)
}

/*Среднее арифметическое положительных элементов массива
 [-1,2,3,4,-5,-2,-4]
 avg -> сумма_всех/количество*/