package main

import (
	"fmt"
	"math/rand"
	"time"
)

//функция генерации рандомного массива. На выходе дает  массив
func generateOneArray(e, m int) []int {
	arr := []int{}
	rand.Seed(time.Now().UnixNano())
	//цикл для заполнения массива
	for i := 0; i <= m; i++ {
		v := rand.Intn(e)
		arr = append(arr, v)
	}
	return arr
}

//функция разделения массива на два массива. На выходе дает два массива
func generateTwoDimArray(e, m, n int) ([][]int{}) {
        for i := 0; i < n; i++ {
		arr := []int{}
		arr1 := [][]int{}
		twodimarray := [][]int{}
		arr = generateOneArray(e, m)
		arr1 = [i][arr]int{}
		twodimarray = append(twodimarray, arr1)
	}
	return twodimarray
}

func main() {
	arr := [][]int{}
	e := 10
	n := 3
	m := 4
	arr = generateTwoDimArray(e, m, n)
	fmt.Println(arr)

}

/*1) создать фукнции которая генерирует двумерный массив исходя из параметров n,m
n = 3
m = 4
1 2 3 4
4 5 6 3
4 1 2 3*/
