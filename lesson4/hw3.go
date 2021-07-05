package main

import "fmt"

func main() {
	var s, e int
	fmt.Scan(&s, &e)
	arr := []int{}
	s_arr := []int{}
	avg := 0
	n_arr := 0

	//цикл для заполнения массива
	for i:=s; i<=e; i++ {
		arr = append(arr, i)
	}
	l_arr := len(arr)
	//цикл для нахождения суммы элементов
	for i:=0; i<l_arr; i++ {
		n_arr = n_arr + arr[i]
	}
	//средняя ариф элементов массива
	avg = n_arr / l_arr
	//цикл для выч. элементов>avg
	for i:=0; i<l_arr; i++ {
		if arr[i] < avg {s_arr = append(s_arr, arr[i])}
	}
	fmt.Println("меньше среднего арифметического",s_arr)

}

/*Элементы массива, которые меньше среднего арифметического
вводится start,end
генерация массива от start до end*/