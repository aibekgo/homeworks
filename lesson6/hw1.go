package main

import (
	"fmt"
	"math/rand"
	"time"
)

func generateArr(e, l int) (int ,[]int) {
	arr := []int{}
	rand.Seed(time.Now().UnixNano())
	//цикл для заполнения массива
	for i := 0; i <= l; i++ {
		v := rand.Intn(e)
		arr = append(arr, v)
	}
	le := len(arr)
	return le, arr
}

//сумма массива
func GetSumOfArray(a []int) int {
	sumi := 0
	for i := 0; i < len(a); i++ {
		sumi = sumi + a[i]
	}
	return sumi
}

//произведение массива
func GetMultiOfArray(a []int) int {
	multi := 0
	for i := 0; i < len(a); i++ {
		multi = multi * a[i]
	}
	return multi
}

//среднее ариф массива
func GetAvrOfArray(s, b int) int {
	avr := 0
	avr = s / b
	return avr
}
func main() {
	le, k := generateArr(10, 6)
	//s := GetSumOfArray(k)
	fmt.Println(k)
	fmt.Println(GetSumOfArray(k))
	fmt.Println(GetMultiOfArray(k))
	fmt.Println(GetAvrOfArray(GetSumOfArray(k),le))

}

/*1)необходимо написать три фукнции одна считает сумму всех элементов в массиве
произведение,среднее ариф*/
