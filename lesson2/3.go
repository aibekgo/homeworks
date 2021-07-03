package main

import (
	"fmt"
	"math"
)

func main() {
	var a, n float64
	var res, res1 string
	res1 = " Целый остаток"
	n = a/2
	//m = a%2
	s := math.Mod(a, 2)
	//res = " Четное число"
	//res2 = " Не четное число"
	fmt.Scanf("%d", &a)
	if s=0 {res = " Четное число"
	} else { res = " Не четное число"
	}
   fmt.Printf("%s %d %s %d", res, a, res1, n)
}



/*как найти целый остаток от деления?(%)
вводится одно число A
и необходимо проверить является ли оно четным или нет*/