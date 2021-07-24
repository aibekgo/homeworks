package main

import "fmt"

func main() {
	//var result string
	arr := []int{4, 16, 19, 31, 2}
	arrE := []int{}
	arrO := []int{}
	for i := 0; i < len(arr); i++ {
		element := arr[i]
		if element%2 == 0 {
			arrE = append(arrE, element)
		} else {
			arrO = append(arrO, element)
		}
	}
	//result := 'No'
	fmt.Println(arr)
	fmt.Println(arrE)
	fmt.Println(arrO)
	if len(arrE) >= len(arrO) {fmt.Println("Yes")
	} else {fmt.Println("No")

	}
}