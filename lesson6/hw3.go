package main

import (
	"fmt"
	"math/rand"
	"time"
)
//функция генерации рандомного массива. На выходе дает длину массива и сам массив
func generateArra(e, l int) (int ,[]int) {
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

//функция разделения массива на два массива. На выходе дает два массива
func splitArray(le int, arr []int) ([]int ,[]int) {
	arr1 := []int{}
	arr2 := []int{}
	if le % 2 != 0 {
		arr = append(arr,0)
		le +=1
	}
	p := le / 2
	arr1 = arr[:p]
	arr2 = arr[p:]
	return arr1, arr2
}


func main() {
	var n, le int
	//введите лимит рандома и длину массива
	fmt.Scan(&n,&le)
	//генерирует массив вызывая функцию generateArra
	le, arr := generateArra(n, le)
	arr1 := []int{}
	arr2 := []int{}
	//передаем длину массива le и массив arr
	arr1, arr2 = splitArray(le, arr)
	fmt.Println(arr)
	fmt.Println(arr1)
	fmt.Println(arr2)
}
/*3)фукнция которая позволяет разделить массив на два одинаковых по длине массива
[1,2,3,4,5]
[1,2,3]
[4,5,0]*/
