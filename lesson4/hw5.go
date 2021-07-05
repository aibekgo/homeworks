package main

import "fmt"

func main() {
	arr := []int{}
	s := 2
	e := 99
	for i := s; i <= e; i++ {
		arr = append(arr, i)
		//for j := 2; j < 9; j++ {
		//	fmt.Print(arr[i]/j, "\t")
		}
		fmt.Println(arr)
	}

/*Количество кратных чисел от 2 до 99 числам от 2 до 9*/