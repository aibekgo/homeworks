package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	//библиотека для рандома
	rand.Seed(time.Now().UnixNano())
	//интервал/лимит рандома, здесь от 0 до 20
	limit := 20
	arr := []int{}
	arr_e := []int{}
	arr_o := []int{}
	var n, s, e, s_m, e_m, sum, avg int
	fmt.Scan(&n)
	//заполняем массив
	for i := 0; i < n; i++ {
		k := rand.Intn(limit)
		arr = append(arr, k)
	}
	fmt.Println(arr)
	//задаем старт и энд для забора нашего диапазона
	fmt.Scan(&s, &e)
	arr1 := arr[:s]
	arr2 := arr[e:]
	arr = append(arr1, arr2...)
	fmt.Println(arr)
	//цикл для суммы
	for i := 0; i < len(arr); i++ {
		sum = sum + arr[i]
	}
	//среднее арифметическое
	avg = sum / len(arr)
	//максимальное значение и его индекс
	maxi := arr[0]
	maxiIndex := 0
	for i := 0; i < len(arr); i++ {
		if arr[i] > maxi {
			maxi = arr[i]
			maxiIndex = i
		}
	}
	//минимальное значение и его индекс
	mini := arr[0]
	miniIndex := 0
	for i := 0; i < len(arr); i++ {
		if arr[i] < mini {
			mini = arr[i]
			miniIndex = i
		}
	}
	//разница между макс и мин
	delta := maxi - mini
	//определение старта и энд массива мин_макс
	if miniIndex < maxiIndex {
		s_m = miniIndex
		e_m = maxiIndex
	} else {
		s_m = maxiIndex
		e_m = miniIndex
	}
	//цикл для перебора четных и не четных элементов с диапазона мин макс
	for i := s_m; i <= e_m; i++ {
		if arr[i] % 2 == 0 {
			arr_e = append(arr_e, arr[i])
		} else {
			arr_o = append(arr_o, arr[i])
		}
	}
	//выводим результаты
	fmt.Println("sum =", sum, "avg=", avg, "max=", maxi, "min=", mini, "delta=", delta)
	fmt.Println("четные элементы", arr_e)
	fmt.Println("не четные элементы", arr_o)

}

/* Массив генерируется автоматически
вводится количество элементов в массиве:n
вводится диапозон чисел которые необходимо взять:4 10
найти сумму среднее арифметическое в этом диапозоне
найти максимальное минимальное
найти разность между макс и мин
и среди элементов между макс и мин найти четные и нечетные и поместить в разные массивы*/
