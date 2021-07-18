package main

import "fmt"

func GetSumOfArr(a []int) int {
	sumi := 0
	for i := 0; i < len(a); i++ {
		sumi = sumi + a[i]
	}
	return sumi
}

func removeFromArr(a []int, value int) []int {
	index := 0
	for i := 0; i < len(a); i++ {
		if i == value {
			index = i
			break
		}
	}
	part1 := a[:index]
	part2 := a[index+1:]
	part1 = append(part1, part2...)
	return part1
}

func main() {
	arr1 := []int{1, 2, 3, 4, 5, 2}
	fmt.Println(arr1)
	arr1 = removeFromArr(arr1, 2)
	fmt.Println(arr1)
}