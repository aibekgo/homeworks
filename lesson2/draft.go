package main

import (
	"fmt"
	"math"
)

func main() {
	var a float64
	var res string
	fmt.Scanf("%d", &a)
	//res1 = " Целый остаток"
	//n = a/2
	//m = a%2
	s := math.Mod(a, 2)
	res = " Четное число"
	//res2 = " Не четное число"

	fmt.Printf("%s %f", res, s)
}



/*как найти целый остаток от деления?(%)
вводится одно число A
и необходимо проверить является ли оно четным или нет*/