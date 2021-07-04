package main

import "fmt"

func main() {
	var a, b, c, max, min, avr int
	var res, res1, res2 string
	res = " Максимальное число"
	res1 = " Минимальное число"
	res2 = " Среднее значение"
	fmt.Scanf("%d %d %d", &a, &b, &c)
	if a >= b && a >= c {
		max = a
	} else if b >= a && b >= c {
		max = b
	} else if c >= a && c >= b {
		max = c
	}
	if a <= b && a <= c {
		min = a
	} else if b <= a && b <= c {
		min = b
	} else if c <= a && c <= b {
		min = c
	}

	if a != max && a != min {
		avr = a
	} else if b != max && b != min {
		avr = b
	} else if c != max && c != min {
		avr = c
	} else {
		avr = a
	}
	fmt.Printf("%s -> %d \n", res, max)
	fmt.Printf("%s -> %d \n", res1, min)
	fmt.Printf("%s -> %d \n", res2, avr)
}

/*вводятся a,b,c
найти маскимальное из чисел
найти минимальное из чисел
среднее число это то число которое больше первого но меньше второго*/