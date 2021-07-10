package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7}
	arr1 := arr[:4]
	arr2 := arr[5:]
	arr = append(arr1, arr2...)
	fmt.Println(arr)

}