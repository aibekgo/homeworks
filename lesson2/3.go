package main

import (
	"fmt"
)

func main() {
	var a int
	fmt.Scan(&a)
	c := a % 2
	fmt.Println(c)
	if c == 0 {
		fmt.Println("Четное число")
	} else {
		fmt.Println("Не четное число")
	}
}


/*как найти целый остаток от деления?(%)
вводится одно число A
и необходимо проверить является ли оно четным или нет*/