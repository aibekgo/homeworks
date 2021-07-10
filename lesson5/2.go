package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	limit := 100
	arr := []int{}
	arr_i := []int{}
	var n, s, e int
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		k := rand.Intn(limit)
		arr = append(arr, k)
	}
	maxi := arr[0]
	maxiIndex := 0
	for i := 0; i < n; i++ {
		if arr[i] > maxi {
			maxi = arr[i]
			maxiIndex = i
		}
	}
	mini := arr[0]
	miniIndex := 0
	for i := 0; i < n; i++ {
		if arr[i] < mini {
			mini = arr[i]
			miniIndex = i
		}
	}

	if miniIndex < maxiIndex {
		s = miniIndex
		e = maxiIndex
	} else {
		s = maxiIndex
		e = miniIndex
	}

	for i := s + 1; i < e; i++ {
		arr_i = append(arr_i, arr[i])
	}

	fmt.Println(arr)
	fmt.Println(mini, miniIndex)
	fmt.Println(maxi, maxiIndex)
	fmt.Println(arr_i)

}
