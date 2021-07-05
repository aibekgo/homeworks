package main

import "fmt"

func main() {
	var s, e int
	fmt.Scan(&s, &e)
	arr := []int{}
	n_arr :=0

	//цикл для заполнения массива
	for i := s; i <= e; i++ {
		arr = append(arr, i)
	}
	l_arr := len(arr)
	fmt.Println(arr)
	fmt.Println(l_arr)
	for i:=0; i<l_arr; i++ {
		n_arr = n_arr + arr[i]
	}
	fmt.Println(n_arr)
}