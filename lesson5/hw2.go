package main

import "fmt"

func main() {
	arr := []int{1, 1, 2, 3, 2, 3, 4}
	values := []int{}
	repeats := []int{}
	//uniqe := []int{}

	for i := 0; i < len(arr); i++ {
		n := 0
		t := arr[i]
		for j := 0; j < len(arr); j++ {
		if arr[j] == t {
			values = append(values, t)
			n = n + 1
		}
		}
		repeats = append(repeats, n)
		//if arr[i]==t {uniqe = append(uniqe, t)}

	}
//fmt.Println(uniqe)
fmt.Println(arr)
fmt.Println(repeats)

}




/*в массиве посчитать ск ко встречается каждый элемент в массиве

arr := []int{1, 1, 2, 3, 2, 3, 4}

values :=[1,2,3,4]
repeats:=[2,2,2,1]*/
