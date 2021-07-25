package main

import "fmt"

func WarmWeatherCounter(arr []int) int {
	c := 0
	m := 0
	for i := 0; i < len(arr); i++ {
		element := arr[i]
		if element < 1 {
			c = 0
		} else {
			c = c + 1
		}
		if c > m {
			m = c
		}

	}
	return m

}

func main() {
	arr1 := []int{-20, 30, -40, 50, 10, -10}
	arr2 := []int{10, 20, 30, 1, -10, 1, 2, 3}
	arr3 := []int{-10, 0, -10, 0, -10}
	fmt.Println(WarmWeatherCounter(arr1))
	fmt.Println(WarmWeatherCounter(arr2))
	fmt.Println(WarmWeatherCounter(arr3))

}
