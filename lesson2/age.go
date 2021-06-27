package main

import "fmt"

func main() {
	var age int
	fmt.Scan(&age)
	if age >= 18 { fmt.Println("ты уже взрослый") } else {fmt.Println("ты еще ребенок")}

}
