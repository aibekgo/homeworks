package main

import "fmt"

func main() {
	var a, b, n int
	fmt.Scanln(&a,&b)
	if a==0 || b==0 {
		fmt.Println("встрелися 0")
	    } else { n = a%b }
	if n == 0 {fmt.Println("делится")
		}  else if n != 0 {fmt.Println("не делится прикинь и вот остаток ", n )
		}  //else {fmt.Println("встрелися 0")}

}
/*Вводятся два целых числа не равных нулю. Проверить делится ли первое на второе.
Вывести на экран сообщение об этом, а также остаток (если он есть)

Если первое число нацело делится на второе, то вывести сообщение об этом.
Иначе вывести сообщение о том, что первое число не делится на второе, найти остаток от деления и также вывести его*/